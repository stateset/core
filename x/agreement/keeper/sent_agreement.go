package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/agreement/types"
)

// GetSentAgreementCount get the total number of sentAgreement
func (k Keeper) GetSentAgreementCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentAgreementCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSentAgreementCount set the total number of sentAgreement
func (k Keeper) SetSentAgreementCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentAgreementCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSentAgreement appends a sentAgreement in the store with a new id and update the count
func (k Keeper) AppendSentAgreement(
	ctx sdk.Context,
	sentAgreement types.SentAgreement,
) uint64 {
	// Create the sentAgreement
	count := k.GetSentAgreementCount(ctx)

	// Set the ID of the appended value
	sentAgreement.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentAgreementKey))
	appendedValue := k.cdc.MustMarshal(&sentAgreement)
	store.Set(GetSentAgreementIDBytes(sentAgreement.Id), appendedValue)

	// Update sentAgreement count
	k.SetSentAgreementCount(ctx, count+1)

	return count
}

// SetSentAgreement set a specific sentAgreement in the store
func (k Keeper) SetSentAgreement(ctx sdk.Context, sentAgreement types.SentAgreement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentAgreementKey))
	b := k.cdc.MustMarshal(&sentAgreement)
	store.Set(GetSentAgreementIDBytes(sentAgreement.Id), b)
}

// GetSentAgreement returns a sentAgreement from its id
func (k Keeper) GetSentAgreement(ctx sdk.Context, id uint64) (val types.SentAgreement, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentAgreementKey))
	b := store.Get(GetSentAgreementIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSentAgreement removes a sentAgreement from the store
func (k Keeper) RemoveSentAgreement(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentAgreementKey))
	store.Delete(GetSentAgreementIDBytes(id))
}

// GetAllSentAgreement returns all sentAgreement
func (k Keeper) GetAllSentAgreement(ctx sdk.Context) (list []types.SentAgreement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentAgreementKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SentAgreement
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSentAgreementIDBytes returns the byte representation of the ID
func GetSentAgreementIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSentAgreementIDFromBytes returns ID in uint64 format from a byte array
func GetSentAgreementIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
