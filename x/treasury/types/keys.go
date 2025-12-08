package types

import "encoding/binary"

const (
	ModuleName = "treasury"
	StoreKey   = ModuleName
	RouterKey  = ModuleName

	EventTypeSnapshotRecorded = "reserve_snapshot_recorded"

	AttributeKeySnapshotID = "snapshot_id"
	AttributeKeyReporter   = "reporter"
	AttributeKeyTotal      = "total_supply"
)

var (
	SnapshotsKeyPrefix = []byte{0x01}
)

func SnapshotStoreKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, SnapshotsKeyPrefix...), bz...)
}
