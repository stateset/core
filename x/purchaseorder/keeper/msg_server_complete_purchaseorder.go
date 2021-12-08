package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/purchaseorder/types"
)

func (k msgServer) CompletePurchaseorder(goCtx context.Context, msg *types.MsgCompletePurchaseorder) (*types.MsgCompletePurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCompletePurchaseorderResponse{}, nil
}
