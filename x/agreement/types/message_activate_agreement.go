package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgActivateAgreement{}

func NewMsgActivateAgreement(creator string, id uint64) *MsgActivateAgreement {
	return &MsgActivateAgreement{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgActivateAgreement) Route() string {
	return RouterKey
}

func (msg *MsgActivateAgreement) Type() string {
	return "ActivateAgreement"
}

func (msg *MsgActivateAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgActivateAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgActivateAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
