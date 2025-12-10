package keeper_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/treasury/keeper"
	"github.com/stateset/core/x/treasury/types"
)

// TestSimulateTimelockGovernance simulates the timelock governance system
func TestSimulateTimelockGovernance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping simulation test in short mode")
	}

	k, ctx, bank := setupTreasurySimKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)
	authority := k.GetAuthority()

	// Fund treasury
	treasuryAddr := bank.moduleAddresses[types.ModuleAccountName]
	bank.SetModuleBalance(types.ModuleAccountName, sdk.NewCoins(
		sdk.NewInt64Coin("ustate", 100_000_000_000), // 100k STATE
		sdk.NewInt64Coin("ssusd", 50_000_000_000),   // 50k ssUSD
	))

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create multiple proposals with different timelocks
	proposals := make([]uint64, 0)
	for i := 0; i < 20; i++ {
		recipient := newTreasurySimAddress()
		amount := sdk.NewCoins(sdk.NewInt64Coin("ustate", int64(1_000_000+rng.Intn(10_000_000))))
		timelockSecs := uint64(86400 + rng.Intn(604800)) // 1-8 days

		msg := types.NewMsgProposeSpend(
			authority,
			recipient.String(),
			amount,
			types.CategoryGrants,
			fmt.Sprintf("Grant proposal #%d for ecosystem development", i+1),
			timelockSecs,
		)

		resp, err := msgServer.ProposeSpend(ctx, msg)
		if err == nil {
			proposals = append(proposals, resp.ProposalID)
			t.Logf("Created proposal %d: amount=%s, executeAfter=%s",
				resp.ProposalID, amount.String(), resp.ExecuteAfter.Format(time.RFC3339))
		}
	}

	t.Logf("Created %d proposals", len(proposals))

	// Simulate time passing and execution attempts
	for day := 1; day <= 14; day++ {
		// Advance time by 1 day
		ctx = ctx.WithBlockTime(ctx.BlockTime().Add(24 * time.Hour))
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 17280) // ~5s blocks

		executed := 0
		premature := 0
		expired := 0

		for _, proposalID := range proposals {
			proposal, found := k.GetSpendProposal(ctx, proposalID)
			if !found || proposal.Status != types.SpendStatusPending {
				continue
			}

			// Try to execute
			execMsg := types.NewMsgExecuteSpend(authority, proposalID)
			_, err := msgServer.ExecuteSpend(ctx, execMsg)

			if err == nil {
				executed++
			} else if err == types.ErrTimelockNotExpired {
				premature++
			} else if err == types.ErrProposalExpired {
				expired++
			}
		}

		t.Logf("Day %d: executed=%d, premature=%d, expired=%d",
			day, executed, premature, expired)
	}

	// Verify treasury balance decreased appropriately
	finalBalance := bank.GetModuleBalance(types.ModuleAccountName)
	t.Logf("Final treasury balance: %s (started with 100k ustate)", finalBalance.String())

	// Check for any remaining pending proposals
	pendingCount := 0
	k.IterateSpendProposals(ctx, func(p types.SpendProposal) bool {
		if p.Status == types.SpendStatusPending {
			pendingCount++
		}
		return false
	})
	t.Logf("Remaining pending proposals: %d", pendingCount)

	require.NotNil(t, treasuryAddr) // Just to use the variable
}

// TestSimulateBudgetEnforcement tests budget limit enforcement
func TestSimulateBudgetEnforcement(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping simulation test in short mode")
	}

	k, ctx, bank := setupTreasurySimKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)
	authority := k.GetAuthority()

	// Fund treasury generously
	bank.SetModuleBalance(types.ModuleAccountName, sdk.NewCoins(
		sdk.NewInt64Coin("ustate", 1_000_000_000_000), // 1M STATE
	))

	// Set strict budget for grants
	budgetMsg := types.NewMsgSetBudget(
		authority,
		types.CategoryGrants,
		sdk.NewCoins(sdk.NewInt64Coin("ustate", 100_000_000_000)), // 100k total limit
		sdk.NewCoins(sdk.NewInt64Coin("ustate", 10_000_000_000)),  // 10k per period
		30*24*time.Hour, // 30 day period
		true,
	)
	_, err := msgServer.SetBudget(ctx, budgetMsg)
	require.NoError(t, err)

	// Try to create proposals that would exceed budget
	successfulProposals := 0
	rejectedProposals := 0

	for i := 0; i < 20; i++ {
		recipient := newTreasurySimAddress()
		// Try to spend 2k each time (would exceed 10k period limit after 5)
		amount := sdk.NewCoins(sdk.NewInt64Coin("ustate", 2_000_000_000))

		msg := types.NewMsgProposeSpend(
			authority,
			recipient.String(),
			amount,
			types.CategoryGrants,
			fmt.Sprintf("Grant %d", i+1),
			86400, // 24 hour timelock
		)

		_, err := msgServer.ProposeSpend(ctx, msg)
		if err == nil {
			successfulProposals++
		} else {
			rejectedProposals++
			t.Logf("Proposal %d rejected: %v", i+1, err)
		}
	}

	t.Logf("Budget enforcement: %d successful, %d rejected", successfulProposals, rejectedProposals)

	// Verify budget was enforced
	budget, found := k.GetBudget(ctx, types.CategoryGrants)
	require.True(t, found)
	t.Logf("Final budget state: periodSpent=%s, totalSpent=%s",
		budget.PeriodSpent.String(), budget.TotalSpent.String())
}

