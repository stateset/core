package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/circuit/types"
)

// Keeper manages circuit breaker state
type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	authority  string
}

// NewKeeper creates a new circuit keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authority string,
) Keeper {
	return Keeper{
		storeKey:  storeKey,
		cdc:       cdc,
		authority: authority,
	}
}

// GetAuthority returns the module authority
func (k Keeper) GetAuthority() string {
	return k.authority
}

// ============================================================================
// Params
// ============================================================================

// GetParams returns the module parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("params"))
	if len(bz) == 0 {
		return types.DefaultParams()
	}
	var params types.Params
	if err := json.Unmarshal(bz, &params); err != nil {
		ctx.Logger().Error("failed to unmarshal circuit params", "error", err)
		return types.DefaultParams()
	}
	return params
}

// SetParams sets the module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	bz, err := json.Marshal(params)
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal circuit params")
	}
	store.Set([]byte("params"), bz)
	return nil
}

// ============================================================================
// Global Circuit State
// ============================================================================

// GetCircuitState returns the global circuit state
func (k Keeper) GetCircuitState(ctx sdk.Context) types.CircuitState {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.CircuitStateKey)
	if len(bz) == 0 {
		return types.CircuitState{GlobalPaused: false}
	}
	state, err := types.UnmarshalCircuitState(bz)
	if err != nil {
		ctx.Logger().Error("failed to unmarshal circuit state", "error", err)
		return types.CircuitState{GlobalPaused: false}
	}
	return state
}

// SetCircuitState sets the global circuit state
func (k Keeper) SetCircuitState(ctx sdk.Context, state types.CircuitState) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := state.Marshal()
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal circuit state")
	}
	store.Set(types.CircuitStateKey, bz)
	return nil
}

// IsGloballyPaused returns true if the system is globally paused
func (k Keeper) IsGloballyPaused(ctx sdk.Context) bool {
	state := k.GetCircuitState(ctx)
	if !state.GlobalPaused {
		return false
	}

	// Check for auto-resume
	if !state.AutoResumeAt.IsZero() && ctx.BlockTime().After(state.AutoResumeAt) {
		// Auto-resume
		state.GlobalPaused = false
		if err := k.SetCircuitState(ctx, state); err != nil {
			ctx.Logger().Error("failed to auto-resume circuit state", "error", err)
		}
		return false
	}

	return true
}

// PauseSystem pauses the entire system
func (k Keeper) PauseSystem(ctx sdk.Context, authority, reason string, durationSeconds int64) error {
	state := k.GetCircuitState(ctx)
	if state.GlobalPaused {
		return types.ErrAlreadyPaused
	}

	params := k.GetParams(ctx)
	if durationSeconds > params.MaxPauseDuration {
		return types.ErrInvalidDuration
	}

	state.GlobalPaused = true
	state.PausedAt = ctx.BlockTime()
	state.PausedBy = authority
	state.Reason = reason

	if durationSeconds > 0 {
		state.AutoResumeAt = ctx.BlockTime().Add(time.Duration(durationSeconds) * time.Second)
	} else {
		state.AutoResumeAt = time.Time{}
	}

	if err := k.SetCircuitState(ctx, state); err != nil {
		return err
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"system_paused",
			sdk.NewAttribute("authority", authority),
			sdk.NewAttribute("reason", reason),
		),
	)

	return nil
}

// ResumeSystem resumes the system from pause
func (k Keeper) ResumeSystem(ctx sdk.Context, authority string) error {
	state := k.GetCircuitState(ctx)
	if !state.GlobalPaused {
		return types.ErrNotPaused
	}

	state.GlobalPaused = false
	state.AutoResumeAt = time.Time{}
	if err := k.SetCircuitState(ctx, state); err != nil {
		return err
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"system_resumed",
			sdk.NewAttribute("authority", authority),
		),
	)

	return nil
}

// ============================================================================
// Module Circuit State
// ============================================================================

