package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

var _ sdk.Msg = &MsgCreateTimedoutInvoice{}

func NewMsgCreateTimedoutInvoice(creator string, did string, chain string) *MsgCreateTimedoutInvoice {
	return &MsgCreateTimedoutInvoice{
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgCreateTimedoutInvoice) Route() string {
	return RouterKey
}

func (msg *MsgCreateTimedoutInvoice) Type() string {
	return "CreateTimedoutInvoice"
}

func (msg *MsgCreateTimedoutInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTimedoutInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTimedoutInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTimedoutInvoice{}

func NewMsgUpdateTimedoutInvoice(creator string, id uint64, did string, chain string) *MsgUpdateTimedoutInvoice {
	return &MsgUpdateTimedoutInvoice{
		Id:      id,
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgUpdateTimedoutInvoice) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTimedoutInvoice) Type() string {
	return "UpdateTimedoutInvoice"
}

func (msg *MsgUpdateTimedoutInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTimedoutInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTimedoutInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTimedoutInvoice{}

func NewMsgDeleteTimedoutInvoice(creator string, id uint64) *MsgDeleteTimedoutInvoice {
	return &MsgDeleteTimedoutInvoice{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTimedoutInvoice) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTimedoutInvoice) Type() string {
	return "DeleteTimedoutInvoice"
}

func (msg *MsgDeleteTimedoutInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTimedoutInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTimedoutInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
