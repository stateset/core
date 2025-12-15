package keeper

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stablecoin/types"
)

// ============================================================================
// Savings Parameters
// ============================================================================

// GetSavingsParams retrieves savings parameters.
func (k Keeper) GetSavingsParams(ctx sdk.Context) types.SavingsParams {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SavingsParamsKey)
	if len(bz) == 0 {
		return DefaultSavingsParams()
	}
	var params types.SavingsParams
	types.MustUnmarshalJSON(bz, &params)
	return params
}

// SetSavingsParams stores savings parameters.
func (k Keeper) SetSavingsParams(ctx sdk.Context, params types.SavingsParams) error {
	if params.SavingsRateBps > 5000 {
		return errorsmod.Wrap(types.ErrInvalidReserve, "savings rate cannot exceed 50% APY")
	}
	if params.MinDeposit.IsNegative() {
		return errorsmod.Wrap(types.ErrInvalidReserve, "min deposit cannot be negative")
	}
	if params.AccrualIntervalSeconds < 0 {
		return errorsmod.Wrap(types.ErrInvalidReserve, "accrual interval cannot be negative")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.SavingsParamsKey, types.MustMarshalJSON(params))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSavingsParamsUpdate,
			sdk.NewAttribute(types.AttributeKeySavingsRate, fmt.Sprintf("%d", params.SavingsRateBps)),
		),
	)
	return nil
}

// DefaultSavingsParams returns default savings parameters.
func DefaultSavingsParams() types.SavingsParams {
	return types.SavingsParams{
		Enabled:                false, // Disabled by default
		SavingsRateBps:         500,   // 5% APY
		MinDeposit:             sdkmath.NewInt(1_000_000), // 1 ssUSD minimum
		AccrualIntervalSeconds: 86400, // Daily accrual
	}
}

// ============================================================================
// Savings Deposits
// ============================================================================

// GetSavingsDeposit retrieves a savings deposit for a given depositor.
func (k Keeper) GetSavingsDeposit(ctx sdk.Context, depositor string) (types.SavingsDeposit, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SavingsDepositKey(depositor))
	if len(bz) == 0 {
		return types.SavingsDeposit{}, false
	}
	var deposit types.SavingsDeposit
	types.MustUnmarshalJSON(bz, &deposit)
	return deposit, true
}

// SetSavingsDeposit stores a savings deposit.
func (k Keeper) SetSavingsDeposit(ctx sdk.Context, deposit types.SavingsDeposit) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SavingsDepositKey(deposit.Depositor), types.MustMarshalJSON(deposit))
}

// DeleteSavingsDeposit removes a savings deposit.
func (k Keeper) DeleteSavingsDeposit(ctx sdk.Context, depositor string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SavingsDepositKey(depositor))
}

// IterateSavingsDeposits iterates over all savings deposits.
func (k Keeper) IterateSavingsDeposits(ctx sdk.Context, cb func(types.SavingsDeposit) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SavingsDepositKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var deposit types.SavingsDeposit
		types.MustUnmarshalJSON(iter.Value(), &deposit)
		if cb(deposit) {
			break
		}
	}
}

// ============================================================================
// Savings Statistics
// ============================================================================

// GetSavingsStats retrieves savings statistics.
func (k Keeper) GetSavingsStats(ctx sdk.Context) types.SavingsStats {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SavingsStatsKey)
	if len(bz) == 0 {
		return types.SavingsStats{
			TotalDeposits:     sdkmath.ZeroInt(),
			TotalInterestPaid: sdkmath.ZeroInt(),
			DepositorCount:    0,
		}
	}
	var stats types.SavingsStats
	types.MustUnmarshalJSON(bz, &stats)
	return stats
}

// SetSavingsStats stores savings statistics.
func (k Keeper) SetSavingsStats(ctx sdk.Context, stats types.SavingsStats) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SavingsStatsKey, types.MustMarshalJSON(stats))
}

// ============================================================================
// Savings Operations
// ============================================================================

