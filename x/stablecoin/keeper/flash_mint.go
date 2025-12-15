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
// Flash Mint Parameters
// ============================================================================

// GetFlashMintParams retrieves flash mint parameters.
func (k Keeper) GetFlashMintParams(ctx sdk.Context) types.FlashMintParams {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FlashMintParamsKey)
	if len(bz) == 0 {
		return DefaultFlashMintParams()
	}
	var params types.FlashMintParams
	types.MustUnmarshalJSON(bz, &params)
	return params
}

// SetFlashMintParams stores flash mint parameters.
func (k Keeper) SetFlashMintParams(ctx sdk.Context, params types.FlashMintParams) error {
	if params.FeeBps > 1000 {
		return errorsmod.Wrap(types.ErrInvalidReserve, "flash mint fee cannot exceed 10%")
	}
	if params.MaxFlashMint.IsNegative() {
		return errorsmod.Wrap(types.ErrInvalidReserve, "max flash mint cannot be negative")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.FlashMintParamsKey, types.MustMarshalJSON(params))
	return nil
}

// DefaultFlashMintParams returns default flash mint parameters.
func DefaultFlashMintParams() types.FlashMintParams {
	return types.FlashMintParams{
		Enabled:      true,
		FeeBps:       9, // 0.09% fee (similar to MakerDAO)
		MaxFlashMint: sdkmath.NewInt(100_000_000_000_000), // 100M ssUSD max per flash mint
	}
}

// ============================================================================
// Flash Mint Statistics
// ============================================================================

// GetFlashMintStats retrieves flash mint statistics.
func (k Keeper) GetFlashMintStats(ctx sdk.Context) types.FlashMintStats {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FlashMintStatsKey)
	if len(bz) == 0 {
		return types.FlashMintStats{
			TotalFlashMinted:   sdkmath.ZeroInt(),
			TotalFeesCollected: sdkmath.ZeroInt(),
		}
	}
	var stats types.FlashMintStats
	types.MustUnmarshalJSON(bz, &stats)
	return stats
}

// SetFlashMintStats stores flash mint statistics.
func (k Keeper) SetFlashMintStats(ctx sdk.Context, stats types.FlashMintStats) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.FlashMintStatsKey, types.MustMarshalJSON(stats))
}

// ============================================================================
// Flash Mint Session Tracking
// ============================================================================

// FlashMintSession tracks an active flash mint within a transaction.
type FlashMintSession struct {
	Sender         string
	Amount         sdkmath.Int
	Fee            sdkmath.Int
	AmountToReturn sdkmath.Int
}

// GetFlashMintSession retrieves the active flash mint session for a sender.
func (k Keeper) GetFlashMintSession(ctx sdk.Context, sender string) (FlashMintSession, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.FlashMintSessionKey, []byte(sender)...)
	bz := store.Get(key)
	if len(bz) == 0 {
		return FlashMintSession{}, false
	}
	var session FlashMintSession
	types.MustUnmarshalJSON(bz, &session)
	return session, true
}

// SetFlashMintSession stores an active flash mint session.
func (k Keeper) SetFlashMintSession(ctx sdk.Context, session FlashMintSession) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.FlashMintSessionKey, []byte(session.Sender)...)
	store.Set(key, types.MustMarshalJSON(session))
}

// DeleteFlashMintSession removes a flash mint session.
func (k Keeper) DeleteFlashMintSession(ctx sdk.Context, sender string) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.FlashMintSessionKey, []byte(sender)...)
	store.Delete(key)
}

// HasActiveFlashMintSession checks if there's an active flash mint session.
func (k Keeper) HasActiveFlashMintSession(ctx sdk.Context, sender string) bool {
	store := ctx.KVStore(k.storeKey)
	key := append(types.FlashMintSessionKey, []byte(sender)...)
	return store.Has(key)
}

// ============================================================================
// Flash Mint Operations
// ============================================================================

