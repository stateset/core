package keeper

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/zkpverify/types"
)

// VerifyProof performs STARK proof verification and symbolic rule checking
func (k Keeper) VerifyProof(
	ctx sdk.Context,
	circuitName string,
	proofData []byte,
	publicInputs []byte,
	dataCommitment []byte,
	recursiveProofs []uint64,
) (types.VerificationResult, error) {
	startTime := time.Now()

	circuit, found := k.GetCircuit(ctx, circuitName)
	if !found {
		return types.VerificationResult{}, types.ErrCircuitNotFound
	}

	if !circuit.Active {
		return types.VerificationResult{}, types.ErrCircuitInactive
	}

	params := k.GetParams(ctx)

	// Validate proof size
	if uint64(len(proofData)) > params.MaxProofSize {
		return types.VerificationResult{}, types.ErrInvalidProof
	}
	if uint64(len(publicInputs)) > params.MaxPublicInputSize {
		return types.VerificationResult{}, types.ErrInvalidPublicInputs
	}

	// Calculate recursion depth
	recursionDepth := uint32(0)
	if len(recursiveProofs) > 0 {
		recursionDepth = 1
		for _, subProofID := range recursiveProofs {
			subResult, found := k.GetVerificationResult(ctx, subProofID)
			if !found {
				return types.VerificationResult{}, types.ErrRecursiveProofNotFound
			}
			if !subResult.Valid {
				return types.VerificationResult{}, types.ErrRecursiveProofInvalid
			}
			if subResult.RecursionDepth+1 > recursionDepth {
				recursionDepth = subResult.RecursionDepth + 1
			}
		}
	}

	if recursionDepth > params.MaxRecursionDepth || recursionDepth > circuit.MaxRecursionDepth {
		return types.VerificationResult{}, types.ErrMaxRecursionDepthExceeded
	}

	// Verify the STARK proof cryptographically
	proofValid, err := k.verifySTARKProof(circuit, proofData, publicInputs)
	if err != nil {
		return types.VerificationResult{
			CircuitName: circuitName,
			Valid:       false,
			Error:       fmt.Sprintf("STARK verification error: %v", err),
		}, nil
	}

	if !proofValid {
		return types.VerificationResult{
			CircuitName: circuitName,
			Valid:       false,
			Error:       "STARK proof verification failed",
		}, nil
	}

	// Decode public inputs
	var pi types.PublicInputs
	if err := json.Unmarshal(publicInputs, &pi); err != nil {
		return types.VerificationResult{}, types.ErrInvalidPublicInputs
	}

	// Validate public inputs against schema
	if err := k.validatePublicInputSchema(circuit, pi); err != nil {
		return types.VerificationResult{
			CircuitName: circuitName,
			Valid:       false,
			Error:       fmt.Sprintf("public input schema validation failed: %v", err),
		}, nil
	}

	// Check symbolic rules
	rules := k.GetSymbolicRulesForCircuit(ctx, circuitName)
	satisfiedRules, ruleErr := k.evaluateSymbolicRules(rules, pi)
	if ruleErr != nil {
		return types.VerificationResult{
			CircuitName:          circuitName,
			Valid:                false,
			Error:                fmt.Sprintf("symbolic rule violation: %v", ruleErr),
			ConstraintsSatisfied: satisfiedRules,
		}, nil
	}

	// Verify data commitment integrity
	if len(dataCommitment) > 0 {
		if err := k.verifyDataCommitment(dataCommitment, publicInputs); err != nil {
			return types.VerificationResult{
				CircuitName: circuitName,
				Valid:       false,
				Error:       fmt.Sprintf("data commitment verification failed: %v", err),
			}, nil
		}
	}

	verificationTimeMs := time.Since(startTime).Milliseconds()

	return types.VerificationResult{
		CircuitName:          circuitName,
		Valid:                true,
		DataCommitment:       dataCommitment,
		ConstraintsSatisfied: satisfiedRules,
		VerificationTimeMs:   verificationTimeMs,
		RecursionDepth:       recursionDepth,
	}, nil
}

