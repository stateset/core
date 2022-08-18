package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveRefund = "approve_refund"

var _ sdk.Msg = &MsgApproveRefund{}

func NewMsgApproveRefund(creator string, id uint64) *MsgApproveRefund {
	return &MsgApproveRefund{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgApproveRefund) Route() string {
	return RouterKey
}

func (msg *MsgApproveRefund) Type() string {
	return TypeMsgApproveRefund
}

func (msg *MsgApproveRefund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveRefund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveRefund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
