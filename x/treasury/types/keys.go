package types

import "encoding/binary"

const (
	ModuleName        = "treasury"
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	ModuleAccountName = ModuleName

	// Event types
	EventTypeSnapshotRecorded  = "reserve_snapshot_recorded"
	EventTypeFundAllocated     = "fund_allocated"
	EventTypeFundDisbursed     = "fund_disbursed"
	EventTypeSpendProposed     = "spend_proposed"
	EventTypeSpendExecuted     = "spend_executed"
	EventTypeSpendCancelled    = "spend_cancelled"
	EventTypeBudgetSet         = "budget_set"
	EventTypeRevenueReceived   = "revenue_received"

	// Attribute keys
	AttributeKeySnapshotID   = "snapshot_id"
	AttributeKeyReporter     = "reporter"
	AttributeKeyTotal        = "total_supply"
	AttributeKeyRecipient    = "recipient"
	AttributeKeyAmount       = "amount"
	AttributeKeyProposalID   = "proposal_id"
	AttributeKeyCategory     = "category"
	AttributeKeySource       = "source"
	AttributeKeyExecuteAfter = "execute_after"

	// Spend proposal statuses
	SpendStatusPending   = "pending"
	SpendStatusExecuted  = "executed"
	SpendStatusCancelled = "cancelled"
	SpendStatusExpired   = "expired"

	// Budget categories
	CategoryDevelopment   = "development"
	CategoryMarketing     = "marketing"
	CategoryOperations    = "operations"
	CategoryGrants        = "grants"
	CategorySecurity      = "security"
	CategoryInfrastructure = "infrastructure"
	CategoryReserve       = "reserve"
)

var (
	SnapshotsKeyPrefix     = []byte{0x01}
	ParamsKey              = []byte{0x02}
	SpendProposalKeyPrefix = []byte{0x03}
	BudgetKeyPrefix        = []byte{0x04}
	AllocationKeyPrefix    = []byte{0x05}
	RevenueRecordKeyPrefix = []byte{0x06}
)

func SnapshotStoreKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, SnapshotsKeyPrefix...), bz...)
}

func SpendProposalKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, SpendProposalKeyPrefix...), bz...)
}

func BudgetKey(category string) []byte {
	return append(append([]byte{}, BudgetKeyPrefix...), []byte(category)...)
}

func AllocationKey(recipient string) []byte {
	return append(append([]byte{}, AllocationKeyPrefix...), []byte(recipient)...)
}

func RevenueRecordKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, RevenueRecordKeyPrefix...), bz...)
}
