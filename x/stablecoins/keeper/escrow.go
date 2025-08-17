package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/x/stablecoins/types"
)

// EscrowAccount represents an escrow account for order payments
type EscrowAccount struct {
	OrderID     string    `json:"order_id"`
	Amount      sdk.Coin  `json:"amount"`
	CustomerAddr string   `json:"customer_addr"`
	MerchantAddr string   `json:"merchant_addr"`
	CreatedAt   time.Time `json:"created_at"`
	Timeout     time.Time `json:"timeout"`
	Released    bool      `json:"released"`
}

// EscrowStablecoin transfers stablecoins to an escrow account for an order
func (k Keeper) EscrowStablecoin(ctx sdk.Context, from sdk.AccAddress, orderId string, amount sdk.Coin) error {
	// Check if escrow already exists for this order
	if k.hasEscrow(ctx, orderId) {
		return sdkerrors.Wrap(types.ErrEscrowAlreadyExists, orderId)
	}

	// Validate stablecoin
	if !k.IsValidStablecoin(ctx, amount.Denom) {
		return sdkerrors.Wrap(types.ErrInvalidStablecoin, amount.Denom)
	}

	// Check balance
	balance := k.bankKeeper.GetBalance(ctx, from, amount.Denom)
	if balance.IsLT(amount) {
		return sdkerrors.Wrap(types.ErrInsufficientBalance, fmt.Sprintf("required: %s, available: %s", amount.String(), balance.String()))
	}

	// Transfer to module account (escrow)
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddr == nil {
		return sdkerrors.Wrap(types.ErrModuleAccountNotFound, types.ModuleName)
	}

	err := k.bankKeeper.SendCoins(ctx, from, moduleAddr, sdk.NewCoins(amount))
	if err != nil {
		return sdkerrors.Wrap(types.ErrEscrowTransferFailed, err.Error())
	}

	// Store escrow information
	escrow := EscrowAccount{
		OrderID:      orderId,
		Amount:       amount,
		CustomerAddr: from.String(),
		CreatedAt:    ctx.BlockTime(),
		Timeout:      ctx.BlockTime().Add(time.Hour * 24 * 30), // 30 days default timeout
		Released:     false,
	}
	
	k.setEscrow(ctx, orderId, escrow)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_escrowed",
			sdk.NewAttribute("order_id", orderId),
			sdk.NewAttribute("customer", from.String()),
			sdk.NewAttribute("amount", amount.String()),
			sdk.NewAttribute("timeout", escrow.Timeout.String()),
		),
	)

	return nil
}

// ReleaseEscrow releases escrowed stablecoins to the merchant
func (k Keeper) ReleaseEscrow(ctx sdk.Context, orderId string, to sdk.AccAddress) error {
	// Get escrow
	escrow, found := k.getEscrow(ctx, orderId)
	if !found {
		return sdkerrors.Wrap(types.ErrEscrowNotFound, orderId)
	}

	// Check if already released
	if escrow.Released {
		return sdkerrors.Wrap(types.ErrEscrowAlreadyReleased, orderId)
	}

	// Check timeout
	if ctx.BlockTime().After(escrow.Timeout) {
		return sdkerrors.Wrap(types.ErrEscrowTimeout, orderId)
	}

	// Transfer from module account to merchant
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddr == nil {
		return sdkerrors.Wrap(types.ErrModuleAccountNotFound, types.ModuleName)
	}

	err := k.bankKeeper.SendCoins(ctx, moduleAddr, to, sdk.NewCoins(escrow.Amount))
	if err != nil {
		return sdkerrors.Wrap(types.ErrEscrowReleaseFailed, err.Error())
	}

	// Mark as released
	escrow.Released = true
	escrow.MerchantAddr = to.String()
	k.setEscrow(ctx, orderId, escrow)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"escrow_released",
			sdk.NewAttribute("order_id", orderId),
			sdk.NewAttribute("merchant", to.String()),
			sdk.NewAttribute("amount", escrow.Amount.String()),
		),
	)

	return nil
}

// RefundEscrow refunds escrowed stablecoins back to the customer
func (k Keeper) RefundEscrow(ctx sdk.Context, orderId string, to sdk.AccAddress) error {
	// Get escrow
	escrow, found := k.getEscrow(ctx, orderId)
	if !found {
		return sdkerrors.Wrap(types.ErrEscrowNotFound, orderId)
	}

	// Check if already released
	if escrow.Released {
		return sdkerrors.Wrap(types.ErrEscrowAlreadyReleased, orderId)
	}

	// Transfer from module account back to customer
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddr == nil {
		return sdkerrors.Wrap(types.ErrModuleAccountNotFound, types.ModuleName)
	}

	err := k.bankKeeper.SendCoins(ctx, moduleAddr, to, sdk.NewCoins(escrow.Amount))
	if err != nil {
		return sdkerrors.Wrap(types.ErrRefundFailed, err.Error())
	}

	// Mark as released (refunded)
	escrow.Released = true
	k.setEscrow(ctx, orderId, escrow)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"escrow_refunded",
			sdk.NewAttribute("order_id", orderId),
			sdk.NewAttribute("customer", to.String()),
			sdk.NewAttribute("amount", escrow.Amount.String()),
		),
	)

	return nil
}

