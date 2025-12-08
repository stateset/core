package keeper

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/oracle/types"
)

var (
	paramsKey           = []byte("params")
	configKeyPrefix     = []byte("config:")
	providerKeyPrefix   = []byte("provider:")
	priceHistoryPrefix  = []byte("history:")
)

// Keeper provides state access to the oracle module.
type Keeper struct {
	storeKey  storetypes.StoreKey
	cdc       codec.BinaryCodec
	authority string
}

// NewKeeper creates a new oracle keeper instance.
func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority string) Keeper {
	return Keeper{
		storeKey:  key,
		cdc:       cdc,
		authority: authority,
	}
}

// GetAuthority returns the current authority allowed to update prices.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// SetAuthority updates the address allowed to submit price updates.
func (k *Keeper) SetAuthority(authority string) {
	k.authority = authority
}

// ============================================================================
// Parameters
// ============================================================================

// GetParams returns the oracle module parameters
func (k Keeper) GetParams(ctx context.Context) types.OracleParams {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz := store.Get(paramsKey)
	if len(bz) == 0 {
		return types.DefaultOracleParams()
	}
	var params types.OracleParams
	if err := json.Unmarshal(bz, &params); err != nil {
		// Log error and return defaults for safety
		sdkCtx.Logger().Error("failed to unmarshal oracle params", "error", err)
		return types.DefaultOracleParams()
	}
	return params
}

// SetParams sets the oracle module parameters
func (k Keeper) SetParams(ctx context.Context, params types.OracleParams) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	bz, err := json.Marshal(params)
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal oracle params")
	}
	store.Set(paramsKey, bz)
	return nil
}

// ============================================================================
// Oracle Config (per-denom)
// ============================================================================

// GetOracleConfig returns the oracle config for a denom
func (k Keeper) GetOracleConfig(ctx context.Context, denom string) (types.OracleConfig, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), configKeyPrefix)
	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return types.OracleConfig{}, false
	}
	var config types.OracleConfig
	if err := json.Unmarshal(bz, &config); err != nil {
		sdkCtx.Logger().Error("failed to unmarshal oracle config", "denom", denom, "error", err)
		return types.OracleConfig{}, false
	}
	return config, true
}

// SetOracleConfig sets the oracle config for a denom
func (k Keeper) SetOracleConfig(ctx context.Context, config types.OracleConfig) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), configKeyPrefix)
	bz, err := json.Marshal(config)
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal oracle config")
	}
	store.Set([]byte(config.Denom), bz)
	return nil
}

// ============================================================================
// Oracle Providers
// ============================================================================

// GetProvider returns an oracle provider by address
func (k Keeper) GetProvider(ctx context.Context, address string) (types.OracleProvider, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), providerKeyPrefix)
	bz := store.Get([]byte(address))
	if len(bz) == 0 {
		return types.OracleProvider{}, false
	}
	var provider types.OracleProvider
	if err := json.Unmarshal(bz, &provider); err != nil {
		sdkCtx.Logger().Error("failed to unmarshal oracle provider", "address", address, "error", err)
		return types.OracleProvider{}, false
	}
	return provider, true
}

// SetProvider sets an oracle provider
func (k Keeper) SetProvider(ctx context.Context, provider types.OracleProvider) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), providerKeyPrefix)
	bz, err := json.Marshal(provider)
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal oracle provider")
	}
	store.Set([]byte(provider.Address), bz)
	return nil
}

// RemoveProvider removes an oracle provider
func (k Keeper) RemoveProvider(ctx context.Context, address string) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), providerKeyPrefix)
	store.Delete([]byte(address))
}

// IterateProviders iterates over all providers
func (k Keeper) IterateProviders(ctx context.Context, cb func(types.OracleProvider) bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), providerKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var provider types.OracleProvider
		if err := json.Unmarshal(iterator.Value(), &provider); err != nil {
			sdkCtx.Logger().Error("failed to unmarshal oracle provider during iteration", "error", err)
			continue
		}
		if cb(provider) {
			break
		}
	}
}

// IsAuthorizedProvider checks if an address is an authorized oracle provider
func (k Keeper) IsAuthorizedProvider(ctx context.Context, address string) bool {
	// Check if it's the governance authority
	if address == k.authority {
		return true
	}

	// Check if it's a registered active provider
	provider, found := k.GetProvider(ctx, address)
	if !found {
		return false
	}

	return provider.IsActive && !provider.Slashed
}

// ============================================================================
// Price Operations with Security
// ============================================================================

// SetPrice records the price for a denom.
func (k Keeper) SetPrice(ctx context.Context, price types.Price) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.PriceKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&price)
	store.Set([]byte(price.Denom), bz)
}

