package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/stateset/core/x/zkpverify/keeper"
	"github.com/stateset/core/x/zkpverify/types"
)

func setupKeeperTest(t *testing.T) (keeper.Keeper, sdk.Context) {
	t.Helper()

	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(cdc, storeKey, "stateset1authority")

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	return k, ctx
}

func TestRegisterCircuit(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	circuit := types.Circuit{
		Name:            "test-circuit",
		VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
		ProofSystem:     types.ProofSystemSTARK,
		Description:     "Test circuit for unit tests",
	}

	err := k.RegisterCircuit(ctx, circuit)
	require.NoError(t, err)

	// Verify circuit was stored
	stored, found := k.GetCircuit(ctx, "test-circuit")
	require.True(t, found)
	require.Equal(t, "test-circuit", stored.Name)
	require.True(t, stored.Active)
	require.Equal(t, types.ProofSystemSTARK, stored.ProofSystem)
	require.NotZero(t, stored.CreatedAt)
}

func TestRegisterCircuitDuplicate(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	circuit := types.Circuit{
		Name:            "test-circuit",
		VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
		ProofSystem:     types.ProofSystemSTARK,
	}

	err := k.RegisterCircuit(ctx, circuit)
	require.NoError(t, err)

	// Duplicate should fail
	err = k.RegisterCircuit(ctx, circuit)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrCircuitAlreadyExists)
}

func TestRegisterCircuitInvalid(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	testCases := []struct {
		name    string
		circuit types.Circuit
	}{
		{
			name: "empty name",
			circuit: types.Circuit{
				Name:            "",
				VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
				ProofSystem:     types.ProofSystemSTARK,
			},
		},
		{
			name: "empty verification key",
			circuit: types.Circuit{
				Name:            "test",
				VerificationKey: []byte{},
				ProofSystem:     types.ProofSystemSTARK,
			},
		},
		{
			name: "invalid proof system",
			circuit: types.Circuit{
				Name:            "test",
				VerificationKey: []byte("vk-data"),
				ProofSystem:     "invalid",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := k.RegisterCircuit(ctx, tc.circuit)
			require.Error(t, err)
		})
	}
}

func TestDeactivateCircuit(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	circuit := types.Circuit{
		Name:            "test-circuit",
		VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
		ProofSystem:     types.ProofSystemSTARK,
	}

	err := k.RegisterCircuit(ctx, circuit)
	require.NoError(t, err)

	// Deactivate
	err = k.DeactivateCircuit(ctx, "test-circuit")
	require.NoError(t, err)

	// Verify deactivated
	stored, found := k.GetCircuit(ctx, "test-circuit")
	require.True(t, found)
	require.False(t, stored.Active)
}

func TestDeactivateCircuitNotFound(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	err := k.DeactivateCircuit(ctx, "nonexistent")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrCircuitNotFound)
}

func TestGetAllCircuits(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	circuits := []types.Circuit{
		{Name: "circuit-1", VerificationKey: []byte("vk-data-at-least-32-bytes-long-1"), ProofSystem: types.ProofSystemSTARK},
		{Name: "circuit-2", VerificationKey: []byte("vk-data-at-least-32-bytes-long-2"), ProofSystem: types.ProofSystemSTARK},
		{Name: "circuit-3", VerificationKey: []byte("vk-data-at-least-32-bytes-long-3"), ProofSystem: types.ProofSystemSTARK},
	}

	for _, c := range circuits {
		err := k.RegisterCircuit(ctx, c)
		require.NoError(t, err)
	}

	all := k.GetAllCircuits(ctx)
	require.Len(t, all, 3)
}

