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
