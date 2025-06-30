package main

import (
	"math/big"
	"strconv"
	"strings"

	diaOracleRandomness "github.com/diadata-org/decentral-data-feeder/contracts/DIAOracleRandomness"
	"github.com/diadata-org/decentral-data-feeder/pkg/onchain"
	scraper "github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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
	source = utils.Getenv("SOURCE", "Randamu")
)

func main() {

	// wg := sync.WaitGroup{}

	// Feeder mechanics
	privateKeyHex := utils.Getenv("PRIVATE_KEY", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	backupNode := utils.Getenv("BACKUP_NODE", "")

	conn, err := utils.MakeEthClient(blockchainNode, backupNode)
	if err != nil {
		log.Fatalf("MakeEthClient: %v", err)
	}

	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "100640"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v", err)
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var DS scraper.DataScraper
	switch source {
	case scraper.COINMARKETCAP:
		DS = scraper.NewDataScraper(scraper.COINMARKETCAP)
		// TO DO: add oracle updater

	case scraper.COINGECKO:
		DS = scraper.NewDataScraper(scraper.COINGECKO)
		// TO DO: add oracle updater

	case scraper.RANDAMU:
		DS = scraper.NewDataScraper(scraper.RANDAMU)

		var contract diaOracleRandomness.DIAOracleRandomness
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}
		onchain.OracleUpdateExecutor(auth, c, chainId, source, DS.DataChannel(), DS.UpdateDoneChannel())

	case scraper.TWELVEDATA:
		DS = scraper.NewDataScraper(scraper.TWELVEDATA)

		var contract diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}
		onchain.OracleUpdateExecutor(auth, c, chainId, source, DS.DataChannel(), DS.UpdateDoneChannel())

	}

}
