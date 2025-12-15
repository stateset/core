package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stablecoin/types"
)

// ============================================================================
// Auction Parameters
// ============================================================================

// GetAuctionParams retrieves auction parameters.
func (k Keeper) GetAuctionParams(ctx sdk.Context) types.AuctionParams {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AuctionParamsKey)
	if len(bz) == 0 {
		return DefaultAuctionParams()
	}
	var params types.AuctionParams
	types.MustUnmarshalJSON(bz, &params)
	return params
}

// SetAuctionParams stores auction parameters.
func (k Keeper) SetAuctionParams(ctx sdk.Context, params types.AuctionParams) error {
	if params.StartPriceMultiplierBps < 10000 {
		return errorsmod.Wrap(types.ErrInvalidReserve, "start price multiplier must be >= 100%")
	}
	if params.EndPriceMultiplierBps > params.StartPriceMultiplierBps {
		return errorsmod.Wrap(types.ErrInvalidReserve, "end price must be <= start price")
	}
	if params.LiquidationPenaltyBps > 5000 {
		return errorsmod.Wrap(types.ErrInvalidReserve, "liquidation penalty cannot exceed 50%")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.AuctionParamsKey, types.MustMarshalJSON(params))
	return nil
}

// DefaultAuctionParams returns default auction parameters.
func DefaultAuctionParams() types.AuctionParams {
	return types.AuctionParams{
		Enabled:                  true,
		Duration:                 time.Hour * 6, // 6 hour auctions
		StartPriceMultiplierBps:  13000,         // Start at 130% of oracle price
		EndPriceMultiplierBps:    8000,          // End at 80% of oracle price
		LiquidationPenaltyBps:    1300,          // 13% penalty
	}
}

// ============================================================================
// Auction State Management
// ============================================================================

func (k Keeper) getNextAuctionID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.NextAuctionIDKey) {
		return 1
	}
	return binary.BigEndian.Uint64(store.Get(types.NextAuctionIDKey))
}

func (k Keeper) setNextAuctionID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextAuctionIDKey, bz)
}

// GetAuction retrieves an auction by ID.
func (k Keeper) GetAuction(ctx sdk.Context, id uint64) (types.DutchAuction, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AuctionKeyPrefix)
	bz := store.Get(mustBz(id))
	if len(bz) == 0 {
		return types.DutchAuction{}, false
	}
	var auction types.DutchAuction
	types.MustUnmarshalJSON(bz, &auction)
	return auction, true
}

// SetAuction stores an auction.
func (k Keeper) SetAuction(ctx sdk.Context, auction types.DutchAuction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AuctionKeyPrefix)
	store.Set(mustBz(auction.Id), types.MustMarshalJSON(auction))

	// Track active auctions
	if auction.Status == types.AuctionStatus_AUCTION_STATUS_ACTIVE {
		activeStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveAuctionKeyPrefix)
		activeStore.Set(mustBz(auction.Id), []byte{1})
	} else {
		k.removeActiveAuction(ctx, auction.Id)
	}
}

func (k Keeper) removeActiveAuction(ctx sdk.Context, id uint64) {
	activeStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveAuctionKeyPrefix)
	activeStore.Delete(mustBz(id))
}

// IterateAuctions iterates over all auctions.
func (k Keeper) IterateAuctions(ctx sdk.Context, cb func(types.DutchAuction) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AuctionKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var auction types.DutchAuction
		types.MustUnmarshalJSON(iter.Value(), &auction)
		if cb(auction) {
			break
		}
	}
}

// IterateActiveAuctions iterates over active auctions.
func (k Keeper) IterateActiveAuctions(ctx sdk.Context, cb func(types.DutchAuction) bool) {
	activeStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveAuctionKeyPrefix)
	iter := activeStore.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		id := binary.BigEndian.Uint64(iter.Key())
		auction, found := k.GetAuction(ctx, id)
		if found && cb(auction) {
			break
		}
	}
}

// ============================================================================
// Dutch Auction Operations
// ============================================================================

