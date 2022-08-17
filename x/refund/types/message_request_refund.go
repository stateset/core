package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestRefund = "request_refund"

var _ sdk.Msg = &MsgRequestRefund{}

func NewMsgRequestRefund(creator string, did string, amount string, fee string, deadline string) *MsgRequestRefund {
	return &MsgRequestRefund{
		Creator:  creator,
		Did:      did,
		Amount:   amount,
		Fee:      fee,
		Deadline: deadline,
	}
}

func (msg *MsgRequestRefund) Route() string {
	return RouterKey
}

func (msg *MsgRequestRefund) Type() string {
	return TypeMsgRequestRefund
}

func (msg *MsgRequestRefund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestRefund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestRefund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
