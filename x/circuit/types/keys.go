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

// CircuitStatus represents the state of a circuit breaker
type CircuitStatus uint8

const (
	// CircuitClosed means the circuit is functioning normally
	CircuitClosed CircuitStatus = iota
	// CircuitOpen means the circuit is tripped and operations are blocked
	CircuitOpen
	// CircuitHalfOpen means the circuit is in recovery mode
	CircuitHalfOpen
)

func (s CircuitStatus) String() string {
	switch s {
	case CircuitClosed:
		return "closed"
	case CircuitOpen:
		return "open"
	case CircuitHalfOpen:
		return "half_open"
	default:
		return "unknown"
	}
}
