package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterMsgServer registers the stablecoin Msg service with the router.
func RegisterMsgServer(router grpc.ServiceRegistrar, srv MsgServer) {
	router.RegisterService(&Msg_ServiceDesc, srv)
}

// Msg_ServiceDesc describes the stablecoin Msg service for the legacy router.
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.stablecoin.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVault",
			Handler:    _Msg_CreateVault_Handler,
		},
		{
			MethodName: "DepositCollateral",
			Handler:    _Msg_DepositCollateral_Handler,
		},
		{
			MethodName: "WithdrawCollateral",
			Handler:    _Msg_WithdrawCollateral_Handler,
		},
		{
			MethodName: "MintStablecoin",
			Handler:    _Msg_MintStablecoin_Handler,
		},
		{
			MethodName: "RepayStablecoin",
			Handler:    _Msg_RepayStablecoin_Handler,
		},
		{
			MethodName: "LiquidateVault",
			Handler:    _Msg_LiquidateVault_Handler,
		},
		{
			MethodName: "DepositReserve",
			Handler:    _Msg_DepositReserve_Handler,
		},
		{
			MethodName: "RequestRedemption",
			Handler:    _Msg_RequestRedemption_Handler,
		},
		{
			MethodName: "ExecuteRedemption",
			Handler:    _Msg_ExecuteRedemption_Handler,
		},
		{
			MethodName: "CancelRedemption",
			Handler:    _Msg_CancelRedemption_Handler,
		},
		{
			MethodName: "UpdateReserveParams",
			Handler:    _Msg_UpdateReserveParams_Handler,
		},
		{
			MethodName: "RecordAttestation",
			Handler:    _Msg_RecordAttestation_Handler,
		},
		{
			MethodName: "SetApprovedAttester",
			Handler:    _Msg_SetApprovedAttester_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/stablecoin",
}

func _Msg_CreateVault_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCreateVault)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVault(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/CreateVault",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVault(ctx, request.(*MsgCreateVault))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_DepositCollateral_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgDepositCollateral)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositCollateral(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/DepositCollateral",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositCollateral(ctx, request.(*MsgDepositCollateral))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_WithdrawCollateral_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgWithdrawCollateral)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawCollateral(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/WithdrawCollateral",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawCollateral(ctx, request.(*MsgWithdrawCollateral))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_MintStablecoin_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgMintStablecoin)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintStablecoin(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/MintStablecoin",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).MintStablecoin(ctx, request.(*MsgMintStablecoin))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_RepayStablecoin_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRepayStablecoin)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RepayStablecoin(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/RepayStablecoin",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RepayStablecoin(ctx, request.(*MsgRepayStablecoin))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_LiquidateVault_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgLiquidateVault)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LiquidateVault(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/LiquidateVault",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).LiquidateVault(ctx, request.(*MsgLiquidateVault))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_DepositReserve_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgDepositReserve)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositReserve(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/DepositReserve",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositReserve(ctx, request.(*MsgDepositReserve))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_RequestRedemption_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRequestRedemption)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestRedemption(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/RequestRedemption",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestRedemption(ctx, request.(*MsgRequestRedemption))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_ExecuteRedemption_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgExecuteRedemption)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteRedemption(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/ExecuteRedemption",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteRedemption(ctx, request.(*MsgExecuteRedemption))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_CancelRedemption_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgCancelRedemption)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelRedemption(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/CancelRedemption",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelRedemption(ctx, request.(*MsgCancelRedemption))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_UpdateReserveParams_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgUpdateReserveParams)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateReserveParams(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/UpdateReserveParams",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateReserveParams(ctx, request.(*MsgUpdateReserveParams))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_RecordAttestation_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgRecordAttestation)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RecordAttestation(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/RecordAttestation",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).RecordAttestation(ctx, request.(*MsgRecordAttestation))
	}
	return interceptor(ctx, req, info, handler)
}

func _Msg_SetApprovedAttester_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(MsgSetApprovedAttester)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetApprovedAttester(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.stablecoin.Msg/SetApprovedAttester",
	}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(MsgServer).SetApprovedAttester(ctx, request.(*MsgSetApprovedAttester))
	}
	return interceptor(ctx, req, info, handler)
}
