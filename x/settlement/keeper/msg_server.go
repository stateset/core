package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/settlement/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the settlement MsgServer interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// InstantTransfer handles instant stablecoin transfers
func (m msgServer) InstantTransfer(goCtx context.Context, msg *types.MsgInstantTransfer) (*types.MsgInstantTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	settlementId, err := m.Keeper.InstantTransfer(ctx, msg.Sender, msg.Recipient, msg.Amount, msg.Reference, msg.Metadata)
	if err != nil {
		return nil, err
	}

	return &types.MsgInstantTransferResponse{
		SettlementId: settlementId,
		TxHash:       "", // Will be filled by client
	}, nil
}

// CreateEscrow creates an escrow settlement
func (m msgServer) CreateEscrow(goCtx context.Context, msg *types.MsgCreateEscrow) (*types.MsgCreateEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Convert duration to seconds
	expirationSeconds := int64(msg.ExpiresIn.Seconds())

	settlementId, err := m.Keeper.CreateEscrow(ctx, msg.Sender, msg.Recipient, msg.Amount, msg.Reference, msg.Metadata, expirationSeconds)
	if err != nil {
		return nil, err
	}

	// Get the settlement to return the calculated expiration
	settlement, _ := m.Keeper.GetSettlement(ctx, settlementId)

	return &types.MsgCreateEscrowResponse{
		SettlementId: settlementId,
		ExpiresAt:    settlement.ExpiresAt,
	}, nil
}

// ReleaseEscrow releases escrowed funds to recipient
func (m msgServer) ReleaseEscrow(goCtx context.Context, msg *types.MsgReleaseEscrow) (*types.MsgReleaseEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, types.ErrInvalidSettlement
	}

	if err := m.Keeper.ReleaseEscrow(ctx, msg.SettlementId, senderAddr); err != nil {
		return nil, err
	}

	return &types.MsgReleaseEscrowResponse{}, nil
}

// RefundEscrow refunds escrowed funds to sender
func (m msgServer) RefundEscrow(goCtx context.Context, msg *types.MsgRefundEscrow) (*types.MsgRefundEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	recipientAddr, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, types.ErrInvalidSettlement
	}

	if err := m.Keeper.RefundEscrow(ctx, msg.SettlementId, recipientAddr, msg.Reason); err != nil {
		return nil, err
	}

	return &types.MsgRefundEscrowResponse{}, nil
}

// CreateBatch creates a batch settlement
func (m msgServer) CreateBatch(goCtx context.Context, msg *types.MsgCreateBatch) (*types.MsgCreateBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Only module authority can create batches since this message can move funds for many senders.
	if msg.Authority != m.Keeper.GetAuthority() {
		return nil, types.ErrUnauthorized
	}

	batchId, settlementIds, err := m.Keeper.CreateBatch(ctx, msg.Merchant, msg.Senders, msg.Amounts, msg.References)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateBatchResponse{
		BatchId:       batchId,
		SettlementIds: settlementIds,
	}, nil
}

// SettleBatch settles a batch of payments
func (m msgServer) SettleBatch(goCtx context.Context, msg *types.MsgSettleBatch) (*types.MsgSettleBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.Keeper.SettleBatch(ctx, msg.BatchId, msg.Authority); err != nil {
		return nil, err
	}

	batch, _ := m.Keeper.GetBatch(ctx, msg.BatchId)

	return &types.MsgSettleBatchResponse{
		TotalAmount: batch.TotalAmount,
		TotalFees:   batch.TotalFees,
		NetAmount:   batch.NetAmount,
	}, nil
}

// OpenChannel opens a payment channel
func (m msgServer) OpenChannel(goCtx context.Context, msg *types.MsgOpenChannel) (*types.MsgOpenChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	channelId, err := m.Keeper.OpenChannel(ctx, msg.Sender, msg.Recipient, msg.Deposit, msg.ExpiresInBlocks)
	if err != nil {
		return nil, err
	}

	return &types.MsgOpenChannelResponse{
		ChannelId:       channelId,
		ExpiresAtHeight: ctx.BlockHeight() + msg.ExpiresInBlocks,
	}, nil
}

