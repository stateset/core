package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) CreateSentInvoice(goCtx context.Context, msg *types.MsgCreateSentInvoice) (*types.MsgCreateSentInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var sentInvoice = types.SentInvoice{
		Creator: msg.Creator,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	id := k.AppendSentInvoice(
		ctx,
		sentInvoice,
	)

	return &types.MsgCreateSentInvoiceResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSentInvoice(goCtx context.Context, msg *types.MsgUpdateSentInvoice) (*types.MsgUpdateSentInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var sentInvoice = types.SentInvoice{
		Creator: msg.Creator,
		Id:      msg.Id,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	// Checks that the element exists
	val, found := k.GetSentInvoice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(errorsmod.ErrUnauthorized, "incorrect owner")
	}

	k.SetSentInvoice(ctx, sentInvoice)

	return &types.MsgUpdateSentInvoiceResponse{}, nil
}

func (k msgServer) DeleteSentInvoice(goCtx context.Context, msg *types.MsgDeleteSentInvoice) (*types.MsgDeleteSentInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetSentInvoice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(errorsmod.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSentInvoice(ctx, msg.Id)

	return &types.MsgDeleteSentInvoiceResponse{}, nil
}
