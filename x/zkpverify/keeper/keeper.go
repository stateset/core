package keeper

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/zkpverify/types"
)

// Keeper maintains the zkpverify module state
type Keeper struct {
	storeKey  storetypes.StoreKey
	cdc       codec.BinaryCodec
	authority string
}

// NewKeeper creates a new zkpverify keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	authority string,
) Keeper {
	return Keeper{
		storeKey:  key,
		cdc:       cdc,
		authority: authority,
	}
}

// GetAuthority returns the module authority address
func (k Keeper) GetAuthority() string {
	return k.authority
}

// ----------------------------------------------------------------------------
// Parameters
// ----------------------------------------------------------------------------

// GetParams returns the module parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return types.DefaultParams()
	}
	var params types.Params
	json.Unmarshal(bz, &params)
	return params
}

// SetParams sets the module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	bz, _ := json.Marshal(params)
	store.Set(types.ParamsKey, bz)
	return nil
}

// ----------------------------------------------------------------------------
// Circuits
// ----------------------------------------------------------------------------

// RegisterCircuit registers a new verification circuit
func (k Keeper) RegisterCircuit(ctx sdk.Context, circuit types.Circuit) error {
	if err := circuit.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetCircuitKey(circuit.Name)

	if store.Has(key) {
		return types.ErrCircuitAlreadyExists
	}

	circuit.CreatedAt = ctx.BlockTime().Unix()
	circuit.Active = true
	circuit.ProofSystem = types.ProofSystemSTARK

	bz, _ := json.Marshal(circuit)
	store.Set(key, bz)
	return nil
}

// GetCircuit retrieves a circuit by name
func (k Keeper) GetCircuit(ctx sdk.Context, name string) (types.Circuit, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetCircuitKey(name))
	if bz == nil {
		return types.Circuit{}, false
	}
	var circuit types.Circuit
	json.Unmarshal(bz, &circuit)
	return circuit, true
}

// DeactivateCircuit deactivates a circuit
func (k Keeper) DeactivateCircuit(ctx sdk.Context, name string) error {
	circuit, found := k.GetCircuit(ctx, name)
	if !found {
		return types.ErrCircuitNotFound
	}

	circuit.Active = false

	store := ctx.KVStore(k.storeKey)
	bz, _ := json.Marshal(circuit)
	store.Set(types.GetCircuitKey(name), bz)
	return nil
}

// GetAllCircuits returns all registered circuits
func (k Keeper) GetAllCircuits(ctx sdk.Context) []types.Circuit {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, types.CircuitKeyPrefix)
	defer iterator.Close()

	var circuits []types.Circuit
	for ; iterator.Valid(); iterator.Next() {
		var circuit types.Circuit
		json.Unmarshal(iterator.Value(), &circuit)
		circuits = append(circuits, circuit)
	}
	return circuits
}

// ----------------------------------------------------------------------------
// Symbolic Rules
// ----------------------------------------------------------------------------

// RegisterSymbolicRule registers a symbolic logic rule
func (k Keeper) RegisterSymbolicRule(ctx sdk.Context, rule types.SymbolicRule) error {
	if err := rule.Validate(); err != nil {
		return err
	}

	// Verify circuit exists
	_, found := k.GetCircuit(ctx, rule.CircuitName)
	if !found {
		return types.ErrCircuitNotFound
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetSymbolicRuleKey(rule.CircuitName, rule.Name)

	if store.Has(key) {
		return types.ErrRuleAlreadyExists
	}

	rule.CreatedAt = ctx.BlockTime().Unix()
	rule.Active = true

	bz, _ := json.Marshal(rule)
	store.Set(key, bz)
	return nil
}

// GetSymbolicRule retrieves a symbolic rule
func (k Keeper) GetSymbolicRule(ctx sdk.Context, circuitName, ruleName string) (types.SymbolicRule, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetSymbolicRuleKey(circuitName, ruleName))
	if bz == nil {
		return types.SymbolicRule{}, false
	}
	var rule types.SymbolicRule
	json.Unmarshal(bz, &rule)
	return rule, true
}

// GetSymbolicRulesForCircuit returns all rules for a circuit
func (k Keeper) GetSymbolicRulesForCircuit(ctx sdk.Context, circuitName string) []types.SymbolicRule {
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.SymbolicRuleKeyPrefix, []byte(circuitName+"/")...)
	iterator := storetypes.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	var rules []types.SymbolicRule
	for ; iterator.Valid(); iterator.Next() {
		var rule types.SymbolicRule
		json.Unmarshal(iterator.Value(), &rule)
		if rule.Active {
			rules = append(rules, rule)
		}
	}
	return rules
}

// ----------------------------------------------------------------------------
// Proofs
// ----------------------------------------------------------------------------

// getNextProofID returns the next proof ID and increments the counter
func (k Keeper) getNextProofID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ProofCountKey)
	var id uint64
	if bz != nil {
		id = binary.BigEndian.Uint64(bz)
	}
	id++
	newBz := make([]byte, 8)
	binary.BigEndian.PutUint64(newBz, id)
	store.Set(types.ProofCountKey, newBz)
	return id
}

