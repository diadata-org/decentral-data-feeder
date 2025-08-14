package scraper

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	calendar "github.com/scmhub/calendar"
)

const (
	TWELVEDATA                  = "TwelveData"
	DIA_TICKER_SEPARATOR        = "-"
	TWELVEDATA_TICKER_SEPARATOR = "/"

	NyseOpen    dataType = "NyseOpen"
	Equities    dataType = "Equities"
	Fiat        dataType = "Fiat"
	Commodities dataType = "Commodities"
	ETF         dataType = "ETF"
)

var (
	twelvedataUpdateSeconds int64
	twelvedataApiBaseString = "https://api.twelvedata.com/"
	diadataApiBaseString    = "https://api.diadata.org/v1/"
)

type twelvedataAPIFXResponse struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
}

type twelvedataAPIStockResponse struct {
	Price string `json:"price"`
}

type twelvedataQuoteResponse struct {
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	Timestamp  int64  `json:"last_quote_at"`
	Price      string `json:"close"`
	MarketOpen bool   `json:"is_market_open"`
}

type dataType string

type TwelvedataQuote struct {
	Symbol      string    `json:"Ticker"`
	Name        string    `json:"Name"`
	Price       float64   `json:"Price"`
	Time        time.Time `json:"Timestamp"`
	NYSEHoliday bool      `json:"Nyse_Holiday"`
	NYSEOpen    bool      `json:"Nyse_Open"`
	Type        dataType
	Source      string
}

type TwelvedataScraper struct {
	dataChannel               chan []byte
	updateDoneChannel         chan bool
	ticker                    *time.Ticker
	twelvedataStockSymbols    []string
	twelvedataStockMarketOpen bool
	twelvedataFXTickers       []string
	twelvedataCommodities     []string
	twelvedataETFs            []string
	apiKey                    string
}

func init() {
	var err error
	twelvedataUpdateSeconds, err = strconv.ParseInt(utils.Getenv("UPDATE_SECONDS", "30"), 10, 64)
	if err != nil {
		log.Error("Parse UPDATE_SECONDS: ", err)
	}
}

func NewTwelvedataScraper() *TwelvedataScraper {

	s := &TwelvedataScraper{
		ticker:                 time.NewTicker(time.Duration(twelvedataUpdateSeconds) * time.Second),
		twelvedataStockSymbols: strings.Split(utils.Getenv("STOCK_SYMBOLS", "AAPL,MSFT,NVDA,AMZN,GOOG,META,TSLA,WMT,JPM,V"), ","),
		twelvedataFXTickers:    strings.Split(utils.Getenv("FX_TICKERS", "USD-EUR,USD-GBP,USD-CHF,USD-JPY"), ","),
		twelvedataCommodities:  strings.Split(utils.Getenv("COMMODITIES", "XAU-USD"), ","),
		twelvedataETFs:         strings.Split(utils.Getenv("ETF", ""), ","),
		apiKey:                 utils.Getenv("TWELVEDATA_API_KEY", ""),
	}

	s.dataChannel = make(chan []byte)
	s.updateDoneChannel = make(chan bool)

	go s.mainLoop()

	return s

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *TwelvedataScraper) mainLoop() {

	// Initial run.
	err := scraper.UpdateQuotations()
	if err != nil {
		log.Error(err)
	}
	scraper.updateDoneChannel <- true

	// Update every @twelvedataUpdateSeconds.
	for range scraper.ticker.C {
		err := scraper.UpdateQuotations()
		if err != nil {
			log.Error(err)
			continue
		}
		scraper.updateDoneChannel <- true
	}

}

