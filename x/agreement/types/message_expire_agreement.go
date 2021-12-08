package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgExpireAgreement{}

func NewMsgExpireAgreement(creator string, id uint64) *MsgExpireAgreement {
	return &MsgExpireAgreement{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgExpireAgreement) Route() string {
	return RouterKey
}

func (msg *MsgExpireAgreement) Type() string {
	return "ExpireAgreement"
}

func (msg *MsgExpireAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExpireAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExpireAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
