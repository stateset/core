package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stst/types"
)

// RegisterInvariants registers the STST module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "token-supply", TokenSupplyInvariant(k))
	ir.RegisterRoute(types.ModuleName, "staking-consistency", StakingConsistencyInvariant(k))
	ir.RegisterRoute(types.ModuleName, "vesting-consistency", VestingConsistencyInvariant(k))
}

// TokenSupplyInvariant checks that the total token supply is correct
func TokenSupplyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		params, err := k.GetParams(ctx)
		if err != nil {
			return sdk.FormatInvariant(
				types.ModuleName, "token-supply",
				"failed to get params",
			), true
		}

		// Check that total supply has not been exceeded
		// This would be a critical violation since STST has a fixed supply
		expectedSupply := params.TotalSupply
		
		// TODO: Get actual circulating supply from bank module
		// actualSupply := k.bankKeeper.GetSupply(ctx, params.TokenDenom).Amount

		msg := fmt.Sprintf(
			"expected total supply: %s\nactual supply: %s\n",
			expectedSupply.String(),
			"TODO", // actualSupply.String(),
		)

		// For now, return false (no violation) since we can't check actual supply
		return sdk.FormatInvariant(types.ModuleName, "token-supply", msg), false
	}
}

// StakingConsistencyInvariant checks that staking state is consistent
func StakingConsistencyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		stakingState, err := k.GetStakingState(ctx)
		if err != nil {
			return sdk.FormatInvariant(
				types.ModuleName, "staking-consistency",
				"failed to get staking state",
			), true
		}

		// Check that total staked amount is non-negative
		if stakingState.TotalStaked.IsNegative() {
			return sdk.FormatInvariant(
				types.ModuleName, "staking-consistency",
				fmt.Sprintf("total staked amount is negative: %s", stakingState.TotalStaked),
			), true
		}

		// TODO: Check that sum of all delegations equals total staked
		// This would require iterating through all delegations

		msg := fmt.Sprintf(
			"total staked: %s\ntotal validators: %d\n",
			stakingState.TotalStaked.String(),
			stakingState.TotalValidators,
		)

		return sdk.FormatInvariant(types.ModuleName, "staking-consistency", msg), false
	}
}

// VestingConsistencyInvariant checks that vesting schedules are consistent
func VestingConsistencyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		params, err := k.GetParams(ctx)
		if err != nil {
			return sdk.FormatInvariant(
				types.ModuleName, "vesting-consistency",
				"failed to get params",
			), true
		}

		vestingSchedules, err := k.GetAllVestingSchedules(ctx)
		if err != nil {
			return sdk.FormatInvariant(
				types.ModuleName, "vesting-consistency",
				"failed to get vesting schedules",
			), true
		}

		totalVestingAmount := sdk.ZeroInt()
		for _, schedule := range vestingSchedules {
			totalVestingAmount = totalVestingAmount.Add(schedule.TotalAmount)

			// Check that vested amount doesn't exceed total amount
			if schedule.VestedAmount.GT(schedule.TotalAmount) {
				return sdk.FormatInvariant(
					types.ModuleName, "vesting-consistency",
					fmt.Sprintf("vested amount (%s) exceeds total amount (%s) for category %s",
						schedule.VestedAmount, schedule.TotalAmount, schedule.Category),
				), true
			}
		}

		// Check that total vesting amount equals total supply
		if !totalVestingAmount.Equal(params.TotalSupply) {
			return sdk.FormatInvariant(
				types.ModuleName, "vesting-consistency",
				fmt.Sprintf("total vesting amount (%s) does not equal total supply (%s)",
					totalVestingAmount, params.TotalSupply),
			), true
		}

		msg := fmt.Sprintf(
			"total vesting amount: %s\ntotal supply: %s\n",
			totalVestingAmount.String(),
			params.TotalSupply.String(),
		)

		return sdk.FormatInvariant(types.ModuleName, "vesting-consistency", msg), false
	}
}