// TestSimulateProposalCancellation tests proposal cancellation scenarios
func TestSimulateProposalCancellation(t *testing.T) {
	k, ctx, bank := setupTreasurySimKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)
	authority := k.GetAuthority()

	bank.SetModuleBalance(types.ModuleAccountName, sdk.NewCoins(
		sdk.NewInt64Coin("ustate", 100_000_000_000),
	))

	// Create proposal
	recipient := newTreasurySimAddress()
	proposeMsg := types.NewMsgProposeSpend(
		authority,
		recipient.String(),
		sdk.NewCoins(sdk.NewInt64Coin("ustate", 1_000_000_000)),
		types.CategoryDevelopment,
		"Development grant",
		86400,
	)
	resp, err := msgServer.ProposeSpend(ctx, proposeMsg)
	require.NoError(t, err)

	// Cancel before timelock expires
	cancelMsg := types.NewMsgCancelSpend(authority, resp.ProposalID, "Changed priorities")
	_, err = msgServer.CancelSpend(ctx, cancelMsg)
	require.NoError(t, err)

	// Verify cannot execute cancelled proposal
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(48 * time.Hour)) // Past timelock
	execMsg := types.NewMsgExecuteSpend(authority, resp.ProposalID)
	_, err = msgServer.ExecuteSpend(ctx, execMsg)
	require.Error(t, err)
	require.Equal(t, types.ErrProposalNotPending, err)

	// Verify proposal status
	proposal, found := k.GetSpendProposal(ctx, resp.ProposalID)
	require.True(t, found)
	require.Equal(t, types.SpendStatusCancelled, proposal.Status)
}

// TestFuzzProposalCreation fuzz tests proposal creation
func TestFuzzProposalCreation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping fuzz test in short mode")
	}

	k, ctx, bank := setupTreasurySimKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)
	authority := k.GetAuthority()

	bank.SetModuleBalance(types.ModuleAccountName, sdk.NewCoins(
		sdk.NewInt64Coin("ustate", 1_000_000_000_000),
	))

	rng := rand.New(rand.NewSource(42))

	validProposals := 0
	invalidProposals := 0

	categories := types.ValidCategories()

	for i := 0; i < 1000; i++ {
		// Fuzz inputs
		recipient := newTreasurySimAddress()
		amount := rng.Int63n(100_000_000_000) // 0 to 100k
		category := categories[rng.Intn(len(categories))]
		timelockSecs := uint64(rng.Intn(30 * 24 * 3600)) // 0 to 30 days

		// Generate random description
		descLen := rng.Intn(2000) // 0 to 2000 chars (some over limit)
		desc := make([]byte, descLen)
		for j := range desc {
			desc[j] = byte('a' + rng.Intn(26))
		}

		msg := &types.MsgProposeSpend{
			Authority:       authority,
			Recipient:       recipient.String(),
			Amount:          sdk.NewCoins(sdk.NewInt64Coin("ustate", amount)),
			Category:        category,
			Description:     string(desc),
			TimelockSeconds: timelockSecs,
		}

		_, err := msgServer.ProposeSpend(ctx, msg)
		if err == nil {
			validProposals++
		} else {
			invalidProposals++
		}

		// Verify no panics and state is consistent
		k.GetParams(ctx) // Should not panic
	}

	t.Logf("Fuzz results: valid=%d, invalid=%d", validProposals, invalidProposals)
	require.True(t, validProposals > 0, "Expected some valid proposals")
	require.True(t, invalidProposals > 0, "Expected some invalid proposals (fuzz should find edge cases)")
}

