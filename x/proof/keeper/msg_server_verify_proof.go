package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/proof/types"
	groth16 "github.com/Consensys/gnark"
)

func (k msgServer) VerifyProof(goCtx context.Context, msg *types.MsgVerifyProof) (*types.MsgVerifyProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proof := msg.Proof
	publicWitness := msg.PublicWitness

	vk, err := groth16.VerifierKeyFromPkFile("path/to/verifier_key.pk")
	if err != nil {
		return nil, err
	}

	// Check if proof is valid.
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		return nil, err
	}

	return &types.MsgVerifyProofResponse{}, nil
}

