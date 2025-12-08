package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/zkpverify/types"
)

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl returns the QueryServer implementation
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{keeper: keeper}
}

var _ types.QueryServer = queryServer{}

// Circuit returns circuit details by name
func (q queryServer) Circuit(
	goCtx context.Context,
	req *types.QueryCircuitRequest,
) (*types.QueryCircuitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	circuit, found := q.keeper.GetCircuit(ctx, req.Name)
	if !found {
		return nil, types.ErrCircuitNotFound
	}

	return &types.QueryCircuitResponse{
		Circuit: circuit,
	}, nil
}

// Circuits returns all registered circuits
func (q queryServer) Circuits(
	goCtx context.Context,
	req *types.QueryCircuitsRequest,
) (*types.QueryCircuitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allCircuits := q.keeper.GetAllCircuits(ctx)

	if req.ActiveOnly {
		var activeCircuits []types.Circuit
		for _, c := range allCircuits {
			if c.Active {
				activeCircuits = append(activeCircuits, c)
			}
		}
		return &types.QueryCircuitsResponse{Circuits: activeCircuits}, nil
	}

	return &types.QueryCircuitsResponse{Circuits: allCircuits}, nil
}

// Proof returns proof details by ID
func (q queryServer) Proof(
	goCtx context.Context,
	req *types.QueryProofRequest,
) (*types.QueryProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proof, found := q.keeper.GetProof(ctx, req.ProofID)
	if !found {
		return nil, types.ErrProofNotFound
	}

	return &types.QueryProofResponse{
		Proof: proof,
	}, nil
}

// VerificationResult returns the verification result for a proof
func (q queryServer) VerificationResult(
	goCtx context.Context,
	req *types.QueryVerificationResultRequest,
) (*types.QueryVerificationResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result, found := q.keeper.GetVerificationResult(ctx, req.ProofID)
	if !found {
		return nil, types.ErrProofNotFound
	}

	return &types.QueryVerificationResultResponse{
		Result: result,
	}, nil
}

// SymbolicRules returns all symbolic rules for a circuit
func (q queryServer) SymbolicRules(
	goCtx context.Context,
	req *types.QuerySymbolicRulesRequest,
) (*types.QuerySymbolicRulesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	rules := q.keeper.GetSymbolicRulesForCircuit(ctx, req.CircuitName)

	return &types.QuerySymbolicRulesResponse{
		Rules: rules,
	}, nil
}

// Params returns module parameters
func (q queryServer) Params(
	goCtx context.Context,
	_ *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.keeper.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
