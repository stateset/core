package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/stateset/core/x/stst/types"
)

// Keeper of the STST module
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	logger     log.Logger

	// Keepers
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
	
	// Authority is the address capable of executing governance proposals
	authority string
	
	// Module account permissions
	maccPerms map[string][]string
}

// NewKeeper creates a new STST Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey, memKey storetypes.StoreKey,
	logger log.Logger,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	authority string,
) *Keeper {
	// Ensure the authority address is valid
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		logger:        logger,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
		authority:     authority,
		maccPerms: map[string][]string{
			types.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		},
	}
}

// GetAuthority returns the module's authority
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", "x/"+types.ModuleName)
}

// SetParams sets the parameters for the module
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)
	return nil
}

// GetParams gets the parameters for the module
func (k Keeper) GetParams(ctx context.Context) (types.Params, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return types.DefaultParams(), nil
	}

	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params, nil
}

// SetStakingState sets the staking state
func (k Keeper) SetStakingState(ctx context.Context, state types.StakingState) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&state)
	store.Set(types.StakingStateKey, bz)
	return nil
}

// GetStakingState gets the staking state
func (k Keeper) GetStakingState(ctx context.Context) (types.StakingState, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := store.Get(types.StakingStateKey)
	if bz == nil {
		return types.DefaultStakingState(), nil
	}

	var state types.StakingState
	k.cdc.MustUnmarshal(bz, &state)
	return state, nil
}

// SetFeeBurnState sets the fee burn state
func (k Keeper) SetFeeBurnState(ctx context.Context, state types.FeeBurnState) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&state)
	store.Set(types.FeeBurnStateKey, bz)
	return nil
}

// GetFeeBurnState gets the fee burn state
func (k Keeper) GetFeeBurnState(ctx context.Context) (types.FeeBurnState, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := store.Get(types.FeeBurnStateKey)
	if bz == nil {
		return types.DefaultFeeBurnState(), nil
	}

	var state types.FeeBurnState
	k.cdc.MustUnmarshal(bz, &state)
	return state, nil
}

// SetVestingSchedule sets a vesting schedule
func (k Keeper) SetVestingSchedule(ctx context.Context, schedule types.VestingSchedule) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&schedule)
	key := types.VestingScheduleKey(schedule.Category)
	store.Set(key, bz)
	return nil
}

// GetVestingSchedule gets a vesting schedule by category
func (k Keeper) GetVestingSchedule(ctx context.Context, category string) (types.VestingSchedule, bool, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	key := types.VestingScheduleKey(category)
	bz := store.Get(key)
	if bz == nil {
		return types.VestingSchedule{}, false, nil
	}

	var schedule types.VestingSchedule
	k.cdc.MustUnmarshal(bz, &schedule)
	return schedule, true, nil
}

// GetAllVestingSchedules gets all vesting schedules
func (k Keeper) GetAllVestingSchedules(ctx context.Context) ([]types.VestingSchedule, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetVestingScheduleIteratorKey())
	defer iterator.Close()

	var schedules []types.VestingSchedule
	for ; iterator.Valid(); iterator.Next() {
		var schedule types.VestingSchedule
		k.cdc.MustUnmarshal(iterator.Value(), &schedule)
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

// StakeTokens handles staking STST tokens
func (k Keeper) StakeTokens(ctx context.Context, stakerAddr, validatorAddr string, amount sdk.Coin) error {
	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}

	// Validate amount
	if amount.Denom != params.TokenDenom {
		return fmt.Errorf("invalid token denomination: expected %s, got %s", params.TokenDenom, amount.Denom)
	}

	if amount.Amount.LT(params.MinStakingAmount) {
		return fmt.Errorf("staking amount %s is below minimum %s", amount.Amount, params.MinStakingAmount)
	}

	// Transfer tokens from staker to module account
	stakerAccAddr, err := sdk.AccAddressFromBech32(stakerAddr)
	if err != nil {
		return fmt.Errorf("invalid staker address: %w", err)
	}

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if err := k.bankKeeper.SendCoins(ctx, stakerAccAddr, moduleAddr, sdk.NewCoins(amount)); err != nil {
		return fmt.Errorf("failed to transfer tokens: %w", err)
	}

	// Update staking state
	stakingState, err := k.GetStakingState(ctx)
	if err != nil {
		return err
	}

	stakingState.TotalStaked = stakingState.TotalStaked.Add(amount.Amount)
	if err := k.SetStakingState(ctx, stakingState); err != nil {
		return err
	}

	// Store delegation information
	delegation := types.Delegation{
		ValidatorAddress: validatorAddr,
		Amount:           amount,
		Rewards:          sdk.NewCoins(),
	}

	return k.setDelegation(ctx, stakerAddr, delegation)
}

