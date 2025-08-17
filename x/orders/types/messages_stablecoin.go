package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Stablecoin payment message types

// MsgPayWithStablecoin defines the message for paying with stablecoins
type MsgPayWithStablecoin struct {
	Creator                 string      `json:"creator"`
	OrderId                 string      `json:"order_id"`
	StablecoinDenom         string      `json:"stablecoin_denom"`
	StablecoinAmount        sdk.Int     `json:"stablecoin_amount"`
	CustomerAddress         string      `json:"customer_address"`
	MerchantAddress         string      `json:"merchant_address"`
	ExchangeRate           sdk.Dec     `json:"exchange_rate"`
	UseEscrow              bool        `json:"use_escrow"`
	ConfirmationsRequired  uint64      `json:"confirmations_required"`
	EscrowTimeout          *time.Time  `json:"escrow_timeout,omitempty"`
}

func (msg *MsgPayWithStablecoin) Route() string {
	return RouterKey
}

func (msg *MsgPayWithStablecoin) Type() string {
	return "pay_with_stablecoin"
}

func (msg *MsgPayWithStablecoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPayWithStablecoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPayWithStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	if msg.StablecoinDenom == "" {
		return sdkerrors.Wrap(ErrInvalidStablecoin, "stablecoin denomination cannot be empty")
	}

	if msg.StablecoinAmount.IsNil() || msg.StablecoinAmount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAmount, "stablecoin amount must be positive")
	}

	_, err = sdk.AccAddressFromBech32(msg.CustomerAddress)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidCustomer, "invalid customer address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.MerchantAddress)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidMerchant, "invalid merchant address (%s)", err)
	}

	if msg.ExchangeRate.IsNil() || msg.ExchangeRate.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidAmount, "exchange rate must be positive")
	}

	return nil
}

// MsgPayWithStablecoinResponse defines the response for MsgPayWithStablecoin
type MsgPayWithStablecoinResponse struct {
	TxHash    string    `json:"tx_hash"`
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// MsgConfirmStablecoinPayment defines the message for confirming stablecoin payments
type MsgConfirmStablecoinPayment struct {
	Creator          string `json:"creator"`
	OrderId          string `json:"order_id"`
	ConfirmationCount uint64 `json:"confirmation_count"`
	BlockHeight      uint64 `json:"block_height"`
}

func (msg *MsgConfirmStablecoinPayment) Route() string {
	return RouterKey
}

func (msg *MsgConfirmStablecoinPayment) Type() string {
	return "confirm_stablecoin_payment"
}

func (msg *MsgConfirmStablecoinPayment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConfirmStablecoinPayment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfirmStablecoinPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	return nil
}

// MsgConfirmStablecoinPaymentResponse defines the response for MsgConfirmStablecoinPayment
type MsgConfirmStablecoinPaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// MsgRefundStablecoinPayment defines the message for refunding stablecoin payments
type MsgRefundStablecoinPayment struct {
	Creator         string  `json:"creator"`
	OrderId         string  `json:"order_id"`
	CustomerAddress string  `json:"customer_address"`
	RefundAmount    sdk.Int `json:"refund_amount"`
	Reason          string  `json:"reason"`
	PartialRefund   bool    `json:"partial_refund"`
}

func (msg *MsgRefundStablecoinPayment) Route() string {
	return RouterKey
}

func (msg *MsgRefundStablecoinPayment) Type() string {
	return "refund_stablecoin_payment"
}

func (msg *MsgRefundStablecoinPayment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRefundStablecoinPayment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRefundStablecoinPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	_, err = sdk.AccAddressFromBech32(msg.CustomerAddress)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidCustomer, "invalid customer address (%s)", err)
	}

	if msg.RefundAmount.IsNil() || msg.RefundAmount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAmount, "refund amount must be positive")
	}

	return nil
}

// MsgRefundStablecoinPaymentResponse defines the response for MsgRefundStablecoinPayment
type MsgRefundStablecoinPaymentResponse struct {
	TxHash  string `json:"tx_hash"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// MsgReleaseEscrow defines the message for releasing escrow
type MsgReleaseEscrow struct {
	Creator string `json:"creator"`
	OrderId string `json:"order_id"`
}

func (msg *MsgReleaseEscrow) Route() string {
	return RouterKey
}

func (msg *MsgReleaseEscrow) Type() string {
	return "release_escrow"
}

func (msg *MsgReleaseEscrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReleaseEscrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReleaseEscrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	return nil
}

// MsgReleaseEscrowResponse defines the response for MsgReleaseEscrow
type MsgReleaseEscrowResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}