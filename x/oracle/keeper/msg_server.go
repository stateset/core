package keeper

import (
	"context"
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/stateset/core/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// SubmitPriceFeed handles price feed submissions from oracle providers
func (k msgServer) SubmitPriceFeed(goCtx context.Context, msg *types.MsgSubmitPriceFeed) (*types.MsgSubmitPriceFeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Verify provider is registered and active
	provider, found := k.GetOracleProvider(ctx, msg.Provider)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "provider %s is not registered", msg.Provider)
	}
	if !provider.Active {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "provider %s is not active", msg.Provider)
	}
	
	// Validate price submission against deviation thresholds
	if err := k.ValidatePriceSubmission(ctx, msg.Asset, msg.Price, msg.Provider); err != nil {
		// Slash provider for bad submission
		params := k.GetParams(ctx)
		providerAddr, _ := sdk.AccAddressFromBech32(msg.Provider)
		slashCoins := sdk.NewCoins(sdk.NewCoin("state", params.SlashAmount))
		
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, providerAddr, types.ModuleName, slashCoins); err != nil {
			return nil, sdkerrors.Wrapf(err, "failed to slash provider")
		}
		
		return nil, sdkerrors.Wrapf(types.ErrInvalidPrice, "price validation failed: %s", err)
	}
	
	// Create price feed
	feed := types.PriceFeed{
		FeedID:     fmt.Sprintf("%s-%s-%d", msg.Asset, msg.Provider, ctx.BlockHeight()),
		Asset:      msg.Asset,
		Price:      msg.Price,
		Provider:   msg.Provider,
		Timestamp:  ctx.BlockTime(),
		Expiry:     msg.Expiry,
		Confidence: msg.Confidence,
		Volume:     msg.Volume,
	}
	
	// Store price feed
	if err := k.SetPriceFeed(ctx, feed); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to store price feed")
	}
	
	// Update provider's last update time
	provider.LastUpdate = ctx.BlockTime()
	if err := k.SetOracleProvider(ctx, provider); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to update provider")
	}
	
	// Reward provider for valid submission
	params := k.GetParams(ctx)
	providerAddr, _ := sdk.AccAddressFromBech32(msg.Provider)
	rewardCoins := sdk.NewCoins(sdk.NewCoin("state", params.RewardAmount))
	
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, providerAddr, rewardCoins); err != nil {
		// Log error but don't fail the transaction
		ctx.Logger().Error("failed to send rewards", "error", err)
	}
	
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"price_feed_submitted",
			sdk.NewAttribute("feed_id", feed.FeedID),
			sdk.NewAttribute("asset", msg.Asset),
			sdk.NewAttribute("price", msg.Price.String()),
			sdk.NewAttribute("provider", msg.Provider),
			sdk.NewAttribute("confidence", msg.Confidence.String()),
		),
	)
	
	return &types.MsgSubmitPriceFeedResponse{
		FeedId:  feed.FeedID,
		Success: true,
	}, nil
}

// RegisterOracle registers a new oracle provider
func (k msgServer) RegisterOracle(goCtx context.Context, msg *types.MsgRegisterOracle) (*types.MsgRegisterOracleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Verify authority
	if msg.Authority != k.GetAuthority() {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}
	
	// Check if provider already exists
	if _, found := k.GetOracleProvider(ctx, msg.ProviderAddr); found {
		return nil, sdkerrors.Wrapf(types.ErrOracleAlreadyExists, "oracle provider %s already exists", msg.ProviderAddr)
	}
	
	// Create oracle provider
	provider := types.OracleProvider{
		Name:           msg.Name,
		Address:        msg.ProviderAddr,
		Active:         true,
		Priority:       msg.Priority,
		MinSubmissions: msg.MinSubmissions,
		MaxDeviation:   msg.MaxDeviation,
		LastUpdate:     ctx.BlockTime(),
		Reputation:     sdk.NewDecWithPrec(5, 1), // Start with 0.5 reputation
	}
	
	// Store oracle provider
	if err := k.SetOracleProvider(ctx, provider); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to store oracle provider")
	}
	
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"oracle_registered",
			sdk.NewAttribute("name", msg.Name),
			sdk.NewAttribute("address", msg.ProviderAddr),
			sdk.NewAttribute("priority", fmt.Sprintf("%d", msg.Priority)),
		),
	)
	
	return &types.MsgRegisterOracleResponse{
		Success: true,
	}, nil
}

