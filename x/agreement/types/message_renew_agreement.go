package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRenewAgreement{}

func NewMsgRenewAgreement(creator string, id uint64) *MsgRenewAgreement {
	return &MsgRenewAgreement{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRenewAgreement) Route() string {
	return RouterKey
}

func (msg *MsgRenewAgreement) Type() string {
	return "RenewAgreement"
}

func (msg *MsgRenewAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRenewAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRenewAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
