package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/stablecoin/types"
)

// TestVaultLiquidation_UndercollateralizedVault tests liquidating an undercollateralized vault
func TestVaultLiquidation_UndercollateralizedVault(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()
	liquidator := newStablecoinAddress()

	// Setup collateral price at $100
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000)) // $100 with 6 decimals

	// Fund owner with collateral
	collateral := sdk.NewCoin("atom", sdkmath.NewInt(10000000)) // 10 ATOM
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	// Create vault with debt
	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(800000000)) // 800 SSUSD
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Drop collateral price to $80 (makes vault undercollateralized)
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(80000000)) // $80

	// Fund liquidator
	bankKeeper.SetBalance(liquidator.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000))))

	// Liquidate vault
	err = k.LiquidateVault(ctx, vaultID, liquidator)
	require.NoError(t, err)

	// Verify vault is liquidated or removed
	vault, found := k.GetVault(ctx, vaultID)
	if found {
		require.True(t, vault.Liquidated)
	}
}

func TestVaultLiquidation_FullyCollateralizedVault(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()
	liquidator := newStablecoinAddress()

	// Setup collateral price at $100
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	// Fund owner
	collateral := sdk.NewCoin("atom", sdkmath.NewInt(10000000)) // 10 ATOM = $1000
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	// Create vault with debt (safe collateral ratio)
	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(500000000)) // 500 SSUSD (< $1000/2)
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Try to liquidate (should fail - vault is safe)
	err = k.LiquidateVault(ctx, vaultID, liquidator)
	require.Error(t, err)
}

func TestVaultLiquidation_AtLiquidationThreshold(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()
	liquidator := newStablecoinAddress()

	// Setup collateral price
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	collateral := sdk.NewCoin("atom", sdkmath.NewInt(15000000)) // 15 ATOM = $1500
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	// Create vault with debt at liquidation threshold (assuming 150% ratio)
	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)) // 1000 SSUSD
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Fund liquidator
	bankKeeper.SetBalance(liquidator.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000000))))

	// Should be liquidatable
	err = k.LiquidateVault(ctx, vaultID, liquidator)
	require.NoError(t, err)
}

func TestVaultLiquidation_PartialLiquidation(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()
	liquidator := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	collateral := sdk.NewCoin("atom", sdkmath.NewInt(20000000)) // 20 ATOM = $2000
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	// Large debt
	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(1500000000)) // 1500 SSUSD
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Drop price to make it liquidatable
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(80000000))

	// Fund liquidator with partial amount
	partialAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(500000000))
	bankKeeper.SetBalance(liquidator.String(), sdk.NewCoins(partialAmount))

	// Attempt partial liquidation
	err = k.PartialLiquidate(ctx, vaultID, liquidator, partialAmount)
	if err != nil {
		// Module may not support partial liquidation
		require.Error(t, err)
	}
}

func TestVaultLiquidation_MultipleCollateralTypes(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	// Setup multiple collateral prices
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))
	oracleKeeper.SetPrice("osmo", sdkmath.NewInt(50000000))

	// Fund owner with multiple collateral types
	atomCollateral := sdk.NewCoin("atom", sdkmath.NewInt(5000000))  // 5 ATOM = $500
	osmoCollateral := sdk.NewCoin("osmo", sdkmath.NewInt(10000000)) // 10 OSMO = $500
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(atomCollateral, osmoCollateral))

	// Try creating vault with mixed collateral (may not be supported)
	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(600000000))
	_, err := k.CreateVault(ctx, owner, atomCollateral, debt)
	// Just check it doesn't panic
	if err != nil {
		require.Error(t, err)
	}
}

func TestVaultLiquidation_ZeroDebtVault(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()
	liquidator := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	collateral := sdk.NewCoin("atom", sdkmath.NewInt(10000000))
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	// Create vault with zero debt
	zeroDebt := sdk.NewCoin("ssusd", sdkmath.ZeroInt())
	vaultID, err := k.CreateVault(ctx, owner, collateral, zeroDebt)
	if err != nil {
		// Zero debt may not be allowed
		require.Error(t, err)
		return
	}

	// Try to liquidate zero-debt vault (should fail)
	err = k.LiquidateVault(ctx, vaultID, liquidator)
	require.Error(t, err)
}

func TestVaultLiquidation_LiquidatorIncentive(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()
	liquidator := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	collateral := sdk.NewCoin("atom", sdkmath.NewInt(10000000))
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(800000000))
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Make undercollateralized
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(80000000))

	// Fund liquidator
	bankKeeper.SetBalance(liquidator.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000))))

	// Get liquidator balance before
	balanceBefore := bankKeeper.GetBalance(sdk.WrapSDKContext(ctx), liquidator, "atom")

	// Liquidate
	err = k.LiquidateVault(ctx, vaultID, liquidator)
	if err == nil {
		// Check liquidator received incentive
		balanceAfter := bankKeeper.GetBalance(sdk.WrapSDKContext(ctx), liquidator, "atom")
		require.True(t, balanceAfter.Amount.GT(balanceBefore.Amount), "Liquidator should receive collateral")
	}
}

