package keeper

import (
	"context"
	"encoding/binary"
	"time"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/treasury/types"
)

var (
	nextIDKey = []byte{0x02}
)

// Keeper manages reserve snapshots and treasury authority.
type Keeper struct {
	storeKey  storetypes.StoreKey
	authority string
}

func NewKeeper(_ codec.BinaryCodec, key storetypes.StoreKey, authority string) Keeper {
	return Keeper{storeKey: key, authority: authority}
}

func (k Keeper) GetAuthority() string { return k.authority }

func (k *Keeper) SetAuthority(authority string) { k.authority = authority }

func (k Keeper) getNextID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(nextIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(nextIDKey, bz)
}

// RecordSnapshot persists a new reserve snapshot and returns its ID.
func (k Keeper) RecordSnapshot(ctx context.Context, snapshot types.ReserveSnapshot) uint64 {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	nextID := k.getNextID(sdkCtx)
	snapshot.Id = nextID
	if snapshot.Timestamp.IsZero() {
		snapshot.Timestamp = time.Unix(sdkCtx.BlockTime().Unix(), 0)
	}

	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&snapshot)
	store.Set(mustWriteUint64(nextID), bz)

	k.setNextID(sdkCtx, nextID+1)
	return nextID
}

// GetSnapshot fetches a snapshot by id.
func (k Keeper) GetSnapshot(ctx context.Context, id uint64) (types.ReserveSnapshot, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	bz := store.Get(mustWriteUint64(id))
	if len(bz) == 0 {
		return types.ReserveSnapshot{}, false
	}
	var snapshot types.ReserveSnapshot
	types.ModuleCdc.MustUnmarshalJSON(bz, &snapshot)
	return snapshot, true
}

// GetLatestSnapshot returns the newest snapshot if any.
func (k Keeper) GetLatestSnapshot(ctx context.Context) (types.ReserveSnapshot, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	iterator := store.ReverseIterator(nil, nil)
	defer iterator.Close()
	if iterator.Valid() {
		var snapshot types.ReserveSnapshot
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &snapshot)
		return snapshot, true
	}
	return types.ReserveSnapshot{}, false
}

// IterateSnapshots iterates over stored snapshots.
func (k Keeper) IterateSnapshots(ctx context.Context, cb func(types.ReserveSnapshot) bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var snapshot types.ReserveSnapshot
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &snapshot)
		if cb(snapshot) {
			break
		}
	}
}

// InitGenesis initializes from genesis state.
func (k Keeper) InitGenesis(ctx context.Context, state *types.GenesisState) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if state == nil {
		state = types.DefaultGenesis()
	}
	k.authority = state.Authority
	k.setNextID(sdkCtx, state.NextID)
	for _, snapshot := range state.Snapshots {
		store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
		bz := types.ModuleCdc.MustMarshalJSON(&snapshot)
		store.Set(mustWriteUint64(snapshot.Id), bz)
	}
}

// ExportGenesis exports the module state.
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Authority = k.authority
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	state.NextID = k.getNextID(sdkCtx)
	k.IterateSnapshots(ctx, func(snapshot types.ReserveSnapshot) bool {
		state.Snapshots = append(state.Snapshots, snapshot)
		return false
	})
	return state
}

func mustWriteUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
