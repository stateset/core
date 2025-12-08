package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/payments/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the payments QueryServer interface
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

// Payment returns a payment by ID
func (q queryServer) Payment(goCtx context.Context, req *types.QueryPaymentRequest) (*types.QueryPaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	payment, found := q.Keeper.GetPayment(ctx, req.Id)
	if !found {
		return nil, types.ErrPaymentNotFound
	}

	return &types.QueryPaymentResponse{
		Payment: payment,
	}, nil
}

// Payments returns all payments with pagination
func (q queryServer) Payments(goCtx context.Context, req *types.QueryPaymentsRequest) (*types.QueryPaymentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	limit := req.Limit
	if limit == 0 || limit > 100 {
		limit = 100
	}
	offset := req.Offset

	var payments []types.PaymentIntent
	var total uint64

	q.Keeper.IteratePayments(ctx, func(p types.PaymentIntent) bool {
		if total >= offset && uint64(len(payments)) < limit {
			payments = append(payments, p)
		}
		total++
		return false
	})

	return &types.QueryPaymentsResponse{
		Payments: payments,
		Total:    total,
	}, nil
}

// PaymentsByPayer returns payments for a specific payer
func (q queryServer) PaymentsByPayer(goCtx context.Context, req *types.QueryPaymentsByPayerRequest) (*types.QueryPaymentsByPayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	limit := req.Limit
	if limit == 0 || limit > 100 {
		limit = 100
	}
	offset := req.Offset

	var payments []types.PaymentIntent
	var matched uint64

	q.Keeper.IteratePayments(ctx, func(p types.PaymentIntent) bool {
		if p.Payer == req.Payer {
			if matched >= offset && uint64(len(payments)) < limit {
				payments = append(payments, p)
			}
			matched++
		}
		return false
	})

	return &types.QueryPaymentsByPayerResponse{
		Payments: payments,
		Total:    matched,
	}, nil
}

// PaymentsByPayee returns payments for a specific payee
func (q queryServer) PaymentsByPayee(goCtx context.Context, req *types.QueryPaymentsByPayeeRequest) (*types.QueryPaymentsByPayeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	limit := req.Limit
	if limit == 0 || limit > 100 {
		limit = 100
	}
	offset := req.Offset

	var payments []types.PaymentIntent
	var matched uint64

	q.Keeper.IteratePayments(ctx, func(p types.PaymentIntent) bool {
		if p.Payee == req.Payee {
			if matched >= offset && uint64(len(payments)) < limit {
				payments = append(payments, p)
			}
			matched++
		}
		return false
	})

	return &types.QueryPaymentsByPayeeResponse{
		Payments: payments,
		Total:    matched,
	}, nil
}

// PaymentsByStatus returns payments filtered by status
func (q queryServer) PaymentsByStatus(goCtx context.Context, req *types.QueryPaymentsByStatusRequest) (*types.QueryPaymentsByStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	limit := req.Limit
	if limit == 0 || limit > 100 {
		limit = 100
	}
	offset := req.Offset

	var payments []types.PaymentIntent
	var matched uint64

	q.Keeper.IteratePayments(ctx, func(p types.PaymentIntent) bool {
		if p.Status == req.Status {
			if matched >= offset && uint64(len(payments)) < limit {
				payments = append(payments, p)
			}
			matched++
		}
		return false
	})

	return &types.QueryPaymentsByStatusResponse{
		Payments: payments,
		Total:    matched,
	}, nil
}