// GetModuleCircuitState returns the circuit state for a specific module
func (k Keeper) GetModuleCircuitState(ctx sdk.Context, moduleName string) (types.ModuleCircuitState, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ModuleCircuitKeyPrefix)
	bz := store.Get([]byte(moduleName))
	if len(bz) == 0 {
		return types.ModuleCircuitState{}, false
	}
	state, err := types.UnmarshalModuleCircuitState(bz)
	if err != nil {
		ctx.Logger().Error("failed to unmarshal module circuit state", "module", moduleName, "error", err)
		return types.ModuleCircuitState{}, false
	}
	return state, true
}

// SetModuleCircuitState sets the circuit state for a specific module
func (k Keeper) SetModuleCircuitState(ctx sdk.Context, state types.ModuleCircuitState) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ModuleCircuitKeyPrefix)
	bz, err := state.Marshal()
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal module circuit state")
	}
	store.Set([]byte(state.ModuleName), bz)
	return nil
}

// IsModuleCircuitOpen checks if a module's circuit is open
func (k Keeper) IsModuleCircuitOpen(ctx sdk.Context, moduleName string) bool {
	state, found := k.GetModuleCircuitState(ctx, moduleName)
	if !found {
		return false
	}

	// Check if circuit is in recovery mode
	if state.Status == types.CircuitHalfOpen {
		if ctx.BlockTime().After(state.RecoveryTime) {
			// Recovery period over, reset to closed
			state.Status = types.CircuitClosed
			state.FailureCount = 0
			if err := k.SetModuleCircuitState(ctx, state); err != nil {
				ctx.Logger().Error("failed to reset module circuit state", "module", moduleName, "error", err)
			}
			return false
		}
	}

	return state.Status == types.CircuitOpen
}

// IsMessageDisabled checks if a specific message type is disabled
func (k Keeper) IsMessageDisabled(ctx sdk.Context, moduleName, msgType string) bool {
	state, found := k.GetModuleCircuitState(ctx, moduleName)
	if !found {
		return false
	}

	if state.Status != types.CircuitOpen {
		return false
	}

	// If no specific messages disabled, all are disabled
	if len(state.DisabledMessages) == 0 {
		return true
	}

	// Check if this specific message is disabled
	for _, disabled := range state.DisabledMessages {
		if disabled == msgType {
			return true
		}
	}

	return false
}

// TripCircuit trips a module's circuit breaker
func (k Keeper) TripCircuit(ctx sdk.Context, moduleName, reason, authority string, disableMessages []string) error {
	params := k.GetParams(ctx)

	state := types.ModuleCircuitState{
		ModuleName:       moduleName,
		Status:           types.CircuitOpen,
		TrippedAt:        ctx.BlockTime(),
		TrippedBy:        authority,
		Reason:           reason,
		FailureCount:     0,
		FailureThreshold: params.DefaultFailureThreshold,
		RecoveryTime:     ctx.BlockTime().Add(time.Duration(params.DefaultRecoveryPeriod) * time.Second),
		DisabledMessages: disableMessages,
	}

	if err := k.SetModuleCircuitState(ctx, state); err != nil {
		return err
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"circuit_tripped",
			sdk.NewAttribute("module", moduleName),
			sdk.NewAttribute("reason", reason),
			sdk.NewAttribute("authority", authority),
		),
	)

	return nil
}

// ResetCircuit resets a module's circuit breaker
func (k Keeper) ResetCircuit(ctx sdk.Context, moduleName, authority string) error {
	state, found := k.GetModuleCircuitState(ctx, moduleName)
	if !found {
		return types.ErrModuleNotFound
	}

	state.Status = types.CircuitClosed
	state.FailureCount = 0
	state.DisabledMessages = nil

	if err := k.SetModuleCircuitState(ctx, state); err != nil {
		return err
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"circuit_reset",
			sdk.NewAttribute("module", moduleName),
			sdk.NewAttribute("authority", authority),
		),
	)

	return nil
}

