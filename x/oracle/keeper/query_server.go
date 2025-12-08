package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/oracle/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the oracle QueryServer interface
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

// Price returns the current price for a denom
func (q queryServer) Price(req *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	// For oracle queries, we need to get context from somewhere
	// Since this is a simplified implementation, we'll return an error
	// In production, this would be called via gRPC with proper context
	return nil, types.ErrPriceNotFound
}

// Prices returns all prices with pagination
func (q queryServer) Prices(req *types.QueryPricesRequest) (*types.QueryPricesResponse, error) {
	return nil, types.ErrPriceNotFound
}

// OracleConfig returns the oracle configuration for a denom
func (q queryServer) OracleConfig(req *types.QueryOracleConfigRequest) (*types.QueryOracleConfigResponse, error) {
	return nil, types.ErrInvalidDenom
}

// Provider returns an oracle provider by address
func (q queryServer) Provider(req *types.QueryProviderRequest) (*types.QueryProviderResponse, error) {
	return nil, types.ErrProviderNotFound
}

// Providers returns all oracle providers
func (q queryServer) Providers(req *types.QueryProvidersRequest) (*types.QueryProvidersResponse, error) {
	return nil, types.ErrProviderNotFound
}

// Params returns the module parameters
func (q queryServer) Params(req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	return nil, nil
}

// Context-aware query implementations

// PriceWithContext returns the current price for a denom
func (q queryServer) PriceWithContext(goCtx context.Context, req *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	price, found := q.Keeper.GetPrice(ctx, req.Denom)
	if !found {
		return nil, types.ErrPriceNotFound
	}

	return &types.QueryPriceResponse{
		Price: price,
	}, nil
}

// PricesWithContext returns all prices with pagination
func (q queryServer) PricesWithContext(goCtx context.Context, req *types.QueryPricesRequest) (*types.QueryPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.PriceHistorySize)
	if maxLimit == 0 {
		maxLimit = 100
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var prices []types.Price
	var total uint64

	q.Keeper.IteratePrices(ctx, func(p types.Price) bool {
		if total >= offset && uint64(len(prices)) < limit {
			prices = append(prices, p)
		}
		total++
		return false
	})

	return &types.QueryPricesResponse{
		Prices: prices,
		Total:  total,
	}, nil
}

// OracleConfigWithContext returns the oracle configuration for a denom
func (q queryServer) OracleConfigWithContext(goCtx context.Context, req *types.QueryOracleConfigRequest) (*types.QueryOracleConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	config, found := q.Keeper.GetOracleConfig(ctx, req.Denom)
	if !found {
		return nil, types.ErrInvalidDenom
	}

	return &types.QueryOracleConfigResponse{
		Config: config,
	}, nil
}

// ProviderWithContext returns an oracle provider by address
func (q queryServer) ProviderWithContext(goCtx context.Context, req *types.QueryProviderRequest) (*types.QueryProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := q.Keeper.GetProvider(ctx, req.Address)
	if !found {
		return nil, types.ErrProviderNotFound
	}

	return &types.QueryProviderResponse{
		Provider: provider,
	}, nil
}

// ProvidersWithContext returns all oracle providers
func (q queryServer) ProvidersWithContext(goCtx context.Context, req *types.QueryProvidersRequest) (*types.QueryProvidersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.MaxProviders)
	if maxLimit == 0 {
		maxLimit = 100
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var providers []types.OracleProvider
	var total uint64

	q.Keeper.IterateProviders(ctx, func(p types.OracleProvider) bool {
		if total >= offset && uint64(len(providers)) < limit {
			providers = append(providers, p)
		}
		total++
		return false
	})

	return &types.QueryProvidersResponse{
		Providers: providers,
		Total:     total,
	}, nil
}

// ParamsWithContext returns the module parameters
func (q queryServer) ParamsWithContext(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
