package ibc

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

// IBCHooks provides IBC middleware hooks for cross-chain settlement
type IBCHooks struct {
	settlementKeeper SettlementKeeper
}

// SettlementKeeper defines the expected settlement keeper interface
type SettlementKeeper interface {
	ProcessCrossChainSettlement(ctx sdk.Context, sender, receiver string, amount sdk.Coins, sourceChannel, destChannel string) error
	GetCrossChainEscrow(ctx sdk.Context, escrowID string) (CrossChainEscrow, bool)
	CompleteCrossChainEscrow(ctx sdk.Context, escrowID string, packet channeltypes.Packet) error
	RefundCrossChainEscrow(ctx sdk.Context, escrowID string) error
}

// CrossChainEscrow represents a cross-chain escrow
type CrossChainEscrow struct {
	ID            string      `json:"id"`
	Sender        string      `json:"sender"`
	Receiver      string      `json:"receiver"`
	Amount        sdk.Coins   `json:"amount"`
	SourceChain   string      `json:"source_chain"`
	DestChain     string      `json:"dest_chain"`
	SourceChannel string      `json:"source_channel"`
	DestChannel   string      `json:"dest_channel"`
	Status        string      `json:"status"`
	CreatedAt     int64       `json:"created_at"`
	ExpiresAt     int64       `json:"expires_at"`
}

// CrossChainSettlementMemo represents the memo structure for cross-chain settlements
type CrossChainSettlementMemo struct {
	Type       string `json:"type"`              // "settlement", "escrow_create", "escrow_release"
	EscrowID   string `json:"escrow_id,omitempty"`
	Receiver   string `json:"receiver,omitempty"`
	Conditions string `json:"conditions,omitempty"`
}

// NewIBCHooks creates a new IBC hooks instance
func NewIBCHooks(sk SettlementKeeper) *IBCHooks {
	return &IBCHooks{
		settlementKeeper: sk,
	}
}

// OnRecvPacket processes incoming IBC transfer packets
func (h *IBCHooks) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// Parse the transfer packet data
	var data transfertypes.FungibleTokenPacketData
	if err := json.Unmarshal(packet.GetData(), &data); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// Check if this is a settlement-related transfer
	if data.Memo == "" {
		return nil // Not a settlement transfer, continue with normal processing
	}

	var memo CrossChainSettlementMemo
	if err := json.Unmarshal([]byte(data.Memo), &memo); err != nil {
		return nil // Invalid memo format, continue with normal processing
	}

	// Handle different settlement types
	switch memo.Type {
	case "settlement":
		return h.handleSettlement(ctx, packet, data, memo)
	case "escrow_release":
		return h.handleEscrowRelease(ctx, packet, data, memo)
	default:
		return nil // Unknown type, continue with normal processing
	}
}

// handleSettlement processes a cross-chain settlement
func (h *IBCHooks) handleSettlement(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data transfertypes.FungibleTokenPacketData,
	memo CrossChainSettlementMemo,
) ibcexported.Acknowledgement {
	// Parse amount
	amount, ok := sdkmath.NewIntFromString(data.Amount)
	if !ok {
		return channeltypes.NewErrorAcknowledgement(errorsmod.Wrap(ErrInvalidAmount, "failed to parse amount"))
	}

	coins := sdk.NewCoins(sdk.NewCoin(data.Denom, amount))

	// Process the cross-chain settlement
	if err := h.settlementKeeper.ProcessCrossChainSettlement(
		ctx,
		data.Sender,
		data.Receiver,
		coins,
		packet.SourceChannel,
		packet.DestinationChannel,
	); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"cross_chain_settlement",
			sdk.NewAttribute("sender", data.Sender),
			sdk.NewAttribute("receiver", data.Receiver),
			sdk.NewAttribute("amount", data.Amount),
			sdk.NewAttribute("denom", data.Denom),
			sdk.NewAttribute("source_channel", packet.SourceChannel),
			sdk.NewAttribute("dest_channel", packet.DestinationChannel),
		),
	)

	return nil // Success, continue with normal transfer processing
}