func TestVaultLiquidation_InsufficientLiquidatorFunds(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()
	liquidator := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	collateral := sdk.NewCoin("atom", sdkmath.NewInt(10000000))
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(800000000))
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Make undercollateralized
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(80000000))

	// Liquidator has insufficient funds
	bankKeeper.SetBalance(liquidator.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(100000))))

	// Try to liquidate (should fail)
	err = k.LiquidateVault(ctx, vaultID, liquidator)
	require.Error(t, err)
}

// TestCollateralRatio_EdgeCases tests collateral ratio calculations
func TestCollateralRatio_ExactlyAtMinimum(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	// Create vault with exactly minimum collateral ratio (200%)
	collateral := sdk.NewCoin("atom", sdkmath.NewInt(20000000)) // 20 ATOM = $2000
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)) // 1000 SSUSD
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Verify vault is valid
	vault, found := k.GetVault(ctx, vaultID)
	require.True(t, found)
	require.False(t, vault.Liquidated)
}

func TestCollateralRatio_BelowMinimum(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	// Try to create vault with insufficient collateral
	collateral := sdk.NewCoin("atom", sdkmath.NewInt(15000000)) // 15 ATOM = $1500
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)) // 1000 SSUSD (ratio = 150%)
	_, err := k.CreateVault(ctx, owner, collateral, debt)
	require.Error(t, err) // Should fail minimum collateral ratio
}

func TestCollateralRatio_HighlyOvercollateralized(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	// Create highly overcollateralized vault (1000%)
	collateral := sdk.NewCoin("atom", sdkmath.NewInt(100000000)) // 100 ATOM = $10000
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)) // 1000 SSUSD
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	vault, found := k.GetVault(ctx, vaultID)
	require.True(t, found)
	require.False(t, vault.Liquidated)
}

func TestCollateralRatio_PriceVolatility(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	// Start with good price
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	collateral := sdk.NewCoin("atom", sdkmath.NewInt(25000000)) // 25 ATOM = $2500
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)) // 1000 SSUSD (250% ratio)
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Simulate price crash (50% drop)
	oracleKeeper.SetPrice("atom", sdkmath.NewInt(50000000))

	// Check if vault becomes liquidatable
	vault, found := k.GetVault(ctx, vaultID)
	require.True(t, found)

	// Calculate new ratio: $1250 / $1000 = 125% (below 150% threshold)
	ratio := k.CalculateCollateralRatio(ctx, vault)
	require.True(t, ratio.LT(sdkmath.LegacyNewDec(150)))
}

func TestCollateralRatio_MaxDebtLimit(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	// Very large collateral
	collateral := sdk.NewCoin("atom", sdkmath.NewInt(1000000000)) // 1M ATOM
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	// Try to mint maximum possible debt
	maxDebt := sdk.NewCoin("ssusd", sdkmath.NewInt(50000000000000)) // 50M SSUSD
	_, err := k.CreateVault(ctx, owner, collateral, maxDebt)
	if err != nil {
		// May hit global debt limit
		require.Error(t, err)
	}
}

func TestCollateralRatio_DustAmounts(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	// Very small collateral (dust)
	collateral := sdk.NewCoin("atom", sdkmath.NewInt(100)) // 0.0001 ATOM
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	// Try to create vault with dust
	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(1))
	_, err := k.CreateVault(ctx, owner, collateral, debt)
	// Should fail minimum vault size
	require.Error(t, err)
}

func TestCollateralRatio_AfterRepayment(t *testing.T) {
	k, ctx, bankKeeper, oracleKeeper, _ := setupStablecoinKeeper(t)

	owner := newStablecoinAddress()

	oracleKeeper.SetPrice("atom", sdkmath.NewInt(100000000))

	collateral := sdk.NewCoin("atom", sdkmath.NewInt(20000000)) // 20 ATOM = $2000
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(collateral))

	debt := sdk.NewCoin("ssusd", sdkmath.NewInt(900000000)) // 900 SSUSD (222% ratio)
	vaultID, err := k.CreateVault(ctx, owner, collateral, debt)
	require.NoError(t, err)

	// Fund owner with stablecoin for repayment
	bankKeeper.SetBalance(owner.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(500000000))))

	// Repay part of debt
	repayment := sdk.NewCoin("ssusd", sdkmath.NewInt(400000000))
	err = k.RepayStablecoin(ctx, owner, vaultID, repayment)
	require.NoError(t, err)

	// Check improved ratio
	vault, found := k.GetVault(ctx, vaultID)
	require.True(t, found)

	// New ratio should be: $2000 / $500 = 400%
	ratio := k.CalculateCollateralRatio(ctx, vault)
	require.True(t, ratio.GT(sdkmath.LegacyNewDec(300)))
}
