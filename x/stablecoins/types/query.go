package types

import (
	"context"
	
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc"
)

// Query request and response types

type QueryParamsRequest struct{}

func (*QueryParamsRequest) ProtoMessage() {}
func (m *QueryParamsRequest) Reset() { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return "" }

type QueryParamsResponse struct {
	Params Params `json:"params"`
}

func (*QueryParamsResponse) ProtoMessage() {}
func (m *QueryParamsResponse) Reset() { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return "" }

type QueryStablecoinRequest struct {
	Denom string `json:"denom"`
}

func (*QueryStablecoinRequest) ProtoMessage() {}
func (m *QueryStablecoinRequest) Reset() { *m = QueryStablecoinRequest{} }
func (m *QueryStablecoinRequest) String() string { return "" }

type QueryStablecoinResponse struct {
	Stablecoin Stablecoin `json:"stablecoin"`
}

func (*QueryStablecoinResponse) ProtoMessage() {}
func (m *QueryStablecoinResponse) Reset() { *m = QueryStablecoinResponse{} }
func (m *QueryStablecoinResponse) String() string { return "" }

type QueryStablecoinsRequest struct {
	Pagination *query.PageRequest `json:"pagination,omitempty"`
}

func (*QueryStablecoinsRequest) ProtoMessage() {}
func (m *QueryStablecoinsRequest) Reset() { *m = QueryStablecoinsRequest{} }
func (m *QueryStablecoinsRequest) String() string { return "" }

type QueryStablecoinsResponse struct {
	Stablecoins []Stablecoin        `json:"stablecoins"`
	Pagination  *query.PageResponse `json:"pagination,omitempty"`
}

func (*QueryStablecoinsResponse) ProtoMessage() {}
func (m *QueryStablecoinsResponse) Reset() { *m = QueryStablecoinsResponse{} }
func (m *QueryStablecoinsResponse) String() string { return "" }

type QueryStablecoinsByIssuerRequest struct {
	Issuer     string              `json:"issuer"`
	Pagination *query.PageRequest `json:"pagination,omitempty"`
}

func (*QueryStablecoinsByIssuerRequest) ProtoMessage() {}
func (m *QueryStablecoinsByIssuerRequest) Reset() { *m = QueryStablecoinsByIssuerRequest{} }
func (m *QueryStablecoinsByIssuerRequest) String() string { return "" }

type QueryStablecoinsByIssuerResponse struct {
	Stablecoins []Stablecoin        `json:"stablecoins"`
	Pagination  *query.PageResponse `json:"pagination,omitempty"`
}

func (*QueryStablecoinsByIssuerResponse) ProtoMessage() {}
func (m *QueryStablecoinsByIssuerResponse) Reset() { *m = QueryStablecoinsByIssuerResponse{} }
func (m *QueryStablecoinsByIssuerResponse) String() string { return "" }

type QueryStablecoinSupplyRequest struct {
	Denom string `json:"denom"`
}

func (*QueryStablecoinSupplyRequest) ProtoMessage() {}
func (m *QueryStablecoinSupplyRequest) Reset() { *m = QueryStablecoinSupplyRequest{} }
func (m *QueryStablecoinSupplyRequest) String() string { return "" }

type QueryStablecoinSupplyResponse struct {
	Supply StablecoinSupply `json:"supply"`
}

func (*QueryStablecoinSupplyResponse) ProtoMessage() {}
func (m *QueryStablecoinSupplyResponse) Reset() { *m = QueryStablecoinSupplyResponse{} }
func (m *QueryStablecoinSupplyResponse) String() string { return "" }

// Additional query types

type QueryPriceDataRequest struct {
	Denom string `json:"denom"`
}

func (*QueryPriceDataRequest) ProtoMessage() {}
func (m *QueryPriceDataRequest) Reset() { *m = QueryPriceDataRequest{} }
func (m *QueryPriceDataRequest) String() string { return "" }

type QueryPriceDataResponse struct {
	PriceData []PriceData `json:"price_data"`
}

