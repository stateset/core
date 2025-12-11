package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the orders Msg service with the router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the orders Msg service for the routing layer.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.core.orders.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateOrder", Handler: _Msg_CreateOrder_Handler},
		{MethodName: "ConfirmOrder", Handler: _Msg_ConfirmOrder_Handler},
		{MethodName: "PayOrder", Handler: _Msg_PayOrder_Handler},
		{MethodName: "ShipOrder", Handler: _Msg_ShipOrder_Handler},
		{MethodName: "DeliverOrder", Handler: _Msg_DeliverOrder_Handler},
		{MethodName: "CompleteOrder", Handler: _Msg_CompleteOrder_Handler},
		{MethodName: "CancelOrder", Handler: _Msg_CancelOrder_Handler},
		{MethodName: "RefundOrder", Handler: _Msg_RefundOrder_Handler},
		{MethodName: "OpenDispute", Handler: _Msg_OpenDispute_Handler},
		{MethodName: "ResolveDispute", Handler: _Msg_ResolveDispute_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/orders",
}

func _Msg_CreateOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCreateOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/CreateOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateOrder(ctx, request.(*MsgCreateOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ConfirmOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgConfirmOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConfirmOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/ConfirmOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ConfirmOrder(ctx, request.(*MsgConfirmOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_PayOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgPayOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PayOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/PayOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).PayOrder(ctx, request.(*MsgPayOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ShipOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgShipOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ShipOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/ShipOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ShipOrder(ctx, request.(*MsgShipOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_DeliverOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgDeliverOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeliverOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/DeliverOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).DeliverOrder(ctx, request.(*MsgDeliverOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_CompleteOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCompleteOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CompleteOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/CompleteOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CompleteOrder(ctx, request.(*MsgCompleteOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_CancelOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCancelOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/CancelOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelOrder(ctx, request.(*MsgCancelOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_RefundOrder_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRefundOrder)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RefundOrder(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/RefundOrder"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RefundOrder(ctx, request.(*MsgRefundOrder))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_OpenDispute_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgOpenDispute)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OpenDispute(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/OpenDispute"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).OpenDispute(ctx, request.(*MsgOpenDispute))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ResolveDispute_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgResolveDispute)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ResolveDispute(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.core.orders.Msg/ResolveDispute"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ResolveDispute(ctx, request.(*MsgResolveDispute))
	}
	return interceptor(ctx, req, info, handler)
}
