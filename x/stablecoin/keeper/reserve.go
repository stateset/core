package keeper

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	oracletypes "github.com/stateset/core/x/oracle/types"
	"github.com/stateset/core/x/stablecoin/types"
)

// ============================================================================
// Reserve State
// ============================================================================

// GetReserve retrieves the current reserve state
func (k Keeper) GetReserve(ctx sdk.Context) types.Reserve {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.ReserveKey) {
		return types.Reserve{
			TotalDeposited:    sdk.NewCoins(),
			TotalValue:        sdkmath.ZeroInt(),
			TotalMinted:       sdkmath.ZeroInt(),
			LastUpdatedHeight: ctx.BlockHeight(),
			LastUpdatedTime:   ctx.BlockTime(),
		}
	}
	var reserve types.Reserve
	types.ModuleCdc.MustUnmarshalJSON(store.Get(types.ReserveKey), &reserve)
	reserve.TotalMinted = k.getStablecoinSupply(ctx)
	return reserve
}

// SetReserve updates the reserve state
func (k Keeper) SetReserve(ctx sdk.Context, reserve types.Reserve) {
	reserve.TotalMinted = k.getStablecoinSupply(ctx)
	reserve.LastUpdatedHeight = ctx.BlockHeight()
	reserve.LastUpdatedTime = ctx.BlockTime()
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ReserveKey, types.ModuleCdc.MustMarshalJSON(&reserve))
}

// UpdateReserveValue recalculates reserve value from oracle prices
func (k Keeper) UpdateReserveValue(ctx sdk.Context) error {
	reserve := k.GetReserve(ctx)
	params := k.GetReserveParams(ctx)

	totalValue := sdkmath.ZeroInt()
	wrappedCtx := sdk.WrapSDKContext(ctx)

	for _, coin := range reserve.TotalDeposited {
		ttConfig, found := params.GetTokenizedTreasury(coin.Denom)
		if !found || !ttConfig.Active {
			continue
		}

		// Get price from oracle
		price, err := k.oracleKeeper.GetPriceDecSafe(wrappedCtx, ttConfig.OracleDenom)
		if err != nil {
			// Use fallback price of 1 USD only for cash-equivalent reserves.
			if ttConfig.UnderlyingType == types.ReserveAssetCash {
				price = sdkmath.LegacyOneDec()
			} else {
				// Skip assets without a reliable price to avoid overstating reserves.
				continue
			}
		}

		// Calculate value with haircut
		rawValue := price.MulInt(coin.Amount).TruncateInt()
		haircutMultiplier := sdkmath.LegacyNewDec(10000 - int64(ttConfig.HaircutBps)).Quo(sdkmath.LegacyNewDec(10000))
		adjustedValue := haircutMultiplier.MulInt(rawValue).TruncateInt()

		totalValue = totalValue.Add(adjustedValue)
	}

	reserve.TotalValue = totalValue
	k.SetReserve(ctx, reserve)
	return nil
}

func (k Keeper) getStablecoinSupply(ctx sdk.Context) sdkmath.Int {
	supply := k.bankKeeper.GetSupply(sdk.WrapSDKContext(ctx), types.StablecoinDenom)
	return supply.Amount
}

// ============================================================================
// Deposit Management
// ============================================================================

func (k Keeper) getNextDepositID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.NextDepositIDKey) {
		return 1
	}
	return binary.BigEndian.Uint64(store.Get(types.NextDepositIDKey))
}

func (k Keeper) setNextDepositID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextDepositIDKey, bz)
}

// GetReserveDeposit retrieves a reserve deposit by ID
func (k Keeper) GetReserveDeposit(ctx sdk.Context, id uint64) (types.ReserveDeposit, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ReserveDepositKeyPrefix)
	bz := store.Get(mustBz(id))
	if len(bz) == 0 {
		return types.ReserveDeposit{}, false
	}
	var deposit types.ReserveDeposit
	types.ModuleCdc.MustUnmarshalJSON(bz, &deposit)
	return deposit, true
}

func (k Keeper) setReserveDeposit(ctx sdk.Context, deposit types.ReserveDeposit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ReserveDepositKeyPrefix)
	store.Set(mustBz(deposit.Id), types.ModuleCdc.MustMarshalJSON(&deposit))
}

