package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the circuit Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the circuit Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.circuit.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Params", Handler: _Query_Params_Handler},
		{MethodName: "CircuitState", Handler: _Query_CircuitState_Handler},
		{MethodName: "ModuleCircuit", Handler: _Query_ModuleCircuit_Handler},
		{MethodName: "RateLimits", Handler: _Query_RateLimits_Handler},
		{MethodName: "LiquidationProtection", Handler: _Query_LiquidationProtection_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/circuit",
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryParamsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Query/Params"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, request.(*QueryParamsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_CircuitState_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryCircuitStateRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CircuitState(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Query/CircuitState"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).CircuitState(ctx, request.(*QueryCircuitStateRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_ModuleCircuit_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryModuleCircuitRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModuleCircuit(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Query/ModuleCircuit"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).ModuleCircuit(ctx, request.(*QueryModuleCircuitRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_RateLimits_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryRateLimitsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RateLimits(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Query/RateLimits"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).RateLimits(ctx, request.(*QueryRateLimitsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_LiquidationProtection_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryLiquidationProtectionRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidationProtection(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Query/LiquidationProtection"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidationProtection(ctx, request.(*QueryLiquidationProtectionRequest))
	}
	return interceptor(ctx, req, info, handler)
}
