package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the stablecoin Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the stablecoin Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.stablecoin.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Params", Handler: _Query_Params_Handler},
		{MethodName: "Vault", Handler: _Query_Vault_Handler},
		{MethodName: "Vaults", Handler: _Query_Vaults_Handler},
		{MethodName: "ReserveParams", Handler: _Query_ReserveParams_Handler},
		{MethodName: "Reserve", Handler: _Query_Reserve_Handler},
		{MethodName: "TotalReserves", Handler: _Query_TotalReserves_Handler},
		{MethodName: "ReserveDeposit", Handler: _Query_ReserveDeposit_Handler},
		{MethodName: "ReserveDeposits", Handler: _Query_ReserveDeposits_Handler},
		{MethodName: "RedemptionRequest", Handler: _Query_RedemptionRequest_Handler},
		{MethodName: "RedemptionRequests", Handler: _Query_RedemptionRequests_Handler},
		{MethodName: "LatestAttestation", Handler: _Query_LatestAttestation_Handler},
		{MethodName: "Attestation", Handler: _Query_Attestation_Handler},
		{MethodName: "DailyStats", Handler: _Query_DailyStats_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/stablecoin",
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryParamsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/Params"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, request.(*QueryParamsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Vault_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryVaultRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vault(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/Vault"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Vault(ctx, request.(*QueryVaultRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Vaults_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryVaultsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vaults(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/Vaults"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Vaults(ctx, request.(*QueryVaultsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_ReserveParams_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryReserveParamsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReserveParams(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/ReserveParams"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).ReserveParams(ctx, request.(*QueryReserveParamsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Reserve_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryReserveRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Reserve(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/Reserve"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Reserve(ctx, request.(*QueryReserveRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_TotalReserves_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryTotalReservesRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalReserves(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/TotalReserves"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalReserves(ctx, request.(*QueryTotalReservesRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_ReserveDeposit_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryReserveDepositRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReserveDeposit(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/ReserveDeposit"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).ReserveDeposit(ctx, request.(*QueryReserveDepositRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_ReserveDeposits_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryReserveDepositsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReserveDeposits(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/ReserveDeposits"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).ReserveDeposits(ctx, request.(*QueryReserveDepositsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_RedemptionRequest_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryRedemptionRequestRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RedemptionRequest(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/RedemptionRequest"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).RedemptionRequest(ctx, request.(*QueryRedemptionRequestRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_RedemptionRequests_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryRedemptionRequestsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RedemptionRequests(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/RedemptionRequests"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).RedemptionRequests(ctx, request.(*QueryRedemptionRequestsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_LatestAttestation_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryLatestAttestationRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LatestAttestation(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/LatestAttestation"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).LatestAttestation(ctx, request.(*QueryLatestAttestationRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Attestation_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryAttestationRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Attestation(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/Attestation"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Attestation(ctx, request.(*QueryAttestationRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_DailyStats_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryDailyStatsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DailyStats(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.stablecoin.Query/DailyStats"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).DailyStats(ctx, request.(*QueryDailyStatsRequest))
	}
	return interceptor(ctx, req, info, handler)
}
