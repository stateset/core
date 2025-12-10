package types

import "context"

// QueryServer defines the query server interface for the stablecoin module
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	Vault(context.Context, *QueryVaultRequest) (*QueryVaultResponse, error)
	Vaults(context.Context, *QueryVaultsRequest) (*QueryVaultsResponse, error)
	ReserveParams(context.Context, *QueryReserveParamsRequest) (*QueryReserveParamsResponse, error)
	Reserve(context.Context, *QueryReserveRequest) (*QueryReserveResponse, error)
	TotalReserves(context.Context, *QueryTotalReservesRequest) (*QueryTotalReservesResponse, error)
	ReserveDeposit(context.Context, *QueryReserveDepositRequest) (*QueryReserveDepositResponse, error)
	ReserveDeposits(context.Context, *QueryReserveDepositsRequest) (*QueryReserveDepositsResponse, error)
	RedemptionRequest(context.Context, *QueryRedemptionRequestRequest) (*QueryRedemptionRequestResponse, error)
	RedemptionRequests(context.Context, *QueryRedemptionRequestsRequest) (*QueryRedemptionRequestsResponse, error)
	LatestAttestation(context.Context, *QueryLatestAttestationRequest) (*QueryLatestAttestationResponse, error)
	Attestation(context.Context, *QueryAttestationRequest) (*QueryAttestationResponse, error)
	DailyStats(context.Context, *QueryDailyStatsRequest) (*QueryDailyStatsResponse, error)
}

// Params query
type QueryParamsRequest struct{}
type QueryParamsResponse struct {
	Params Params `json:"params"`
}

// Vault queries
type QueryVaultRequest struct {
	VaultId uint64 `json:"vault_id"`
}
type QueryVaultResponse struct {
	Vault Vault `json:"vault"`
}

type QueryVaultsRequest struct {
	Owner string `json:"owner,omitempty"`
}
type QueryVaultsResponse struct {
	Vaults []Vault `json:"vaults"`
}

// Reserve queries
type QueryReserveParamsRequest struct{}
type QueryReserveParamsResponse struct {
	Params ReserveParams `json:"params"`
}

type QueryReserveRequest struct{}
type QueryReserveResponse struct {
	Reserve Reserve `json:"reserve"`
}

type QueryTotalReservesRequest struct{}
type QueryTotalReservesResponse struct {
	TotalReserves TotalReserves `json:"total_reserves"`
}

// Reserve deposit queries
type QueryReserveDepositRequest struct {
	DepositId uint64 `json:"deposit_id"`
}
type QueryReserveDepositResponse struct {
	Deposit ReserveDeposit `json:"deposit"`
}

type QueryReserveDepositsRequest struct {
	Depositor string `json:"depositor,omitempty"`
}
type QueryReserveDepositsResponse struct {
	Deposits []ReserveDeposit `json:"deposits"`
}

// Redemption queries
type QueryRedemptionRequestRequest struct {
	RedemptionId uint64 `json:"redemption_id"`
}
type QueryRedemptionRequestResponse struct {
	Redemption RedemptionRequest `json:"redemption"`
}

type QueryRedemptionRequestsRequest struct {
	Status string `json:"status,omitempty"`
}
type QueryRedemptionRequestsResponse struct {
	Redemptions []RedemptionRequest `json:"redemptions"`
}

// Attestation queries
type QueryLatestAttestationRequest struct{}
type QueryLatestAttestationResponse struct {
	Attestation OffChainReserveAttestation `json:"attestation"`
}

type QueryAttestationRequest struct {
	AttestationId uint64 `json:"attestation_id"`
}
type QueryAttestationResponse struct {
	Attestation OffChainReserveAttestation `json:"attestation"`
}

// Daily stats query
type QueryDailyStatsRequest struct{}
type QueryDailyStatsResponse struct {
	Stats DailyMintStats `json:"stats"`
}