// CreateAuction creates a new Dutch auction for liquidated collateral.
func (k Keeper) CreateAuction(ctx sdk.Context, vaultID uint64, owner string, collateral sdk.Coin, debtToCover sdkmath.Int) (uint64, error) {
	params := k.GetAuctionParams(ctx)
	if !params.Enabled {
		return 0, errorsmod.Wrap(types.ErrMintPaused, "Dutch auctions are disabled")
	}

	// Get oracle price for collateral
	wrappedCtx := sdk.WrapSDKContext(ctx)
	price, err := k.oracleKeeper.GetPriceDecSafe(wrappedCtx, collateral.Denom)
	if err != nil {
		return 0, errorsmod.Wrapf(types.ErrPriceNotFound, "cannot create auction without oracle price for %s", collateral.Denom)
	}

	// Calculate start and end prices
	startMultiplier := sdkmath.LegacyNewDec(int64(params.StartPriceMultiplierBps)).Quo(sdkmath.LegacyNewDec(10000))
	endMultiplier := sdkmath.LegacyNewDec(int64(params.EndPriceMultiplierBps)).Quo(sdkmath.LegacyNewDec(10000))

	startPrice := price.Mul(startMultiplier)
	endPrice := price.Mul(endMultiplier)

	auctionID := k.getNextAuctionID(ctx)
	auction := types.DutchAuction{
		Id:             auctionID,
		VaultId:        vaultID,
		Owner:          owner,
		Collateral:     collateral,
		DebtToCover:    debtToCover,
		StartPrice:     startPrice,
		EndPrice:       endPrice,
		StartedAt:      ctx.BlockTime(),
		Duration:       params.Duration,
		Status:         types.AuctionStatus_AUCTION_STATUS_ACTIVE,
		CollateralSold: sdkmath.ZeroInt(),
		DebtRaised:     sdkmath.ZeroInt(),
	}

	k.SetAuction(ctx, auction)
	k.setNextAuctionID(ctx, auctionID+1)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAuctionCreated,
			sdk.NewAttribute(types.AttributeKeyAuctionID, fmt.Sprintf("%d", auctionID)),
			sdk.NewAttribute(types.AttributeKeyVaultID, fmt.Sprintf("%d", vaultID)),
			sdk.NewAttribute(types.AttributeKeyOwner, owner),
			sdk.NewAttribute(types.AttributeKeyCollateral, collateral.String()),
			sdk.NewAttribute(types.AttributeKeyDebtToCover, debtToCover.String()),
			sdk.NewAttribute(types.AttributeKeyStartPrice, startPrice.String()),
			sdk.NewAttribute(types.AttributeKeyEndPrice, endPrice.String()),
		),
	)

	return auctionID, nil
}

// GetCurrentAuctionPrice calculates the current price in a Dutch auction.
func (k Keeper) GetCurrentAuctionPrice(ctx sdk.Context, auction types.DutchAuction) sdkmath.LegacyDec {
	now := ctx.BlockTime()
	elapsed := now.Sub(auction.StartedAt)
	duration := auction.Duration

	if elapsed >= duration {
		return auction.EndPrice
	}

	// Linear price decay: price = startPrice - (startPrice - endPrice) * (elapsed / duration)
	elapsedRatio := sdkmath.LegacyNewDec(int64(elapsed)).Quo(sdkmath.LegacyNewDec(int64(duration)))
	priceDiff := auction.StartPrice.Sub(auction.EndPrice)
	currentPrice := auction.StartPrice.Sub(priceDiff.Mul(elapsedRatio))

	return currentPrice
}

