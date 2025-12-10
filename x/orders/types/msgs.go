package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateOrder{}
var _ sdk.Msg = &MsgConfirmOrder{}
var _ sdk.Msg = &MsgPayOrder{}
var _ sdk.Msg = &MsgShipOrder{}
var _ sdk.Msg = &MsgDeliverOrder{}
var _ sdk.Msg = &MsgCompleteOrder{}
var _ sdk.Msg = &MsgCancelOrder{}
var _ sdk.Msg = &MsgRefundOrder{}
var _ sdk.Msg = &MsgOpenDispute{}
var _ sdk.Msg = &MsgResolveDispute{}

// MsgCreateOrder creates a new order.
type MsgCreateOrder struct {
	Customer     string      `json:"customer"`
	Merchant     string      `json:"merchant"`
	Items        []OrderItem `json:"items"`
	ShippingInfo ShippingInfo `json:"shipping_info"`
	Metadata     string      `json:"metadata,omitempty"`
}

func (m *MsgCreateOrder) Reset()         { *m = MsgCreateOrder{} }
func (m *MsgCreateOrder) String() string { return fmt.Sprintf("MsgCreateOrder{customer:%s,merchant:%s}", m.Customer, m.Merchant) }
func (*MsgCreateOrder) ProtoMessage()    {}

func NewMsgCreateOrder(customer, merchant string, items []OrderItem, shippingInfo ShippingInfo, metadata string) *MsgCreateOrder {
	return &MsgCreateOrder{
		Customer:     customer,
		Merchant:     merchant,
		Items:        items,
		ShippingInfo: shippingInfo,
		Metadata:     metadata,
	}
}

func (msg MsgCreateOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Customer); err != nil {
		return ErrInvalidCustomer
	}
	if _, err := sdk.AccAddressFromBech32(msg.Merchant); err != nil {
		return ErrInvalidMerchant
	}
	if len(msg.Items) == 0 {
		return ErrEmptyItems
	}
	return nil
}

func (msg MsgCreateOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Customer)
	return []sdk.AccAddress{addr}
}

// MsgConfirmOrder confirms an order by the merchant.
type MsgConfirmOrder struct {
	Merchant string `json:"merchant"`
	OrderId  uint64 `json:"order_id"`
}

func (m *MsgConfirmOrder) Reset()         { *m = MsgConfirmOrder{} }
func (m *MsgConfirmOrder) String() string { return fmt.Sprintf("MsgConfirmOrder{merchant:%s,order:%d}", m.Merchant, m.OrderId) }
func (*MsgConfirmOrder) ProtoMessage()    {}

func NewMsgConfirmOrder(merchant string, orderId uint64) *MsgConfirmOrder {
	return &MsgConfirmOrder{
		Merchant: merchant,
		OrderId:  orderId,
	}
}

func (msg MsgConfirmOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Merchant); err != nil {
		return ErrInvalidMerchant
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	return nil
}

func (msg MsgConfirmOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Merchant)
	return []sdk.AccAddress{addr}
}

// MsgPayOrder pays for an order using stablecoin.
type MsgPayOrder struct {
	Customer string   `json:"customer"`
	OrderId  uint64   `json:"order_id"`
	Amount   sdk.Coin `json:"amount"`
	UseEscrow bool    `json:"use_escrow"`
}

func (m *MsgPayOrder) Reset()         { *m = MsgPayOrder{} }
func (m *MsgPayOrder) String() string { return fmt.Sprintf("MsgPayOrder{customer:%s,order:%d,amount:%s}", m.Customer, m.OrderId, m.Amount.String()) }
func (*MsgPayOrder) ProtoMessage()    {}

func NewMsgPayOrder(customer string, orderId uint64, amount sdk.Coin, useEscrow bool) *MsgPayOrder {
	return &MsgPayOrder{
		Customer:  customer,
		OrderId:   orderId,
		Amount:    amount,
		UseEscrow: useEscrow,
	}
}

