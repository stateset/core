package keeper

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/treasury/types"
)

var (
	nextIDKey         = []byte{0x10}
	nextProposalIDKey = []byte{0x11}
	nextRevenueIDKey  = []byte{0x12}
	paramsKey         = []byte{0x13}
)

// BankKeeper defines the expected bank keeper interface
type BankKeeper interface {
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
}

// AccountKeeper defines the expected account keeper interface
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// Keeper manages treasury funds, spend proposals, and budgets.
type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      storetypes.StoreKey
	authority     string
	bankKeeper    BankKeeper
	accountKeeper AccountKeeper
}

// NewKeeper creates a new treasury keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	authority string,
	bankKeeper BankKeeper,
	accountKeeper AccountKeeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		authority:     authority,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

func (k Keeper) GetAuthority() string { return k.authority }

func (k *Keeper) SetAuthority(authority string) { k.authority = authority }

// ============================================================================
// Parameters
// ============================================================================

// GetParams returns the treasury module parameters
func (k Keeper) GetParams(ctx context.Context) types.TreasuryParams {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := store.Get(paramsKey)
	if len(bz) == 0 {
		return types.DefaultTreasuryParams()
	}
	var params types.TreasuryParams
	if err := json.Unmarshal(bz, &params); err != nil {
		return types.DefaultTreasuryParams()
	}
	return params
}

// SetParams sets the treasury module parameters
func (k Keeper) SetParams(ctx context.Context, params types.TreasuryParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz, err := json.Marshal(params)
	if err != nil {
		return err
	}
	store.Set(paramsKey, bz)
	return nil
}

// ============================================================================
// Treasury Balance
// ============================================================================

// GetTreasuryBalance returns the current treasury module balance
func (k Keeper) GetTreasuryBalance(ctx context.Context) sdk.Coins {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleAccountName)
	if moduleAddr == nil {
		return sdk.NewCoins()
	}
	return k.bankKeeper.GetAllBalances(ctx, moduleAddr)
}

// ============================================================================
// ID Management
// ============================================================================

func (k Keeper) getNextID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(nextIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(nextIDKey, bz)
}

func (k Keeper) getNextProposalID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(nextProposalIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextProposalID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(nextProposalIDKey, bz)
}

func (k Keeper) getNextRevenueID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(nextRevenueIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextRevenueID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(nextRevenueIDKey, bz)
}

// ============================================================================
// Spend Proposals (with Timelock)
// ============================================================================

// CreateSpendProposal creates a new time-locked spend proposal
func (k Keeper) CreateSpendProposal(ctx context.Context, proposer, recipient string, amount sdk.Coins, category, description string, timelockDuration time.Duration) (uint64, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(ctx)

	// Validate timelock duration
	if timelockDuration < params.MinTimelockDuration {
		return 0, errorsmod.Wrapf(types.ErrTimelockTooShort,
			"timelock %v is less than minimum %v", timelockDuration, params.MinTimelockDuration)
	}
	if timelockDuration > params.MaxTimelockDuration {
		return 0, errorsmod.Wrapf(types.ErrTimelockTooShort,
			"timelock %v exceeds maximum %v", timelockDuration, params.MaxTimelockDuration)
	}

	// Check pending proposal count
	pendingCount := k.countPendingProposals(ctx)
	if pendingCount >= params.MaxPendingProposals {
		return 0, types.ErrMaxProposalsReached
	}

	// Check budget allows this spend
	budget, found := k.GetBudget(ctx, category)
	if found {
		if err := budget.CanSpend(amount, sdkCtx.BlockTime()); err != nil {
			return 0, err
		}
	}

	// Check treasury has sufficient funds
	balance := k.GetTreasuryBalance(ctx)
	if !balance.IsAllGTE(amount) {
		return 0, errorsmod.Wrapf(types.ErrInsufficientFunds,
			"treasury balance %s is less than requested %s", balance.String(), amount.String())
	}

	now := sdkCtx.BlockTime()
	proposalID := k.getNextProposalID(sdkCtx)

	proposal := types.SpendProposal{
		Id:           proposalID,
		Proposer:     proposer,
		Recipient:    recipient,
		Amount:       amount,
		Category:     category,
		Description:  description,
		Status:       types.SpendStatusPending,
		CreatedAt:    now,
		ExecuteAfter: now.Add(timelockDuration),
		ExpiresAt:    now.Add(timelockDuration).Add(params.ProposalExpiryDuration),
	}

	if err := k.setSpendProposal(ctx, proposal); err != nil {
		return 0, err
	}

	k.setNextProposalID(sdkCtx, proposalID+1)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSpendProposed,
			sdk.NewAttribute(types.AttributeKeyProposalID, mustUint64ToString(proposalID)),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyCategory, category),
			sdk.NewAttribute(types.AttributeKeyExecuteAfter, proposal.ExecuteAfter.Format(time.RFC3339)),
		),
	)

	return proposalID, nil
}