func TestRegisterSymbolicRule(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	// First register a circuit
	circuit := types.Circuit{
		Name:            "test-circuit",
		VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
		ProofSystem:     types.ProofSystemSTARK,
	}
	err := k.RegisterCircuit(ctx, circuit)
	require.NoError(t, err)

	// Register a rule
	rule := types.SymbolicRule{
		Name:        "test-rule",
		CircuitName: "test-circuit",
		RuleType:    types.RuleTypeImplication,
		Conditions: []types.Condition{
			{Field: "amount", Operator: "gt", Value: "0"},
		},
		Description: "Test rule",
	}

	err = k.RegisterSymbolicRule(ctx, rule)
	require.NoError(t, err)

	// Verify rule was stored
	stored, found := k.GetSymbolicRule(ctx, "test-circuit", "test-rule")
	require.True(t, found)
	require.Equal(t, "test-rule", stored.Name)
	require.True(t, stored.Active)
}

func TestRegisterSymbolicRuleCircuitNotFound(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	rule := types.SymbolicRule{
		Name:        "test-rule",
		CircuitName: "nonexistent",
		RuleType:    types.RuleTypeImplication,
		Conditions:  []types.Condition{{Field: "a", Operator: "eq", Value: "b"}},
	}

	err := k.RegisterSymbolicRule(ctx, rule)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrCircuitNotFound)
}

func TestRegisterSymbolicRuleDuplicate(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	circuit := types.Circuit{
		Name:            "test-circuit",
		VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
		ProofSystem:     types.ProofSystemSTARK,
	}
	k.RegisterCircuit(ctx, circuit)

	rule := types.SymbolicRule{
		Name:        "test-rule",
		CircuitName: "test-circuit",
		RuleType:    types.RuleTypeImplication,
		Conditions:  []types.Condition{{Field: "a", Operator: "eq", Value: "b"}},
	}

	err := k.RegisterSymbolicRule(ctx, rule)
	require.NoError(t, err)

	// Duplicate should fail
	err = k.RegisterSymbolicRule(ctx, rule)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrRuleAlreadyExists)
}

func TestGetSymbolicRulesForCircuit(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	// Register circuit
	circuit := types.Circuit{
		Name:            "test-circuit",
		VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
		ProofSystem:     types.ProofSystemSTARK,
	}
	k.RegisterCircuit(ctx, circuit)

	// Register multiple rules
	for i := 1; i <= 3; i++ {
		rule := types.SymbolicRule{
			Name:        "rule-" + string(rune('0'+i)),
			CircuitName: "test-circuit",
			RuleType:    types.RuleTypeConjunction,
			Conditions:  []types.Condition{{Field: "x", Operator: "eq", Value: "1"}},
		}
		err := k.RegisterSymbolicRule(ctx, rule)
		require.NoError(t, err)
	}

	rules := k.GetSymbolicRulesForCircuit(ctx, "test-circuit")
	require.Len(t, rules, 3)
}

func TestStoreAndGetProof(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	proof := &types.Proof{
		CircuitName:    "test-circuit",
		ProofData:      []byte("proof-data-bytes"),
		PublicInputs:   []byte(`{"fields": {"amount": 100}}`),
		DataCommitment: []byte("commitment-hash"),
		Submitter:      "stateset1submitter",
	}

	id := k.StoreProof(ctx, proof)
	require.Equal(t, uint64(1), id)
	require.Equal(t, uint64(1), proof.ID)
	require.NotZero(t, proof.SubmittedAt)
	require.NotZero(t, proof.SubmittedHeight)

	// Retrieve proof
	stored, found := k.GetProof(ctx, 1)
	require.True(t, found)
	require.Equal(t, "test-circuit", stored.CircuitName)
	require.Equal(t, []byte("proof-data-bytes"), stored.ProofData)
	require.Equal(t, "stateset1submitter", stored.Submitter)
}

func TestStoreMultipleProofs(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	for i := 1; i <= 5; i++ {
		proof := &types.Proof{
			CircuitName:  "test-circuit",
			ProofData:    []byte("proof-data"),
			PublicInputs: []byte(`{}`),
		}
		id := k.StoreProof(ctx, proof)
		require.Equal(t, uint64(i), id)
	}

	// Verify all proofs exist
	for i := uint64(1); i <= 5; i++ {
		_, found := k.GetProof(ctx, i)
		require.True(t, found)
	}
}

