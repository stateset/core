package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/stateset/core/x/stablecoin/types"
)

// ============================================================================
// PSM Configuration
// ============================================================================

// GetPSMConfig retrieves the PSM configuration for a given denom.
func (k Keeper) GetPSMConfig(ctx sdk.Context, denom string) (types.PSMConfig, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PSMConfigKey(denom))
	if len(bz) == 0 {
		return types.PSMConfig{}, false
	}
	var config types.PSMConfig
	types.MustUnmarshalJSON(bz, &config)
	return config, true
}

// SetPSMConfig stores the PSM configuration for a given denom.
func (k Keeper) SetPSMConfig(ctx sdk.Context, config types.PSMConfig) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PSMConfigKey(config.Denom), types.MustMarshalJSON(config))
}

// GetAllPSMConfigs returns all PSM configurations.
func (k Keeper) GetAllPSMConfigs(ctx sdk.Context) []types.PSMConfig {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PSMConfigKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()

	var configs []types.PSMConfig
	for ; iter.Valid(); iter.Next() {
		var config types.PSMConfig
		types.MustUnmarshalJSON(iter.Value(), &config)
		configs = append(configs, config)
	}
	return configs
}

// ============================================================================
// PSM State
// ============================================================================

// GetPSMState retrieves the PSM state for a given denom.
func (k Keeper) GetPSMState(ctx sdk.Context, denom string) types.PSMState {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PSMStateKey(denom))
	if len(bz) == 0 {
		return types.PSMState{
			Denom:          denom,
			TotalDeposited: sdkmath.ZeroInt(),
			TotalMinted:    sdkmath.ZeroInt(),
		}
	}
	var state types.PSMState
	types.MustUnmarshalJSON(bz, &state)
	return state
}

// SetPSMState stores the PSM state for a given denom.
func (k Keeper) SetPSMState(ctx sdk.Context, state types.PSMState) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PSMStateKey(state.Denom), types.MustMarshalJSON(state))
}

// ============================================================================
// PSM Swap Operations
// ============================================================================

// PSMSwapIn swaps a stablecoin (e.g., USDC) for ssUSD at 1:1 minus fee.
// This is used when ssUSD is trading above $1.00 - arbitrageurs mint ssUSD with USDC.
func (k Keeper) PSMSwapIn(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coin) (sdkmath.Int, sdkmath.Int, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Get PSM config for this denom
	config, found := k.GetPSMConfig(ctx, amount.Denom)
	if !found {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "PSM not configured for %s", amount.Denom)
	}
	if !config.Active {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "PSM for %s is inactive", amount.Denom)
	}

	// Check debt ceiling
	state := k.GetPSMState(ctx, amount.Denom)
	newTotalMinted := state.TotalMinted.Add(amount.Amount)
	if newTotalMinted.GT(config.DebtCeiling) {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrDailyMintLimitExceeded,
			"PSM debt ceiling exceeded: current %s + mint %s > ceiling %s",
			state.TotalMinted, amount.Amount, config.DebtCeiling)
	}

	// Validate price is near $1.00 (safety check)
	price, err := k.oracleKeeper.GetPriceDecSafe(wrappedCtx, config.OracleDenom)
	if err == nil {
		// Allow 5% deviation for PSM stablecoins
		upperBound := sdkmath.LegacyNewDecWithPrec(105, 2) // 1.05
		lowerBound := sdkmath.LegacyNewDecWithPrec(95, 2)  // 0.95
		if price.GT(upperBound) || price.LT(lowerBound) {
			return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrPriceNotFound,
				"PSM asset %s price %s outside safety bounds (0.95-1.05)", config.OracleDenom, price)
		}
	}
	// Note: If oracle fails, we proceed with 1:1 assumption (PSM stablecoins should be pegged)

	// Calculate fee
	feeRate := sdkmath.LegacyNewDec(int64(config.MintFeeBps)).Quo(sdkmath.LegacyNewDec(10000))
	feeAmount := feeRate.MulInt(amount.Amount).TruncateInt()
	ssusdToMint := amount.Amount.Sub(feeAmount)

	// Transfer stablecoin from sender to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, sender, types.ModuleAccountName, sdk.NewCoins(amount)); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Send fee to fee collector (in the input stablecoin)
	if !feeAmount.IsZero() {
		feeCoin := sdk.NewCoin(amount.Denom, feeAmount)
		if err := k.bankKeeper.SendCoinsFromModuleToModule(wrappedCtx, types.ModuleAccountName, authtypes.FeeCollectorName, sdk.NewCoins(feeCoin)); err != nil {
			return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
		}
	}

	// Mint ssUSD to sender
	mintCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, ssusdToMint))
	if err := k.bankKeeper.MintCoins(wrappedCtx, types.ModuleAccountName, mintCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, sender, mintCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Update PSM state
	state.TotalDeposited = state.TotalDeposited.Add(amount.Amount.Sub(feeAmount))
	state.TotalMinted = state.TotalMinted.Add(ssusdToMint)
	k.SetPSMState(ctx, state)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePSMSwapIn,
			sdk.NewAttribute(types.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyInputDenom, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.Amount.String()),
			sdk.NewAttribute(types.AttributeKeySsusdAmount, ssusdToMint.String()),
			sdk.NewAttribute(types.AttributeKeySwapFee, feeAmount.String()),
		),
	)

	return ssusdToMint, feeAmount, nil
}

