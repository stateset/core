package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCancelPurchaseorder{}

func NewMsgCancelPurchaseorder(creator string, id uint64) *MsgCancelPurchaseorder {
	return &MsgCancelPurchaseorder{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgCancelPurchaseorder) Type() string {
	return "CancelPurchaseorder"
}

func (msg *MsgCancelPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
