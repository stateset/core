package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

var _ sdk.Msg = &MsgCreateTimedoutPurchaseorder{}

func NewMsgCreateTimedoutPurchaseorder(creator string, did string, chain string) *MsgCreateTimedoutPurchaseorder {
	return &MsgCreateTimedoutPurchaseorder{
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgCreateTimedoutPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgCreateTimedoutPurchaseorder) Type() string {
	return "CreateTimedoutPurchaseorder"
}

func (msg *MsgCreateTimedoutPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTimedoutPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTimedoutPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTimedoutPurchaseorder{}

func NewMsgUpdateTimedoutPurchaseorder(creator string, id uint64, did string, chain string) *MsgUpdateTimedoutPurchaseorder {
	return &MsgUpdateTimedoutPurchaseorder{
		Id:      id,
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgUpdateTimedoutPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTimedoutPurchaseorder) Type() string {
	return "UpdateTimedoutPurchaseorder"
}

func (msg *MsgUpdateTimedoutPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTimedoutPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTimedoutPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTimedoutPurchaseorder{}

func NewMsgDeleteTimedoutPurchaseorder(creator string, id uint64) *MsgDeleteTimedoutPurchaseorder {
	return &MsgDeleteTimedoutPurchaseorder{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTimedoutPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTimedoutPurchaseorder) Type() string {
	return "DeleteTimedoutPurchaseorder"
}

func (msg *MsgDeleteTimedoutPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTimedoutPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTimedoutPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(errorsmod.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