// RecordFailure records a failure for a module, potentially tripping the circuit
func (k Keeper) RecordFailure(ctx sdk.Context, moduleName string) {
	params := k.GetParams(ctx)

	state, found := k.GetModuleCircuitState(ctx, moduleName)
	if !found {
		state = types.ModuleCircuitState{
			ModuleName:       moduleName,
			Status:           types.CircuitClosed,
			FailureThreshold: params.DefaultFailureThreshold,
		}
	}

	state.FailureCount++

	if state.FailureCount >= state.FailureThreshold {
		state.Status = types.CircuitOpen
		state.TrippedAt = ctx.BlockTime()
		state.TrippedBy = "automatic"
		state.Reason = "failure threshold exceeded"
		state.RecoveryTime = ctx.BlockTime().Add(time.Duration(params.DefaultRecoveryPeriod) * time.Second)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"circuit_auto_tripped",
				sdk.NewAttribute("module", moduleName),
				sdk.NewAttribute("failure_count", fmt.Sprintf("%d", state.FailureCount)),
			),
		)
	}

	if err := k.SetModuleCircuitState(ctx, state); err != nil {
		ctx.Logger().Error("failed to record circuit failure", "module", moduleName, "error", err)
	}
}

// RecordSuccess records a success, potentially resetting failure count
func (k Keeper) RecordSuccess(ctx sdk.Context, moduleName string) {
	state, found := k.GetModuleCircuitState(ctx, moduleName)
	if !found {
		return
	}

	// Reset failure count on success
	if state.FailureCount > 0 {
		state.FailureCount = 0
		if err := k.SetModuleCircuitState(ctx, state); err != nil {
			ctx.Logger().Error("failed to record circuit success", "module", moduleName, "error", err)
		}
	}
}

// ============================================================================
// Rate Limiting
// ============================================================================

// getRateLimitKey generates the key for rate limit state
func (k Keeper) getRateLimitKey(configName, address string) []byte {
	if address == "" {
		return []byte(configName)
	}
	return []byte(configName + ":" + address)
}

// GetRateLimitState returns the rate limit state
func (k Keeper) GetRateLimitState(ctx sdk.Context, configName, address string) (types.RateLimitState, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RateLimitKeyPrefix)
	key := k.getRateLimitKey(configName, address)
	bz := store.Get(key)
	if len(bz) == 0 {
		return types.RateLimitState{}, false
	}
	state, err := types.UnmarshalRateLimitState(bz)
	if err != nil {
		ctx.Logger().Error("failed to unmarshal rate limit state", "config", configName, "address", address, "error", err)
		return types.RateLimitState{}, false
	}
	return state, true
}

// SetRateLimitState sets the rate limit state
func (k Keeper) SetRateLimitState(ctx sdk.Context, state types.RateLimitState) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RateLimitKeyPrefix)
	key := k.getRateLimitKey(state.ConfigName, state.Address)
	bz, err := state.Marshal()
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal rate limit state")
	}
	store.Set(key, bz)
	return nil
}

// CheckRateLimit checks if a request is within rate limits
func (k Keeper) CheckRateLimit(ctx sdk.Context, configName, address, msgType string) error {
	params := k.GetParams(ctx)

	// Find the config
	var config *types.RateLimitConfig
	for _, c := range params.RateLimits {
		if c.Name == configName {
			config = &c
			break
		}
	}

	if config == nil || !config.Enabled {
		return nil // No limit or disabled
	}

	// Check if this message type is covered
	if len(config.MessageTypes) > 0 {
		found := false
		for _, mt := range config.MessageTypes {
			if mt == msgType {
				found = true
				break
			}
		}
		if !found {
			return nil // Message type not covered
		}
	}

	// Get or create state
	stateAddr := ""
	if config.PerAddress {
		stateAddr = address
	}

	state, found := k.GetRateLimitState(ctx, configName, stateAddr)
	if !found {
		state = types.RateLimitState{
			ConfigName:   configName,
			Address:      stateAddr,
			RequestCount: 0,
			WindowStart:  ctx.BlockTime(),
		}
	}

	// Check if window has expired
	windowDuration := time.Duration(config.WindowSeconds) * time.Second
	if ctx.BlockTime().Sub(state.WindowStart) >= windowDuration {
		// Reset window
		state.RequestCount = 0
		state.WindowStart = ctx.BlockTime()
	}

	// Check limit
	if state.RequestCount >= config.MaxRequests {
		return types.ErrRateLimitExceeded
	}

	// Increment and save
	state.RequestCount++
	if err := k.SetRateLimitState(ctx, state); err != nil {
		ctx.Logger().Error("failed to update rate limit state", "error", err)
		// Don't fail the request due to state storage error
	}

	return nil
}

