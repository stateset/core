package types

const (
	// ModuleName defines the module name
	ModuleName = "purchaseorder"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_purchaseorder"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PurchaseorderKey      = "Purchaseorder-value-"
	PurchaseorderCountKey = "Purchaseorder-count-"
)

const (
	SentPurchaseorderKey      = "SentPurchaseorder-value-"
	SentPurchaseorderCountKey = "SentPurchaseorder-count-"
)

const (
	TimedoutPurchaseorderKey      = "TimedoutPurchaseorder-value-"
	TimedoutPurchaseorderCountKey = "TimedoutPurchaseorder-count-"
)
