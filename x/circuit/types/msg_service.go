package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// MsgServer defines the circuit breaker message server interface.
type MsgServer interface {
	PauseSystem(context.Context, *MsgPauseSystem) (*MsgPauseSystemResponse, error)
	ResumeSystem(context.Context, *MsgResumeSystem) (*MsgResumeSystemResponse, error)
	TripCircuit(context.Context, *MsgTripCircuit) (*MsgTripCircuitResponse, error)
	ResetCircuit(context.Context, *MsgResetCircuit) (*MsgResetCircuitResponse, error)
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// Response types.
type MsgPauseSystemResponse struct{}
type MsgResumeSystemResponse struct{}
type MsgTripCircuitResponse struct{}
type MsgResetCircuitResponse struct{}
type MsgUpdateParamsResponse struct{}

// RegisterMsgServer registers the circuit Msg service with the router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the circuit Msg service for the legacy router.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.circuit.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "PauseSystem", Handler: _Msg_PauseSystem_Handler},
		{MethodName: "ResumeSystem", Handler: _Msg_ResumeSystem_Handler},
		{MethodName: "TripCircuit", Handler: _Msg_TripCircuit_Handler},
		{MethodName: "ResetCircuit", Handler: _Msg_ResetCircuit_Handler},
		{MethodName: "UpdateParams", Handler: _Msg_UpdateParams_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/circuit",
}

func _Msg_PauseSystem_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgPauseSystem)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PauseSystem(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Msg/PauseSystem"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).PauseSystem(ctx, request.(*MsgPauseSystem))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ResumeSystem_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgResumeSystem)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ResumeSystem(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Msg/ResumeSystem"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ResumeSystem(ctx, request.(*MsgResumeSystem))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_TripCircuit_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgTripCircuit)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TripCircuit(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Msg/TripCircuit"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).TripCircuit(ctx, request.(*MsgTripCircuit))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ResetCircuit_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgResetCircuit)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ResetCircuit(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Msg/ResetCircuit"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ResetCircuit(ctx, request.(*MsgResetCircuit))
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
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.circuit.Msg/UpdateParams"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, request.(*MsgUpdateParams))
	}
	return interceptor(ctx, req, info, handler)
}
