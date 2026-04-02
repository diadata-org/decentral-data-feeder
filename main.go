package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/diadata-org/decentral-data-feeder/pkg/metrics"
	"github.com/diadata-org/decentral-data-feeder/pkg/onchain"
	scraper "github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	diaoraclev3 "github.com/diadata-org/lumina-library/contracts/lumina/diaoraclev3"
	log "github.com/sirupsen/logrus"
)

var (
	source = utils.Getenv("SOURCE", "")
)

func main() {

	// Set up blockchain connections and contracts.
	deployedContract, conn, chainId, privateKey, auth := utils.SetupOnchain()

	// Start collecting and pushing metrics.
	metrics.StartMetrics(conn, privateKey, deployedContract, chainId)

	decimalsOracleValue, err := strconv.Atoi(utils.Getenv("DECIMALS_ORACLE_VALUE", "18"))
	if err != nil {
		log.Errorf("parse DECIMALS_ORACLE_VALUE: %v", err)
		decimalsOracleValue = 18
	}
	log.Infof("Using DECIMALS_ORACLE_VALUE: %d", decimalsOracleValue)

	// ----------------------------------- Metrics -----------------------------------

	var DS scraper.DataScraper
	switch source {
	// case scraper.COINMARKETCAP:
	//  DS = scraper.NewDataScraper(scraper.COINMARKETCAP)
	//  // TO DO: add oracle updater

	// case scraper.COINGECKO:
	//  DS = scraper.NewDataScraper(scraper.COINGECKO)
	//  // TO DO: add oracle updater

	case scraper.TWELVEDATA:
		DS = scraper.NewDataScraper(scraper.TWELVEDATA)

		var contract diaoraclev3.DIAOracleV3
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}
		onchain.OracleUpdateExecutor(auth, c, chainId, source, decimalsOracleValue, DS.DataChannel(), DS.UpdateDoneChannel())

	case scraper.PARTICULA:
		DS = scraper.NewDataScraper(scraper.PARTICULA)

		var contract diaoraclev3.DIAOracleV3
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}
		onchain.OracleUpdateExecutor(auth, c, chainId, source, decimalsOracleValue, DS.DataChannel(), DS.UpdateDoneChannel())

	case scraper.RWAWS:
		var contract diaoraclev3.DIAOracleV3
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}

		s := scraper.NewRWAWSScraper(auth, c, chainId, source)
		defer s.Close()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}

}
