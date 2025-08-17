package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "ssusd"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey = "mem_ssusd"
)

var (
	ParamsKey = []byte{0x01}
	
	CollateralPositionKeyPrefix  = []byte{0x02}
	AgentWalletKeyPrefix         = []byte{0x03}
	OraclePriceKeyPrefix         = []byte{0x04}
	LiquidationAuctionKeyPrefix  = []byte{0x05}
	StabilityPoolKey             = []byte{0x06}
	SystemMetricsKey             = []byte{0x07}
	NextPositionIDKey            = []byte{0x08}
	NextAuctionIDKey             = []byte{0x09}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func CollateralPositionKey(owner string) []byte {
	return append(CollateralPositionKeyPrefix, []byte(owner)...)
}

func AgentWalletKey(agentID string) []byte {
	return append(AgentWalletKeyPrefix, []byte(agentID)...)
}

func OraclePriceKey(asset string) []byte {
	return append(OraclePriceKeyPrefix, []byte(asset)...)
}

func LiquidationAuctionKey(auctionID string) []byte {
	return append(LiquidationAuctionKeyPrefix, []byte(auctionID)...)
}