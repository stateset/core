package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stateset/core/x/treasury/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl returns a QueryServer implementation backed by the treasury keeper.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{keeper: k}
}

// Params returns the module parameters
func (q queryServer) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.keeper.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// SpendProposal returns a spend proposal by ID
func (q queryServer) SpendProposal(goCtx context.Context, req *types.QuerySpendProposalRequest) (*types.QuerySpendProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal, found := q.keeper.GetSpendProposal(ctx, req.ProposalId)
	if !found {
		return nil, status.Error(codes.NotFound, "proposal not found")
	}

	return &types.QuerySpendProposalResponse{Proposal: proposal}, nil
}

// SpendProposals returns all spend proposals with optional status filter
func (q queryServer) SpendProposals(goCtx context.Context, req *types.QuerySpendProposalsRequest) (*types.QuerySpendProposalsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var proposals []types.SpendProposal
	q.keeper.IterateSpendProposals(ctx, func(proposal types.SpendProposal) bool {
		if req.Status == "" || proposal.Status == req.Status {
			proposals = append(proposals, proposal)
		}
		return false
	})

	return &types.QuerySpendProposalsResponse{Proposals: proposals}, nil
}

// Budget returns a budget by category
func (q queryServer) Budget(goCtx context.Context, req *types.QueryBudgetRequest) (*types.QueryBudgetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	budget, found := q.keeper.GetBudget(ctx, req.Category)
	if !found {
		return nil, status.Error(codes.NotFound, "budget not found")
	}

	return &types.QueryBudgetResponse{Budget: budget}, nil
}

// Budgets returns all budgets
func (q queryServer) Budgets(goCtx context.Context, req *types.QueryBudgetsRequest) (*types.QueryBudgetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var budgets []types.Budget
	q.keeper.IterateBudgets(ctx, func(budget types.Budget) bool {
		budgets = append(budgets, budget)
		return false
	})

	return &types.QueryBudgetsResponse{Budgets: budgets}, nil
}

// TreasuryBalance returns the treasury balance
func (q queryServer) TreasuryBalance(goCtx context.Context, req *types.QueryTreasuryBalanceRequest) (*types.QueryTreasuryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	balance := q.keeper.GetTreasuryBalance(ctx)
	return &types.QueryTreasuryBalanceResponse{Balance: balance}, nil
}

// LatestSnapshot returns the latest reserve snapshot
func (q queryServer) LatestSnapshot(goCtx context.Context, req *types.QueryLatestSnapshotRequest) (*types.QueryLatestSnapshotResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	snapshot, found := q.keeper.GetLatestSnapshot(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "no snapshots found")
	}

	return &types.QueryLatestSnapshotResponse{Snapshot: snapshot}, nil
}

// Allocation returns a specific allocation by recipient
func (q queryServer) Allocation(goCtx context.Context, req *types.QueryAllocationRequest) (*types.QueryAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allocation, found := q.keeper.GetAllocation(ctx, req.Recipient)
	if !found {
		return nil, status.Error(codes.NotFound, "allocation not found")
	}

	return &types.QueryAllocationResponse{Allocation: allocation}, nil
}
