package onchain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	diaOracleRandomness "github.com/diadata-org/decentral-data-feeder/contracts/DIAOracleRandomness"
	"github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	"github.com/diadata-org/decentral-data-feeder/pkg/utils"
	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

var txWriteMu sync.Mutex

func init() {
	log = logrus.New()
	loglevel, err := logrus.ParseLevel(utils.Getenv("LOG_LEVEL_UPDATER", "info"))
	if err != nil {
		log.Errorf("Parse log level: %v.", err)
	}
	log.SetLevel(loglevel)
}

func OracleUpdateExecutor(
	conn *ethclient.Client,
	auth *bind.TransactOpts,
	contractAny any,
	chainId int64,
	source string,
	dataChannel <-chan []byte,
	updateDoneChannel <-chan bool,
) {
	var keys []string
	var values []int64

	_ = chainId

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
						continue
					}
					randomBytes.RequestID = randamuResponse.RequestId

					err = updateRandomnessOracle(conn, *contractAny.(*diaOracleRandomness.DIAOracleRandomness), auth, randomBytes)
					if err != nil {
						log.Errorf("updater - Failed to update Oracle: %v.", err)
					}

				case 1:
					var randomIntRange scraper.RandamuIntRangeResponse
					err = json.Unmarshal(randamuResponse.Response, &randomIntRange)
					if err != nil {
						log.Errorf("Unmarshal Randamu random intRange: %v.", err)
						continue
					}
					randomIntRange.RequestID = randamuResponse.RequestId

					err = updateRandomnessOracle(conn, *contractAny.(*diaOracleRandomness.DIAOracleRandomness), auth, randomIntRange)
					if err != nil {
						log.Errorf("updater - Failed to update Oracle: %v.", err)
					}

				case 2:
					var randomInts scraper.RandamuIntResponse
					err = json.Unmarshal(randamuResponse.Response, &randomInts)
					if err != nil {
						log.Errorf("Unmarshal Randamu random ints: %v.", err)
						continue
					}
					randomInts.RequestID = randamuResponse.RequestId

					log.Info("obtained random Ints: ", randomInts.Ints)
					log.Infof("requestID -- round -- seed -- signature: %s -- %v -- %s -- %s",
						randomInts.RequestID.String(),
						randomInts.Metadata.Round,
						randomInts.Metadata.Seed,
						randomInts.Metadata.Signature,
					)

					err = updateRandomnessOracle(conn, *contractAny.(*diaOracleRandomness.DIAOracleRandomness), auth, randomInts)
					if err != nil {
						log.Errorf("updater - Failed to update Oracle: %v.", err)
					}
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
			case *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService:
				err := updateOracleMultiValues(conn, *contract, auth, keys, values, time.Now().Unix())
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

func updateRandomnessOracle(
	conn *ethclient.Client,
	contract diaOracleRandomness.DIAOracleRandomness,
	auth *bind.TransactOpts,
	data any,
) error {
	var gasPrice *big.Int

	switch data := data.(type) {
	case scraper.RandamuBytesResponse:
		txWriteMu.Lock()
		defer txWriteMu.Unlock()

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

		receipt, err := bind.WaitMined(context.Background(), conn, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("randomness bytes tx failed: %s", tx.Hash().Hex())
		}

	case scraper.RandamuIntRangeResponse:
		var bigInts []*big.Int
		for _, i := range data.Ints {
			i64, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				return err
			}
			bigInts = append(bigInts, big.NewInt(i64))
		}

		txWriteMu.Lock()
		defer txWriteMu.Unlock()

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

		receipt, err := bind.WaitMined(context.Background(), conn, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("randomness int range tx failed: %s", tx.Hash().Hex())
		}

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

		txWriteMu.Lock()
		defer txWriteMu.Unlock()

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

		receipt, err := bind.WaitMined(context.Background(), conn, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("randomness int array tx failed: %s", tx.Hash().Hex())
		}
	}

	return nil
}

func updateOracleMultiValues(
	conn *ethclient.Client,
	contract diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	auth *bind.TransactOpts,
	keys []string,
	values []int64,
	timestamp int64,
) error {
	var cValues []*big.Int
	var gasPrice *big.Int

	for _, value := range values {
		// Create compressed argument with values/timestamps
		cValue := big.NewInt(value)
		cValue = cValue.Lsh(cValue, 128)
		cValue = cValue.Add(cValue, big.NewInt(timestamp))
		cValues = append(cValues, cValue)
	}

	txWriteMu.Lock()
	defer txWriteMu.Unlock()

	tx, err := contract.SetMultipleValues(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, keys, cValues)
	if err != nil {
		return err
	}

	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	if tx.To() != nil {
		log.Infof("updater - Tx To: %s.", tx.To().String())
	}
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())

	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("oracle update tx failed: %s", tx.Hash().Hex())
	}

	log.Infof("updater - Tx mined successfully: %s.", tx.Hash().Hex())
	return nil
}

func logTx(tx *types.Transaction) {
	log.Infof("updater - Gas price: %d.", tx.GasPrice())
	log.Infof("updater - Nonce: %d.", tx.Nonce())
	if tx.To() != nil {
		log.Infof("updater - Tx To: %s.", tx.To().String())
	}
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
}