// UnstakeTokens handles unstaking STST tokens
func (k Keeper) UnstakeTokens(ctx context.Context, stakerAddr, validatorAddr string, amount sdk.Coin) (int64, error) {
	params, err := k.GetParams(ctx)
	if err != nil {
		return 0, err
	}

	// Get current delegation
	delegation, found, err := k.getDelegation(ctx, stakerAddr, validatorAddr)
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, fmt.Errorf("delegation not found")
	}

	if delegation.Amount.Amount.LT(amount.Amount) {
		return 0, fmt.Errorf("insufficient staked amount: have %s, requested %s", 
			delegation.Amount.Amount, amount.Amount)
	}

	// Calculate completion time (21 days unbonding period)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	completionTime := sdkCtx.BlockTime().Add(time.Hour * 24 * 21).Unix()

	// Update delegation
	delegation.Amount = delegation.Amount.Sub(amount)
	if delegation.Amount.IsZero() {
		if err := k.deleteDelegation(ctx, stakerAddr, validatorAddr); err != nil {
			return 0, err
		}
	} else {
		if err := k.setDelegation(ctx, stakerAddr, delegation); err != nil {
			return 0, err
		}
	}

	// Update staking state
	stakingState, err := k.GetStakingState(ctx)
	if err != nil {
		return 0, err
	}

	stakingState.TotalStaked = stakingState.TotalStaked.Sub(amount.Amount)
	if err := k.SetStakingState(ctx, stakingState); err != nil {
		return 0, err
	}

	// Add to unstaking queue
	if err := k.addToUnstakingQueue(ctx, completionTime, stakerAddr, amount); err != nil {
		return 0, err
	}

	return completionTime, nil
}

// BurnTokens burns tokens as part of the deflationary mechanism
func (k Keeper) BurnTokens(ctx context.Context, amount sdk.Coin) error {
	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}

	if amount.Denom != params.TokenDenom {
		return fmt.Errorf("invalid token denomination: expected %s, got %s", params.TokenDenom, amount.Denom)
	}

	// Burn tokens from module account
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return fmt.Errorf("failed to burn tokens: %w", err)
	}

	// Update fee burn state
	feeBurnState, err := k.GetFeeBurnState(ctx)
	if err != nil {
		return err
	}

	feeBurnState.TotalBurned = feeBurnState.TotalBurned.Add(amount.Amount)

	// Add to burn rate history
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	burnRateEntry := types.BurnRateEntry{
		BlockHeight:   sdkCtx.BlockHeight(),
		BurnRate:      params.FeeBurnRate,
		AmountBurned:  amount.Amount,
	}
	feeBurnState.BurnRateHistory = append(feeBurnState.BurnRateHistory, burnRateEntry)

	return k.SetFeeBurnState(ctx, feeBurnState)
}

