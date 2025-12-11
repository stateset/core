package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the zkpverify Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the zkpverify Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.zkpverify.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Circuit", Handler: _Query_Circuit_Handler},
		{MethodName: "Circuits", Handler: _Query_Circuits_Handler},
		{MethodName: "Proof", Handler: _Query_Proof_Handler},
		{MethodName: "VerificationResult", Handler: _Query_VerificationResult_Handler},
		{MethodName: "SymbolicRules", Handler: _Query_SymbolicRules_Handler},
		{MethodName: "Params", Handler: _Query_Params_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/zkpverify",
}

func _Query_Circuit_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryCircuitRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Circuit(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Query/Circuit"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Circuit(ctx, request.(*QueryCircuitRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Circuits_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryCircuitsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Circuits(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Query/Circuits"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Circuits(ctx, request.(*QueryCircuitsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Proof_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryProofRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proof(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Query/Proof"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Proof(ctx, request.(*QueryProofRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_VerificationResult_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryVerificationResultRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VerificationResult(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Query/VerificationResult"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).VerificationResult(ctx, request.(*QueryVerificationResultRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_SymbolicRules_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QuerySymbolicRulesRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SymbolicRules(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Query/SymbolicRules"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).SymbolicRules(ctx, request.(*QuerySymbolicRulesRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryParamsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Query/Params"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, request.(*QueryParamsRequest))
	}
	return interceptor(ctx, req, info, handler)
}
