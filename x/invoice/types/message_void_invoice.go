package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

const TypeMsgVoidInvoice = "void_invoice"

var _ sdk.Msg = &MsgVoidInvoice{}

func NewMsgVoidInvoice(creator string, id uint64) *MsgVoidInvoice {
	return &MsgVoidInvoice{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgVoidInvoice) Route() string {
	return RouterKey
}

func (msg *MsgVoidInvoice) Type() string {
	return TypeMsgVoidInvoice
}

func (msg *MsgVoidInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoidInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoidInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
