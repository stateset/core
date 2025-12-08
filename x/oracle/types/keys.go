package types

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey is the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message routing key for the oracle module
	RouterKey = ModuleName

	// QuerierRoute is the legacy querier route for the oracle module
	QuerierRoute = ModuleName

	// Event types emitted by the oracle module.
	EventTypePriceUpdated = "price_updated"

	// Attribute keys for oracle events.
	AttributeKeyDenom  = "denom"
	AttributeKeyPrice  = "price"
	AttributeKeySource = "source"
)

var (
	PriceKeyPrefix = []byte{0x01}
)

func PriceStoreKey(denom string) []byte {
	key := append([]byte{}, PriceKeyPrefix...)
	return append(key, []byte(denom)...)
}
