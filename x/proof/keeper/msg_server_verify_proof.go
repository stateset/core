package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/proof/types"
)

func (k msgServer) VerifyProof(goCtx context.Context, msg *types.MsgVerifyProof) (*types.MsgVerifyProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgVerifyProofResponse{}, nil
}
