package types

const (
	// ModuleName defines the module name
	ModuleName = "settlement"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// ModuleAccountName defines the account name for holding escrowed funds
	ModuleAccountName = "settlement"

	// StablecoinDenom is the default stablecoin denomination for settlements
	StablecoinDenom = "ssusd"
)

// Key prefixes for store keys
var (
	// SettlementKeyPrefix is the prefix for settlement storage
	SettlementKeyPrefix = []byte{0x01}

	// BatchKeyPrefix is the prefix for batch settlement storage
	BatchKeyPrefix = []byte{0x02}

	// NextSettlementIDKey is the key for the next settlement ID
	NextSettlementIDKey = []byte{0x03}

	// NextBatchIDKey is the key for the next batch ID
	NextBatchIDKey = []byte{0x04}

	// MerchantKeyPrefix is the prefix for merchant configurations
	MerchantKeyPrefix = []byte{0x05}

	// ChannelKeyPrefix is the prefix for payment channel storage
	ChannelKeyPrefix = []byte{0x06}

	// NextChannelIDKey is the key for the next channel ID
	NextChannelIDKey = []byte{0x07}

	// ParamsKey is the key for module parameters
	ParamsKey = []byte{0x08}

	// FeeCollectorKey is the key for accumulated fees
	FeeCollectorKey = []byte{0x09}
)

// Event types
const (
	EventTypeSettlementCreated   = "settlement_created"
	EventTypeSettlementCompleted = "settlement_completed"
	EventTypeSettlementFailed    = "settlement_failed"
	EventTypeSettlementRefunded  = "settlement_refunded"
	EventTypeBatchCreated        = "batch_created"
	EventTypeBatchSettled        = "batch_settled"
	EventTypeInstantTransfer     = "instant_transfer"
	EventTypeChannelOpened       = "channel_opened"
	EventTypeChannelClosed       = "channel_closed"
	EventTypeChannelUpdated      = "channel_updated"
	EventTypeFeeCollected        = "fee_collected"
	EventTypeEscrowExpired       = "escrow_expired"
	EventTypeChannelExpired      = "channel_expired"
)

// Event attribute keys
const (
	AttributeKeySettlementID = "settlement_id"
	AttributeKeyBatchID      = "batch_id"
	AttributeKeyChannelID    = "channel_id"
	AttributeKeySender       = "sender"
	AttributeKeyRecipient    = "recipient"
	AttributeKeyAmount       = "amount"
	AttributeKeyFee          = "fee"
	AttributeKeyStatus       = "status"
	AttributeKeyMerchant     = "merchant"
	AttributeKeyOrderID      = "order_id"
	AttributeKeyReference    = "reference"
)
