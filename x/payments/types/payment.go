package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PaymentStatus enumerates the lifecycle states.
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSettled   PaymentStatus = "settled"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

// PaymentIntent holds metadata for pending settlements.
type PaymentIntent struct {
	Id            uint64        `json:"id" yaml:"id"`
	Payer         string        `json:"payer" yaml:"payer"`
	Payee         string        `json:"payee" yaml:"payee"`
	Amount        sdk.Coin      `json:"amount" yaml:"amount"`
	Status        PaymentStatus `json:"status" yaml:"status"`
	Metadata      string        `json:"metadata" yaml:"metadata"`
	CreatedHeight int64         `json:"created_height" yaml:"created_height"`
	CreatedTime   time.Time     `json:"created_time" yaml:"created_time"`
	SettledHeight int64         `json:"settled_height" yaml:"settled_height"`
	SettledTime   time.Time     `json:"settled_time" yaml:"settled_time"`
}

func (p PaymentIntent) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(p.Payer); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(p.Payee); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if !p.Amount.IsValid() || p.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidPayment, "amount must be positive")
	}
	return nil
}
