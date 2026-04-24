package onchain

import (
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	"github.com/diadata-org/decentral-data-feeder/pkg/utils"
	diaoraclev3 "github.com/diadata-org/lumina-library/contracts/lumina/diaoraclev3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

var (
	keys                  []string
	values                []*big.Int
	isFirstRun            = true
	batchSizeOracleUpdate int
)

func init() {
	log = logrus.New()
	loglevel, err := logrus.ParseLevel(utils.Getenv("LOG_LEVEL_UPDATER", "info"))
	if err != nil {
		log.Errorf("Parse log level: %v.", err)
	}
	log.SetLevel(loglevel)

	batchSizeOracleUpdate, err = strconv.Atoi(utils.Getenv("BATCH_SIZE_ORACLE_UPDATE", "20"))
	if err != nil {
		log.Errorf("Parse BATCH_SIZE_ORACLE_UPDATE: %v. Using default value of 20", err)
		batchSizeOracleUpdate = 20
	}
	log.Infof("Batch size for oracle update: %d", batchSizeOracleUpdate)
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
	decimalsOracleValue int,
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
					nyseOpen := big.NewInt(0)
					if twelvedataResponse.NYSEOpen {
						nyseOpen = utils.ScaleInt(1, decimalsOracleValue)
					}
					values = append(values, nyseOpen)

					// holiday.
					keys = append(keys, "US_Holiday")
					nyseHoliday := big.NewInt(0)
					if twelvedataResponse.NYSEHoliday {
						nyseHoliday = utils.ScaleInt(1, decimalsOracleValue)
					}
					values = append(values, nyseHoliday)
				} else {
					log.Info("got rwa data: ", twelvedataResponse)
					keys = append(keys, twelvedataResponse.Symbol)
					values = append(values, utils.ScaleFloat(twelvedataResponse.Price, decimalsOracleValue))
				}

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
					values = append(values, utils.ScaleInt(value, decimalsOracleValue))
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
					marketOpen := big.NewInt(0)
					if rwaResponse.MarketOpen {
						marketOpen = utils.ScaleInt(1, decimalsOracleValue)
					}
					values = append(values, marketOpen)

					keys = append(keys, "Market_Holiday")
					marketHoliday := big.NewInt(0)
					if rwaResponse.MarketHoliday {
						marketHoliday = utils.ScaleInt(1, decimalsOracleValue)
					}
					values = append(values, marketHoliday)
				}

				keys = append(keys, rwaResponse.Symbol)
				values = append(values, utils.ScaleFloat(rwaResponse.Price, decimalsOracleValue))
			}

		case <-updateDoneChannel:

			// update oracle with collected keys and values.
			log.Infof("OracleUpdateExecutor collected %v responses. make oracle update...", len(values))

			switch contract := contractAny.(type) {
			case diaoraclev3.DIAOracleV3:
				err := updateOracleMultiValues(contract, auth, keys, values, time.Now().Unix(), isFirstRun)
				if err != nil {
					log.Warnf("updater - Failed to update Oracle: %v.", err)
					return
				}

			default:
				log.Errorf("DEBUG: Unknown contract type: %T", contractAny)
			}

			// reset keys and values for next update.
			keys = []string{}
			values = []*big.Int{}
			isFirstRun = false // Mark first run as complete
		}
	}
}

func OracleUpdateExecutorForHighFrequencyScraper(
	auth *bind.TransactOpts,
	contractAny any,
	chainId int64,
	source string,
	decimalsOracleValue int,
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
					marketOpen := big.NewInt(0)
					if rwaResponse.MarketOpen {
						marketOpen = utils.ScaleInt(1, decimalsOracleValue)
					}
					values = append(values, marketOpen)

					keys = append(keys, "Market_Holiday")
					marketHoliday := big.NewInt(0)
					if rwaResponse.MarketHoliday {
						marketHoliday = utils.ScaleInt(1, decimalsOracleValue)
					}
					values = append(values, marketHoliday)
				} else {
					keys = append(keys, rwaResponse.Symbol)
					values = append(values, utils.ScaleFloat(rwaResponse.Price, decimalsOracleValue))
				}
			}
		case <-updateDoneChannel:
			log.Infof("OracleUpdateExecutorForHighFrequencyScraper collected %v responses. make oracle update...", len(values))

			keysSnapshot := keys
			valuesSnapshot := values
			isFirstRunSnapshot := isFirstRun

			switch contract := contractAny.(type) {
			case diaoraclev3.DIAOracleV3:
				go func() {
					err := updateOracleMultiValues(contract, auth, keysSnapshot, valuesSnapshot, time.Now().Unix(), isFirstRunSnapshot)
					if err != nil {
						log.Warnf("updater - Failed to update Oracle: %v.", err)
					}
				}()
			default:
				log.Errorf("OracleUpdateExecutorForHighFrequencyScraper - Unknown contract type: %T", contractAny)
			}

			// reset keys and values for next update.
			keys = []string{}
			values = []*big.Int{}
			isFirstRun = false // Mark first run as complete
		}
	}
}

func updateOracleMultiValues(
	contract diaoraclev3.DIAOracleV3,
	auth *bind.TransactOpts,
	keys []string,
	values []*big.Int,
	timestamp int64,
	isFirstRun bool,
) error {

	var cValues []*big.Int
	var gasPrice *big.Int

	for _, value := range values {
		// Create compressed argument with values/timestamps
		cValue := value
		cValue = cValue.Lsh(cValue, 128)
		cValue = cValue.Add(cValue, big.NewInt(timestamp))
		cValues = append(cValues, cValue)
	}

	totalItems := len(keys)

	if isFirstRun {
		if totalItems > batchSizeOracleUpdate {
			log.Infof("updater - First run: Total items: %d, processing in batches of %d", totalItems, batchSizeOracleUpdate)
		}

		for i := 0; i < totalItems; i += batchSizeOracleUpdate {
			end := i + batchSizeOracleUpdate
			if end > totalItems {
				end = totalItems
			}

			batchKeys := keys[i:end]
			batchValues := cValues[i:end]

			log.Infof("updater - Processing batch %d: items %d to %d", (i/batchSizeOracleUpdate)+1, i+1, end)

			// Write values to smart contract
			tx, err := contract.SetMultipleValues(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, batchKeys, batchValues)
			if err != nil {
				log.Errorf("updater - Failed to update batch %d: %v", (i/batchSizeOracleUpdate)+1, err)
				return err
			}

			logTx(tx)
			log.Infof("updater - Successfully processed batch %d", (i/batchSizeOracleUpdate)+1)
		}
	} else {
		log.Infof("updater - Subsequent run: processing all %d items in single transaction", totalItems)

		tx, err := contract.SetMultipleValues(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: gasPrice,
		}, keys, cValues)
		if err != nil {
			log.Errorf("updater - Failed to update oracle: %v", err)
			return err
		}

		logTx(tx)
		log.Infof("updater - Successfully processed all items in single transaction")
	}

	return nil
}

func logTx(tx *types.Transaction) {
	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	// log.Printf("Data: %x\n", tx.Data())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	log.Infof("updater - Tx To: %s.", tx.To().String())
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
}
