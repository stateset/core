package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the oracle Msg service with the router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the oracle Msg service.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.oracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePrice",
			Handler:    _Msg_UpdatePrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/oracle",
}

func _Msg_UpdatePrice_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgUpdatePrice)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePrice(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.oracle.Msg/UpdatePrice",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePrice(ctx, request.(*MsgUpdatePrice))
	}
	return interceptor(ctx, req, info, handler)
}
