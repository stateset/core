package types

import (
	"time"

	sdkmath "cosmossdk.io/math"
)

// MetricType represents the type of metric.
type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
)

// AlertCondition represents the condition for an alert.
type AlertCondition string

const (
	AlertConditionGreaterThan AlertCondition = "gt"
	AlertConditionLessThan    AlertCondition = "lt"
	AlertConditionEquals      AlertCondition = "eq"
	AlertConditionNotEquals   AlertCondition = "ne"
)

// AlertSeverity represents the severity of an alert.
type AlertSeverity string

const (
	AlertSeverityInfo     AlertSeverity = "info"
	AlertSeverityWarning  AlertSeverity = "warning"
	AlertSeverityCritical AlertSeverity = "critical"
)

// DefaultSystemMetrics returns default system metrics.
func DefaultSystemMetrics() SystemMetrics {
	return SystemMetrics{
		ModuleHealth:          make(map[string]*ModuleHealth),
		TotalCollateralValue:  sdkmath.ZeroInt(),
		TotalDebtValue:        sdkmath.ZeroInt(),
		SystemCollateralRatio: sdkmath.LegacyZeroDec(),
		TotalSettlementVolume: sdkmath.ZeroInt(),
	}
}

// DefaultAlertConfigs returns default alert configurations.
func DefaultAlertConfigs() []AlertConfig {
	return []AlertConfig{
		{
			Name:       "high_error_rate",
			MetricName: "transactions_failed",
			Condition:  AlertConditionGreaterThan,
			Threshold:  sdkmath.LegacyNewDecWithPrec(1, 1),
			Duration:   5 * time.Minute,
			Severity:   AlertSeverityWarning,
			Enabled:    true,
		},
		{
			Name:       "circuit_trips",
			MetricName: "circuit_trips",
			Condition:  AlertConditionGreaterThan,
			Threshold:  sdkmath.LegacyNewDec(10),
			Duration:   1 * time.Minute,
			Severity:   AlertSeverityCritical,
			Enabled:    true,
		},
		{
			Name:       "rate_limit_hits",
			MetricName: "rate_limit_hits",
			Condition:  AlertConditionGreaterThan,
			Threshold:  sdkmath.LegacyNewDec(100),
			Duration:   1 * time.Minute,
			Severity:   AlertSeverityInfo,
			Enabled:    true,
		},
	}
}

// GenesisState defines the metrics module's genesis state.
type GenesisState struct {
	SystemMetrics SystemMetrics `json:"system_metrics"`
	AlertConfigs  []AlertConfig `json:"alert_configs"`
	Alerts        []Alert       `json:"alerts"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return "" }
func (m *GenesisState) ProtoMessage()  {}

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SystemMetrics: DefaultSystemMetrics(),
		AlertConfigs:  DefaultAlertConfigs(),
		Alerts:        []Alert{},
	}
}

func UnmarshalSystemMetrics(bz []byte) (SystemMetrics, error) {
	var metrics SystemMetrics
	if err := metrics.Unmarshal(bz); err != nil {
		return SystemMetrics{}, err
	}
	return metrics, nil
}

func UnmarshalCounter(bz []byte) (Counter, error) {
	var counter Counter
	if err := counter.Unmarshal(bz); err != nil {
		return Counter{}, err
	}
	return counter, nil
}

func UnmarshalGauge(bz []byte) (Gauge, error) {
	var gauge Gauge
	if err := gauge.Unmarshal(bz); err != nil {
		return Gauge{}, err
	}
	return gauge, nil
}
