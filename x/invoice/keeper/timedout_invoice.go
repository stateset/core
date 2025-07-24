package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/types"
)

// GetTimedoutInvoiceCount get the total number of timedoutInvoice
func (k Keeper) GetTimedoutInvoiceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutInvoiceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTimedoutInvoiceCount set the total number of timedoutInvoice
func (k Keeper) SetTimedoutInvoiceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutInvoiceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTimedoutInvoice appends a timedoutInvoice in the store with a new id and update the count
func (k Keeper) AppendTimedoutInvoice(
	ctx sdk.Context,
	timedoutInvoice types.TimedoutInvoice,
) uint64 {
	// Create the timedoutInvoice
	count := k.GetTimedoutInvoiceCount(ctx)

	// Set the ID of the appended value
	timedoutInvoice.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutInvoiceKey))
	appendedValue := k.cdc.MustMarshal(&timedoutInvoice)
	store.Set(GetTimedoutInvoiceIDBytes(timedoutInvoice.Id), appendedValue)

	// Update timedoutInvoice count
	k.SetTimedoutInvoiceCount(ctx, count+1)

	return count
}

// SetTimedoutInvoice set a specific timedoutInvoice in the store
func (k Keeper) SetTimedoutInvoice(ctx sdk.Context, timedoutInvoice types.TimedoutInvoice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutInvoiceKey))
	b := k.cdc.MustMarshal(&timedoutInvoice)
	store.Set(GetTimedoutInvoiceIDBytes(timedoutInvoice.Id), b)
}

// GetTimedoutInvoice returns a timedoutInvoice from its id
func (k Keeper) GetTimedoutInvoice(ctx sdk.Context, id uint64) (val types.TimedoutInvoice, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutInvoiceKey))
	b := store.Get(GetTimedoutInvoiceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimedoutInvoice removes a timedoutInvoice from the store
func (k Keeper) RemoveTimedoutInvoice(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutInvoiceKey))
	store.Delete(GetTimedoutInvoiceIDBytes(id))
}

// GetAllTimedoutInvoice returns all timedoutInvoice
func (k Keeper) GetAllTimedoutInvoice(ctx sdk.Context) (list []types.TimedoutInvoice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutInvoiceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimedoutInvoice
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTimedoutInvoiceIDBytes returns the byte representation of the ID
func GetTimedoutInvoiceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTimedoutInvoiceIDFromBytes returns ID in uint64 format from a byte array
func GetTimedoutInvoiceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