// CloseChannel closes a payment channel
func (m msgServer) CloseChannel(goCtx context.Context, msg *types.MsgCloseChannel) (*types.MsgCloseChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	closerAddr, err := sdk.AccAddressFromBech32(msg.Closer)
	if err != nil {
		return nil, types.ErrInvalidSettlement
	}

	finalBalance, err := m.Keeper.CloseChannel(ctx, msg.ChannelId, closerAddr)
	if err != nil {
		return nil, err
	}

	return &types.MsgCloseChannelResponse{
		FinalBalance: finalBalance,
	}, nil
}

// ClaimChannel claims funds from a payment channel
func (m msgServer) ClaimChannel(goCtx context.Context, msg *types.MsgClaimChannel) (*types.MsgClaimChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	recipientAddr, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, types.ErrInvalidSettlement
	}

	if err := m.Keeper.ClaimChannel(ctx, msg.ChannelId, recipientAddr, msg.Amount, msg.Nonce, msg.Signature); err != nil {
		return nil, err
	}

	channel, _ := m.Keeper.GetChannel(ctx, msg.ChannelId)

	return &types.MsgClaimChannelResponse{
		AmountClaimed:    msg.Amount,
		RemainingBalance: channel.Balance,
	}, nil
}

// RegisterMerchant registers a new merchant
func (m msgServer) RegisterMerchant(goCtx context.Context, msg *types.MsgRegisterMerchant) (*types.MsgRegisterMerchantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Only the merchant itself or module authority may register a merchant configuration.
	if msg.Authority != msg.Merchant && msg.Authority != m.Keeper.GetAuthority() {
		return nil, types.ErrUnauthorized
	}

	config := types.MerchantConfig{
		Address:        msg.Merchant,
		Name:           msg.Name,
		FeeRateBps:     msg.FeeRateBps,
		MinSettlement:  msg.MinSettlement,
		MaxSettlement:  msg.MaxSettlement,
		BatchEnabled:   msg.BatchEnabled,
		BatchThreshold: msg.BatchThreshold,
		WebhookUrl:     msg.WebhookUrl,
		IsActive:       true,
		RegisteredAt:   time.Time{}, // Will be set in keeper
	}

	if err := m.Keeper.RegisterMerchant(ctx, config); err != nil {
		return nil, err
	}

	return &types.MsgRegisterMerchantResponse{}, nil
}

// UpdateMerchant updates a merchant configuration
func (m msgServer) UpdateMerchant(goCtx context.Context, msg *types.MsgUpdateMerchant) (*types.MsgUpdateMerchantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Only the merchant itself or module authority may update merchant configuration.
	if msg.Authority != msg.Merchant && msg.Authority != m.Keeper.GetAuthority() {
		return nil, types.ErrUnauthorized
	}

	updates := make(map[string]interface{})

	if msg.Name != "" {
		updates["name"] = msg.Name
	}
	if msg.FeeRateBps > 0 {
		updates["fee_rate_bps"] = msg.FeeRateBps
	}
	if msg.BatchEnabled != nil {
		updates["batch_enabled"] = *msg.BatchEnabled
	}
	if msg.IsActive != nil {
		updates["is_active"] = *msg.IsActive
	}

	if err := m.Keeper.UpdateMerchant(ctx, msg.Merchant, updates); err != nil {
		return nil, err
	}

	return &types.MsgUpdateMerchantResponse{}, nil
}

// InstantCheckout handles streamlined checkout for ecommerce
func (m msgServer) InstantCheckout(goCtx context.Context, msg *types.MsgInstantCheckout) (*types.MsgInstantCheckoutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	settlementId, netAmount, fee, err := m.Keeper.InstantCheckout(
		ctx,
		msg.Customer,
		msg.Merchant,
		msg.Amount,
		msg.OrderReference,
		msg.UseEscrow,
		msg.Metadata,
	)
	if err != nil {
		return nil, err
	}

	status := "completed"
	if msg.UseEscrow {
		status = "escrowed"
	}

	return &types.MsgInstantCheckoutResponse{
		SettlementId: settlementId,
		Status:       status,
		NetAmount:    netAmount,
		Fee:          fee,
	}, nil
}

// PartialRefund handles partial refunds for settlements
func (m msgServer) PartialRefund(goCtx context.Context, msg *types.MsgPartialRefund) (*types.MsgPartialRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	remainingAmount, err := m.Keeper.PartialRefund(
		ctx,
		msg.Authority,
		msg.SettlementId,
		msg.RefundAmount,
		msg.Reason,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgPartialRefundResponse{
		RefundedAmount:  msg.RefundAmount,
		RemainingAmount: remainingAmount,
	}, nil
}
