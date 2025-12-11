package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the feemarket Msg service with the router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the feemarket Msg service for the legacy router.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.feemarket.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "UpdateParams", Handler: _Msg_UpdateParams_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/feemarket",
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgUpdateParams)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.feemarket.Msg/UpdateParams"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, request.(*MsgUpdateParams))
	}
	return interceptor(ctx, req, info, handler)
}
