package keeper

import (
	"context"

	groth16 "github.com/consensys/gnark/std/groth16_bls12377"
	"github.com/stateset/core/x/proof/types"
)

func (k msgServer) VerifyProof(goCtx context.Context, msg *types.MsgVerifyProof) (*types.MsgVerifyProofResponse, error) {

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

