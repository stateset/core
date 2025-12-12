package types

// Order status constants
type OrderStatus = string

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
type PaymentStatus = string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusEscrowed  PaymentStatus = "escrowed"
	PaymentStatusCaptured  PaymentStatus = "captured"
	PaymentStatusReleased  PaymentStatus = "released"
	PaymentStatusRefunded  PaymentStatus = "refunded"
	PaymentStatusFailed    PaymentStatus = "failed"
)

// Dispute status constants
type DisputeStatus = string

const (
	DisputeStatusOpen       DisputeStatus = "open"
	DisputeStatusUnderReview DisputeStatus = "under_review"
	DisputeStatusResolved   DisputeStatus = "resolved"
	DisputeStatusEscalated  DisputeStatus = "escalated"
)

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
