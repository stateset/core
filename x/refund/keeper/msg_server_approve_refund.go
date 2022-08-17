package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/refund/types"
)

func (k msgServer) ApproveRefund(goCtx context.Context, msg *types.MsgApproveRefund) (*types.MsgApproveRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApproveRefundResponse{}, nil
}
