package types

import (
	"context"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

// QueryServer is the server API for Query service
type QueryServer interface {
	// Price queries the current price for an asset
	Price(context.Context, *QueryPriceRequest) (*QueryPriceResponse, error)
	// PriceFeed queries a specific price feed
	PriceFeed(context.Context, *QueryPriceFeedRequest) (*QueryPriceFeedResponse, error)
	// Oracle queries a specific oracle provider
	Oracle(context.Context, *QueryOracleRequest) (*QueryOracleResponse, error)
	// Oracles queries all oracle providers
	Oracles(context.Context, *QueryOraclesRequest) (*QueryOraclesResponse, error)
	// PriceHistory queries price history for an asset
	PriceHistory(context.Context, *QueryPriceHistoryRequest) (*QueryPriceHistoryResponse, error)
	// Params queries the oracle parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// Query request and response types

type QueryPriceRequest struct {
	Asset string `json:"asset"`
}

type QueryPriceResponse struct {
	Price       sdk.Dec          `json:"price"`
	Aggregated  AggregatedPrice  `json:"aggregated"`
}

type QueryPriceFeedRequest struct {
	FeedId string `json:"feed_id"`
}

type QueryPriceFeedResponse struct {
	PriceFeed PriceFeed `json:"price_feed"`
}

type QueryOracleRequest struct {
	Address string `json:"address"`
}

type QueryOracleResponse struct {
	Oracle OracleProvider `json:"oracle"`
}

type QueryOraclesRequest struct {
	ActiveOnly bool `json:"active_only"`
}

type QueryOraclesResponse struct {
	Oracles []OracleProvider `json:"oracles"`
}

type QueryPriceHistoryRequest struct {
	Asset string `json:"asset"`
	Limit uint32 `json:"limit"`
}

type QueryPriceHistoryResponse struct {
	History PriceHistory `json:"history"`
}

type QueryParamsRequest struct{}

type QueryParamsResponse struct {
	Params Params `json:"params"`
}

// QueryClient is the client API for Query service
type QueryClient interface {
	Price(ctx context.Context, in *QueryPriceRequest, opts ...grpc.CallOption) (*QueryPriceResponse, error)
	PriceFeed(ctx context.Context, in *QueryPriceFeedRequest, opts ...grpc.CallOption) (*QueryPriceFeedResponse, error)
	Oracle(ctx context.Context, in *QueryOracleRequest, opts ...grpc.CallOption) (*QueryOracleResponse, error)
	Oracles(ctx context.Context, in *QueryOraclesRequest, opts ...grpc.CallOption) (*QueryOraclesResponse, error)
	PriceHistory(ctx context.Context, in *QueryPriceHistoryRequest, opts ...grpc.CallOption) (*QueryPriceHistoryResponse, error)
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

// NewQueryClient creates a new QueryClient
func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Price(ctx context.Context, in *QueryPriceRequest, opts ...grpc.CallOption) (*QueryPriceResponse, error) {
	out := new(QueryPriceResponse)
	err := c.cc.Invoke(ctx, "/stateset.oracle.v1.Query/Price", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PriceFeed(ctx context.Context, in *QueryPriceFeedRequest, opts ...grpc.CallOption) (*QueryPriceFeedResponse, error) {
	out := new(QueryPriceFeedResponse)
	err := c.cc.Invoke(ctx, "/stateset.oracle.v1.Query/PriceFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Oracle(ctx context.Context, in *QueryOracleRequest, opts ...grpc.CallOption) (*QueryOracleResponse, error) {
	out := new(QueryOracleResponse)
	err := c.cc.Invoke(ctx, "/stateset.oracle.v1.Query/Oracle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Oracles(ctx context.Context, in *QueryOraclesRequest, opts ...grpc.CallOption) (*QueryOraclesResponse, error) {
	out := new(QueryOraclesResponse)
	err := c.cc.Invoke(ctx, "/stateset.oracle.v1.Query/Oracles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PriceHistory(ctx context.Context, in *QueryPriceHistoryRequest, opts ...grpc.CallOption) (*QueryPriceHistoryResponse, error) {
	out := new(QueryPriceHistoryResponse)
	err := c.cc.Invoke(ctx, "/stateset.oracle.v1.Query/PriceHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/stateset.oracle.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterQueryHandlerClient registers the http handlers for service Query to "mux"
func RegisterQueryHandlerClient(ctx context.Context, mux *grpc.ServeMux, client QueryClient) error {
	// Implementation would be here in production
	return nil
}

// RegisterQueryServer registers the Query service
func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// Implementation would be here in production
}