package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/zkpverify/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns the MsgServer implementation
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// RegisterCircuit registers a new STARK verification circuit
func (m msgServer) RegisterCircuit(
	goCtx context.Context,
	msg *types.MsgRegisterCircuit,
) (*types.MsgRegisterCircuitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Only authority can register circuits
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "only authority can register circuits")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Compute constraint hash from verification key
	vkHash := sha256.Sum256(msg.VerificationKey)
	constraintHash := hex.EncodeToString(vkHash[:])

	circuit := types.Circuit{
		Name:              msg.Name,
		VerificationKey:   msg.VerificationKey,
		ProofSystem:       types.ProofSystemSTARK,
		PublicInputSchema: msg.PublicInputSchema,
		ConstraintHash:    constraintHash,
		Owner:             msg.Authority,
		Description:       msg.Description,
		MaxRecursionDepth: msg.MaxRecursionDepth,
	}

	if err := m.keeper.RegisterCircuit(ctx, circuit); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCircuitRegistered,
			sdk.NewAttribute(types.AttributeKeyCircuitName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyProofSystem, types.ProofSystemSTARK),
			sdk.NewAttribute(types.AttributeKeyConstraintHash, constraintHash),
		),
	)

	return &types.MsgRegisterCircuitResponse{
		CircuitName:    msg.Name,
		ConstraintHash: constraintHash,
	}, nil
}

// DeactivateCircuit deactivates an existing circuit
func (m msgServer) DeactivateCircuit(
	goCtx context.Context,
	msg *types.MsgDeactivateCircuit,
) (*types.MsgDeactivateCircuitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Only authority can deactivate circuits
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "only authority can deactivate circuits")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.DeactivateCircuit(ctx, msg.CircuitName); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCircuitDeactivated,
			sdk.NewAttribute(types.AttributeKeyCircuitName, msg.CircuitName),
		),
	)

	return &types.MsgDeactivateCircuitResponse{}, nil
}

// RegisterSymbolicRule registers a symbolic logic rule for constraint verification
func (m msgServer) RegisterSymbolicRule(
	goCtx context.Context,
	msg *types.MsgRegisterSymbolicRule,
) (*types.MsgRegisterSymbolicRuleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Only authority can register rules
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "only authority can register symbolic rules")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	rule := types.SymbolicRule{
		Name:        msg.RuleName,
		CircuitName: msg.CircuitName,
		RuleType:    msg.RuleType,
		Conditions:  msg.Conditions,
		Description: msg.Description,
	}

	if err := m.keeper.RegisterSymbolicRule(ctx, rule); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRuleRegistered,
			sdk.NewAttribute(types.AttributeKeyCircuitName, msg.CircuitName),
			sdk.NewAttribute(types.AttributeKeyRuleName, msg.RuleName),
		),
	)

	return &types.MsgRegisterSymbolicRuleResponse{
		RuleName: msg.RuleName,
	}, nil
}

