package feemarket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/feemarket/keeper"
	"github.com/stateset/core/x/feemarket/types"
)

// InitGenesis initializes state from genesis data.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}
	if err := state.Validate(); err != nil {
		panic(err)
	}
	k.SetParams(ctx, state.Params)
	k.SetBaseFee(ctx, state.BaseFee)
}

// ExportGenesis exports the current state.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:  k.GetParams(ctx),
		BaseFee: k.GetBaseFee(ctx),
	}
}
