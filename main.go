package main

import (
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	diaOracleRandomness "github.com/diadata-org/decentral-data-feeder/contracts/DIAOracleRandomness"
	"github.com/diadata-org/decentral-data-feeder/pkg/metrics"
	"github.com/diadata-org/decentral-data-feeder/pkg/onchain"
	scraper "github.com/diadata-org/decentral-data-feeder/pkg/scraper"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
	"github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"
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

// GetImageVersion returns the Docker image version from environment variable
func GetImageVersion() string {
	// Get version from IMAGE_TAG environment variable
	version := os.Getenv("IMAGE_TAG")

	if version == "" {
		version = "unknown" // fallback if not set
		log.Info("No version found, using 'unknown'")
	}

	return version
}

func main() {

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

	// ----------------------------------- Metrics -----------------------------------
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Failed to get hostname: %v", err)
	}

	// Check if metrics pushing to Pushgateway is enabled
	pushgatewayURL := os.Getenv("PUSHGATEWAY_URL")
	authUser := os.Getenv("PUSHGATEWAY_USER")
	authPassword := os.Getenv("PUSHGATEWAY_PASSWORD")
	pushgatewayEnabled := pushgatewayURL != "" && authUser != "" && authPassword != ""

	// Check if Prometheus HTTP server is enabled
	enablePrometheusServer := utils.Getenv("ENABLE_METRICS_SERVER", "false")
	prometheusServerEnabled := strings.ToLower(enablePrometheusServer) == "true"

	// Get the node operator ID from the environment variable (optional)
	nodeOperatorName := utils.Getenv("NODE_OPERATOR_NAME", "")

	// Create the job name for metrics (used for both modes)
	jobName := metrics.MakeJobName(hostname, nodeOperatorName)

	// Get image version using our local function
	imageVersion := GetImageVersion()
	log.Infof("Image version: %s", imageVersion)

	// Set default pushgateway URL if enabled
	if pushgatewayEnabled {
		if pushgatewayURL == "" {
			pushgatewayURL = "https://pushgateway-auth.diadata.org"
		}
		log.Info("Metrics pushing enabled. Pushing to: ", pushgatewayURL)
	} else {
		log.Info("Metrics pushing to Pushgateway disabled")
	}

	// Create metrics object
	m := metrics.NewMetrics(
		prometheus.NewRegistry(),
		pushgatewayURL,
		jobName,
		authUser,
		authPassword,
		chainId,
		imageVersion,
	)

	// Start Prometheus HTTP server if enabled
	if prometheusServerEnabled {
		metricsPort := utils.Getenv("METRICS_PORT", "9090")
		go metrics.StartPrometheusServer(m, metricsPort)
		log.Info("Prometheus HTTP server enabled on port:", metricsPort)
	} else {
		log.Info("Prometheus HTTP server disabled")
	}

	// Record start time for uptime calculation
	startTime := time.Now()

	// Move metrics setup here, right before the blocking call
	// Only setup metrics collection if metrics are enabled and metrics object exists
	if pushgatewayEnabled && m != nil {
		// Set the static contract label for Prometheus monitoring
		m.Contract.WithLabelValues(deployedContract).Set(1)

		// TO DO: Adapt this to decentral-data-feeder repository. In particular, move it to
		// the respective scraper in order to fetch the symbols, if existent.
		// exchangePairsList := strings.Split(exchangePairsEnv, ",")
		// for _, pair := range exchangePairsList {
		// 	pair = strings.TrimSpace(pair) // Clean whitespace
		// 	if pair != "" {
		// 		m.ExchangePairs.WithLabelValues(pair).Set(1)
		// 	}
		// }

		// Push metrics to Pushgateway if enabled
		go metrics.PushMetricsToPushgateway(m, startTime, conn, privateKey, deployedContract)
	}

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
