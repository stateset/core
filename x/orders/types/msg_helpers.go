package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	if !msg.RefundAmount.IsPositive() {
		return ErrInvalidAmount
	}
	return nil
}

func (msg MsgRefundOrder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Merchant)
	return []sdk.AccAddress{addr}
}

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
	return nil
}

func (msg MsgOpenDispute) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Customer)
	return []sdk.AccAddress{addr}
}

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
