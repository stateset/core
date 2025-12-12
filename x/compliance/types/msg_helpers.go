package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgUpsertProfile(authority string, profile Profile) *MsgUpsertProfile {
	return &MsgUpsertProfile{
		Authority: authority,
		Profile:   profile,
	}
}

func (m MsgUpsertProfile) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidAddress, fmt.Sprintf("authority: %v", err))
	}
	return m.Profile.ValidateBasic()
}

func (m MsgUpsertProfile) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgSetSanction(authority, address string, sanction bool, reason string) *MsgSetSanction {
	return &MsgSetSanction{
		Authority: authority,
		Address:   address,
		Sanction:  sanction,
		Reason:    reason,
	}
}

func (m MsgSetSanction) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidAddress, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errorsmod.Wrap(ErrInvalidAddress, err.Error())
	}
	return nil
}

func (m MsgSetSanction) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
