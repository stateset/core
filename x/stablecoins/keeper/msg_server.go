package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/stablecoins/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateStablecoin(goCtx context.Context, msg *types.MsgCreateStablecoin) (*types.MsgCreateStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.CreateStablecoin(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateStablecoinResponse{
		Denom: msg.Denom,
	}, nil
}

func (k msgServer) UpdateStablecoin(goCtx context.Context, msg *types.MsgUpdateStablecoin) (*types.MsgUpdateStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the existing stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can update
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can update stablecoin")
	}

	// Update fields
	if msg.Name != "" {
		stablecoin.Name = msg.Name
	}
	if msg.Description != "" {
		stablecoin.Description = msg.Description
	}
	if msg.PegInfo != nil {
		stablecoin.PegInfo = msg.PegInfo
	}
	if msg.FeeInfo != nil {
		stablecoin.FeeInfo = msg.FeeInfo
	}
	if msg.AccessControl != nil {
		stablecoin.AccessControl = msg.AccessControl
	}
	if msg.Metadata != "" {
		stablecoin.Metadata = msg.Metadata
	}

	// Save updated stablecoin
	k.Keeper.SetStablecoin(ctx, stablecoin)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_updated",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgUpdateStablecoinResponse{}, nil
}

func (k msgServer) MintStablecoin(goCtx context.Context, msg *types.MsgMintStablecoin) (*types.MsgMintStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.MintStablecoin(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintStablecoinResponse{}, nil
}

func (k msgServer) BurnStablecoin(goCtx context.Context, msg *types.MsgBurnStablecoin) (*types.MsgBurnStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.BurnStablecoin(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnStablecoinResponse{}, nil
}

func (k msgServer) PauseStablecoin(goCtx context.Context, msg *types.MsgPauseStablecoin) (*types.MsgPauseStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can pause
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can pause stablecoin")
	}

	// Update pause status based on operation
	switch msg.Operation {
	case "mint":
		stablecoin.MintPaused = true
	case "burn":
		stablecoin.BurnPaused = true
	case "transfer":
		stablecoin.TransferPaused = true
	case "all":
		stablecoin.MintPaused = true
		stablecoin.BurnPaused = true
		stablecoin.TransferPaused = true
	default:
		return nil, sdkerrors.Wrap(types.ErrOperationNotAllowed, "invalid operation type")
	}

	// Save updated stablecoin
	k.Keeper.SetStablecoin(ctx, stablecoin)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_paused",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("operation", msg.Operation),
			sdk.NewAttribute("reason", msg.Reason),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgPauseStablecoinResponse{}, nil
}

func (k msgServer) UnpauseStablecoin(goCtx context.Context, msg *types.MsgUnpauseStablecoin) (*types.MsgUnpauseStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can unpause
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can unpause stablecoin")
	}

	// Update pause status based on operation
	switch msg.Operation {
	case "mint":
		stablecoin.MintPaused = false
	case "burn":
		stablecoin.BurnPaused = false
	case "transfer":
		stablecoin.TransferPaused = false
	case "all":
		stablecoin.MintPaused = false
		stablecoin.BurnPaused = false
		stablecoin.TransferPaused = false
	default:
		return nil, sdkerrors.Wrap(types.ErrOperationNotAllowed, "invalid operation type")
	}

	// Save updated stablecoin
	k.Keeper.SetStablecoin(ctx, stablecoin)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_unpaused",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("operation", msg.Operation),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgUnpauseStablecoinResponse{}, nil
}

func (k msgServer) UpdatePriceData(goCtx context.Context, msg *types.MsgUpdatePriceData) (*types.MsgUpdatePriceDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - admin or authorized oracle
	if msg.Creator != stablecoin.Admin {
		// In a real implementation, we would check if the creator is an authorized oracle
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin or authorized oracle can update price data")
	}

	// Create price data entry
	priceData := types.PriceData{
		Denom:       msg.Denom,
		Price:       msg.Price,
		Source:      msg.Source,
		Timestamp:   ctx.BlockTime(),
		BlockHeight: uint64(ctx.BlockHeight()),
	}

	// Store price data
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&priceData)
	store.Set(append(types.PriceDataKeyPrefix, types.PriceDataKey(msg.Denom)...), b)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"price_data_updated",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("price", msg.Price.String()),
			sdk.NewAttribute("source", msg.Source),
			sdk.NewAttribute("updater", msg.Creator),
		),
	)

	return &types.MsgUpdatePriceDataResponse{}, nil
}