// DepositSavings deposits ssUSD into the savings module.
func (k Keeper) DepositSavings(ctx sdk.Context, depositor sdk.AccAddress, amount sdkmath.Int) (sdkmath.Int, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return sdkmath.ZeroInt(), err
	}

	params := k.GetSavingsParams(ctx)
	if !params.Enabled {
		return sdkmath.ZeroInt(), errorsmod.Wrap(types.ErrMintPaused, "savings rate is disabled")
	}

	if amount.LT(params.MinDeposit) {
		return sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrBelowMinimumMint,
			"deposit amount %s below minimum %s", amount, params.MinDeposit)
	}

	// Transfer ssUSD from depositor to module
	depositCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, amount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, depositor, types.ModuleAccountName, depositCoins); err != nil {
		return sdkmath.ZeroInt(), err
	}

	depositorStr := depositor.String()
	deposit, found := k.GetSavingsDeposit(ctx, depositorStr)

	stats := k.GetSavingsStats(ctx)

	if found {
		// Accrue any pending interest first
		deposit = k.accrueInterest(ctx, deposit, params)
		deposit.Principal = deposit.Principal.Add(amount)
	} else {
		// New depositor
		deposit = types.SavingsDeposit{
			Depositor:       depositorStr,
			Principal:       amount,
			AccruedInterest: sdkmath.ZeroInt(),
			LastAccrualTime: ctx.BlockTime(),
			DepositedAt:     ctx.BlockTime(),
		}
		stats.DepositorCount++
	}

	k.SetSavingsDeposit(ctx, deposit)

	// Update stats
	stats.TotalDeposits = stats.TotalDeposits.Add(amount)
	k.SetSavingsStats(ctx, stats)

	totalDeposit := deposit.Principal.Add(deposit.AccruedInterest)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSavingsDeposit,
			sdk.NewAttribute(types.AttributeKeyDepositor, depositorStr),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyPrincipal, deposit.Principal.String()),
			sdk.NewAttribute(types.AttributeKeyTotalDeposits, totalDeposit.String()),
		),
	)

	return totalDeposit, nil
}

// WithdrawSavings withdraws ssUSD from the savings module.
func (k Keeper) WithdrawSavings(ctx sdk.Context, depositor sdk.AccAddress, amount sdkmath.Int) (sdkmath.Int, sdkmath.Int, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	params := k.GetSavingsParams(ctx)
	depositorStr := depositor.String()

	deposit, found := k.GetSavingsDeposit(ctx, depositorStr)
	if !found {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrap(types.ErrVaultNotFound, "no savings deposit found")
	}

	// Accrue pending interest
	deposit = k.accrueInterest(ctx, deposit, params)

	totalBalance := deposit.Principal.Add(deposit.AccruedInterest)
	if amount.GT(totalBalance) {
		amount = totalBalance // Cap at total balance
	}

	// Calculate how much comes from interest vs principal
	interestWithdrawn := sdkmath.ZeroInt()
	principalWithdrawn := sdkmath.ZeroInt()

	if amount.LTE(deposit.AccruedInterest) {
		// All from interest
		interestWithdrawn = amount
		deposit.AccruedInterest = deposit.AccruedInterest.Sub(amount)
	} else {
		// All interest + some principal
		interestWithdrawn = deposit.AccruedInterest
		principalWithdrawn = amount.Sub(interestWithdrawn)
		deposit.AccruedInterest = sdkmath.ZeroInt()
		deposit.Principal = deposit.Principal.Sub(principalWithdrawn)
	}

	// Mint interest (it was virtual) and send to user
	if interestWithdrawn.IsPositive() {
		interestCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, interestWithdrawn))
		if err := k.bankKeeper.MintCoins(wrappedCtx, types.ModuleAccountName, interestCoins); err != nil {
			return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, depositor, interestCoins); err != nil {
			return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
		}
	}

	// Return principal from module
	if principalWithdrawn.IsPositive() {
		principalCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, principalWithdrawn))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, depositor, principalCoins); err != nil {
			return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
		}
	}

	// Update or delete deposit
	stats := k.GetSavingsStats(ctx)
	if deposit.Principal.IsZero() && deposit.AccruedInterest.IsZero() {
		k.DeleteSavingsDeposit(ctx, depositorStr)
		stats.DepositorCount--
	} else {
		k.SetSavingsDeposit(ctx, deposit)
	}

	// Update stats
	stats.TotalDeposits = stats.TotalDeposits.Sub(principalWithdrawn)
	if stats.TotalDeposits.IsNegative() {
		stats.TotalDeposits = sdkmath.ZeroInt()
	}
	stats.TotalInterestPaid = stats.TotalInterestPaid.Add(interestWithdrawn)
	k.SetSavingsStats(ctx, stats)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSavingsWithdraw,
			sdk.NewAttribute(types.AttributeKeyDepositor, depositorStr),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyPrincipal, principalWithdrawn.String()),
			sdk.NewAttribute(types.AttributeKeyInterest, interestWithdrawn.String()),
		),
	)

	return amount, interestWithdrawn, nil
}

