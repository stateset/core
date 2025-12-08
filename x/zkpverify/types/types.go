package types

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Circuit represents a registered STARK verification circuit
type Circuit struct {
	Name              string            `json:"name"`
	VerificationKey   []byte            `json:"verification_key"`
	ProofSystem       string            `json:"proof_system"` // "stark"
	PublicInputSchema []PublicInputField `json:"public_input_schema"`
	ConstraintHash    string            `json:"constraint_hash"`
	Owner             string            `json:"owner"`
	Active            bool              `json:"active"`
	CreatedAt         int64             `json:"created_at"`
	Description       string            `json:"description"`
	MaxRecursionDepth uint32            `json:"max_recursion_depth"`
}

// PublicInputField defines expected structure of public inputs
type PublicInputField struct {
	Name     string `json:"name"`
	Type     string `json:"type"` // "field", "hash", "uint64", "bytes"
	Required bool   `json:"required"`
}

// Validate performs basic validation on the circuit
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

// SymbolicRule defines a logical constraint that must be satisfied
type SymbolicRule struct {
	Name        string       `json:"name"`
	CircuitName string       `json:"circuit_name"`
	RuleType    RuleType     `json:"rule_type"`
	Conditions  []Condition  `json:"conditions"`
	Description string       `json:"description"`
	Active      bool         `json:"active"`
	CreatedAt   int64        `json:"created_at"`
}

// RuleType defines the type of symbolic rule
type RuleType string

const (
	RuleTypeImplication   RuleType = "implication"   // if A then B
	RuleTypeConjunction   RuleType = "conjunction"   // A and B and C
	RuleTypeDisjunction   RuleType = "disjunction"   // A or B or C
	RuleTypeNegation      RuleType = "negation"      // not A
	RuleTypeUniversal     RuleType = "universal"     // for all X: P(X)
	RuleTypeExistential   RuleType = "existential"   // exists X: P(X)
	RuleTypeEquality      RuleType = "equality"      // A == B
	RuleTypeInequality    RuleType = "inequality"    // A != B
	RuleTypeComparison    RuleType = "comparison"    // A < B, A > B, etc.
	RuleTypeSetMembership RuleType = "set_membership" // A in {x, y, z}
)

// Condition represents a single condition in a rule
type Condition struct {
	Field    string `json:"field"`    // Field name from public inputs
	Operator string `json:"operator"` // "eq", "neq", "gt", "lt", "gte", "lte", "in", "not_in"
	Value    string `json:"value"`    // Expected value or reference
	RefField string `json:"ref_field"` // Reference to another field (for field-to-field comparison)
}

// Validate performs basic validation on the rule
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

// Proof represents a submitted ZKP for verification
type Proof struct {
	ID              uint64   `json:"id"`
	CircuitName     string   `json:"circuit_name"`
	ProofData       []byte   `json:"proof_data"`
	PublicInputs    []byte   `json:"public_inputs"`
	DataCommitment  []byte   `json:"data_commitment"` // Merkle root of off-chain data
	RecursiveProofs []uint64 `json:"recursive_proofs"` // IDs of aggregated sub-proofs
	Submitter       string   `json:"submitter"`
	SubmittedAt     int64    `json:"submitted_at"`
	SubmittedHeight int64    `json:"submitted_height"`
}

// ComputeProofHash returns a hash of the proof for reference
func (p Proof) ComputeProofHash() string {
	h := sha256.New()
	h.Write([]byte(p.CircuitName))
	h.Write(p.ProofData)
	h.Write(p.PublicInputs)
	h.Write(p.DataCommitment)
	return hex.EncodeToString(h.Sum(nil))
}

// Validate performs basic validation on the proof
func (p Proof) Validate() error {
	if p.CircuitName == "" {
		return ErrEmptyCircuitName
	}
	if len(p.ProofData) == 0 {
		return ErrEmptyProofData
	}
	return nil
}

// VerificationResult stores the outcome of proof verification
type VerificationResult struct {
	ProofID              uint64   `json:"proof_id"`
	CircuitName          string   `json:"circuit_name"`
	Valid                bool     `json:"valid"`
	VerifiedAtHeight     int64    `json:"verified_at_height"`
	VerifiedAt           int64    `json:"verified_at"`
	DataCommitment       []byte   `json:"data_commitment"`
	ConstraintsSatisfied []string `json:"constraints_satisfied"`
	Error                string   `json:"error,omitempty"`
	VerificationTimeMs   int64    `json:"verification_time_ms"`
	RecursionDepth       uint32   `json:"recursion_depth"`
	Challenged           bool     `json:"challenged"`
	ChallengeDeadline    int64    `json:"challenge_deadline"`
}

// DataCommitmentRecord tracks off-chain data commitments (validium style)
type DataCommitmentRecord struct {
	Commitment      []byte `json:"commitment"`
	DataHash        []byte `json:"data_hash"`
	ProofID         uint64 `json:"proof_id"`
	CommittedAt     int64  `json:"committed_at"`
	CommittedHeight int64  `json:"committed_height"`
	DataURI         string `json:"data_uri,omitempty"` // Optional pointer to DA layer
}

// Params defines the module parameters
type Params struct {
	MaxProofSize          uint64        `json:"max_proof_size"`
	MaxPublicInputSize    uint64        `json:"max_public_input_size"`
	MaxRecursionDepth     uint32        `json:"max_recursion_depth"`
	ChallengeWindow       time.Duration `json:"challenge_window"`
	MinVerificationKeySize uint64       `json:"min_verification_key_size"`
	MaxCircuitsPerOwner   uint32        `json:"max_circuits_per_owner"`
	ProofSubmissionFee    string        `json:"proof_submission_fee"`
}

// DefaultParams returns default module parameters
func DefaultParams() Params {
	return Params{
		MaxProofSize:          1024 * 1024,     // 1MB
		MaxPublicInputSize:    64 * 1024,       // 64KB
		MaxRecursionDepth:     16,              // Max 16 levels of recursion
		ChallengeWindow:       24 * time.Hour,  // 24 hour challenge window
		MinVerificationKeySize: 32,             // Minimum VK size
		MaxCircuitsPerOwner:   100,             // Max circuits per owner
		ProofSubmissionFee:    "1000state",     // Fee for submitting proofs
	}
}

// Validate validates the parameters
func (p Params) Validate() error {
	if p.MaxProofSize == 0 {
		return ErrInvalidProof
	}
	if p.MaxRecursionDepth == 0 {
		return ErrMaxRecursionDepthExceeded
	}
	return nil
}

// PublicInputs represents decoded public inputs from a proof
type PublicInputs struct {
	Fields map[string]interface{} `json:"fields"`
}

// GetField retrieves a field value from public inputs
func (pi PublicInputs) GetField(name string) (interface{}, bool) {
	v, ok := pi.Fields[name]
	return v, ok
}

// GetStringField retrieves a string field value
func (pi PublicInputs) GetStringField(name string) (string, bool) {
	v, ok := pi.Fields[name]
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}

// GetUint64Field retrieves a uint64 field value
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
