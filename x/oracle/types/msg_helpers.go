package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMsgUpdatePrice creates a new MsgUpdatePrice instance.
func NewMsgUpdatePrice(authority, denom string, price sdkmath.LegacyDec) *MsgUpdatePrice {
	return &MsgUpdatePrice{
		Authority: authority,
		Denom:     denom,
		Price:     price,
	}
}

func (m MsgUpdatePrice) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrapf(ErrInvalidAuthority, "authority %s", err)
	}
	if len(m.Denom) == 0 {
		return errorsmod.Wrap(ErrInvalidDenom, "denom cannot be empty")
	}
	if !m.Price.IsPositive() {
		return errorsmod.Wrap(ErrInvalidPrice, "price must be positive")
	}
	return nil
}

func (m MsgUpdatePrice) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
