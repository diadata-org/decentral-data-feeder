package main

import (
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

	var contract diaoraclev3.DIAOracleV3
	cAny, err := onchain.DeployOrBindContract(deployedContract, conn, auth, contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}

	c, ok := cAny.(diaoraclev3.DIAOracleV3)
	if !ok {
		log.Fatalf("unexpected contract type: %T", cAny)
	}

	var wg sync.WaitGroup
	for _, source := range sources {
		wg.Add(1)
		go func(source string, chainID int64, auth *bind.TransactOpts, contract diaoraclev3.DIAOracleV3) {
			defer wg.Done()

			switch source {
			case scraper.RWAWS:
				handleRWAWS(auth, contract, chainID, source, decimalsOracleValue)
			default:
				handleSource(auth, contract, chainID, source, decimalsOracleValue)
			}

		}(source, chainId, auth, c)
	}
	wg.Wait()
}

func handleSource(auth *bind.TransactOpts, contract diaoraclev3.DIAOracleV3, chainId int64, source string, decimalsOracleValue int) {
	DS := scraper.NewDataScraper(source)
	if DS == nil {
		log.Errorf("Unknown source: %s", source)
		return
	}
	onchain.OracleUpdateExecutor(auth, contract, chainId, source, decimalsOracleValue, DS.DataChannel(), DS.UpdateDoneChannel())
}

func handleRWAWS(auth *bind.TransactOpts, contract diaoraclev3.DIAOracleV3, chainId int64, source string, decimalsOracleValue int) {
	s := scraper.NewRWAWSScraper(auth, contract, chainId, source, int64(decimalsOracleValue))
	defer s.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