// IterateReserveDeposits iterates over all reserve deposits
func (k Keeper) IterateReserveDeposits(ctx sdk.Context, cb func(types.ReserveDeposit) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ReserveDepositKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var deposit types.ReserveDeposit
		types.ModuleCdc.MustUnmarshalJSON(iter.Value(), &deposit)
		if cb(deposit) {
			break
		}
	}
}

// DepositReserve deposits tokenized treasuries and mints ssUSD
func (k Keeper) DepositReserve(ctx sdk.Context, depositor sdk.AccAddress, amount sdk.Coin) (uint64, sdkmath.Int, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return 0, sdkmath.ZeroInt(), err
	}

	params := k.GetReserveParams(ctx)

	// Check if minting is paused
	if params.MintPaused {
		return 0, sdkmath.ZeroInt(), types.ErrMintPaused
	}

	// Validate tokenized treasury is approved
	ttConfig, found := params.GetTokenizedTreasury(amount.Denom)
	if !found {
		return 0, sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "denom %s not approved", amount.Denom)
	}
	if !ttConfig.Active {
		return 0, sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "denom %s is inactive", amount.Denom)
	}

	// Check KYC if required
	if params.RequireKyc {
		if err := k.complianceKeeper.AssertCompliant(wrappedCtx, depositor); err != nil {
			return 0, sdkmath.ZeroInt(), errorsmod.Wrap(types.ErrKYCRequired, err.Error())
		}
	}

	// Get price from oracle
	price, err := k.oracleKeeper.GetPriceDecSafe(wrappedCtx, ttConfig.OracleDenom)
	if err != nil {
		// Only allow a 1 USD fallback for cash-equivalent reserves.
		if ttConfig.UnderlyingType == types.ReserveAssetCash {
			price = sdkmath.LegacyOneDec()
		} else {
			if errors.Is(err, oracletypes.ErrPriceStale) {
				return 0, sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrPriceStale, "price for %s is stale", ttConfig.OracleDenom)
			}
			return 0, sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrPriceNotFound, "price not available for %s", ttConfig.OracleDenom)
		}
	}

	// Safety check: Price must be positive
	if price.IsNil() || !price.IsPositive() {
		return 0, sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrPriceNotFound, "invalid price %s for %s", price, ttConfig.OracleDenom)
	}

	// Calculate USD value (with haircut)
	rawValue := price.MulInt(amount.Amount).TruncateInt()
	haircutMultiplier := sdkmath.LegacyNewDec(10000 - int64(ttConfig.HaircutBps)).Quo(sdkmath.LegacyNewDec(10000))
	usdValue := haircutMultiplier.MulInt(rawValue).TruncateInt()

	// Calculate ssUSD to mint (after fee)
	feeMultiplier := sdkmath.LegacyNewDec(10000 - int64(params.MintFeeBps)).Quo(sdkmath.LegacyNewDec(10000))
	ssusdToMint := feeMultiplier.MulInt(usdValue).TruncateInt()

	// Check minimum mint amount
	if ssusdToMint.LT(params.MinMintAmount) {
		return 0, sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrBelowMinimumMint, "mint amount %s below minimum %s", ssusdToMint, params.MinMintAmount)
	}

	// Check daily limit (0 means no limit)
	dailyStats := k.GetDailyMintStats(ctx)
	if params.MaxDailyMint.IsPositive() && dailyStats.TotalMinted.Add(ssusdToMint).GT(params.MaxDailyMint) {
		return 0, sdkmath.ZeroInt(), types.ErrDailyMintLimitExceeded
	}

	// Check allocation limit
	if err := k.checkAllocationLimit(ctx, amount, ttConfig); err != nil {
		return 0, sdkmath.ZeroInt(), err
	}

	// Transfer tokenized treasury to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, depositor, types.ModuleAccountName, sdk.NewCoins(amount)); err != nil {
		return 0, sdkmath.ZeroInt(), err
	}

	// Mint ssUSD to depositor
	mintCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, ssusdToMint))
	if err := k.bankKeeper.MintCoins(wrappedCtx, types.ModuleAccountName, mintCoins); err != nil {
		return 0, sdkmath.ZeroInt(), err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, depositor, mintCoins); err != nil {
		return 0, sdkmath.ZeroInt(), err
	}

	// Update reserve state
	reserve := k.GetReserve(ctx)
	reserve.TotalDeposited = reserve.TotalDeposited.Add(amount)
	reserve.TotalValue = reserve.TotalValue.Add(usdValue)
	reserve.TotalMinted = reserve.TotalMinted.Add(ssusdToMint)
	k.SetReserve(ctx, reserve)

	// Create deposit record
	depositID := k.getNextDepositID(ctx)
	deposit := types.ReserveDeposit{
		Id:          depositID,
		Depositor:   depositor.String(),
		Amount:      amount,
		UsdValue:    usdValue,
		SsusdMinted: ssusdToMint,
		DepositedAt: ctx.BlockTime(),
		Status:      types.DepositStatusActive,
	}
	k.setReserveDeposit(ctx, deposit)
	k.setNextDepositID(ctx, depositID+1)

	// Update daily stats
	dailyStats.TotalMinted = dailyStats.TotalMinted.Add(ssusdToMint)
	k.SetDailyMintStats(ctx, dailyStats)

	// Emit event
	feeAmount := usdValue.Sub(ssusdToMint)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeReserveDeposit,
			sdk.NewAttribute(types.AttributeKeyDepositor, depositor.String()),
			sdk.NewAttribute(types.AttributeKeyDepositID, fmt.Sprintf("%d", depositID)),
			sdk.NewAttribute(types.AttributeKeyReserveAsset, amount.String()),
			sdk.NewAttribute(types.AttributeKeyPrice, price.String()),
			sdk.NewAttribute(types.AttributeKeyUsdValue, usdValue.String()),
			sdk.NewAttribute(types.AttributeKeySsusdAmount, ssusdToMint.String()),
			sdk.NewAttribute(types.AttributeKeyFeeAmount, feeAmount.String()),
			sdk.NewAttribute(types.AttributeKeyReserveRatio, fmt.Sprintf("%d", reserve.GetReserveRatio())),
		),
	)

	return depositID, ssusdToMint, nil
}

