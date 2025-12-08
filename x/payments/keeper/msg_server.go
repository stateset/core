package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/payments/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the payments MsgServer backed by keeper logic.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{keeper: k}
}

func (m msgServer) CreatePayment(goCtx context.Context, msg *types.MsgCreatePayment) (*types.MsgCreatePaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	intent := types.PaymentIntent{
		Payer:    msg.Payer,
		Payee:    msg.Payee,
		Amount:   msg.Amount,
		Metadata: msg.Metadata,
	}
	if err := intent.ValidateBasic(); err != nil {
		return nil, err
	}

	id, err := m.keeper.CreatePayment(ctx, intent)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreated,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Payer),
			sdk.NewAttribute(types.AttributeKeyPayer, msg.Payer),
			sdk.NewAttribute(types.AttributeKeyPayee, msg.Payee),
			sdk.NewAttribute(types.AttributeKeyID, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreatePaymentResponse{PaymentId: id}, nil
}

func (m msgServer) SettlePayment(goCtx context.Context, msg *types.MsgSettlePayment) (*types.MsgSettlePaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	payee, err := sdk.AccAddressFromBech32(msg.Payee)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidPayment, err.Error())
	}

	if err := m.keeper.SettlePayment(ctx, msg.PaymentId, payee); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSettled,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Payee),
			sdk.NewAttribute(types.AttributeKeyPayee, msg.Payee),
			sdk.NewAttribute(types.AttributeKeyID, strconv.FormatUint(msg.PaymentId, 10)),
		),
	)

	return &types.MsgSettlePaymentResponse{}, nil
}

func (m msgServer) CancelPayment(goCtx context.Context, msg *types.MsgCancelPayment) (*types.MsgCancelPaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	payer, err := sdk.AccAddressFromBech32(msg.Payer)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidPayment, err.Error())
	}

	if err := m.keeper.CancelPayment(ctx, msg.PaymentId, payer); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCancelled,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Payer),
			sdk.NewAttribute(types.AttributeKeyPayer, msg.Payer),
			sdk.NewAttribute(types.AttributeKeyID, strconv.FormatUint(msg.PaymentId, 10)),
		),
	)

	return &types.MsgCancelPaymentResponse{}, nil
}
