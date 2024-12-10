package main

import (
	"strconv"
	"sync"
	"time"

	models "github.com/diadata-org/decentral-data-feeder/pkg/models"
	scraper "github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	// Separator for entries in the environment variables, i.e. Binance:BTC-USDT,KuCoin:BTC-USDT.
	ENV_SEPARATOR = ","
	// Separator for a pair ticker's assets, i.e. BTC-USDT.
	PAIR_TICKER_SEPARATOR = "-"
	// Separator for a pair on a given exchange, i.e. Binance:BTC-USDT.
	EXCHANGE_PAIR_SEPARATOR = ":"
)

var (
	// Comma separated list of exchanges. Only used in case pairs are read from config files.
	source = utils.Getenv("SOURCE", "Coingecko")
	// Comma separated list of assets which should be retrieved from @source.
	// It is the responsability of each exchange scraper to determine the correct format for the corresponding API calls.
	assets = utils.Getenv("ASSETS", "bitcoin,ethereum")
)

func main() {

	triggerChannel := make(chan time.Time)
	wg := sync.WaitGroup{}

	// // Feeder mechanics
	// privateKeyHex := utils.Getenv("PRIVATE_KEY", "")
	// deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	// blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "https://testnet-rpc.diadata.org")
	// backupNode := utils.Getenv("BACKUP_NODE", "https://testnet-rpc.diadata.org")

	// conn, err := ethclient.Dial(blockchainNode)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// }
	// connBackup, err := ethclient.Dial(backupNode)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the backup Ethereum client: %v", err)
	// }
	// chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "10640"), 10, 64)
	// if err != nil {
	// 	log.Fatalf("Failed to parse chainId: %v", err)
	// }

	// Frequency for the trigger ticker initiating the computation of filter values.
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "20"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v", err)
	}

	// privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	// privateKey, err := crypto.HexToECDSA(privateKeyHex)
	// if err != nil {
	// 	log.Fatalf("Failed to load private key: %v", err)
	// }

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	// if err != nil {
	// 	log.Fatalf("Failed to create authorized transactor: %v", err)
	// }

	// var contract, contractBackup *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService
	// err = onchain.DeployOrBindContract(deployedContract, conn, connBackup, auth, &contract, &contractBackup)
	// if err != nil {
	// 	log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	// }

	// Use a ticker for triggering the processing.
	// This is for testing purposes for now. Could also be request based or other trigger types.
	triggerTick := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for tick := range triggerTick.C {
			triggerChannel <- tick
		}
	}()

	var DS scraper.DataScraper
	switch source {
	case scraper.COINMARKETCAP:
		DS = scraper.NewDataScraper(scraper.COINMARKETCAP, &triggerChannel)
	case scraper.COINGECKO:
		DS = scraper.NewDataScraper(scraper.COINGECKO, &triggerChannel)
	}

	wg.Add(1)
	defer wg.Wait()
	go handleData(DS.DataChannel(), DS, &wg)
	// onchain.OracleUpdateExecutor(auth, contract, conn, chainId, filtersChannel)

}

func handleData(dataChannel chan models.Data, scraper scraper.DataScraper, wg *sync.WaitGroup) {
	for {
		select {
		case data := <-dataChannel:
			log.Info("data received: ", data)
			// TO DO: case close scraper.
		}
	}
}
