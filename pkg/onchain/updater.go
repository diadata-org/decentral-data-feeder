package onchain

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	"github.com/diadata-org/decentral-data-feeder/pkg/utils"
	diaoraclev3 "github.com/diadata-org/lumina-library/contracts/lumina/diaoraclev3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

var (
	keys   []string
	values []int64
)

func init() {
	log = logrus.New()
	loglevel, err := logrus.ParseLevel(utils.Getenv("LOG_LEVEL_UPDATER", "info"))
	if err != nil {
		log.Errorf("Parse log level: %v.", err)
	}
	log.SetLevel(loglevel)
}

func OracleUpdateExecutor(
	// publishedPrices map[string]float64,
	// newPrices map[string]float64,
	// deviationPermille int,
	auth *bind.TransactOpts,
	contractAny any,
	chainId int64,
	// compatibilityMode bool,
	source string,
	dataChannel <-chan []byte,
	updateDoneChannel <-chan bool,
) {

	for {

		select {
		case data := <-dataChannel:

			switch source {
			case scraper.TWELVEDATA:
				var twelvedataResponse scraper.TwelvedataQuote
				err := json.Unmarshal(data, &twelvedataResponse)
				if err != nil {
					log.Error("Unmarshal Twelvedata response: ", err)
					continue
				}

				if twelvedataResponse.Type == scraper.NyseOpen {
					// business time.
					keys = append(keys, "US_Open")
					nyseOpen := int64(0)
					if twelvedataResponse.NYSEOpen {
						nyseOpen = int64(1)
					}
					values = append(values, nyseOpen)

					// holiday.
					keys = append(keys, "US_Holiday")
					nyseHoliday := int64(0)
					if twelvedataResponse.NYSEHoliday {
						nyseHoliday = int64(1)
					}
					values = append(values, nyseHoliday)
				}

				log.Info("got rwa data: ", twelvedataResponse)
				keys = append(keys, twelvedataResponse.Symbol)
				values = append(values, int64(twelvedataResponse.Price*1e5))

			case scraper.PARTICULA:
				tokenRating := make(map[string]int64)
				err := json.Unmarshal(data, &tokenRating)
				if err != nil {
					log.Error("Unmarshal particula response: ", err)
					continue
				}

				log.Info("got particula token rating data: ", tokenRating)
				for key, value := range tokenRating {
					keys = append(keys, key)
					values = append(values, value)
				}

			case scraper.RWAWS:
				var rwaResponse scraper.RWAWSQuote
				err := json.Unmarshal(data, &rwaResponse)
				if err != nil {
					log.Error("Unmarshal RWAWS response: ", err)
					continue
				}

				log.Info("got rwa ws data: ", rwaResponse)

				if rwaResponse.Type == scraper.Equities || rwaResponse.Type == scraper.ETF {
					keys = append(keys, "Market_Open")
					marketOpen := int64(0)
					if rwaResponse.MarketOpen {
						marketOpen = int64(1)
					}
					values = append(values, marketOpen)

					keys = append(keys, "Market_Holiday")
					marketHoliday := int64(0)
					if rwaResponse.MarketHoliday {
						marketHoliday = int64(1)
					}
					values = append(values, marketHoliday)
				}

				keys = append(keys, rwaResponse.Symbol)
				values = append(values, int64(rwaResponse.Price*1e5))
			}

		case <-updateDoneChannel:

			// update oracle with collected keys and values.
			log.Infof("collected %v responses. make oracle update...", len(values))
			switch contract := contractAny.(type) {
			case *diaoraclev3.DiaOracleV3MultiupdateService:
				err := updateOracleMultiValues(*contract, auth, keys, values, time.Now().Unix())
				if err != nil {
					log.Warnf("updater - Failed to update Oracle: %v.", err)
					return
				}

			}

			// reset keys and values for next update.
			keys = []string{}
			values = []int64{}
		}
	}
}

func OracleUpdateExecutorForHighFrequencyScraper(
	auth *bind.TransactOpts,
	contractAny any,
	chainId int64,
	source string,
	dataChannel <-chan []byte,
	updateDoneChannel <-chan bool,
) {

	for {
		select {
		case data := <-dataChannel:

			switch source {
			case scraper.RWAWS:
				var rwaResponse scraper.RWAWSQuote
				err := json.Unmarshal(data, &rwaResponse)
				if err != nil {
					log.Error("Unmarshal RWAWS response: ", err)
					continue
				}

				log.Info("got rwa ws data: ", rwaResponse)

				if rwaResponse.Type == scraper.Equities || rwaResponse.Type == scraper.ETF {
					keys = append(keys, "Market_Open")
					marketOpen := int64(0)
					if rwaResponse.MarketOpen {
						marketOpen = int64(1)
					}
					values = append(values, marketOpen)

					keys = append(keys, "Market_Holiday")
					marketHoliday := int64(0)
					if rwaResponse.MarketHoliday {
						marketHoliday = int64(1)
					}
					values = append(values, marketHoliday)
				}

				keys = append(keys, rwaResponse.Symbol)
				values = append(values, int64(rwaResponse.Price*1e5))
			}
		case <-updateDoneChannel:
			log.Infof("collected %v responses. make oracle update...", len(values))

			keysSnapshot := keys
			valuesSnapshot := values

			switch contract := contractAny.(type) {
			case *diaoraclev3.DiaOracleV3MultiupdateService:
				go func() {
					err := updateOracleMultiValues(*contract, auth, keysSnapshot, valuesSnapshot, time.Now().Unix())
					if err != nil {
						log.Warnf("updater - Failed to update Oracle: %v.", err)
					}
				}()
			}

			// reset keys and values for next update.
			keys = []string{}
			values = []int64{}
		}
	}
}

func updateOracleMultiValues(
	contract diaoraclev3.DiaOracleV3MultiupdateService,
	auth *bind.TransactOpts,
	keys []string,
	values []int64,
	timestamp int64) error {

	var cValues []*big.Int
	var gasPrice *big.Int
	var err error

	for _, value := range values {
		// Create compressed argument with values/timestamps
		cValue := big.NewInt(value)
		cValue = cValue.Lsh(cValue, 128)
		cValue = cValue.Add(cValue, big.NewInt(timestamp))
		cValues = append(cValues, cValue)
	}

	// Write values to smart contract
	tx, err := contract.SetMultipleValues(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, keys, cValues)
	// check if tx is sendable then fgo backup
	if err != nil {
		// backup in here
		return err
	}

	logTx(tx)
	return nil
}

func logTx(tx *types.Transaction) {
	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	// log.Printf("Data: %x\n", tx.Data())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	log.Infof("updater - Tx To: %s.", tx.To().String())
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
}