// StoreProof stores a proof submission
func (k Keeper) StoreProof(ctx sdk.Context, proof *types.Proof) uint64 {
	proof.ID = k.getNextProofID(ctx)
	proof.SubmittedAt = ctx.BlockTime().Unix()
	proof.SubmittedHeight = ctx.BlockHeight()

	store := ctx.KVStore(k.storeKey)
	bz, _ := json.Marshal(proof)
	store.Set(types.GetProofKey(proof.ID), bz)
	return proof.ID
}

// GetProof retrieves a proof by ID
func (k Keeper) GetProof(ctx sdk.Context, id uint64) (types.Proof, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetProofKey(id))
	if bz == nil {
		return types.Proof{}, false
	}
	var proof types.Proof
	json.Unmarshal(bz, &proof)
	return proof, true
}

// ----------------------------------------------------------------------------
// Verification Results
// ----------------------------------------------------------------------------

// StoreVerificationResult stores a verification result
func (k Keeper) StoreVerificationResult(ctx sdk.Context, result types.VerificationResult) {
	params := k.GetParams(ctx)
	result.VerifiedAtHeight = ctx.BlockHeight()
	result.VerifiedAt = ctx.BlockTime().Unix()
	result.ChallengeDeadline = ctx.BlockTime().Add(params.ChallengeWindow).Unix()

	store := ctx.KVStore(k.storeKey)
	bz, _ := json.Marshal(result)
	store.Set(types.GetVerificationResultKey(result.ProofID), bz)
}

// GetVerificationResult retrieves a verification result
func (k Keeper) GetVerificationResult(ctx sdk.Context, proofID uint64) (types.VerificationResult, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVerificationResultKey(proofID))
	if bz == nil {
		return types.VerificationResult{}, false
	}
	var result types.VerificationResult
	json.Unmarshal(bz, &result)
	return result, true
}

// UpdateVerificationResult updates an existing result (for challenges)
func (k Keeper) UpdateVerificationResult(ctx sdk.Context, result types.VerificationResult) {
	store := ctx.KVStore(k.storeKey)
	bz, _ := json.Marshal(result)
	store.Set(types.GetVerificationResultKey(result.ProofID), bz)
}

// IsProofValid checks if a proof is valid and not challenged
func (k Keeper) IsProofValid(ctx sdk.Context, proofID uint64) bool {
	result, found := k.GetVerificationResult(ctx, proofID)
	if !found {
		return false
	}
	return result.Valid && !result.Challenged
}

// IsProofFinalized checks if a proof is past its challenge window
func (k Keeper) IsProofFinalized(ctx sdk.Context, proofID uint64) bool {
	result, found := k.GetVerificationResult(ctx, proofID)
	if !found {
		return false
	}
	return result.Valid && !result.Challenged && ctx.BlockTime().Unix() > result.ChallengeDeadline
}

// ----------------------------------------------------------------------------
// Data Commitments
// ----------------------------------------------------------------------------

// StoreDataCommitment stores a data commitment record
func (k Keeper) StoreDataCommitment(ctx sdk.Context, record types.DataCommitmentRecord) {
	record.CommittedAt = ctx.BlockTime().Unix()
	record.CommittedHeight = ctx.BlockHeight()

	store := ctx.KVStore(k.storeKey)
	bz, _ := json.Marshal(record)
	store.Set(types.GetDataCommitmentKey(record.Commitment), bz)
}

// GetDataCommitment retrieves a data commitment record
func (k Keeper) GetDataCommitment(ctx sdk.Context, commitment []byte) (types.DataCommitmentRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetDataCommitmentKey(commitment))
	if bz == nil {
		return types.DataCommitmentRecord{}, false
	}
	var record types.DataCommitmentRecord
	json.Unmarshal(bz, &record)
	return record, true
}

// ----------------------------------------------------------------------------
// Genesis
// ----------------------------------------------------------------------------

// InitGenesis initializes the module state from genesis
func (k Keeper) InitGenesis(ctx sdk.Context, gs *types.GenesisState) error {
	if err := k.SetParams(ctx, gs.Params); err != nil {
		return err
	}

	for _, circuit := range gs.Circuits {
		if err := k.RegisterCircuit(ctx, circuit); err != nil {
			return err
		}
	}

	for _, rule := range gs.SymbolicRules {
		if err := k.RegisterSymbolicRule(ctx, rule); err != nil {
			return err
		}
	}

	// Set proof count
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, gs.ProofCount)
	store.Set(types.ProofCountKey, bz)

	return nil
}

// ExportGenesis exports the module state to genesis
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ProofCountKey)
	var proofCount uint64
	if bz != nil {
		proofCount = binary.BigEndian.Uint64(bz)
	}

	// Get all symbolic rules
	var allRules []types.SymbolicRule
	iterator := storetypes.KVStorePrefixIterator(store, types.SymbolicRuleKeyPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var rule types.SymbolicRule
		json.Unmarshal(iterator.Value(), &rule)
		allRules = append(allRules, rule)
	}

	return &types.GenesisState{
		Params:        k.GetParams(ctx),
		Circuits:      k.GetAllCircuits(ctx),
		SymbolicRules: allRules,
		ProofCount:    proofCount,
	}
}

// ----------------------------------------------------------------------------
// Helper: Get current time for external use
// ----------------------------------------------------------------------------

func (k Keeper) GetBlockTime(ctx sdk.Context) time.Time {
	return ctx.BlockTime()
}