// SubmitProof submits a STARK proof for verification
func (m msgServer) SubmitProof(
	goCtx context.Context,
	msg *types.MsgSubmitProof,
) (*types.MsgSubmitProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Store the proof
	proof := &types.Proof{
		CircuitName:     msg.CircuitName,
		ProofData:       msg.ProofData,
		PublicInputs:    msg.PublicInputs,
		DataCommitment:  msg.DataCommitment,
		RecursiveProofs: msg.RecursiveProofs,
		Submitter:       msg.Submitter,
	}

	proofID := m.keeper.StoreProof(ctx, proof)

	// Verify the proof
	result, err := m.keeper.VerifyProof(
		ctx,
		msg.CircuitName,
		msg.ProofData,
		msg.PublicInputs,
		msg.DataCommitment,
		msg.RecursiveProofs,
	)

	if err != nil {
		// Store failed result
		result = types.VerificationResult{
			ProofID:     proofID,
			CircuitName: msg.CircuitName,
			Valid:       false,
			Error:       err.Error(),
		}
		m.keeper.StoreVerificationResult(ctx, result)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProofRejected,
				sdk.NewAttribute(types.AttributeKeyProofID, fmt.Sprintf("%d", proofID)),
				sdk.NewAttribute(types.AttributeKeyCircuitName, msg.CircuitName),
				sdk.NewAttribute(types.AttributeKeyError, err.Error()),
			),
		)

		return &types.MsgSubmitProofResponse{
			ProofID: proofID,
			Valid:   false,
			Error:   err.Error(),
		}, nil
	}

	// Store successful result
	result.ProofID = proofID
	m.keeper.StoreVerificationResult(ctx, result)

	// Store data commitment record
	if len(msg.DataCommitment) > 0 {
		commitmentRecord := types.DataCommitmentRecord{
			Commitment: msg.DataCommitment,
			ProofID:    proofID,
		}
		m.keeper.StoreDataCommitment(ctx, commitmentRecord)
	}

	// Emit appropriate event
	if result.Valid {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProofVerified,
				sdk.NewAttribute(types.AttributeKeyProofID, fmt.Sprintf("%d", proofID)),
				sdk.NewAttribute(types.AttributeKeyCircuitName, msg.CircuitName),
				sdk.NewAttribute(types.AttributeKeyValid, "true"),
				sdk.NewAttribute(types.AttributeKeyDataCommitment, hex.EncodeToString(msg.DataCommitment)),
				sdk.NewAttribute(types.AttributeKeyRecursionDepth, fmt.Sprintf("%d", result.RecursionDepth)),
				sdk.NewAttribute(types.AttributeKeyVerificationTime, fmt.Sprintf("%d", result.VerificationTimeMs)),
			),
		)

		// Emit recursive aggregation event if applicable
		if len(msg.RecursiveProofs) > 0 {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeRecursiveAggregated,
					sdk.NewAttribute(types.AttributeKeyProofID, fmt.Sprintf("%d", proofID)),
					sdk.NewAttribute(types.AttributeKeySubProofCount, fmt.Sprintf("%d", len(msg.RecursiveProofs))),
					sdk.NewAttribute(types.AttributeKeyRecursionDepth, fmt.Sprintf("%d", result.RecursionDepth)),
				),
			)
		}
	} else {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProofRejected,
				sdk.NewAttribute(types.AttributeKeyProofID, fmt.Sprintf("%d", proofID)),
				sdk.NewAttribute(types.AttributeKeyCircuitName, msg.CircuitName),
				sdk.NewAttribute(types.AttributeKeyError, result.Error),
			),
		)
	}

	return &types.MsgSubmitProofResponse{
		ProofID:              proofID,
		Valid:                result.Valid,
		ConstraintsSatisfied: result.ConstraintsSatisfied,
		Error:                result.Error,
	}, nil
}

// ChallengeProof challenges a previously verified proof with fraud evidence
func (m msgServer) ChallengeProof(
	goCtx context.Context,
	msg *types.MsgChallengeProof,
) (*types.MsgChallengeProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	challengeAccepted, err := m.keeper.ProcessChallenge(ctx, msg.ProofID, msg.FraudProof)

	if err != nil {
		return &types.MsgChallengeProofResponse{
			ChallengeAccepted: false,
			ProofInvalidated:  false,
			Reason:            err.Error(),
		}, nil
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeProofChallenged,
			sdk.NewAttribute(types.AttributeKeyProofID, fmt.Sprintf("%d", msg.ProofID)),
			sdk.NewAttribute(types.AttributeKeyValid, fmt.Sprintf("%t", !challengeAccepted)),
		),
	)

	return &types.MsgChallengeProofResponse{
		ChallengeAccepted: challengeAccepted,
		ProofInvalidated:  challengeAccepted,
		Reason:            msg.Reason,
	}, nil
}

// UpdateParams updates module parameters
func (m msgServer) UpdateParams(
	goCtx context.Context,
	msg *types.MsgUpdateParams,
) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "only authority can update params")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.SetParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
