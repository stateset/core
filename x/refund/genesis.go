package refund

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/refund/keeper"
	"github.com/stateset/core/x/refund/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the refund
	for _, elem := range genState.RefundList {
		k.SetRefund(ctx, elem)
	}

	// Set refund count
	k.SetRefundCount(ctx, genState.RefundCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RefundList = k.GetAllRefund(ctx)
	genesis.RefundCount = k.GetRefundCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}