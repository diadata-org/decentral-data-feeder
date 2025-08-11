package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/tidwall/gjson"
)

const COINMARKETCAP = "Coinmarketcap"
const COINMARKETCAP_API_BASE_URL = "https://pro-api.coinmarketcap.com/v1"

type CMCScraper struct {
	apiKey            string
	assetNames        string
	apiValuePath      string
	dataChannel       chan []byte
	updateDoneChannel chan bool
}

type CMCResponse struct {
	Symbol    string
	Value     float64
	Timestamp time.Time
}

func NewCMCScraper() *CMCScraper {
	var scraper CMCScraper
	scraper.apiKey = utils.Getenv("CMC_PRO_API_KEY", "")
	scraper.assetNames = utils.Getenv("CMC_ASSET_NAMES", "1,2,3")
	scraper.apiValuePath = utils.Getenv("CMC_API_VALUE_PATH", "quote.USD.price")
	// scraper.apiValuePath = utils.Getenv("CMC_API_VALUE_PATH", "circulating_supply")
	scraper.dataChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)

	go scraper.run()
	return &scraper
}

func (scraper *CMCScraper) run() {
	tick := time.NewTicker(time.Duration(120 * time.Second))
	for range tick.C {
		scraper.getCmcPrices(scraper.assetNames, scraper.apiKey)
	}
}

func (scraper *CMCScraper) getCmcPrices(assetNames, apiKey string) error {
	url := COINMARKETCAP_API_BASE_URL + "/cryptocurrency/quotes/latest?id=" + assetNames + "&CMC_PRO_API_KEY=" + scraper.apiKey
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("CMC API call with return code %d", response.StatusCode)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	data := gjson.GetBytes(contents, "data")

	data.ForEach(func(key, value gjson.Result) bool {
		var resp CMCResponse
		resp.Symbol = gjson.Get(value.String(), "symbol").String()
		resp.Value = gjson.Get(value.String(), scraper.apiValuePath).Float()
		resp.Timestamp, err = time.Parse(time.RFC3339, gjson.Get(value.String(), "quote.USD.last_updated").String())
		if err != nil {
			log.Error("CMC - Parse timestamp: ", err)
		}
		data, err := json.Marshal(resp)
		if err != nil {
			log.Error("Marshal response: ", err)
		}

		scraper.DataChannel() <- data
		return true
	})

	return nil
}

func (scraper *CMCScraper) Close() error {
	log.Info("TO DO.")
	return nil
}

func (scraper *CMCScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}

func (scraper *CMCScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}
