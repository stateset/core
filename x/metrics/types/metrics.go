package types

import (
	"context"
	"encoding/json"
	"time"

	sdkmath "cosmossdk.io/math"
)

// QueryServer defines the interface for metrics queries
type QueryServer interface {
	SystemMetrics(ctx context.Context, req *QuerySystemMetricsRequest) (*QuerySystemMetricsResponse, error)
	Counter(ctx context.Context, req *QueryCounterRequest) (*QueryCounterResponse, error)
	Gauge(ctx context.Context, req *QueryGaugeRequest) (*QueryGaugeResponse, error)
	ModuleHealth(ctx context.Context, req *QueryModuleHealthRequest) (*QueryModuleHealthResponse, error)
	Alerts(ctx context.Context, req *QueryAlertsRequest) (*QueryAlertsResponse, error)
	PrometheusMetrics(ctx context.Context, req *QueryPrometheusMetricsRequest) (*QueryPrometheusMetricsResponse, error)
}

type QuerySystemMetricsRequest struct{}
func (*QuerySystemMetricsRequest) Reset()         {}
func (*QuerySystemMetricsRequest) String() string { return "QuerySystemMetricsRequest" }
func (*QuerySystemMetricsRequest) ProtoMessage()  {}

type QuerySystemMetricsResponse struct {
	Metrics *SystemMetrics `json:"metrics"`
}
func (*QuerySystemMetricsResponse) Reset()         {}
func (*QuerySystemMetricsResponse) String() string { return "QuerySystemMetricsResponse" }
func (*QuerySystemMetricsResponse) ProtoMessage()  {}

type QueryCounterRequest struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels,omitempty"`
}
func (*QueryCounterRequest) Reset()         {}
func (*QueryCounterRequest) String() string { return "QueryCounterRequest" }
func (*QueryCounterRequest) ProtoMessage()  {}

type QueryCounterResponse struct {
	Counter *Counter `json:"counter"`
}
func (*QueryCounterResponse) Reset()         {}
func (*QueryCounterResponse) String() string { return "QueryCounterResponse" }
func (*QueryCounterResponse) ProtoMessage()  {}

type QueryGaugeRequest struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels,omitempty"`
}
func (*QueryGaugeRequest) Reset()         {}
func (*QueryGaugeRequest) String() string { return "QueryGaugeRequest" }
func (*QueryGaugeRequest) ProtoMessage()  {}

type QueryGaugeResponse struct {
	Gauge *Gauge `json:"gauge"`
}
func (*QueryGaugeResponse) Reset()         {}
func (*QueryGaugeResponse) String() string { return "QueryGaugeResponse" }
func (*QueryGaugeResponse) ProtoMessage()  {}

type QueryModuleHealthRequest struct {
	Module string `json:"module"`
}
func (*QueryModuleHealthRequest) Reset()         {}
func (*QueryModuleHealthRequest) String() string { return "QueryModuleHealthRequest" }
func (*QueryModuleHealthRequest) ProtoMessage()  {}

type QueryModuleHealthResponse struct {
	Health map[string]ModuleHealth `json:"health"`
}
func (*QueryModuleHealthResponse) Reset()         {}
func (*QueryModuleHealthResponse) String() string { return "QueryModuleHealthResponse" }
func (*QueryModuleHealthResponse) ProtoMessage()  {}

type QueryAlertsRequest struct {
	Active bool `json:"active"`
}
func (*QueryAlertsRequest) Reset()         {}
func (*QueryAlertsRequest) String() string { return "QueryAlertsRequest" }
func (*QueryAlertsRequest) ProtoMessage()  {}

type QueryAlertsResponse struct {
	Alerts       []Alert       `json:"alerts"`
	AlertConfigs []AlertConfig `json:"alert_configs,omitempty"`
}
func (*QueryAlertsResponse) Reset()         {}
func (*QueryAlertsResponse) String() string { return "QueryAlertsResponse" }
func (*QueryAlertsResponse) ProtoMessage()  {}

// QueryPrometheusMetricsRequest is the request type for prometheus metrics
type QueryPrometheusMetricsRequest struct{}
func (*QueryPrometheusMetricsRequest) Reset()         {}
func (*QueryPrometheusMetricsRequest) String() string { return "QueryPrometheusMetricsRequest" }
func (*QueryPrometheusMetricsRequest) ProtoMessage()  {}

