package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PurchaseorderList:         []Purchaseorder{},
		SentPurchaseorderList:     []SentPurchaseorder{},
		TimedoutPurchaseorderList: []TimedoutPurchaseorder{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in purchaseorder
	purchaseorderIdMap := make(map[uint64]bool)
	purchaseorderCount := gs.GetPurchaseorderCount()
	for _, elem := range gs.PurchaseorderList {
		if _, ok := purchaseorderIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for purchaseorder")
		}
		if elem.Id >= purchaseorderCount {
			return fmt.Errorf("purchaseorder id should be lower or equal than the last id")
		}
		purchaseorderIdMap[elem.Id] = true
	}
	// Check for duplicated ID in sentPurchaseorder
	sentPurchaseorderIdMap := make(map[uint64]bool)
	sentPurchaseorderCount := gs.GetSentPurchaseorderCount()
	for _, elem := range gs.SentPurchaseorderList {
		if _, ok := sentPurchaseorderIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for sentPurchaseorder")
		}
		if elem.Id >= sentPurchaseorderCount {
			return fmt.Errorf("sentPurchaseorder id should be lower or equal than the last id")
		}
		sentPurchaseorderIdMap[elem.Id] = true
	}
	// Check for duplicated ID in timedoutPurchaseorder
	timedoutPurchaseorderIdMap := make(map[uint64]bool)
	timedoutPurchaseorderCount := gs.GetTimedoutPurchaseorderCount()
	for _, elem := range gs.TimedoutPurchaseorderList {
		if _, ok := timedoutPurchaseorderIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for timedoutPurchaseorder")
		}
		if elem.Id >= timedoutPurchaseorderCount {
			return fmt.Errorf("timedoutPurchaseorder id should be lower or equal than the last id")
		}
		timedoutPurchaseorderIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
