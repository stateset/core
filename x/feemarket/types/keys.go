package types

const (
	// ModuleName defines the module name
	ModuleName = "feemarket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for the module
	RouterKey = ModuleName

	// MemStoreKey is the in-memory store key
	MemStoreKey = "mem_feemarket"
)

var (
	ParamsKey    = []byte{0x01}
	BaseFeeKey   = []byte{0x02}
	HistoryKey   = []byte{0x03}
	CircuitKey   = []byte{0x04}
	LatestGasKey = []byte{0x05}
)
