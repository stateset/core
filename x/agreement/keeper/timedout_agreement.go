package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/agreement/types"
)

// GetTimedoutAgreementCount get the total number of timedoutAgreement
func (k Keeper) GetTimedoutAgreementCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutAgreementCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTimedoutAgreementCount set the total number of timedoutAgreement
func (k Keeper) SetTimedoutAgreementCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutAgreementCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTimedoutAgreement appends a timedoutAgreement in the store with a new id and update the count
func (k Keeper) AppendTimedoutAgreement(
	ctx sdk.Context,
	timedoutAgreement types.TimedoutAgreement,
) uint64 {
	// Create the timedoutAgreement
	count := k.GetTimedoutAgreementCount(ctx)

	// Set the ID of the appended value
	timedoutAgreement.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutAgreementKey))
	appendedValue := k.cdc.MustMarshal(&timedoutAgreement)
	store.Set(GetTimedoutAgreementIDBytes(timedoutAgreement.Id), appendedValue)

	// Update timedoutAgreement count
	k.SetTimedoutAgreementCount(ctx, count+1)

	return count
}

// SetTimedoutAgreement set a specific timedoutAgreement in the store
func (k Keeper) SetTimedoutAgreement(ctx sdk.Context, timedoutAgreement types.TimedoutAgreement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutAgreementKey))
	b := k.cdc.MustMarshal(&timedoutAgreement)
	store.Set(GetTimedoutAgreementIDBytes(timedoutAgreement.Id), b)
}

// GetTimedoutAgreement returns a timedoutAgreement from its id
func (k Keeper) GetTimedoutAgreement(ctx sdk.Context, id uint64) (val types.TimedoutAgreement, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutAgreementKey))
	b := store.Get(GetTimedoutAgreementIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimedoutAgreement removes a timedoutAgreement from the store
func (k Keeper) RemoveTimedoutAgreement(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutAgreementKey))
	store.Delete(GetTimedoutAgreementIDBytes(id))
}

// GetAllTimedoutAgreement returns all timedoutAgreement
func (k Keeper) GetAllTimedoutAgreement(ctx sdk.Context) (list []types.TimedoutAgreement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutAgreementKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimedoutAgreement
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTimedoutAgreementIDBytes returns the byte representation of the ID
func GetTimedoutAgreementIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTimedoutAgreementIDFromBytes returns ID in uint64 format from a byte array
func GetTimedoutAgreementIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
