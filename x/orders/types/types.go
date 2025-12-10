package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Order status constants
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
	OrderStatusDisputed   OrderStatus = "disputed"
)

// Payment status constants
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusEscrowed  PaymentStatus = "escrowed"
	PaymentStatusCaptured  PaymentStatus = "captured"
	PaymentStatusReleased  PaymentStatus = "released"
	PaymentStatusRefunded  PaymentStatus = "refunded"
	PaymentStatusFailed    PaymentStatus = "failed"
)

// Dispute status constants
type DisputeStatus string

const (
	DisputeStatusOpen       DisputeStatus = "open"
	DisputeStatusUnderReview DisputeStatus = "under_review"
	DisputeStatusResolved   DisputeStatus = "resolved"
	DisputeStatusEscalated  DisputeStatus = "escalated"
)

// Order represents a customer order.
type Order struct {
	Id              uint64         `json:"id"`
	Customer        string         `json:"customer"`
	Merchant        string         `json:"merchant"`
	Status          OrderStatus    `json:"status"`
	Items           []OrderItem    `json:"items"`
	Subtotal        sdk.Coin       `json:"subtotal"`
	ShippingCost    sdk.Coin       `json:"shipping_cost"`
	TaxAmount       sdk.Coin       `json:"tax_amount"`
	DiscountAmount  sdk.Coin       `json:"discount_amount"`
	TotalAmount     sdk.Coin       `json:"total_amount"`
	PaymentInfo     PaymentInfo    `json:"payment_info"`
	ShippingInfo    ShippingInfo   `json:"shipping_info"`
	Metadata        string         `json:"metadata"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	PaidAt          time.Time      `json:"paid_at,omitempty"`
	ShippedAt       time.Time      `json:"shipped_at,omitempty"`
	DeliveredAt     time.Time      `json:"delivered_at,omitempty"`
	CompletedAt     time.Time      `json:"completed_at,omitempty"`
	ExpiresAt       time.Time      `json:"expires_at"`
	SettlementId    uint64         `json:"settlement_id,omitempty"`
	DisputeId       uint64         `json:"dispute_id,omitempty"`
}

func (o *Order) Reset()         { *o = Order{} }
func (o *Order) String() string { return "Order" }
func (*Order) ProtoMessage()    {}

// OrderItem represents an individual item in an order.
type OrderItem struct {
	Id          string   `json:"id"`
	ProductId   string   `json:"product_id"`
	ProductName string   `json:"product_name"`
	Quantity    uint64   `json:"quantity"`
	UnitPrice   sdk.Coin `json:"unit_price"`
	TotalPrice  sdk.Coin `json:"total_price"`
	Variant     string   `json:"variant,omitempty"`
	Metadata    string   `json:"metadata,omitempty"`
}

// PaymentInfo contains payment details for an order.
type PaymentInfo struct {
	Status          PaymentStatus `json:"status"`
	Method          string        `json:"method"` // stablecoin, escrow, instant
	TransactionId   string        `json:"transaction_id,omitempty"`
	SettlementId    uint64        `json:"settlement_id,omitempty"`
	EscrowId        uint64        `json:"escrow_id,omitempty"`
	PaidAmount      sdk.Coin      `json:"paid_amount,omitempty"`
	RefundedAmount  sdk.Coin      `json:"refunded_amount,omitempty"`
	FeeAmount       sdk.Coin      `json:"fee_amount,omitempty"`
	PaidAt          time.Time     `json:"paid_at,omitempty"`
}

// ShippingInfo contains shipping details for an order.
type ShippingInfo struct {
	Address         Address   `json:"address"`
	Method          string    `json:"method"`
	Carrier         string    `json:"carrier,omitempty"`
	TrackingNumber  string    `json:"tracking_number,omitempty"`
	EstimatedDelivery time.Time `json:"estimated_delivery,omitempty"`
	ActualDelivery  time.Time `json:"actual_delivery,omitempty"`
}

// Address represents a shipping or billing address.
type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2,omitempty"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Name       string `json:"name"`
	Phone      string `json:"phone,omitempty"`
}

// Dispute represents a dispute on an order.
type Dispute struct {
	Id          uint64        `json:"id"`
	OrderId     uint64        `json:"order_id"`
	Customer    string        `json:"customer"`
	Merchant    string        `json:"merchant"`
	Reason      string        `json:"reason"`
	Description string        `json:"description"`
	Evidence    []string      `json:"evidence,omitempty"`
	Status      DisputeStatus `json:"status"`
	Resolution  string        `json:"resolution,omitempty"`
	ResolvedBy  string        `json:"resolved_by,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	ResolvedAt  time.Time     `json:"resolved_at,omitempty"`
	Amount      sdk.Coin      `json:"amount"`
}

func (d *Dispute) Reset()         { *d = Dispute{} }
func (d *Dispute) String() string { return "Dispute" }
func (*Dispute) ProtoMessage()    {}

// IsValidTransition checks if a status transition is valid.
func (o *Order) IsValidTransition(newStatus OrderStatus) bool {
	validTransitions := map[OrderStatus][]OrderStatus{
		OrderStatusPending:   {OrderStatusConfirmed, OrderStatusCancelled},
		OrderStatusConfirmed: {OrderStatusPaid, OrderStatusCancelled},
		OrderStatusPaid:      {OrderStatusShipped, OrderStatusRefunded, OrderStatusDisputed},
		OrderStatusShipped:   {OrderStatusDelivered, OrderStatusDisputed},
		OrderStatusDelivered: {OrderStatusCompleted, OrderStatusDisputed},
		OrderStatusDisputed:  {OrderStatusRefunded, OrderStatusCompleted},
		// Terminal states
		OrderStatusCompleted: {},
		OrderStatusCancelled: {},
		OrderStatusRefunded:  {},
	}

	allowed, exists := validTransitions[o.Status]
	if !exists {
		return false
	}

	for _, status := range allowed {
		if status == newStatus {
			return true
		}
	}
	return false
}

// CanBeRefunded checks if an order can be refunded.
func (o *Order) CanBeRefunded() bool {
	return o.Status == OrderStatusPaid ||
		o.Status == OrderStatusShipped ||
		o.Status == OrderStatusDelivered ||
		o.Status == OrderStatusDisputed
}

// CanBeDisputed checks if an order can have a dispute opened.
func (o *Order) CanBeDisputed() bool {
	return o.Status == OrderStatusPaid ||
		o.Status == OrderStatusShipped ||
		o.Status == OrderStatusDelivered
}
