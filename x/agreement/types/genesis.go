package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SentAgreementList:     []SentAgreement{},
		TimedoutAgreementList: []TimedoutAgreement{},
		AgreementList:         []Agreement{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in sentAgreement
	sentAgreementIdMap := make(map[uint64]bool)
	sentAgreementCount := gs.GetSentAgreementCount()
	for _, elem := range gs.SentAgreementList {
		if _, ok := sentAgreementIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for sentAgreement")
		}
		if elem.Id >= sentAgreementCount {
			return fmt.Errorf("sentAgreement id should be lower or equal than the last id")
		}
		sentAgreementIdMap[elem.Id] = true
	}
	// Check for duplicated ID in timedoutAgreement
	timedoutAgreementIdMap := make(map[uint64]bool)
	timedoutAgreementCount := gs.GetTimedoutAgreementCount()
	for _, elem := range gs.TimedoutAgreementList {
		if _, ok := timedoutAgreementIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for timedoutAgreement")
		}
		if elem.Id >= timedoutAgreementCount {
			return fmt.Errorf("timedoutAgreement id should be lower or equal than the last id")
		}
		timedoutAgreementIdMap[elem.Id] = true
	}
	// Check for duplicated ID in agreement
	agreementIdMap := make(map[uint64]bool)
	agreementCount := gs.GetAgreementCount()
	for _, elem := range gs.AgreementList {
		if _, ok := agreementIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for agreement")
		}
		if elem.Id >= agreementCount {
			return fmt.Errorf("agreement id should be lower or equal than the last id")
		}
		agreementIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