// FlashMint initiates a flash mint operation.
// This mints ssUSD to the sender, which must be returned (plus fee) via FlashMintCallback
// within the same transaction. This is enforced via an EndBlocker check.
func (k Keeper) FlashMint(ctx sdk.Context, sender sdk.AccAddress, amount sdkmath.Int) (sdkmath.Int, sdkmath.Int, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	params := k.GetFlashMintParams(ctx)
	if !params.Enabled {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrap(types.ErrMintPaused, "flash minting is disabled")
	}

	if amount.GT(params.MaxFlashMint) {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrDailyMintLimitExceeded,
			"flash mint amount %s exceeds max %s", amount, params.MaxFlashMint)
	}

	senderStr := sender.String()

	// Check for existing session (no nested flash mints)
	if k.HasActiveFlashMintSession(ctx, senderStr) {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrap(types.ErrInvalidReserve, "nested flash mints not allowed")
	}

	// Calculate fee
	feeRate := sdkmath.LegacyNewDec(int64(params.FeeBps)).Quo(sdkmath.LegacyNewDec(10000))
	fee := feeRate.MulInt(amount).TruncateInt()
	amountToReturn := amount.Add(fee)

	// Mint ssUSD to sender
	mintCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, amount))
	if err := k.bankKeeper.MintCoins(wrappedCtx, types.ModuleAccountName, mintCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, sender, mintCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Create session to track the flash mint
	session := FlashMintSession{
		Sender:         senderStr,
		Amount:         amount,
		Fee:            fee,
		AmountToReturn: amountToReturn,
	}
	k.SetFlashMintSession(ctx, session)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeFlashMint,
			sdk.NewAttribute(types.AttributeKeySender, senderStr),
			sdk.NewAttribute(types.AttributeKeyFlashMintAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyFlashMintFee, fee.String()),
		),
	)

	return amount, fee, nil
}

// FlashMintCallback completes a flash mint by returning the minted amount plus fee.
// This must be called in the same transaction as FlashMint.
func (k Keeper) FlashMintCallback(ctx sdk.Context, sender sdk.AccAddress, amountToReturn sdkmath.Int) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return err
	}

	senderStr := sender.String()

	// Get the flash mint session
	session, found := k.GetFlashMintSession(ctx, senderStr)
	if !found {
		return errorsmod.Wrap(types.ErrVaultNotFound, "no active flash mint session")
	}

	// Verify amount
	if amountToReturn.LT(session.AmountToReturn) {
		return errorsmod.Wrapf(types.ErrInvalidAmount,
			"insufficient return amount: got %s, required %s", amountToReturn, session.AmountToReturn)
	}

	// Transfer ssUSD from sender to module
	returnCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, amountToReturn))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, sender, types.ModuleAccountName, returnCoins); err != nil {
		return err
	}

	// Burn the principal
	principalCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, session.Amount))
	if err := k.bankKeeper.BurnCoins(wrappedCtx, types.ModuleAccountName, principalCoins); err != nil {
		return err
	}

	// Send fee to fee collector
	actualFee := amountToReturn.Sub(session.Amount)
	if actualFee.IsPositive() {
		feeCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, actualFee))
		if err := k.bankKeeper.SendCoinsFromModuleToModule(wrappedCtx, types.ModuleAccountName, authtypes.FeeCollectorName, feeCoins); err != nil {
			return err
		}
	}

	// Update stats
	stats := k.GetFlashMintStats(ctx)
	stats.TotalFlashMinted = stats.TotalFlashMinted.Add(session.Amount)
	stats.TotalFeesCollected = stats.TotalFeesCollected.Add(actualFee)
	k.SetFlashMintStats(ctx, stats)

	// Delete session
	k.DeleteFlashMintSession(ctx, senderStr)

	// Emit callback event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeFlashMintCallback,
			sdk.NewAttribute(types.AttributeKeySender, senderStr),
			sdk.NewAttribute(types.AttributeKeyAmount, amountToReturn.String()),
			sdk.NewAttribute(types.AttributeKeyFlashMintFee, actualFee.String()),
		),
	)

	return nil
}

