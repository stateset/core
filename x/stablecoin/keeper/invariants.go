package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stablecoin/types"
)

// RegisterInvariants registers all stablecoin module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "reserve-backing", ReserveBackingInvariant(k))
	ir.RegisterRoute(types.ModuleName, "total-supply-match", TotalSupplyMatchInvariant(k))
	ir.RegisterRoute(types.ModuleName, "vault-collateralization", VaultCollateralizationInvariant(k))
	ir.RegisterRoute(types.ModuleName, "deposit-consistency", DepositConsistencyInvariant(k))
}

// AllInvariants runs all invariants of the stablecoin module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := ReserveBackingInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		res, stop = TotalSupplyMatchInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		res, stop = VaultCollateralizationInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return DepositConsistencyInvariant(k)(ctx)
	}
}

// ReserveBackingInvariant checks that total reserves >= total minted ssUSD
func ReserveBackingInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		reserve := k.GetReserve(ctx)
		params := k.GetReserveParams(ctx)

		// Skip check if nothing minted
		if reserve.TotalMinted.IsZero() {
			return "", false
		}

		// Calculate reserve ratio
		reserveRatio := reserve.GetReserveRatio()

		// Reserve ratio must be at least the minimum (100% = 10000 bps)
		if reserveRatio < params.MinReserveRatioBps {
			return sdk.FormatInvariant(
				types.ModuleName,
				"reserve-backing",
				fmt.Sprintf(
					"reserve backing violation: reserve ratio %d bps < minimum %d bps (reserves: %s, minted: %s)",
					reserveRatio, params.MinReserveRatioBps,
					reserve.TotalValue, reserve.TotalMinted,
				),
			), true
		}

		return "", false
	}
}

// TotalSupplyMatchInvariant checks that tracked minted amount matches actual supply
func TotalSupplyMatchInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		reserve := k.GetReserve(ctx)

		supply := k.getStablecoinSupply(ctx)
		if reserve.TotalMinted.IsNegative() {
			return sdk.FormatInvariant(
				types.ModuleName,
				"total-supply-match",
				fmt.Sprintf("TotalMinted is negative: %s", reserve.TotalMinted),
			), true
		}

		if !reserve.TotalMinted.Equal(supply) {
			return sdk.FormatInvariant(
				types.ModuleName,
				"total-supply-match",
				fmt.Sprintf("tracked minted %s does not match bank supply %s", reserve.TotalMinted, supply),
			), true
		}

		return "", false
	}
}

// VaultCollateralizationInvariant checks that all vaults are properly collateralized
func VaultCollateralizationInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		params := k.GetParams(ctx)

		var brokenVaults []string

		k.IterateVaults(ctx, func(vault types.Vault) bool {
			// Skip vaults with no debt
			if vault.Debt.IsZero() {
				return false
			}

			// Get collateral param
			cp, ok := params.GetCollateralParam(vault.CollateralDenom)
			if !ok {
				brokenVaults = append(brokenVaults, fmt.Sprintf(
					"vault %d has unsupported collateral type %s",
					vault.Id, vault.CollateralDenom,
				))
				return false
			}

			// Get oracle price
			wrappedCtx := sdk.WrapSDKContext(ctx)
			price, err := k.oracleKeeper.GetPriceDec(wrappedCtx, vault.CollateralDenom)
			if err != nil {
				// Can't verify without price - skip
				return false
			}

			// Calculate collateral value
			collateralValue := vault.Collateral.Amount.ToLegacyDec().Mul(price)
			requiredValue := sdkmath.LegacyNewDecFromInt(vault.Debt).Mul(cp.LiquidationRatio)

			// If underwater, mark as broken
			if collateralValue.LT(requiredValue) {
				brokenVaults = append(brokenVaults, fmt.Sprintf(
					"vault %d is undercollateralized: value %s < required %s",
					vault.Id, collateralValue, requiredValue,
				))
			}

			return false
		})

		if len(brokenVaults) > 0 {
			return sdk.FormatInvariant(
				types.ModuleName,
				"vault-collateralization",
				fmt.Sprintf("found %d undercollateralized vaults: %v", len(brokenVaults), brokenVaults),
			), true
		}

		return "", false
	}
}

// DepositConsistencyInvariant checks that deposit records match reserve state
func DepositConsistencyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		reserve := k.GetReserve(ctx)

		// Sum all active deposits
		totalFromDeposits := sdk.NewCoins()
		totalMintedFromDeposits := sdkmath.ZeroInt()

		k.IterateReserveDeposits(ctx, func(deposit types.ReserveDeposit) bool {
			if deposit.Status == types.DepositStatusActive {
				totalFromDeposits = totalFromDeposits.Add(deposit.Amount)
				totalMintedFromDeposits = totalMintedFromDeposits.Add(deposit.SsusdMinted)
			}
			return false
		})

		// Check that deposit totals don't exceed reserve totals
		for _, coin := range totalFromDeposits {
			reserveAmount := reserve.TotalDeposited.AmountOf(coin.Denom)
			if coin.Amount.GT(reserveAmount) {
				return sdk.FormatInvariant(
					types.ModuleName,
					"deposit-consistency",
					fmt.Sprintf(
						"deposit total %s exceeds reserve total %s for denom %s",
						coin.Amount, reserveAmount, coin.Denom,
					),
				), true
			}
		}

		return "", false
	}
}
