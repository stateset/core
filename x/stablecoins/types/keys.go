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