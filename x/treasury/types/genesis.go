package types

type GenesisState struct {
	Authority       string            `json:"authority" yaml:"authority"`
	NextID          uint64            `json:"next_id" yaml:"next_id"`
	NextProposalID  uint64            `json:"next_proposal_id" yaml:"next_proposal_id"`
	NextRevenueID   uint64            `json:"next_revenue_id" yaml:"next_revenue_id"`
	Params          TreasuryParams    `json:"params" yaml:"params"`
	Snapshots       []ReserveSnapshot `json:"snapshots" yaml:"snapshots"`
	SpendProposals  []SpendProposal   `json:"spend_proposals" yaml:"spend_proposals"`
	Budgets         []Budget          `json:"budgets" yaml:"budgets"`
	Allocations     []Allocation      `json:"allocations" yaml:"allocations"`
	RevenueRecords  []RevenueRecord   `json:"revenue_records" yaml:"revenue_records"`
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Authority:       "",
		NextID:          1,
		NextProposalID:  1,
		NextRevenueID:   1,
		Params:          DefaultTreasuryParams(),
		Snapshots:       []ReserveSnapshot{},
		SpendProposals:  []SpendProposal{},
		Budgets:         DefaultBudgets(),
		Allocations:     []Allocation{},
		RevenueRecords:  []RevenueRecord{},
	}
}

// DefaultBudgets returns default budgets for each category
func DefaultBudgets() []Budget {
	budgets := make([]Budget, 0, len(ValidCategories()))
	for _, cat := range ValidCategories() {
		budgets = append(budgets, Budget{
			Category:       cat,
			TotalLimit:     nil, // No total limit by default
			PeriodLimit:    nil, // No period limit by default
			PeriodSpent:    nil,
			TotalSpent:     nil,
			PeriodDuration: 30 * 24 * 60 * 60 * 1e9, // 30 days in nanoseconds
			Enabled:        true,
		})
	}
	return budgets
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	for _, snapshot := range gs.Snapshots {
		if err := snapshot.ValidateBasic(); err != nil {
			return err
		}
	}
	for _, proposal := range gs.SpendProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}
	for _, budget := range gs.Budgets {
		if err := budget.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
