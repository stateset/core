package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/purchaseorder/types"
)

// GetSentPurchaseorderCount get the total number of sentPurchaseorder
func (k Keeper) GetSentPurchaseorderCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentPurchaseorderCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSentPurchaseorderCount set the total number of sentPurchaseorder
func (k Keeper) SetSentPurchaseorderCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentPurchaseorderCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSentPurchaseorder appends a sentPurchaseorder in the store with a new id and update the count
func (k Keeper) AppendSentPurchaseorder(
	ctx sdk.Context,
	sentPurchaseorder types.SentPurchaseorder,
) uint64 {
	// Create the sentPurchaseorder
	count := k.GetSentPurchaseorderCount(ctx)

	// Set the ID of the appended value
	sentPurchaseorder.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPurchaseorderKey))
	appendedValue := k.cdc.MustMarshal(&sentPurchaseorder)
	store.Set(GetSentPurchaseorderIDBytes(sentPurchaseorder.Id), appendedValue)

	// Update sentPurchaseorder count
	k.SetSentPurchaseorderCount(ctx, count+1)

	return count
}

// SetSentPurchaseorder set a specific sentPurchaseorder in the store
func (k Keeper) SetSentPurchaseorder(ctx sdk.Context, sentPurchaseorder types.SentPurchaseorder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPurchaseorderKey))
	b := k.cdc.MustMarshal(&sentPurchaseorder)
	store.Set(GetSentPurchaseorderIDBytes(sentPurchaseorder.Id), b)
}

// GetSentPurchaseorder returns a sentPurchaseorder from its id
func (k Keeper) GetSentPurchaseorder(ctx sdk.Context, id uint64) (val types.SentPurchaseorder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPurchaseorderKey))
	b := store.Get(GetSentPurchaseorderIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSentPurchaseorder removes a sentPurchaseorder from the store
func (k Keeper) RemoveSentPurchaseorder(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPurchaseorderKey))
	store.Delete(GetSentPurchaseorderIDBytes(id))
}

// GetAllSentPurchaseorder returns all sentPurchaseorder
func (k Keeper) GetAllSentPurchaseorder(ctx sdk.Context) (list []types.SentPurchaseorder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentPurchaseorderKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SentPurchaseorder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSentPurchaseorderIDBytes returns the byte representation of the ID
func GetSentPurchaseorderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSentPurchaseorderIDFromBytes returns ID in uint64 format from a byte array
func GetSentPurchaseorderIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