// ExecuteSpendProposal executes a spend proposal after timelock expires
func (k Keeper) ExecuteSpendProposal(ctx context.Context, proposalID uint64) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	proposal, found := k.GetSpendProposal(ctx, proposalID)
	if !found {
		return types.ErrProposalNotFound
	}

	// Check if proposal can be executed
	if err := proposal.CanExecute(sdkCtx.BlockTime()); err != nil {
		return err
	}

	// Check treasury still has sufficient funds
	balance := k.GetTreasuryBalance(ctx)
	if !balance.IsAllGTE(proposal.Amount) {
		return errorsmod.Wrapf(types.ErrInsufficientFunds,
			"treasury balance %s is less than proposal amount %s", balance.String(), proposal.Amount.String())
	}

	// Execute the transfer
	recipientAddr, err := sdk.AccAddressFromBech32(proposal.Recipient)
	if err != nil {
		return errorsmod.Wrap(types.ErrInvalidRecipient, err.Error())
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccountName, recipientAddr, proposal.Amount); err != nil {
		return errorsmod.Wrap(err, "failed to send funds from treasury")
	}

	// Update proposal status
	proposal.Status = types.SpendStatusExecuted
	proposal.ExecutedAt = sdkCtx.BlockTime()
	if err := k.setSpendProposal(ctx, proposal); err != nil {
		return err
	}

	// Update budget spent
	if err := k.recordBudgetSpend(ctx, proposal.Category, proposal.Amount); err != nil {
		sdkCtx.Logger().Error("failed to record budget spend", "error", err)
	}

	// Update allocation tracking
	k.updateAllocation(ctx, proposal.Recipient, proposal.Amount)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSpendExecuted,
			sdk.NewAttribute(types.AttributeKeyProposalID, mustUint64ToString(proposalID)),
			sdk.NewAttribute(types.AttributeKeyRecipient, proposal.Recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, proposal.Amount.String()),
		),
	)

	return nil
}

// CancelSpendProposal cancels a pending spend proposal
func (k Keeper) CancelSpendProposal(ctx context.Context, proposalID uint64, reason string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	proposal, found := k.GetSpendProposal(ctx, proposalID)
	if !found {
		return types.ErrProposalNotFound
	}

	if proposal.Status != types.SpendStatusPending {
		return types.ErrProposalNotPending
	}

	proposal.Status = types.SpendStatusCancelled
	if err := k.setSpendProposal(ctx, proposal); err != nil {
		return err
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSpendCancelled,
			sdk.NewAttribute(types.AttributeKeyProposalID, mustUint64ToString(proposalID)),
			sdk.NewAttribute("reason", reason),
		),
	)

	return nil
}

// GetSpendProposal retrieves a spend proposal by ID
func (k Keeper) GetSpendProposal(ctx context.Context, id uint64) (types.SpendProposal, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SpendProposalKeyPrefix)
	bz := store.Get(mustWriteUint64(id))
	if len(bz) == 0 {
		return types.SpendProposal{}, false
	}
	var proposal types.SpendProposal
	if err := json.Unmarshal(bz, &proposal); err != nil {
		return types.SpendProposal{}, false
	}
	return proposal, true
}

func (k Keeper) setSpendProposal(ctx context.Context, proposal types.SpendProposal) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SpendProposalKeyPrefix)
	bz, err := json.Marshal(proposal)
	if err != nil {
		return err
	}
	store.Set(mustWriteUint64(proposal.Id), bz)
	return nil
}

