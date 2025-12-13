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
	ir.RegisterRoute(types.ModuleName, "redemption-locks", RedemptionLocksInvariant(k))
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
		return RedemptionLocksInvariant(k)(ctx)
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
			price, err := k.oracleKeeper.GetPriceDecSafe(wrappedCtx, vault.CollateralDenom)
			if err != nil {
				brokenVaults = append(brokenVaults, fmt.Sprintf(
					"vault %d cannot verify collateralization for %s: %v",
					vault.Id, vault.CollateralDenom, err,
				))
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

// RedemptionLocksInvariant checks that locked reserves match pending redemptions
// and do not exceed the module's tracked reserve totals.
func RedemptionLocksInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		reserve := k.GetReserve(ctx)
		locked := k.GetLockedReserves(ctx)

		// Compute expected locked reserves from pending redemption requests.
		expected := sdk.NewCoins()
		var problems []string

		k.IterateRedemptionRequests(ctx, func(request types.RedemptionRequest) bool {
			if request.Status != types.RedeemStatusPending {
				return false
			}
			if request.OutputAmount.Denom == "" || !request.OutputAmount.Amount.IsPositive() {
				problems = append(problems, fmt.Sprintf("pending redemption %d missing output amount", request.Id))
				return false
			}
			if request.OutputAmount.Denom != request.OutputDenom {
				problems = append(problems, fmt.Sprintf("pending redemption %d output denom mismatch: %s != %s", request.Id, request.OutputAmount.Denom, request.OutputDenom))
				return false
			}
			expected = expected.Add(request.OutputAmount)
			return false
		})

		expected = expected.Sort()
		locked = locked.Sort()

		// locked must equal expected (exact match, order-independent).
		if len(expected) != len(locked) {
			problems = append(problems, fmt.Sprintf("locked reserves length %d != expected %d", len(locked), len(expected)))
		} else {
			for i := range expected {
				if expected[i].Denom != locked[i].Denom || !expected[i].Amount.Equal(locked[i].Amount) {
					problems = append(problems, fmt.Sprintf("locked reserves mismatch at %d: locked %s != expected %s", i, locked[i], expected[i]))
					break
				}
			}
		}

		// locked amounts must be <= total deposited for each denom.
		for _, coin := range locked {
			total := reserve.TotalDeposited.AmountOf(coin.Denom)
			if coin.Amount.GT(total) {
				problems = append(problems, fmt.Sprintf("locked %s%s exceeds total deposited %s%s", coin.Amount, coin.Denom, total, coin.Denom))
			}
		}

		if len(problems) > 0 {
			return sdk.FormatInvariant(
				types.ModuleName,
				"redemption-locks",
				fmt.Sprintf("locked reserve invariant violation: %v", problems),
			), true
		}

		return "", false
	}
}