func (k Keeper) checkAllocationLimit(ctx sdk.Context, amount sdk.Coin, ttConfig types.TokenizedTreasuryConfig) error {
	reserve := k.GetReserve(ctx)

	currentAllocation := reserve.TotalDeposited.AmountOf(amount.Denom)
	newAllocation := currentAllocation.Add(amount.Amount)

	// Calculate total reserve value
	totalValue := reserve.TotalValue
	if totalValue.IsZero() {
		return nil // No limit check needed for first deposit
	}

	// Get price
	wrappedCtx := sdk.WrapSDKContext(ctx)
	price, err := k.oracleKeeper.GetPriceDecSafe(wrappedCtx, ttConfig.OracleDenom)
	if err != nil {
		if ttConfig.UnderlyingType == types.ReserveAssetCash {
			price = sdkmath.LegacyOneDec()
		} else {
			if errors.Is(err, oracletypes.ErrPriceStale) {
				return errorsmod.Wrapf(types.ErrPriceStale, "price for %s is stale", ttConfig.OracleDenom)
			}
			return errorsmod.Wrapf(types.ErrPriceNotFound, "price not available for %s", ttConfig.OracleDenom)
		}
	}

	// Allocation ratios should be computed against haircutted USD values so that
	// totals remain consistent with Reserve.TotalValue.
	haircutMultiplier := sdkmath.LegacyNewDec(10000 - int64(ttConfig.HaircutBps)).Quo(sdkmath.LegacyNewDec(10000))
	newAllocationValue := haircutMultiplier.MulInt(price.MulInt(newAllocation).TruncateInt()).TruncateInt()
	newDepositValue := haircutMultiplier.MulInt(price.MulInt(amount.Amount).TruncateInt()).TruncateInt()

	newTotalValue := totalValue.Add(newDepositValue)
	if newTotalValue.IsZero() {
		return nil
	}

	allocationRatio := newAllocationValue.Mul(sdkmath.NewInt(10000)).Quo(newTotalValue)

	if allocationRatio.GT(sdkmath.NewInt(int64(ttConfig.MaxAllocationBps))) {
		return errorsmod.Wrapf(types.ErrAllocationLimitExceeded,
			"allocation of %s would be %d bps, max is %d bps",
			amount.Denom, allocationRatio.Int64(), ttConfig.MaxAllocationBps)
	}

	return nil
}

