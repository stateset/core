package types

import (
	sdkmath "cosmossdk.io/math"
)

// FeeHistoryEntry captures recent block usage for oracle suggestions.
type FeeHistoryEntry struct {
	BlockHeight int64             `json:"block_height" yaml:"block_height"`
	GasUsed     uint64            `json:"gas_used" yaml:"gas_used"`
	GasLimit    uint64            `json:"gas_limit" yaml:"gas_limit"`
	BaseFee     sdkmath.LegacyDec `json:"base_fee" yaml:"base_fee"`
}

func (e *FeeHistoryEntry) Reset()         { *e = FeeHistoryEntry{} }
func (e *FeeHistoryEntry) String() string { return "FeeHistoryEntry" }
func (*FeeHistoryEntry) ProtoMessage()    {}

// GenesisState defines the module genesis state.
type GenesisState struct {
	Params     Params            `json:"params" yaml:"params"`
	BaseFee    sdkmath.LegacyDec `json:"base_fee" yaml:"base_fee"`
	FeeHistory []FeeHistoryEntry `json:"fee_history" yaml:"fee_history"`
}

func (gs *GenesisState) Reset()         { *gs = GenesisState{} }
func (gs *GenesisState) String() string { return "GenesisState" }
func (*GenesisState) ProtoMessage()     {}

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
