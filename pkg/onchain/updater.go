package onchain

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	"github.com/diadata-org/decentral-data-feeder/pkg/utils"
	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/sirupsen/logrus"
)

var (
	log    *logrus.Logger
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
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	chainId int64,
	// compatibilityMode bool,
	source string,
	dataChannel <-chan []byte,
	updateDoneChannel chan bool,
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
				log.Info("got twelvedata data: ", twelvedataResponse)
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
						log.Error("Unmarshal Randamu random bytes: ", err)
					} else {
						log.Info("final result-----------------: ", randomBytes)
					}
				case 1:
					var randomIntRange scraper.RandamuIntRangeResponse
					err = json.Unmarshal(randamuResponse.Response, &randomIntRange)
					if err != nil {
						log.Error("Unmarshal Randamu random intRange: ", err)
					} else {
						log.Info("final result-----------------: ", randomIntRange)
					}
				case 2:
					var randomInts scraper.RandamuIntResponse
					err = json.Unmarshal(randamuResponse.Response, &randomInts)
					if err != nil {
						log.Error("Unmarshal Randamu random ints: ", err)
					} else {
						log.Info("final result-----------------: ", randomInts)
					}
				}

			}

		case <-updateDoneChannel:
			// TO DO: We might need a switch here in order to account for different types of contracts.
			// update oracle with collected keys and values.
			log.Infof("collected %v responses. make oracle update...", len(values))
			err := updateOracleMultiValues(contract, auth, keys, values, time.Now().Unix())
			if err != nil {
				log.Warnf("updater - Failed to update Oracle: %v.", err)
				return
			}
			// reset keys and values for next update.
			keys = []string{}
			values = []int64{}
		}

	}
}

func updateOracleMultiValues(
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
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