func (*QueryPriceDataResponse) ProtoMessage() {}
func (m *QueryPriceDataResponse) Reset() { *m = QueryPriceDataResponse{} }
func (m *QueryPriceDataResponse) String() string { return "" }

type QueryReserveInfoRequest struct {
	Denom string `json:"denom"`
}

func (*QueryReserveInfoRequest) ProtoMessage() {}
func (m *QueryReserveInfoRequest) Reset() { *m = QueryReserveInfoRequest{} }
func (m *QueryReserveInfoRequest) String() string { return "" }

type QueryReserveInfoResponse struct {
	ReserveInfo *ReserveInfo `json:"reserve_info"`
}

func (*QueryReserveInfoResponse) ProtoMessage() {}
func (m *QueryReserveInfoResponse) Reset() { *m = QueryReserveInfoResponse{} }
func (m *QueryReserveInfoResponse) String() string { return "" }

type QueryIsWhitelistedRequest struct {
	Denom   string `json:"denom"`
	Address string `json:"address"`
}

func (*QueryIsWhitelistedRequest) ProtoMessage() {}
func (m *QueryIsWhitelistedRequest) Reset() { *m = QueryIsWhitelistedRequest{} }
func (m *QueryIsWhitelistedRequest) String() string { return "" }

type QueryIsWhitelistedResponse struct {
	Whitelisted bool `json:"whitelisted"`
}

func (*QueryIsWhitelistedResponse) ProtoMessage() {}
func (m *QueryIsWhitelistedResponse) Reset() { *m = QueryIsWhitelistedResponse{} }
func (m *QueryIsWhitelistedResponse) String() string { return "" }

type QueryIsBlacklistedRequest struct {
	Denom   string `json:"denom"`
	Address string `json:"address"`
}

func (*QueryIsBlacklistedRequest) ProtoMessage() {}
func (m *QueryIsBlacklistedRequest) Reset() { *m = QueryIsBlacklistedRequest{} }
func (m *QueryIsBlacklistedRequest) String() string { return "" }

type QueryIsBlacklistedResponse struct {
	Blacklisted bool `json:"blacklisted"`
}

func (*QueryIsBlacklistedResponse) ProtoMessage() {}
func (m *QueryIsBlacklistedResponse) Reset() { *m = QueryIsBlacklistedResponse{} }
func (m *QueryIsBlacklistedResponse) String() string { return "" }

type QueryStablecoinStatsRequest struct {
	Denom string `json:"denom"`
}

func (*QueryStablecoinStatsRequest) ProtoMessage() {}
func (m *QueryStablecoinStatsRequest) Reset() { *m = QueryStablecoinStatsRequest{} }
func (m *QueryStablecoinStatsRequest) String() string { return "" }

type QueryStablecoinStatsResponse struct {
	Stats *StablecoinStats `json:"stats"`
}

func (*QueryStablecoinStatsResponse) ProtoMessage() {}
func (m *QueryStablecoinStatsResponse) Reset() { *m = QueryStablecoinStatsResponse{} }
func (m *QueryStablecoinStatsResponse) String() string { return "" }

type StablecoinStats struct {
	TotalSupply     string `json:"total_supply"`
	CirculatingSupply string `json:"circulating_supply"`
	ReserveBalance  string `json:"reserve_balance"`
	NumberOfHolders uint64 `json:"number_of_holders"`
	TransactionCount uint64 `json:"transaction_count"`
}

// QueryServer is the server API for Query service
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	Stablecoin(context.Context, *QueryStablecoinRequest) (*QueryStablecoinResponse, error)
	Stablecoins(context.Context, *QueryStablecoinsRequest) (*QueryStablecoinsResponse, error)
	StablecoinsByIssuer(context.Context, *QueryStablecoinsByIssuerRequest) (*QueryStablecoinsByIssuerResponse, error)
	StablecoinSupply(context.Context, *QueryStablecoinSupplyRequest) (*QueryStablecoinSupplyResponse, error)
	PriceData(context.Context, *QueryPriceDataRequest) (*QueryPriceDataResponse, error)
	ReserveInfo(context.Context, *QueryReserveInfoRequest) (*QueryReserveInfoResponse, error)
	IsWhitelisted(context.Context, *QueryIsWhitelistedRequest) (*QueryIsWhitelistedResponse, error)
	IsBlacklisted(context.Context, *QueryIsBlacklistedRequest) (*QueryIsBlacklistedResponse, error)
	StablecoinStats(context.Context, *QueryStablecoinStatsRequest) (*QueryStablecoinStatsResponse, error)
}