// ============================================================================
// Locked Reserves (for pending redemptions)
// ============================================================================

// GetLockedReserves returns reserve assets locked for pending redemptions.
func (k Keeper) GetLockedReserves(ctx sdk.Context) sdk.Coins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LockedReservesKey)
	if len(bz) == 0 {
		return sdk.NewCoins()
	}

	var locked sdk.Coins
	if err := json.Unmarshal(bz, &locked); err != nil {
		panic(fmt.Sprintf("failed to unmarshal locked reserves: %v", err))
	}
	return locked.Sort()
}

func (k Keeper) setLockedReserves(ctx sdk.Context, locked sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	locked = locked.Sort()
	if len(locked) == 0 {
		store.Delete(types.LockedReservesKey)
		return
	}
	bz, err := json.Marshal(locked)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal locked reserves: %v", err))
	}
	store.Set(types.LockedReservesKey, bz)
}

func (k Keeper) lockReserves(ctx sdk.Context, coins sdk.Coins) error {
	if len(coins) == 0 {
		return nil
	}

	reserve := k.GetReserve(ctx)
	locked := k.GetLockedReserves(ctx)

	for _, coin := range coins {
		if !coin.IsValid() || !coin.Amount.IsPositive() {
			return errorsmod.Wrapf(types.ErrInvalidReserve, "invalid lock amount %s", coin)
		}

		total := reserve.TotalDeposited.AmountOf(coin.Denom)
		alreadyLocked := locked.AmountOf(coin.Denom)
		available := total.Sub(alreadyLocked)
		if available.IsNegative() {
			return errorsmod.Wrapf(types.ErrInvalidReserve, "locked reserves exceed total for %s", coin.Denom)
		}
		if coin.Amount.GT(available) {
			return errorsmod.Wrapf(
				types.ErrInsufficientReserves,
				"requested %s %s but only %s available (total %s, locked %s)",
				coin.Amount, coin.Denom, available, total, alreadyLocked,
			)
		}
	}

	locked = locked.Add(coins...).Sort()
	k.setLockedReserves(ctx, locked)
	return nil
}

func (k Keeper) unlockReserves(ctx sdk.Context, coins sdk.Coins) error {
	if len(coins) == 0 {
		return nil
	}

	locked := k.GetLockedReserves(ctx)
	for _, coin := range coins {
		if !coin.IsValid() || coin.Amount.IsNegative() {
			return errorsmod.Wrapf(types.ErrInvalidReserve, "invalid unlock amount %s", coin)
		}
		if coin.Amount.IsZero() {
			continue
		}
		lockedAmt := locked.AmountOf(coin.Denom)
		if lockedAmt.LT(coin.Amount) {
			return errorsmod.Wrapf(types.ErrInvalidReserve,
				"locked reserves underflow for %s: locked %s, unlocking %s",
				coin.Denom, lockedAmt, coin.Amount)
		}
	}

	locked = locked.Sub(coins...).Sort()
	k.setLockedReserves(ctx, locked)
	return nil
}

// ============================================================================
// Redemption Management
// ============================================================================

func (k Keeper) getNextRedemptionID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.NextRedemptionIDKey) {
		return 1
	}
	return binary.BigEndian.Uint64(store.Get(types.NextRedemptionIDKey))
}

func (k Keeper) setNextRedemptionID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextRedemptionIDKey, bz)
}

// GetRedemptionRequest retrieves a redemption request by ID
func (k Keeper) GetRedemptionRequest(ctx sdk.Context, id uint64) (types.RedemptionRequest, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RedemptionRequestKeyPrefix)
	bz := store.Get(mustBz(id))
	if len(bz) == 0 {
		return types.RedemptionRequest{}, false
	}
	var request types.RedemptionRequest
	types.ModuleCdc.MustUnmarshalJSON(bz, &request)
	return request, true
}

func (k Keeper) setRedemptionRequest(ctx sdk.Context, request types.RedemptionRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RedemptionRequestKeyPrefix)
	store.Set(mustBz(request.Id), types.ModuleCdc.MustMarshalJSON(&request))
}

