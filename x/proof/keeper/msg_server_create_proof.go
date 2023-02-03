package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/proof/types"
)

func (k msgServer) CreateProof(goCtx context.Context, msg *types.MsgCreateProof) (*types.MsgCreateProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateProofResponse{}, nil
}