// BidAuction places a bid on a Dutch auction.
func (k Keeper) BidAuction(ctx sdk.Context, bidder sdk.AccAddress, auctionID uint64, maxCollateral, maxSSUSD sdkmath.Int) (sdkmath.Int, sdkmath.Int, sdkmath.LegacyDec, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	if err := k.ensureModuleAccount(ctx); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), err
	}

	auction, found := k.GetAuction(ctx, auctionID)
	if !found {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), errorsmod.Wrap(types.ErrVaultNotFound, "auction not found")
	}

	if auction.Status != types.AuctionStatus_AUCTION_STATUS_ACTIVE {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), errorsmod.Wrap(types.ErrVaultNotFound, "auction is not active")
	}

	// Check if auction has expired
	if ctx.BlockTime().After(auction.StartedAt.Add(auction.Duration)) {
		auction.Status = types.AuctionStatus_AUCTION_STATUS_EXPIRED
		k.SetAuction(ctx, auction)
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), errorsmod.Wrap(types.ErrVaultNotFound, "auction has expired")
	}

	// Get current price
	currentPrice := k.GetCurrentAuctionPrice(ctx, auction)

	// Calculate how much collateral remains
	remainingCollateral := auction.Collateral.Amount.Sub(auction.CollateralSold)
	if remainingCollateral.IsZero() {
		auction.Status = types.AuctionStatus_AUCTION_STATUS_COMPLETED
		k.SetAuction(ctx, auction)
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), errorsmod.Wrap(types.ErrVaultNotFound, "auction has no remaining collateral")
	}

	// Calculate remaining debt to cover
	remainingDebt := auction.DebtToCover.Sub(auction.DebtRaised)

	// Determine collateral to purchase
	collateralToPurchase := maxCollateral
	if collateralToPurchase.GT(remainingCollateral) {
		collateralToPurchase = remainingCollateral
	}

	// Calculate ssUSD cost
	ssusdCost := currentPrice.MulInt(collateralToPurchase).TruncateInt()
	if ssusdCost.GT(maxSSUSD) {
		// Adjust collateral to match max ssUSD
		collateralToPurchase = sdkmath.LegacyNewDecFromInt(maxSSUSD).Quo(currentPrice).TruncateInt()
		ssusdCost = currentPrice.MulInt(collateralToPurchase).TruncateInt()
	}

	if collateralToPurchase.IsZero() {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), errorsmod.Wrap(types.ErrInvalidAmount, "calculated collateral purchase is zero")
	}

	// Cap at remaining debt if debt is almost covered
	if ssusdCost.GT(remainingDebt) && remainingDebt.IsPositive() {
		ssusdCost = remainingDebt
		collateralToPurchase = sdkmath.LegacyNewDecFromInt(ssusdCost).Quo(currentPrice).TruncateInt()
	}

	// Transfer ssUSD from bidder to module and burn
	ssusdCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, ssusdCost))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, bidder, types.ModuleAccountName, ssusdCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), err
	}
	if err := k.bankKeeper.BurnCoins(wrappedCtx, types.ModuleAccountName, ssusdCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), err
	}

	// Transfer collateral to bidder
	collateralCoins := sdk.NewCoins(sdk.NewCoin(auction.Collateral.Denom, collateralToPurchase))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, bidder, collateralCoins); err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdkmath.LegacyZeroDec(), err
	}

	// Update auction state
	auction.CollateralSold = auction.CollateralSold.Add(collateralToPurchase)
	auction.DebtRaised = auction.DebtRaised.Add(ssusdCost)

	// Check if auction is complete
	newRemainingCollateral := auction.Collateral.Amount.Sub(auction.CollateralSold)
	newRemainingDebt := auction.DebtToCover.Sub(auction.DebtRaised)

	if newRemainingCollateral.IsZero() || newRemainingDebt.IsZero() {
		auction.Status = types.AuctionStatus_AUCTION_STATUS_COMPLETED

		// Return any remaining collateral to original owner
		if newRemainingCollateral.IsPositive() {
			ownerAddr, err := sdk.AccAddressFromBech32(auction.Owner)
			if err == nil {
				remainingCoins := sdk.NewCoins(sdk.NewCoin(auction.Collateral.Denom, newRemainingCollateral))
				k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, ownerAddr, remainingCoins)
			}
		}

		// Emit completion event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeAuctionCompleted,
				sdk.NewAttribute(types.AttributeKeyAuctionID, fmt.Sprintf("%d", auctionID)),
				sdk.NewAttribute(types.AttributeKeyCollateralSold, auction.CollateralSold.String()),
				sdk.NewAttribute(types.AttributeKeyDebtRaised, auction.DebtRaised.String()),
			),
		)
	}

	k.SetAuction(ctx, auction)

	// Emit bid event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAuctionBid,
			sdk.NewAttribute(types.AttributeKeyAuctionID, fmt.Sprintf("%d", auctionID)),
			sdk.NewAttribute(types.AttributeKeyBidder, bidder.String()),
			sdk.NewAttribute(types.AttributeKeyCollateral, collateralToPurchase.String()),
			sdk.NewAttribute(types.AttributeKeySsusdAmount, ssusdCost.String()),
			sdk.NewAttribute(types.AttributeKeyCurrentPrice, currentPrice.String()),
		),
	)

	return collateralToPurchase, ssusdCost, currentPrice, nil
}

