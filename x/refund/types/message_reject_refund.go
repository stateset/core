package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectRefund = "reject_refund"

var _ sdk.Msg = &MsgRejectRefund{}

func NewMsgRejectRefund(creator string, id uint64) *MsgRejectRefund {
	return &MsgRejectRefund{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRejectRefund) Route() string {
	return RouterKey
}

func (msg *MsgRejectRefund) Type() string {
	return TypeMsgRejectRefund
}

func (msg *MsgRejectRefund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRejectRefund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectRefund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
