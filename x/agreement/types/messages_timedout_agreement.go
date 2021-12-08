package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTimedoutAgreement{}

func NewMsgCreateTimedoutAgreement(creator string, did string, chain string) *MsgCreateTimedoutAgreement {
	return &MsgCreateTimedoutAgreement{
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgCreateTimedoutAgreement) Route() string {
	return RouterKey
}

func (msg *MsgCreateTimedoutAgreement) Type() string {
	return "CreateTimedoutAgreement"
}

func (msg *MsgCreateTimedoutAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTimedoutAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTimedoutAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTimedoutAgreement{}

func NewMsgUpdateTimedoutAgreement(creator string, id uint64, did string, chain string) *MsgUpdateTimedoutAgreement {
	return &MsgUpdateTimedoutAgreement{
		Id:      id,
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgUpdateTimedoutAgreement) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTimedoutAgreement) Type() string {
	return "UpdateTimedoutAgreement"
}

func (msg *MsgUpdateTimedoutAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTimedoutAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTimedoutAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTimedoutAgreement{}

func NewMsgDeleteTimedoutAgreement(creator string, id uint64) *MsgDeleteTimedoutAgreement {
	return &MsgDeleteTimedoutAgreement{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTimedoutAgreement) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTimedoutAgreement) Type() string {
	return "DeleteTimedoutAgreement"
}

func (msg *MsgDeleteTimedoutAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTimedoutAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTimedoutAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
