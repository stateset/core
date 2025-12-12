package types

const (
	// ModuleName defines the module name
	ModuleName = "circuit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName
)

var (
	// CircuitStateKey is the key for storing global circuit state
	CircuitStateKey = []byte{0x01}

	// ModuleCircuitKeyPrefix is the prefix for module-specific circuit states
	ModuleCircuitKeyPrefix = []byte{0x02}

	// RateLimitKeyPrefix is the prefix for rate limit tracking
	RateLimitKeyPrefix = []byte{0x03}

	// RateLimitConfigKeyPrefix is the prefix for rate limit configurations
	RateLimitConfigKeyPrefix = []byte{0x04}
)

const (
	CircuitClosed   = CircuitStatus_CIRCUIT_CLOSED
	CircuitOpen     = CircuitStatus_CIRCUIT_OPEN
	CircuitHalfOpen = CircuitStatus_CIRCUIT_HALF_OPEN
)