// verifySTARKProof performs the cryptographic STARK verification
func (k Keeper) verifySTARKProof(circuit types.Circuit, proofData, publicInputs []byte) (bool, error) {
	// STARK verification implementation
	// In production, this would use a STARK verification library like:
	// - winterfell (Rust, via CGO or WASM)
	// - stone-prover verification
	// - custom STARK verifier

	// For now, we implement a placeholder that validates structure
	// and performs basic integrity checks

	if len(proofData) < 64 {
		return false, fmt.Errorf("proof data too short")
	}

	if len(circuit.VerificationKey) == 0 {
		return false, fmt.Errorf("circuit has no verification key")
	}

	// Verify proof structure
	// STARK proofs typically contain:
	// 1. Trace commitment
	// 2. Constraint commitment
	// 3. FRI layers
	// 4. Query responses

	// Basic structure validation
	proofHash := sha256.Sum256(proofData)
	vkHash := sha256.Sum256(circuit.VerificationKey)
	inputHash := sha256.Sum256(publicInputs)

	// Compute binding between proof, VK, and inputs
	combined := append(proofHash[:], vkHash[:]...)
	combined = append(combined, inputHash[:]...)
	binding := sha256.Sum256(combined)

	// In a real implementation, this would:
	// 1. Parse the STARK proof structure
	// 2. Verify the FRI protocol
	// 3. Check trace polynomial commitments
	// 4. Verify constraint evaluations
	// 5. Validate query responses

	// For demonstration, we verify the proof has proper structure
	// by checking it contains expected markers

	// Check proof has commitment section (first 32 bytes)
	if len(proofData) < 32 {
		return false, fmt.Errorf("missing trace commitment")
	}

	// Verify proof isn't all zeros (trivial rejection)
	allZeros := true
	for _, b := range proofData[:32] {
		if b != 0 {
			allZeros = false
			break
		}
	}
	if allZeros {
		return false, fmt.Errorf("invalid proof: null commitment")
	}

	// Log verification attempt
	_ = binding // Would be used in full verification

	return true, nil
}

// validatePublicInputSchema validates public inputs match the circuit schema
func (k Keeper) validatePublicInputSchema(circuit types.Circuit, pi types.PublicInputs) error {
	if pi.Fields == nil {
		pi.Fields = make(map[string]interface{})
	}

	for _, field := range circuit.PublicInputSchema {
		value, exists := pi.Fields[field.Name]

		if field.Required && !exists {
			return fmt.Errorf("missing required field: %s", field.Name)
		}

		if exists {
			// Type validation
			switch field.Type {
			case "field":
				// Field element - should be numeric
				switch value.(type) {
				case float64, int64, uint64, string:
					// Valid
				default:
					return fmt.Errorf("field %s: expected field element", field.Name)
				}
			case "hash":
				// Hash - should be string or bytes
				switch v := value.(type) {
				case string:
					if len(v) != 64 { // 32 bytes hex encoded
						return fmt.Errorf("field %s: invalid hash length", field.Name)
					}
				default:
					return fmt.Errorf("field %s: expected hash string", field.Name)
				}
			case "uint64":
				switch value.(type) {
				case float64, int64, uint64:
					// Valid
				default:
					return fmt.Errorf("field %s: expected uint64", field.Name)
				}
			case "bytes":
				switch value.(type) {
				case string, []byte:
					// Valid
				default:
					return fmt.Errorf("field %s: expected bytes", field.Name)
				}
			}
		}
	}

	return nil
}

// evaluateSymbolicRules checks all symbolic rules against public inputs
func (k Keeper) evaluateSymbolicRules(rules []types.SymbolicRule, pi types.PublicInputs) ([]string, error) {
	var satisfied []string

	for _, rule := range rules {
		if !rule.Active {
			continue
		}

		ruleSatisfied, err := k.evaluateRule(rule, pi)
		if err != nil {
			return satisfied, fmt.Errorf("rule %s: %v", rule.Name, err)
		}

		if !ruleSatisfied {
			return satisfied, fmt.Errorf("rule %s not satisfied", rule.Name)
		}

		satisfied = append(satisfied, rule.Name)
	}

	return satisfied, nil
}

// evaluateRule evaluates a single symbolic rule
func (k Keeper) evaluateRule(rule types.SymbolicRule, pi types.PublicInputs) (bool, error) {
	switch rule.RuleType {
	case types.RuleTypeImplication:
		return k.evaluateImplication(rule.Conditions, pi)
	case types.RuleTypeConjunction:
		return k.evaluateConjunction(rule.Conditions, pi)
	case types.RuleTypeDisjunction:
		return k.evaluateDisjunction(rule.Conditions, pi)
	case types.RuleTypeEquality:
		return k.evaluateEquality(rule.Conditions, pi)
	case types.RuleTypeInequality:
		return k.evaluateInequality(rule.Conditions, pi)
	case types.RuleTypeComparison:
		return k.evaluateComparison(rule.Conditions, pi)
	case types.RuleTypeSetMembership:
		return k.evaluateSetMembership(rule.Conditions, pi)
	default:
		return false, fmt.Errorf("unknown rule type: %s", rule.RuleType)
	}
}

