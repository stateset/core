package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/feemarket/types"
)

var _ types.QueryServer = Keeper{}

// BaseFee returns the current base fee.
func (k Keeper) BaseFee(goCtx context.Context, req *types.QueryBaseFeeRequest) (*types.QueryBaseFeeResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	baseFee := k.GetBaseFee(ctx)

	return &types.QueryBaseFeeResponse{
		BaseFee: baseFee,
	}, nil
}

// Params returns the module parameters.
func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}

// GasPrice returns the recommended gas price based on current base fee and priority.
func (k Keeper) GasPrice(goCtx context.Context, req *types.QueryGasPriceRequest) (*types.QueryGasPriceResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	baseFee := k.GetBaseFee(ctx)

	// Calculate gas price based on priority
	gasPrice := k.calculateGasPrice(baseFee, req.Priority)

	return &types.QueryGasPriceResponse{
		GasPrice: gasPrice,
	}, nil
}

// EstimateFee estimates the total fee for a transaction with given gas limit.
func (k Keeper) EstimateFee(goCtx context.Context, req *types.QueryEstimateFeeRequest) (*types.QueryEstimateFeeResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	if req.GasLimit == 0 {
		return nil, types.ErrInvalidGasLimit
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	baseFee := k.GetBaseFee(ctx)

	// Calculate gas price with priority (used for potential future enhancements)
	_ = k.calculateGasPrice(baseFee, req.Priority)

	// Calculate priority fee component
	priorityMultiplier := k.getPriorityMultiplier(req.Priority)
	priorityFee := baseFee.Mul(priorityMultiplier)

	// Base fee component: baseFee * gasLimit
	baseFeeComponent := baseFee.MulInt64(int64(req.GasLimit))

	// Priority fee component: priorityFee * gasLimit
	priorityFeeComponent := priorityFee.MulInt64(int64(req.GasLimit))

	// Total estimated fee
	estimatedFee := baseFeeComponent.Add(priorityFeeComponent)

	return &types.QueryEstimateFeeResponse{
		EstimatedFee:         estimatedFee,
		BaseFeeComponent:     baseFeeComponent,
		PriorityFeeComponent: priorityFeeComponent,
	}, nil
}

// FeeHistory returns historical fee data.
func (k Keeper) FeeHistory(goCtx context.Context, req *types.QueryFeeHistoryRequest) (*types.QueryFeeHistoryResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Default limit
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100 // Max limit
	}

	// Get fee history from store
	history := k.GetFeeHistory(ctx, limit)

	return &types.QueryFeeHistoryResponse{
		FeeHistory: history,
	}, nil
}

// calculateGasPrice calculates the recommended gas price based on base fee and priority level.
func (k Keeper) calculateGasPrice(baseFee sdkmath.LegacyDec, priority string) sdkmath.LegacyDec {
	multiplier := k.getPriorityMultiplier(priority)
	return baseFee.Mul(sdkmath.LegacyOneDec().Add(multiplier))
}

// getPriorityMultiplier returns the multiplier for the given priority level.
func (k Keeper) getPriorityMultiplier(priority string) sdkmath.LegacyDec {
	switch priority {
	case "low":
		return sdkmath.LegacyZeroDec() // 0% increase (just base fee)
	case "medium":
		return sdkmath.LegacyMustNewDecFromStr("0.25") // 25% increase
	case "high":
		return sdkmath.LegacyMustNewDecFromStr("0.50") // 50% increase
	default:
		return sdkmath.LegacyMustNewDecFromStr("0.10") // 10% increase (standard)
	}
}

// GetFeeHistory retrieves historical fee data from the store.
func (k Keeper) GetFeeHistory(ctx sdk.Context, limit uint64) []types.FeeHistoryEntry {
	store := ctx.KVStore(k.storeKey)
	history := make([]types.FeeHistoryEntry, 0, limit)

	// Iterate through stored history entries (most recent first)
	currentHeight := ctx.BlockHeight()
	for i := int64(0); i < int64(limit); i++ {
		height := currentHeight - i
		if height <= 0 {
			break
		}

		key := types.GetFeeHistoryKey(height)
		if !store.Has(key) {
			continue
		}

		var entry types.FeeHistoryEntry
		bz := store.Get(key)
		k.cdc.MustUnmarshal(bz, &entry)
		history = append(history, entry)
	}

	return history
}

// StoreFeeHistory stores a fee history entry for the current block.
func (k Keeper) StoreFeeHistory(ctx sdk.Context, entry types.FeeHistoryEntry) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeHistoryKey(entry.BlockHeight)
	bz := k.cdc.MustMarshal(&entry)
	store.Set(key, bz)

	// Optionally prune old entries (keep last 1000 blocks)
	k.pruneOldFeeHistory(ctx)
}

// pruneOldFeeHistory removes fee history entries older than 1000 blocks.
func (k Keeper) pruneOldFeeHistory(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	currentHeight := ctx.BlockHeight()
	pruneHeight := currentHeight - 1000

	if pruneHeight > 0 {
		key := types.GetFeeHistoryKey(pruneHeight)
		if store.Has(key) {
			store.Delete(key)
		}
	}
}
