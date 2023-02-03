package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/proof/types"
)

func (k msgServer) CreateProof(goCtx context.Context, msg *types.MsgCreateProof) (*types.MsgCreateProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var proof = types.Proof{
		Id:     	msg.Id,
		Did:        msg.Did,
		Uri: 		msg.Uri,
		Hash:       msg.Hash,
		State:      "created",
	}

	k.AppendProof(ctx, proof)

	return &types.MsgCreateProofResponse{}, nil
}
