package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

var _ sdk.Msg = &MsgTerminateAgreement{}

func NewMsgTerminateAgreement(creator string, id uint64) *MsgTerminateAgreement {
	return &MsgTerminateAgreement{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgTerminateAgreement) Route() string {
	return RouterKey
}

func (msg *MsgTerminateAgreement) Type() string {
	return "TerminateAgreement"
}

func (msg *MsgTerminateAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTerminateAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTerminateAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