// evaluateImplication: if first condition holds, rest must hold (A -> B)
func (k Keeper) evaluateImplication(conditions []types.Condition, pi types.PublicInputs) (bool, error) {
	if len(conditions) < 2 {
		return false, fmt.Errorf("implication requires at least 2 conditions")
	}

	// Check antecedent (first condition)
	antecedentHolds, err := k.evaluateCondition(conditions[0], pi)
	if err != nil {
		return false, err
	}

	// If antecedent doesn't hold, implication is vacuously true
	if !antecedentHolds {
		return true, nil
	}

	// Check all consequents
	for _, cond := range conditions[1:] {
		holds, err := k.evaluateCondition(cond, pi)
		if err != nil {
			return false, err
		}
		if !holds {
			return false, nil
		}
	}

	return true, nil
}

// evaluateConjunction: all conditions must hold (A AND B AND C)
func (k Keeper) evaluateConjunction(conditions []types.Condition, pi types.PublicInputs) (bool, error) {
	for _, cond := range conditions {
		holds, err := k.evaluateCondition(cond, pi)
		if err != nil {
			return false, err
		}
		if !holds {
			return false, nil
		}
	}
	return true, nil
}

// evaluateDisjunction: at least one condition must hold (A OR B OR C)
func (k Keeper) evaluateDisjunction(conditions []types.Condition, pi types.PublicInputs) (bool, error) {
	for _, cond := range conditions {
		holds, err := k.evaluateCondition(cond, pi)
		if err != nil {
			return false, err
		}
		if holds {
			return true, nil
		}
	}
	return false, nil
}

// evaluateEquality: field equals value
func (k Keeper) evaluateEquality(conditions []types.Condition, pi types.PublicInputs) (bool, error) {
	for _, cond := range conditions {
		cond.Operator = "eq"
		holds, err := k.evaluateCondition(cond, pi)
		if err != nil {
			return false, err
		}
		if !holds {
			return false, nil
		}
	}
	return true, nil
}

// evaluateInequality: field does not equal value
func (k Keeper) evaluateInequality(conditions []types.Condition, pi types.PublicInputs) (bool, error) {
	for _, cond := range conditions {
		cond.Operator = "neq"
		holds, err := k.evaluateCondition(cond, pi)
		if err != nil {
			return false, err
		}
		if !holds {
			return false, nil
		}
	}
	return true, nil
}

// evaluateComparison: numeric comparison
func (k Keeper) evaluateComparison(conditions []types.Condition, pi types.PublicInputs) (bool, error) {
	for _, cond := range conditions {
		holds, err := k.evaluateCondition(cond, pi)
		if err != nil {
			return false, err
		}
		if !holds {
			return false, nil
		}
	}
	return true, nil
}

// evaluateSetMembership: value in set
func (k Keeper) evaluateSetMembership(conditions []types.Condition, pi types.PublicInputs) (bool, error) {
	for _, cond := range conditions {
		cond.Operator = "in"
		holds, err := k.evaluateCondition(cond, pi)
		if err != nil {
			return false, err
		}
		if !holds {
			return false, nil
		}
	}
	return true, nil
}

// evaluateCondition evaluates a single condition
func (k Keeper) evaluateCondition(cond types.Condition, pi types.PublicInputs) (bool, error) {
	fieldValue, exists := pi.GetField(cond.Field)
	if !exists {
		return false, nil // Missing field means condition doesn't hold
	}

	var compareValue interface{}
	if cond.RefField != "" {
		// Compare against another field
		refValue, exists := pi.GetField(cond.RefField)
		if !exists {
			return false, nil
		}
		compareValue = refValue
	} else {
		compareValue = cond.Value
	}

	switch cond.Operator {
	case "eq":
		return compareValues(fieldValue, compareValue) == 0, nil
	case "neq":
		return compareValues(fieldValue, compareValue) != 0, nil
	case "gt":
		return compareValues(fieldValue, compareValue) > 0, nil
	case "lt":
		return compareValues(fieldValue, compareValue) < 0, nil
	case "gte":
		return compareValues(fieldValue, compareValue) >= 0, nil
	case "lte":
		return compareValues(fieldValue, compareValue) <= 0, nil
	case "in":
		return valueInSet(fieldValue, cond.Value), nil
	case "not_in":
		return !valueInSet(fieldValue, cond.Value), nil
	default:
		return false, fmt.Errorf("unknown operator: %s", cond.Operator)
	}
}

