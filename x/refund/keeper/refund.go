package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/refund/types"
)

// GetRefundCount get the total number of refund
func (k Keeper) GetRefundCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RefundCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRefundCount set the total number of refund
func (k Keeper) SetRefundCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RefundCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRefund appends a refund in the store with a new id and update the count
func (k Keeper) AppendRefund(
	ctx sdk.Context,
	refund types.Refund,
) uint64 {
	// Create the refund
	count := k.GetRefundCount(ctx)

	// Set the ID of the appended value
	refund.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RefundKey))
	appendedValue := k.cdc.MustMarshal(&refund)
	store.Set(GetRefundIDBytes(refund.Id), appendedValue)

	// Update refund count
	k.SetRefundCount(ctx, count+1)

	return count
}

// SetRefund set a specific refund in the store
func (k Keeper) SetRefund(ctx sdk.Context, refund types.Refund) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RefundKey))
	b := k.cdc.MustMarshal(&refund)
	store.Set(GetRefundIDBytes(refund.Id), b)
}

// GetRefund returns a refund from its id
func (k Keeper) GetRefund(ctx sdk.Context, id uint64) (val types.Refund, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RefundKey))
	b := store.Get(GetRefundIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRefund removes a refund from the store
func (k Keeper) RemoveRefund(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RefundKey))
	store.Delete(GetRefundIDBytes(id))
}

// GetAllRefund returns all refund
func (k Keeper) GetAllRefund(ctx sdk.Context) (list []types.Refund) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RefundKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Refund
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRefundIDBytes returns the byte representation of the ID
func GetRefundIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRefundIDFromBytes returns ID in uint64 format from a byte array
func GetRefundIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
