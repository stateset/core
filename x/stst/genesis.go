package stst

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stst/keeper"
	"github.com/stateset/core/x/stst/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set the module parameters
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Set the staking state
	if err := k.SetStakingState(ctx, genState.StakingState); err != nil {
		panic(err)
	}

	// Set the fee burn state
	if err := k.SetFeeBurnState(ctx, genState.FeeBurnState); err != nil {
		panic(err)
	}

	// Set vesting schedules and initialize their start times
	currentTime := ctx.BlockTime().Unix()
	for _, schedule := range genState.VestingSchedules {
		// Set start time to current time if not already set
		if schedule.StartTime == 0 {
			schedule.StartTime = currentTime
		}
		
		if err := k.SetVestingSchedule(ctx, schedule); err != nil {
			panic(err)
		}
	}

	// Note: Initial token supply and module account setup 
	// should be handled by the bank module and genesis configuration
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}

	stakingState, err := k.GetStakingState(ctx)
	if err != nil {
		panic(err)
	}

	feeBurnState, err := k.GetFeeBurnState(ctx)
	if err != nil {
		panic(err)
	}

	vestingSchedules, err := k.GetAllVestingSchedules(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Params:           params,
		StakingState:     stakingState,
		FeeBurnState:     feeBurnState,
		VestingSchedules: vestingSchedules,
	}
}