package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/purchaseorder/types"
)

// GetTimedoutPurchaseorderCount get the total number of timedoutPurchaseorder
func (k Keeper) GetTimedoutPurchaseorderCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutPurchaseorderCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTimedoutPurchaseorderCount set the total number of timedoutPurchaseorder
func (k Keeper) SetTimedoutPurchaseorderCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutPurchaseorderCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTimedoutPurchaseorder appends a timedoutPurchaseorder in the store with a new id and update the count
func (k Keeper) AppendTimedoutPurchaseorder(
	ctx sdk.Context,
	timedoutPurchaseorder types.TimedoutPurchaseorder,
) uint64 {
	// Create the timedoutPurchaseorder
	count := k.GetTimedoutPurchaseorderCount(ctx)

	// Set the ID of the appended value
	timedoutPurchaseorder.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPurchaseorderKey))
	appendedValue := k.cdc.MustMarshal(&timedoutPurchaseorder)
	store.Set(GetTimedoutPurchaseorderIDBytes(timedoutPurchaseorder.Id), appendedValue)

	// Update timedoutPurchaseorder count
	k.SetTimedoutPurchaseorderCount(ctx, count+1)

	return count
}

// SetTimedoutPurchaseorder set a specific timedoutPurchaseorder in the store
func (k Keeper) SetTimedoutPurchaseorder(ctx sdk.Context, timedoutPurchaseorder types.TimedoutPurchaseorder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPurchaseorderKey))
	b := k.cdc.MustMarshal(&timedoutPurchaseorder)
	store.Set(GetTimedoutPurchaseorderIDBytes(timedoutPurchaseorder.Id), b)
}

// GetTimedoutPurchaseorder returns a timedoutPurchaseorder from its id
func (k Keeper) GetTimedoutPurchaseorder(ctx sdk.Context, id uint64) (val types.TimedoutPurchaseorder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPurchaseorderKey))
	b := store.Get(GetTimedoutPurchaseorderIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimedoutPurchaseorder removes a timedoutPurchaseorder from the store
func (k Keeper) RemoveTimedoutPurchaseorder(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPurchaseorderKey))
	store.Delete(GetTimedoutPurchaseorderIDBytes(id))
}

// GetAllTimedoutPurchaseorder returns all timedoutPurchaseorder
func (k Keeper) GetAllTimedoutPurchaseorder(ctx sdk.Context) (list []types.TimedoutPurchaseorder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPurchaseorderKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimedoutPurchaseorder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTimedoutPurchaseorderIDBytes returns the byte representation of the ID
func GetTimedoutPurchaseorderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTimedoutPurchaseorderIDFromBytes returns ID in uint64 format from a byte array
func GetTimedoutPurchaseorderIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
