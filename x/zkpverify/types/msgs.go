package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgRegisterCircuit)(nil)
	_ sdk.Msg = (*MsgDeactivateCircuit)(nil)
	_ sdk.Msg = (*MsgRegisterSymbolicRule)(nil)
	_ sdk.Msg = (*MsgSubmitProof)(nil)
	_ sdk.Msg = (*MsgChallengeProof)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
)

// MsgRegisterCircuit registers a new STARK verification circuit
type MsgRegisterCircuit struct {
	Authority         string             `json:"authority"`
	Name              string             `json:"name"`
	VerificationKey   []byte             `json:"verification_key"`
	PublicInputSchema []PublicInputField `json:"public_input_schema"`
	Description       string             `json:"description"`
	MaxRecursionDepth uint32             `json:"max_recursion_depth"`
}

func (m *MsgRegisterCircuit) Reset()         { *m = MsgRegisterCircuit{} }
func (m *MsgRegisterCircuit) String() string { return "MsgRegisterCircuit" }
func (m *MsgRegisterCircuit) ProtoMessage()  {}

// TypeURL returns the type URL for this message (implements proto.Message)
func (m *MsgRegisterCircuit) TypeURL() string { return "/zkpverify.MsgRegisterCircuit" }

func (msg MsgRegisterCircuit) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

func (msg MsgRegisterCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrUnauthorized
	}
	if msg.Name == "" {
		return ErrEmptyCircuitName
	}
	if len(msg.VerificationKey) == 0 {
		return ErrInvalidVerificationKey
	}
	return nil
}

// MsgDeactivateCircuit deactivates an existing circuit
type MsgDeactivateCircuit struct {
	Authority   string `json:"authority"`
	CircuitName string `json:"circuit_name"`
}

func (m *MsgDeactivateCircuit) Reset()         { *m = MsgDeactivateCircuit{} }
func (m *MsgDeactivateCircuit) String() string { return "MsgDeactivateCircuit" }
func (m *MsgDeactivateCircuit) ProtoMessage()  {}

// TypeURL returns the type URL for this message (implements proto.Message)
func (m *MsgDeactivateCircuit) TypeURL() string { return "/zkpverify.MsgDeactivateCircuit" }

func (msg MsgDeactivateCircuit) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

func (msg MsgDeactivateCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrUnauthorized
	}
	if msg.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	return nil
}

// MsgRegisterSymbolicRule registers a logical constraint rule for a circuit
type MsgRegisterSymbolicRule struct {
	Authority   string      `json:"authority"`
	CircuitName string      `json:"circuit_name"`
	RuleName    string      `json:"rule_name"`
	RuleType    RuleType    `json:"rule_type"`
	Conditions  []Condition `json:"conditions"`
	Description string      `json:"description"`
}

func (m *MsgRegisterSymbolicRule) Reset()         { *m = MsgRegisterSymbolicRule{} }
func (m *MsgRegisterSymbolicRule) String() string { return "MsgRegisterSymbolicRule" }
func (m *MsgRegisterSymbolicRule) ProtoMessage()  {}

// TypeURL returns the type URL for this message (implements proto.Message)
func (m *MsgRegisterSymbolicRule) TypeURL() string { return "/zkpverify.MsgRegisterSymbolicRule" }

func (msg MsgRegisterSymbolicRule) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

func (msg MsgRegisterSymbolicRule) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrUnauthorized
	}
	if msg.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	if msg.RuleName == "" {
		return ErrInvalidRuleDefinition
	}
	if len(msg.Conditions) == 0 {
		return ErrInvalidRuleDefinition
	}
	return nil
}

// MsgSubmitProof submits a STARK proof for verification
type MsgSubmitProof struct {
	Submitter       string   `json:"submitter"`
	CircuitName     string   `json:"circuit_name"`
	ProofData       []byte   `json:"proof_data"`
	PublicInputs    []byte   `json:"public_inputs"`
	DataCommitment  []byte   `json:"data_commitment"`
	RecursiveProofs []uint64 `json:"recursive_proofs,omitempty"`
}

func (m *MsgSubmitProof) Reset()         { *m = MsgSubmitProof{} }
func (m *MsgSubmitProof) String() string { return "MsgSubmitProof" }
func (m *MsgSubmitProof) ProtoMessage()  {}

// TypeURL returns the type URL for this message (implements proto.Message)
func (m *MsgSubmitProof) TypeURL() string { return "/zkpverify.MsgSubmitProof" }

func (msg MsgSubmitProof) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Submitter)
	return []sdk.AccAddress{addr}
}

func (msg MsgSubmitProof) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Submitter); err != nil {
		return ErrUnauthorized
	}
	if msg.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	if len(msg.ProofData) == 0 {
		return ErrEmptyProofData
	}
	if len(msg.DataCommitment) == 0 {
		return ErrInvalidDataCommitment
	}
	return nil
}

// MsgChallengeProof challenges a previously verified proof
type MsgChallengeProof struct {
	Challenger string `json:"challenger"`
	ProofID    uint64 `json:"proof_id"`
	FraudProof []byte `json:"fraud_proof"`
	Reason     string `json:"reason"`
}

func (m *MsgChallengeProof) Reset()         { *m = MsgChallengeProof{} }
func (m *MsgChallengeProof) String() string { return "MsgChallengeProof" }
func (m *MsgChallengeProof) ProtoMessage()  {}

// TypeURL returns the type URL for this message (implements proto.Message)
func (m *MsgChallengeProof) TypeURL() string { return "/zkpverify.MsgChallengeProof" }

func (msg MsgChallengeProof) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Challenger)
	return []sdk.AccAddress{addr}
}

func (msg MsgChallengeProof) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Challenger); err != nil {
		return ErrUnauthorized
	}
	if msg.ProofID == 0 {
		return ErrProofNotFound
	}
	if len(msg.FraudProof) == 0 {
		return ErrInvalidChallenge
	}
	return nil
}

// MsgUpdateParams updates the module parameters
type MsgUpdateParams struct {
	Authority string `json:"authority"`
	Params    Params `json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return "MsgUpdateParams" }
func (m *MsgUpdateParams) ProtoMessage()  {}

// TypeURL returns the type URL for this message (implements proto.Message)
func (m *MsgUpdateParams) TypeURL() string { return "/zkpverify.MsgUpdateParams" }

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrUnauthorized
	}
	return msg.Params.Validate()
}

// Response types

type MsgRegisterCircuitResponse struct {
	CircuitName    string `json:"circuit_name"`
	ConstraintHash string `json:"constraint_hash"`
}

type MsgDeactivateCircuitResponse struct{}

type MsgRegisterSymbolicRuleResponse struct {
	RuleName string `json:"rule_name"`
}

type MsgSubmitProofResponse struct {
	ProofID            uint64   `json:"proof_id"`
	Valid              bool     `json:"valid"`
	ConstraintsSatisfied []string `json:"constraints_satisfied"`
	Error              string   `json:"error,omitempty"`
}

type MsgChallengeProofResponse struct {
	ChallengeAccepted bool   `json:"challenge_accepted"`
	ProofInvalidated  bool   `json:"proof_invalidated"`
	Reason            string `json:"reason"`
}

type MsgUpdateParamsResponse struct{}