func (msg MsgPayOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Customer); err != nil {
		return ErrInvalidCustomer
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	if !msg.Amount.IsPositive() {
		return ErrInvalidAmount
	}
	return nil
}

func (msg MsgPayOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Customer)
	return []sdk.AccAddress{addr}
}

// MsgShipOrder marks an order as shipped.
type MsgShipOrder struct {
	Merchant       string `json:"merchant"`
	OrderId        uint64 `json:"order_id"`
	Carrier        string `json:"carrier"`
	TrackingNumber string `json:"tracking_number"`
}

func (m *MsgShipOrder) Reset()         { *m = MsgShipOrder{} }
func (m *MsgShipOrder) String() string { return fmt.Sprintf("MsgShipOrder{merchant:%s,order:%d}", m.Merchant, m.OrderId) }
func (*MsgShipOrder) ProtoMessage()    {}

func NewMsgShipOrder(merchant string, orderId uint64, carrier, trackingNumber string) *MsgShipOrder {
	return &MsgShipOrder{
		Merchant:       merchant,
		OrderId:        orderId,
		Carrier:        carrier,
		TrackingNumber: trackingNumber,
	}
}

func (msg MsgShipOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Merchant); err != nil {
		return ErrInvalidMerchant
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	return nil
}

func (msg MsgShipOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Merchant)
	return []sdk.AccAddress{addr}
}

// MsgDeliverOrder marks an order as delivered (can be triggered by customer or delivery confirmation).
type MsgDeliverOrder struct {
	Signer  string `json:"signer"`
	OrderId uint64 `json:"order_id"`
}

func (m *MsgDeliverOrder) Reset()         { *m = MsgDeliverOrder{} }
func (m *MsgDeliverOrder) String() string { return fmt.Sprintf("MsgDeliverOrder{signer:%s,order:%d}", m.Signer, m.OrderId) }
func (*MsgDeliverOrder) ProtoMessage()    {}

func NewMsgDeliverOrder(signer string, orderId uint64) *MsgDeliverOrder {
	return &MsgDeliverOrder{
		Signer:  signer,
		OrderId: orderId,
	}
}

func (msg MsgDeliverOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return ErrUnauthorized
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	return nil
}

func (msg MsgDeliverOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// MsgCompleteOrder marks an order as completed (auto-releases escrow if applicable).
type MsgCompleteOrder struct {
	Customer string `json:"customer"`
	OrderId  uint64 `json:"order_id"`
}

func (m *MsgCompleteOrder) Reset()         { *m = MsgCompleteOrder{} }
func (m *MsgCompleteOrder) String() string { return fmt.Sprintf("MsgCompleteOrder{customer:%s,order:%d}", m.Customer, m.OrderId) }
func (*MsgCompleteOrder) ProtoMessage()    {}

func NewMsgCompleteOrder(customer string, orderId uint64) *MsgCompleteOrder {
	return &MsgCompleteOrder{
		Customer: customer,
		OrderId:  orderId,
	}
}

func (msg MsgCompleteOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Customer); err != nil {
		return ErrInvalidCustomer
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	return nil
}

func (msg MsgCompleteOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Customer)
	return []sdk.AccAddress{addr}
}

// MsgCancelOrder cancels an order before payment.
type MsgCancelOrder struct {
	Signer  string `json:"signer"`
	OrderId uint64 `json:"order_id"`
	Reason  string `json:"reason"`
}

func (m *MsgCancelOrder) Reset()         { *m = MsgCancelOrder{} }
func (m *MsgCancelOrder) String() string { return fmt.Sprintf("MsgCancelOrder{signer:%s,order:%d}", m.Signer, m.OrderId) }
func (*MsgCancelOrder) ProtoMessage()    {}

func NewMsgCancelOrder(signer string, orderId uint64, reason string) *MsgCancelOrder {
	return &MsgCancelOrder{
		Signer:  signer,
		OrderId: orderId,
		Reason:  reason,
	}
}

