package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/circuit/types"
)

var _ types.QueryServer = Keeper{}

// Params returns the module parameters
func (k Keeper) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// CircuitState returns the global circuit state
func (k Keeper) CircuitState(ctx context.Context, req *types.QueryCircuitStateRequest) (*types.QueryCircuitStateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	state := k.GetCircuitState(sdkCtx)
	return &types.QueryCircuitStateResponse{State: state}, nil
}

// ModuleCircuit returns a specific module's circuit state
func (k Keeper) ModuleCircuit(ctx context.Context, req *types.QueryModuleCircuitRequest) (*types.QueryModuleCircuitResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	state, _ := k.GetModuleCircuitState(sdkCtx, req.ModuleName)
	return &types.QueryModuleCircuitResponse{State: state}, nil
}

// RateLimits returns rate limit status for an address
func (k Keeper) RateLimits(ctx context.Context, req *types.QueryRateLimitsRequest) (*types.QueryRateLimitsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get all rate limit configs
	params := k.GetParams(sdkCtx)
	var states []types.RateLimitState

	for _, config := range params.RateLimits {
		if config.PerAddress {
			state, _ := k.GetRateLimitState(sdkCtx, config.Name, req.Address)
			states = append(states, state)
		}
	}

	return &types.QueryRateLimitsResponse{RateLimits: states}, nil
}

// LiquidationProtection returns the liquidation surge protection state
func (k Keeper) LiquidationProtection(ctx context.Context, req *types.QueryLiquidationProtectionRequest) (*types.QueryLiquidationProtectionResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	protection := k.GetLiquidationProtection(sdkCtx)
	return &types.QueryLiquidationProtectionResponse{Protection: protection}, nil
}
