package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                DefaultParams(),
		CommerceTransactions:  []CommerceTransaction{},
		FinancialInstruments:  []FinancialInstrument{},
		// this line is used by starport scaffolding # genesis/types/default
		GlobalTradeStatistics: GlobalTradeStatistics{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in commerceTransaction
	commerceTransactionIndexMap := make(map[string]struct{})

	for _, elem := range gs.CommerceTransactions {
		index := string(CommerceTransactionKey(elem.ID))
		if _, ok := commerceTransactionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for commerceTransaction")
		}
		commerceTransactionIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in financialInstrument
	financialInstrumentIndexMap := make(map[string]struct{})

	for _, elem := range gs.FinancialInstruments {
		index := string(FinancialInstrumentKey(elem.ID))
		if _, ok := financialInstrumentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for financialInstrument")
		}
		financialInstrumentIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// GenesisState defines the commerce module's genesis state.
type GenesisState struct {
	Params                Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	CommerceTransactions  []CommerceTransaction   `protobuf:"bytes,2,rep,name=commerceTransactions,proto3" json:"commerceTransactions"`
	FinancialInstruments  []FinancialInstrument   `protobuf:"bytes,3,rep,name=financialInstruments,proto3" json:"financialInstruments"`
	GlobalTradeStatistics GlobalTradeStatistics   `protobuf:"bytes,4,opt,name=globalTradeStatistics,proto3" json:"globalTradeStatistics"`
	// this line is used by starport scaffolding # genesis/proto/state
}

// GlobalTradeStatistics represents the global trade statistics
type GlobalTradeStatistics struct {
	TotalTransactions    uint64  `json:"total_transactions"`
	TotalVolume          string  `json:"total_volume"`
	ActiveUsers          uint64  `json:"active_users"`
	LastUpdated          string  `json:"last_updated"`
}

// Params defines the parameters for the commerce module.
type Params struct {
	// Add module parameters here
	MaxTransactionSize     string `json:"max_transaction_size"`
	MaxTransactionsPerDay  uint64 `json:"max_transactions_per_day"`
	ComplianceEnabled      bool   `json:"compliance_enabled"`
	AnalyticsEnabled       bool   `json:"analytics_enabled"`
	FeeOptimizationEnabled bool   `json:"fee_optimization_enabled"`
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		MaxTransactionSize:     "1000000000000", // 1 trillion units
		MaxTransactionsPerDay:  10000,
		ComplianceEnabled:      true,
		AnalyticsEnabled:       true,
		FeeOptimizationEnabled: true,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.MaxTransactionsPerDay == 0 {
		return fmt.Errorf("max transactions per day must be positive")
	}
	return nil
}