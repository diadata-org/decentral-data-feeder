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
	values []*big.Float
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
	decimal int64,
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
					nyseOpen := float64(0)
					if twelvedataResponse.NYSEOpen {
						nyseOpen = float64(1)
					}
					values = append(values, new(big.Float).SetFloat64(nyseOpen))

					// holiday.
					keys = append(keys, "US_Holiday")
					nyseHoliday := float64(0)
					if twelvedataResponse.NYSEHoliday {
						nyseHoliday = float64(1)
					}
					values = append(values, new(big.Float).SetFloat64(nyseHoliday))
				}

				log.Info("got rwa data: ", twelvedataResponse)
				keys = append(keys, twelvedataResponse.Symbol)
				values = append(values, new(big.Float).SetFloat64(twelvedataResponse.Price))

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
					values = append(values, new(big.Float).SetFloat64(float64(value)))
				}
			}

		case <-updateDoneChannel:

			// update oracle with collected keys and values.
			log.Infof("collected %v responses. make oracle update...", len(values))
			switch contract := contractAny.(type) {
			case *diaoraclev3.DiaOracleV3MultiupdateService:
				err := updateOracleMultiValues(*contract, auth, keys, values, time.Now().Unix(), decimal)
				if err != nil {
					log.Warnf("updater - Failed to update Oracle: %v.", err)
					return
				}

			}

			// reset keys and values for next update.
			keys = []string{}
			values = []*big.Float{}
		}
	}
}

func updateOracleMultiValues(
	contract diaoraclev3.DiaOracleV3MultiupdateService,
	auth *bind.TransactOpts,
	keys []string,
	values []*big.Float,
	timestamp int64,
	decimals int64) error {

	var cValues []*big.Int
	var gasPrice *big.Int
	var err error

	multiplier := new(big.Float).SetInt(
		new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil),
	)

	for _, value := range values {
		scaled := new(big.Float).Mul(value, multiplier)
		scaledInt, _ := scaled.Int(nil)

		cValue := new(big.Int).Lsh(scaledInt, 128)
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
