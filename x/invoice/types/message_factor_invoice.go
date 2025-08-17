package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

var _ sdk.Msg = &MsgFactorInvoice{}

func NewMsgFactorInvoice(creator string, id uint64) *MsgFactorInvoice {
	return &MsgFactorInvoice{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgFactorInvoice) Route() string {
	return RouterKey
}

func (msg *MsgFactorInvoice) Type() string {
	return "FactorInvoice"
}

func (msg *MsgFactorInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFactorInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFactorInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