func TestGetProofNotFound(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	_, found := k.GetProof(ctx, 999)
	require.False(t, found)
}

func TestStoreAndGetVerificationResult(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	// Store a proof first
	proof := &types.Proof{
		CircuitName:  "test-circuit",
		ProofData:    []byte("proof"),
		PublicInputs: []byte(`{}`),
	}
	proofID := k.StoreProof(ctx, proof)

	// Store verification result
	result := types.VerificationResult{
		ProofID:            proofID,
		Valid:              true,
		CircuitName:        "test-circuit",
		VerificationTimeMs: 100,
	}
	k.StoreVerificationResult(ctx, result)

	// Retrieve result
	stored, found := k.GetVerificationResult(ctx, proofID)
	require.True(t, found)
	require.True(t, stored.Valid)
	require.Equal(t, "test-circuit", stored.CircuitName)
	require.NotZero(t, stored.VerifiedAt)
	require.NotZero(t, stored.ChallengeDeadline)
}

func TestIsProofValid(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	proof := &types.Proof{
		CircuitName:  "test-circuit",
		ProofData:    []byte("proof"),
		PublicInputs: []byte(`{}`),
	}
	proofID := k.StoreProof(ctx, proof)

	// Before verification result, should be invalid
	require.False(t, k.IsProofValid(ctx, proofID))

	// Store valid result
	result := types.VerificationResult{
		ProofID:     proofID,
		Valid:       true,
		CircuitName: "test-circuit",
	}
	k.StoreVerificationResult(ctx, result)

	// Now should be valid
	require.True(t, k.IsProofValid(ctx, proofID))

	// Mark as challenged
	result.Challenged = true
	k.UpdateVerificationResult(ctx, result)

	// Should be invalid when challenged
	require.False(t, k.IsProofValid(ctx, proofID))
}

func TestIsProofFinalized(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	// Set params with short challenge window
	params := types.DefaultParams()
	params.ChallengeWindow = 60 * time.Second
	k.SetParams(ctx, params)

	proof := &types.Proof{
		CircuitName:  "test-circuit",
		ProofData:    []byte("proof"),
		PublicInputs: []byte(`{}`),
	}
	proofID := k.StoreProof(ctx, proof)

	result := types.VerificationResult{
		ProofID:     proofID,
		Valid:       true,
		CircuitName: "test-circuit",
	}
	k.StoreVerificationResult(ctx, result)

	// Not finalized yet (within challenge window)
	require.False(t, k.IsProofFinalized(ctx, proofID))

	// Move time forward past challenge window
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(61 * time.Second))

	// Now should be finalized
	require.True(t, k.IsProofFinalized(ctx, proofID))
}

func TestStoreAndGetDataCommitment(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	commitment := []byte("sha256-commitment-hash-here")
	record := types.DataCommitmentRecord{
		Commitment: commitment,
		DataHash:   []byte("data-hash"),
		DataURI:    "https://example.com/data",
	}

	k.StoreDataCommitment(ctx, record)

	// Retrieve commitment
	stored, found := k.GetDataCommitment(ctx, commitment)
	require.True(t, found)
	require.Equal(t, []byte("data-hash"), stored.DataHash)
	require.Equal(t, "https://example.com/data", stored.DataURI)
	require.NotZero(t, stored.CommittedAt)
	require.NotZero(t, stored.CommittedHeight)
}

func TestGetDataCommitmentNotFound(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	_, found := k.GetDataCommitment(ctx, []byte("nonexistent"))
	require.False(t, found)
}

func TestParams(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	// Default params
	params := k.GetParams(ctx)
	require.NotZero(t, params.MaxProofSize)

	// Update params
	newParams := types.Params{
		MaxProofSize:      2048,
		MaxRecursionDepth: 5,
		ChallengeWindow:   time.Hour,
	}
	err := k.SetParams(ctx, newParams)
	require.NoError(t, err)

	// Verify update
	stored := k.GetParams(ctx)
	require.Equal(t, uint64(2048), stored.MaxProofSize)
	require.Equal(t, uint32(5), stored.MaxRecursionDepth)
}

