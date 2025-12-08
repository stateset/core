package types

import (
	"context"
)

// MsgServer defines the zkpverify message server interface
type MsgServer interface {
	// RegisterCircuit registers a new STARK verification circuit
	RegisterCircuit(context.Context, *MsgRegisterCircuit) (*MsgRegisterCircuitResponse, error)

	// DeactivateCircuit deactivates an existing circuit
	DeactivateCircuit(context.Context, *MsgDeactivateCircuit) (*MsgDeactivateCircuitResponse, error)

	// RegisterSymbolicRule registers a symbolic logic rule for constraint verification
	RegisterSymbolicRule(context.Context, *MsgRegisterSymbolicRule) (*MsgRegisterSymbolicRuleResponse, error)

	// SubmitProof submits a STARK proof for verification
	SubmitProof(context.Context, *MsgSubmitProof) (*MsgSubmitProofResponse, error)

	// ChallengeProof challenges a previously verified proof with fraud evidence
	ChallengeProof(context.Context, *MsgChallengeProof) (*MsgChallengeProofResponse, error)

	// UpdateParams updates module parameters
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// QueryServer defines the zkpverify query server interface
type QueryServer interface {
	// Circuit returns circuit details by name
	Circuit(context.Context, *QueryCircuitRequest) (*QueryCircuitResponse, error)

	// Circuits returns all registered circuits
	Circuits(context.Context, *QueryCircuitsRequest) (*QueryCircuitsResponse, error)

	// Proof returns proof details by ID
	Proof(context.Context, *QueryProofRequest) (*QueryProofResponse, error)

	// VerificationResult returns the verification result for a proof
	VerificationResult(context.Context, *QueryVerificationResultRequest) (*QueryVerificationResultResponse, error)

	// SymbolicRules returns all symbolic rules for a circuit
	SymbolicRules(context.Context, *QuerySymbolicRulesRequest) (*QuerySymbolicRulesResponse, error)

	// Params returns module parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// Query request/response types

type QueryCircuitRequest struct {
	Name string `json:"name"`
}

type QueryCircuitResponse struct {
	Circuit Circuit `json:"circuit"`
}

type QueryCircuitsRequest struct {
	ActiveOnly bool `json:"active_only"`
}

type QueryCircuitsResponse struct {
	Circuits []Circuit `json:"circuits"`
}

type QueryProofRequest struct {
	ProofID uint64 `json:"proof_id"`
}

type QueryProofResponse struct {
	Proof Proof `json:"proof"`
}

type QueryVerificationResultRequest struct {
	ProofID uint64 `json:"proof_id"`
}

type QueryVerificationResultResponse struct {
	Result VerificationResult `json:"result"`
}

type QuerySymbolicRulesRequest struct {
	CircuitName string `json:"circuit_name"`
}

type QuerySymbolicRulesResponse struct {
	Rules []SymbolicRule `json:"rules"`
}

type QueryParamsRequest struct{}

type QueryParamsResponse struct {
	Params Params `json:"params"`
}
