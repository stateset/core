package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFinancePurchaseorder{}

func NewMsgFinancePurchaseorder(creator string, id uint64) *MsgFinancePurchaseorder {
	return &MsgFinancePurchaseorder{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgFinancePurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgFinancePurchaseorder) Type() string {
	return "FinancePurchaseorder"
}

func (msg *MsgFinancePurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinancePurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinancePurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