func TestParamsValidation(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	// Invalid params should fail
	invalidParams := types.Params{
		MaxProofSize:      0, // Invalid
		MaxRecursionDepth: 5,
		ChallengeWindow:   time.Hour,
	}
	err := k.SetParams(ctx, invalidParams)
	require.Error(t, err)
}

func TestGenesisExportImport(t *testing.T) {
	k, ctx := setupKeeperTest(t)

	// Set up state
	circuit := types.Circuit{
		Name:            "test-circuit",
		VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
		ProofSystem:     types.ProofSystemSTARK,
	}
	k.RegisterCircuit(ctx, circuit)

	rule := types.SymbolicRule{
		Name:        "test-rule",
		CircuitName: "test-circuit",
		RuleType:    types.RuleTypeImplication,
		Conditions:  []types.Condition{{Field: "x", Operator: "eq", Value: "1"}},
	}
	k.RegisterSymbolicRule(ctx, rule)

	proof := &types.Proof{
		CircuitName:  "test-circuit",
		ProofData:    []byte("proof"),
		PublicInputs: []byte(`{}`),
	}
	k.StoreProof(ctx, proof)

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.Len(t, genesis.Circuits, 1)
	require.Len(t, genesis.SymbolicRules, 1)
	require.Equal(t, uint64(1), genesis.ProofCount)

	// Create new keeper and import
	k2, ctx2 := setupKeeperTest(t)
	err := k2.InitGenesis(ctx2, genesis)
	require.NoError(t, err)

	// Verify state was imported
	importedCircuit, found := k2.GetCircuit(ctx2, "test-circuit")
	require.True(t, found)
	require.Equal(t, "test-circuit", importedCircuit.Name)

	importedRule, found := k2.GetSymbolicRule(ctx2, "test-circuit", "test-rule")
	require.True(t, found)
	require.Equal(t, "test-rule", importedRule.Name)
}

func TestGetAuthority(t *testing.T) {
	k, _ := setupKeeperTest(t)
	require.Equal(t, "stateset1authority", k.GetAuthority())
}

func TestGetBlockTime(t *testing.T) {
	k, ctx := setupKeeperTest(t)
	blockTime := k.GetBlockTime(ctx)
	require.False(t, blockTime.IsZero())
}

// Benchmark tests
func BenchmarkStoreProof(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	k := keeper.NewKeeper(cdc, storeKey, "authority")
	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, Time: time.Now()}, false, log.NewNopLogger())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proof := &types.Proof{
			CircuitName:  "test-circuit",
			ProofData:    []byte("proof-data-bytes-for-benchmark"),
			PublicInputs: []byte(`{"fields": {}}`),
		}
		k.StoreProof(ctx, proof)
	}
}

func BenchmarkGetProof(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	k := keeper.NewKeeper(cdc, storeKey, "authority")
	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, Time: time.Now()}, false, log.NewNopLogger())

	// Pre-populate with proofs
	for i := 0; i < 1000; i++ {
		proof := &types.Proof{
			CircuitName:  "test-circuit",
			ProofData:    []byte("proof-data"),
			PublicInputs: []byte(`{}`),
		}
		k.StoreProof(ctx, proof)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetProof(ctx, uint64((i%1000)+1))
	}
}

func BenchmarkRegisterCircuit(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	k := keeper.NewKeeper(cdc, storeKey, "authority")
	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, Time: time.Now()}, false, log.NewNopLogger())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		circuit := types.Circuit{
			Name:            "circuit-" + sdkmath.NewInt(int64(i)).String(),
			VerificationKey: []byte("vk-data-at-least-32-bytes-long!!"),
			ProofSystem:     types.ProofSystemSTARK,
		}
		k.RegisterCircuit(ctx, circuit)
	}
}
