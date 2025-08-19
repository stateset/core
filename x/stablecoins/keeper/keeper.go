package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	
	"github.com/stateset/core/x/stablecoins/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramstore.GetParamSet(ctx, &p)
	return p
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// ----- Stablecoin CRUD Operations -----

// SetStablecoin set a specific stablecoin in the store from its index
func (k Keeper) SetStablecoin(ctx sdk.Context, stablecoin types.Stablecoin) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&stablecoin)
	store.Set(types.StablecoinKey(stablecoin.Denom), b)
}

// GetStablecoin returns a stablecoin from its index
func (k Keeper) GetStablecoin(ctx sdk.Context, denom string) (val types.Stablecoin, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.StablecoinKey(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStablecoin removes a stablecoin from the store
func (k Keeper) RemoveStablecoin(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.StablecoinKey(denom))
}

// GetAllStablecoin returns all stablecoin
func (k Keeper) GetAllStablecoin(ctx sdk.Context) (list []types.Stablecoin) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.StablecoinKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Stablecoin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// ----- Business Logic -----

// CreateStablecoin creates a new stablecoin
func (k Keeper) CreateStablecoin(ctx sdk.Context, msg *types.MsgCreateStablecoin) error {
	// Check if stablecoin already exists
	_, found := k.GetStablecoin(ctx, msg.Denom)
	if found {
		return sdkerrors.Wrap(types.ErrStablecoinAlreadyExists, msg.Denom)
	}

	// Validate parameters
	params := k.GetParams(ctx)
	
	// Check max stablecoins limit
	allStablecoins := k.GetAllStablecoin(ctx)
	if uint64(len(allStablecoins)) >= params.MaxStablecoins {
		return errorsmod.Wrapf(types.ErrOperationNotAllowed, "maximum number of stablecoins reached: %d", params.MaxStablecoins)
	}

	// Validate max supply against parameters
	maxSupply, ok := sdk.NewIntFromString(params.MaxInitialSupply)
	if !ok {
		return sdkerrors.Wrap(types.ErrInvalidAmount, "invalid max initial supply parameter")
	}
	if msg.MaxSupply.GT(maxSupply) {
		return sdkerrors.Wrap(types.ErrExceedsMaxSupply, "max supply exceeds maximum allowed")
	}

	// Create the stablecoin
	stablecoin := types.NewStablecoin(
		msg.Denom,
		msg.Name,
		msg.Symbol,
		msg.Decimals,
		msg.Description,
		msg.Creator, // issuer
		msg.Creator, // admin
		msg.MaxSupply,
		msg.PegInfo,
		msg.ReserveInfo,
		msg.StabilityMechanism,
		msg.FeeInfo,
		msg.AccessControl,
		msg.Metadata,
	)

	// Set denom metadata for bank module
	k.setDenomMetadata(ctx, stablecoin)

	// Save stablecoin
	k.SetStablecoin(ctx, stablecoin)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_created",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("name", msg.Name),
			sdk.NewAttribute("symbol", msg.Symbol),
			sdk.NewAttribute("issuer", msg.Creator),
		),
	)

	return nil
}

// MintStablecoin mints new stablecoin tokens
func (k Keeper) MintStablecoin(ctx sdk.Context, msg *types.MsgMintStablecoin) error {
	// Get stablecoin
	stablecoin, found := k.GetStablecoin(ctx, msg.Denom)
	if !found {
		return sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check if stablecoin is active
	if !stablecoin.Active {
		return sdkerrors.Wrap(types.ErrStablecoinInactive, msg.Denom)
	}

	// Check if minting is paused
	if stablecoin.MintPaused {
		return sdkerrors.Wrap(types.ErrStablecoinPaused, "minting is paused")
	}

	// Check authorization - only issuer can mint
	if msg.Creator != stablecoin.Issuer && msg.Creator != stablecoin.Admin {
		return sdkerrors.Wrap(types.ErrUnauthorized, "only issuer or admin can mint")
	}

	// Check if recipient is whitelisted (if whitelist is enabled)
	if stablecoin.AccessControl != nil && stablecoin.AccessControl.WhitelistEnabled {
		if !k.IsWhitelisted(ctx, msg.Denom, msg.Recipient) {
			return sdkerrors.Wrap(types.ErrAddressNotWhitelisted, msg.Recipient)
		}
	}

	// Check if recipient is blacklisted
	if stablecoin.AccessControl != nil && stablecoin.AccessControl.BlacklistEnabled {
		if k.IsBlacklisted(ctx, msg.Denom, msg.Recipient) {
			return sdkerrors.Wrap(types.ErrAddressBlacklisted, msg.Recipient)
		}
	}

	// Check max supply
	newTotalSupply := stablecoin.TotalSupply.Add(msg.Amount)
	if newTotalSupply.GT(stablecoin.MaxSupply) {
		return sdkerrors.Wrap(types.ErrExceedsMaxSupply, "minting would exceed max supply")
	}

	// Mint coins
	mintCoins := sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.Amount))
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if err != nil {
		return err
	}

	// Send to recipient
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddr, mintCoins)
	if err != nil {
		return err
	}

	// Update stablecoin total supply
	stablecoin.TotalSupply = newTotalSupply
	k.SetStablecoin(ctx, stablecoin)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_minted",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("amount", msg.Amount.String()),
			sdk.NewAttribute("recipient", msg.Recipient),
			sdk.NewAttribute("new_total_supply", newTotalSupply.String()),
		),
	)

	return nil
}

