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

// Event types
const (
	EventTypeUpdateParams   = "update_params"
	EventTypeFeeValidation  = "fee_validation"
)

// Attribute keys
const (
	AttributeKeyAuthority      = "authority"
	AttributeKeyEnabled        = "enabled"
	AttributeKeyMinBaseFee     = "min_base_fee"
	AttributeKeyMaxBaseFee     = "max_base_fee"
	AttributeKeyBaseFee        = "base_fee"
	AttributeKeyGasLimit       = "gas_limit"
	AttributeKeyMinRequiredFee = "min_required_fee"
	AttributeKeyActualFee      = "actual_fee"
)
