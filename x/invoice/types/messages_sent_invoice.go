package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	errorsmod "cosmossdk.io/errors"
)

var _ sdk.Msg = &MsgCreateSentInvoice{}

func NewMsgCreateSentInvoice(creator string, did string, chain string) *MsgCreateSentInvoice {
	return &MsgCreateSentInvoice{
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgCreateSentInvoice) Route() string {
	return RouterKey
}

func (msg *MsgCreateSentInvoice) Type() string {
	return "CreateSentInvoice"
}

func (msg *MsgCreateSentInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSentInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSentInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSentInvoice{}

func NewMsgUpdateSentInvoice(creator string, id uint64, did string, chain string) *MsgUpdateSentInvoice {
	return &MsgUpdateSentInvoice{
		Id:      id,
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgUpdateSentInvoice) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSentInvoice) Type() string {
	return "UpdateSentInvoice"
}

func (msg *MsgUpdateSentInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSentInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSentInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSentInvoice{}

func NewMsgDeleteSentInvoice(creator string, id uint64) *MsgDeleteSentInvoice {
	return &MsgDeleteSentInvoice{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSentInvoice) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSentInvoice) Type() string {
	return "DeleteSentInvoice"
}

func (msg *MsgDeleteSentInvoice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSentInvoice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSentInvoice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
