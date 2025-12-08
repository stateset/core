package types

import (
	"context"
)

// QueryServer defines the query server interface for the circuit module
type QueryServer interface {
	// Params queries the module parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// CircuitState queries the global circuit state
	CircuitState(context.Context, *QueryCircuitStateRequest) (*QueryCircuitStateResponse, error)
	// ModuleCircuit queries a specific module's circuit state
	ModuleCircuit(context.Context, *QueryModuleCircuitRequest) (*QueryModuleCircuitResponse, error)
	// RateLimits queries rate limit status for an address
	RateLimits(context.Context, *QueryRateLimitsRequest) (*QueryRateLimitsResponse, error)
	// LiquidationProtection queries the liquidation surge protection state
	LiquidationProtection(context.Context, *QueryLiquidationProtectionRequest) (*QueryLiquidationProtectionResponse, error)
}

// Query request/response types

type QueryParamsRequest struct{}

type QueryParamsResponse struct {
	Params Params `json:"params"`
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return "" }
func (m *QueryParamsRequest) ProtoMessage()  {}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return "" }
func (m *QueryParamsResponse) ProtoMessage()  {}

type QueryCircuitStateRequest struct{}

type QueryCircuitStateResponse struct {
	State CircuitState `json:"state"`
}

func (m *QueryCircuitStateRequest) Reset()         { *m = QueryCircuitStateRequest{} }
func (m *QueryCircuitStateRequest) String() string { return "" }
func (m *QueryCircuitStateRequest) ProtoMessage()  {}

func (m *QueryCircuitStateResponse) Reset()         { *m = QueryCircuitStateResponse{} }
func (m *QueryCircuitStateResponse) String() string { return "" }
func (m *QueryCircuitStateResponse) ProtoMessage()  {}

type QueryModuleCircuitRequest struct {
	ModuleName string `json:"module_name"`
}

type QueryModuleCircuitResponse struct {
	State ModuleCircuitState `json:"state"`
}

func (m *QueryModuleCircuitRequest) Reset()         { *m = QueryModuleCircuitRequest{} }
func (m *QueryModuleCircuitRequest) String() string { return "" }
func (m *QueryModuleCircuitRequest) ProtoMessage()  {}

func (m *QueryModuleCircuitResponse) Reset()         { *m = QueryModuleCircuitResponse{} }
func (m *QueryModuleCircuitResponse) String() string { return "" }
func (m *QueryModuleCircuitResponse) ProtoMessage()  {}

type QueryRateLimitsRequest struct {
	Address string `json:"address"`
}

type QueryRateLimitsResponse struct {
	RateLimits []RateLimitState `json:"rate_limits"`
}

func (m *QueryRateLimitsRequest) Reset()         { *m = QueryRateLimitsRequest{} }
func (m *QueryRateLimitsRequest) String() string { return "" }
func (m *QueryRateLimitsRequest) ProtoMessage()  {}

func (m *QueryRateLimitsResponse) Reset()         { *m = QueryRateLimitsResponse{} }
func (m *QueryRateLimitsResponse) String() string { return "" }
func (m *QueryRateLimitsResponse) ProtoMessage()  {}

type QueryLiquidationProtectionRequest struct{}

type QueryLiquidationProtectionResponse struct {
	Protection LiquidationSurgeProtection `json:"protection"`
}

func (m *QueryLiquidationProtectionRequest) Reset()         { *m = QueryLiquidationProtectionRequest{} }
func (m *QueryLiquidationProtectionRequest) String() string { return "" }
func (m *QueryLiquidationProtectionRequest) ProtoMessage()  {}

func (m *QueryLiquidationProtectionResponse) Reset()         { *m = QueryLiquidationProtectionResponse{} }
func (m *QueryLiquidationProtectionResponse) String() string { return "" }
func (m *QueryLiquidationProtectionResponse) ProtoMessage()  {}
