package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SettlementStatus represents the status of a settlement
type SettlementStatus string

const (
	SettlementStatusPending    SettlementStatus = "pending"
	SettlementStatusProcessing SettlementStatus = "processing"
	SettlementStatusCompleted  SettlementStatus = "completed"
	SettlementStatusFailed     SettlementStatus = "failed"
	SettlementStatusRefunded   SettlementStatus = "refunded"
	SettlementStatusCancelled  SettlementStatus = "cancelled"
)

// SettlementType represents the type of settlement
type SettlementType string

const (
	SettlementTypeInstant   SettlementType = "instant"
	SettlementTypeEscrow    SettlementType = "escrow"
	SettlementTypeBatch     SettlementType = "batch"
	SettlementTypeRecurring SettlementType = "recurring"
)

// Settlement represents a stablecoin settlement between parties
type Settlement struct {
	// Unique identifier for the settlement
	Id uint64 `json:"id" yaml:"id"`

	// The type of settlement
	Type SettlementType `json:"type" yaml:"type"`

	// Sender's address (payer)
	Sender string `json:"sender" yaml:"sender"`

	// Recipient's address (merchant/payee)
	Recipient string `json:"recipient" yaml:"recipient"`

	// Amount to settle (in stablecoin)
	Amount sdk.Coin `json:"amount" yaml:"amount"`

	// Fee charged for the settlement
	Fee sdk.Coin `json:"fee" yaml:"fee"`

	// Net amount after fee
	NetAmount sdk.Coin `json:"net_amount" yaml:"net_amount"`

	// Current status
	Status SettlementStatus `json:"status" yaml:"status"`

	// External reference (order ID, invoice ID, etc.)
	Reference string `json:"reference" yaml:"reference"`

	// Metadata for the settlement
	Metadata string `json:"metadata" yaml:"metadata"`

	// Block height when created
	CreatedHeight int64 `json:"created_height" yaml:"created_height"`

	// Time when created
	CreatedTime time.Time `json:"created_time" yaml:"created_time"`

	// Block height when settled
	SettledHeight int64 `json:"settled_height" yaml:"settled_height"`

	// Time when settled
	SettledTime time.Time `json:"settled_time" yaml:"settled_time"`

	// Expiration time for escrow settlements
	ExpiresAt time.Time `json:"expires_at,omitempty" yaml:"expires_at,omitempty"`

	// Batch ID if part of a batch settlement
	BatchId uint64 `json:"batch_id,omitempty" yaml:"batch_id,omitempty"`
}

// ProtoMessage implements proto.Message
func (s *Settlement) ProtoMessage() {}

// Reset implements proto.Message
func (s *Settlement) Reset() { *s = Settlement{} }

// String implements proto.Message
func (s *Settlement) String() string {
	return fmt.Sprintf("Settlement{id:%d, type:%s, sender:%s, recipient:%s, amount:%s, status:%s}",
		s.Id, s.Type, s.Sender, s.Recipient, s.Amount.String(), s.Status)
}

// BatchSettlement represents a batch of settlements processed together
type BatchSettlement struct {
	// Unique identifier for the batch
	Id uint64 `json:"id" yaml:"id"`

	// Merchant receiving the batch settlement
	Merchant string `json:"merchant" yaml:"merchant"`

	// List of settlement IDs in this batch
	SettlementIds []uint64 `json:"settlement_ids" yaml:"settlement_ids"`

	// Total amount in the batch
	TotalAmount sdk.Coin `json:"total_amount" yaml:"total_amount"`

	// Total fees for the batch
	TotalFees sdk.Coin `json:"total_fees" yaml:"total_fees"`

	// Net amount after fees
	NetAmount sdk.Coin `json:"net_amount" yaml:"net_amount"`

	// Number of settlements in batch
	Count uint64 `json:"count" yaml:"count"`

	// Status of the batch
	Status SettlementStatus `json:"status" yaml:"status"`

	// Block height when created
	CreatedHeight int64 `json:"created_height" yaml:"created_height"`

	// Time when created
	CreatedTime time.Time `json:"created_time" yaml:"created_time"`

	// Block height when settled
	SettledHeight int64 `json:"settled_height" yaml:"settled_height"`

	// Time when settled
	SettledTime time.Time `json:"settled_time" yaml:"settled_time"`
}

// ProtoMessage implements proto.Message
func (b *BatchSettlement) ProtoMessage() {}

// Reset implements proto.Message
func (b *BatchSettlement) Reset() { *b = BatchSettlement{} }

// String implements proto.Message
func (b *BatchSettlement) String() string {
	return fmt.Sprintf("BatchSettlement{id:%d, merchant:%s, count:%d, total:%s, status:%s}",
		b.Id, b.Merchant, b.Count, b.TotalAmount.String(), b.Status)
}

