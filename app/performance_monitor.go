package app

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// Block metrics
	blockProcessingTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "stateset_block_processing_seconds",
		Help:    "Time taken to process a block",
		Buckets: prometheus.DefBuckets,
	}, []string{"phase"})

	blockHeight = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_block_height",
		Help: "Current block height",
	})

	blockSize = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "stateset_block_size_bytes",
		Help:    "Size of blocks in bytes",
		Buckets: prometheus.ExponentialBuckets(1024, 2, 20), // 1KB to 1GB
	})

	// Transaction metrics
	transactionCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "stateset_transactions_total",
		Help: "Total number of transactions processed",
	}, []string{"status", "type"})

	transactionProcessingTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "stateset_transaction_processing_seconds",
		Help:    "Time taken to process individual transactions",
		Buckets: prometheus.DefBuckets,
	}, []string{"type"})

	gasUsed = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "stateset_gas_used",
		Help:    "Gas used per transaction",
		Buckets: prometheus.ExponentialBuckets(1000, 10, 10),
	}, []string{"type"})

	// State metrics
	stateSize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stateset_state_size_bytes",
		Help: "Size of state data in bytes",
	}, []string{"module"})

	queryLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "stateset_query_latency_seconds",
		Help:    "Latency of state queries",
		Buckets: prometheus.DefBuckets,
	}, []string{"query_type"})

	// Mempool metrics
	mempoolSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_mempool_size",
		Help: "Number of transactions in mempool",
	})

	mempoolBytes = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_mempool_bytes",
		Help: "Total size of mempool in bytes",
	})

	// Consensus metrics
	validatorCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_validator_count",
		Help: "Number of active validators",
	})

	votingPower = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stateset_voting_power",
		Help: "Voting power distribution",
	}, []string{"validator"})

	// Network metrics
	peerCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_peer_count",
		Help: "Number of connected peers",
	})

	// Fee metrics
	feeAmount = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "stateset_fee_amount",
		Help:    "Transaction fee amounts",
		Buckets: prometheus.ExponentialBuckets(1, 10, 10),
	}, []string{"denom"})
)

// PerformanceMonitor tracks blockchain performance metrics
type PerformanceMonitor struct {
	enabled bool
}

// NewPerformanceMonitor creates a new performance monitor
func NewPerformanceMonitor(enabled bool) *PerformanceMonitor {
	return &PerformanceMonitor{
		enabled: enabled,
	}
}

// RecordBlockMetrics records block-level metrics
func (pm *PerformanceMonitor) RecordBlockMetrics(ctx sdk.Context, processingTime time.Duration, blockSizeBytes int64) {
	if !pm.enabled {
		return
	}

	blockHeight.Set(float64(ctx.BlockHeight()))
	blockProcessingTime.WithLabelValues("total").Observe(processingTime.Seconds())
	blockSize.Observe(float64(blockSizeBytes))

	telemetry.SetGauge(float32(ctx.BlockHeight()), "block", "height")
	telemetry.SetGauge(float32(blockSizeBytes), "block", "size")
}

// RecordTransactionMetrics records transaction-level metrics
func (pm *PerformanceMonitor) RecordTransactionMetrics(ctx sdk.Context, tx sdk.Tx, status string, processingTime time.Duration) {
	if !pm.enabled {
		return
	}

	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		msgType := sdk.MsgTypeURL(msg)
		transactionCount.WithLabelValues(status, msgType).Inc()
		transactionProcessingTime.WithLabelValues(msgType).Observe(processingTime.Seconds())

		if feeTx, ok := tx.(sdk.FeeTx); ok {
			gasUsed.WithLabelValues(msgType).Observe(float64(feeTx.GetGas()))
			
			fees := feeTx.GetFee()
			for _, fee := range fees {
				feeAmount.WithLabelValues(fee.Denom).Observe(float64(fee.Amount.Int64()))
			}
		}
	}
}

// RecordQueryMetrics records query performance metrics
func (pm *PerformanceMonitor) RecordQueryMetrics(queryType string, latency time.Duration) {
	if !pm.enabled {
		return
	}

	queryLatency.WithLabelValues(queryType).Observe(latency.Seconds())
}

// RecordMempoolMetrics records mempool metrics
func (pm *PerformanceMonitor) RecordMempoolMetrics(size int, bytes int64) {
	if !pm.enabled {
		return
	}

	mempoolSize.Set(float64(size))
	mempoolBytes.Set(float64(bytes))
}

// RecordValidatorMetrics records validator-related metrics
func (pm *PerformanceMonitor) RecordValidatorMetrics(validators []ValidatorInfo) {
	if !pm.enabled {
		return
	}

	validatorCount.Set(float64(len(validators)))
	
	for _, val := range validators {
		votingPower.WithLabelValues(val.Address).Set(float64(val.VotingPower))
	}
}

// RecordStateMetrics records state size metrics
func (pm *PerformanceMonitor) RecordStateMetrics(module string, sizeBytes int64) {
	if !pm.enabled {
		return
	}

	stateSize.WithLabelValues(module).Set(float64(sizeBytes))
}

// RecordPeerCount records the number of connected peers
func (pm *PerformanceMonitor) RecordPeerCount(count int) {
	if !pm.enabled {
		return
	}

	peerCount.Set(float64(count))
}

// ValidatorInfo contains validator information for metrics
type ValidatorInfo struct {
	Address     string
	VotingPower int64
}

// BlockProcessingTimer helps time different phases of block processing
type BlockProcessingTimer struct {
	monitor *PerformanceMonitor
	start   time.Time
	phase   string
}

// NewBlockProcessingTimer creates a new timer for block processing
func (pm *PerformanceMonitor) NewBlockProcessingTimer(phase string) *BlockProcessingTimer {
	return &BlockProcessingTimer{
		monitor: pm,
		start:   time.Now(),
		phase:   phase,
	}
}

// End records the elapsed time for the block processing phase
func (t *BlockProcessingTimer) End() {
	if t.monitor.enabled {
		elapsed := time.Since(t.start)
		blockProcessingTime.WithLabelValues(t.phase).Observe(elapsed.Seconds())
	}
}

// WithMetrics wraps a function to record its execution time
func (pm *PerformanceMonitor) WithMetrics(ctx context.Context, metricName string, fn func() error) error {
	if !pm.enabled {
		return fn()
	}

	start := time.Now()
	err := fn()
	duration := time.Since(start)

	status := "success"
	if err != nil {
		status = "error"
	}

	telemetry.SetGaugeWithLabels(
		[]string{"operation", "duration"},
		float32(duration.Milliseconds()),
		[]telemetry.Label{
			{Name: "operation", Value: metricName},
			{Name: "status", Value: status},
		},
	)

	return err
}

// GetMetricsHandler returns a Prometheus metrics handler
func (pm *PerformanceMonitor) GetMetricsHandler() http.Handler {
	return promhttp.Handler()
}