// UpdateOracle updates an existing oracle provider
func (k msgServer) UpdateOracle(goCtx context.Context, msg *types.MsgUpdateOracle) (*types.MsgUpdateOracleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Verify authority
	if msg.Authority != k.GetAuthority() {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}
	
	// Get existing provider
	provider, found := k.GetOracleProvider(ctx, msg.ProviderAddr)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOracleNotFound, "oracle provider %s not found", msg.ProviderAddr)
	}
	
	// Update provider fields
	provider.Active = msg.Active
	provider.Priority = msg.Priority
	provider.MinSubmissions = msg.MinSubmissions
	provider.MaxDeviation = msg.MaxDeviation
	
	// Store updated provider
	if err := k.SetOracleProvider(ctx, provider); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to update oracle provider")
	}
	
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"oracle_updated",
			sdk.NewAttribute("address", msg.ProviderAddr),
			sdk.NewAttribute("active", fmt.Sprintf("%t", msg.Active)),
			sdk.NewAttribute("priority", fmt.Sprintf("%d", msg.Priority)),
		),
	)
	
	return &types.MsgUpdateOracleResponse{
		Success: true,
	}, nil
}

// RemoveOracle removes an oracle provider
func (k msgServer) RemoveOracle(goCtx context.Context, msg *types.MsgRemoveOracle) (*types.MsgRemoveOracleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Verify authority
	if msg.Authority != k.GetAuthority() {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}
	
	// Check if provider exists
	if _, found := k.GetOracleProvider(ctx, msg.ProviderAddr); !found {
		return nil, sdkerrors.Wrapf(types.ErrOracleNotFound, "oracle provider %s not found", msg.ProviderAddr)
	}
	
	// Remove oracle provider
	k.RemoveOracleProvider(ctx, msg.ProviderAddr)
	
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"oracle_removed",
			sdk.NewAttribute("address", msg.ProviderAddr),
			sdk.NewAttribute("reason", msg.Reason),
		),
	)
	
	return &types.MsgRemoveOracleResponse{
		Success: true,
	}, nil
}

// RequestPrice handles price update requests
func (k msgServer) RequestPrice(goCtx context.Context, msg *types.MsgRequestPrice) (*types.MsgRequestPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Check if aggregated price exists
	aggregated, found := k.GetAggregatedPrice(ctx, msg.Asset)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPriceNotAvailable, "no price available for asset %s", msg.Asset)
	}
	
	// Check if price needs update
	needsUpdate := ctx.BlockTime().After(aggregated.NextUpdate) || msg.Urgent
	
	if needsUpdate {
		// Emit event to trigger oracle providers to submit new prices
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"price_update_requested",
				sdk.NewAttribute("asset", msg.Asset),
				sdk.NewAttribute("requester", msg.Requester),
				sdk.NewAttribute("urgent", fmt.Sprintf("%t", msg.Urgent)),
				sdk.NewAttribute("current_price", aggregated.Price.String()),
			),
		)
		
		// In a real implementation, this would trigger off-chain oracle providers
		// to fetch and submit new prices
	}
	
	return &types.MsgRequestPriceResponse{
		Asset:       msg.Asset,
		Price:       aggregated.Price,
		LastUpdate:  aggregated.LastUpdate,
		NextUpdate:  aggregated.NextUpdate,
		Confidence:  aggregated.Confidence,
		NumProviders: aggregated.NumProviders,
	}, nil
}

// GetParams returns the oracle module parameters
func (k msgServer) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}