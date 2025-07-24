package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/types"
)

// GetInvoiceCount get the total number of invoice
func (k Keeper) GetInvoiceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.InvoiceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetInvoiceCount set the total number of invoice
func (k Keeper) SetInvoiceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.InvoiceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendInvoice appends a invoice in the store with a new id and update the count
func (k Keeper) AppendInvoice(
	ctx sdk.Context,
	invoice types.Invoice,
) uint64 {
	// Create the invoice
	count := k.GetInvoiceCount(ctx)

	// Set the ID of the appended value
	invoice.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvoiceKey))
	appendedValue := k.cdc.MustMarshal(&invoice)
	store.Set(GetInvoiceIDBytes(invoice.Id), appendedValue)

	// Update invoice count
	k.SetInvoiceCount(ctx, count+1)

	return count
}

// SetInvoice set a specific invoice in the store
func (k Keeper) SetInvoice(ctx sdk.Context, invoice types.Invoice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvoiceKey))
	b := k.cdc.MustMarshal(&invoice)
	store.Set(GetInvoiceIDBytes(invoice.Id), b)
}

// GetInvoice returns a invoice from its id
func (k Keeper) GetInvoice(ctx sdk.Context, id uint64) (val types.Invoice, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvoiceKey))
	b := store.Get(GetInvoiceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInvoice removes a invoice from the store
func (k Keeper) RemoveInvoice(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvoiceKey))
	store.Delete(GetInvoiceIDBytes(id))
}

// GetAllInvoice returns all invoice
func (k Keeper) GetAllInvoice(ctx sdk.Context) (list []types.Invoice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvoiceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Invoice
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetInvoiceIDBytes returns the byte representation of the ID
func GetInvoiceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetInvoiceIDFromBytes returns ID in uint64 format from a byte array
func GetInvoiceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
