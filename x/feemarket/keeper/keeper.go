package keeper

import (
	"encoding/binary"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/feemarket/types"
)

// Keeper maintains the fee market state.
type Keeper struct {
	cdc       codec.BinaryCodec
	storeKey  storetypes.StoreKey
	authority string
}

// NewKeeper creates a new fee market keeper instance.
func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority string) Keeper {
	return Keeper{
		cdc:       cdc,
		storeKey:  key,
		authority: authority,
	}
}

// GetAuthority returns the module authority address (used for param changes).
func (k Keeper) GetAuthority() string {
	return k.authority
}

// GetParams fetches module parameters or returns defaults if unset.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.ParamsKey) {
		return types.DefaultParams()
	}
	var params types.Params
	types.ModuleCdc.MustUnmarshalJSON(store.Get(types.ParamsKey), &params)
	return params
}

// SetParams stores module parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ParamsKey, types.ModuleCdc.MustMarshalJSON(&params))
}

// GetBaseFee returns the current base fee, falling back to the initial default.
func (k Keeper) GetBaseFee(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.BaseFeeKey) {
		return types.DefaultInitialBaseFee
	}
	var dec sdk.Dec
	types.ModuleCdc.MustUnmarshalJSON(store.Get(types.BaseFeeKey), &dec)
	return dec
}

// SetBaseFee persists the current base fee.
func (k Keeper) SetBaseFee(ctx sdk.Context, baseFee sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BaseFeeKey, types.ModuleCdc.MustMarshalJSON(&baseFee))
}

// UpdateBaseFee computes and persists the next base fee using the provided gas usage and block gas limit.
func (k Keeper) UpdateBaseFee(ctx sdk.Context, gasUsed uint64, maxBlockGas uint64) sdk.Dec {
	params := k.GetParams(ctx)
	current := k.GetBaseFee(ctx)
	next := types.ComputeNextBaseFee(current, gasUsed, params, maxBlockGas)
	k.SetBaseFee(ctx, next)
	k.setLatestGas(ctx, gasUsed)
	return next
}

// setLatestGas stores the last block's gas usage for quick oracle responses.
func (k Keeper) setLatestGas(ctx sdk.Context, gasUsed uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, gasUsed)
	store.Set(types.LatestGasKey, bz)
}

// GetLatestGas returns the last recorded gas usage.
func (k Keeper) GetLatestGas(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.LatestGasKey) {
		return 0
	}
	return binary.BigEndian.Uint64(store.Get(types.LatestGasKey))
}
