package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/purchaseorder/types"
)

// GetPurchaseorderCount get the total number of purchaseorder
func (k Keeper) GetPurchaseorderCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PurchaseorderCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPurchaseorderCount set the total number of purchaseorder
func (k Keeper) SetPurchaseorderCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PurchaseorderCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPurchaseorder appends a purchaseorder in the store with a new id and update the count
func (k Keeper) AppendPurchaseorder(
	ctx sdk.Context,
	purchaseorder types.Purchaseorder,
) uint64 {
	// Create the purchaseorder
	count := k.GetPurchaseorderCount(ctx)

	// Set the ID of the appended value
	purchaseorder.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PurchaseorderKey))
	appendedValue := k.cdc.MustMarshal(&purchaseorder)
	store.Set(GetPurchaseorderIDBytes(purchaseorder.Id), appendedValue)

	// Update purchaseorder count
	k.SetPurchaseorderCount(ctx, count+1)

	return count
}

// SetPurchaseorder set a specific purchaseorder in the store
func (k Keeper) SetPurchaseorder(ctx sdk.Context, purchaseorder types.Purchaseorder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PurchaseorderKey))
	b := k.cdc.MustMarshal(&purchaseorder)
	store.Set(GetPurchaseorderIDBytes(purchaseorder.Id), b)
}

// GetPurchaseorder returns a purchaseorder from its id
func (k Keeper) GetPurchaseorder(ctx sdk.Context, id uint64) (val types.Purchaseorder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PurchaseorderKey))
	b := store.Get(GetPurchaseorderIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePurchaseorder removes a purchaseorder from the store
func (k Keeper) RemovePurchaseorder(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PurchaseorderKey))
	store.Delete(GetPurchaseorderIDBytes(id))
}

// GetAllPurchaseorder returns all purchaseorder
func (k Keeper) GetAllPurchaseorder(ctx sdk.Context) (list []types.Purchaseorder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PurchaseorderKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Purchaseorder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPurchaseorderIDBytes returns the byte representation of the ID
func GetPurchaseorderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPurchaseorderIDFromBytes returns ID in uint64 format from a byte array
func GetPurchaseorderIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
