package types

const (
	// ModuleName defines the module name
	ModuleName = "metrics"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName
)

var (
	// MetricsKeyPrefix is the prefix for metrics data
	MetricsKeyPrefix = []byte{0x01}

	// CounterKeyPrefix is the prefix for counter metrics
	CounterKeyPrefix = []byte{0x02}

	// GaugeKeyPrefix is the prefix for gauge metrics
	GaugeKeyPrefix = []byte{0x03}

	// HistogramKeyPrefix is the prefix for histogram metrics
	HistogramKeyPrefix = []byte{0x04}

	// AlertKeyPrefix is the prefix for alert configurations
	AlertKeyPrefix = []byte{0x05}
)
