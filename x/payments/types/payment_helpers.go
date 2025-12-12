package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PaymentStatus enumerates the lifecycle states.
type PaymentStatus = string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSettled   PaymentStatus = "settled"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

func (p PaymentIntent) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(p.Payer); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(p.Payee); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if p.Payer == p.Payee {
		return errorsmod.Wrap(ErrInvalidPayment, "payer and payee cannot be the same")
	}
	if !p.Amount.IsValid() || p.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidPayment, "amount must be positive")
	}
	return nil
}