// QueryPrometheusMetricsResponse is the response type for prometheus metrics
type QueryPrometheusMetricsResponse struct {
	Metrics string `json:"metrics"`
}
func (*QueryPrometheusMetricsResponse) Reset()         {}
func (*QueryPrometheusMetricsResponse) String() string { return "QueryPrometheusMetricsResponse" }
func (*QueryPrometheusMetricsResponse) ProtoMessage()  {}

// MetricType represents the type of metric
type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
)

// Metric represents a single metric data point
type Metric struct {
	Name      string            `json:"name"`
	Type      MetricType        `json:"type"`
	Value     sdkmath.LegacyDec `json:"value"`
	Labels    map[string]string `json:"labels,omitempty"`
	Timestamp time.Time         `json:"timestamp"`
	Height    int64             `json:"height"`
}

// Counter represents a monotonically increasing counter
type Counter struct {
	Name   string            `json:"name"`
	Value  uint64            `json:"value"`
	Labels map[string]string `json:"labels,omitempty"`
}

// Gauge represents a value that can go up and down
type Gauge struct {
	Name      string            `json:"name"`
	Value     sdkmath.LegacyDec `json:"value"`
	Labels    map[string]string `json:"labels,omitempty"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// HistogramBucket represents a single bucket in a histogram
type HistogramBucket struct {
	LowerBound sdkmath.LegacyDec `json:"lower_bound"`
	UpperBound sdkmath.LegacyDec `json:"upper_bound"`
	Count      uint64            `json:"count"`
}

// Histogram represents a distribution of values
type Histogram struct {
	Name      string            `json:"name"`
	Buckets   []HistogramBucket `json:"buckets"`
	Sum       sdkmath.LegacyDec `json:"sum"`
	Count     uint64            `json:"count"`
	Labels    map[string]string `json:"labels,omitempty"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// SystemMetrics represents overall system metrics
type SystemMetrics struct {
	// Block metrics
	LastBlockHeight   int64         `json:"last_block_height"`
	LastBlockTime     time.Time     `json:"last_block_time"`
	AverageBlockTime  time.Duration `json:"average_block_time"`
	TransactionsTotal uint64        `json:"transactions_total"`

	// Module health
	ModuleHealth map[string]ModuleHealth `json:"module_health"`

	// Economic metrics
	TotalCollateralValue sdkmath.Int `json:"total_collateral_value"`
	TotalDebtValue       sdkmath.Int `json:"total_debt_value"`
	SystemCollateralRatio sdkmath.LegacyDec `json:"system_collateral_ratio"`

	// Settlement metrics
	TotalSettlements      uint64      `json:"total_settlements"`
	TotalSettlementVolume sdkmath.Int `json:"total_settlement_volume"`
	ActiveEscrows         uint64      `json:"active_escrows"`
	ActiveChannels        uint64      `json:"active_channels"`

	// Oracle metrics
	PricesUpdated   uint64 `json:"prices_updated"`
	StalePriceCount uint64 `json:"stale_price_count"`

	// Security metrics
	CircuitTrips     uint64 `json:"circuit_trips"`
	RateLimitHits    uint64 `json:"rate_limit_hits"`
	ComplianceBlocks uint64 `json:"compliance_blocks"`
}

// ModuleHealth represents the health status of a module
type ModuleHealth struct {
	Module       string  `json:"module"`
	Status       string  `json:"status"` // healthy, degraded, unhealthy
	ErrorRate    float64 `json:"error_rate"`
	Latency      float64 `json:"latency_ms"`
	LastError    string  `json:"last_error,omitempty"`
	LastErrorAt  time.Time `json:"last_error_at,omitempty"`
	Transactions uint64  `json:"transactions"`
}

// AlertConfig defines an alert configuration
type AlertConfig struct {
	Name       string            `json:"name"`
	MetricName string            `json:"metric_name"`
	Condition  AlertCondition    `json:"condition"`
	Threshold  sdkmath.LegacyDec `json:"threshold"`
	Duration   time.Duration     `json:"duration"`
	Severity   AlertSeverity     `json:"severity"`
	Enabled    bool              `json:"enabled"`
}

// AlertCondition represents the condition for an alert
type AlertCondition string

const (
	AlertConditionGreaterThan AlertCondition = "gt"
	AlertConditionLessThan    AlertCondition = "lt"
	AlertConditionEquals      AlertCondition = "eq"
	AlertConditionNotEquals   AlertCondition = "ne"
)

// AlertSeverity represents the severity of an alert
type AlertSeverity string

const (
	AlertSeverityInfo     AlertSeverity = "info"
	AlertSeverityWarning  AlertSeverity = "warning"
	AlertSeverityCritical AlertSeverity = "critical"
)

// Alert represents a triggered alert
type Alert struct {
	ID          string            `json:"id"`
	ConfigName  string            `json:"config_name"`
	MetricName  string            `json:"metric_name"`
	Value       sdkmath.LegacyDec `json:"value"`
	Threshold   sdkmath.LegacyDec `json:"threshold"`
	Severity    AlertSeverity     `json:"severity"`
	Message     string            `json:"message"`
	TriggeredAt time.Time         `json:"triggered_at"`
	ResolvedAt  time.Time         `json:"resolved_at,omitempty"`
	Resolved    bool              `json:"resolved"`
}

// PerformanceStats represents performance statistics
type PerformanceStats struct {
	Module         string  `json:"module"`
	Operation      string  `json:"operation"`
	TotalCalls     uint64  `json:"total_calls"`
	TotalDuration  float64 `json:"total_duration_ms"`
	AverageDuration float64 `json:"average_duration_ms"`
	MinDuration    float64 `json:"min_duration_ms"`
	MaxDuration    float64 `json:"max_duration_ms"`
	ErrorCount     uint64  `json:"error_count"`
	LastCalled     time.Time `json:"last_called"`
}

// DefaultSystemMetrics returns default system metrics
func DefaultSystemMetrics() SystemMetrics {
	return SystemMetrics{
		ModuleHealth:          make(map[string]ModuleHealth),
		TotalCollateralValue:  sdkmath.ZeroInt(),
		TotalDebtValue:        sdkmath.ZeroInt(),
		SystemCollateralRatio: sdkmath.LegacyZeroDec(),
		TotalSettlementVolume: sdkmath.ZeroInt(),
	}
}

// DefaultAlertConfigs returns default alert configurations
func DefaultAlertConfigs() []AlertConfig {
	return []AlertConfig{
		{
			Name:       "high_error_rate",
			MetricName: "error_rate",
			Condition:  AlertConditionGreaterThan,
			Threshold:  sdkmath.LegacyNewDecWithPrec(5, 2), // 5%
			Duration:   5 * time.Minute,
			Severity:   AlertSeverityWarning,
			Enabled:    true,
		},
		{
			Name:       "low_collateral_ratio",
			MetricName: "system_collateral_ratio",
			Condition:  AlertConditionLessThan,
			Threshold:  sdkmath.LegacyNewDecWithPrec(15, 1), // 1.5
			Duration:   1 * time.Minute,
			Severity:   AlertSeverityCritical,
			Enabled:    true,
		},
		{
			Name:       "stale_prices",
			MetricName: "stale_price_count",
			Condition:  AlertConditionGreaterThan,
			Threshold:  sdkmath.LegacyNewDec(0),
			Duration:   10 * time.Minute,
			Severity:   AlertSeverityWarning,
			Enabled:    true,
		},
		{
			Name:       "high_rate_limit_hits",
			MetricName: "rate_limit_hits",
			Condition:  AlertConditionGreaterThan,
			Threshold:  sdkmath.LegacyNewDec(100),
			Duration:   1 * time.Minute,
			Severity:   AlertSeverityInfo,
			Enabled:    true,
		},
	}
}

// GenesisState defines the metrics module's genesis state
type GenesisState struct {
	SystemMetrics SystemMetrics `json:"system_metrics"`
	AlertConfigs  []AlertConfig `json:"alert_configs"`
	Alerts        []Alert       `json:"alerts"`
}

// Proto message methods for GenesisState
func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return "" }
func (m *GenesisState) ProtoMessage()  {}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SystemMetrics: DefaultSystemMetrics(),
		AlertConfigs:  DefaultAlertConfigs(),
		Alerts:        []Alert{},
	}
}

// Marshal helpers
func (m SystemMetrics) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func UnmarshalSystemMetrics(bz []byte) (SystemMetrics, error) {
	var m SystemMetrics
	err := json.Unmarshal(bz, &m)
	return m, err
}

func (c Counter) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func UnmarshalCounter(bz []byte) (Counter, error) {
	var c Counter
	err := json.Unmarshal(bz, &c)
	return c, err
}

func (g Gauge) Marshal() ([]byte, error) {
	return json.Marshal(g)
}

func UnmarshalGauge(bz []byte) (Gauge, error) {
	var g Gauge
	err := json.Unmarshal(bz, &g)
	return g, err
}
