package keeper_test

import (
	"testing"

	"github.com/stateset/core/x/zkpverify/types"
)

func TestCircuitValidation(t *testing.T) {
	tests := []struct {
		name    string
		circuit types.Circuit
		wantErr bool
	}{
		{
			name: "valid circuit",
			circuit: types.Circuit{
				Name:            "test-circuit",
				VerificationKey: []byte("test-vk-data-at-least-32-bytes-long"),
				ProofSystem:     types.ProofSystemSTARK,
			},
			wantErr: false,
		},
		{
			name: "empty name",
			circuit: types.Circuit{
				Name:            "",
				VerificationKey: []byte("test-vk"),
				ProofSystem:     types.ProofSystemSTARK,
			},
			wantErr: true,
		},
		{
			name: "empty verification key",
			circuit: types.Circuit{
				Name:            "test-circuit",
				VerificationKey: []byte{},
				ProofSystem:     types.ProofSystemSTARK,
			},
			wantErr: true,
		},
		{
			name: "invalid proof system",
			circuit: types.Circuit{
				Name:            "test-circuit",
				VerificationKey: []byte("test-vk"),
				ProofSystem:     "groth16", // We only support STARK
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.circuit.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circuit.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSymbolicRuleValidation(t *testing.T) {
	tests := []struct {
		name    string
		rule    types.SymbolicRule
		wantErr bool
	}{
		{
			name: "valid rule",
			rule: types.SymbolicRule{
				Name:        "test-rule",
				CircuitName: "test-circuit",
				RuleType:    types.RuleTypeImplication,
				Conditions: []types.Condition{
					{Field: "amount", Operator: "gt", Value: "0"},
					{Field: "approved", Operator: "eq", Value: "true"},
				},
			},
			wantErr: false,
		},
		{
			name: "empty name",
			rule: types.SymbolicRule{
				Name:        "",
				CircuitName: "test-circuit",
				RuleType:    types.RuleTypeImplication,
				Conditions:  []types.Condition{{Field: "a", Operator: "eq", Value: "b"}},
			},
			wantErr: true,
		},
		{
			name: "empty circuit name",
			rule: types.SymbolicRule{
				Name:        "test-rule",
				CircuitName: "",
				RuleType:    types.RuleTypeImplication,
				Conditions:  []types.Condition{{Field: "a", Operator: "eq", Value: "b"}},
			},
			wantErr: true,
		},
		{
			name: "no conditions",
			rule: types.SymbolicRule{
				Name:        "test-rule",
				CircuitName: "test-circuit",
				RuleType:    types.RuleTypeImplication,
				Conditions:  []types.Condition{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("SymbolicRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProofValidation(t *testing.T) {
	tests := []struct {
		name    string
		proof   types.Proof
		wantErr bool
	}{
		{
			name: "valid proof",
			proof: types.Proof{
				CircuitName: "test-circuit",
				ProofData:   []byte("proof-data"),
			},
			wantErr: false,
		},
		{
			name: "empty circuit name",
			proof: types.Proof{
				CircuitName: "",
				ProofData:   []byte("proof-data"),
			},
			wantErr: true,
		},
		{
			name: "empty proof data",
			proof: types.Proof{
				CircuitName: "test-circuit",
				ProofData:   []byte{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.proof.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Proof.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultParams(t *testing.T) {
	params := types.DefaultParams()

	if params.MaxProofSize == 0 {
		t.Error("MaxProofSize should not be 0")
	}
	if params.MaxRecursionDepth == 0 {
		t.Error("MaxRecursionDepth should not be 0")
	}
	if params.ChallengeWindow == 0 {
		t.Error("ChallengeWindow should not be 0")
	}

	// Validate should pass for default params
	if err := params.Validate(); err != nil {
		t.Errorf("DefaultParams().Validate() error = %v", err)
	}
}

func TestGenesisValidation(t *testing.T) {
	// Default genesis should be valid
	genesis := types.DefaultGenesis()
	if err := genesis.Validate(); err != nil {
		t.Errorf("DefaultGenesis().Validate() error = %v", err)
	}

	// Genesis with circuits
	genesisWithCircuit := &types.GenesisState{
		Params: types.DefaultParams(),
		Circuits: []types.Circuit{
			{
				Name:            "circuit1",
				VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
				ProofSystem:     types.ProofSystemSTARK,
				Active:          true,
			},
		},
		SymbolicRules: []types.SymbolicRule{
			{
				Name:        "rule1",
				CircuitName: "circuit1",
				RuleType:    types.RuleTypeConjunction,
				Conditions:  []types.Condition{{Field: "x", Operator: "eq", Value: "1"}},
				Active:      true,
			},
		},
	}
	if err := genesisWithCircuit.Validate(); err != nil {
		t.Errorf("Genesis with circuit Validate() error = %v", err)
	}

	// Genesis with rule referencing non-existent circuit should fail
	genesisInvalid := &types.GenesisState{
		Params:   types.DefaultParams(),
		Circuits: []types.Circuit{},
		SymbolicRules: []types.SymbolicRule{
			{
				Name:        "rule1",
				CircuitName: "non-existent",
				RuleType:    types.RuleTypeConjunction,
				Conditions:  []types.Condition{{Field: "x", Operator: "eq", Value: "1"}},
			},
		},
	}
	if err := genesisInvalid.Validate(); err == nil {
		t.Error("Expected error for rule referencing non-existent circuit")
	}
}

func TestPublicInputs(t *testing.T) {
	pi := types.PublicInputs{
		Fields: map[string]interface{}{
			"amount":     float64(100),
			"hash":       "abc123def456",
			"is_valid":   true,
			"count":      uint64(42),
		},
	}

	// Test GetField
	if _, ok := pi.GetField("amount"); !ok {
		t.Error("GetField should find 'amount'")
	}
	if _, ok := pi.GetField("nonexistent"); ok {
		t.Error("GetField should not find 'nonexistent'")
	}

	// Test GetStringField
	if hash, ok := pi.GetStringField("hash"); !ok || hash != "abc123def456" {
		t.Errorf("GetStringField('hash') = %v, %v", hash, ok)
	}

	// Test GetUint64Field
	if count, ok := pi.GetUint64Field("count"); !ok || count != 42 {
		t.Errorf("GetUint64Field('count') = %v, %v", count, ok)
	}
}

func TestProofHash(t *testing.T) {
	proof := types.Proof{
		CircuitName:    "test-circuit",
		ProofData:      []byte("proof-data"),
		PublicInputs:   []byte(`{"fields": {}}`),
		DataCommitment: []byte("commitment"),
	}

	hash1 := proof.ComputeProofHash()
	hash2 := proof.ComputeProofHash()

	if hash1 != hash2 {
		t.Error("ComputeProofHash should be deterministic")
	}

	if len(hash1) != 64 { // SHA256 hex = 64 chars
		t.Errorf("ComputeProofHash should return 64 char hex, got %d", len(hash1))
	}

	// Different proof should have different hash
	proof2 := types.Proof{
		CircuitName:    "test-circuit-2",
		ProofData:      []byte("different-proof"),
		PublicInputs:   []byte(`{"fields": {}}`),
		DataCommitment: []byte("commitment"),
	}

	if proof.ComputeProofHash() == proof2.ComputeProofHash() {
		t.Error("Different proofs should have different hashes")
	}
}
