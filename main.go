package main

import (
	diaOracleRandomness "github.com/diadata-org/decentral-data-feeder/contracts/DIAOracleRandomness"
	"github.com/diadata-org/decentral-data-feeder/pkg/metrics"
	"github.com/diadata-org/decentral-data-feeder/pkg/onchain"
	scraper "github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
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

	// Set up blockchain connections and contracts.
	deployedContract, conn, chainId, privateKey, auth := utils.SetupOnchain()

	// Start collecting and pushing metrics.
	metrics.StartMetrics(conn, privateKey, deployedContract, chainId)

	// ----------------------------------- Metrics -----------------------------------

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
