package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the feemarket Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the feemarket Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.feemarket.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "BaseFee", Handler: _Query_BaseFee_Handler},
		{MethodName: "Params", Handler: _Query_Params_Handler},
		{MethodName: "GasPrice", Handler: _Query_GasPrice_Handler},
		{MethodName: "EstimateFee", Handler: _Query_EstimateFee_Handler},
		{MethodName: "FeeHistory", Handler: _Query_FeeHistory_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/feemarket",
}

func _Query_BaseFee_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryBaseFeeRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BaseFee(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.feemarket.Query/BaseFee"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).BaseFee(ctx, request.(*QueryBaseFeeRequest))
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
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.feemarket.Query/Params"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, request.(*QueryParamsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_GasPrice_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryGasPriceRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GasPrice(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.feemarket.Query/GasPrice"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).GasPrice(ctx, request.(*QueryGasPriceRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_EstimateFee_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryEstimateFeeRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EstimateFee(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.feemarket.Query/EstimateFee"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).EstimateFee(ctx, request.(*QueryEstimateFeeRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_FeeHistory_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryFeeHistoryRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeHistory(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.feemarket.Query/FeeHistory"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeHistory(ctx, request.(*QueryFeeHistoryRequest))
	}
	return interceptor(ctx, req, info, handler)
}
