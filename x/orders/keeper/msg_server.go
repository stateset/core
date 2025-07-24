package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/orders/types"
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

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Generate order ID
	orderID := k.Keeper.GetNextOrderID(ctx)

	// Create order
	order := types.Order{
		Id:              fmt.Sprintf("ORDER-%d", orderID),
		Customer:        msg.Customer,
		Merchant:        msg.Merchant,
		Status:          "pending",
		TotalAmount:     msg.TotalAmount,
		Currency:        msg.Currency,
		Items:           msg.Items,
		ShippingInfo:    msg.ShippingInfo,
		PaymentInfo:     msg.PaymentInfo,
		Metadata:        msg.Metadata,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DueDate:         msg.DueDate,
		FulfillmentType: msg.FulfillmentType,
		Source:          msg.Source,
		Discounts:       msg.Discounts,
		TaxInfo:         msg.TaxInfo,
	}

	// Store the order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_created",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("customer", order.Customer),
			sdk.NewAttribute("merchant", order.Merchant),
			sdk.NewAttribute("status", order.Status),
		),
	)

	return &types.MsgCreateOrderResponse{
		OrderId: order.Id,
	}, nil
}

func (k msgServer) UpdateOrder(goCtx context.Context, msg *types.MsgUpdateOrder) (*types.MsgUpdateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (only customer or merchant can update)
	if msg.Creator != order.Customer && msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only customer or merchant can update order")
	}

	// Update fields
	if msg.Items != nil {
		order.Items = msg.Items
	}
	if msg.ShippingInfo != nil {
		order.ShippingInfo = msg.ShippingInfo
	}
	if msg.PaymentInfo != nil {
		order.PaymentInfo = msg.PaymentInfo
	}
	if msg.Metadata != "" {
		order.Metadata = msg.Metadata
	}
	if msg.DueDate != nil {
		order.DueDate = msg.DueDate
	}

	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_updated",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("updated_by", msg.Creator),
		),
	)

	return &types.MsgUpdateOrderResponse{}, nil
}

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (only customer or merchant can cancel)
	if msg.Creator != order.Customer && msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only customer or merchant can cancel order")
	}

	// Check if order can be cancelled
	if order.Status != "pending" && order.Status != "confirmed" {
		return nil, sdkerrors.Wrap(types.ErrOrderNotCancellable, fmt.Sprintf("order status: %s", order.Status))
	}

	// Update order status
	order.Status = "cancelled"
	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_cancelled",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("cancelled_by", msg.Creator),
			sdk.NewAttribute("reason", msg.Reason),
		),
	)

	return &types.MsgCancelOrderResponse{}, nil
}

func (k msgServer) FulfillOrder(goCtx context.Context, msg *types.MsgFulfillOrder) (*types.MsgFulfillOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (only merchant can fulfill)
	if msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only merchant can fulfill order")
	}

	// Check if order can be fulfilled
	if order.Status != "confirmed" && order.Status != "pending" {
		return nil, sdkerrors.Wrap(types.ErrInvalidStatus, fmt.Sprintf("order status: %s", order.Status))
	}

	// Update order status and shipping info
	order.Status = "shipped"
	order.UpdatedAt = time.Now()
	
	if order.ShippingInfo == nil {
		order.ShippingInfo = &types.ShippingInfo{}
	}
	order.ShippingInfo.TrackingNumber = msg.TrackingNumber
	order.ShippingInfo.Carrier = msg.Carrier

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_fulfilled",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("tracking_number", msg.TrackingNumber),
			sdk.NewAttribute("carrier", msg.Carrier),
		),
	)

	return &types.MsgFulfillOrderResponse{}, nil
}

func (k msgServer) RefundOrder(goCtx context.Context, msg *types.MsgRefundOrder) (*types.MsgRefundOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (only merchant can process refund)
	if msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only merchant can process refund")
	}

	// Update order status
	if msg.PartialRefund {
		order.Status = "partially_refunded"
	} else {
		order.Status = "refunded"
	}
	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_refunded",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("refund_amount", msg.RefundAmount.String()),
			sdk.NewAttribute("reason", msg.Reason),
			sdk.NewAttribute("partial", fmt.Sprintf("%t", msg.PartialRefund)),
		),
	)

	return &types.MsgRefundOrderResponse{}, nil
}

func (k msgServer) UpdateOrderStatus(goCtx context.Context, msg *types.MsgUpdateOrderStatus) (*types.MsgUpdateOrderStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (only merchant can update status)
	if msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only merchant can update order status")
	}

	// Update order status
	order.Status = msg.Status
	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_status_updated",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("old_status", order.Status),
			sdk.NewAttribute("new_status", msg.Status),
			sdk.NewAttribute("notes", msg.Notes),
		),
	)

	return &types.MsgUpdateOrderStatusResponse{}, nil
}