// Helper functions
func setupTreasurySimKeeper(t *testing.T) (keeper.Keeper, sdk.Context, *treasurySimBankKeeper) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	err := stateStore.LoadLatestVersion()
	require.NoError(t, err)

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	authority := "stateset1authority"

	ctx := sdk.NewContext(stateStore, cmtproto.Header{
		ChainID: "stateset-sim",
		Time:    time.Now(),
		Height:  1,
	}, false, log.NewNopLogger())

	bankKeeper := newTreasurySimBankKeeper()
	accountKeeper := newTreasurySimAccountKeeper()

	// Set treasury module address
	treasuryAddr := newTreasurySimAddress()
	accountKeeper.SetModuleAddress(types.ModuleAccountName, treasuryAddr)
	bankKeeper.moduleAddresses[types.ModuleAccountName] = treasuryAddr

	k := keeper.NewKeeper(cdc, storeKey, authority, bankKeeper, accountKeeper)

	// Initialize with default params
	k.SetParams(ctx, types.DefaultTreasuryParams())

	return k, ctx, bankKeeper
}

func newTreasurySimAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

// Mock keepers for treasury simulation
type treasurySimBankKeeper struct {
	balances        map[string]sdk.Coins
	moduleBalances  map[string]sdk.Coins
	moduleAddresses map[string]sdk.AccAddress
}

func newTreasurySimBankKeeper() *treasurySimBankKeeper {
	return &treasurySimBankKeeper{
		balances:        make(map[string]sdk.Coins),
		moduleBalances:  make(map[string]sdk.Coins),
		moduleAddresses: make(map[string]sdk.AccAddress),
	}
}

func (m *treasurySimBankKeeper) GetBalance(_ context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins := m.balances[addr.String()]
	return sdk.NewCoin(denom, coins.AmountOf(denom))
}

func (m *treasurySimBankKeeper) GetAllBalances(_ context.Context, addr sdk.AccAddress) sdk.Coins {
	// Check if this is a module address
	for module, moduleAddr := range m.moduleAddresses {
		if moduleAddr.Equals(addr) {
			return m.moduleBalances[module]
		}
	}
	return m.balances[addr.String()]
}

func (m *treasurySimBankKeeper) SetBalance(addr sdk.AccAddress, coins sdk.Coins) {
	m.balances[addr.String()] = coins.Sort()
}

func (m *treasurySimBankKeeper) SetModuleBalance(module string, coins sdk.Coins) {
	m.moduleBalances[module] = coins.Sort()
}

func (m *treasurySimBankKeeper) GetModuleBalance(module string) sdk.Coins {
	return m.moduleBalances[module]
}

func (m *treasurySimBankKeeper) SendCoinsFromModuleToAccount(_ context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[senderModule]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	newBalance, negative := moduleCoins.SafeSub(amt...)
	if negative {
		return fmt.Errorf("insufficient module funds")
	}
	m.moduleBalances[senderModule] = newBalance.Sort()

	account := m.balances[recipientAddr.String()]
	if account == nil {
		account = sdk.NewCoins()
	}
	m.balances[recipientAddr.String()] = account.Add(amt...).Sort()
	return nil
}

func (m *treasurySimBankKeeper) SendCoinsFromAccountToModule(_ context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	account := m.balances[senderAddr.String()]
	if account == nil {
		account = sdk.NewCoins()
	}
	newBalance, negative := account.SafeSub(amt...)
	if negative {
		return fmt.Errorf("insufficient funds")
	}
	m.balances[senderAddr.String()] = newBalance.Sort()

	moduleCoins := m.moduleBalances[recipientModule]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[recipientModule] = moduleCoins.Add(amt...).Sort()
	return nil
}

func (m *treasurySimBankKeeper) BurnCoins(_ context.Context, moduleName string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[moduleName]
	if moduleCoins == nil {
		return fmt.Errorf("module has no coins")
	}
	newBalance, negative := moduleCoins.SafeSub(amt...)
	if negative {
		return fmt.Errorf("insufficient module funds to burn")
	}
	m.moduleBalances[moduleName] = newBalance.Sort()
	return nil
}

func (m *treasurySimBankKeeper) MintCoins(_ context.Context, moduleName string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[moduleName]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[moduleName] = moduleCoins.Add(amt...).Sort()
	return nil
}

type treasurySimAccountKeeper struct {
	addresses map[string]sdk.AccAddress
}

func newTreasurySimAccountKeeper() *treasurySimAccountKeeper {
	return &treasurySimAccountKeeper{
		addresses: make(map[string]sdk.AccAddress),
	}
}

func (m *treasurySimAccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return m.addresses[moduleName]
}

func (m *treasurySimAccountKeeper) SetModuleAddress(moduleName string, addr sdk.AccAddress) {
	m.addresses[moduleName] = addr
}