// BurnStablecoin burns stablecoin tokens
func (k Keeper) BurnStablecoin(ctx sdk.Context, msg *types.MsgBurnStablecoin) error {
	// Get stablecoin
	stablecoin, found := k.GetStablecoin(ctx, msg.Denom)
	if !found {
		return sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check if stablecoin is active
	if !stablecoin.Active {
		return sdkerrors.Wrap(types.ErrStablecoinInactive, msg.Denom)
	}

	// Check if burning is paused
	if stablecoin.BurnPaused {
		return sdkerrors.Wrap(types.ErrStablecoinPaused, "burning is paused")
	}

	// Get creator's balance
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}

	balance := k.bankKeeper.GetBalance(ctx, creatorAddr, msg.Denom)
	if balance.Amount.LT(msg.Amount) {
		return sdkerrors.Wrap(types.ErrInsufficientFunds, "insufficient balance to burn")
	}

	// Send coins to module account
	burnCoins := sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.Amount))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, burnCoins)
	if err != nil {
		return err
	}

	// Burn coins
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		return err
	}

	// Update stablecoin total supply
	stablecoin.TotalSupply = stablecoin.TotalSupply.Sub(msg.Amount)
	k.SetStablecoin(ctx, stablecoin)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_burned",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("amount", msg.Amount.String()),
			sdk.NewAttribute("burner", msg.Creator),
			sdk.NewAttribute("new_total_supply", stablecoin.TotalSupply.String()),
		),
	)

	return nil
}

// ----- Access Control -----

// CheckIsWhitelisted checks if an address is whitelisted for a stablecoin
func (k Keeper) CheckIsWhitelisted(ctx sdk.Context, denom, address string) bool {
	store := ctx.KVStore(k.storeKey)
	key := append(types.WhitelistKeyPrefix, types.WhitelistKey(denom, address)...)
	return store.Has(key)
}

// CheckIsBlacklisted checks if an address is blacklisted for a stablecoin
func (k Keeper) CheckIsBlacklisted(ctx sdk.Context, denom, address string) bool {
	store := ctx.KVStore(k.storeKey)
	key := append(types.BlacklistKeyPrefix, types.BlacklistKey(denom, address)...)
	return store.Has(key)
}

// WhitelistAddress adds an address to the whitelist
func (k Keeper) WhitelistAddress(ctx sdk.Context, denom, address string) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.WhitelistKeyPrefix, types.WhitelistKey(denom, address)...)
	store.Set(key, []byte{1})
}

// BlacklistAddress adds an address to the blacklist
func (k Keeper) BlacklistAddress(ctx sdk.Context, denom, address, reason string) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.BlacklistKeyPrefix, types.BlacklistKey(denom, address)...)
	store.Set(key, []byte(reason))
}

// ----- Helper functions -----

// setDenomMetadata sets the denomination metadata for the bank module
func (k Keeper) setDenomMetadata(ctx sdk.Context, stablecoin types.Stablecoin) {
	// This would set the denom metadata for display purposes
	// Implementation depends on the specific requirements
}

// GetStablecoinCount returns the total number of stablecoins
func (k Keeper) GetStablecoinCount(ctx sdk.Context) uint64 {
	return uint64(len(k.GetAllStablecoin(ctx)))
}

// GetStoreKey returns the store key
func (k Keeper) GetStoreKey() storetypes.StoreKey {
	return k.storeKey
}

// GetCodec returns the codec
func (k Keeper) GetCodec() codec.BinaryCodec {
	return k.cdc
}