// IterateRedemptionRequests iterates over all redemption requests
func (k Keeper) IterateRedemptionRequests(ctx sdk.Context, cb func(types.RedemptionRequest) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RedemptionRequestKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var request types.RedemptionRequest
		types.ModuleCdc.MustUnmarshalJSON(iter.Value(), &request)
		if cb(request) {
			break
		}
	}
}

// RequestRedemption requests redemption of ssUSD for tokenized treasuries
func (k Keeper) RequestRedemption(ctx sdk.Context, requester sdk.AccAddress, ssusdAmount sdkmath.Int, outputDenom string) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	params := k.GetReserveParams(ctx)

	// Check if redemption is paused
	if params.RedeemPaused {
		return 0, types.ErrRedeemPaused
	}

	// Validate output denom is approved
	ttConfig, found := params.GetTokenizedTreasury(outputDenom)
	if !found || !ttConfig.Active {
		return 0, errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "output denom %s not approved", outputDenom)
	}

	// Check KYC if required
	if params.RequireKyc {
		if err := k.complianceKeeper.AssertCompliant(wrappedCtx, requester); err != nil {
			return 0, errorsmod.Wrap(types.ErrKYCRequired, err.Error())
		}
	}

	// Check minimum redeem amount
	if ssusdAmount.LT(params.MinRedeemAmount) {
		return 0, errorsmod.Wrapf(types.ErrBelowMinimumRedeem, "redeem amount %s below minimum %s", ssusdAmount, params.MinRedeemAmount)
	}

	// Check daily limit (0 means no limit)
	dailyStats := k.GetDailyMintStats(ctx)
	if params.MaxDailyRedeem.IsPositive() && dailyStats.TotalRedeemed.Add(ssusdAmount).GT(params.MaxDailyRedeem) {
		return 0, types.ErrDailyRedeemLimitExceeded
	}

	// Check reserve has sufficient output denom (net of pending locks).
	reserve := k.GetReserve(ctx)
	locked := k.GetLockedReserves(ctx)
	outputAvailable := reserve.TotalDeposited.AmountOf(outputDenom).Sub(locked.AmountOf(outputDenom))
	if outputAvailable.IsNegative() {
		return 0, errorsmod.Wrapf(types.ErrInvalidReserve, "locked reserves exceed total for %s", outputDenom)
	}

	// Calculate output amount (apply fee)
	feeMultiplier := sdkmath.LegacyNewDec(10000 - int64(params.RedeemFeeBps)).Quo(sdkmath.LegacyNewDec(10000))
	ssusdAfterFee := feeMultiplier.MulInt(ssusdAmount).TruncateInt()

	// Get output price
	price, err := k.oracleKeeper.GetPriceDecSafe(wrappedCtx, ttConfig.OracleDenom)
	if err != nil {
		if ttConfig.UnderlyingType == types.ReserveAssetCash {
			price = sdkmath.LegacyOneDec()
		} else {
			if errors.Is(err, oracletypes.ErrPriceStale) {
				return 0, errorsmod.Wrapf(types.ErrPriceStale, "price for %s is stale", ttConfig.OracleDenom)
			}
			return 0, errorsmod.Wrapf(types.ErrPriceNotFound, "price not available for %s", ttConfig.OracleDenom)
		}
	}

	// Calculate output tokens needed
	outputAmount := sdkmath.LegacyNewDecFromInt(ssusdAfterFee).Quo(price).TruncateInt()

	if outputAmount.GT(outputAvailable) {
		return 0, errorsmod.Wrapf(types.ErrInsufficientReserves,
			"requested %s %s but only %s available (locked %s)", outputAmount, outputDenom, outputAvailable, locked.AmountOf(outputDenom))
	}

	// Transfer ssUSD to module (burned at request time)
	ssusdCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, ssusdAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, requester, types.ModuleAccountName, ssusdCoins); err != nil {
		return 0, err
	}

	// Burn ssUSD immediately
	if err := k.bankKeeper.BurnCoins(wrappedCtx, types.ModuleAccountName, ssusdCoins); err != nil {
		return 0, err
	}

	// Create redemption request
	redemptionID := k.getNextRedemptionID(ctx)
	executableAfter := ctx.BlockTime().Add(params.RedemptionDelay)

	request := types.RedemptionRequest{
		Id:              redemptionID,
		Requester:       requester.String(),
		SsusdAmount:     ssusdAmount,
		OutputDenom:     outputDenom,
		RequestedAt:     ctx.BlockTime(),
		ExecutableAfter: executableAfter,
		Status:          types.RedeemStatusPending,
		OutputAmount:    sdk.NewCoin(outputDenom, outputAmount),
	}

	// Store request and bump ID regardless of delay so IDs/statistics remain consistent.
	k.setRedemptionRequest(ctx, request)
	k.setNextRedemptionID(ctx, redemptionID+1)

	// Lock reserves so delayed redemptions cannot be starved by later withdrawals.
	if err := k.lockReserves(ctx, sdk.NewCoins(request.OutputAmount)); err != nil {
		return 0, err
	}

	// Update daily stats regardless of delay (requests are counted when made).
	dailyStats.TotalRedeemed = dailyStats.TotalRedeemed.Add(ssusdAmount)
	k.SetDailyMintStats(ctx, dailyStats)

	// Emit requested event regardless of delay.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRedemptionRequested,
			sdk.NewAttribute(types.AttributeKeyDepositor, requester.String()),
			sdk.NewAttribute(types.AttributeKeyRedemptionID, fmt.Sprintf("%d", redemptionID)),
			sdk.NewAttribute(types.AttributeKeySsusdAmount, ssusdAmount.String()),
			sdk.NewAttribute(types.AttributeKeyReserveAsset, outputDenom),
		),
	)

	// If no delay, execute immediately after storing.
	if params.RedemptionDelay == 0 {
		_, err := k.executeRedemption(ctx, &request, outputAmount)
		return redemptionID, err
	}

	return redemptionID, nil
}

