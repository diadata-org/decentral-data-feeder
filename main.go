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
	decimals, err := strconv.ParseInt(utils.Getenv("DECIMALS", "18"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse decimals: %v", err)
	}
	// Set up blockchain connections and contracts.
	deployedContract, conn, chainId, privateKey, auth := utils.SetupOnchain()

	// Start collecting and pushing metrics.
	metrics.StartMetrics(conn, privateKey, deployedContract, chainId)

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

		var contract diaoraclev3.DiaOracleV3MultiupdateService
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}
		onchain.OracleUpdateExecutor(auth, c, chainId, source, decimals, DS.DataChannel(), DS.UpdateDoneChannel())

	case scraper.PARTICULA:
		DS = scraper.NewDataScraper(scraper.PARTICULA)

		var contract diaoraclev3.DiaOracleV3MultiupdateService
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}
		onchain.OracleUpdateExecutor(auth, c, chainId, source, decimals, DS.DataChannel(), DS.UpdateDoneChannel())

	case scraper.RWAWS:
		var contract diaoraclev3.DiaOracleV3MultiupdateService
		c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
		if err != nil {
			log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
		}

		s := scraper.NewRWAWSScraper(auth, c, chainId, source, decimals)
		defer s.Close()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}

}