// LiquidateVaultWithAuction liquidates a vault by creating a Dutch auction.
// This replaces the instant liquidation method.
func (k Keeper) LiquidateVaultWithAuction(ctx sdk.Context, liquidator sdk.AccAddress, vaultID uint64) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	vault, found := k.GetVault(ctx, vaultID)
	if !found {
		return 0, types.ErrVaultNotFound
	}

	params := k.GetParams(ctx)
	cp, ok := params.GetCollateralParam(vault.CollateralDenom)
	if !ok {
		return 0, types.ErrUnsupportedCollateral
	}

	// Check if vault is under-collateralized
	if err := k.assertCollateralization(ctx, vault.Collateral, vault.Debt, cp); err == nil {
		return 0, errorsmod.Wrap(types.ErrVaultHealthy, "vault still healthy")
	}

	auctionParams := k.GetAuctionParams(ctx)
	if !auctionParams.Enabled {
		// Fall back to instant liquidation if auctions are disabled
		_, err := k.LiquidateVault(ctx, liquidator, vaultID)
		return 0, err
	}

	// Calculate debt with liquidation penalty
	penaltyMultiplier := sdkmath.LegacyNewDec(10000 + int64(auctionParams.LiquidationPenaltyBps)).Quo(sdkmath.LegacyNewDec(10000))
	debtWithPenalty := penaltyMultiplier.MulInt(vault.Debt).TruncateInt()

	// Transfer collateral from module (it's already there from vault creation)
	// The collateral stays in the module account during the auction

	// Create the auction
	auctionID, err := k.CreateAuction(ctx, vaultID, vault.Owner, vault.Collateral, debtWithPenalty)
	if err != nil {
		return 0, err
	}

	// Remove the vault (debt will be covered by auction proceeds)
	k.removeVault(ctx, vaultID)

	// Emit liquidation event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVaultLiquidated,
			sdk.NewAttribute(sdk.AttributeKeySender, liquidator.String()),
			sdk.NewAttribute(types.AttributeKeyLiquidator, liquidator.String()),
			sdk.NewAttribute(types.AttributeKeyVaultID, fmt.Sprintf("%d", vaultID)),
			sdk.NewAttribute(types.AttributeKeyAuctionID, fmt.Sprintf("%d", auctionID)),
			sdk.NewAttribute(types.AttributeKeyCollateral, vault.Collateral.String()),
			sdk.NewAttribute(types.AttributeKeyDebtToCover, debtWithPenalty.String()),
		),
	)

	// Reward liquidator with a small incentive (optional - can be disabled)
	// This ensures liquidators are incentivized to trigger liquidations
	_ = wrappedCtx // Use if liquidator incentive is implemented

	return auctionID, nil
}

// ProcessExpiredAuctions processes expired auctions in the EndBlocker.
func (k Keeper) ProcessExpiredAuctions(ctx sdk.Context) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	var expiredAuctions []types.DutchAuction
	k.IterateActiveAuctions(ctx, func(auction types.DutchAuction) bool {
		if ctx.BlockTime().After(auction.StartedAt.Add(auction.Duration)) {
			expiredAuctions = append(expiredAuctions, auction)
		}
		return false
	})

	for _, auction := range expiredAuctions {
		auction.Status = types.AuctionStatus_AUCTION_STATUS_EXPIRED

		// Return remaining collateral to original owner
		remainingCollateral := auction.Collateral.Amount.Sub(auction.CollateralSold)
		if remainingCollateral.IsPositive() {
			ownerAddr, err := sdk.AccAddressFromBech32(auction.Owner)
			if err == nil {
				remainingCoins := sdk.NewCoins(sdk.NewCoin(auction.Collateral.Denom, remainingCollateral))
				k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, ownerAddr, remainingCoins)
			}
		}

		k.SetAuction(ctx, auction)

		// Emit expiry event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeAuctionExpired,
				sdk.NewAttribute(types.AttributeKeyAuctionID, fmt.Sprintf("%d", auction.Id)),
				sdk.NewAttribute(types.AttributeKeyCollateralSold, auction.CollateralSold.String()),
				sdk.NewAttribute(types.AttributeKeyDebtRaised, auction.DebtRaised.String()),
			),
		)
	}
}

// UpdateAuctionParams updates auction parameters (governance only).
func (k Keeper) UpdateAuctionParams(ctx sdk.Context, authority string, params types.AuctionParams) error {
	if authority != k.GetAuthority() {
		return errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", k.GetAuthority(), authority)
	}
	return k.SetAuctionParams(ctx, params)
}

// GetActiveAuctions returns all active auctions.
func (k Keeper) GetActiveAuctions(ctx sdk.Context) []types.DutchAuction {
	var auctions []types.DutchAuction
	k.IterateActiveAuctions(ctx, func(auction types.DutchAuction) bool {
		auctions = append(auctions, auction)
		return false
	})
	return auctions
}