func (msg MsgCancelOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return ErrUnauthorized
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	return nil
}

func (msg MsgCancelOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// MsgRefundOrder refunds an order (by merchant or via dispute resolution).
type MsgRefundOrder struct {
	Merchant     string   `json:"merchant"`
	OrderId      uint64   `json:"order_id"`
	RefundAmount sdk.Coin `json:"refund_amount"`
	Reason       string   `json:"reason"`
	FullRefund   bool     `json:"full_refund"`
}

func (m *MsgRefundOrder) Reset()         { *m = MsgRefundOrder{} }
func (m *MsgRefundOrder) String() string { return fmt.Sprintf("MsgRefundOrder{merchant:%s,order:%d}", m.Merchant, m.OrderId) }
func (*MsgRefundOrder) ProtoMessage()    {}

func NewMsgRefundOrder(merchant string, orderId uint64, refundAmount sdk.Coin, reason string, fullRefund bool) *MsgRefundOrder {
	return &MsgRefundOrder{
		Merchant:     merchant,
		OrderId:      orderId,
		RefundAmount: refundAmount,
		Reason:       reason,
		FullRefund:   fullRefund,
	}
}

func (msg MsgRefundOrder) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Merchant); err != nil {
		return ErrInvalidMerchant
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	return nil
}

func (msg MsgRefundOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Merchant)
	return []sdk.AccAddress{addr}
}

// MsgOpenDispute opens a dispute on an order.
type MsgOpenDispute struct {
	Customer    string   `json:"customer"`
	OrderId     uint64   `json:"order_id"`
	Reason      string   `json:"reason"`
	Description string   `json:"description"`
	Evidence    []string `json:"evidence,omitempty"`
}

func (m *MsgOpenDispute) Reset()         { *m = MsgOpenDispute{} }
func (m *MsgOpenDispute) String() string { return fmt.Sprintf("MsgOpenDispute{customer:%s,order:%d}", m.Customer, m.OrderId) }
func (*MsgOpenDispute) ProtoMessage()    {}

func NewMsgOpenDispute(customer string, orderId uint64, reason, description string, evidence []string) *MsgOpenDispute {
	return &MsgOpenDispute{
		Customer:    customer,
		OrderId:     orderId,
		Reason:      reason,
		Description: description,
		Evidence:    evidence,
	}
}

func (msg MsgOpenDispute) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Customer); err != nil {
		return ErrInvalidCustomer
	}
	if msg.OrderId == 0 {
		return ErrInvalidOrder
	}
	if msg.Reason == "" {
		return ErrInvalidOrder
	}
	return nil
}

func (msg MsgOpenDispute) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Customer)
	return []sdk.AccAddress{addr}
}

// MsgResolveDispute resolves a dispute (by authority/arbitrator).
type MsgResolveDispute struct {
	Authority    string   `json:"authority"`
	DisputeId    uint64   `json:"dispute_id"`
	Resolution   string   `json:"resolution"`
	RefundAmount sdk.Coin `json:"refund_amount"`
	ToCustomer   bool     `json:"to_customer"`
}

func (m *MsgResolveDispute) Reset()         { *m = MsgResolveDispute{} }
func (m *MsgResolveDispute) String() string { return fmt.Sprintf("MsgResolveDispute{authority:%s,dispute:%d}", m.Authority, m.DisputeId) }
func (*MsgResolveDispute) ProtoMessage()    {}

func NewMsgResolveDispute(authority string, disputeId uint64, resolution string, refundAmount sdk.Coin, toCustomer bool) *MsgResolveDispute {
	return &MsgResolveDispute{
		Authority:    authority,
		DisputeId:    disputeId,
		Resolution:   resolution,
		RefundAmount: refundAmount,
		ToCustomer:   toCustomer,
	}
}

func (msg MsgResolveDispute) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrUnauthorized
	}
	if msg.DisputeId == 0 {
		return ErrDisputeNotFound
	}
	return nil
}

func (msg MsgResolveDispute) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}