func (k msgServer) UpdateReserves(goCtx context.Context, msg *types.MsgUpdateReserves) (*types.MsgUpdateReservesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can update reserves
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can update reserves")
	}

	// Update reserve information
	if stablecoin.ReserveInfo == nil {
		stablecoin.ReserveInfo = &types.ReserveInfo{}
	}
	stablecoin.ReserveInfo.ReserveAssets = msg.ReserveAssets

	// Save updated stablecoin
	k.Keeper.SetStablecoin(ctx, stablecoin)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"reserves_updated",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgUpdateReservesResponse{}, nil
}

func (k msgServer) WhitelistAddress(goCtx context.Context, msg *types.MsgWhitelistAddress) (*types.MsgWhitelistAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can manage whitelist
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can manage whitelist")
	}

	// Add to whitelist
	k.Keeper.WhitelistAddress(ctx, msg.Denom, msg.Address)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"address_whitelisted",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgWhitelistAddressResponse{}, nil
}

func (k msgServer) BlacklistAddress(goCtx context.Context, msg *types.MsgBlacklistAddress) (*types.MsgBlacklistAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can manage blacklist
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can manage blacklist")
	}

	// Add to blacklist
	k.Keeper.BlacklistAddress(ctx, msg.Denom, msg.Address, msg.Reason)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"address_blacklisted",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("reason", msg.Reason),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgBlacklistAddressResponse{}, nil
}

func (k msgServer) RemoveFromWhitelist(goCtx context.Context, msg *types.MsgRemoveFromWhitelist) (*types.MsgRemoveFromWhitelistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can manage whitelist
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can manage whitelist")
	}

	// Remove from whitelist
	store := ctx.KVStore(k.storeKey)
	key := append(types.WhitelistKeyPrefix, types.WhitelistKey(msg.Denom, msg.Address)...)
	store.Delete(key)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"address_removed_from_whitelist",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgRemoveFromWhitelistResponse{}, nil
}

func (k msgServer) RemoveFromBlacklist(goCtx context.Context, msg *types.MsgRemoveFromBlacklist) (*types.MsgRemoveFromBlacklistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the stablecoin
	stablecoin, found := k.Keeper.GetStablecoin(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, msg.Denom)
	}

	// Check authorization - only admin can manage blacklist
	if msg.Creator != stablecoin.Admin {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only admin can manage blacklist")
	}

	// Remove from blacklist
	store := ctx.KVStore(k.storeKey)
	key := append(types.BlacklistKeyPrefix, types.BlacklistKey(msg.Denom, msg.Address)...)
	store.Delete(key)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"address_removed_from_blacklist",
			sdk.NewAttribute("denom", msg.Denom),
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("admin", msg.Creator),
		),
	)

	return &types.MsgRemoveFromBlacklistResponse{}, nil
}

func (k msgServer) InitializeSSUSD(goCtx context.Context, msg *types.MsgInitializeSSUSD) (*types.MsgInitializeSSUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create ssUSD stablecoin engine
	engine := NewSSUSDStablecoinEngine(&k.Keeper)

	// Initialize ssUSD
	err := engine.InitializeSSUSD(ctx)
	if err != nil {
		return nil, err
	}

	return &types.MsgInitializeSSUSDResponse{
		Success: true,
	}, nil
}

func (k msgServer) IssueSSUSD(goCtx context.Context, msg *types.MsgIssueSSUSD) (*types.MsgIssueSSUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create ssUSD stablecoin engine
	engine := NewSSUSDStablecoinEngine(&k.Keeper)

	// Create issue request
	request := SSUSDIssueRequest{
		Requester:      msg.Creator,
		Amount:         msg.Amount,
		ReservePayment: msg.ReservePayment,
		RequestTime:    ctx.BlockTime(),
	}

	// Issue ssUSD
	err := engine.IssueSSUSD(ctx, request)
	if err != nil {
		return nil, err
	}

	return &types.MsgIssueSSUSDResponse{
		AmountIssued: msg.Amount,
	}, nil
}

func (k msgServer) RedeemSSUSD(goCtx context.Context, msg *types.MsgRedeemSSUSD) (*types.MsgRedeemSSUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create ssUSD stablecoin engine
	engine := NewSSUSDStablecoinEngine(&k.Keeper)

	// Create redeem request
	request := SSUSDRedeemRequest{
		Requester:      msg.Creator,
		SSUSDAmount:    msg.SSUSDAmount,
		PreferredAsset: msg.PreferredAsset,
		RequestTime:    ctx.BlockTime(),
	}

	// Redeem ssUSD
	err := engine.RedeemSSUSD(ctx, request)
	if err != nil {
		return nil, err
	}

	return &types.MsgRedeemSSUSDResponse{
		AmountRedeemed: msg.SSUSDAmount,
	}, nil
}