// PSMSwapOut swaps ssUSD for a stablecoin (e.g., USDC) at 1:1 minus fee.
// This is used when ssUSD is trading below $1.00 - arbitrageurs redeem ssUSD for USDC.
func (k Keeper) PSMSwapOut(ctx sdk.Context, sender sdk.AccAddress, ssusdAmount sdkmath.Int, outputDenom string) (sdkmath.Int, sdkmath.Int, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Get PSM config for this denom
	config, found := k.GetPSMConfig(ctx, outputDenom)
	if !found {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "PSM not configured for %s", outputDenom)
	}
	if !config.Active {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "PSM for %s is inactive", outputDenom)
	}

	// Get current state
	state := k.GetPSMState(ctx, outputDenom)

	// Calculate fee
	feeRate := sdkmath.LegacyNewDec(int64(config.RedeemFeeBps)).Quo(sdkmath.LegacyNewDec(10000))
	feeAmount := feeRate.MulInt(ssusdAmount).TruncateInt()
	outputAmount := ssusdAmount.Sub(feeAmount)

	// Check PSM has sufficient liquidity
	if outputAmount.GT(state.TotalDeposited) {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrInsufficientReserves,
			"PSM has insufficient %s: requested %s, available %s",
			outputDenom, outputAmount, state.TotalDeposited)
	}

	// Transfer ssUSD from sender to module and burn
	ssusdCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, ssusdAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, sender, types.ModuleAccountName, ssusdCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}
	if err := k.bankKeeper.BurnCoins(wrappedCtx, types.ModuleAccountName, ssusdCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Transfer output stablecoin to sender
	outputCoins := sdk.NewCoins(sdk.NewCoin(outputDenom, outputAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, sender, outputCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Send fee to fee collector (in ssUSD terms, deducted from output)
	// The fee stays in the PSM as additional backing

	// Update PSM state
	state.TotalDeposited = state.TotalDeposited.Sub(outputAmount)
	state.TotalMinted = state.TotalMinted.Sub(ssusdAmount)
	if state.TotalMinted.IsNegative() {
		state.TotalMinted = sdkmath.ZeroInt()
	}
	k.SetPSMState(ctx, state)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePSMSwapOut,
			sdk.NewAttribute(types.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeySsusdAmount, ssusdAmount.String()),
			sdk.NewAttribute(types.AttributeKeyOutputDenom, outputDenom),
			sdk.NewAttribute(types.AttributeKeyAmount, outputAmount.String()),
			sdk.NewAttribute(types.AttributeKeySwapFee, feeAmount.String()),
		),
	)

	return outputAmount, feeAmount, nil
}

// UpdatePSMConfigs updates PSM configurations (governance only).
func (k Keeper) UpdatePSMConfigs(ctx sdk.Context, authority string, configs []types.PSMConfig) error {
	if authority != k.GetAuthority() {
		return errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", k.GetAuthority(), authority)
	}

	for _, config := range configs {
		if config.Denom == "" {
			return errorsmod.Wrap(types.ErrInvalidReserve, "PSM denom cannot be empty")
		}
		if config.MintFeeBps > 1000 {
			return errorsmod.Wrap(types.ErrInvalidReserve, "PSM mint fee cannot exceed 10%")
		}
		if config.RedeemFeeBps > 1000 {
			return errorsmod.Wrap(types.ErrInvalidReserve, "PSM redeem fee cannot exceed 10%")
		}
		if config.DebtCeiling.IsNegative() {
			return errorsmod.Wrap(types.ErrInvalidReserve, "PSM debt ceiling cannot be negative")
		}

		k.SetPSMConfig(ctx, config)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePSMConfigUpdate,
			sdk.NewAttribute(types.AttributeKeyAction, "update"),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d configs updated", len(configs))),
		),
	)

	return nil
}

// DefaultPSMConfigs returns default PSM configurations for common stablecoins.
func DefaultPSMConfigs() []types.PSMConfig {
	return []types.PSMConfig{
		{
			Denom:         "ibc/USDC", // Placeholder IBC denom for USDC
			Active:        false,      // Disabled by default, enable via governance
			MintFeeBps:    10,         // 0.1% mint fee
			RedeemFeeBps:  10,         // 0.1% redeem fee
			DebtCeiling:   sdkmath.NewInt(100_000_000_000_000), // 100M USDC ceiling
			OracleDenom:   "USDC",
		},
		{
			Denom:         "ibc/USDT", // Placeholder IBC denom for USDT
			Active:        false,      // Disabled by default
			MintFeeBps:    10,         // 0.1% mint fee
			RedeemFeeBps:  10,         // 0.1% redeem fee
			DebtCeiling:   sdkmath.NewInt(50_000_000_000_000), // 50M USDT ceiling
			OracleDenom:   "USDT",
		},
	}
}
