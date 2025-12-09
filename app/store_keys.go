package app

import (
	storetypes "cosmossdk.io/store/types"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	circuittypes "github.com/stateset/core/x/circuit/types"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	metricstypes "github.com/stateset/core/x/metrics/types"
	oracletypes "github.com/stateset/core/x/oracle/types"
	paymentstypes "github.com/stateset/core/x/payments/types"
	settlementtypes "github.com/stateset/core/x/settlement/types"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
	treasurytypes "github.com/stateset/core/x/treasury/types"
	zkpverifytypes "github.com/stateset/core/x/zkpverify/types"
)

// GetStoreKeys returns all the KVStoreKeys used in the application
func GetStoreKeys() []string {
	return []string{
		// Cosmos SDK core modules
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		consensusparamtypes.StoreKey,
		// IBC modules
		ibcexported.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
		// Stateset custom modules
		oracletypes.StoreKey,
		compliancetypes.StoreKey,
		treasurytypes.StoreKey,
		paymentstypes.StoreKey,
		stablecointypes.StoreKey,
		settlementtypes.StoreKey,
		circuittypes.StoreKey,
		metricstypes.StoreKey,
		zkpverifytypes.StoreKey,
	}
}

// CreateStoreKeys creates all the store keys for the application
func CreateStoreKeys() (
	map[string]*storetypes.KVStoreKey,
	map[string]*storetypes.TransientStoreKey,
	map[string]*storetypes.MemoryStoreKey,
) {
	keys := storetypes.NewKVStoreKeys(GetStoreKeys()...)
	tkeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	return keys, tkeys, memKeys
}

// GetCosmosStoreKeys returns store keys for Cosmos SDK core modules
func GetCosmosStoreKeys() []string {
	return []string{
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		consensusparamtypes.StoreKey,
	}
}

// GetIBCStoreKeys returns store keys for IBC modules
func GetIBCStoreKeys() []string {
	return []string{
		ibcexported.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
	}
}

// GetStatesetStoreKeys returns store keys for Stateset custom modules
func GetStatesetStoreKeys() []string {
	return []string{
		oracletypes.StoreKey,
		compliancetypes.StoreKey,
		treasurytypes.StoreKey,
		paymentstypes.StoreKey,
		stablecointypes.StoreKey,
		settlementtypes.StoreKey,
		circuittypes.StoreKey,
		metricstypes.StoreKey,
		zkpverifytypes.StoreKey,
	}
}
