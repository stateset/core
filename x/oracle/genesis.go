package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/stateset/core/x/oracle/keeper"
	"github.com/stateset/core/x/oracle/types"
)

// InitGenesis initializes the oracle module's state from a provided genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set module parameters
	k.SetParams(ctx, genState.Params)
	
	// Set oracle providers
	for _, provider := range genState.OracleProviders {
		if err := k.SetOracleProvider(ctx, provider); err != nil {
			panic(err)
		}
	}
	
	// Set price feeds
	for _, feed := range genState.PriceFeeds {
		if err := k.SetPriceFeed(ctx, feed); err != nil {
			panic(err)
		}
	}
	
	// Set aggregated prices
	for _, aggregated := range genState.AggregatedPrices {
		k.SetAggregatedPrice(ctx, aggregated)
	}
	
	// Set price histories
	for _, history := range genState.PriceHistories {
		k.SetPriceHistory(ctx, history)
	}
}

// ExportGenesis returns the oracle module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:           k.GetParams(ctx),
		OracleProviders:  k.GetAllOracleProviders(ctx),
		PriceFeeds:       k.GetAllPriceFeeds(ctx),
		AggregatedPrices: k.GetAllAggregatedPrices(ctx),
		PriceHistories:   k.GetAllPriceHistories(ctx),
	}
}