// PayWithStablecoin handles stablecoin payments for orders
func (k msgServer) PayWithStablecoin(goCtx context.Context, msg *types.MsgPayWithStablecoin) (*types.MsgPayWithStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check if payment is already processed
	if order.PaymentInfo != nil && order.PaymentInfo.PaymentStatus == "captured" {
		return nil, sdkerrors.Wrap(types.ErrOrderAlreadyPaid, msg.OrderId)
	}

	// Validate stablecoin
	if !k.stablecoinsKeeper.IsValidStablecoin(ctx, msg.StablecoinDenom) {
		return nil, sdkerrors.Wrap(types.ErrInvalidStablecoin, msg.StablecoinDenom)
	}

	// Parse addresses
	customerAddr, err := sdk.AccAddressFromBech32(msg.CustomerAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid customer address")
	}

	merchantAddr, err := sdk.AccAddressFromBech32(msg.MerchantAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid merchant address")
	}

	// Create stablecoin amount
	stablecoinAmount := sdk.NewCoin(msg.StablecoinDenom, msg.StablecoinAmount)

	// Validate payment amount
	if err := k.stablecoinsKeeper.ValidateStablecoinPayment(ctx, msg.StablecoinDenom, msg.StablecoinAmount); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidPaymentAmount, err.Error())
	}

	var txHash string

	// Handle payment based on escrow preference
	if msg.UseEscrow {
		// Transfer to escrow
		err = k.stablecoinsKeeper.EscrowStablecoin(ctx, customerAddr, msg.OrderId, stablecoinAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrPaymentFailed, err.Error())
		}
		txHash = fmt.Sprintf("escrow_%s_%s", msg.OrderId, msg.StablecoinDenom)
	} else {
		// Direct transfer to merchant
		err = k.stablecoinsKeeper.TransferStablecoin(ctx, customerAddr, merchantAddr, stablecoinAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrPaymentFailed, err.Error())
		}
		txHash = fmt.Sprintf("direct_%s_%s", msg.OrderId, msg.StablecoinDenom)
	}

	// Update order payment info
	if order.PaymentInfo == nil {
		order.PaymentInfo = &types.PaymentInfo{}
	}
	
	order.PaymentInfo.PaymentMethod = "stablecoin"
	order.PaymentInfo.PaymentStatus = "captured"
	order.PaymentInfo.TransactionId = txHash
	order.PaymentInfo.PaymentProcessor = "stateset-stablecoins"
	order.PaymentInfo.AmountPaid = []sdk.Coin{stablecoinAmount}
	order.PaymentInfo.PaymentDate = &time.Now()
	order.PaymentInfo.StablecoinDenom = &msg.StablecoinDenom
	order.PaymentInfo.ExchangeRate = &msg.ExchangeRate
	order.PaymentInfo.UseEscrow = &msg.UseEscrow
	order.PaymentInfo.ConfirmationsRequired = &msg.ConfirmationsRequired
	order.PaymentInfo.EscrowTimeout = msg.EscrowTimeout

	// Update order status
	order.Status = "confirmed"
	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_payment_processed",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("customer", msg.CustomerAddress),
			sdk.NewAttribute("merchant", msg.MerchantAddress),
			sdk.NewAttribute("stablecoin_denom", msg.StablecoinDenom),
			sdk.NewAttribute("amount", msg.StablecoinAmount.String()),
			sdk.NewAttribute("use_escrow", fmt.Sprintf("%t", msg.UseEscrow)),
			sdk.NewAttribute("tx_hash", txHash),
		),
	)

	return &types.MsgPayWithStablecoinResponse{
		TxHash:    txHash,
		Success:   true,
		Message:   "Payment processed successfully",
		Timestamp: time.Now(),
	}, nil
}