func (k Keeper) executeRedemption(ctx sdk.Context, request *types.RedemptionRequest, outputAmount sdkmath.Int) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	requester, err := sdk.AccAddressFromBech32(request.Requester)
	if err != nil {
		return 0, err
	}

	// Transfer output tokens to requester
	outputCoins := sdk.NewCoins(sdk.NewCoin(request.OutputDenom, outputAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, requester, outputCoins); err != nil {
		return 0, err
	}

	// Unlock reserved output now that it has been released.
	if err := k.unlockReserves(ctx, outputCoins); err != nil {
		return 0, err
	}

	// Update reserve state
	reserve := k.GetReserve(ctx)
	reserve.TotalDeposited = reserve.TotalDeposited.Sub(outputCoins...)
	reserve.TotalMinted = reserve.TotalMinted.Sub(request.SsusdAmount)
	k.SetReserve(ctx, reserve)

	// Recalculate reserve value
	k.UpdateReserveValue(ctx)

	// Update request status
	request.Status = types.RedeemStatusExecuted
	request.ExecutedAt = ctx.BlockTime()
	request.OutputAmount = sdk.NewCoin(request.OutputDenom, outputAmount)
	k.setRedemptionRequest(ctx, *request)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRedemptionExecuted,
			sdk.NewAttribute(types.AttributeKeyDepositor, request.Requester),
			sdk.NewAttribute(types.AttributeKeyRedemptionID, fmt.Sprintf("%d", request.Id)),
			sdk.NewAttribute(types.AttributeKeyAmount, outputCoins.String()),
		),
	)

	return request.Id, nil
}

