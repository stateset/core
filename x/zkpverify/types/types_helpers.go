package types

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Validate performs basic validation on the circuit.
func (c Circuit) Validate() error {
	if c.Name == "" {
		return ErrEmptyCircuitName
	}
	if len(c.VerificationKey) == 0 {
		return ErrInvalidVerificationKey
	}
	if c.ProofSystem != ProofSystemSTARK {
		return ErrInvalidProofSystem
	}
	return nil
}

// Validate performs basic validation on the rule.
func (r SymbolicRule) Validate() error {
	if r.Name == "" {
		return ErrInvalidRuleDefinition
	}
	if r.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	if len(r.Conditions) == 0 {
		return ErrInvalidRuleDefinition
	}
	return nil
}

// ComputeProofHash returns a hash of the proof for reference.
func (p Proof) ComputeProofHash() string {
	h := sha256.New()
	h.Write([]byte(p.CircuitName))
	h.Write(p.ProofData)
	h.Write(p.PublicInputs)
	h.Write(p.DataCommitment)
	return hex.EncodeToString(h.Sum(nil))
}

// Validate performs basic validation on the proof.
func (p Proof) Validate() error {
	if p.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	if len(p.ProofData) == 0 {
		return ErrEmptyProofData
	}
	return nil
}

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		MaxProofSize:           1024 * 1024,    // 1MB
		MaxPublicInputSize:     64 * 1024,      // 64KB
		MaxRecursionDepth:      16,             // Max 16 levels of recursion
		ChallengeWindow:        24 * time.Hour, // 24 hour challenge window
		MinVerificationKeySize: 32,             // Minimum VK size
		MaxCircuitsPerOwner:    100,            // Max circuits per owner
		ProofSubmissionFee:     "1000state",    // Fee for submitting proofs
	}
}

// Validate validates the parameters.
func (p Params) Validate() error {
	if p.MaxProofSize == 0 {
		return ErrInvalidProof
	}
	if p.MaxRecursionDepth == 0 {
		return ErrMaxRecursionDepthExceeded
	}
	return nil
}

// PublicInputs represents decoded public inputs from a proof.
type PublicInputs struct {
	Fields map[string]interface{} `json:"fields"`
}

// GetField retrieves a field value from public inputs.
func (pi PublicInputs) GetField(name string) (interface{}, bool) {
	v, ok := pi.Fields[name]
	return v, ok
}

// GetStringField retrieves a string field value.
func (pi PublicInputs) GetStringField(name string) (string, bool) {
	v, ok := pi.Fields[name]
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}

// GetUint64Field retrieves a uint64 field value.
func (pi PublicInputs) GetUint64Field(name string) (uint64, bool) {
	v, ok := pi.Fields[name]
	if !ok {
		return 0, false
	}
	switch n := v.(type) {
	case uint64:
		return n, true
	case float64:
		return uint64(n), true
	case int64:
		return uint64(n), true
	default:
		return 0, false
	}
}
