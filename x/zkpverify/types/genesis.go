package types

// GenesisState defines the zkpverify module's genesis state
type GenesisState struct {
	Params        Params         `json:"params"`
	Circuits      []Circuit      `json:"circuits"`
	SymbolicRules []SymbolicRule `json:"symbolic_rules"`
	ProofCount    uint64         `json:"proof_count"`
}

// ProtoMessage implementations for codec compatibility
func (gs *GenesisState) Reset()         { *gs = GenesisState{} }
func (gs *GenesisState) String() string { return "GenesisState" }
func (gs *GenesisState) ProtoMessage()  {}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:        DefaultParams(),
		Circuits:      []Circuit{},
		SymbolicRules: []SymbolicRule{},
		ProofCount:    0,
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	circuitNames := make(map[string]bool)
	for _, circuit := range gs.Circuits {
		if err := circuit.Validate(); err != nil {
			return err
		}
		if circuitNames[circuit.Name] {
			return ErrCircuitAlreadyExists
		}
		circuitNames[circuit.Name] = true
	}

	for _, rule := range gs.SymbolicRules {
		if err := rule.Validate(); err != nil {
			return err
		}
		if !circuitNames[rule.CircuitName] {
			return ErrCircuitNotFound
		}
	}

	return nil
}
