package types

import "context"

// QueryServer defines the query server interface for the orders module.
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	Order(context.Context, *QueryOrderRequest) (*QueryOrderResponse, error)
	Orders(context.Context, *QueryOrdersRequest) (*QueryOrdersResponse, error)
}

type QueryParamsRequest struct{}

type QueryParamsResponse struct {
	Params Params `json:"params"`
}

type QueryOrderRequest struct {
	Id uint64 `json:"id"`
}

type QueryOrderResponse struct {
	Order Order `json:"order"`
}

type QueryOrdersRequest struct {
	Customer string `json:"customer,omitempty"`
	Merchant string `json:"merchant,omitempty"`
	Status   string `json:"status,omitempty"`
	Offset   uint64 `json:"offset,omitempty"`
	Limit    uint64 `json:"limit,omitempty"`
}

type QueryOrdersResponse struct {
	Orders []Order `json:"orders"`
	Total  uint64  `json:"total"`
}
