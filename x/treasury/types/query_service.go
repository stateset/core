package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the treasury Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the treasury Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.treasury.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Params", Handler: _Query_Params_Handler},
		{MethodName: "SpendProposal", Handler: _Query_SpendProposal_Handler},
		{MethodName: "SpendProposals", Handler: _Query_SpendProposals_Handler},
		{MethodName: "Budget", Handler: _Query_Budget_Handler},
		{MethodName: "Budgets", Handler: _Query_Budgets_Handler},
		{MethodName: "TreasuryBalance", Handler: _Query_TreasuryBalance_Handler},
		{MethodName: "LatestSnapshot", Handler: _Query_LatestSnapshot_Handler},
		{MethodName: "Allocation", Handler: _Query_Allocation_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/treasury",
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryParamsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/Params"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, request.(*QueryParamsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_SpendProposal_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QuerySpendProposalRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SpendProposal(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/SpendProposal"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).SpendProposal(ctx, request.(*QuerySpendProposalRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_SpendProposals_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QuerySpendProposalsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SpendProposals(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/SpendProposals"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).SpendProposals(ctx, request.(*QuerySpendProposalsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Budget_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryBudgetRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Budget(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/Budget"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Budget(ctx, request.(*QueryBudgetRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Budgets_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryBudgetsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Budgets(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/Budgets"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Budgets(ctx, request.(*QueryBudgetsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_TreasuryBalance_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryTreasuryBalanceRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TreasuryBalance(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/TreasuryBalance"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).TreasuryBalance(ctx, request.(*QueryTreasuryBalanceRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_LatestSnapshot_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryLatestSnapshotRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LatestSnapshot(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/LatestSnapshot"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).LatestSnapshot(ctx, request.(*QueryLatestSnapshotRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Allocation_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryAllocationRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Allocation(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.treasury.Query/Allocation"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Allocation(ctx, request.(*QueryAllocationRequest))
	}
	return interceptor(ctx, req, info, handler)
}