// handleEscrowRelease processes a cross-chain escrow release
func (h *IBCHooks) handleEscrowRelease(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data transfertypes.FungibleTokenPacketData,
	memo CrossChainSettlementMemo,
) ibcexported.Acknowledgement {
	// Get the escrow
	escrow, found := h.settlementKeeper.GetCrossChainEscrow(ctx, memo.EscrowID)
	if !found {
		return channeltypes.NewErrorAcknowledgement(errorsmod.Wrap(ErrEscrowNotFound, memo.EscrowID))
	}

	// Verify the release conditions
	if escrow.Status != "active" {
		return channeltypes.NewErrorAcknowledgement(errorsmod.Wrap(ErrInvalidEscrowStatus, "escrow is not active"))
	}

	// Complete the escrow
	if err := h.settlementKeeper.CompleteCrossChainEscrow(ctx, memo.EscrowID, packet); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"cross_chain_escrow_released",
			sdk.NewAttribute("escrow_id", memo.EscrowID),
			sdk.NewAttribute("sender", data.Sender),
			sdk.NewAttribute("receiver", data.Receiver),
		),
	)

	return nil
}

// OnAcknowledgementPacket processes acknowledgements for sent packets
func (h *IBCHooks) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	// Parse the transfer packet data
	var data transfertypes.FungibleTokenPacketData
	if err := json.Unmarshal(packet.GetData(), &data); err != nil {
		return nil // Not our concern
	}

	// Check if this was a settlement-related transfer
	if data.Memo == "" {
		return nil
	}

	var memo CrossChainSettlementMemo
	if err := json.Unmarshal([]byte(data.Memo), &memo); err != nil {
		return nil
	}

	// Parse acknowledgement
	var ack channeltypes.Acknowledgement
	if err := json.Unmarshal(acknowledgement, &ack); err != nil {
		return nil
	}

	// Handle failed transfers
	if !ack.Success() {
		if memo.EscrowID != "" {
			// Refund the escrow
			if err := h.settlementKeeper.RefundCrossChainEscrow(ctx, memo.EscrowID); err != nil {
				return err
			}

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					"cross_chain_escrow_refunded",
					sdk.NewAttribute("escrow_id", memo.EscrowID),
					sdk.NewAttribute("reason", "acknowledgement_failed"),
				),
			)
		}
	}

	return nil
}

// OnTimeoutPacket processes timeout for sent packets
func (h *IBCHooks) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	// Parse the transfer packet data
	var data transfertypes.FungibleTokenPacketData
	if err := json.Unmarshal(packet.GetData(), &data); err != nil {
		return nil
	}

	// Check if this was a settlement-related transfer
	if data.Memo == "" {
		return nil
	}

	var memo CrossChainSettlementMemo
	if err := json.Unmarshal([]byte(data.Memo), &memo); err != nil {
		return nil
	}

	// Handle timeout for escrow transfers
	if memo.EscrowID != "" {
		if err := h.settlementKeeper.RefundCrossChainEscrow(ctx, memo.EscrowID); err != nil {
			return err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"cross_chain_escrow_timeout",
				sdk.NewAttribute("escrow_id", memo.EscrowID),
			),
		)
	}

	return nil
}

// CreateSettlementMemo creates a memo for cross-chain settlement
func CreateSettlementMemo(receiver string) string {
	memo := CrossChainSettlementMemo{
		Type:     "settlement",
		Receiver: receiver,
	}
	bz, _ := json.Marshal(memo)
	return string(bz)
}

// CreateEscrowReleaseMemo creates a memo for escrow release
func CreateEscrowReleaseMemo(escrowID string) string {
	memo := CrossChainSettlementMemo{
		Type:     "escrow_release",
		EscrowID: escrowID,
	}
	bz, _ := json.Marshal(memo)
	return string(bz)
}
