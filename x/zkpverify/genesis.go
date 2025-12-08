package zkpverify

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/zkpverify/keeper"
	"github.com/stateset/core/x/zkpverify/types"
)

// InitGenesis initializes the module's state from a genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) {
	if err := k.InitGenesis(ctx, gs); err != nil {
		panic(err)
	}
}

// ExportGenesis exports the module's state to a genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return k.ExportGenesis(ctx)
}
