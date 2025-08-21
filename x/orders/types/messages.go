package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
)

const (
	TypeMsgCreateOrder       = "create_order"
	TypeMsgUpdateOrder       = "update_order"
	TypeMsgCancelOrder       = "cancel_order"
	TypeMsgFulfillOrder      = "fulfill_order"
	TypeMsgRefundOrder       = "refund_order"
	TypeMsgUpdateOrderStatus = "update_order_status"
)

type MsgCreateOrder struct {
	Creator         string        `json:"creator"`
	Customer        string        `json:"customer"`
	Merchant        string        `json:"merchant"`
	TotalAmount     sdk.Int       `json:"total_amount"`
	Currency        string        `json:"currency"`
	Items           []OrderItem   `json:"items"`
	ShippingInfo    *ShippingInfo `json:"shipping_info,omitempty"`
	PaymentInfo     *PaymentInfo  `json:"payment_info,omitempty"`
	Metadata        string        `json:"metadata"`
	DueDate         *time.Time    `json:"due_date,omitempty"`
	FulfillmentType string        `json:"fulfillment_type"`
	Source          string        `json:"source"`
	Discounts       []Discount    `json:"discounts"`
	TaxInfo         *TaxInfo      `json:"tax_info,omitempty"`
}

func (msg *MsgCreateOrder) Route() string {
	return RouterKey
}

func (msg *MsgCreateOrder) Type() string {
	return TypeMsgCreateOrder
}

func (msg *MsgCreateOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Items) == 0 {
		return sdkerrors.Wrap(ErrEmptyOrderItems, "order must have at least one item")
	}

	if msg.TotalAmount.IsNil() || msg.TotalAmount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAmount, "total amount must be positive")
	}

	return nil
}

type MsgCreateOrderResponse struct {
	OrderId string `json:"order_id"`
}

type MsgUpdateOrder struct {
	Creator      string        `json:"creator"`
	OrderId      string        `json:"order_id"`
	Items        []OrderItem   `json:"items,omitempty"`
	ShippingInfo *ShippingInfo `json:"shipping_info,omitempty"`
	PaymentInfo  *PaymentInfo  `json:"payment_info,omitempty"`
	Metadata     string        `json:"metadata"`
	DueDate      *time.Time    `json:"due_date,omitempty"`
}

func (msg *MsgUpdateOrder) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOrder) Type() string {
	return TypeMsgUpdateOrder
}

func (msg *MsgUpdateOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	return nil
}

type MsgUpdateOrderResponse struct{}

type MsgCancelOrder struct {
	Creator string `json:"creator"`
	OrderId string `json:"order_id"`
	Reason  string `json:"reason"`
}

func (msg *MsgCancelOrder) Route() string {
	return RouterKey
}

func (msg *MsgCancelOrder) Type() string {
	return TypeMsgCancelOrder
}

func (msg *MsgCancelOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	return nil
}

type MsgCancelOrderResponse struct{}

type MsgFulfillOrder struct {
	Creator        string `json:"creator"`
	OrderId        string `json:"order_id"`
	TrackingNumber string `json:"tracking_number"`
	Carrier        string `json:"carrier"`
}

func (msg *MsgFulfillOrder) Route() string {
	return RouterKey
}

func (msg *MsgFulfillOrder) Type() string {
	return TypeMsgFulfillOrder
}

func (msg *MsgFulfillOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFulfillOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFulfillOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	return nil
}

type MsgFulfillOrderResponse struct{}

type MsgRefundOrder struct {
	Creator       string  `json:"creator"`
	OrderId       string  `json:"order_id"`
	RefundAmount  sdk.Int `json:"refund_amount"`
	Reason        string  `json:"reason"`
	PartialRefund bool    `json:"partial_refund"`
}

func (msg *MsgRefundOrder) Route() string {
	return RouterKey
}

func (msg *MsgRefundOrder) Type() string {
	return TypeMsgRefundOrder
}

func (msg *MsgRefundOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRefundOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRefundOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	if msg.RefundAmount.IsNil() || msg.RefundAmount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAmount, "refund amount must be positive")
	}

	return nil
}

type MsgRefundOrderResponse struct{}

type MsgUpdateOrderStatus struct {
	Creator string `json:"creator"`
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
	Notes   string `json:"notes"`
}

func (msg *MsgUpdateOrderStatus) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOrderStatus) Type() string {
	return TypeMsgUpdateOrderStatus
}

func (msg *MsgUpdateOrderStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOrderStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOrderStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.OrderId == "" {
		return sdkerrors.Wrap(ErrInvalidOrderID, "order ID cannot be empty")
	}

	if msg.Status == "" {
		return sdkerrors.Wrap(ErrInvalidStatus, "status cannot be empty")
	}

	return nil
}

type MsgUpdateOrderStatusResponse struct{}

type MsgServer interface {
	CreateOrder(context.Context, *MsgCreateOrder) (*MsgCreateOrderResponse, error)
	UpdateOrder(context.Context, *MsgUpdateOrder) (*MsgUpdateOrderResponse, error)
	CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error)
	FulfillOrder(context.Context, *MsgFulfillOrder) (*MsgFulfillOrderResponse, error)
	RefundOrder(context.Context, *MsgRefundOrder) (*MsgRefundOrderResponse, error)
	UpdateOrderStatus(context.Context, *MsgUpdateOrderStatus) (*MsgUpdateOrderStatusResponse, error)
	PayWithStablecoin(context.Context, *MsgPayWithStablecoin) (*MsgPayWithStablecoinResponse, error)
	ConfirmStablecoinPayment(context.Context, *MsgConfirmStablecoinPayment) (*MsgConfirmStablecoinPaymentResponse, error)
	RefundStablecoinPayment(context.Context, *MsgRefundStablecoinPayment) (*MsgRefundStablecoinPaymentResponse, error)
	ReleaseEscrow(context.Context, *MsgReleaseEscrow) (*MsgReleaseEscrowResponse, error)
}

var _Msg_serviceDesc = sdk.ServiceDesc{
	ServiceName: "stateset.orders.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []sdk.Method{},
	Streams:     []sdk.Stream{},
}