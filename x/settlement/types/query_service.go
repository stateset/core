package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the settlement Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the settlement Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.settlement.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Settlement", Handler: _Query_Settlement_Handler},
		{MethodName: "Settlements", Handler: _Query_Settlements_Handler},
		{MethodName: "SettlementsByStatus", Handler: _Query_SettlementsByStatus_Handler},
		{MethodName: "Batch", Handler: _Query_Batch_Handler},
		{MethodName: "Batches", Handler: _Query_Batches_Handler},
		{MethodName: "Channel", Handler: _Query_Channel_Handler},
		{MethodName: "Channels", Handler: _Query_Channels_Handler},
		{MethodName: "ChannelsByParty", Handler: _Query_ChannelsByParty_Handler},
		{MethodName: "Merchant", Handler: _Query_Merchant_Handler},
		{MethodName: "Merchants", Handler: _Query_Merchants_Handler},
		{MethodName: "Params", Handler: _Query_Params_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/settlement",
}

func _Query_Settlement_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QuerySettlementRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Settlement(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Settlement"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Settlement(ctx, request.(*QuerySettlementRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Settlements_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QuerySettlementsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Settlements(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Settlements"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Settlements(ctx, request.(*QuerySettlementsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_SettlementsByStatus_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QuerySettlementsByStatusRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SettlementsByStatus(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/SettlementsByStatus"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).SettlementsByStatus(ctx, request.(*QuerySettlementsByStatusRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Batch_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryBatchRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Batch(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Batch"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Batch(ctx, request.(*QueryBatchRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Batches_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryBatchesRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Batches(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Batches"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Batches(ctx, request.(*QueryBatchesRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Channel_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryChannelRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Channel(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Channel"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Channel(ctx, request.(*QueryChannelRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Channels_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryChannelsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Channels(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Channels"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Channels(ctx, request.(*QueryChannelsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_ChannelsByParty_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryChannelsByPartyRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChannelsByParty(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/ChannelsByParty"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).ChannelsByParty(ctx, request.(*QueryChannelsByPartyRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Merchant_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryMerchantRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Merchant(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Merchant"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Merchant(ctx, request.(*QueryMerchantRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Merchants_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryMerchantsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Merchants(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Merchants"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Merchants(ctx, request.(*QueryMerchantsRequest))
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
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.settlement.Query/Params"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, request.(*QueryParamsRequest))
	}
	return interceptor(ctx, req, info, handler)
}
