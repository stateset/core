package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stateset/core/x/stablecoin/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl returns a QueryServer implementation backed by the stablecoin keeper.
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

// Vault returns a vault by ID
func (q queryServer) Vault(goCtx context.Context, req *types.QueryVaultRequest) (*types.QueryVaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	vault, found := q.keeper.GetVault(ctx, req.VaultId)
	if !found {
		return nil, status.Error(codes.NotFound, "vault not found")
	}

	return &types.QueryVaultResponse{Vault: vault}, nil
}

// Vaults returns all vaults with optional owner filter
func (q queryServer) Vaults(goCtx context.Context, req *types.QueryVaultsRequest) (*types.QueryVaultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var vaults []types.Vault
	q.keeper.IterateVaults(ctx, func(vault types.Vault) bool {
		if req.Owner == "" || vault.Owner == req.Owner {
			vaults = append(vaults, vault)
		}
		return false
	})

	return &types.QueryVaultsResponse{Vaults: vaults}, nil
}

// ReserveParams returns the reserve parameters
func (q queryServer) ReserveParams(goCtx context.Context, req *types.QueryReserveParamsRequest) (*types.QueryReserveParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.keeper.GetReserveParams(ctx)
	return &types.QueryReserveParamsResponse{Params: params}, nil
}

// Reserve returns the current reserve state
func (q queryServer) Reserve(goCtx context.Context, req *types.QueryReserveRequest) (*types.QueryReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	reserve := q.keeper.GetReserve(ctx)
	return &types.QueryReserveResponse{Reserve: reserve}, nil
}

// TotalReserves returns combined on-chain and off-chain reserves
func (q queryServer) TotalReserves(goCtx context.Context, req *types.QueryTotalReservesRequest) (*types.QueryTotalReservesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	totalReserves := q.keeper.GetTotalReserves(ctx)
	return &types.QueryTotalReservesResponse{TotalReserves: totalReserves}, nil
}

// ReserveDeposit returns a reserve deposit by ID
func (q queryServer) ReserveDeposit(goCtx context.Context, req *types.QueryReserveDepositRequest) (*types.QueryReserveDepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	deposit, found := q.keeper.GetReserveDeposit(ctx, req.DepositId)
	if !found {
		return nil, status.Error(codes.NotFound, "deposit not found")
	}

	return &types.QueryReserveDepositResponse{Deposit: deposit}, nil
}

// ReserveDeposits returns all reserve deposits with optional depositor filter
func (q queryServer) ReserveDeposits(goCtx context.Context, req *types.QueryReserveDepositsRequest) (*types.QueryReserveDepositsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var deposits []types.ReserveDeposit
	q.keeper.IterateReserveDeposits(ctx, func(deposit types.ReserveDeposit) bool {
		if req.Depositor == "" || deposit.Depositor == req.Depositor {
			deposits = append(deposits, deposit)
		}
		return false
	})

	return &types.QueryReserveDepositsResponse{Deposits: deposits}, nil
}

// RedemptionRequest returns a redemption request by ID
func (q queryServer) RedemptionRequest(goCtx context.Context, req *types.QueryRedemptionRequestRequest) (*types.QueryRedemptionRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	redemption, found := q.keeper.GetRedemptionRequest(ctx, req.RedemptionId)
	if !found {
		return nil, status.Error(codes.NotFound, "redemption request not found")
	}

	return &types.QueryRedemptionRequestResponse{Redemption: redemption}, nil
}

// RedemptionRequests returns all redemption requests with optional status filter
func (q queryServer) RedemptionRequests(goCtx context.Context, req *types.QueryRedemptionRequestsRequest) (*types.QueryRedemptionRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var redemptions []types.RedemptionRequest
	q.keeper.IterateRedemptionRequests(ctx, func(redemption types.RedemptionRequest) bool {
		if req.Status == "" || string(redemption.Status) == req.Status {
			redemptions = append(redemptions, redemption)
		}
		return false
	})

	return &types.QueryRedemptionRequestsResponse{Redemptions: redemptions}, nil
}

// LatestAttestation returns the latest off-chain reserve attestation
func (q queryServer) LatestAttestation(goCtx context.Context, req *types.QueryLatestAttestationRequest) (*types.QueryLatestAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	attestation, found := q.keeper.GetLatestAttestation(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "no attestations found")
	}

	return &types.QueryLatestAttestationResponse{Attestation: attestation}, nil
}

// Attestation returns an attestation by ID
func (q queryServer) Attestation(goCtx context.Context, req *types.QueryAttestationRequest) (*types.QueryAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	attestation, found := q.keeper.GetAttestation(ctx, req.AttestationId)
	if !found {
		return nil, status.Error(codes.NotFound, "attestation not found")
	}

	return &types.QueryAttestationResponse{Attestation: attestation}, nil
}

// DailyStats returns the daily mint/redeem statistics
func (q queryServer) DailyStats(goCtx context.Context, req *types.QueryDailyStatsRequest) (*types.QueryDailyStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	stats := q.keeper.GetDailyMintStats(ctx)
	return &types.QueryDailyStatsResponse{Stats: stats}, nil
}
