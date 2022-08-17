package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RefundList: []Refund{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in refund
	refundIdMap := make(map[uint64]bool)
	refundCount := gs.GetRefundCount()
	for _, elem := range gs.RefundList {
		if _, ok := refundIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for refund")
		}
		if elem.Id >= refundCount {
			return fmt.Errorf("refund id should be lower or equal than the last id")
		}
		refundIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}