// PaymentChannel represents a payment channel for streaming payments
type PaymentChannel struct {
	// Unique identifier for the channel
	Id uint64 `json:"id" yaml:"id"`

	// Sender (funding party)
	Sender string `json:"sender" yaml:"sender"`

	// Recipient (receiving party)
	Recipient string `json:"recipient" yaml:"recipient"`

	// Total amount deposited in channel
	Deposit sdk.Coin `json:"deposit" yaml:"deposit"`

	// Amount already spent/claimed
	Spent sdk.Coin `json:"spent" yaml:"spent"`

	// Remaining balance
	Balance sdk.Coin `json:"balance" yaml:"balance"`

	// Is the channel currently open
	IsOpen bool `json:"is_open" yaml:"is_open"`

	// Block height when opened
	OpenedHeight int64 `json:"opened_height" yaml:"opened_height"`

	// Time when opened
	OpenedTime time.Time `json:"opened_time" yaml:"opened_time"`

	// Block height when closed (if closed)
	ClosedHeight int64 `json:"closed_height,omitempty" yaml:"closed_height,omitempty"`

	// Time when closed (if closed)
	ClosedTime time.Time `json:"closed_time,omitempty" yaml:"closed_time,omitempty"`

	// Expiration block height
	ExpiresAtHeight int64 `json:"expires_at_height" yaml:"expires_at_height"`

	// Nonce for replay protection
	Nonce uint64 `json:"nonce" yaml:"nonce"`
}

// ProtoMessage implements proto.Message
func (c *PaymentChannel) ProtoMessage() {}

// Reset implements proto.Message
func (c *PaymentChannel) Reset() { *c = PaymentChannel{} }

// String implements proto.Message
func (c *PaymentChannel) String() string {
	return fmt.Sprintf("PaymentChannel{id:%d, sender:%s, recipient:%s, deposit:%s, balance:%s, open:%v}",
		c.Id, c.Sender, c.Recipient, c.Deposit.String(), c.Balance.String(), c.IsOpen)
}

// MerchantConfig represents merchant-specific configuration
type MerchantConfig struct {
	// Merchant's address
	Address string `json:"address" yaml:"address"`

	// Merchant's display name
	Name string `json:"name" yaml:"name"`

	// Fee rate for this merchant (in basis points, e.g., 30 = 0.30%)
	FeeRateBps uint32 `json:"fee_rate_bps" yaml:"fee_rate_bps"`

	// Minimum settlement amount
	MinSettlement sdk.Coin `json:"min_settlement" yaml:"min_settlement"`

	// Maximum settlement amount
	MaxSettlement sdk.Coin `json:"max_settlement" yaml:"max_settlement"`

	// Whether batch settlements are enabled
	BatchEnabled bool `json:"batch_enabled" yaml:"batch_enabled"`

	// Batch settlement threshold (amount)
	BatchThreshold sdk.Coin `json:"batch_threshold" yaml:"batch_threshold"`

	// Settlement delay (for escrow)
	SettlementDelay time.Duration `json:"settlement_delay" yaml:"settlement_delay"`

	// Is the merchant active
	IsActive bool `json:"is_active" yaml:"is_active"`

	// Webhook URL for notifications (stored but handled off-chain)
	WebhookUrl string `json:"webhook_url,omitempty" yaml:"webhook_url,omitempty"`

	// When the merchant was registered
	RegisteredAt time.Time `json:"registered_at" yaml:"registered_at"`
}

// ProtoMessage implements proto.Message
func (m *MerchantConfig) ProtoMessage() {}

// Reset implements proto.Message
func (m *MerchantConfig) Reset() { *m = MerchantConfig{} }

// String implements proto.Message
func (m *MerchantConfig) String() string {
	return fmt.Sprintf("MerchantConfig{address:%s, name:%s, feeRate:%d, active:%v}",
		m.Address, m.Name, m.FeeRateBps, m.IsActive)
}

// TransferReceipt represents the receipt of a transfer
type TransferReceipt struct {
	// Settlement ID
	SettlementId uint64 `json:"settlement_id" yaml:"settlement_id"`

	// Transaction hash
	TxHash string `json:"tx_hash" yaml:"tx_hash"`

	// Block height
	BlockHeight int64 `json:"block_height" yaml:"block_height"`

	// Timestamp
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`

	// Sender
	Sender string `json:"sender" yaml:"sender"`

	// Recipient
	Recipient string `json:"recipient" yaml:"recipient"`

	// Amount transferred
	Amount sdk.Coin `json:"amount" yaml:"amount"`

	// Fee paid
	Fee sdk.Coin `json:"fee" yaml:"fee"`

	// Reference
	Reference string `json:"reference" yaml:"reference"`
}

// ProtoMessage implements proto.Message
func (r *TransferReceipt) ProtoMessage() {}

// Reset implements proto.Message
func (r *TransferReceipt) Reset() { *r = TransferReceipt{} }

// String implements proto.Message
func (r *TransferReceipt) String() string {
	return fmt.Sprintf("TransferReceipt{settlementId:%d, txHash:%s, amount:%s}",
		r.SettlementId, r.TxHash, r.Amount.String())
}
