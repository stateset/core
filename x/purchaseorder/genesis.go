package purchaseorder

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/purchaseorder/keeper"
	"github.com/stateset/core/x/purchaseorder/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the purchaseorder
	for _, elem := range genState.PurchaseorderList {
		k.SetPurchaseorder(ctx, elem)
	}

	// Set purchaseorder count
	k.SetPurchaseorderCount(ctx, genState.PurchaseorderCount)
	// Set all the sentPurchaseorder
	for _, elem := range genState.SentPurchaseorderList {
		k.SetSentPurchaseorder(ctx, elem)
	}

	// Set sentPurchaseorder count
	k.SetSentPurchaseorderCount(ctx, genState.SentPurchaseorderCount)
	// Set all the timedoutPurchaseorder
	for _, elem := range genState.TimedoutPurchaseorderList {
		k.SetTimedoutPurchaseorder(ctx, elem)
	}

	// Set timedoutPurchaseorder count
	k.SetTimedoutPurchaseorderCount(ctx, genState.TimedoutPurchaseorderCount)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.PurchaseorderList = k.GetAllPurchaseorder(ctx)
	genesis.PurchaseorderCount = k.GetPurchaseorderCount(ctx)
	genesis.SentPurchaseorderList = k.GetAllSentPurchaseorder(ctx)
	genesis.SentPurchaseorderCount = k.GetSentPurchaseorderCount(ctx)
	genesis.TimedoutPurchaseorderList = k.GetAllTimedoutPurchaseorder(ctx)
	genesis.TimedoutPurchaseorderCount = k.GetTimedoutPurchaseorderCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