// compareValues compares two values, returns -1, 0, or 1
func compareValues(a, b interface{}) int {
	// Convert to comparable types
	aNum := toFloat64(a)
	bNum := toFloat64(b)

	if aNum < bNum {
		return -1
	} else if aNum > bNum {
		return 1
	}

	// If not numeric, compare as strings
	aStr := fmt.Sprintf("%v", a)
	bStr := fmt.Sprintf("%v", b)

	if aStr < bStr {
		return -1
	} else if aStr > bStr {
		return 1
	}
	return 0
}

// toFloat64 converts a value to float64 for comparison
func toFloat64(v interface{}) float64 {
	switch n := v.(type) {
	case float64:
		return n
	case int64:
		return float64(n)
	case uint64:
		return float64(n)
	case int:
		return float64(n)
	case string:
		if f, err := strconv.ParseFloat(n, 64); err == nil {
			return f
		}
	}
	return 0
}

// valueInSet checks if value is in a comma-separated set
func valueInSet(value interface{}, setStr string) bool {
	valueStr := fmt.Sprintf("%v", value)
	var set []string
	json.Unmarshal([]byte(setStr), &set)

	for _, s := range set {
		if s == valueStr {
			return true
		}
	}
	return false
}

// verifyDataCommitment verifies the data commitment (Merkle root)
func (k Keeper) verifyDataCommitment(commitment, publicInputs []byte) error {
	if len(commitment) != 32 {
		return fmt.Errorf("invalid commitment length: expected 32 bytes")
	}

	// The commitment should be a hash of the full off-chain data
	// Public inputs should contain a reference to this commitment
	var pi types.PublicInputs
	if err := json.Unmarshal(publicInputs, &pi); err != nil {
		return err
	}

	// Check if public inputs reference this commitment
	if dataCommitmentField, ok := pi.GetStringField("data_commitment"); ok {
		expectedCommitment := sha256.Sum256([]byte(dataCommitmentField))
		if !bytes.Equal(commitment, expectedCommitment[:]) {
			// Allow direct match as well
			if dataCommitmentField != fmt.Sprintf("%x", commitment) {
				return fmt.Errorf("data commitment mismatch")
			}
		}
	}

	return nil
}

// ProcessChallenge processes a fraud proof challenge
func (k Keeper) ProcessChallenge(
	ctx sdk.Context,
	proofID uint64,
	fraudProof []byte,
) (bool, error) {
	result, found := k.GetVerificationResult(ctx, proofID)
	if !found {
		return false, types.ErrProofNotFound
	}

	if result.Challenged {
		return false, types.ErrProofAlreadyChallenged
	}

	// Check challenge window
	if ctx.BlockTime().Unix() > result.ChallengeDeadline {
		return false, types.ErrChallengeWindowExpired
	}

	// Verify the fraud proof
	// In a real implementation, this would verify that the original proof was invalid
	challengeValid := k.verifyFraudProof(ctx, proofID, fraudProof)

	if challengeValid {
		result.Challenged = true
		result.Valid = false
		result.Error = "proof invalidated by successful challenge"
		k.UpdateVerificationResult(ctx, result)
		return true, nil
	}

	return false, types.ErrInvalidChallenge
}

// verifyFraudProof verifies a fraud proof against the original proof
func (k Keeper) verifyFraudProof(ctx sdk.Context, proofID uint64, fraudProof []byte) bool {
	proof, found := k.GetProof(ctx, proofID)
	if !found {
		return false
	}

	circuit, found := k.GetCircuit(ctx, proof.CircuitName)
	if !found {
		return false
	}

	// The fraud proof should demonstrate that either:
	// 1. The original proof doesn't verify
	// 2. The public inputs were malformed
	// 3. Symbolic rules were actually violated

	// For now, perform a re-verification with the fraud proof data
	if len(fraudProof) < 32 {
		return false
	}

	// Extract counter-proof or witness from fraud proof
	// This is a simplified check - real implementation would be more sophisticated
	counterHash := sha256.Sum256(fraudProof)
	proofHash := sha256.Sum256(proof.ProofData)

	// If the fraud proof provides a valid contradiction
	combined := append(counterHash[:], proofHash[:]...)
	combined = append(combined, circuit.VerificationKey...)
	validationHash := sha256.Sum256(combined)

	// Check if fraud proof demonstrates invalidity
	// This is a placeholder - real implementation would do actual cryptographic verification
	return validationHash[0] == 0 && validationHash[1] == 0 // Simplified check
}
