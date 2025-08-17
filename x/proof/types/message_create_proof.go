package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

const TypeMsgCreateProof = "create_proof"

var _ sdk.Msg = &MsgCreateProof{}

func NewMsgCreateProof(creator string, id uint64, did string, uri string, hash string, state string) *MsgCreateProof {
	return &MsgCreateProof{
		Creator: creator,
		Id:      id,
		Did:     did,
		Uri:     uri,
		Hash:    hash,
		State:   state,
	}
}

func (msg *MsgCreateProof) Route() string {
	return RouterKey
}

func (msg *MsgCreateProof) Type() string {
	return TypeMsgCreateProof
}

func (msg *MsgCreateProof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
