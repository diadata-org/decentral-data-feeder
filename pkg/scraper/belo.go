package scraper

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/decentral-data-feeder/pkg/models"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
)

const (
	BELO_BASE_URL    = "https://api.belo.app/public/price"
	BELO_CONFIG_PATH = "belo.json"
)

type BeloResponse struct {
	Ask      string `json:"ask"`
	Bid      string `json:"bid"`
	BaseID   string `json:"baseId"`
	QuoteID  string `json:"quoteId"`
	PairCode string `json:"pairCode"`
	Spread   string `json:"spread"`
}

type BeloQuote struct {
	Ask      float64 `json:"ask"`
	Bid      float64 `json:"bid"`
	BaseID   int     `json:"baseId"`
	QuoteID  int     `json:"quoteId"`
	PairCode string  `json:"pairCode"`
	Spread   float64 `json:"spread"`
}

type BeloScraper struct {
	pairs              []string
	updateTicker       *time.Ticker
	configUpdateTicker *time.Ticker
	branchMarketConfig string
	dataChannel        chan []byte
	updateDoneChannel  chan bool
	pairSeparator      string
}

func NewBeloScraper() *BeloScraper {
	updateSecs, err := strconv.ParseInt(utils.Getenv("BELO_UPDATE_SECONDS", "30"), 10, 64)
	if err != nil {
		log.Error("Parse BELO_UPDATE_SECONDS:", err)
		updateSecs = 30
	}
	configUpdateSeconds, err := strconv.Atoi(utils.Getenv("BELO_CONFIG_UPDATE_SECONDS", "86400"))
	if err != nil {
		log.Errorf("parse BELO_CONFIG_UPDATE_SECONDS: %v", err)
		configUpdateSeconds = 86400
	}

	scraper := &BeloScraper{
		updateTicker:       time.NewTicker(time.Duration(updateSecs) * time.Second),
		configUpdateTicker: time.NewTicker(time.Duration(configUpdateSeconds) * time.Second),
		branchMarketConfig: utils.Getenv("BELO_BRANCH_MARKET_CONFIG", ""),
		pairSeparator:      "/",
	}
	scraper.dataChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)

	err = scraper.updateConfig(BELO_CONFIG_PATH)
	if err != nil {
		log.Fatal("Could not load configuration file: ", err)
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *BeloScraper) mainLoop() {

	// Periodically fetch configuration.
	go func() {
		for range scraper.configUpdateTicker.C {
			err := scraper.updateConfig(BELO_CONFIG_PATH)
			if err != nil {
				log.Errorf("updateConfig %v", err)
			}
		}
	}()

	// Initial run
	err := scraper.UpdatePrices(BELO_BASE_URL)
	if err != nil {
		log.Error("BELO scraper initial update error:", err)
	}
	scraper.updateDoneChannel <- true

	for range scraper.updateTicker.C {
		err := scraper.UpdatePrices(BELO_BASE_URL)
		if err != nil {
			log.Error("BELO scraper update error:", err)
			continue
		}
		scraper.updateDoneChannel <- true
	}
}

func (scraper *BeloScraper) UpdatePrices(url string) error {

	log.Info("update prices for Belo.............")
	br, err := scraper.getPricesFromAPI(url)
	if err != nil {
		return err
	}

	beloQuotes, err := scraper.normalizeBeloResponse(br)
	if err != nil {
		return err
	}

	for _, quote := range beloQuotes {

		b, err := json.Marshal(quote)
		if err != nil {
			log.Error("marshal BELO data:", err)
			continue
		}
		scraper.dataChannel <- b
	}
	return nil
}

// getPricesFromAPI fetches price endpoint and parses JSON.
// Also discards pairs not contained in the configuration file.
func (scraper *BeloScraper) getPricesFromAPI(url string) (br []BeloResponse, err error) {
	var response []byte

	response, _, err = utils.GetRequest(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &br)
	return
}

// Cast API response to final type, restrict to pairs from configuration file.
// Furthermore deduplicate and increment to ARS/USDT.
func (scraper *BeloScraper) normalizeBeloResponse(br []BeloResponse) ([]BeloQuote, error) {
	var bq []BeloQuote

	keyMap := make(map[string]struct{})
	for _, r := range br {
		if !slices.Contains(scraper.pairs, r.PairCode) {
			continue
		}
		if _, ok := keyMap[r.PairCode]; !ok {
			q, err := castResponse(r)
			if err != nil {
				return bq, err
			}
			qIncrement, err := scraper.incrementQuote(q)
			if err != nil {
				log.Error(err)
				continue
			}
			bq = append(bq, qIncrement)
			keyMap[r.PairCode] = struct{}{}
		}
	}
	return bq, nil
}

func (scraper *BeloScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}
func (scraper *BeloScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}
func (scraper *BeloScraper) Close() error {
	return nil
}

func (scraper *BeloScraper) updateConfig(filePath string) error {

	beloConfig, err := models.GetBeloConfig(filePath, scraper.branchMarketConfig)
	if err != nil {
		return err
	}

	scraper.pairs = beloConfig.Pairs
	return nil
}

func (scraper *BeloScraper) incrementQuote(bq BeloQuote) (BeloQuote, error) {
	incrementQuote := bq
	symbols := strings.Split(bq.PairCode, scraper.pairSeparator)
	if len(symbols) != 2 {
		return BeloQuote{}, fmt.Errorf("wrong pair separator")
	}
	incrementQuote.PairCode = symbols[1] + scraper.pairSeparator + symbols[0]
	if bq.Ask == 0 {
		return BeloQuote{}, fmt.Errorf("Ask price is 0 for %s", incrementQuote.PairCode)
	}
	incrementQuote.Ask = 1 / bq.Ask
	if bq.Bid == 0 {
		return BeloQuote{}, fmt.Errorf("Bid price is 0 for %s", incrementQuote.PairCode)
	}
	incrementQuote.Bid = 1 / bq.Bid

	incrementQuote.Spread = (incrementQuote.Ask - incrementQuote.Bid) / ((incrementQuote.Ask + incrementQuote.Bid) / 2) * 100

	return incrementQuote, nil

}

func castResponse(br BeloResponse) (bq BeloQuote, err error) {
	bq.PairCode = br.PairCode
	bq.Ask, err = strconv.ParseFloat(br.Ask, 64)
	if err != nil {
		return
	}
	bq.Bid, err = strconv.ParseFloat(br.Bid, 64)
	if err != nil {
		return
	}
	bq.Spread, err = strconv.ParseFloat(br.Spread, 64)
	if err != nil {
		return
	}
	bq.QuoteID, err = strconv.Atoi(br.QuoteID)
	if err != nil {
		return
	}
	bq.BaseID, err = strconv.Atoi(br.BaseID)
	if err != nil {
		return
	}
	return
}
