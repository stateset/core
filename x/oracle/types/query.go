package types

import "context"

// QueryServer defines the gRPC query server interface for the oracle module.
type QueryServer interface {
	Price(context.Context, *QueryPriceRequest) (*QueryPriceResponse, error)
	Prices(context.Context, *QueryPricesRequest) (*QueryPricesResponse, error)
	OracleConfig(context.Context, *QueryOracleConfigRequest) (*QueryOracleConfigResponse, error)
	Provider(context.Context, *QueryProviderRequest) (*QueryProviderResponse, error)
	Providers(context.Context, *QueryProvidersRequest) (*QueryProvidersResponse, error)
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// QueryPriceRequest is the request for a single price
type QueryPriceRequest struct {
	Denom string `json:"denom"`
}

// QueryPriceResponse is the response containing a single price
type QueryPriceResponse struct {
	Price Price `json:"price"`
}

// QueryPricesRequest is the request for all prices
type QueryPricesRequest struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

// QueryPricesResponse is the response containing multiple prices
type QueryPricesResponse struct {
	Prices []Price `json:"prices"`
	Total  uint64  `json:"total"`
}

// QueryOracleConfigRequest is the request for oracle config
type QueryOracleConfigRequest struct {
	Denom string `json:"denom"`
}

// QueryOracleConfigResponse is the response containing oracle config
type QueryOracleConfigResponse struct {
	Config OracleConfig `json:"config"`
}

// QueryProviderRequest is the request for a provider
type QueryProviderRequest struct {
	Address string `json:"address"`
}

// QueryProviderResponse is the response containing a provider
type QueryProviderResponse struct {
	Provider OracleProvider `json:"provider"`
}

// QueryProvidersRequest is the request for all providers
type QueryProvidersRequest struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

// QueryProvidersResponse is the response containing providers
type QueryProvidersResponse struct {
	Providers []OracleProvider `json:"providers"`
	Total     uint64           `json:"total"`
}

// QueryParamsRequest is the request for params
type QueryParamsRequest struct{}

// QueryParamsResponse is the response containing params
type QueryParamsResponse struct {
	Params OracleParams `json:"params"`
}
