package compliance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/compliance/keeper"
	"github.com/stateset/core/x/compliance/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}
	k.InitGenesis(sdk.WrapSDKContext(ctx), state)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return k.ExportGenesis(sdk.WrapSDKContext(ctx))
}
