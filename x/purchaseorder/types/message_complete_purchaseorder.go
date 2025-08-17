package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

var _ sdk.Msg = &MsgCompletePurchaseorder{}

func NewMsgCompletePurchaseorder(creator string, id uint64) *MsgCompletePurchaseorder {
	return &MsgCompletePurchaseorder{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCompletePurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgCompletePurchaseorder) Type() string {
	return "CompletePurchaseorder"
}

func (msg *MsgCompletePurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCompletePurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCompletePurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
