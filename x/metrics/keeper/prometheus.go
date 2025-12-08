package keeper

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// PrometheusMetrics holds all Prometheus metrics for the Stateset blockchain
type PrometheusMetrics struct {
	// Block metrics
	BlockHeight     prometheus.Gauge
	BlockTime       prometheus.Gauge
	TransactionsTotal prometheus.Counter

	// Economic metrics
	TotalCollateral   prometheus.Gauge
	TotalDebt         prometheus.Gauge
	CollateralRatio   prometheus.Gauge
	StablecoinSupply  prometheus.Gauge

	// Settlement metrics
	SettlementsTotal  prometheus.Counter
	SettlementVolume  prometheus.Counter
	ActiveEscrows     prometheus.Gauge
	ActiveChannels    prometheus.Gauge

	// Oracle metrics
	PriceUpdatesTotal prometheus.Counter
	StalePrices       prometheus.Gauge
	OracleLatency     prometheus.Histogram

	// Security metrics
	CircuitTripsTotal   prometheus.Counter
	RateLimitHitsTotal  prometheus.Counter
	ComplianceBlocksTotal prometheus.Counter
	GlobalPauseActive   prometheus.Gauge

	// Module health
	ModuleErrorRate *prometheus.GaugeVec
	ModuleLatency   *prometheus.HistogramVec
}

// NewPrometheusMetrics creates and registers all Prometheus metrics
func NewPrometheusMetrics(namespace string) *PrometheusMetrics {
	return &PrometheusMetrics{
		// Block metrics
		BlockHeight: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "block_height",
			Help:      "Current block height",
		}),
		BlockTime: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "block_time_seconds",
			Help:      "Average block time in seconds",
		}),
		TransactionsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "transactions_total",
			Help:      "Total number of transactions processed",
		}),

		// Economic metrics
		TotalCollateral: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "total_collateral",
			Help:      "Total collateral value in the system",
		}),
		TotalDebt: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "total_debt",
			Help:      "Total debt value in the system",
		}),
		CollateralRatio: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "collateral_ratio",
			Help:      "System-wide collateral ratio",
		}),
		StablecoinSupply: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "stablecoin_supply",
			Help:      "Total stablecoin supply",
		}),

		// Settlement metrics
		SettlementsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "settlements_total",
			Help:      "Total number of settlements",
		}),
		SettlementVolume: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "settlement_volume_total",
			Help:      "Total settlement volume",
		}),
		ActiveEscrows: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "active_escrows",
			Help:      "Number of active escrows",
		}),
		ActiveChannels: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "active_payment_channels",
			Help:      "Number of active payment channels",
		}),

		// Oracle metrics
		PriceUpdatesTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "price_updates_total",
			Help:      "Total number of price updates",
		}),
		StalePrices: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "stale_prices",
			Help:      "Number of stale prices",
		}),
		OracleLatency: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "oracle_latency_seconds",
			Help:      "Oracle price update latency",
			Buckets:   []float64{0.1, 0.5, 1, 2, 5, 10, 30},
		}),

		// Security metrics
		CircuitTripsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "circuit_trips_total",
			Help:      "Total number of circuit breaker trips",
		}),
		RateLimitHitsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "rate_limit_hits_total",
			Help:      "Total number of rate limit hits",
		}),
		ComplianceBlocksTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "compliance_blocks_total",
			Help:      "Total number of compliance blocks",
		}),
		GlobalPauseActive: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "global_pause_active",
			Help:      "Whether the global pause is active (1 = paused, 0 = running)",
		}),

		// Module health
		ModuleErrorRate: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "module_error_rate",
			Help:      "Error rate per module",
		}, []string{"module"}),
		ModuleLatency: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "module_latency_seconds",
			Help:      "Module operation latency",
			Buckets:   prometheus.DefBuckets,
		}, []string{"module", "operation"}),
	}
}

// DefaultPrometheusMetrics creates metrics with the default "stateset" namespace
func DefaultPrometheusMetrics() *PrometheusMetrics {
	return NewPrometheusMetrics("stateset")
}

// Global metrics instance
var GlobalMetrics *PrometheusMetrics

// InitGlobalMetrics initializes the global metrics instance
func InitGlobalMetrics(namespace string) {
	if GlobalMetrics == nil {
		GlobalMetrics = NewPrometheusMetrics(namespace)
	}
}

// RecordBlockHeight updates the block height metric
func RecordBlockHeight(height int64) {
	if GlobalMetrics != nil {
		GlobalMetrics.BlockHeight.Set(float64(height))
	}
}

// RecordBlockTime updates the block time metric
func RecordBlockTime(seconds float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.BlockTime.Set(seconds)
	}
}

// RecordTransaction increments the transaction counter
func RecordTransaction() {
	if GlobalMetrics != nil {
		GlobalMetrics.TransactionsTotal.Inc()
	}
}

// RecordCircuitTrip increments the circuit trip counter
func RecordCircuitTrip() {
	if GlobalMetrics != nil {
		GlobalMetrics.CircuitTripsTotal.Inc()
	}
}

// RecordRateLimitHit increments the rate limit counter
func RecordRateLimitHit() {
	if GlobalMetrics != nil {
		GlobalMetrics.RateLimitHitsTotal.Inc()
	}
}

// RecordComplianceBlock increments the compliance block counter
func RecordComplianceBlock() {
	if GlobalMetrics != nil {
		GlobalMetrics.ComplianceBlocksTotal.Inc()
	}
}

// SetGlobalPauseActive sets the global pause status
func SetGlobalPauseActive(active bool) {
	if GlobalMetrics != nil {
		if active {
			GlobalMetrics.GlobalPauseActive.Set(1)
		} else {
			GlobalMetrics.GlobalPauseActive.Set(0)
		}
	}
}

// RecordSettlement records a settlement
func RecordSettlement(volume float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.SettlementsTotal.Inc()
		GlobalMetrics.SettlementVolume.Add(volume)
	}
}

// SetActiveEscrows sets the number of active escrows
func SetActiveEscrows(count float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.ActiveEscrows.Set(count)
	}
}

// SetActiveChannels sets the number of active payment channels
func SetActiveChannels(count float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.ActiveChannels.Set(count)
	}
}

// RecordPriceUpdate records a price update
func RecordPriceUpdate() {
	if GlobalMetrics != nil {
		GlobalMetrics.PriceUpdatesTotal.Inc()
	}
}

// SetStalePrices sets the number of stale prices
func SetStalePrices(count float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.StalePrices.Set(count)
	}
}

// SetCollateralMetrics sets the collateral-related metrics
func SetCollateralMetrics(totalCollateral, totalDebt, ratio float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.TotalCollateral.Set(totalCollateral)
		GlobalMetrics.TotalDebt.Set(totalDebt)
		GlobalMetrics.CollateralRatio.Set(ratio)
	}
}

// SetStablecoinSupply sets the stablecoin supply metric
func SetStablecoinSupply(supply float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.StablecoinSupply.Set(supply)
	}
}

// RecordModuleError records an error for a module
func RecordModuleError(module string, errorRate float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.ModuleErrorRate.WithLabelValues(module).Set(errorRate)
	}
}

// RecordModuleLatency records latency for a module operation
func RecordModuleLatency(module, operation string, latencySeconds float64) {
	if GlobalMetrics != nil {
		GlobalMetrics.ModuleLatency.WithLabelValues(module, operation).Observe(latencySeconds)
	}
}
