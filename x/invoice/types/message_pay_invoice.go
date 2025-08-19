package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPayInvoice = "pay_invoice"

var _ sdk.Msg = &MsgPayInvoice{}

func NewMsgPayInvoice(creator string, id uint64) *MsgPayInvoice {
	return &MsgPayInvoice{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgPayInvoice) Route() string {
	return RouterKey
}

func (msg *MsgPayInvoice) Type() string {
	return TypeMsgPayInvoice
}

func (msg *MsgPayInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPayInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPayInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
