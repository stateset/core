package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the settlement Msg service on the gRPC router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the settlement Msg service for the routing layer.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.settlement.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstantTransfer",
			Handler:    _Msg_InstantTransfer_Handler,
		},
		{
			MethodName: "CreateEscrow",
			Handler:    _Msg_CreateEscrow_Handler,
		},
		{
			MethodName: "ReleaseEscrow",
			Handler:    _Msg_ReleaseEscrow_Handler,
		},
		{
			MethodName: "RefundEscrow",
			Handler:    _Msg_RefundEscrow_Handler,
		},
		{
			MethodName: "CreateBatch",
			Handler:    _Msg_CreateBatch_Handler,
		},
		{
			MethodName: "SettleBatch",
			Handler:    _Msg_SettleBatch_Handler,
		},
		{
			MethodName: "OpenChannel",
			Handler:    _Msg_OpenChannel_Handler,
		},
		{
			MethodName: "CloseChannel",
			Handler:    _Msg_CloseChannel_Handler,
		},
		{
			MethodName: "ClaimChannel",
			Handler:    _Msg_ClaimChannel_Handler,
		},
		{
			MethodName: "RegisterMerchant",
			Handler:    _Msg_RegisterMerchant_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _Msg_UpdateMerchant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/settlement",
}

func _Msg_InstantTransfer_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgInstantTransfer)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InstantTransfer(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/InstantTransfer",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).InstantTransfer(ctx, request.(*MsgInstantTransfer))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_CreateEscrow_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCreateEscrow)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateEscrow(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/CreateEscrow",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateEscrow(ctx, request.(*MsgCreateEscrow))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ReleaseEscrow_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgReleaseEscrow)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ReleaseEscrow(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/ReleaseEscrow",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ReleaseEscrow(ctx, request.(*MsgReleaseEscrow))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_RefundEscrow_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRefundEscrow)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RefundEscrow(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/RefundEscrow",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RefundEscrow(ctx, request.(*MsgRefundEscrow))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_CreateBatch_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCreateBatch)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBatch(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/CreateBatch",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBatch(ctx, request.(*MsgCreateBatch))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_SettleBatch_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgSettleBatch)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SettleBatch(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/SettleBatch",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).SettleBatch(ctx, request.(*MsgSettleBatch))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_OpenChannel_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgOpenChannel)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OpenChannel(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/OpenChannel",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).OpenChannel(ctx, request.(*MsgOpenChannel))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_CloseChannel_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCloseChannel)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseChannel(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/CloseChannel",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseChannel(ctx, request.(*MsgCloseChannel))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ClaimChannel_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgClaimChannel)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimChannel(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/ClaimChannel",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimChannel(ctx, request.(*MsgClaimChannel))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_RegisterMerchant_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRegisterMerchant)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterMerchant(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/RegisterMerchant",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterMerchant(ctx, request.(*MsgRegisterMerchant))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_UpdateMerchant_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgUpdateMerchant)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMerchant(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.settlement.Msg/UpdateMerchant",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMerchant(ctx, request.(*MsgUpdateMerchant))
	}
	return interceptor(ctx, req, info, handler)
}