// ConfirmStablecoinPayment confirms a stablecoin payment after required confirmations
func (k msgServer) ConfirmStablecoinPayment(goCtx context.Context, msg *types.MsgConfirmStablecoinPayment) (*types.MsgConfirmStablecoinPaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (only merchant can confirm)
	if msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only merchant can confirm payment")
	}

	// Check if payment exists and is pending
	if order.PaymentInfo == nil || order.PaymentInfo.PaymentMethod != "stablecoin" {
		return nil, sdkerrors.Wrap(types.ErrNoStablecoinPayment, msg.OrderId)
	}

	// Update payment status
	order.PaymentInfo.PaymentStatus = "confirmed"
	order.PaymentInfo.ConfirmationCount = &msg.ConfirmationCount
	order.PaymentInfo.ConfirmationBlockHeight = &msg.BlockHeight
	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_payment_confirmed",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("confirmation_count", fmt.Sprintf("%d", msg.ConfirmationCount)),
			sdk.NewAttribute("block_height", fmt.Sprintf("%d", msg.BlockHeight)),
		),
	)

	return &types.MsgConfirmStablecoinPaymentResponse{
		Success: true,
		Message: "Payment confirmed successfully",
	}, nil
}

// RefundStablecoinPayment processes a stablecoin refund
func (k msgServer) RefundStablecoinPayment(goCtx context.Context, msg *types.MsgRefundStablecoinPayment) (*types.MsgRefundStablecoinPaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (only merchant can process refund)
	if msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only merchant can process refund")
	}

	// Check if stablecoin payment exists
	if order.PaymentInfo == nil || order.PaymentInfo.PaymentMethod != "stablecoin" {
		return nil, sdkerrors.Wrap(types.ErrNoStablecoinPayment, msg.OrderId)
	}

	// Parse customer address
	customerAddr, err := sdk.AccAddressFromBech32(msg.CustomerAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid customer address")
	}

	// Create refund amount
	refundAmount := sdk.NewCoin(*order.PaymentInfo.StablecoinDenom, msg.RefundAmount)

	var txHash string

	// Process refund based on original payment method
	if order.PaymentInfo.UseEscrow != nil && *order.PaymentInfo.UseEscrow {
		// Refund from escrow
		err = k.stablecoinsKeeper.RefundEscrow(ctx, msg.OrderId, customerAddr)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrRefundFailed, err.Error())
		}
		txHash = fmt.Sprintf("refund_escrow_%s", msg.OrderId)
	} else {
		// Direct refund from merchant
		merchantAddr, err := sdk.AccAddressFromBech32(order.Merchant)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid merchant address")
		}
		
		err = k.stablecoinsKeeper.TransferStablecoin(ctx, merchantAddr, customerAddr, refundAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrRefundFailed, err.Error())
		}
		txHash = fmt.Sprintf("refund_direct_%s", msg.OrderId)
	}

	// Update order status
	if msg.PartialRefund {
		order.Status = "partially_refunded"
	} else {
		order.Status = "refunded"
	}
	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stablecoin_refund_processed",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("customer", msg.CustomerAddress),
			sdk.NewAttribute("refund_amount", msg.RefundAmount.String()),
			sdk.NewAttribute("partial", fmt.Sprintf("%t", msg.PartialRefund)),
			sdk.NewAttribute("reason", msg.Reason),
			sdk.NewAttribute("tx_hash", txHash),
		),
	)

	return &types.MsgRefundStablecoinPaymentResponse{
		TxHash:  txHash,
		Success: true,
		Message: "Refund processed successfully",
	}, nil
}

// ReleaseEscrow releases escrowed stablecoins to merchant
func (k msgServer) ReleaseEscrow(goCtx context.Context, msg *types.MsgReleaseEscrow) (*types.MsgReleaseEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing order
	order, found := k.Keeper.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrderNotFound, msg.OrderId)
	}

	// Check authorization (customer, merchant, or authorized party can release)
	if msg.Creator != order.Customer && msg.Creator != order.Merchant {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "only customer or merchant can release escrow")
	}

	// Check if escrow exists
	if order.PaymentInfo == nil || order.PaymentInfo.UseEscrow == nil || !*order.PaymentInfo.UseEscrow {
		return nil, sdkerrors.Wrap(types.ErrNoEscrow, msg.OrderId)
	}

	// Parse merchant address
	merchantAddr, err := sdk.AccAddressFromBech32(order.Merchant)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid merchant address")
	}

	// Release escrow to merchant
	err = k.stablecoinsKeeper.ReleaseEscrow(ctx, msg.OrderId, merchantAddr)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrEscrowReleaseFailed, err.Error())
	}

	// Update order status
	order.Status = "completed"
	order.UpdatedAt = time.Now()

	// Store updated order
	k.Keeper.SetOrder(ctx, order)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"escrow_released",
			sdk.NewAttribute("order_id", order.Id),
			sdk.NewAttribute("merchant", order.Merchant),
			sdk.NewAttribute("released_by", msg.Creator),
		),
	)

	return &types.MsgReleaseEscrowResponse{
		Success: true,
		Message: "Escrow released successfully",
	}, nil
}