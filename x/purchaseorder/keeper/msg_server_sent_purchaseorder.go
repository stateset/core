package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/x/purchaseorder/types"
)

func (k msgServer) CreateSentPurchaseorder(goCtx context.Context, msg *types.MsgCreateSentPurchaseorder) (*types.MsgCreateSentPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var sentPurchaseorder = types.SentPurchaseorder{
		Creator: msg.Creator,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	id := k.AppendSentPurchaseorder(
		ctx,
		sentPurchaseorder,
	)

	return &types.MsgCreateSentPurchaseorderResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSentPurchaseorder(goCtx context.Context, msg *types.MsgUpdateSentPurchaseorder) (*types.MsgUpdateSentPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var sentPurchaseorder = types.SentPurchaseorder{
		Creator: msg.Creator,
		Id:      msg.Id,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	// Checks that the element exists
	val, found := k.GetSentPurchaseorder(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(errorsmod.ErrUnauthorized, "incorrect owner")
	}

	k.SetSentPurchaseorder(ctx, sentPurchaseorder)

	return &types.MsgUpdateSentPurchaseorderResponse{}, nil
}

func (k msgServer) DeleteSentPurchaseorder(goCtx context.Context, msg *types.MsgDeleteSentPurchaseorder) (*types.MsgDeleteSentPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetSentPurchaseorder(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(errorsmod.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSentPurchaseorder(ctx, msg.Id)

	return &types.MsgDeleteSentPurchaseorderResponse{}, nil
}
