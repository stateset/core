package types

import (
	"context"
)

// QueryServer defines the query server interface for the settlement module
type QueryServer interface {
	// Settlement returns a settlement by ID
	Settlement(ctx context.Context, req *QuerySettlementRequest) (*QuerySettlementResponse, error)
	// Settlements returns all settlements with pagination
	Settlements(ctx context.Context, req *QuerySettlementsRequest) (*QuerySettlementsResponse, error)
	// SettlementsByStatus returns settlements filtered by status
	SettlementsByStatus(ctx context.Context, req *QuerySettlementsByStatusRequest) (*QuerySettlementsByStatusResponse, error)
	// Batch returns a batch by ID
	Batch(ctx context.Context, req *QueryBatchRequest) (*QueryBatchResponse, error)
	// Batches returns all batches with pagination
	Batches(ctx context.Context, req *QueryBatchesRequest) (*QueryBatchesResponse, error)
	// Channel returns a payment channel by ID
	Channel(ctx context.Context, req *QueryChannelRequest) (*QueryChannelResponse, error)
	// Channels returns all payment channels with pagination
	Channels(ctx context.Context, req *QueryChannelsRequest) (*QueryChannelsResponse, error)
	// ChannelsByParty returns channels for a specific sender or recipient
	ChannelsByParty(ctx context.Context, req *QueryChannelsByPartyRequest) (*QueryChannelsByPartyResponse, error)
	// Merchant returns a merchant configuration by address
	Merchant(ctx context.Context, req *QueryMerchantRequest) (*QueryMerchantResponse, error)
	// Merchants returns all merchants with pagination
	Merchants(ctx context.Context, req *QueryMerchantsRequest) (*QueryMerchantsResponse, error)
	// Params returns the module parameters
	Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error)
}

// Query request/response types

type QuerySettlementRequest struct {
	Id uint64 `json:"id"`
}

type QuerySettlementResponse struct {
	Settlement Settlement `json:"settlement"`
}

type QuerySettlementsRequest struct {
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
}

type QuerySettlementsResponse struct {
	Settlements []Settlement `json:"settlements"`
	Total       uint64       `json:"total"`
}

type QuerySettlementsByStatusRequest struct {
	Status SettlementStatus `json:"status"`
	Offset uint64           `json:"offset,omitempty"`
	Limit  uint64           `json:"limit,omitempty"`
}

type QuerySettlementsByStatusResponse struct {
	Settlements []Settlement `json:"settlements"`
	Total       uint64       `json:"total"`
}

type QueryBatchRequest struct {
	Id uint64 `json:"id"`
}

type QueryBatchResponse struct {
	Batch BatchSettlement `json:"batch"`
}

type QueryBatchesRequest struct {
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
}

type QueryBatchesResponse struct {
	Batches []BatchSettlement `json:"batches"`
	Total   uint64            `json:"total"`
}

type QueryChannelRequest struct {
	Id uint64 `json:"id"`
}

type QueryChannelResponse struct {
	Channel PaymentChannel `json:"channel"`
}

type QueryChannelsRequest struct {
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
}

type QueryChannelsResponse struct {
	Channels []PaymentChannel `json:"channels"`
	Total    uint64           `json:"total"`
}

type QueryChannelsByPartyRequest struct {
	Address string `json:"address"`
	Offset  uint64 `json:"offset,omitempty"`
	Limit   uint64 `json:"limit,omitempty"`
}

type QueryChannelsByPartyResponse struct {
	Channels []PaymentChannel `json:"channels"`
	Total    uint64           `json:"total"`
}

type QueryMerchantRequest struct {
	Address string `json:"address"`
}

type QueryMerchantResponse struct {
	Merchant MerchantConfig `json:"merchant"`
}

type QueryMerchantsRequest struct {
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
}

type QueryMerchantsResponse struct {
	Merchants []MerchantConfig `json:"merchants"`
	Total     uint64           `json:"total"`
}

type QueryParamsRequest struct{}

type QueryParamsResponse struct {
	Params Params `json:"params"`
}

// RegisterQueryServer registers the query server
func RegisterQueryServer(s interface{}, srv QueryServer) {
	// Registration handled by module
}
