package scraper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	models "github.com/diadata-org/decentral-data-feeder/pkg/models"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/tidwall/gjson"
)

const COINMARKETCAP = "Coinmarketcap"
const COINMARKETCAP_API_BASE_URL = "https://pro-api.coinmarketcap.com/v1"

type CMCScraper struct {
	apiKey       string
	assetNames   string
	apiValuePath string
	dataChannel  chan models.Data
}

func NewCMCScraper(triggerChannel *chan time.Time) *CMCScraper {
	var scraper CMCScraper
	scraper.apiKey = utils.Getenv("CMC_PRO_API_KEY", "")
	scraper.assetNames = utils.Getenv("CMC_ASSET_NAMES", "1,2,3")
	scraper.apiValuePath = utils.Getenv("CMC_API_VALUE_PATH", "quote.USD.price")
	// scraper.apiValuePath = utils.Getenv("CMC_API_VALUE_PATH", "circulating_supply")
	scraper.dataChannel = make(chan models.Data)

	go scraper.run(triggerChannel)
	return &scraper
}

func (scraper *CMCScraper) run(triggerChannel *chan time.Time) {
	for range *triggerChannel {
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
	if 200 != response.StatusCode {
		return errors.New(fmt.Sprintf("CMC API call with return code %d", response.StatusCode))
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	data := gjson.GetBytes(contents, "data")

	data.ForEach(func(key, value gjson.Result) bool {
		var d models.Data
		d.Symbol = gjson.Get(value.String(), "symbol").String()
		d.Value = gjson.Get(value.String(), scraper.apiValuePath).Float()
		d.Timestamp, err = time.Parse(time.RFC3339, gjson.Get(value.String(), "quote.USD.last_updated").String())
		if err != nil {
			log.Error("CMC - Parse timestamp: ", err)
		}
		d.Source = COINMARKETCAP
		scraper.DataChannel() <- d
		return true
	})

	return nil
}

func (scraper *CMCScraper) Close() error {
	log.Info("TO DO.")
	return nil
}

func (scraper *CMCScraper) DataChannel() chan models.Data {
	return scraper.dataChannel
}
