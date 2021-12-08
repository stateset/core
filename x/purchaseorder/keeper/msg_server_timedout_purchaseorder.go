package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/purchaseorder/types"
)

func (k msgServer) CreateTimedoutPurchaseorder(goCtx context.Context, msg *types.MsgCreateTimedoutPurchaseorder) (*types.MsgCreateTimedoutPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var timedoutPurchaseorder = types.TimedoutPurchaseorder{
		Creator: msg.Creator,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	id := k.AppendTimedoutPurchaseorder(
		ctx,
		timedoutPurchaseorder,
	)

	return &types.MsgCreateTimedoutPurchaseorderResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTimedoutPurchaseorder(goCtx context.Context, msg *types.MsgUpdateTimedoutPurchaseorder) (*types.MsgUpdateTimedoutPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var timedoutPurchaseorder = types.TimedoutPurchaseorder{
		Creator: msg.Creator,
		Id:      msg.Id,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	// Checks that the element exists
	val, found := k.GetTimedoutPurchaseorder(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTimedoutPurchaseorder(ctx, timedoutPurchaseorder)

	return &types.MsgUpdateTimedoutPurchaseorderResponse{}, nil
}

func (k msgServer) DeleteTimedoutPurchaseorder(goCtx context.Context, msg *types.MsgDeleteTimedoutPurchaseorder) (*types.MsgDeleteTimedoutPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTimedoutPurchaseorder(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTimedoutPurchaseorder(ctx, msg.Id)

	return &types.MsgDeleteTimedoutPurchaseorderResponse{}, nil
}
