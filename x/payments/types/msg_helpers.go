package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgCreatePayment(payer, payee string, amount sdk.Coin, metadata string) *MsgCreatePayment {
	return &MsgCreatePayment{Payer: payer, Payee: payee, Amount: amount, Metadata: metadata}
}

func (m MsgCreatePayment) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Payer); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(m.Payee); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidPayment, "amount must be positive")
	}
	return nil
}

func (m MsgCreatePayment) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Payer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgSettlePayment(payee string, paymentID uint64) *MsgSettlePayment {
	return &MsgSettlePayment{Payee: payee, PaymentId: paymentID}
}

func (m MsgSettlePayment) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Payee); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if m.PaymentId == 0 {
		return errorsmod.Wrap(ErrInvalidPayment, "payment id required")
	}
	return nil
}

func (m MsgSettlePayment) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Payee)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgCancelPayment(payer string, paymentID uint64, reason string) *MsgCancelPayment {
	return &MsgCancelPayment{Payer: payer, PaymentId: paymentID, Reason: reason}
}

func (m MsgCancelPayment) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Payer); err != nil {
		return errorsmod.Wrap(ErrInvalidPayment, err.Error())
	}
	if m.PaymentId == 0 {
		return errorsmod.Wrap(ErrInvalidPayment, "payment id required")
	}
	return nil
}

func (m MsgCancelPayment) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Payer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