// Update retrieves new coin information from the twelvedata API and stores it to influx
func (scraper *TwelvedataScraper) UpdateQuotations() error {

	nyseHoliday, marketTime, err := isMarketTime()
	if err != nil {
		log.Error("isMarketOpen: ", err)
	}
	if nyseHoliday {
		scraper.twelvedataStockMarketOpen = false
	}

	// Check if markets are open.
	if !nyseHoliday && marketTime {
		quote, err := scraper.getTwelveQuote("AAPL")
		if err != nil {
			log.Error("getTwelveQuote: ", err)
		}
		scraper.twelvedataStockMarketOpen = quote.MarketOpen
	} else {
		scraper.twelvedataStockMarketOpen = false
	}

	// Send information on nyse open or closed.
	var quoteOpen TwelvedataQuote
	quoteOpen.Type = NyseOpen
	quoteOpen.NYSEHoliday = nyseHoliday
	quoteOpen.NYSEOpen = marketTime
	quoteOpen.Time = time.Now()
	quoteBytes, err := json.Marshal(quoteOpen)
	if err != nil {
		log.Error("marshal stock data: ", err)
	}
	scraper.dataChannel <- quoteBytes

	if scraper.twelvedataStockMarketOpen {
		log.Infof("Executing stock data update for %v symbols", getNumSymbols(scraper.twelvedataStockSymbols))
		for _, symbol := range scraper.twelvedataStockSymbols {
			if symbol == "" {
				continue
			}

			quote, err := getRwaQuoteFromDia(symbol, "Equities")
			// Fetch data from twelvedata API in case diadata API fails and a twelvedata api key was provided.
			if err != nil && scraper.apiKey != "" {
				log.Warn("")
				quotation, err := scraper.getTwelveStockPrice(symbol)
				if err != nil {
					log.Error("getTwelveStockPrice: ", err)
				}
				price, err := strconv.ParseFloat(quotation.Price, 64)
				if err != nil {
					log.Error("Parse Float for: ", symbol)
				}
				quote.Symbol = symbol
				quote.Price = price
				quote.Time = time.Now()
			}
			quote.Source = TWELVEDATA
			quote.Type = Equities

			quoteBytes, err := json.Marshal(quote)
			if err != nil {
				log.Error("marshal stock data: ", err)
			}
			scraper.dataChannel <- quoteBytes
		}
	} else {
		log.Info("Stock market closed. No updates.")
	}

	log.Infof("Executing fx data update for %v symbols", getNumSymbols(scraper.twelvedataFXTickers))
	for _, ticker := range scraper.twelvedataFXTickers {
		if ticker == "" {
			continue
		}

		quote, err := getRwaQuoteFromDia(ticker, "Fiat")
		if err != nil && scraper.apiKey != "" {
			log.Warnf("quote for %s not available in diadata api. Switch to twelvdata api.", ticker)
			quotation, err := scraper.getTwelveFXData(ticker)
			if err != nil {
				log.Error("getTwelveFXData: ", err)
			}
			quote.Symbol = ticker
			quote.Price = quotation.Rate
			quote.Time = time.Unix(quotation.Timestamp, 0)
		} else if err != nil {
			continue
		}

		quote.Source = TWELVEDATA
		quote.Type = Fiat

		quoteBytes, err := json.Marshal(quote)
		if err != nil {
			log.Error("marshal fx data: ", err)
		}

		scraper.dataChannel <- quoteBytes
	}

	log.Infof("Executing commodities data update for %v symbols.", getNumSymbols(scraper.twelvedataCommodities))
	for _, ticker := range scraper.twelvedataCommodities {
		if ticker == "" {
			continue
		}

		quote, err := getRwaQuoteFromDia(ticker, "Commodities")
		if err != nil && scraper.apiKey != "" {
			log.Warnf("quote for %s not available in diadata api. Switch to twelvdata api.", ticker)
			quotation, err := scraper.getTwelveQuote(ticker)
			if err != nil {
				log.Error("getTwelveCommoditiesData: ", err)
			}

			price, err := strconv.ParseFloat(quotation.Price, 64)
			if err != nil {
				log.Errorf("parse price for %s", quotation.Symbol)
			}
			quote = TwelvedataQuote{
				Symbol: ticker,
				Price:  price,
				Time:   time.Unix(quotation.Timestamp, 0),
				Name:   quotation.Name,
			}
		} else if err != nil {
			continue
		}

		quote.Source = TWELVEDATA
		quote.Type = Commodities

		quoteBytes, err := json.Marshal(quote)
		if err != nil {
			log.Error("marshal commodities data: ", err)
		}
		scraper.dataChannel <- quoteBytes
	}

	log.Infof("Executing ETF data update for %v symbols", getNumSymbols(scraper.twelvedataETFs))
	for _, ticker := range scraper.twelvedataETFs {
		if ticker == "" {
			continue
		}

		quote, err := getRwaQuoteFromDia(ticker, "ETF")
		if err != nil && scraper.apiKey != "" {
			log.Warnf("quote for %s not available in diadata api. Switch to twelvdata api.", ticker)
			quotation, err := scraper.getTwelveQuote(ticker)
			if err != nil {
				log.Error("getTwelveETFData: ", err)
			}

			price, err := strconv.ParseFloat(quotation.Price, 64)
			if err != nil {
				log.Errorf("parse price for %s", quotation.Symbol)
			}
			quote = TwelvedataQuote{
				Symbol: ticker,
				Name:   quotation.Name,
				Price:  price,
				Time:   time.Unix(quotation.Timestamp, 0),
			}
		} else if err != nil {
			continue
		}

		quote.Source = TWELVEDATA
		quote.Type = ETF

		quoteBytes, err := json.Marshal(quote)
		if err != nil {
			log.Error("marshal etf data: ", err)
		}
		scraper.dataChannel <- quoteBytes
	}

	return nil

}

