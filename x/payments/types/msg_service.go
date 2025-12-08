package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the payments Msg service on the legacy router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the payments Msg service for the routing layer.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.payments.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _Msg_CreatePayment_Handler,
		},
		{
			MethodName: "SettlePayment",
			Handler:    _Msg_SettlePayment_Handler,
		},
		{
			MethodName: "CancelPayment",
			Handler:    _Msg_CancelPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/payments",
}

func _Msg_CreatePayment_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCreatePayment)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePayment(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.payments.Msg/CreatePayment",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePayment(ctx, request.(*MsgCreatePayment))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_SettlePayment_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgSettlePayment)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SettlePayment(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.payments.Msg/SettlePayment",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).SettlePayment(ctx, request.(*MsgSettlePayment))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_CancelPayment_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCancelPayment)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelPayment(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.payments.Msg/CancelPayment",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelPayment(ctx, request.(*MsgCancelPayment))
	}
	return interceptor(ctx, req, info, handler)
}
