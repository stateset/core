package orders

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/orders/keeper"
	"github.com/stateset/core/x/orders/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the orders
	for _, elem := range genState.Orders {
		k.SetOrder(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.Orders = k.GetAllOrders(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}