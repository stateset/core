package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

const TypeMsgVerifyProof = "verify_proof"

var _ sdk.Msg = &MsgVerifyProof{}

func NewMsgVerifyProof(creator string, proof string, publicWitness string) *MsgVerifyProof {
	return &MsgVerifyProof{
		Creator:       creator,
		Proof:         proof,
		PublicWitness: publicWitness,
	}
}

func (msg *MsgVerifyProof) Route() string {
	return RouterKey
}

func (msg *MsgVerifyProof) Type() string {
	return TypeMsgVerifyProof
}

func (msg *MsgVerifyProof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVerifyProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVerifyProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
