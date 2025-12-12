package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/metrics/types"
)

// Keeper manages metrics state
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

// NewKeeper creates a new metrics keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// ============================================================================
// System Metrics
// ============================================================================

// GetSystemMetrics returns the current system metrics
func (k Keeper) GetSystemMetrics(ctx sdk.Context) types.SystemMetrics {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.MetricsKeyPrefix)
	if len(bz) == 0 {
		return types.DefaultSystemMetrics()
	}
	metrics, _ := types.UnmarshalSystemMetrics(bz)
	return metrics
}

// SetSystemMetrics sets the system metrics
func (k Keeper) SetSystemMetrics(ctx sdk.Context, metrics types.SystemMetrics) {
	store := ctx.KVStore(k.storeKey)
	bz, _ := metrics.Marshal()
	store.Set(types.MetricsKeyPrefix, bz)
}

// UpdateBlockMetrics updates block-related metrics
func (k Keeper) UpdateBlockMetrics(ctx sdk.Context) {
	metrics := k.GetSystemMetrics(ctx)

	// Calculate average block time
	if metrics.LastBlockHeight > 0 && !metrics.LastBlockTime.IsZero() {
		timeDiff := ctx.BlockTime().Sub(metrics.LastBlockTime)
		if metrics.AverageBlockTime == 0 {
			metrics.AverageBlockTime = timeDiff
		} else {
			// Exponential moving average
			alpha := 0.1
			metrics.AverageBlockTime = time.Duration(
				float64(metrics.AverageBlockTime)*(1-alpha) + float64(timeDiff)*alpha,
			)
		}
	}

	metrics.LastBlockHeight = ctx.BlockHeight()
	metrics.LastBlockTime = ctx.BlockTime()

	k.SetSystemMetrics(ctx, metrics)
}

// ============================================================================
// Counters
// ============================================================================

func (k Keeper) getCounterKey(name string, labels map[string]string) []byte {
	key := name
	if len(labels) > 0 {
		labelsJson, _ := json.Marshal(labels)
		key = fmt.Sprintf("%s:%s", name, string(labelsJson))
	}
	return []byte(key)
}

// GetCounter returns a counter value
func (k Keeper) GetCounter(ctx sdk.Context, name string, labels map[string]string) types.Counter {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CounterKeyPrefix)
	key := k.getCounterKey(name, labels)
	bz := store.Get(key)
	if len(bz) == 0 {
		return types.Counter{Name: name, Value: 0, Labels: labels}
	}
	counter, _ := types.UnmarshalCounter(bz)
	return counter
}

// SetCounter sets a counter value
func (k Keeper) SetCounter(ctx sdk.Context, counter types.Counter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CounterKeyPrefix)
	key := k.getCounterKey(counter.Name, counter.Labels)
	bz, _ := counter.Marshal()
	store.Set(key, bz)
}

// IncrementCounter increments a counter by 1
func (k Keeper) IncrementCounter(ctx sdk.Context, name string, labels map[string]string) {
	counter := k.GetCounter(ctx, name, labels)
	counter.Value++
	k.SetCounter(ctx, counter)
}

// AddToCounter adds a value to a counter
func (k Keeper) AddToCounter(ctx sdk.Context, name string, labels map[string]string, value uint64) {
	counter := k.GetCounter(ctx, name, labels)
	counter.Value += value
	k.SetCounter(ctx, counter)
}

// ============================================================================
// Gauges
// ============================================================================

func (k Keeper) getGaugeKey(name string, labels map[string]string) []byte {
	key := name
	if len(labels) > 0 {
		labelsJson, _ := json.Marshal(labels)
		key = fmt.Sprintf("%s:%s", name, string(labelsJson))
	}
	return []byte(key)
}

// GetGauge returns a gauge value
func (k Keeper) GetGauge(ctx sdk.Context, name string, labels map[string]string) types.Gauge {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GaugeKeyPrefix)
	key := k.getGaugeKey(name, labels)
	bz := store.Get(key)
	if len(bz) == 0 {
		return types.Gauge{Name: name, Value: sdkmath.LegacyZeroDec(), Labels: labels}
	}
	gauge, _ := types.UnmarshalGauge(bz)
	return gauge
}

// SetGauge sets a gauge value
func (k Keeper) SetGauge(ctx sdk.Context, gauge types.Gauge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GaugeKeyPrefix)
	key := k.getGaugeKey(gauge.Name, gauge.Labels)
	gauge.UpdatedAt = ctx.BlockTime()
	bz, _ := gauge.Marshal()
	store.Set(key, bz)
}

// UpdateGauge updates a gauge with a new value
func (k Keeper) UpdateGauge(ctx sdk.Context, name string, labels map[string]string, value sdkmath.LegacyDec) {
	gauge := types.Gauge{
		Name:      name,
		Value:     value,
		Labels:    labels,
		UpdatedAt: ctx.BlockTime(),
	}
	k.SetGauge(ctx, gauge)
}