// CheckPendingFlashMints checks for unclosed flash mint sessions and reverts them.
// This should be called in the EndBlocker to ensure atomicity.
func (k Keeper) CheckPendingFlashMints(ctx sdk.Context) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FlashMintSessionKey)
	iter := store.Iterator(nil, nil)
	defer iter.Close()

	var sessionsToRevert []FlashMintSession
	for ; iter.Valid(); iter.Next() {
		var session FlashMintSession
		types.MustUnmarshalJSON(iter.Value(), &session)
		sessionsToRevert = append(sessionsToRevert, session)
	}

	// Revert any unclosed sessions by burning the minted amount
	// In practice, this should not happen if the transaction is properly structured
	// because the callback must be included in the same tx
	for _, session := range sessionsToRevert {
		senderAddr, err := sdk.AccAddressFromBech32(session.Sender)
		if err != nil {
			// Delete invalid session
			k.DeleteFlashMintSession(ctx, session.Sender)
			continue
		}

		// Try to recover the minted amount from sender
		balance := k.bankKeeper.GetBalance(wrappedCtx, senderAddr, types.StablecoinDenom)
		if balance.Amount.GTE(session.Amount) {
			// Sender still has the funds, burn them
			burnCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, session.Amount))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, senderAddr, types.ModuleAccountName, burnCoins); err == nil {
				k.bankKeeper.BurnCoins(wrappedCtx, types.ModuleAccountName, burnCoins)
			}
		}
		// If sender doesn't have funds, this is a protocol loss
		// This should never happen in practice due to tx atomicity

		k.DeleteFlashMintSession(ctx, session.Sender)
	}

	return nil
}

// UpdateFlashMintParams updates flash mint parameters (governance only).
func (k Keeper) UpdateFlashMintParams(ctx sdk.Context, authority string, params types.FlashMintParams) error {
	if authority != k.GetAuthority() {
		return errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", k.GetAuthority(), authority)
	}
	return k.SetFlashMintParams(ctx, params)
}

// GetMaxFlashMint returns the maximum allowed flash mint amount.
func (k Keeper) GetMaxFlashMint(ctx sdk.Context) sdkmath.Int {
	params := k.GetFlashMintParams(ctx)
	return params.MaxFlashMint
}

// GetFlashMintFee calculates the flash mint fee for a given amount.
func (k Keeper) GetFlashMintFee(ctx sdk.Context, amount sdkmath.Int) sdkmath.Int {
	params := k.GetFlashMintParams(ctx)
	feeRate := sdkmath.LegacyNewDec(int64(params.FeeBps)).Quo(sdkmath.LegacyNewDec(10000))
	return feeRate.MulInt(amount).TruncateInt()
}

// SimulateFlashMint simulates a flash mint without executing it.
func (k Keeper) SimulateFlashMint(ctx sdk.Context, amount sdkmath.Int) (fee sdkmath.Int, totalToReturn sdkmath.Int, err error) {
	params := k.GetFlashMintParams(ctx)
	if !params.Enabled {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrap(types.ErrMintPaused, "flash minting is disabled")
	}

	if amount.GT(params.MaxFlashMint) {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrDailyMintLimitExceeded,
			"flash mint amount %s exceeds max %s", amount, params.MaxFlashMint)
	}

	fee = k.GetFlashMintFee(ctx, amount)
	totalToReturn = amount.Add(fee)
	return fee, totalToReturn, nil
}

// FlashMintArbitrage is a helper function for common arbitrage patterns.
// It flash mints ssUSD, allows arbitrary operations via callback, and ensures repayment.
// The callback data contains instructions for what to do with the minted funds.
func (k Keeper) FlashMintArbitrage(ctx sdk.Context, sender sdk.AccAddress, amount sdkmath.Int, callbackData []byte) error {
	// This would be implemented with WASM or IBC callbacks in a full implementation
	// For now, the basic FlashMint + FlashMintCallback flow handles this
	return fmt.Errorf("flash mint arbitrage requires WASM or IBC callback implementation")
}
