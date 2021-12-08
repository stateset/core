package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		InvoiceList:         []Invoice{},
		SentInvoiceList:     []SentInvoice{},
		TimedoutInvoiceList: []TimedoutInvoice{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in invoice
	invoiceIdMap := make(map[uint64]bool)
	invoiceCount := gs.GetInvoiceCount()
	for _, elem := range gs.InvoiceList {
		if _, ok := invoiceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for invoice")
		}
		if elem.Id >= invoiceCount {
			return fmt.Errorf("invoice id should be lower or equal than the last id")
		}
		invoiceIdMap[elem.Id] = true
	}
	// Check for duplicated ID in sentInvoice
	sentInvoiceIdMap := make(map[uint64]bool)
	sentInvoiceCount := gs.GetSentInvoiceCount()
	for _, elem := range gs.SentInvoiceList {
		if _, ok := sentInvoiceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for sentInvoice")
		}
		if elem.Id >= sentInvoiceCount {
			return fmt.Errorf("sentInvoice id should be lower or equal than the last id")
		}
		sentInvoiceIdMap[elem.Id] = true
	}
	// Check for duplicated ID in timedoutInvoice
	timedoutInvoiceIdMap := make(map[uint64]bool)
	timedoutInvoiceCount := gs.GetTimedoutInvoiceCount()
	for _, elem := range gs.TimedoutInvoiceList {
		if _, ok := timedoutInvoiceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for timedoutInvoice")
		}
		if elem.Id >= timedoutInvoiceCount {
			return fmt.Errorf("timedoutInvoice id should be lower or equal than the last id")
		}
		timedoutInvoiceIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