// ExecutePendingRedemption executes a pending redemption after delay
func (k Keeper) ExecutePendingRedemption(ctx sdk.Context, redemptionID uint64) error {
	request, found := k.GetRedemptionRequest(ctx, redemptionID)
	if !found {
		return types.ErrRedemptionNotFound
	}

	if request.Status != types.RedeemStatusPending {
		return errorsmod.Wrapf(types.ErrRedemptionNotFound, "redemption %d is not pending", redemptionID)
	}

	if ctx.BlockTime().Before(request.ExecutableAfter) {
		return errorsmod.Wrapf(types.ErrRedemptionNotReady,
			"redemption executable after %s, current time %s",
			request.ExecutableAfter, ctx.BlockTime())
	}

	params := k.GetReserveParams(ctx)
	ttConfig, found := params.GetTokenizedTreasury(request.OutputDenom)
	if !found || !ttConfig.Active {
		return errorsmod.Wrapf(types.ErrUnsupportedReserveAsset, "output denom %s not approved", request.OutputDenom)
	}

	wrappedCtx := sdk.WrapSDKContext(ctx)
	// Re-check compliance if required.
	if params.RequireKyc {
		requesterAddr, err := sdk.AccAddressFromBech32(request.Requester)
		if err != nil {
			return errorsmod.Wrap(types.ErrInvalidReserve, "invalid requester address")
		}
		if err := k.complianceKeeper.AssertCompliant(wrappedCtx, requesterAddr); err != nil {
			return errorsmod.Wrap(types.ErrKYCRequired, err.Error())
		}
	}

	if request.OutputAmount.Denom != request.OutputDenom || request.OutputAmount.Amount.IsNegative() {
		return errorsmod.Wrap(types.ErrInvalidReserve, "invalid stored output amount")
	}
	if request.OutputAmount.Amount.IsZero() {
		return errorsmod.Wrap(types.ErrInvalidReserve, "missing stored output amount")
	}

	_, err := k.executeRedemption(ctx, &request, request.OutputAmount.Amount)
	return err
}

// CancelRedemption cancels a pending redemption (authority only)
func (k Keeper) CancelRedemption(ctx sdk.Context, authority string, redemptionID uint64) error {
	request, found := k.GetRedemptionRequest(ctx, redemptionID)
	if !found {
		return types.ErrRedemptionNotFound
	}

	if request.Status != types.RedeemStatusPending {
		return errorsmod.Wrapf(types.ErrRedemptionNotFound, "redemption %d is not pending", redemptionID)
	}

	// Verify authority
	if authority != k.GetAuthority() {
		return errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", k.GetAuthority(), authority)
	}

	// Release locked reserves back to the available pool.
	if request.OutputAmount.Denom != "" && request.OutputAmount.Amount.IsPositive() {
		if err := k.unlockReserves(ctx, sdk.NewCoins(request.OutputAmount)); err != nil {
			return err
		}
	}

	// Refund ssUSD to requester
	wrappedCtx := sdk.WrapSDKContext(ctx)
	requester, _ := sdk.AccAddressFromBech32(request.Requester)
	ssusdCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, request.SsusdAmount))

	// Mint back the ssUSD (it was already burned)
	if err := k.bankKeeper.MintCoins(wrappedCtx, types.ModuleAccountName, ssusdCoins); err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, requester, ssusdCoins); err != nil {
		return err
	}

	request.Status = types.RedeemStatusCancelled
	k.setRedemptionRequest(ctx, request)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRedemptionCancelled,
			sdk.NewAttribute(types.AttributeKeyRedemptionID, fmt.Sprintf("%d", redemptionID)),
		),
	)

	return nil
}

// ============================================================================
// Daily Stats
// ============================================================================

// GetDailyMintStats gets daily mint/redeem stats (resets daily)
func (k Keeper) GetDailyMintStats(ctx sdk.Context) types.DailyMintStats {
	today := ctx.BlockTime().Format("2006-01-02")
	store := ctx.KVStore(k.storeKey)
	key := types.DailyMintStatsKey(today)

	if !store.Has(key) {
		return types.DailyMintStats{
			Date:          today,
			TotalMinted:   sdkmath.ZeroInt(),
			TotalRedeemed: sdkmath.ZeroInt(),
		}
	}

	var stats types.DailyMintStats
	types.ModuleCdc.MustUnmarshalJSON(store.Get(key), &stats)
	return stats
}

// SetDailyMintStats updates daily mint stats
func (k Keeper) SetDailyMintStats(ctx sdk.Context, stats types.DailyMintStats) {
	store := ctx.KVStore(k.storeKey)
	key := types.DailyMintStatsKey(stats.Date)
	store.Set(key, types.ModuleCdc.MustMarshalJSON(&stats))
}

// ============================================================================
// Off-Chain Attestations
// ============================================================================

func (k Keeper) getNextAttestationID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.NextAttestationIDKey) {
		return 1
	}
	return binary.BigEndian.Uint64(store.Get(types.NextAttestationIDKey))
}

func (k Keeper) setNextAttestationID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextAttestationIDKey, bz)
}

// IsApprovedAttester checks if an address is an approved attester
func (k Keeper) IsApprovedAttester(ctx sdk.Context, addr string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ApprovedAttesterKey(addr))
}

