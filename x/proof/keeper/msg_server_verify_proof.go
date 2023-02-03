package keeper

import (
	"context"
	ecc "github.com/consensys/gnark-crypto/ecc"
	groth16 "github.com/consensys/gnark/backend/groth16"
	"github.com/stateset/core/x/proof/types"
)

func (k msgServer) VerifyProof(goCtx context.Context, msg *types.MsgVerifyProof) (*types.MsgVerifyProofResponse, error) {

	proof := msg.Proof
	publicWitness := msg.PublicWitness
	
	// Set the curve to use.
	curve := ecc.BN254.init();

	// Create a new verifying key.
	vk, err := groth16.NewVerifyingKey(curve)
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

