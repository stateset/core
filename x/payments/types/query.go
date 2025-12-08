package types

import (
	"context"
)

// QueryServer defines the query server interface for the payments module
type QueryServer interface {
	// Payment returns a payment by ID
	Payment(ctx context.Context, req *QueryPaymentRequest) (*QueryPaymentResponse, error)
	// Payments returns all payments with pagination
	Payments(ctx context.Context, req *QueryPaymentsRequest) (*QueryPaymentsResponse, error)
	// PaymentsByPayer returns payments for a specific payer
	PaymentsByPayer(ctx context.Context, req *QueryPaymentsByPayerRequest) (*QueryPaymentsByPayerResponse, error)
	// PaymentsByPayee returns payments for a specific payee
	PaymentsByPayee(ctx context.Context, req *QueryPaymentsByPayeeRequest) (*QueryPaymentsByPayeeResponse, error)
	// PaymentsByStatus returns payments filtered by status
	PaymentsByStatus(ctx context.Context, req *QueryPaymentsByStatusRequest) (*QueryPaymentsByStatusResponse, error)
}

// Query request/response types

type QueryPaymentRequest struct {
	Id uint64 `json:"id"`
}

type QueryPaymentResponse struct {
	Payment PaymentIntent `json:"payment"`
}

type QueryPaymentsRequest struct {
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
}

type QueryPaymentsResponse struct {
	Payments []PaymentIntent `json:"payments"`
	Total    uint64          `json:"total"`
}

type QueryPaymentsByPayerRequest struct {
	Payer  string `json:"payer"`
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
}

type QueryPaymentsByPayerResponse struct {
	Payments []PaymentIntent `json:"payments"`
	Total    uint64          `json:"total"`
}

type QueryPaymentsByPayeeRequest struct {
	Payee  string `json:"payee"`
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
}

type QueryPaymentsByPayeeResponse struct {
	Payments []PaymentIntent `json:"payments"`
	Total    uint64          `json:"total"`
}

type QueryPaymentsByStatusRequest struct {
	Status PaymentStatus `json:"status"`
	Offset uint64        `json:"offset,omitempty"`
	Limit  uint64        `json:"limit,omitempty"`
}

type QueryPaymentsByStatusResponse struct {
	Payments []PaymentIntent `json:"payments"`
	Total    uint64          `json:"total"`
}

// RegisterQueryServer registers the query server
func RegisterQueryServer(s interface{}, srv QueryServer) {
	// Registration handled by module
}
