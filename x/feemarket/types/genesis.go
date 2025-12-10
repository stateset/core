package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FeeHistoryEntry captures recent block usage for oracle suggestions.
type FeeHistoryEntry struct {
	Height      int64   `json:"height" yaml:"height"`
	GasUsed     uint64  `json:"gas_used" yaml:"gas_used"`
	BaseFee     sdk.Dec `json:"base_fee" yaml:"base_fee"`
	PriorityFee sdk.Dec `json:"priority_fee" yaml:"priority_fee"`
}

// GenesisState defines the module genesis state.
type GenesisState struct {
	Params     Params            `json:"params" yaml:"params"`
	BaseFee    sdk.Dec           `json:"base_fee" yaml:"base_fee"`
	FeeHistory []FeeHistoryEntry `json:"fee_history" yaml:"fee_history"`
}

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:     DefaultParams(),
		BaseFee:    DefaultInitialBaseFee,
		FeeHistory: []FeeHistoryEntry{},
	}
}

// Validate performs basic genesis validation.
func (gs GenesisState) Validate() error {
	if err := gs.Params.ValidateBasic(); err != nil {
		return err
	}
	if gs.BaseFee.IsNegative() {
		return ErrInvalidBaseFee
	}
	return nil
}
