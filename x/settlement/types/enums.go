package types

// SettlementStatus represents the status of a settlement.
type SettlementStatus string

const (
	SettlementStatusPending    SettlementStatus = "pending"
	SettlementStatusProcessing SettlementStatus = "processing"
	SettlementStatusCompleted  SettlementStatus = "completed"
	SettlementStatusFailed     SettlementStatus = "failed"
	SettlementStatusRefunded   SettlementStatus = "refunded"
	SettlementStatusCancelled  SettlementStatus = "cancelled"
)

// SettlementType represents the type of settlement.
type SettlementType string

const (
	SettlementTypeInstant   SettlementType = "instant"
	SettlementTypeEscrow    SettlementType = "escrow"
	SettlementTypeBatch     SettlementType = "batch"
	SettlementTypeRecurring SettlementType = "recurring"
)
