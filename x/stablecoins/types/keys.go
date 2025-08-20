package types

const (
	// ModuleName defines the module name
	ModuleName = "stablecoins"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stablecoins"
)

var (
	StablecoinKeyPrefix        = []byte{0x01}
	PriceDataKeyPrefix         = []byte{0x02}
	MintRequestKeyPrefix       = []byte{0x03}
	BurnRequestKeyPrefix       = []byte{0x04}
	WhitelistKeyPrefix         = []byte{0x05}
	BlacklistKeyPrefix         = []byte{0x06}
	StablecoinCountKey         = []byte{0x07}
	TotalMarketCapKey          = []byte{0x08}
	EscrowKeyPrefix            = []byte{0x09}
	
	// Working capital prefixes
	WorkingCapitalLoanKeyPrefix    = []byte{0x10}
	WorkingCapitalRequestKeyPrefix = []byte{0x11}
	WorkingCapitalPoolKeyPrefix    = []byte{0x12}
	RepaymentScheduleKeyPrefix     = []byte{0x13}
	CreditLineKeyPrefix            = []byte{0x14}
	BusinessProfileKeyPrefix       = []byte{0x15}
)

// StablecoinKey returns the store key to retrieve a Stablecoin from the index fields
func StablecoinKey(denom string) []byte {
	var key []byte
	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)
	return key
}

// PriceDataKey returns the store key for price data
func PriceDataKey(denom string) []byte {
	var key []byte
	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)
	return key
}

// MintRequestKey returns the store key for mint requests
func MintRequestKey(id string) []byte {
	var key []byte
	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)
	return key
}

// BurnRequestKey returns the store key for burn requests
func BurnRequestKey(id string) []byte {
	var key []byte
	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)
	return key
}

// WhitelistKey returns the store key for whitelist entries
func WhitelistKey(denom, address string) []byte {
	var key []byte
	denomBytes := []byte(denom)
	addressBytes := []byte(address)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)
	return key
}

// BlacklistKey returns the store key for blacklist entries
func BlacklistKey(denom, address string) []byte {
	var key []byte
	denomBytes := []byte(denom)
	addressBytes := []byte(address)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)
	return key
}

// EscrowKey returns the store key for escrow entries
func EscrowKey(orderId string) []byte {
	var key []byte
	orderIdBytes := []byte(orderId)
	key = append(key, orderIdBytes...)
	key = append(key, []byte("/")...)
	return key
}

// WorkingCapitalLoanKey returns the store key for working capital loans
func WorkingCapitalLoanKey(loanId string) []byte {
	var key []byte
	loanIdBytes := []byte(loanId)
	key = append(key, loanIdBytes...)
	key = append(key, []byte("/")...)
	return key
}

// WorkingCapitalRequestKey returns the store key for working capital requests
func WorkingCapitalRequestKey(requestId string) []byte {
	var key []byte
	requestIdBytes := []byte(requestId)
	key = append(key, requestIdBytes...)
	key = append(key, []byte("/")...)
	return key
}

// WorkingCapitalPoolKey returns the store key for capital pools
func WorkingCapitalPoolKey(poolId string) []byte {
	var key []byte
	poolIdBytes := []byte(poolId)
	key = append(key, poolIdBytes...)
	key = append(key, []byte("/")...)
	return key
}

// RepaymentScheduleKey returns the store key for repayment schedules
func RepaymentScheduleKey(loanId string, installmentNumber int) []byte {
	var key []byte
	loanIdBytes := []byte(loanId)
	key = append(key, loanIdBytes...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(string(rune(installmentNumber)))...)
	key = append(key, []byte("/")...)
	return key
}

// CreditLineKey returns the store key for credit lines
func CreditLineKey(creditLineId string) []byte {
	var key []byte
	creditLineIdBytes := []byte(creditLineId)
	key = append(key, creditLineIdBytes...)
	key = append(key, []byte("/")...)
	return key
}

// BusinessProfileKey returns the store key for business profiles
func BusinessProfileKey(businessId string) []byte {
	var key []byte
	businessIdBytes := []byte(businessId)
	key = append(key, businessIdBytes...)
	key = append(key, []byte("/")...)
	return key
}