// SetPriceWithValidation validates and sets a price with security checks
func (k Keeper) SetPriceWithValidation(ctx context.Context, provider string, denom string, amount sdkmath.LegacyDec) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get or create config
	config, found := k.GetOracleConfig(ctx, denom)
	if !found {
		config = types.DefaultOracleConfig(denom)
	}

	if !config.Enabled {
		return types.ErrConfigDisabled
	}

	// Get existing price for deviation check
	existingPrice, hasExisting := k.GetPrice(ctx, denom)

	// Check update frequency
	if hasExisting && !existingPrice.UpdatedAt.IsZero() {
		minInterval := time.Duration(config.MinUpdateIntervalSeconds) * time.Second
		if sdkCtx.BlockTime().Sub(existingPrice.UpdatedAt) < minInterval {
			return types.ErrUpdateTooFrequent
		}
	}

	// Check deviation
	if hasExisting && !existingPrice.Amount.IsZero() {
		deviation := types.CalculateDeviation(existingPrice.Amount, amount)
		maxDeviation := sdkmath.LegacyNewDec(int64(config.MaxDeviationBps))

		if deviation.GT(maxDeviation) {
			// Record provider failure
			k.recordProviderFailure(ctx, provider)

			// Emit warning event
			sdkCtx.EventManager().EmitEvent(
				sdk.NewEvent(
					"oracle_deviation_rejected",
					sdk.NewAttribute("denom", denom),
					sdk.NewAttribute("old_price", existingPrice.Amount.String()),
					sdk.NewAttribute("new_price", amount.String()),
					sdk.NewAttribute("deviation_bps", deviation.String()),
					sdk.NewAttribute("max_deviation_bps", maxDeviation.String()),
					sdk.NewAttribute("provider", provider),
				),
			)

			return errorsmod.Wrapf(types.ErrDeviationTooLarge,
				"deviation %s bps exceeds max %s bps",
				deviation.String(), maxDeviation.String())
		}
	}

	// Create and store new price
	newPrice := types.Price{
		Denom:       denom,
		Amount:      amount,
		LastUpdater: provider,
		LastHeight:  sdkCtx.BlockHeight(),
		UpdatedAt:   sdkCtx.BlockTime(),
	}

	k.SetPrice(ctx, newPrice)

	// Record in price history
	k.recordPriceHistory(ctx, newPrice)

	// Update provider stats
	k.recordProviderSuccess(ctx, provider)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"price_updated",
			sdk.NewAttribute("denom", denom),
			sdk.NewAttribute("price", amount.String()),
			sdk.NewAttribute("provider", provider),
		),
	)

	return nil
}

// GetPrice returns the price for a denom if present.
func (k Keeper) GetPrice(ctx context.Context, denom string) (types.Price, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.PriceKeyPrefix)
	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return types.Price{}, false
	}

	var price types.Price
	types.ModuleCdc.MustUnmarshalJSON(bz, &price)
	return price, true
}

// GetPriceWithStalenessCheck returns a price with staleness validation
func (k Keeper) GetPriceWithStalenessCheck(ctx context.Context, denom string) (types.Price, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	price, found := k.GetPrice(ctx, denom)
	if !found {
		return types.Price{}, types.ErrPriceNotFound
	}

	// Get config for staleness threshold
	config, found := k.GetOracleConfig(ctx, denom)
	if !found {
		config = types.DefaultOracleConfig(denom)
	}

	stalenessThreshold := time.Duration(config.StalenessThresholdSeconds) * time.Second
	if price.IsStale(sdkCtx.BlockTime(), stalenessThreshold) {
		return price, errorsmod.Wrapf(types.ErrPriceStale,
			"price for %s is stale (last updated: %s)",
			denom, price.UpdatedAt.String())
	}

	return price, nil
}

// IteratePrices iterates over stored prices.
func (k Keeper) IteratePrices(ctx context.Context, cb func(price types.Price) bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.PriceKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var price types.Price
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &price)
		if cb(price) {
			break
		}
	}
}

// DeletePrice removes a stored price entry.
func (k Keeper) DeletePrice(ctx context.Context, denom string) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.PriceKeyPrefix)
	store.Delete([]byte(denom))
}

// GetPriceDec returns the sdk.Dec price or error if not found.
func (k Keeper) GetPriceDec(ctx context.Context, denom string) (sdkmath.LegacyDec, error) {
	price, ok := k.GetPrice(ctx, denom)
	if !ok {
		return sdkmath.LegacyDec{}, types.ErrPriceNotFound
	}
	return price.Amount, nil
}

// GetPriceDecSafe returns the price with staleness check
func (k Keeper) GetPriceDecSafe(ctx context.Context, denom string) (sdkmath.LegacyDec, error) {
	price, err := k.GetPriceWithStalenessCheck(ctx, denom)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	return price.Amount, nil
}

