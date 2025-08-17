package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) CreateTimedoutInvoice(goCtx context.Context, msg *types.MsgCreateTimedoutInvoice) (*types.MsgCreateTimedoutInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var timedoutInvoice = types.TimedoutInvoice{
		Creator: msg.Creator,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	id := k.AppendTimedoutInvoice(
		ctx,
		timedoutInvoice,
	)

	return &types.MsgCreateTimedoutInvoiceResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTimedoutInvoice(goCtx context.Context, msg *types.MsgUpdateTimedoutInvoice) (*types.MsgUpdateTimedoutInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var timedoutInvoice = types.TimedoutInvoice{
		Creator: msg.Creator,
		Id:      msg.Id,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	// Checks that the element exists
	val, found := k.GetTimedoutInvoice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(errorsmod.ErrUnauthorized, "incorrect owner")
	}

	k.SetTimedoutInvoice(ctx, timedoutInvoice)

	return &types.MsgUpdateTimedoutInvoiceResponse{}, nil
}

func (k msgServer) DeleteTimedoutInvoice(goCtx context.Context, msg *types.MsgDeleteTimedoutInvoice) (*types.MsgDeleteTimedoutInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTimedoutInvoice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(errorsmod.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTimedoutInvoice(ctx, msg.Id)

	return &types.MsgDeleteTimedoutInvoiceResponse{}, nil
}
