package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type Metrics struct {
	uptime         prometheus.Gauge
	cpuUsage       prometheus.Gauge
	memoryUsage    prometheus.Gauge
	Contract       *prometheus.GaugeVec
	ExchangePairs  *prometheus.GaugeVec
	gasBalance     prometheus.Gauge
	lastUpdateTime prometheus.Gauge
	chainID        prometheus.Gauge
	imageVersion   *prometheus.GaugeVec
	pushGatewayURL string
	jobName        string
	authUser       string
	authPassword   string
}

func NewMetrics(reg *prometheus.Registry, pushGatewayURL, jobName, authUser, authPassword string, chainID int64, imageVersion string) *Metrics {
	m := &Metrics{
		uptime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "feeder",
			Name:      "uptime_hours",
			Help:      "Feeder Uptime in hours.",
		}),
		cpuUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "feeder",
			Name:      "cpu_usage_percent",
			Help:      "Feeder CPU usage in percent.",
		}),
		memoryUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "feeder",
			Name:      "memory_usage_megabytes",
			Help:      "Feeder Memory usage in megabytes.",
		}),
		Contract: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "feeder",
				Name:      "contract_info",
				Help:      "Feeder contract information.",
			},
			[]string{"contract"}, // Label to store the contract address
		),
		ExchangePairs: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "feeder",
				Name:      "exchange_pairs",
				Help:      "List of exchange pairs to be pushed as labels for each Feeder.",
			},
			[]string{"exchange_pair"}, // Label to store each exchange pair
		),
		gasBalance: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "feeder",
			Name:      "gas_balance",
			Help:      "Gas wallet balance in DIA.",
		}),
		lastUpdateTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "feeder",
			Name:      "last_update_time",
			Help:      "Last update time in UTC timestamp.'",
		}),
		chainID: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "feeder",
			Name:      "chain_id",
			Help:      "The chain ID of the blockchain being monitored.",
		}),
		imageVersion: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "feeder",
				Name:      "image_version",
				Help:      "The Docker image version of the feeder.",
			},
			[]string{"version"},
		),
		pushGatewayURL: pushGatewayURL,
		jobName:        jobName,
		authUser:       authUser,
		authPassword:   authPassword,
	}
	reg.MustRegister(m.uptime)
	reg.MustRegister(m.cpuUsage)
	reg.MustRegister(m.memoryUsage)
	reg.MustRegister(m.Contract)
	reg.MustRegister(m.ExchangePairs)
	reg.MustRegister(m.gasBalance)
	reg.MustRegister(m.lastUpdateTime)
	reg.MustRegister(m.chainID)
	reg.MustRegister(m.imageVersion)

	m.chainID.Set(float64(chainID))
	m.imageVersion.WithLabelValues(imageVersion).Set(1)

	return m
}

func StartPrometheusServer(m *Metrics, port string) {
	if m == nil {
		log.Errorf("Cannot start metrics server: metrics object is nil")
		return
	}

	// Register metrics with the default registry
	prometheus.DefaultRegisterer.MustRegister(m.uptime)
	prometheus.DefaultRegisterer.MustRegister(m.cpuUsage)
	prometheus.DefaultRegisterer.MustRegister(m.memoryUsage)
	prometheus.DefaultRegisterer.MustRegister(m.Contract)
	prometheus.DefaultRegisterer.MustRegister(m.ExchangePairs)
	prometheus.DefaultRegisterer.MustRegister(m.gasBalance)
	prometheus.DefaultRegisterer.MustRegister(m.lastUpdateTime)
	prometheus.DefaultRegisterer.MustRegister(m.chainID)
	prometheus.DefaultRegisterer.MustRegister(m.imageVersion)

	log.Printf("Starting metrics server on :%s", port)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Failed to start metrics server: %v", err)
	}
}
