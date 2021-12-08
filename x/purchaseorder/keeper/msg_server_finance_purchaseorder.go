package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/purchaseorder/types"
)

func (k msgServer) FinancePurchaseorder(goCtx context.Context, msg *types.MsgFinancePurchaseorder) (*types.MsgFinancePurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgFinancePurchaseorderResponse{}, nil
}
