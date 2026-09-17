package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/decentral-data-feeder/pkg/models"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/tidwall/gjson"
)

const (
	DENARIO_PRICE_URL   = "https://price.denario.swiss/api"
	DENARIO_RESERVE_URL = "https://reserve.denario.swiss/api"
	DENARIO_CONFIG_PATH = "denario.json"

	DENARIO_PRICE_FIELD   = "priceAsk"
	DENARIO_RESERVE_FIELD = "amountOunces"

	denarioRequestTimeout = 10 * time.Second

	DenarioPrice   dataType = "Price"
	DenarioReserve dataType = "ProofOfReserve"
)

type DenarioQuote struct {
	Key   string   `json:"Key"`
	Value float64  `json:"Value"`
	Type  dataType `json:"Type"`
}

type DenarioScraper struct {
	priceURL           string
	reserveURL         string
	authKey            string
	httpClient         *http.Client
	assetsMu           sync.RWMutex
	prices             []string
	reserves           []string
	updateTicker       *time.Ticker
	configUpdateTicker *time.Ticker
	branchMarketConfig string
	dataChannel        chan []byte
	updateDoneChannel  chan bool
}

func NewDenarioScraper() *DenarioScraper {
	updateSecs, err := strconv.ParseInt(utils.Getenv("DENARIO_UPDATE_SECONDS", "60"), 10, 64)
	if err != nil {
		log.Errorf("parse DENARIO_UPDATE_SECONDS: %v", err)
		updateSecs = 60
	}
	configUpdateSeconds, err := strconv.Atoi(utils.Getenv("DENARIO_CONFIG_UPDATE_SECONDS", "86400"))
	if err != nil {
		log.Errorf("parse DENARIO_CONFIG_UPDATE_SECONDS: %v", err)
		configUpdateSeconds = 86400
	}

	scraper := &DenarioScraper{
		priceURL:           DENARIO_PRICE_URL,
		reserveURL:         DENARIO_RESERVE_URL,
		authKey:            utils.Getenv("DENARIO_AUTH_KEY", ""),
		httpClient:         &http.Client{Timeout: denarioRequestTimeout},
		updateTicker:       time.NewTicker(time.Duration(updateSecs) * time.Second),
		configUpdateTicker: time.NewTicker(time.Duration(configUpdateSeconds) * time.Second),
		branchMarketConfig: utils.Getenv("DENARIO_BRANCH_MARKET_CONFIG", ""),
	}
	scraper.dataChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)

	if scraper.authKey == "" {
		log.Fatal("DENARIO_AUTH_KEY is not set.")
	}

	err = scraper.updateConfig(DENARIO_CONFIG_PATH)
	if err != nil {
		log.Fatal("Could not load configuration file: ", err)
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *DenarioScraper) mainLoop() {

	// Periodically fetch configuration.
	go func() {
		for range scraper.configUpdateTicker.C {
			err := scraper.updateConfig(DENARIO_CONFIG_PATH)
			if err != nil {
				log.Errorf("updateConfig %v", err)
			}
		}
	}()

	// Initial run
	err := scraper.UpdateValues()
	if err != nil {
		log.Error("Denario scraper initial update error: ", err)
	} else {
		scraper.updateDoneChannel <- true
	}

	for range scraper.updateTicker.C {
		err := scraper.UpdateValues()
		if err != nil {
			log.Error("Denario scraper update error: ", err)
			continue
		}
		scraper.updateDoneChannel <- true
	}
}

// UpdateValues fetches all prices and reserves from the configuration file and sends them to the data channel.
// Assets that fail are logged and skipped. Returns an error if no value was sent at all, so that no empty
// oracle update is triggered.
func (scraper *DenarioScraper) UpdateValues() error {

	log.Info("update values for Denario.............")
	prices, reserves := scraper.getAssets()

	sent := scraper.sendValues(scraper.priceURL, prices, DENARIO_PRICE_FIELD, DenarioPrice)
	sent += scraper.sendValues(scraper.reserveURL, reserves, DENARIO_RESERVE_FIELD, DenarioReserve)
	if sent == 0 {
		return fmt.Errorf("no values for %d prices and %d reserves", len(prices), len(reserves))
	}
	return nil
}

func (scraper *DenarioScraper) sendValues(baseURL string, assets []string, field string, t dataType) (sent int) {
	for _, asset := range assets {
		value, err := scraper.getValue(baseURL+"/"+asset, field)
		if err != nil {
			log.Errorf("get Denario %s for %s: %v", field, asset, err)
			continue
		}

		b, err := json.Marshal(DenarioQuote{Key: asset, Value: value, Type: t})
		if err != nil {
			log.Error("marshal Denario data: ", err)
			continue
		}
		scraper.dataChannel <- b
		sent++
	}
	return
}

func (scraper *DenarioScraper) getValue(url string, field string) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("X-AUTH-TOKEN", scraper.authKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := scraper.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP Response Error %d: %s", resp.StatusCode, body)
	}

	result := gjson.GetBytes(body, field)
	if !result.Exists() {
		return 0, fmt.Errorf("no %s in response: %s", field, body)
	}
	if result.Float() <= 0 {
		return 0, fmt.Errorf("non-positive %s: %s", field, result.Raw)
	}
	return result.Float(), nil
}

func (scraper *DenarioScraper) updateConfig(filePath string) error {

	denarioConfig, err := models.GetDenarioConfig(filePath, scraper.branchMarketConfig)
	if err != nil {
		return err
	}

	// Runs concurrently to UpdateValues.
	scraper.assetsMu.Lock()
	defer scraper.assetsMu.Unlock()
	scraper.prices = denarioConfig.Prices
	scraper.reserves = denarioConfig.Reserves
	return nil
}

func (scraper *DenarioScraper) getAssets() (prices []string, reserves []string) {
	scraper.assetsMu.RLock()
	defer scraper.assetsMu.RUnlock()
	return scraper.prices, scraper.reserves
}

func (scraper *DenarioScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}
func (scraper *DenarioScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}
func (scraper *DenarioScraper) Close() error {
	return nil
}
