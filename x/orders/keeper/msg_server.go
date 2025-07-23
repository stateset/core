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