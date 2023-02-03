package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/proof/types"
)

// GetProofCount get the total number of proof
func (k Keeper) GetProofCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProofCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetProofCount set the total number of proof
func (k Keeper) SetProofCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProofCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendProof appends a proof in the store with a new id and update the count
func (k Keeper) AppendProof(
	ctx sdk.Context,
	proof types.Proof,
) uint64 {
	// Create the proof
	count := k.GetProofCount(ctx)

	// Set the ID of the appended value
	proof.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	appendedValue := k.cdc.MustMarshal(&proof)
	store.Set(GetProofIDBytes(proof.Id), appendedValue)

	// Update proof count
	k.SetProofCount(ctx, count+1)

	return count
}

// SetProof set a specific proof in the store
func (k Keeper) SetProof(ctx sdk.Context, proof types.Proof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	b := k.cdc.MustMarshal(&proof)
	store.Set(GetProofIDBytes(proof.Id), b)
}

// GetProof returns a proof from its id
func (k Keeper) GetProof(ctx sdk.Context, id uint64) (val types.Proof, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	b := store.Get(GetProofIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProof removes a proof from the store
func (k Keeper) RemoveProof(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	store.Delete(GetProofIDBytes(id))
}

// GetAllProof returns all proof
func (k Keeper) GetAllProof(ctx sdk.Context) (list []types.Proof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proof
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetProofIDBytes returns the byte representation of the ID
func GetProofIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetProofIDFromBytes returns ID in uint64 format from a byte array
func GetProofIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
