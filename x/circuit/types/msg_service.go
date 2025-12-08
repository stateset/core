package types

import (
	"context"

	"google.golang.org/grpc"
)

// MsgServer defines the circuit breaker message server interface
type MsgServer interface {
	PauseSystem(context.Context, *MsgPauseSystem) (*MsgPauseSystemResponse, error)
	ResumeSystem(context.Context, *MsgResumeSystem) (*MsgResumeSystemResponse, error)
	TripCircuit(context.Context, *MsgTripCircuit) (*MsgTripCircuitResponse, error)
	ResetCircuit(context.Context, *MsgResetCircuit) (*MsgResetCircuitResponse, error)
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// Response types
type MsgPauseSystemResponse struct{}
type MsgResumeSystemResponse struct{}
type MsgTripCircuitResponse struct{}
type MsgResetCircuitResponse struct{}
type MsgUpdateParamsResponse struct{}

// RegisterMsgServer registers the message server with gRPC
func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	// Note: Full gRPC registration would require protobuf-generated code
	// This is a placeholder for non-protobuf implementation
}