func (k Keeper) countPendingProposals(ctx context.Context) uint32 {
	var count uint32
	k.IterateSpendProposals(ctx, func(p types.SpendProposal) bool {
		if p.Status == types.SpendStatusPending {
			count++
		}
		return false
	})
	return count
}

// IterateSpendProposals iterates over all spend proposals
func (k Keeper) IterateSpendProposals(ctx context.Context, cb func(types.SpendProposal) bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SpendProposalKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var proposal types.SpendProposal
		if err := json.Unmarshal(iterator.Value(), &proposal); err != nil {
			continue
		}
		if cb(proposal) {
			break
		}
	}
}

// ============================================================================
// Budgets
// ============================================================================

// GetBudget retrieves a budget by category
func (k Keeper) GetBudget(ctx context.Context, category string) (types.Budget, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.BudgetKeyPrefix)
	bz := store.Get([]byte(category))
	if len(bz) == 0 {
		return types.Budget{}, false
	}
	var budget types.Budget
	if err := json.Unmarshal(bz, &budget); err != nil {
		return types.Budget{}, false
	}
	return budget, true
}

// SetBudget sets a budget for a category
func (k Keeper) SetBudget(ctx context.Context, budget types.Budget) error {
	if err := budget.ValidateBasic(); err != nil {
		return err
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.BudgetKeyPrefix)
	bz, err := json.Marshal(budget)
	if err != nil {
		return err
	}
	store.Set([]byte(budget.Category), bz)

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBudgetSet,
			sdk.NewAttribute(types.AttributeKeyCategory, budget.Category),
		),
	)

	return nil
}

func (k Keeper) recordBudgetSpend(ctx context.Context, category string, amount sdk.Coins) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	budget, found := k.GetBudget(ctx, category)
	if !found {
		// No budget tracking for this category
		return nil
	}

	// Reset period if needed
	if sdkCtx.BlockTime().After(budget.PeriodStart.Add(budget.PeriodDuration)) {
		budget.PeriodStart = sdkCtx.BlockTime()
		budget.PeriodSpent = sdk.NewCoins()
	}

	budget.PeriodSpent = budget.PeriodSpent.Add(amount...)
	budget.TotalSpent = budget.TotalSpent.Add(amount...)

	return k.SetBudget(ctx, budget)
}

// IterateBudgets iterates over all budgets
func (k Keeper) IterateBudgets(ctx context.Context, cb func(types.Budget) bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.BudgetKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var budget types.Budget
		if err := json.Unmarshal(iterator.Value(), &budget); err != nil {
			continue
		}
		if cb(budget) {
			break
		}
	}
}

// ============================================================================
// Allocations
// ============================================================================

func (k Keeper) updateAllocation(ctx context.Context, recipient string, amount sdk.Coins) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.AllocationKeyPrefix)

	var allocation types.Allocation
	bz := store.Get([]byte(recipient))
	if len(bz) > 0 {
		json.Unmarshal(bz, &allocation)
	} else {
		allocation = types.Allocation{Recipient: recipient}
	}

	allocation.TotalAllocated = allocation.TotalAllocated.Add(amount...)
	allocation.TotalDisbursed = allocation.TotalDisbursed.Add(amount...)
	allocation.LastUpdated = sdkCtx.BlockTime()

	if newBz, err := json.Marshal(allocation); err == nil {
		store.Set([]byte(recipient), newBz)
	}
}

// GetAllocation retrieves allocation info for a recipient
func (k Keeper) GetAllocation(ctx context.Context, recipient string) (types.Allocation, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.AllocationKeyPrefix)
	bz := store.Get([]byte(recipient))
	if len(bz) == 0 {
		return types.Allocation{}, false
	}
	var allocation types.Allocation
	if err := json.Unmarshal(bz, &allocation); err != nil {
		return types.Allocation{}, false
	}
	return allocation, true
}

// ============================================================================
// Revenue Tracking
// ============================================================================

// RecordRevenue records incoming revenue to the treasury
func (k Keeper) RecordRevenue(ctx context.Context, source string, amount sdk.Coins, metadata string) uint64 {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	revenueID := k.getNextRevenueID(sdkCtx)

	record := types.RevenueRecord{
		Id:        revenueID,
		Source:    source,
		Amount:    amount,
		Timestamp: sdkCtx.BlockTime(),
		Metadata:  metadata,
	}

	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.RevenueRecordKeyPrefix)
	if bz, err := json.Marshal(record); err == nil {
		store.Set(mustWriteUint64(revenueID), bz)
	}

	k.setNextRevenueID(sdkCtx, revenueID+1)

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRevenueReceived,
			sdk.NewAttribute(types.AttributeKeySource, source),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	)

	return revenueID
}