// GetEscrowBalance returns the escrowed balance for an order
func (k Keeper) GetEscrowBalance(ctx sdk.Context, orderId string) (sdk.Coin, bool) {
	escrow, found := k.getEscrow(ctx, orderId)
	if !found {
		return sdk.Coin{}, false
	}
	
	if escrow.Released {
		return sdk.Coin{}, false
	}

	return escrow.Amount, true
}

// TransferStablecoin transfers stablecoins between accounts with validation
func (k Keeper) TransferStablecoin(ctx sdk.Context, from, to sdk.AccAddress, amount sdk.Coin) error {
	// Validate stablecoin
	if !k.IsValidStablecoin(ctx, amount.Denom) {
		return sdkerrors.Wrap(types.ErrInvalidStablecoin, amount.Denom)
	}

	// Check if from address is blacklisted
	if k.IsBlacklisted(ctx, amount.Denom, from.String()) {
		return sdkerrors.Wrap(types.ErrAddressBlacklisted, from.String())
	}

	// Check if to address is blacklisted
	if k.IsBlacklisted(ctx, amount.Denom, to.String()) {
		return sdkerrors.Wrap(types.ErrAddressBlacklisted, to.String())
	}

	// Get stablecoin to check requirements
	stablecoin, found := k.GetStablecoin(ctx, amount.Denom)
	if !found {
		return sdkerrors.Wrap(types.ErrStablecoinNotFound, amount.Denom)
	}

	// Check whitelist requirement
	if stablecoin.RequireWhitelist {
		if !k.IsWhitelisted(ctx, amount.Denom, from.String()) {
			return sdkerrors.Wrap(types.ErrAddressNotWhitelisted, from.String())
		}
		if !k.IsWhitelisted(ctx, amount.Denom, to.String()) {
			return sdkerrors.Wrap(types.ErrAddressNotWhitelisted, to.String())
		}
	}

	// Perform transfer
	err := k.bankKeeper.SendCoins(ctx, from, to, sdk.NewCoins(amount))
	if err != nil {
		return sdkerrors.Wrap(types.ErrTransferFailed, err.Error())
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_transferred",
			sdk.NewAttribute("from", from.String()),
			sdk.NewAttribute("to", to.String()),
			sdk.NewAttribute("amount", amount.String()),
		),
	)

	return nil
}

// IsValidStablecoin checks if a denomination is a valid stablecoin
func (k Keeper) IsValidStablecoin(ctx sdk.Context, denom string) bool {
	_, found := k.GetStablecoin(ctx, denom)
	return found
}

// ValidateStablecoinPayment validates a stablecoin payment
func (k Keeper) ValidateStablecoinPayment(ctx sdk.Context, denom string, amount sdk.Int) error {
	// Check if stablecoin exists
	stablecoin, found := k.GetStablecoin(ctx, denom)
	if !found {
		return sdkerrors.Wrap(types.ErrStablecoinNotFound, denom)
	}

	// Check if stablecoin is active
	if !stablecoin.Active {
		return sdkerrors.Wrap(types.ErrStablecoinInactive, denom)
	}

	// Check minimum amount if set
	if stablecoin.MinAmount != nil && amount.LT(*stablecoin.MinAmount) {
		return sdkerrors.Wrap(types.ErrAmountBelowMinimum, fmt.Sprintf("minimum: %s, provided: %s", stablecoin.MinAmount.String(), amount.String()))
	}

	// Check maximum amount if set
	if stablecoin.MaxAmount != nil && amount.GT(*stablecoin.MaxAmount) {
		return sdkerrors.Wrap(types.ErrAmountAboveMaximum, fmt.Sprintf("maximum: %s, provided: %s", stablecoin.MaxAmount.String(), amount.String()))
	}

	return nil
}

// GetAllStablecoins returns all stablecoins (alias for compatibility)
func (k Keeper) GetAllStablecoins(ctx sdk.Context) []types.Stablecoin {
	return k.GetAllStablecoin(ctx)
}

// Internal helper methods

// setEscrow stores escrow information
func (k Keeper) setEscrow(ctx sdk.Context, orderId string, escrow EscrowAccount) {
	store := ctx.KVStore(k.storeKey)
	key := types.EscrowKey(orderId)
	bz := k.cdc.MustMarshal(&escrow)
	store.Set(key, bz)
}

// getEscrow retrieves escrow information
func (k Keeper) getEscrow(ctx sdk.Context, orderId string) (EscrowAccount, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.EscrowKey(orderId)
	bz := store.Get(key)
	if bz == nil {
		return EscrowAccount{}, false
	}

	var escrow EscrowAccount
	k.cdc.MustUnmarshal(bz, &escrow)
	return escrow, true
}

// hasEscrow checks if escrow exists for an order
func (k Keeper) hasEscrow(ctx sdk.Context, orderId string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.EscrowKey(orderId)
	return store.Has(key)
}