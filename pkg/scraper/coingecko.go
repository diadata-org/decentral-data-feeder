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

const (
	COINGECKO              = "Coingecko"
	COINGECKO_API_BASE_URL = "https://pro-api.coingecko.com/api/v3"
)

type CGScraper struct {
	apiKey            string
	assetNames        string
	apiValuePath      string
	dataChannel       chan []byte
	updateDoneChannel chan bool
}

type CGResponse struct {
	Symbol    string
	Value     float64
	Timestamp time.Time
}

func NewCGScraper() *CGScraper {
	var scraper CGScraper
	scraper.apiKey = utils.Getenv("CG_PRO_API_KEY", "")
	scraper.assetNames = utils.Getenv("CG_ASSET_NAMES", "bitcoin,ethereum")
	scraper.apiValuePath = utils.Getenv("CG_API_VALUE_PATH", "usd")
	// scraper.apiValuePath = utils.Getenv("CG_API_VALUE_PATH", "usd_24h_vol")
	scraper.dataChannel = make(chan []byte)
	scraper.updateDoneChannel = make(chan bool)

	go scraper.run()
	return &scraper
}

func (scraper *CGScraper) run() {
	tick := time.NewTicker(time.Duration(120 * time.Second))
	for range tick.C {
		err := scraper.getCGPrice(scraper.assetNames, scraper.apiKey)
		if err != nil {
			log.Errorf("getCGPrice for %s: %v", scraper.assetNames, err)
		}
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
	if response.StatusCode != 200 {
		return fmt.Errorf("coingecko API call with return code %d", response.StatusCode)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	data := gjson.ParseBytes(contents)

	data.ForEach(func(key, value gjson.Result) bool {
		var resp CGResponse
		resp.Symbol = key.String()
		resp.Value = gjson.Get(value.String(), scraper.apiValuePath).Float()
		resp.Timestamp = time.Unix(gjson.Get(value.String(), "last_updated_at").Int(), 0)
		data, err := json.Marshal(resp)
		if err != nil {
			log.Error("Marshal data: ", err)
		}
		scraper.DataChannel() <- data
		return true
	})

	return nil
}

func (scraper *CGScraper) Close() error {
	log.Info("TO DO.")
	return nil
}

func (scraper *CGScraper) DataChannel() chan []byte {
	return scraper.dataChannel
}

func (scraper *CGScraper) UpdateDoneChannel() chan bool {
	return scraper.updateDoneChannel
}
