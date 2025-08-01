package types

const (
	// ModuleName defines the module name
	ModuleName = "commerce"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_commerce"
)

const (
	CommerceTransactionKeyPrefix    = "CommerceTransaction/value/"
	FinancialInstrumentKeyPrefix    = "FinancialInstrument/value/"
	GlobalTradeStatisticsKeyPrefix  = "GlobalTradeStatistics/value/"
	PaymentRouteKeyPrefix           = "PaymentRoute/value/"
	ComplianceReportKeyPrefix       = "ComplianceReport/value/"
	AnalyticsDataKeyPrefix          = "AnalyticsData/value/"
)

// CommerceTransactionKey returns the store key to retrieve a CommerceTransaction from the index fields
func CommerceTransactionKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// FinancialInstrumentKey returns the store key to retrieve a FinancialInstrument from the index fields
func FinancialInstrumentKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// KeyPrefix returns the key prefix for a given prefix string
func KeyPrefix(p string) []byte {
	return []byte(p)
}

// Event types
const (
	EventTypeCommerceTransactionCreated      = "commerce_transaction_created"
	EventTypeCommerceTransactionCompleted    = "commerce_transaction_completed"
	EventTypeCommerceTransactionFailed       = "commerce_transaction_failed"
	EventTypeTradeFinanceInstrumentCreated   = "trade_finance_instrument_created"
	EventTypeTradeFinanceInstrumentUtilized  = "trade_finance_instrument_utilized"
	EventTypePaymentRouteOptimized           = "payment_route_optimized"
	EventTypeComplianceCheckCompleted        = "compliance_check_completed"
	EventTypeAnalyticsUpdated               = "analytics_updated"
)

// Attribute keys
const (
	AttributeKeyTransactionID       = "transaction_id"
	AttributeKeyTransactionType     = "transaction_type"
	AttributeKeyAmount              = "amount"
	AttributeKeyInstrumentID        = "instrument_id"
	AttributeKeyInstrumentType      = "instrument_type"
	AttributeKeyIssuer              = "issuer"
	AttributeKeyBeneficiary         = "beneficiary"
	AttributeKeyRouteType           = "route_type"
	AttributeKeyRouteCost           = "route_cost"
	AttributeKeyComplianceScore     = "compliance_score"
	AttributeKeyRiskLevel           = "risk_level"
)

// Error types
const (
	ErrTransactionNotFound        = "transaction not found"
	ErrInvalidTransactionStatus   = "invalid transaction status"
	ErrInstrumentNotFound         = "instrument not found"
	ErrInsufficientCollateral     = "insufficient collateral"
	ErrComplianceViolation        = "compliance violation"
	ErrRouteOptimizationFailed    = "route optimization failed"
	ErrInvalidPaymentRoute        = "invalid payment route"
)