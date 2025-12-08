package types

const (
	ModuleName = "zkpverify"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
)

// Store key prefixes
var (
	CircuitKeyPrefix            = []byte{0x01}
	ProofKeyPrefix              = []byte{0x02}
	VerificationResultKeyPrefix = []byte{0x03}
	DataCommitmentKeyPrefix     = []byte{0x04}
	SymbolicRuleKeyPrefix       = []byte{0x05}
	ProofCountKey               = []byte{0x06}
	ParamsKey                   = []byte{0x07}
)

// Event types
const (
	EventTypeCircuitRegistered   = "circuit_registered"
	EventTypeCircuitDeactivated  = "circuit_deactivated"
	EventTypeProofSubmitted      = "proof_submitted"
	EventTypeProofVerified       = "proof_verified"
	EventTypeProofRejected       = "proof_rejected"
	EventTypeProofChallenged     = "proof_challenged"
	EventTypeDataCommitment      = "data_commitment"
	EventTypeRuleRegistered      = "rule_registered"
	EventTypeRecursiveAggregated = "recursive_aggregated"
)

// Event attribute keys
const (
	AttributeKeyCircuitName      = "circuit_name"
	AttributeKeyProofID          = "proof_id"
	AttributeKeyProofSystem      = "proof_system"
	AttributeKeySubmitter        = "submitter"
	AttributeKeyValid            = "valid"
	AttributeKeyError            = "error"
	AttributeKeyDataCommitment   = "data_commitment"
	AttributeKeyConstraintHash   = "constraint_hash"
	AttributeKeyRecursionDepth   = "recursion_depth"
	AttributeKeySubProofCount    = "sub_proof_count"
	AttributeKeyRuleName         = "rule_name"
	AttributeKeyRulesSatisfied   = "rules_satisfied"
	AttributeKeyVerificationTime = "verification_time_ms"
)

// Proof systems supported
const (
	ProofSystemSTARK = "stark"
)

// GetCircuitKey returns the store key for a circuit by name
func GetCircuitKey(name string) []byte {
	return append(CircuitKeyPrefix, []byte(name)...)
}

// GetProofKey returns the store key for a proof by ID
func GetProofKey(id uint64) []byte {
	return append(ProofKeyPrefix, uint64ToBytes(id)...)
}

// GetVerificationResultKey returns the store key for a verification result
func GetVerificationResultKey(proofID uint64) []byte {
	return append(VerificationResultKeyPrefix, uint64ToBytes(proofID)...)
}

// GetDataCommitmentKey returns the store key for a data commitment
func GetDataCommitmentKey(commitment []byte) []byte {
	return append(DataCommitmentKeyPrefix, commitment...)
}

// GetSymbolicRuleKey returns the store key for a symbolic rule
func GetSymbolicRuleKey(circuitName, ruleName string) []byte {
	key := append(SymbolicRuleKeyPrefix, []byte(circuitName)...)
	key = append(key, []byte("/")...)
	return append(key, []byte(ruleName)...)
}

// uint64ToBytes converts uint64 to big-endian bytes
func uint64ToBytes(n uint64) []byte {
	b := make([]byte, 8)
	b[0] = byte(n >> 56)
	b[1] = byte(n >> 48)
	b[2] = byte(n >> 40)
	b[3] = byte(n >> 32)
	b[4] = byte(n >> 24)
	b[5] = byte(n >> 16)
	b[6] = byte(n >> 8)
	b[7] = byte(n)
	return b
}
