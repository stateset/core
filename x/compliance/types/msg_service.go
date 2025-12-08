package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the compliance Msg service with the given router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the compliance Msg service for the legacy router.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.compliance.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertProfile",
			Handler:    _Msg_UpsertProfile_Handler,
		},
		{
			MethodName: "SetSanction",
			Handler:    _Msg_SetSanction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/compliance",
}

func _Msg_UpsertProfile_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgUpsertProfile)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpsertProfile(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.compliance.Msg/UpsertProfile",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).UpsertProfile(ctx, request.(*MsgUpsertProfile))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_SetSanction_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgSetSanction)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetSanction(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.compliance.Msg/SetSanction",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).SetSanction(ctx, request.(*MsgSetSanction))
	}
	return interceptor(ctx, req, info, handler)
}
