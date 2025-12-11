package types

import "context"

// MsgCreateOrderResponse returns the created order id.
type MsgCreateOrderResponse struct {
	OrderId uint64 `json:"order_id"`
}

type MsgConfirmOrderResponse struct{}

type MsgPayOrderResponse struct{}

type MsgShipOrderResponse struct{}

type MsgDeliverOrderResponse struct{}

type MsgCompleteOrderResponse struct{}

type MsgCancelOrderResponse struct{}

type MsgRefundOrderResponse struct{}

// MsgOpenDisputeResponse returns the dispute id.
type MsgOpenDisputeResponse struct {
	DisputeId uint64 `json:"dispute_id"`
}

type MsgResolveDisputeResponse struct{}

// MsgServer defines the gRPC Msg service for orders.
type MsgServer interface {
	CreateOrder(ctx context.Context, msg *MsgCreateOrder) (*MsgCreateOrderResponse, error)
	ConfirmOrder(ctx context.Context, msg *MsgConfirmOrder) (*MsgConfirmOrderResponse, error)
	PayOrder(ctx context.Context, msg *MsgPayOrder) (*MsgPayOrderResponse, error)
	ShipOrder(ctx context.Context, msg *MsgShipOrder) (*MsgShipOrderResponse, error)
	DeliverOrder(ctx context.Context, msg *MsgDeliverOrder) (*MsgDeliverOrderResponse, error)
	CompleteOrder(ctx context.Context, msg *MsgCompleteOrder) (*MsgCompleteOrderResponse, error)
	CancelOrder(ctx context.Context, msg *MsgCancelOrder) (*MsgCancelOrderResponse, error)
	RefundOrder(ctx context.Context, msg *MsgRefundOrder) (*MsgRefundOrderResponse, error)
	OpenDispute(ctx context.Context, msg *MsgOpenDispute) (*MsgOpenDisputeResponse, error)
	ResolveDispute(ctx context.Context, msg *MsgResolveDispute) (*MsgResolveDisputeResponse, error)
}