// CheckAllRateLimits checks all applicable rate limits for a message
func (k Keeper) CheckAllRateLimits(ctx sdk.Context, sender, msgType string) error {
	params := k.GetParams(ctx)

	for _, config := range params.RateLimits {
		if !config.Enabled {
			continue
		}

		// Check if this config applies to this message type
		applies := len(config.MessageTypes) == 0
		if !applies {
			for _, mt := range config.MessageTypes {
				if mt == msgType {
					applies = true
					break
				}
			}
		}

		if applies {
			if err := k.CheckRateLimit(ctx, config.Name, sender, msgType); err != nil {
				return err
			}
		}
	}

	return nil
}

// ============================================================================
// Oracle Deviation Protection
// ============================================================================

// GetOracleDeviationConfig returns the deviation config for a denom
func (k Keeper) GetOracleDeviationConfig(ctx sdk.Context, denom string) (types.OracleDeviationConfig, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("oracle_deviation:"))
	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return types.OracleDeviationConfig{}, false
	}
	var config types.OracleDeviationConfig
	if err := json.Unmarshal(bz, &config); err != nil {
		ctx.Logger().Error("failed to unmarshal oracle deviation config", "denom", denom, "error", err)
		return types.OracleDeviationConfig{}, false
	}
	return config, true
}

// SetOracleDeviationConfig sets the deviation config for a denom
func (k Keeper) SetOracleDeviationConfig(ctx sdk.Context, config types.OracleDeviationConfig) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("oracle_deviation:"))
	bz, err := json.Marshal(config)
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal oracle deviation config")
	}
	store.Set([]byte(config.Denom), bz)
	return nil
}

// ============================================================================
// Liquidation Surge Protection
// ============================================================================

// GetLiquidationProtection returns the liquidation surge protection state
func (k Keeper) GetLiquidationProtection(ctx sdk.Context) types.LiquidationSurgeProtection {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("liquidation_protection"))
	if len(bz) == 0 {
		return types.DefaultLiquidationSurgeProtection()
	}
	var protection types.LiquidationSurgeProtection
	if err := json.Unmarshal(bz, &protection); err != nil {
		ctx.Logger().Error("failed to unmarshal liquidation protection", "error", err)
		return types.DefaultLiquidationSurgeProtection()
	}
	return protection
}

// SetLiquidationProtection sets the liquidation surge protection state
func (k Keeper) SetLiquidationProtection(ctx sdk.Context, protection types.LiquidationSurgeProtection) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := json.Marshal(protection)
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal liquidation protection")
	}
	store.Set([]byte("liquidation_protection"), bz)
	return nil
}

// CheckLiquidationAllowed checks if a liquidation is allowed under surge protection
func (k Keeper) CheckLiquidationAllowed(ctx sdk.Context, value sdkmath.Int) error {
	protection := k.GetLiquidationProtection(ctx)

	// Reset counters if new block
	if ctx.BlockHeight() > protection.LastResetHeight {
		protection.CurrentBlockLiquidations = 0
		protection.CurrentBlockValue = sdkmath.ZeroInt()
		protection.LastResetHeight = ctx.BlockHeight()
	}

	// Check limits
	if protection.CurrentBlockLiquidations >= protection.MaxLiquidationsPerBlock {
		return types.ErrLiquidationSurge
	}

	newValue := protection.CurrentBlockValue.Add(value)
	if newValue.GT(protection.MaxLiquidationValue) {
		return types.ErrLiquidationSurge
	}

	return nil
}

// RecordLiquidation records a liquidation for surge protection tracking
func (k Keeper) RecordLiquidation(ctx sdk.Context, value sdkmath.Int) {
	protection := k.GetLiquidationProtection(ctx)

	// Reset counters if new block
	if ctx.BlockHeight() > protection.LastResetHeight {
		protection.CurrentBlockLiquidations = 0
		protection.CurrentBlockValue = sdkmath.ZeroInt()
		protection.LastResetHeight = ctx.BlockHeight()
	}

	protection.CurrentBlockLiquidations++
	protection.CurrentBlockValue = protection.CurrentBlockValue.Add(value)

	if err := k.SetLiquidationProtection(ctx, protection); err != nil {
		ctx.Logger().Error("failed to record liquidation", "error", err)
	}
}

