package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/refund/types"
)

func (k msgServer) RejectRefund(goCtx context.Context, msg *types.MsgRejectRefund) (*types.MsgRejectRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRejectRefundResponse{}, nil
}