func (scraper *TwelvedataScraper) getTwelveFXData(symbol string) (fxRate twelvedataAPIFXResponse, err error) {
	var response []byte

	// Twelvedata api takes "/" as separator instead of "-".
	symbols := strings.Split(symbol, DIA_TICKER_SEPARATOR)
	if len(symbols) == 2 {
		symbol = symbols[0] + TWELVEDATA_TICKER_SEPARATOR + symbols[1]
	}

	apiURL := twelvedataApiBaseString + "exchange_rate?symbol=" + symbol + "&apikey=" + scraper.apiKey
	response, _, err = utils.GetRequest(apiURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &fxRate)
	return
}

func (scraper *TwelvedataScraper) getTwelveStockPrice(symbol string) (stockPrice twelvedataAPIStockResponse, err error) {
	var response []byte

	apiURL := twelvedataApiBaseString + "price?symbol=" + symbol + "&apikey=" + scraper.apiKey
	response, _, err = utils.GetRequest(apiURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &stockPrice)
	return
}

func (scraper *TwelvedataScraper) getTwelveQuote(symbol string) (commodity twelvedataQuoteResponse, err error) {
	var response []byte

	symbols := strings.Split(symbol, DIA_TICKER_SEPARATOR)
	if len(symbols) == 2 {
		symbol = symbols[0] + TWELVEDATA_TICKER_SEPARATOR + symbols[1]
	}

	apiURL := twelvedataApiBaseString + "quote?symbol=" + symbol + "&apikey=" + scraper.apiKey
	response, _, err = utils.GetRequest(apiURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &commodity)
	return
}

func getRwaQuoteFromDia(symbol string, assetType string) (tq TwelvedataQuote, err error) {

	symbols := strings.Split(symbol, TWELVEDATA_TICKER_SEPARATOR)
	if len(symbols) == 2 {
		symbol = symbols[0] + DIA_TICKER_SEPARATOR + symbols[1]
	}

	url := diadataApiBaseString + "rwa/" + assetType + "/" + symbol
	response, _, err := utils.GetRequest(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &tq)

	// Notation in oracle should be USD/GBP instead of USD-GBP
	if (assetType == string(Fiat) || assetType == string(Commodities)) && len(strings.Split(symbol, DIA_TICKER_SEPARATOR)) == 2 {
		tq.Symbol = strings.Split(symbol, DIA_TICKER_SEPARATOR)[0] + TWELVEDATA_TICKER_SEPARATOR + strings.Split(symbol, DIA_TICKER_SEPARATOR)[1]
	}
	return

}

func (scraper *TwelvedataScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}

func (scraper *TwelvedataScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}

func (scraper *TwelvedataScraper) Close() error {
	return nil
}

func isMarketTime() (holiday bool, marketTime bool, err error) {
	// New Yorck stock exchange calendar.
	nyse := calendar.XNYS()

	if nyse.IsHoliday(time.Now()) {
		holiday = true
		return
	}

	// Load the America/New_York location (EST/EDT with daylight saving)
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return
	}
	now := time.Now().In(loc)

	openTime := time.Date(now.Year(), now.Month(), now.Day(), 8, 50, 0, 0, loc)
	closeTime := time.Date(now.Year(), now.Month(), now.Day(), 16, 10, 0, 0, loc)

	return false, now.After(openTime) && now.Before(closeTime), nil
}

func getNumSymbols(symbols []string) int {
	l := len(symbols)
	if l == 0 {
		return 0
	}
	if l == 1 && symbols[0] == "" {
		return 0
	}
	return l
}
