package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryServer defines the query server interface for the treasury module
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	SpendProposal(context.Context, *QuerySpendProposalRequest) (*QuerySpendProposalResponse, error)
	SpendProposals(context.Context, *QuerySpendProposalsRequest) (*QuerySpendProposalsResponse, error)
	Budget(context.Context, *QueryBudgetRequest) (*QueryBudgetResponse, error)
	Budgets(context.Context, *QueryBudgetsRequest) (*QueryBudgetsResponse, error)
	TreasuryBalance(context.Context, *QueryTreasuryBalanceRequest) (*QueryTreasuryBalanceResponse, error)
	LatestSnapshot(context.Context, *QueryLatestSnapshotRequest) (*QueryLatestSnapshotResponse, error)
	Allocation(context.Context, *QueryAllocationRequest) (*QueryAllocationResponse, error)
}

// Params query
type QueryParamsRequest struct{}
type QueryParamsResponse struct {
	Params TreasuryParams `json:"params"`
}

// Spend proposal queries
type QuerySpendProposalRequest struct {
	ProposalId uint64 `json:"proposal_id"`
}
type QuerySpendProposalResponse struct {
	Proposal SpendProposal `json:"proposal"`
}

type QuerySpendProposalsRequest struct {
	Status string `json:"status,omitempty"`
}
type QuerySpendProposalsResponse struct {
	Proposals []SpendProposal `json:"proposals"`
}

// Budget queries
type QueryBudgetRequest struct {
	Category string `json:"category"`
}
type QueryBudgetResponse struct {
	Budget Budget `json:"budget"`
}

type QueryBudgetsRequest struct{}
type QueryBudgetsResponse struct {
	Budgets []Budget `json:"budgets"`
}

// Treasury balance query
type QueryTreasuryBalanceRequest struct{}
type QueryTreasuryBalanceResponse struct {
	Balance sdk.Coins `json:"balance"`
}

// Snapshot query
type QueryLatestSnapshotRequest struct{}
type QueryLatestSnapshotResponse struct {
	Snapshot ReserveSnapshot `json:"snapshot"`
}

// Allocation query
type QueryAllocationRequest struct {
	Recipient string `json:"recipient"`
}
type QueryAllocationResponse struct {
	Allocation Allocation `json:"allocation"`
}
