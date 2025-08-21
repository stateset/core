package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Query request and response types for working capital

type QueryWorkingCapitalLoanRequest struct {
	LoanId string `json:"loan_id"`
}

type QueryWorkingCapitalLoanResponse struct {
	Loan *WorkingCapitalLoan `json:"loan"`
}

type QueryAllWorkingCapitalLoansRequest struct {
	Borrower   string `json:"borrower,omitempty"`
	Status     string `json:"status,omitempty"`
	Pagination *PageRequest `json:"pagination,omitempty"`
}

type QueryAllWorkingCapitalLoansResponse struct {
	Loans      []WorkingCapitalLoan `json:"loans"`
	Pagination *PageResponse `json:"pagination,omitempty"`
}

type QueryWorkingCapitalRequestRequest struct {
	RequestId string `json:"request_id"`
}

type QueryWorkingCapitalRequestResponse struct {
	Request *WorkingCapitalRequest `json:"request"`
}

type QueryAllWorkingCapitalRequestsRequest struct {
	Requester  string `json:"requester,omitempty"`
	Status     string `json:"status,omitempty"`
	Pagination *PageRequest `json:"pagination,omitempty"`
}

type QueryAllWorkingCapitalRequestsResponse struct {
	Requests   []WorkingCapitalRequest `json:"requests"`
	Pagination *PageResponse `json:"pagination,omitempty"`
}

type QueryWorkingCapitalPoolRequest struct {
	PoolId string `json:"pool_id"`
}

type QueryWorkingCapitalPoolResponse struct {
	Pool *WorkingCapitalPool `json:"pool"`
}

type QueryAllWorkingCapitalPoolsRequest struct {
	StablecoinDenom string `json:"stablecoin_denom,omitempty"`
	Status          string `json:"status,omitempty"`
	Pagination      *PageRequest `json:"pagination,omitempty"`
}

type QueryAllWorkingCapitalPoolsResponse struct {
	Pools      []WorkingCapitalPool `json:"pools"`
	Pagination *PageResponse `json:"pagination,omitempty"`
}

type QueryCreditLineRequest struct {
	CreditLineId string `json:"credit_line_id"`
}

type QueryCreditLineResponse struct {
	CreditLine *CreditLine `json:"credit_line"`
}

type QueryAllCreditLinesRequest struct {
	Borrower   string `json:"borrower,omitempty"`
	Status     string `json:"status,omitempty"`
	Pagination *PageRequest `json:"pagination,omitempty"`
}

type QueryAllCreditLinesResponse struct {
	CreditLines []CreditLine `json:"credit_lines"`
	Pagination  *PageResponse `json:"pagination,omitempty"`
}

type QueryRepaymentScheduleRequest struct {
	LoanId string `json:"loan_id"`
}

type QueryRepaymentScheduleResponse struct {
	Schedule []RepaymentSchedule `json:"schedule"`
}

type QueryLoanStatisticsRequest struct {
	PoolId          string `json:"pool_id,omitempty"`
	StablecoinDenom string `json:"stablecoin_denom,omitempty"`
}

type QueryLoanStatisticsResponse struct {
	TotalLoanAmount       sdk.Coins `json:"total_loan_amount"`
	TotalOutstanding      sdk.Coins `json:"total_outstanding"`
	TotalRepaid           sdk.Coins `json:"total_repaid"`
	ActiveLoans           int64     `json:"active_loans"`
	DefaultedLoans        int64     `json:"defaulted_loans"`
	AverageInterestRate   sdk.Dec   `json:"average_interest_rate"`
	DefaultRate           sdk.Dec   `json:"default_rate"`
}

type QueryPoolUtilizationRequest struct {
	PoolId string `json:"pool_id"`
}

type QueryPoolUtilizationResponse struct {
	PoolId              string    `json:"pool_id"`
	TotalFunds          sdk.Coins `json:"total_funds"`
	AvailableFunds      sdk.Coins `json:"available_funds"`
	LentFunds           sdk.Coins `json:"lent_funds"`
	UtilizationRate     sdk.Dec   `json:"utilization_rate"`
	AverageInterestRate sdk.Dec   `json:"average_interest_rate"`
	DefaultRate         sdk.Dec   `json:"default_rate"`
}

type QueryBorrowerProfileRequest struct {
	Borrower string `json:"borrower"`
}

type QueryBorrowerProfileResponse struct {
	Profile           *BusinessProfile      `json:"profile"`
	ActiveLoans       []WorkingCapitalLoan  `json:"active_loans"`
	CreditLines       []CreditLine          `json:"credit_lines"`
	TotalBorrowed     sdk.Coins            `json:"total_borrowed"`
	TotalRepaid       sdk.Coins            `json:"total_repaid"`
	OutstandingAmount sdk.Coins            `json:"outstanding_amount"`
	CreditScore       int                  `json:"credit_score"`
	RepaymentHistory  string               `json:"repayment_history"`
}

// Pagination types (if not already defined elsewhere)
type PageRequest struct {
	Key        []byte `json:"key,omitempty"`
	Offset     uint64 `json:"offset,omitempty"`
	Limit      uint64 `json:"limit,omitempty"`
	CountTotal bool   `json:"count_total,omitempty"`
	Reverse    bool   `json:"reverse,omitempty"`
}

type PageResponse struct {
	NextKey []byte `json:"next_key,omitempty"`
	Total   uint64 `json:"total,omitempty"`
}