package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

var _ sdk.Msg = &MsgCreateSentAgreement{}

func NewMsgCreateSentAgreement(creator string, did string, chain string) *MsgCreateSentAgreement {
	return &MsgCreateSentAgreement{
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgCreateSentAgreement) Route() string {
	return RouterKey
}

func (msg *MsgCreateSentAgreement) Type() string {
	return "CreateSentAgreement"
}

func (msg *MsgCreateSentAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSentAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSentAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSentAgreement{}

func NewMsgUpdateSentAgreement(creator string, id uint64, did string, chain string) *MsgUpdateSentAgreement {
	return &MsgUpdateSentAgreement{
		Id:      id,
		Creator: creator,
		Did:     did,
		Chain:   chain,
	}
}

func (msg *MsgUpdateSentAgreement) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSentAgreement) Type() string {
	return "UpdateSentAgreement"
}

func (msg *MsgUpdateSentAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSentAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSentAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSentAgreement{}

func NewMsgDeleteSentAgreement(creator string, id uint64) *MsgDeleteSentAgreement {
	return &MsgDeleteSentAgreement{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSentAgreement) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSentAgreement) Type() string {
	return "DeleteSentAgreement"
}

func (msg *MsgDeleteSentAgreement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSentAgreement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSentAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
