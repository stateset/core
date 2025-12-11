package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/orders/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the orders MsgServer backed by keeper logic.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{keeper: k}
}

func (m msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	id, err := m.keeper.CreateOrder(ctx, msg.Customer, msg.Merchant, msg.Items, msg.ShippingInfo, msg.Metadata)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateOrderResponse{OrderId: id}, nil
}

func (m msgServer) ConfirmOrder(goCtx context.Context, msg *types.MsgConfirmOrder) (*types.MsgConfirmOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.ConfirmOrder(ctx, msg.Merchant, msg.OrderId); err != nil {
		return nil, err
	}
	return &types.MsgConfirmOrderResponse{}, nil
}

func (m msgServer) PayOrder(goCtx context.Context, msg *types.MsgPayOrder) (*types.MsgPayOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.PayOrder(ctx, msg.Customer, msg.OrderId, msg.Amount, msg.UseEscrow); err != nil {
		return nil, err
	}
	return &types.MsgPayOrderResponse{}, nil
}

func (m msgServer) ShipOrder(goCtx context.Context, msg *types.MsgShipOrder) (*types.MsgShipOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.ShipOrder(ctx, msg.Merchant, msg.OrderId, msg.Carrier, msg.TrackingNumber); err != nil {
		return nil, err
	}
	return &types.MsgShipOrderResponse{}, nil
}

func (m msgServer) DeliverOrder(goCtx context.Context, msg *types.MsgDeliverOrder) (*types.MsgDeliverOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.DeliverOrder(ctx, msg.Signer, msg.OrderId); err != nil {
		return nil, err
	}
	return &types.MsgDeliverOrderResponse{}, nil
}

func (m msgServer) CompleteOrder(goCtx context.Context, msg *types.MsgCompleteOrder) (*types.MsgCompleteOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.CompleteOrder(ctx, msg.Customer, msg.OrderId); err != nil {
		return nil, err
	}
	return &types.MsgCompleteOrderResponse{}, nil
}

func (m msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.CancelOrder(ctx, msg.Signer, msg.OrderId, msg.Reason); err != nil {
		return nil, err
	}
	return &types.MsgCancelOrderResponse{}, nil
}

func (m msgServer) RefundOrder(goCtx context.Context, msg *types.MsgRefundOrder) (*types.MsgRefundOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.RefundOrder(ctx, msg.Merchant, msg.OrderId, msg.RefundAmount, msg.Reason, msg.FullRefund); err != nil {
		return nil, err
	}
	return &types.MsgRefundOrderResponse{}, nil
}

func (m msgServer) OpenDispute(goCtx context.Context, msg *types.MsgOpenDispute) (*types.MsgOpenDisputeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	id, err := m.keeper.OpenDispute(ctx, msg.Customer, msg.OrderId, msg.Reason, msg.Description, msg.Evidence)
	if err != nil {
		return nil, err
	}
	return &types.MsgOpenDisputeResponse{DisputeId: id}, nil
}

func (m msgServer) ResolveDispute(goCtx context.Context, msg *types.MsgResolveDispute) (*types.MsgResolveDisputeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.keeper.ResolveDispute(ctx, msg.Authority, msg.DisputeId, msg.Resolution, msg.RefundAmount, msg.ToCustomer); err != nil {
		return nil, err
	}
	return &types.MsgResolveDisputeResponse{}, nil
}
