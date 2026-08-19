package scraper

import (
	"encoding/json"
	"slices"
	"strconv"
	"time"

	"github.com/diadata-org/decentral-data-feeder/pkg/models"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
)

const (
	XLSD_URL         = "https://api.diadata.org/xlsd"
	XLSD_CONFIG_PATH = "xlsd.json"
)

type VToken struct {
	Symbol          string          `json:"Token"`
	FairPrice       float64         `json:"FairPrice"`
	CollateralRatio collateralRatio `json:"Collateralratio"`
	BaseAssetSymbol string          `json:"BaseAssetSymbol"`
	BaseAssetPrice  float64         `json:"BaseAssetPrice"`
	Issuer          string          `json:"Issuer"`
	Timestamp       int64           `json:"TimeStamp"`
}

type collateralRatio struct {
	IssuedToken float64 `json:"IssuedToken"`
	LockedToken float64 `json:"LockedToken"`
	Ratio       float64 `json:"Ratio"`
}

type XLSDScraper struct {
	tokens             []string
	updateTicker       *time.Ticker
	configUpdateTicker *time.Ticker
	branchMarketConfig string
	dataChannel        chan []byte
	updateDoneChannel  chan bool
	pairSeparator      string
}

func NewXLSDScraper() *XLSDScraper {
	updateSecs, err := strconv.ParseInt(utils.Getenv("XLSD_UPDATE_SECONDS", "30"), 10, 64)
	if err != nil {
		log.Error("Parse XLSD_UPDATE_SECONDS:", err)
		updateSecs = 30
	}
	configUpdateSeconds, err := strconv.Atoi(utils.Getenv("XLSD_CONFIG_UPDATE_SECONDS", "86400"))
	if err != nil {
		log.Errorf("parse XLSD_CONFIG_UPDATE_SECONDS: %v", err)
		configUpdateSeconds = 86400
	}

	scraper := &XLSDScraper{
		updateTicker:       time.NewTicker(time.Duration(updateSecs) * time.Second),
		configUpdateTicker: time.NewTicker(time.Duration(configUpdateSeconds) * time.Second),
		branchMarketConfig: utils.Getenv("XLSD_BRANCH_MARKET_CONFIG", ""),
		pairSeparator:      "/",
	}
	scraper.dataChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)

	err = scraper.updateConfig(XLSD_CONFIG_PATH)
	if err != nil {
		log.Fatal("Could not load configuration file: ", err)
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *XLSDScraper) mainLoop() {

	// Periodically fetch configuration.
	go func() {
		for range scraper.configUpdateTicker.C {
			err := scraper.updateConfig(XLSD_CONFIG_PATH)
			if err != nil {
				log.Errorf("updateConfig %v", err)
			}
		}
	}()

	// Initial run
	err := scraper.UpdatePrices(XLSD_URL)
	if err != nil {
		log.Error("XLSD scraper initial update error:", err)
	}
	scraper.updateDoneChannel <- true

	for range scraper.updateTicker.C {
		err := scraper.UpdatePrices(XLSD_URL)
		if err != nil {
			log.Error("XLSD scraper update error:", err)
			continue
		}
		scraper.updateDoneChannel <- true
	}
}

func (scraper *XLSDScraper) UpdatePrices(url string) error {

	log.Info("update prices for XLSD.............")
	allTokens, err := getXLSD()
	if err != nil {
		return err
	}
	tokens := scraper.restrictToConfigTokens(allTokens)

	for _, token := range tokens {

		b, err := json.Marshal(token)
		if err != nil {
			log.Error("marshal XLSD data:", err)
			continue
		}
		scraper.dataChannel <- b
	}
	return nil
}

func getXLSD() (resp []VToken, err error) {
	data, _, err := utils.GetRequest(XLSD_URL)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)

	return
}

func (scraper *XLSDScraper) updateConfig(filePath string) error {

	xlsdConfig, err := models.GetXLSDConfig(filePath, scraper.branchMarketConfig)
	if err != nil {
		return err
	}

	scraper.tokens = xlsdConfig.Tokens
	return nil
}

func (scraper *XLSDScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}
func (scraper *XLSDScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}
func (scraper *XLSDScraper) Close() error {
	return nil
}

// Restrict all tokens from API response to subset of tokens as given by config file.
func (scraper *XLSDScraper) restrictToConfigTokens(allTokens []VToken) (tokens []VToken) {
	for _, token := range allTokens {
		if slices.Contains(scraper.tokens, token.Symbol) {
			tokens = append(tokens, token)
		}
	}
	return
}
