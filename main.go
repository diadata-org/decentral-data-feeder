package main

import (
	"crypto/ecdsa"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/diadata-org/decentral-data-feeder/pkg/metrics"
	"github.com/diadata-org/decentral-data-feeder/pkg/onchain"
	scraper "github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/diadata-org/lumina-library/contracts/lumina/diaoraclev3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
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

	sources, err := scraper.GetSourcesFromEnv("SOURCES")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for _, source := range sources {
		wg.Add(1)
		go func(source string, deployedContract string, conn *ethclient.Client, chainID int64, privateKey *ecdsa.PrivateKey, auth *bind.TransactOpts) {
			defer wg.Done()

			switch source {
			case scraper.TWELVEDATA:
				handleTwelveData(deployedContract, conn, auth, chainId, source, decimalsOracleValue)
			case scraper.PARTICULA:
				handleParticula(deployedContract, conn, auth, chainId, source, decimalsOracleValue)
			case scraper.BELO:
				handleBelo(deployedContract, conn, auth, chainID, source, decimalsOracleValue)
			case scraper.RWAWS:
				handleRWAWS(deployedContract, conn, auth, chainId, source, decimalsOracleValue)
			}

		}(source, deployedContract, conn, chainId, privateKey, auth)
	}
	wg.Wait()
}

func handleTwelveData(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, chainId int64, source string, decimalsOracleValue int) {
	DS := scraper.NewDataScraper(scraper.TWELVEDATA)
	var contract diaoraclev3.DIAOracleV3
	c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}
	onchain.OracleUpdateExecutor(auth, c, chainId, source, decimalsOracleValue, DS.DataChannel(), DS.UpdateDoneChannel())
}

func handleBelo(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, chainId int64, source string, decimalsOracleValue int) {
	DS := scraper.NewDataScraper(scraper.BELO)
	var contract diaoraclev3.DIAOracleV3
	c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}
	onchain.OracleUpdateExecutor(auth, c, chainId, source, decimalsOracleValue, DS.DataChannel(), DS.UpdateDoneChannel())
}

func handleParticula(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, chainId int64, source string, decimalsOracleValue int) {
	DS := scraper.NewDataScraper(scraper.PARTICULA)
	var contract diaoraclev3.DIAOracleV3
	c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}
	onchain.OracleUpdateExecutor(auth, c, chainId, source, decimalsOracleValue, DS.DataChannel(), DS.UpdateDoneChannel())
}

func handleRWAWS(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, chainId int64, source string, decimalsOracleValue int) {
	var contract diaoraclev3.DIAOracleV3
	c, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}

	s := scraper.NewRWAWSScraper(auth, c, chainId, source, int64(decimalsOracleValue))
	defer s.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}