package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/agreement/types"
)

// GetAgreementCount get the total number of agreement
func (k Keeper) GetAgreementCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AgreementCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAgreementCount set the total number of agreement
func (k Keeper) SetAgreementCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AgreementCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAgreement appends a agreement in the store with a new id and update the count
func (k Keeper) AppendAgreement(
	ctx sdk.Context,
	agreement types.Agreement,
) uint64 {
	// Create the agreement
	count := k.GetAgreementCount(ctx)

	// Set the ID of the appended value
	agreement.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AgreementKey))
	appendedValue := k.cdc.MustMarshal(&agreement)
	store.Set(GetAgreementIDBytes(agreement.Id), appendedValue)

	// Update agreement count
	k.SetAgreementCount(ctx, count+1)

	return count
}

// SetAgreement set a specific agreement in the store
func (k Keeper) SetAgreement(ctx sdk.Context, agreement types.Agreement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AgreementKey))
	b := k.cdc.MustMarshal(&agreement)
	store.Set(GetAgreementIDBytes(agreement.Id), b)
}

// GetAgreement returns a agreement from its id
func (k Keeper) GetAgreement(ctx sdk.Context, id uint64) (val types.Agreement, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AgreementKey))
	b := store.Get(GetAgreementIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAgreement removes a agreement from the store
func (k Keeper) RemoveAgreement(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AgreementKey))
	store.Delete(GetAgreementIDBytes(id))
}

// GetAllAgreement returns all agreement
func (k Keeper) GetAllAgreement(ctx sdk.Context) (list []types.Agreement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AgreementKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Agreement
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAgreementIDBytes returns the byte representation of the ID
func GetAgreementIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAgreementIDFromBytes returns ID in uint64 format from a byte array
func GetAgreementIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
