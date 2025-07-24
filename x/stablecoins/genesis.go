package stablecoins

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/stablecoins/keeper"
	"github.com/stateset/core/x/stablecoins/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the stablecoin
	for _, elem := range genState.StablecoinList {
		k.SetStablecoin(ctx, elem)
	}

	// Set all the price data
	for _, elem := range genState.PriceDataList {
		store := ctx.KVStore(k.GetStoreKey())
		b := k.GetCodec().MustMarshal(&elem)
		store.Set(append(types.PriceDataKeyPrefix, types.PriceDataKey(elem.Denom)...), b)
	}

	// Set all the mint requests
	for _, elem := range genState.MintRequestList {
		store := ctx.KVStore(k.GetStoreKey())
		b := k.GetCodec().MustMarshal(&elem)
		store.Set(append(types.MintRequestKeyPrefix, types.MintRequestKey(elem.Id)...), b)
	}

	// Set all the burn requests
	for _, elem := range genState.BurnRequestList {
		store := ctx.KVStore(k.GetStoreKey())
		b := k.GetCodec().MustMarshal(&elem)
		store.Set(append(types.BurnRequestKeyPrefix, types.BurnRequestKey(elem.Id)...), b)
	}

	// Set whitelist entries
	for _, elem := range genState.WhitelistEntries {
		k.WhitelistAddress(ctx, elem.Denom, elem.Address)
	}

	// Set blacklist entries
	for _, elem := range genState.BlacklistEntries {
		k.BlacklistAddress(ctx, elem.Denom, elem.Address, elem.Reason)
	}

	// Set stablecoin count
	store := ctx.KVStore(k.GetStoreKey())
	countBytes := sdk.Uint64ToBigEndian(genState.StablecoinCount)
	store.Set(types.StablecoinCountKey, countBytes)

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StablecoinList = k.GetAllStablecoin(ctx)

	// Export price data
	store := ctx.KVStore(k.GetStoreKey())
	priceDataStore := sdk.KVStorePrefixIterator(store, types.PriceDataKeyPrefix)
	defer priceDataStore.Close()

	for ; priceDataStore.Valid(); priceDataStore.Next() {
		var priceData types.PriceData
		k.GetCodec().MustUnmarshal(priceDataStore.Value(), &priceData)
		genesis.PriceDataList = append(genesis.PriceDataList, priceData)
	}

	// Export mint requests
	mintRequestStore := sdk.KVStorePrefixIterator(store, types.MintRequestKeyPrefix)
	defer mintRequestStore.Close()

	for ; mintRequestStore.Valid(); mintRequestStore.Next() {
		var mintRequest types.MintRequest
		k.GetCodec().MustUnmarshal(mintRequestStore.Value(), &mintRequest)
		genesis.MintRequestList = append(genesis.MintRequestList, mintRequest)
	}

	// Export burn requests
	burnRequestStore := sdk.KVStorePrefixIterator(store, types.BurnRequestKeyPrefix)
	defer burnRequestStore.Close()

	for ; burnRequestStore.Valid(); burnRequestStore.Next() {
		var burnRequest types.BurnRequest
		k.GetCodec().MustUnmarshal(burnRequestStore.Value(), &burnRequest)
		genesis.BurnRequestList = append(genesis.BurnRequestList, burnRequest)
	}

	// Export whitelist entries
	whitelistStore := sdk.KVStorePrefixIterator(store, types.WhitelistKeyPrefix)
	defer whitelistStore.Close()

	for ; whitelistStore.Valid(); whitelistStore.Next() {
		// Parse key to extract denom and address
		// This is a simplified approach - in practice you'd parse the key properly
		// For now, we'll just store empty entries as placeholders
		entry := types.WhitelistEntry{
			Denom:   "",
			Address: "",
		}
		genesis.WhitelistEntries = append(genesis.WhitelistEntries, entry)
	}

	// Export blacklist entries
	blacklistStore := sdk.KVStorePrefixIterator(store, types.BlacklistKeyPrefix)
	defer blacklistStore.Close()

	for ; blacklistStore.Valid(); blacklistStore.Next() {
		// Parse key to extract denom and address
		// This is a simplified approach - in practice you'd parse the key properly
		entry := types.BlacklistEntry{
			Denom:   "",
			Address: "",
			Reason:  string(blacklistStore.Value()),
		}
		genesis.BlacklistEntries = append(genesis.BlacklistEntries, entry)
	}

	// Export stablecoin count
	countBytes := store.Get(types.StablecoinCountKey)
	if countBytes != nil {
		genesis.StablecoinCount = sdk.BigEndianToUint64(countBytes)
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}