// ============================================================================
// Module Health
// ============================================================================

// UpdateModuleHealth updates the health status of a module
func (k Keeper) UpdateModuleHealth(ctx sdk.Context, module string, status string, errorRate float64, latency float64) {
	metrics := k.GetSystemMetrics(ctx)

	if metrics.ModuleHealth == nil {
		metrics.ModuleHealth = make(map[string]*types.ModuleHealth)
	}

	health := metrics.ModuleHealth[module]
	if health == nil {
		health = &types.ModuleHealth{Module: module}
		metrics.ModuleHealth[module] = health
	}
	health.Status = status
	health.ErrorRate = errorRate
	health.Latency = latency
	health.Transactions++
	k.SetSystemMetrics(ctx, metrics)
}

// RecordModuleError records an error for a module
func (k Keeper) RecordModuleError(ctx sdk.Context, module string, errorMsg string) {
	metrics := k.GetSystemMetrics(ctx)

	if metrics.ModuleHealth == nil {
		metrics.ModuleHealth = make(map[string]*types.ModuleHealth)
	}

	health := metrics.ModuleHealth[module]
	if health == nil {
		health = &types.ModuleHealth{Module: module}
		metrics.ModuleHealth[module] = health
	}
	health.LastError = errorMsg
	health.LastErrorAt = ctx.BlockTime()

	// Update error rate (simple moving average)
	health.ErrorRate = (health.ErrorRate*float64(health.Transactions) + 1) / float64(health.Transactions+1)
	health.Transactions++

	if health.ErrorRate > 0.1 {
		health.Status = "unhealthy"
	} else if health.ErrorRate > 0.05 {
		health.Status = "degraded"
	} else {
		health.Status = "healthy"
	}

	k.SetSystemMetrics(ctx, metrics)
}

// ============================================================================
// Economic Metrics
// ============================================================================

// UpdateCollateralMetrics updates collateral-related metrics
func (k Keeper) UpdateCollateralMetrics(ctx sdk.Context, totalCollateral, totalDebt sdkmath.Int) {
	metrics := k.GetSystemMetrics(ctx)

	metrics.TotalCollateralValue = totalCollateral
	metrics.TotalDebtValue = totalDebt

	if !totalDebt.IsZero() {
		metrics.SystemCollateralRatio = sdkmath.LegacyNewDecFromInt(totalCollateral).Quo(
			sdkmath.LegacyNewDecFromInt(totalDebt),
		)
	} else {
		metrics.SystemCollateralRatio = sdkmath.LegacyZeroDec()
	}

	k.SetSystemMetrics(ctx, metrics)
}

// ============================================================================
// Settlement Metrics
// ============================================================================

// RecordSettlement records a settlement transaction
func (k Keeper) RecordSettlement(ctx sdk.Context, amount sdkmath.Int) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.TotalSettlements++
	metrics.TotalSettlementVolume = metrics.TotalSettlementVolume.Add(amount)
	k.SetSystemMetrics(ctx, metrics)
}

// UpdateActiveEscrows updates the count of active escrows
func (k Keeper) UpdateActiveEscrows(ctx sdk.Context, count uint64) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.ActiveEscrows = count
	k.SetSystemMetrics(ctx, metrics)
}

// UpdateActiveChannels updates the count of active payment channels
func (k Keeper) UpdateActiveChannels(ctx sdk.Context, count uint64) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.ActiveChannels = count
	k.SetSystemMetrics(ctx, metrics)
}

// ============================================================================
// Security Metrics
// ============================================================================

// RecordCircuitTrip records a circuit breaker trip
func (k Keeper) RecordCircuitTrip(ctx sdk.Context) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.CircuitTrips++
	k.SetSystemMetrics(ctx, metrics)

	k.IncrementCounter(ctx, "circuit_trips", nil)
}

// RecordRateLimitHit records a rate limit hit
func (k Keeper) RecordRateLimitHit(ctx sdk.Context, limitName string) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.RateLimitHits++
	k.SetSystemMetrics(ctx, metrics)

	k.IncrementCounter(ctx, "rate_limit_hits", map[string]string{"limit": limitName})
}

// RecordComplianceBlock records a compliance block
func (k Keeper) RecordComplianceBlock(ctx sdk.Context, reason string) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.ComplianceBlocks++
	k.SetSystemMetrics(ctx, metrics)

	k.IncrementCounter(ctx, "compliance_blocks", map[string]string{"reason": reason})
}

// ============================================================================
// Oracle Metrics
// ============================================================================

// RecordPriceUpdate records an oracle price update
func (k Keeper) RecordPriceUpdate(ctx sdk.Context, denom string) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.PricesUpdated++
	k.SetSystemMetrics(ctx, metrics)

	k.IncrementCounter(ctx, "price_updates", map[string]string{"denom": denom})
}

// UpdateStalePriceCount updates the count of stale prices
func (k Keeper) UpdateStalePriceCount(ctx sdk.Context, count uint64) {
	metrics := k.GetSystemMetrics(ctx)
	metrics.StalePriceCount = count
	k.SetSystemMetrics(ctx, metrics)
}

