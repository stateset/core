package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the treasury Msg service with the router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the treasury Msg service.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.treasury.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordReserve",
			Handler:    _Msg_RecordReserve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/treasury",
}

func _Msg_RecordReserve_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRecordReserve)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RecordReserve(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.treasury.Msg/RecordReserve",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RecordReserve(ctx, request.(*MsgRecordReserve))
	}
	return interceptor(ctx, req, info, handler)
}