// QueryClient is the client API for Query service
type QueryClient interface {
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	Stablecoin(ctx context.Context, in *QueryStablecoinRequest, opts ...grpc.CallOption) (*QueryStablecoinResponse, error)
	Stablecoins(ctx context.Context, in *QueryStablecoinsRequest, opts ...grpc.CallOption) (*QueryStablecoinsResponse, error)
	StablecoinsByIssuer(ctx context.Context, in *QueryStablecoinsByIssuerRequest, opts ...grpc.CallOption) (*QueryStablecoinsByIssuerResponse, error)
	StablecoinSupply(ctx context.Context, in *QueryStablecoinSupplyRequest, opts ...grpc.CallOption) (*QueryStablecoinSupplyResponse, error)
	PriceData(ctx context.Context, in *QueryPriceDataRequest, opts ...grpc.CallOption) (*QueryPriceDataResponse, error)
	ReserveInfo(ctx context.Context, in *QueryReserveInfoRequest, opts ...grpc.CallOption) (*QueryReserveInfoResponse, error)
	IsWhitelisted(ctx context.Context, in *QueryIsWhitelistedRequest, opts ...grpc.CallOption) (*QueryIsWhitelistedResponse, error)
	IsBlacklisted(ctx context.Context, in *QueryIsBlacklistedRequest, opts ...grpc.CallOption) (*QueryIsBlacklistedResponse, error)
	StablecoinStats(ctx context.Context, in *QueryStablecoinStatsRequest, opts ...grpc.CallOption) (*QueryStablecoinStatsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Stablecoin(ctx context.Context, in *QueryStablecoinRequest, opts ...grpc.CallOption) (*QueryStablecoinResponse, error) {
	out := new(QueryStablecoinResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/Stablecoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Stablecoins(ctx context.Context, in *QueryStablecoinsRequest, opts ...grpc.CallOption) (*QueryStablecoinsResponse, error) {
	out := new(QueryStablecoinsResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/Stablecoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StablecoinsByIssuer(ctx context.Context, in *QueryStablecoinsByIssuerRequest, opts ...grpc.CallOption) (*QueryStablecoinsByIssuerResponse, error) {
	out := new(QueryStablecoinsByIssuerResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/StablecoinsByIssuer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StablecoinSupply(ctx context.Context, in *QueryStablecoinSupplyRequest, opts ...grpc.CallOption) (*QueryStablecoinSupplyResponse, error) {
	out := new(QueryStablecoinSupplyResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/StablecoinSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PriceData(ctx context.Context, in *QueryPriceDataRequest, opts ...grpc.CallOption) (*QueryPriceDataResponse, error) {
	out := new(QueryPriceDataResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/PriceData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ReserveInfo(ctx context.Context, in *QueryReserveInfoRequest, opts ...grpc.CallOption) (*QueryReserveInfoResponse, error) {
	out := new(QueryReserveInfoResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/ReserveInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IsWhitelisted(ctx context.Context, in *QueryIsWhitelistedRequest, opts ...grpc.CallOption) (*QueryIsWhitelistedResponse, error) {
	out := new(QueryIsWhitelistedResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/IsWhitelisted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IsBlacklisted(ctx context.Context, in *QueryIsBlacklistedRequest, opts ...grpc.CallOption) (*QueryIsBlacklistedResponse, error) {
	out := new(QueryIsBlacklistedResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/IsBlacklisted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StablecoinStats(ctx context.Context, in *QueryStablecoinStatsRequest, opts ...grpc.CallOption) (*QueryStablecoinStatsResponse, error) {
	out := new(QueryStablecoinStatsResponse)
	err := c.cc.Invoke(ctx, "/stateset.stablecoins.v1.Query/StablecoinStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}