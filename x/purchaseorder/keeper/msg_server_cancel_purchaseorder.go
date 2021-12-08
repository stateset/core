package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/purchaseorder/types"
)

func (k msgServer) CancelPurchaseorder(goCtx context.Context, msg *types.MsgCancelPurchaseorder) (*types.MsgCancelPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	purchaseorder, found := k.GetPurchaseorder(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if purchaseorder.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongPurchaseOrderState, "%v", purchaseorder.State)
	}

	purchaseorder.State = "cancelled"

	k.SetPurchaseorder(ctx, purchaseorder)

	return &types.MsgCancelPurchaseorderResponse{}, nil
}
