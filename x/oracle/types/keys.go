package types

const (
	ModuleName = "oracle"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey = "mem_oracle"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PriceFeedKey = "PriceFeed/value/"
	PriceFeedCountKey = "PriceFeed/count/"
	OracleProviderKey = "OracleProvider/value/"
	AggregatedPriceKey = "AggregatedPrice/value/"
	PriceHistoryKey = "PriceHistory/value/"
)