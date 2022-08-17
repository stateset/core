package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/refund/types"
)

func (k msgServer) RequestRefund(goCtx context.Context, msg *types.MsgRequestRefund) (*types.MsgRequestRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestRefundResponse{}, nil
}
