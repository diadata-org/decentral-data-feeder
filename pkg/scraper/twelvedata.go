package scraper

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
)

const (
	TWELVEDATA = "TwelveData"
)

var (
	twelvedataUpdateSeconds int64
	twelvedataApiBaseString = "https://api.twelvedata.com/"
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

type TwelvedataQuote struct {
	Symbol   string
	Name     string
	Price    float64
	Time     time.Time
	DataType string
	Source   string
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
		twelvedataFXTickers:    strings.Split(utils.Getenv("FX_TICKERS", "USD/EUR,USD/GBP,USD/CHF,USD/JPY"), ","),
		twelvedataCommodities:  strings.Split(utils.Getenv("COMMODITIES", ""), ","),
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

	marketTime, err := isMarketTime()
	if err != nil {
		log.Error("isMarketOpen: ", err)
	}
	if marketTime {
		// Check if markets are open.
		quote, err := scraper.getTwelveQuote("AAPL")
		if err != nil {
			log.Error("getTwelveQuote: ", err)
		}
		scraper.twelvedataStockMarketOpen = quote.MarketOpen
	} else {
		scraper.twelvedataStockMarketOpen = false
	}

	log.Infof("Executing stock data update for %v symbols", len(scraper.twelvedataStockSymbols))
	if scraper.twelvedataStockMarketOpen {
		for _, symbol := range scraper.twelvedataStockSymbols {
			if symbol == "" {
				continue
			}
			quotation, err := scraper.getTwelveStockPrice(symbol)
			if err != nil {
				log.Error("getTwelveStockPrice: ", err)
			}
			price, err := strconv.ParseFloat(quotation.Price, 64)
			if err != nil {
				log.Error("Parse Float for: ", symbol)
			}

			quote := TwelvedataQuote{
				Symbol:   symbol,
				Price:    price,
				Time:     time.Now(),
				Source:   TWELVEDATA,
				DataType: "Stock",
			}

			quoteBytes, err := json.Marshal(quote)
			if err != nil {
				log.Error("marshal stock data: ", err)
			}
			scraper.dataChannel <- quoteBytes
		}
	}

	log.Infof("Executing fx data update for %v symbols", len(scraper.twelvedataFXTickers))
	for _, ticker := range scraper.twelvedataFXTickers {
		if ticker == "" {
			continue
		}
		quotation, err := scraper.getTwelveFXData(ticker)
		if err != nil {
			log.Error("getTwelveFXData: ", err)
		}

		quote := TwelvedataQuote{
			Symbol:   ticker,
			Price:    quotation.Rate,
			Time:     time.Unix(quotation.Timestamp, 0),
			Source:   TWELVEDATA,
			DataType: "FX",
		}
		quoteBytes, err := json.Marshal(quote)
		if err != nil {
			log.Error("marshal fx data: ", err)
		}

		scraper.dataChannel <- quoteBytes
	}

	log.Infof("Executing commodities data update for %v symbols.", len(scraper.twelvedataCommodities))
	for _, ticker := range scraper.twelvedataCommodities {
		if ticker == "" {
			continue
		}
		quotation, err := scraper.getTwelveQuote(ticker)
		if err != nil {
			log.Error("getTwelveFXData: ", err)
		}

		price, err := strconv.ParseFloat(quotation.Price, 64)
		if err != nil {
			log.Errorf("parse price for %s", quotation.Symbol)
		}

		quote := TwelvedataQuote{
			Symbol:   ticker,
			Name:     quotation.Name,
			Price:    price,
			Time:     time.Unix(quotation.Timestamp, 0),
			Source:   TWELVEDATA,
			DataType: "Commodity",
		}
		quoteBytes, err := json.Marshal(quote)
		if err != nil {
			log.Error("marshal commodities data: ", err)
		}
		scraper.dataChannel <- quoteBytes
	}

	log.Infof("Executing ETF data update for %v symbols", len(scraper.twelvedataETFs))
	for _, ticker := range scraper.twelvedataETFs {
		if ticker == "" {
			continue
		}
		quotation, err := scraper.getTwelveQuote(ticker)
		if err != nil {
			log.Error("getTwelveFXData: ", err)
		}

		price, err := strconv.ParseFloat(quotation.Price, 64)
		if err != nil {
			log.Errorf("parse price for %s", quotation.Symbol)
		}

		quote := TwelvedataQuote{
			Symbol:   ticker,
			Name:     quotation.Name,
			Price:    price,
			Time:     time.Unix(quotation.Timestamp, 0),
			Source:   TWELVEDATA,
			DataType: "ETF",
		}
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

	apiURL := twelvedataApiBaseString + "quote?symbol=" + symbol + "&apikey=" + scraper.apiKey
	response, _, err = utils.GetRequest(apiURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &commodity)
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

func isMarketTime() (bool, error) {
	// Load the America/New_York location (EST/EDT with daylight saving)
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return false, err
	}
	now := time.Now().In(loc)

	openTime := time.Date(now.Year(), now.Month(), now.Day(), 8, 50, 0, 0, loc)
	closeTime := time.Date(now.Year(), now.Month(), now.Day(), 16, 10, 0, 0, loc)

	return now.After(openTime) && now.Before(closeTime), nil
}
