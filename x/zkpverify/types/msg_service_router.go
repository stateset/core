package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the zkpverify Msg service with the router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the zkpverify Msg service for the legacy router.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.zkpverify.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "RegisterCircuit", Handler: _Msg_RegisterCircuit_Handler},
		{MethodName: "DeactivateCircuit", Handler: _Msg_DeactivateCircuit_Handler},
		{MethodName: "RegisterSymbolicRule", Handler: _Msg_RegisterSymbolicRule_Handler},
		{MethodName: "SubmitProof", Handler: _Msg_SubmitProof_Handler},
		{MethodName: "ChallengeProof", Handler: _Msg_ChallengeProof_Handler},
		{MethodName: "UpdateParams", Handler: _Msg_UpdateParams_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/zkpverify",
}

func _Msg_RegisterCircuit_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRegisterCircuit)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterCircuit(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Msg/RegisterCircuit"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterCircuit(ctx, request.(*MsgRegisterCircuit))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_DeactivateCircuit_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgDeactivateCircuit)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeactivateCircuit(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Msg/DeactivateCircuit"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).DeactivateCircuit(ctx, request.(*MsgDeactivateCircuit))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_RegisterSymbolicRule_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRegisterSymbolicRule)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterSymbolicRule(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Msg/RegisterSymbolicRule"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterSymbolicRule(ctx, request.(*MsgRegisterSymbolicRule))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_SubmitProof_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgSubmitProof)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProof(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Msg/SubmitProof"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProof(ctx, request.(*MsgSubmitProof))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ChallengeProof_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgChallengeProof)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ChallengeProof(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Msg/ChallengeProof"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ChallengeProof(ctx, request.(*MsgChallengeProof))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgUpdateParams)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.zkpverify.Msg/UpdateParams"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, request.(*MsgUpdateParams))
	}
	return interceptor(ctx, req, info, handler)
}
