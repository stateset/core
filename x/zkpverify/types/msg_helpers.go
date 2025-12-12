package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m MsgRegisterCircuit) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgRegisterCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return ErrUnauthorized
	}
	if m.Name == "" {
		return ErrEmptyCircuitName
	}
	if len(m.VerificationKey) == 0 {
		return ErrInvalidVerificationKey
	}
	return nil
}

func (m MsgDeactivateCircuit) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgDeactivateCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return ErrUnauthorized
	}
	if m.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	return nil
}

func (m MsgRegisterSymbolicRule) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgRegisterSymbolicRule) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return ErrUnauthorized
	}
	if m.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	if m.RuleName == "" {
		return ErrInvalidRuleDefinition
	}
	if len(m.Conditions) == 0 {
		return ErrInvalidRuleDefinition
	}
	return nil
}

func (m MsgSubmitProof) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Submitter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgSubmitProof) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Submitter); err != nil {
		return ErrUnauthorized
	}
	if m.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	if len(m.ProofData) == 0 {
		return ErrEmptyProofData
	}
	if len(m.DataCommitment) == 0 {
		return ErrInvalidDataCommitment
	}
	return nil
}

func (m MsgChallengeProof) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Challenger)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgChallengeProof) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Challenger); err != nil {
		return ErrUnauthorized
	}
	if m.ProofID == 0 {
		return ErrProofNotFound
	}
	if len(m.FraudProof) == 0 {
		return ErrInvalidChallenge
	}
	return nil
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return ErrUnauthorized
	}
	return m.Params.Validate()
}