// ============================================================================
// Authority Check
// ============================================================================

// IsAuthorized checks if an address is authorized to control circuits
func (k Keeper) IsAuthorized(ctx sdk.Context, address string) bool {
	// Check if it's the module authority (governance)
	if address == k.authority {
		return true
	}

	// Check if it's in the params authorities list
	params := k.GetParams(ctx)
	for _, auth := range params.Authorities {
		if auth == address {
			return true
		}
	}

	return false
}

// ============================================================================
// Genesis
// ============================================================================

// InitGenesis initializes the module state from genesis
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) error {
	if state == nil {
		state = types.DefaultGenesis()
	}

	if err := k.SetParams(ctx, state.Params); err != nil {
		return errorsmod.Wrap(err, "failed to set params during genesis")
	}
	if err := k.SetCircuitState(ctx, state.CircuitState); err != nil {
		return errorsmod.Wrap(err, "failed to set circuit state during genesis")
	}

	for _, mc := range state.ModuleCircuits {
		if err := k.SetModuleCircuitState(ctx, mc); err != nil {
			return errorsmod.Wrapf(err, "failed to set module circuit state for %s", mc.ModuleName)
		}
	}

	for _, rls := range state.RateLimitStates {
		if err := k.SetRateLimitState(ctx, rls); err != nil {
			return errorsmod.Wrapf(err, "failed to set rate limit state for %s", rls.ConfigName)
		}
	}

	for _, odc := range state.OracleDeviationConfigs {
		if err := k.SetOracleDeviationConfig(ctx, odc); err != nil {
			return errorsmod.Wrapf(err, "failed to set oracle deviation config for %s", odc.Denom)
		}
	}

	if err := k.SetLiquidationProtection(ctx, state.LiquidationProtection); err != nil {
		return errorsmod.Wrap(err, "failed to set liquidation protection during genesis")
	}

	return nil
}

// ExportGenesis exports the module state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Params = k.GetParams(ctx)
	state.CircuitState = k.GetCircuitState(ctx)
	state.LiquidationProtection = k.GetLiquidationProtection(ctx)

	// Iterate module circuits
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ModuleCircuitKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var mc types.ModuleCircuitState
		if err := json.Unmarshal(iterator.Value(), &mc); err != nil {
			ctx.Logger().Error("failed to unmarshal module circuit during export", "error", err)
			continue
		}
		state.ModuleCircuits = append(state.ModuleCircuits, mc)
	}

	// Iterate rate limit states
	rlStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.RateLimitKeyPrefix)
	rlIterator := rlStore.Iterator(nil, nil)
	defer rlIterator.Close()
	for ; rlIterator.Valid(); rlIterator.Next() {
		var rls types.RateLimitState
		if err := json.Unmarshal(rlIterator.Value(), &rls); err != nil {
			ctx.Logger().Error("failed to unmarshal rate limit state during export", "error", err)
			continue
		}
		state.RateLimitStates = append(state.RateLimitStates, rls)
	}

	return state
}

// ============================================================================
// Context-based helpers for ante handler
// ============================================================================

// CheckCircuitBreakers is a helper for the ante handler to check all circuit breakers
func (k Keeper) CheckCircuitBreakers(ctx context.Context, moduleName, msgType, sender string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check global pause
	if k.IsGloballyPaused(sdkCtx) {
		return types.ErrGlobalPause
	}

	// Check module circuit
	if k.IsModuleCircuitOpen(sdkCtx, moduleName) {
		if k.IsMessageDisabled(sdkCtx, moduleName, msgType) {
			return types.ErrCircuitOpen
		}
	}

	// Check rate limits
	if err := k.CheckAllRateLimits(sdkCtx, sender, msgType); err != nil {
		return err
	}

	return nil
}