// ============================================================================
// Price History
// ============================================================================

// recordPriceHistory records a price point in history
func (k Keeper) recordPriceHistory(ctx context.Context, price types.Price) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(ctx)

	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), priceHistoryPrefix)

	var history types.PriceHistory
	bz := store.Get([]byte(price.Denom))
	if len(bz) > 0 {
		if err := json.Unmarshal(bz, &history); err != nil {
			sdkCtx.Logger().Error("failed to unmarshal price history", "denom", price.Denom, "error", err)
			// Initialize new history on error
			history = types.PriceHistory{
				Denom:   price.Denom,
				Prices:  []types.PricePoint{},
				MaxSize: params.PriceHistorySize,
			}
		}
	} else {
		history = types.PriceHistory{
			Denom:   price.Denom,
			Prices:  []types.PricePoint{},
			MaxSize: params.PriceHistorySize,
		}
	}

	// Add new price point
	history.Prices = append(history.Prices, types.PricePoint{
		Amount:    price.Amount,
		Timestamp: price.UpdatedAt,
		Height:    price.LastHeight,
	})

	// Trim to max size
	if uint32(len(history.Prices)) > history.MaxSize {
		history.Prices = history.Prices[len(history.Prices)-int(history.MaxSize):]
	}

	historyBz, err := json.Marshal(history)
	if err != nil {
		sdkCtx.Logger().Error("failed to marshal price history", "denom", price.Denom, "error", err)
		return
	}
	store.Set([]byte(price.Denom), historyBz)
}

// GetPriceHistory returns the price history for a denom
func (k Keeper) GetPriceHistory(ctx context.Context, denom string) (types.PriceHistory, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), priceHistoryPrefix)

	bz := store.Get([]byte(denom))
	if len(bz) == 0 {
		return types.PriceHistory{}, false
	}

	var history types.PriceHistory
	if err := json.Unmarshal(bz, &history); err != nil {
		sdkCtx.Logger().Error("failed to unmarshal price history", "denom", denom, "error", err)
		return types.PriceHistory{}, false
	}
	return history, true
}

// ============================================================================
// Provider Stats
// ============================================================================

func (k Keeper) recordProviderSuccess(ctx context.Context, address string) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	provider, found := k.GetProvider(ctx, address)
	if !found {
		return
	}

	provider.TotalSubmissions++
	provider.SuccessfulSubmissions++
	if err := k.SetProvider(ctx, provider); err != nil {
		sdkCtx.Logger().Error("failed to update provider stats", "address", address, "error", err)
	}
}

func (k Keeper) recordProviderFailure(ctx context.Context, address string) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	provider, found := k.GetProvider(ctx, address)
	if !found {
		return
	}

	provider.TotalSubmissions++

	// Check if provider should be slashed
	if provider.TotalSubmissions > 10 {
		successRate := float64(provider.SuccessfulSubmissions) / float64(provider.TotalSubmissions)
		if successRate < 0.5 {
			provider.Slashed = true
			provider.SlashCount++
			provider.IsActive = false

			sdkCtx.EventManager().EmitEvent(
				sdk.NewEvent(
					"oracle_provider_slashed",
					sdk.NewAttribute("address", address),
					sdk.NewAttribute("success_rate", fmt.Sprintf("%.0f%%", successRate*100)),
				),
			)
		}
	}

	if err := k.SetProvider(ctx, provider); err != nil {
		sdkCtx.Logger().Error("failed to update provider stats", "address", address, "error", err)
	}
}

// SlashProvider manually slashes an oracle provider
func (k Keeper) SlashProvider(ctx context.Context, address, reason string) error {
	provider, found := k.GetProvider(ctx, address)
	if !found {
		return types.ErrProviderNotFound
	}

	provider.Slashed = true
	provider.SlashCount++
	provider.IsActive = false
	if err := k.SetProvider(ctx, provider); err != nil {
		return errorsmod.Wrap(err, "failed to update slashed provider")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"oracle_provider_slashed",
			sdk.NewAttribute("address", address),
			sdk.NewAttribute("reason", reason),
		),
	)

	return nil
}

// ============================================================================
// Genesis
// ============================================================================

// ExportGenesis exports the module state.
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Authority = k.authority
	k.IteratePrices(ctx, func(price types.Price) bool {
		genesis.Prices = append(genesis.Prices, price)
		return false
	})
	return genesis
}

// InitGenesis initializes state from genesis configuration.
func (k Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) {
	if data != nil {
		k.authority = data.Authority
		for _, price := range data.Prices {
			k.SetPrice(ctx, price)
		}
	}
}

// mustWriteUint64 writes a uint64 into bytes (utility for future use).
func mustWriteUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