// ============================================================================
// Reserve Snapshots (existing functionality)
// ============================================================================

// RecordSnapshot persists a new reserve snapshot and returns its ID.
func (k Keeper) RecordSnapshot(ctx context.Context, snapshot types.ReserveSnapshot) uint64 {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	nextID := k.getNextID(sdkCtx)
	snapshot.Id = nextID
	if snapshot.Timestamp.IsZero() {
		snapshot.Timestamp = time.Unix(sdkCtx.BlockTime().Unix(), 0)
	}

	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&snapshot)
	store.Set(mustWriteUint64(nextID), bz)

	k.setNextID(sdkCtx, nextID+1)
	return nextID
}

// GetSnapshot fetches a snapshot by id.
func (k Keeper) GetSnapshot(ctx context.Context, id uint64) (types.ReserveSnapshot, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	bz := store.Get(mustWriteUint64(id))
	if len(bz) == 0 {
		return types.ReserveSnapshot{}, false
	}
	var snapshot types.ReserveSnapshot
	types.ModuleCdc.MustUnmarshalJSON(bz, &snapshot)
	return snapshot, true
}

// GetLatestSnapshot returns the newest snapshot if any.
func (k Keeper) GetLatestSnapshot(ctx context.Context) (types.ReserveSnapshot, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	iterator := store.ReverseIterator(nil, nil)
	defer iterator.Close()
	if iterator.Valid() {
		var snapshot types.ReserveSnapshot
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &snapshot)
		return snapshot, true
	}
	return types.ReserveSnapshot{}, false
}

// IterateSnapshots iterates over stored snapshots.
func (k Keeper) IterateSnapshots(ctx context.Context, cb func(types.ReserveSnapshot) bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var snapshot types.ReserveSnapshot
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &snapshot)
		if cb(snapshot) {
			break
		}
	}
}

// ============================================================================
// Genesis
// ============================================================================

// InitGenesis initializes from genesis state.
func (k Keeper) InitGenesis(ctx context.Context, state *types.GenesisState) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if state == nil {
		state = types.DefaultGenesis()
	}
	k.authority = state.Authority
	k.setNextID(sdkCtx, state.NextID)
	k.setNextProposalID(sdkCtx, state.NextProposalID)
	k.setNextRevenueID(sdkCtx, state.NextRevenueID)

	if err := k.SetParams(ctx, state.Params); err != nil {
		sdkCtx.Logger().Error("failed to set treasury params", "error", err)
	}

	for _, snapshot := range state.Snapshots {
		store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.SnapshotsKeyPrefix)
		bz := types.ModuleCdc.MustMarshalJSON(&snapshot)
		store.Set(mustWriteUint64(snapshot.Id), bz)
	}

	for _, proposal := range state.SpendProposals {
		k.setSpendProposal(ctx, proposal)
	}

	for _, budget := range state.Budgets {
		k.SetBudget(ctx, budget)
	}
}

// ExportGenesis exports the module state.
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Authority = k.authority
	state.Params = k.GetParams(ctx)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	state.NextID = k.getNextID(sdkCtx)
	state.NextProposalID = k.getNextProposalID(sdkCtx)
	state.NextRevenueID = k.getNextRevenueID(sdkCtx)

	k.IterateSnapshots(ctx, func(snapshot types.ReserveSnapshot) bool {
		state.Snapshots = append(state.Snapshots, snapshot)
		return false
	})

	k.IterateSpendProposals(ctx, func(proposal types.SpendProposal) bool {
		state.SpendProposals = append(state.SpendProposals, proposal)
		return false
	})

	k.IterateBudgets(ctx, func(budget types.Budget) bool {
		state.Budgets = append(state.Budgets, budget)
		return false
	})

	return state
}

// ============================================================================
// Helpers
// ============================================================================

func mustWriteUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func mustUint64ToString(id uint64) string {
	return string(mustWriteUint64(id))
}
