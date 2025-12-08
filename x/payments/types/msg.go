package types

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgCreatePayment = "create_payment"
	TypeMsgSettlePayment = "settle_payment"
	TypeMsgCancelPayment = "cancel_payment"
)

var (
	_ sdk.Msg = (*MsgCreatePayment)(nil)
	_ sdk.Msg = (*MsgSettlePayment)(nil)
	_ sdk.Msg = (*MsgCancelPayment)(nil)
)

type MsgCreatePayment struct {
	Payer    string   `json:"payer" yaml:"payer"`
	Payee    string   `json:"payee" yaml:"payee"`
	Amount   sdk.Coin `json:"amount" yaml:"amount"`
	Metadata string   `json:"metadata" yaml:"metadata"`
}

func (m *MsgCreatePayment) Reset() { *m = MsgCreatePayment{} }
func (m *MsgCreatePayment) String() string {
	return fmt.Sprintf("MsgCreatePayment{%s->%s %s}", m.Payer, m.Payee, m.Amount.String())
}
func (*MsgCreatePayment) ProtoMessage() {}

func NewMsgCreatePayment(payer, payee string, amount sdk.Coin, metadata string) *MsgCreatePayment {
	return &MsgCreatePayment{Payer: payer, Payee: payee, Amount: amount, Metadata: metadata}
}

func (m MsgCreatePayment) Route() string { return RouterKey }
func (m MsgCreatePayment) Type() string  { return TypeMsgCreatePayment }

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

func (m MsgCreatePayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCreatePaymentResponse struct {
	PaymentId uint64 `json:"payment_id"`
}

type MsgSettlePayment struct {
	Payee     string `json:"payee" yaml:"payee"`
	PaymentId uint64 `json:"payment_id" yaml:"payment_id"`
}

func (m *MsgSettlePayment) Reset() { *m = MsgSettlePayment{} }
func (m *MsgSettlePayment) String() string {
	return fmt.Sprintf("MsgSettlePayment{%s %d}", m.Payee, m.PaymentId)
}
func (*MsgSettlePayment) ProtoMessage() {}

func NewMsgSettlePayment(payee string, paymentID uint64) *MsgSettlePayment {
	return &MsgSettlePayment{Payee: payee, PaymentId: paymentID}
}

func (m MsgSettlePayment) Route() string { return RouterKey }
func (m MsgSettlePayment) Type() string  { return TypeMsgSettlePayment }

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

func (m MsgSettlePayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgSettlePaymentResponse struct{}

type MsgCancelPayment struct {
	Payer     string `json:"payer" yaml:"payer"`
	PaymentId uint64 `json:"payment_id" yaml:"payment_id"`
	Reason    string `json:"reason" yaml:"reason"`
}

func (m *MsgCancelPayment) Reset() { *m = MsgCancelPayment{} }
func (m *MsgCancelPayment) String() string {
	return fmt.Sprintf("MsgCancelPayment{%s %d}", m.Payer, m.PaymentId)
}
func (*MsgCancelPayment) ProtoMessage() {}

func NewMsgCancelPayment(payer string, paymentID uint64, reason string) *MsgCancelPayment {
	return &MsgCancelPayment{Payer: payer, PaymentId: paymentID, Reason: reason}
}

func (m MsgCancelPayment) Route() string { return RouterKey }
func (m MsgCancelPayment) Type() string  { return TypeMsgCancelPayment }

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

func (m MsgCancelPayment) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCancelPaymentResponse struct{}

type MsgServer interface {
	CreatePayment(ctx context.Context, msg *MsgCreatePayment) (*MsgCreatePaymentResponse, error)
	SettlePayment(ctx context.Context, msg *MsgSettlePayment) (*MsgSettlePaymentResponse, error)
	CancelPayment(ctx context.Context, msg *MsgCancelPayment) (*MsgCancelPaymentResponse, error)
}
