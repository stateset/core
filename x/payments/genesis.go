package payments

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/payments/keeper"
	"github.com/stateset/core/x/payments/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}
	k.InitGenesis(ctx, state)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return k.ExportGenesis(ctx)
}
