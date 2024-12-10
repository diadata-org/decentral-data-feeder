package scraper

import (
	"fmt"
	"io"
	"net/http"
	"time"

	models "github.com/diadata-org/decentral-data-feeder/pkg/models"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/tidwall/gjson"
)

const (
	COINGECKO              = "Coingecko"
	COINGECKO_API_BASE_URL = "https://pro-api.coingecko.com/api/v3"
)

type CGScraper struct {
	apiKey       string
	assetNames   string
	apiValuePath string
	dataChannel  chan models.Data
}

func NewCGScraper(triggerChannel *chan time.Time) *CGScraper {
	var scraper CGScraper
	scraper.apiKey = utils.Getenv("CG_PRO_API_KEY", "")
	scraper.assetNames = utils.Getenv("CG_ASSET_NAMES", "bitcoin,ethereum")
	scraper.apiValuePath = utils.Getenv("CG_API_VALUE_PATH", "usd")
	// scraper.apiValuePath = utils.Getenv("CG_API_VALUE_PATH", "usd_24h_vol")
	scraper.dataChannel = make(chan models.Data)

	go scraper.run(triggerChannel)
	return &scraper
}

func (scraper *CGScraper) run(triggerChannel *chan time.Time) {
	for range *triggerChannel {
		log.Info("trigger")
		scraper.getCGPrice(scraper.assetNames, scraper.apiKey)
	}
}

func (scraper *CGScraper) getCGPrice(assetName, apiKey string) error {
	url := COINGECKO_API_BASE_URL + "/simple/price?ids=" + assetName +
		"&include_last_updated_at=true&include_24h_vol=true" +
		"&include_24hr_vol=true" +
		"&vs_currencies=usd&x_cg_pro_api_key=" + apiKey

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return fmt.Errorf("Error on coingecko API call with return code %d", response.StatusCode)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	data := gjson.ParseBytes(contents)

	data.ForEach(func(key, value gjson.Result) bool {
		var d models.Data
		d.Symbol = key.String()
		d.Value = gjson.Get(value.String(), scraper.apiValuePath).Float()
		d.Timestamp = time.Unix(gjson.Get(value.String(), "last_updated_at").Int(), 0)

		d.Source = COINGECKO
		scraper.DataChannel() <- d
		return true
	})

	return nil
}

func (scraper *CGScraper) Close() error {
	log.Info("TO DO.")
	return nil
}

func (scraper *CGScraper) DataChannel() chan models.Data {
	return scraper.dataChannel
}