// ClaimSavingsInterest claims accrued interest without withdrawing principal.
func (k Keeper) ClaimSavingsInterest(ctx sdk.Context, depositor sdk.AccAddress) (sdkmath.Int, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return sdkmath.ZeroInt(), err
	}

	params := k.GetSavingsParams(ctx)
	depositorStr := depositor.String()

	deposit, found := k.GetSavingsDeposit(ctx, depositorStr)
	if !found {
		return sdkmath.ZeroInt(), errorsmod.Wrap(types.ErrVaultNotFound, "no savings deposit found")
	}

	// Accrue pending interest
	deposit = k.accrueInterest(ctx, deposit, params)

	interestToClaim := deposit.AccruedInterest
	if interestToClaim.IsZero() {
		return sdkmath.ZeroInt(), nil // Nothing to claim
	}

	// Mint interest and send to user
	interestCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, interestToClaim))
	if err := k.bankKeeper.MintCoins(wrappedCtx, types.ModuleAccountName, interestCoins); err != nil {
		return sdkmath.ZeroInt(), err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, depositor, interestCoins); err != nil {
		return sdkmath.ZeroInt(), err
	}

	// Update deposit
	deposit.AccruedInterest = sdkmath.ZeroInt()
	k.SetSavingsDeposit(ctx, deposit)

	// Update stats
	stats := k.GetSavingsStats(ctx)
	stats.TotalInterestPaid = stats.TotalInterestPaid.Add(interestToClaim)
	k.SetSavingsStats(ctx, stats)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSavingsInterestClaim,
			sdk.NewAttribute(types.AttributeKeyDepositor, depositorStr),
			sdk.NewAttribute(types.AttributeKeyInterest, interestToClaim.String()),
		),
	)

	return interestToClaim, nil
}

// accrueInterest calculates and adds accrued interest to a deposit.
func (k Keeper) accrueInterest(ctx sdk.Context, deposit types.SavingsDeposit, params types.SavingsParams) types.SavingsDeposit {
	if !params.Enabled || params.SavingsRateBps == 0 || deposit.Principal.IsZero() {
		deposit.LastAccrualTime = ctx.BlockTime()
		return deposit
	}

	now := ctx.BlockTime()
	elapsed := now.Sub(deposit.LastAccrualTime)
	if elapsed <= 0 {
		return deposit
	}

	// Calculate interest: Principal * Rate * Time / Year
	// Rate is in bps, so divide by 10000
	// Time is in seconds, year is 365.25 * 24 * 60 * 60 = 31557600 seconds
	secondsPerYear := sdkmath.LegacyNewDec(31557600)
	elapsedSeconds := sdkmath.LegacyNewDec(int64(elapsed.Seconds()))
	rate := sdkmath.LegacyNewDec(int64(params.SavingsRateBps)).Quo(sdkmath.LegacyNewDec(10000))

	// Interest = Principal * Rate * (ElapsedSeconds / SecondsPerYear)
	interest := rate.MulInt(deposit.Principal).Mul(elapsedSeconds).Quo(secondsPerYear).TruncateInt()

	deposit.AccruedInterest = deposit.AccruedInterest.Add(interest)
	deposit.LastAccrualTime = now

	return deposit
}

// AccrueAllSavingsInterest accrues interest for all depositors (called in EndBlocker if needed).
func (k Keeper) AccrueAllSavingsInterest(ctx sdk.Context) {
	params := k.GetSavingsParams(ctx)
	if !params.Enabled {
		return
	}

	// Only accrue at specified intervals
	if params.AccrualIntervalSeconds > 0 {
		// This would be called periodically by governance or a cron mechanism
		// For simplicity, interest is accrued on interaction (deposit/withdraw/claim)
	}
}

// UpdateSavingsParams updates savings parameters (governance only).
func (k Keeper) UpdateSavingsParams(ctx sdk.Context, authority string, params types.SavingsParams) error {
	if authority != k.GetAuthority() {
		return errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", k.GetAuthority(), authority)
	}
	return k.SetSavingsParams(ctx, params)
}

// GetSavingsAccountBalance returns the total balance (principal + accrued interest) for a depositor.
func (k Keeper) GetSavingsAccountBalance(ctx sdk.Context, depositor string) (sdkmath.Int, sdkmath.Int, sdkmath.Int, error) {
	deposit, found := k.GetSavingsDeposit(ctx, depositor)
	if !found {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.ZeroInt(), nil
	}

	params := k.GetSavingsParams(ctx)
	deposit = k.accrueInterest(ctx, deposit, params)

	total := deposit.Principal.Add(deposit.AccruedInterest)
	return deposit.Principal, deposit.AccruedInterest, total, nil
}

// GetCurrentSavingsRate returns the current savings rate.
func (k Keeper) GetCurrentSavingsRate(ctx sdk.Context) (uint32, bool) {
	params := k.GetSavingsParams(ctx)
	return params.SavingsRateBps, params.Enabled
}

// CalculatePendingInterest calculates pending interest for a depositor without updating state.
func (k Keeper) CalculatePendingInterest(ctx sdk.Context, depositor string) (sdkmath.Int, time.Time) {
	deposit, found := k.GetSavingsDeposit(ctx, depositor)
	if !found {
		return sdkmath.ZeroInt(), time.Time{}
	}

	params := k.GetSavingsParams(ctx)
	accrued := k.accrueInterest(ctx, deposit, params)
	newInterest := accrued.AccruedInterest.Sub(deposit.AccruedInterest)

	return newInterest, deposit.LastAccrualTime
}