// CalculateVestedAmount calculates how much is vested for a given schedule
func (k Keeper) CalculateVestedAmount(ctx context.Context, schedule types.VestingSchedule) sdk.Int {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentTime := sdkCtx.BlockTime().Unix()

	// If before cliff, nothing is vested
	cliffTime := schedule.StartTime + int64(schedule.CliffDuration.Seconds())
	if currentTime < cliffTime {
		return sdk.ZeroInt()
	}

	// If after full vesting period, everything is vested
	fullVestingTime := schedule.StartTime + int64(schedule.VestingDuration.Seconds())
	if currentTime >= fullVestingTime {
		return schedule.TotalAmount
	}

	// Calculate linear vesting
	timeElapsed := currentTime - cliffTime
	vestingPeriod := int64(schedule.VestingDuration.Seconds()) - int64(schedule.CliffDuration.Seconds())
	
	if vestingPeriod <= 0 {
		return schedule.TotalAmount
	}

	vestingRatio := sdk.NewDec(timeElapsed).Quo(sdk.NewDec(vestingPeriod))
	vestedAmount := vestingRatio.MulInt(schedule.TotalAmount).TruncateInt()

	return vestedAmount
}

// Helper functions

func (k Keeper) setDelegation(ctx context.Context, stakerAddr string, delegation types.Delegation) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	key := types.DelegationKey(stakerAddr, delegation.ValidatorAddress)
	bz := k.cdc.MustMarshal(&delegation)
	store.Set(key, bz)
	return nil
}

func (k Keeper) getDelegation(ctx context.Context, stakerAddr, validatorAddr string) (types.Delegation, bool, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	key := types.DelegationKey(stakerAddr, validatorAddr)
	bz := store.Get(key)
	if bz == nil {
		return types.Delegation{}, false, nil
	}

	var delegation types.Delegation
	k.cdc.MustUnmarshal(bz, &delegation)
	return delegation, true, nil
}

func (k Keeper) deleteDelegation(ctx context.Context, stakerAddr, validatorAddr string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	key := types.DelegationKey(stakerAddr, validatorAddr)
	store.Delete(key)
	return nil
}

func (k Keeper) addToUnstakingQueue(ctx context.Context, completionTime int64, stakerAddr string, amount sdk.Coin) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	key := types.UnstakingQueueKey(completionTime, stakerAddr)
	bz := k.cdc.MustMarshal(&amount)
	store.Set(key, bz)
	return nil
}

// ProcessUnstakingQueue processes the unstaking queue and returns tokens to users
func (k Keeper) ProcessUnstakingQueue(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentTime := sdkCtx.BlockTime().Unix()

	store := sdkCtx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnstakingQueueKeyPrefix)
	defer iterator.Close()

	var keysToDelete [][]byte

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		
		// Extract completion time from key (first 8 bytes after prefix)
		timeBytes := key[len(types.UnstakingQueueKeyPrefix):len(types.UnstakingQueueKeyPrefix)+8]
		completionTime := int64(sdk.BigEndianToUint64(timeBytes))

		if completionTime <= currentTime {
			// Time to complete unstaking
			var amount sdk.Coin
			k.cdc.MustUnmarshal(iterator.Value(), &amount)

			// Extract staker address from key
			stakerAddrBytes := key[len(types.UnstakingQueueKeyPrefix)+8:]
			stakerAddr := string(stakerAddrBytes)

			// Send tokens back to staker
			stakerAccAddr, err := sdk.AccAddressFromBech32(stakerAddr)
			if err != nil {
				k.Logger().Error("invalid staker address in unstaking queue", "address", stakerAddr, "error", err)
				continue
			}

			moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
			if err := k.bankKeeper.SendCoins(ctx, moduleAddr, stakerAccAddr, sdk.NewCoins(amount)); err != nil {
				k.Logger().Error("failed to send unstaked tokens", "staker", stakerAddr, "amount", amount, "error", err)
				continue
			}

			keysToDelete = append(keysToDelete, key)
		}
	}

	// Delete processed entries
	for _, key := range keysToDelete {
		store.Delete(key)
	}

	return nil
}