// SetApprovedAttester adds or removes an approved attester
func (k Keeper) SetApprovedAttester(ctx sdk.Context, addr string, approved bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ApprovedAttesterKey(addr)
	if approved {
		store.Set(key, []byte{1})
	} else {
		store.Delete(key)
	}
}

// GetLatestAttestation retrieves the latest off-chain attestation
func (k Keeper) GetLatestAttestation(ctx sdk.Context) (types.OffChainReserveAttestation, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OffChainAttestationKeyPrefix)
	iter := store.ReverseIterator(nil, nil)
	defer iter.Close()

	if !iter.Valid() {
		return types.OffChainReserveAttestation{}, false
	}

	var attestation types.OffChainReserveAttestation
	types.ModuleCdc.MustUnmarshalJSON(iter.Value(), &attestation)
	return attestation, true
}

// GetAttestation retrieves an attestation by ID
func (k Keeper) GetAttestation(ctx sdk.Context, id uint64) (types.OffChainReserveAttestation, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OffChainAttestationKeyPrefix)
	bz := store.Get(mustBz(id))
	if len(bz) == 0 {
		return types.OffChainReserveAttestation{}, false
	}
	var attestation types.OffChainReserveAttestation
	types.ModuleCdc.MustUnmarshalJSON(bz, &attestation)
	return attestation, true
}

// RecordAttestation records an off-chain reserve attestation
func (k Keeper) RecordAttestation(ctx sdk.Context, attestation types.OffChainReserveAttestation) (uint64, error) {
	if err := attestation.ValidateBasic(); err != nil {
		return 0, err
	}

	if !k.IsApprovedAttester(ctx, attestation.Attester) {
		return 0, errorsmod.Wrapf(types.ErrInvalidAttester, "attester %s is not approved", attestation.Attester)
	}

	attestationID := k.getNextAttestationID(ctx)
	attestation.Id = attestationID
	attestation.Timestamp = ctx.BlockTime()

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OffChainAttestationKeyPrefix)
	store.Set(mustBz(attestationID), types.ModuleCdc.MustMarshalJSON(&attestation))
	k.setNextAttestationID(ctx, attestationID+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeReserveAttestation,
			sdk.NewAttribute(types.AttributeKeyAttester, attestation.Attester),
			sdk.NewAttribute(types.AttributeKeyUsdValue, attestation.TotalValue.String()),
		),
	)

	return attestationID, nil
}

// IterateAttestations iterates over all attestations
func (k Keeper) IterateAttestations(ctx sdk.Context, cb func(types.OffChainReserveAttestation) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OffChainAttestationKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var attestation types.OffChainReserveAttestation
		types.ModuleCdc.MustUnmarshalJSON(iter.Value(), &attestation)
		if cb(attestation) {
			break
		}
	}
}

// ============================================================================
// Total Reserves Query
// ============================================================================

// GetTotalReserves calculates total reserves (on-chain + off-chain)
func (k Keeper) GetTotalReserves(ctx sdk.Context) types.TotalReserves {
	reserve := k.GetReserve(ctx)
	latestAttestation, hasAttestation := k.GetLatestAttestation(ctx)

	offChainValue := sdkmath.ZeroInt()
	var lastOffChainUpdate time.Time
	if hasAttestation {
		offChainValue = latestAttestation.TotalValue
		lastOffChainUpdate = latestAttestation.Timestamp
	}

	totalValue := reserve.TotalValue.Add(offChainValue)

		// TotalMinted is sourced from the bank module supply via GetReserve.
		totalSupply := reserve.TotalMinted

	lastOnChainUpdate := reserve.LastUpdatedTime
	if lastOnChainUpdate.IsZero() {
		lastOnChainUpdate = ctx.BlockTime()
	}

	totalReserves := types.TotalReserves{
		OnChainValue:       reserve.TotalValue,
		OffChainValue:      offChainValue,
		TotalValue:         totalValue,
		TotalSupply:        totalSupply,
		LastOnChainUpdate:  lastOnChainUpdate,
		LastOffChainUpdate: lastOffChainUpdate,
	}
	totalReserves.ReserveRatioBps = totalReserves.CalculateReserveRatio()

	return totalReserves
}
