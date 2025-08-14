package onchain

import (
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	diaOracleRandomness "github.com/diadata-org/decentral-data-feeder/contracts/DIAOracleRandomness"
	"github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	"github.com/diadata-org/decentral-data-feeder/pkg/utils"
	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
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
					keys = append(keys, "NYSE_Open")
					nyseOpen := int64(0)
					if twelvedataResponse.NYSEOpen {
						nyseOpen = int64(1)
					}
					values = append(values, nyseOpen)

					// holiday.
					keys = append(keys, "NYSE_Holiday")
					nyseHoliday := int64(0)
					if twelvedataResponse.NYSEHoliday {
						nyseHoliday = int64(1)
					}
					values = append(values, nyseHoliday)
				}

				log.Info("got rwa data: ", twelvedataResponse)
				keys = append(keys, twelvedataResponse.Symbol)
				values = append(values, int64(twelvedataResponse.Price*1e5))

			case scraper.RANDAMU:

				var randamuResponse scraper.RandamuResponse
				err := json.Unmarshal(data, &randamuResponse)
				if err != nil {
					log.Error("Unmarshal Randamu response: ", err)
					continue
				}

				switch randamuResponse.Type {
				case 0:
					var randomBytes scraper.RandamuBytesResponse
					err = json.Unmarshal(randamuResponse.Response, &randomBytes)
					if err != nil {
						log.Errorf("Unmarshal Randamu random bytes: %v.", err)
					}
					randomBytes.RequestID = randamuResponse.RequestId

					err = updateRandomnessOracle(*contractAny.(*diaOracleRandomness.DIAOracleRandomness), auth, randomBytes)
					if err != nil {
						log.Errorf("updater - Failed to update Oracle: %v.", err)
					}

				case 1:
					var randomIntRange scraper.RandamuIntRangeResponse
					err = json.Unmarshal(randamuResponse.Response, &randomIntRange)
					if err != nil {
						log.Errorf("Unmarshal Randamu random intRange: %v.", err)
					}
					randomIntRange.RequestID = randamuResponse.RequestId

					err = updateRandomnessOracle(*contractAny.(*diaOracleRandomness.DIAOracleRandomness), auth, randomIntRange)
					if err != nil {
						log.Errorf("updater - Failed to update Oracle: %v.", err)
					}

				case 2:
					var randomInts scraper.RandamuIntResponse
					err = json.Unmarshal(randamuResponse.Response, &randomInts)
					if err != nil {
						log.Errorf("Unmarshal Randamu random ints: %v.", err)
					}
					randomInts.RequestID = randamuResponse.RequestId

					log.Info("obtained random Ints: ", randomInts.Ints)
					log.Infof("requestID -- round -- seed -- signature: %s -- %v -- %s -- %s", randomInts.RequestID.String(), randomInts.Metadata.Round, randomInts.Metadata.Seed, randomInts.Metadata.Signature)

					err = updateRandomnessOracle(*contractAny.(*diaOracleRandomness.DIAOracleRandomness), auth, randomInts)
					if err != nil {
						log.Errorf("updater - Failed to update Oracle: %v.", err)
					}

				}

			}

		case <-updateDoneChannel:

			// update oracle with collected keys and values.
			log.Infof("collected %v responses. make oracle update...", len(values))
			switch contract := contractAny.(type) {
			case *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService:
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

func updateRandomnessOracle(contract diaOracleRandomness.DIAOracleRandomness, auth *bind.TransactOpts, data any) error {
	var gasPrice *big.Int
	switch data := data.(type) {
	case scraper.RandamuBytesResponse:
		tx, err := contract.SetBytes(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: gasPrice,
		},
			data.RequestID,
			data.Randomness,
			data.Metadata.Round,
			data.Metadata.Seed,
			data.Metadata.Signature,
		)
		if err != nil {
			return err
		}
		logTx(tx)

	case scraper.RandamuIntRangeResponse:
		var bigInts []*big.Int
		for _, i := range data.Ints {
			i64, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				return err
			}
			bigInts = append(bigInts, big.NewInt(i64))
		}
		tx, err := contract.SetIntRange(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: gasPrice,
		},
			data.RequestID,
			bigInts,
			data.Metadata.Round,
			data.Metadata.Seed,
			data.Metadata.Signature,
		)
		if err != nil {
			return err
		}
		logTx(tx)

	case scraper.RandamuIntResponse:
		log.Info("update oracle with IntResponse")
		var bigInts []*big.Int
		for _, i := range data.Ints {
			i64, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				return err
			}
			bigInts = append(bigInts, big.NewInt(i64))
		}

		tx, err := contract.SetIntArray(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: gasPrice,
		},
			data.RequestID,
			bigInts,
			data.Metadata.Round,
			data.Metadata.Seed,
			data.Metadata.Signature,
		)
		if err != nil {
			return err
		}
		logTx(tx)

	}

	return nil
}

func updateOracleMultiValues(
	contract diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
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

	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	// log.Printf("Data: %x\n", tx.Data())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	log.Infof("updater - Tx To: %s.", tx.To().String())
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
	return nil
}

func logTx(tx *types.Transaction) {
	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	// log.Printf("Data: %x\n", tx.Data())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	log.Infof("updater - Tx To: %s.", tx.To().String())
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
}
