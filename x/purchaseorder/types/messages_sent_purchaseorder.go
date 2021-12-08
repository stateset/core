package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSentPurchaseorder{}

func NewMsgCreateSentPurchaseorder(creator string, did string, chain string) *MsgCreateSentPurchaseorder {
	return &MsgCreateSentPurchaseorder{
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgCreateSentPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgCreateSentPurchaseorder) Type() string {
	return "CreateSentPurchaseorder"
}

func (msg *MsgCreateSentPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSentPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSentPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSentPurchaseorder{}

func NewMsgUpdateSentPurchaseorder(creator string, id uint64, did string, chain string) *MsgUpdateSentPurchaseorder {
	return &MsgUpdateSentPurchaseorder{
		Id:      id,
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgUpdateSentPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSentPurchaseorder) Type() string {
	return "UpdateSentPurchaseorder"
}

func (msg *MsgUpdateSentPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSentPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSentPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSentPurchaseorder{}

func NewMsgDeleteSentPurchaseorder(creator string, id uint64) *MsgDeleteSentPurchaseorder {
	return &MsgDeleteSentPurchaseorder{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSentPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSentPurchaseorder) Type() string {
	return "DeleteSentPurchaseorder"
}

func (msg *MsgDeleteSentPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSentPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSentPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