// ============================================================================
// Alerts
// ============================================================================

// GetAlertConfigs returns all alert configurations
func (k Keeper) GetAlertConfigs(ctx sdk.Context) []types.AlertConfig {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AlertKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var configs []types.AlertConfig
	for ; iterator.Valid(); iterator.Next() {
		var config types.AlertConfig
		json.Unmarshal(iterator.Value(), &config)
		configs = append(configs, config)
	}

	if len(configs) == 0 {
		return types.DefaultAlertConfigs()
	}

	return configs
}

// SetAlertConfig sets an alert configuration
func (k Keeper) SetAlertConfig(ctx sdk.Context, config types.AlertConfig) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AlertKeyPrefix)
	bz, _ := json.Marshal(config)
	store.Set([]byte(config.Name), bz)
}

// CheckAlerts checks all alert conditions and triggers alerts if needed
func (k Keeper) CheckAlerts(ctx sdk.Context) []types.Alert {
	configs := k.GetAlertConfigs(ctx)
	metrics := k.GetSystemMetrics(ctx)

	var triggeredAlerts []types.Alert

	for _, config := range configs {
		if !config.Enabled {
			continue
		}

		var value sdkmath.LegacyDec
		switch config.MetricName {
		case "system_collateral_ratio":
			value = metrics.SystemCollateralRatio
		case "stale_price_count":
			value = sdkmath.LegacyNewDec(int64(metrics.StalePriceCount))
		case "rate_limit_hits":
			value = sdkmath.LegacyNewDec(int64(metrics.RateLimitHits))
		case "circuit_trips":
			value = sdkmath.LegacyNewDec(int64(metrics.CircuitTrips))
		default:
			continue
		}

		triggered := false
		switch config.Condition {
		case types.AlertConditionGreaterThan:
			triggered = value.GT(config.Threshold)
		case types.AlertConditionLessThan:
			triggered = value.LT(config.Threshold)
		case types.AlertConditionEquals:
			triggered = value.Equal(config.Threshold)
		case types.AlertConditionNotEquals:
			triggered = !value.Equal(config.Threshold)
		}

		if triggered {
			alert := types.Alert{
				Id:          fmt.Sprintf("%s-%d", config.Name, ctx.BlockHeight()),
				ConfigName:  config.Name,
				MetricName:  config.MetricName,
				Value:       value,
				Threshold:   config.Threshold,
				Severity:    config.Severity,
				Message:     fmt.Sprintf("Alert: %s - value %s %s threshold %s", config.Name, value.String(), config.Condition, config.Threshold.String()),
				TriggeredAt: ctx.BlockTime(),
			}
			triggeredAlerts = append(triggeredAlerts, alert)

			// Emit event
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					"alert_triggered",
					sdk.NewAttribute("name", config.Name),
					sdk.NewAttribute("severity", string(config.Severity)),
					sdk.NewAttribute("value", value.String()),
					sdk.NewAttribute("threshold", config.Threshold.String()),
				),
			)
		}
	}

	return triggeredAlerts
}

// ============================================================================
// Genesis
// ============================================================================

// InitGenesis initializes the module state from genesis
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}

	k.SetSystemMetrics(ctx, state.SystemMetrics)

	for _, config := range state.AlertConfigs {
		k.SetAlertConfig(ctx, config)
	}
}

// ExportGenesis exports the module state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		SystemMetrics: k.GetSystemMetrics(ctx),
		AlertConfigs:  k.GetAlertConfigs(ctx),
		Alerts:        []types.Alert{},
	}
}

// ============================================================================
// Context helpers
// ============================================================================

// RecordTransaction records a transaction for metrics
func (k Keeper) RecordTransaction(ctx context.Context, module string, success bool, latencyMs float64) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if success {
		k.IncrementCounter(sdkCtx, "transactions_success", map[string]string{"module": module})
	} else {
		k.IncrementCounter(sdkCtx, "transactions_failed", map[string]string{"module": module})
	}

	// Update module health
	metrics := k.GetSystemMetrics(sdkCtx)
	if metrics.ModuleHealth == nil {
		metrics.ModuleHealth = make(map[string]*types.ModuleHealth)
	}

	health := metrics.ModuleHealth[module]
	if health == nil {
		health = &types.ModuleHealth{Module: module}
		metrics.ModuleHealth[module] = health
	}
	health.Transactions++

	// Update average latency
	health.Latency = (health.Latency*float64(health.Transactions-1) + latencyMs) / float64(health.Transactions)

	if !success {
		health.ErrorRate = (health.ErrorRate*float64(health.Transactions-1) + 1) / float64(health.Transactions)
	} else {
		health.ErrorRate = (health.ErrorRate * float64(health.Transactions-1)) / float64(health.Transactions)
	}

	k.SetSystemMetrics(sdkCtx, metrics)
}
