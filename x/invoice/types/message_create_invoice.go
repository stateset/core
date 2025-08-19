package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateInvoice = "create_invoice"

var _ sdk.Msg = &MsgCreateInvoice{}

func NewMsgCreateInvoice(creator string, id string, did string, amount string, state string, purchaser string) *MsgCreateInvoice {
	return &MsgCreateInvoice{
		Creator:   creator,
		Id:        id,
		Did:       did,
		Amount:    amount,
		State:     state,
		Seller:    creator,
		Purchaser: purchaser,
	}
}

func (msg *MsgCreateInvoice) Route() string {
	return RouterKey
}

func (msg *MsgCreateInvoice) Type() string {
	return TypeMsgCreateInvoice
}

func (msg *MsgCreateInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
