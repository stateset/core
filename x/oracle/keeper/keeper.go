package keeper

import (
	"fmt"
	"time"
	
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	
	"github.com/stateset/core/x/oracle/types"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	paramstore paramtypes.Subspace
	authority  string
	
	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	authority string,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
) *Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		authority:     authority,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

// GetAuthority returns the module's authority
func (k Keeper) GetAuthority() string {
	return k.authority
}

// SetPriceFeed stores a price feed
func (k Keeper) SetPriceFeed(ctx sdk.Context, feed types.PriceFeed) error {
	if err := feed.Validate(); err != nil {
		return err
	}
	
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PriceFeedKey + feed.FeedID)
	value := k.cdc.MustMarshal(&feed)
	store.Set(key, value)
	
	// Update aggregated price
	if err := k.UpdateAggregatedPrice(ctx, feed.Asset); err != nil {
		return err
	}
	
	// Store in history
	if err := k.AddPriceHistory(ctx, feed.Asset, feed.Price, feed.Volume); err != nil {
		return err
	}
	
	return nil
}

// GetPriceFeed retrieves a price feed by ID
func (k Keeper) GetPriceFeed(ctx sdk.Context, feedID string) (types.PriceFeed, bool) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PriceFeedKey + feedID)
	
	value := store.Get(key)
	if value == nil {
		return types.PriceFeed{}, false
	}
	
	var feed types.PriceFeed
	k.cdc.MustUnmarshal(value, &feed)
	return feed, true
}

// GetAllPriceFeeds returns all price feeds
func (k Keeper) GetAllPriceFeeds(ctx sdk.Context) []types.PriceFeed {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, []byte(types.PriceFeedKey))
	defer iterator.Close()
	
	var feeds []types.PriceFeed
	for ; iterator.Valid(); iterator.Next() {
		var feed types.PriceFeed
		k.cdc.MustUnmarshal(iterator.Value(), &feed)
		feeds = append(feeds, feed)
	}
	
	return feeds
}

// GetActivePriceFeedsForAsset returns active price feeds for a specific asset
func (k Keeper) GetActivePriceFeedsForAsset(ctx sdk.Context, asset string) []types.PriceFeed {
	allFeeds := k.GetAllPriceFeeds(ctx)
	currentTime := ctx.BlockTime()
	
	var activeFeeds []types.PriceFeed
	for _, feed := range allFeeds {
		if feed.Asset == asset && !feed.IsExpired(currentTime) {
			activeFeeds = append(activeFeeds, feed)
		}
	}
	
	return activeFeeds
}

// SetOracleProvider stores an oracle provider
func (k Keeper) SetOracleProvider(ctx sdk.Context, provider types.OracleProvider) error {
	if err := provider.Validate(); err != nil {
		return err
	}
	
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.OracleProviderKey + provider.Address)
	value := k.cdc.MustMarshal(&provider)
	store.Set(key, value)
	
	return nil
}

// GetOracleProvider retrieves an oracle provider by address
func (k Keeper) GetOracleProvider(ctx sdk.Context, address string) (types.OracleProvider, bool) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.OracleProviderKey + address)
	
	value := store.Get(key)
	if value == nil {
		return types.OracleProvider{}, false
	}
	
	var provider types.OracleProvider
	k.cdc.MustUnmarshal(value, &provider)
	return provider, true
}

// GetAllOracleProviders returns all oracle providers
func (k Keeper) GetAllOracleProviders(ctx sdk.Context) []types.OracleProvider {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, []byte(types.OracleProviderKey))
	defer iterator.Close()
	
	var providers []types.OracleProvider
	for ; iterator.Valid(); iterator.Next() {
		var provider types.OracleProvider
		k.cdc.MustUnmarshal(iterator.Value(), &provider)
		providers = append(providers, provider)
	}
	
	return providers
}

// GetActiveOracleProviders returns only active oracle providers
func (k Keeper) GetActiveOracleProviders(ctx sdk.Context) []types.OracleProvider {
	allProviders := k.GetAllOracleProviders(ctx)
	
	var activeProviders []types.OracleProvider
	for _, provider := range allProviders {
		if provider.Active {
			activeProviders = append(activeProviders, provider)
		}
	}
	
	return activeProviders
}

// UpdateAggregatedPrice updates the aggregated price for an asset
func (k Keeper) UpdateAggregatedPrice(ctx sdk.Context, asset string) error {
	feeds := k.GetActivePriceFeedsForAsset(ctx, asset)
	if len(feeds) == 0 {
		return fmt.Errorf("no active price feeds for asset %s", asset)
	}
	
	providers := k.GetActiveOracleProviders(ctx)
	providerMap := make(map[string]types.OracleProvider)
	for _, provider := range providers {
		providerMap[provider.Address] = provider
	}
	
	// Calculate aggregated values
	weightedAvg := types.CalculateWeightedAverage(feeds, providerMap)
	median := types.CalculateMedianPrice(feeds)
	stdDev := types.CalculateStandardDeviation(feeds, weightedAvg)
	
	// Calculate confidence based on number of providers and deviation
	confidence := sdk.NewDec(int64(len(feeds))).Quo(sdk.NewDec(int64(len(providers))))
	if stdDev.GT(sdk.NewDecWithPrec(1, 2)) { // If std dev > 0.01
		confidence = confidence.Mul(sdk.NewDecWithPrec(9, 1)) // Reduce confidence by 10%
	}
	
	aggregated := types.AggregatedPrice{
		Asset:            asset,
		Price:            weightedAvg,
		MedianPrice:      median,
		WeightedAvgPrice: weightedAvg,
		StandardDev:      stdDev,
		Confidence:       confidence,
		NumProviders:     uint32(len(feeds)),
		LastUpdate:       ctx.BlockTime(),
		NextUpdate:       ctx.BlockTime().Add(5 * time.Minute), // Default 5 minute update interval
	}
	
	// Store aggregated price
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AggregatedPriceKey + asset)
	value := k.cdc.MustMarshal(&aggregated)
	store.Set(key, value)
	
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"price_updated",
			sdk.NewAttribute("asset", asset),
			sdk.NewAttribute("price", weightedAvg.String()),
			sdk.NewAttribute("median", median.String()),
			sdk.NewAttribute("confidence", confidence.String()),
			sdk.NewAttribute("providers", fmt.Sprintf("%d", len(feeds))),
		),
	)
	
	return nil
}

// GetAggregatedPrice retrieves the aggregated price for an asset
func (k Keeper) GetAggregatedPrice(ctx sdk.Context, asset string) (types.AggregatedPrice, bool) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AggregatedPriceKey + asset)
	
	value := store.Get(key)
	if value == nil {
		return types.AggregatedPrice{}, false
	}
	
	var price types.AggregatedPrice
	k.cdc.MustUnmarshal(value, &price)
	return price, true
}

// GetPrice returns the current price for an asset (simplified interface)
func (k Keeper) GetPrice(ctx sdk.Context, asset string) (sdk.Dec, error) {
	aggregated, found := k.GetAggregatedPrice(ctx, asset)
	if !found {
		return sdk.ZeroDec(), fmt.Errorf("no price available for asset %s", asset)
	}
	
	// Check if price is stale
	if ctx.BlockTime().After(aggregated.NextUpdate) {
		return sdk.ZeroDec(), fmt.Errorf("price for asset %s is stale", asset)
	}
	
	return aggregated.Price, nil
}

// AddPriceHistory adds a price point to historical data
func (k Keeper) AddPriceHistory(ctx sdk.Context, asset string, price sdk.Dec, volume sdk.Int) error {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PriceHistoryKey + asset)
	
	var history types.PriceHistory
	value := store.Get(key)
	if value != nil {
		k.cdc.MustUnmarshal(value, &history)
	} else {
		history = types.PriceHistory{
			Asset:  asset,
			Prices: []types.PricePoint{},
		}
	}
	
	// Add new price point
	history.Prices = append(history.Prices, types.PricePoint{
		Price:     price,
		Timestamp: ctx.BlockTime(),
		Volume:    volume,
	})
	
	// Keep only last 1000 price points
	if len(history.Prices) > 1000 {
		history.Prices = history.Prices[len(history.Prices)-1000:]
	}
	
	history.UpdatedAt = ctx.BlockTime()
	
	// Store updated history
	updatedValue := k.cdc.MustMarshal(&history)
	store.Set(key, updatedValue)
	
	return nil
}

// GetPriceHistory retrieves price history for an asset
func (k Keeper) GetPriceHistory(ctx sdk.Context, asset string, limit uint32) (types.PriceHistory, bool) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PriceHistoryKey + asset)
	
	value := store.Get(key)
	if value == nil {
		return types.PriceHistory{}, false
	}
	
	var history types.PriceHistory
	k.cdc.MustUnmarshal(value, &history)
	
	// Apply limit if specified
	if limit > 0 && uint32(len(history.Prices)) > limit {
		history.Prices = history.Prices[len(history.Prices)-int(limit):]
	}
	
	return history, true
}

// RemoveOracleProvider removes an oracle provider
func (k Keeper) RemoveOracleProvider(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.OracleProviderKey + address)
	store.Delete(key)
}

// UpdateProviderReputation updates an oracle provider's reputation
func (k Keeper) UpdateProviderReputation(ctx sdk.Context, address string, adjustment sdk.Dec) error {
	provider, found := k.GetOracleProvider(ctx, address)
	if !found {
		return fmt.Errorf("oracle provider %s not found", address)
	}
	
	newReputation := provider.Reputation.Add(adjustment)
	if newReputation.LT(sdk.ZeroDec()) {
		newReputation = sdk.ZeroDec()
	} else if newReputation.GT(sdk.OneDec()) {
		newReputation = sdk.OneDec()
	}
	
	provider.Reputation = newReputation
	return k.SetOracleProvider(ctx, provider)
}

// ValidatePriceSubmission validates a price submission against deviation thresholds
func (k Keeper) ValidatePriceSubmission(ctx sdk.Context, asset string, newPrice sdk.Dec, provider string) error {
	// Get current aggregated price
	currentPrice, found := k.GetAggregatedPrice(ctx, asset)
	if !found {
		// First price submission for this asset
		return nil
	}
	
	// Get provider settings
	oracleProvider, found := k.GetOracleProvider(ctx, provider)
	if !found {
		return fmt.Errorf("oracle provider %s not registered", provider)
	}
	
	// Calculate deviation
	deviation := newPrice.Sub(currentPrice.Price).Abs().Quo(currentPrice.Price)
	
	// Check if deviation exceeds threshold
	if deviation.GT(oracleProvider.MaxDeviation) {
		// Penalize reputation for excessive deviation
		k.UpdateProviderReputation(ctx, provider, sdk.NewDecWithPrec(-1, 2)) // -0.01 reputation
		return fmt.Errorf("price deviation %s exceeds maximum allowed %s", deviation, oracleProvider.MaxDeviation)
	}
	
	// Reward good submissions
	if deviation.LT(sdk.NewDecWithPrec(1, 2)) { // Less than 1% deviation
		k.UpdateProviderReputation(ctx, provider, sdk.NewDecWithPrec(1, 3)) // +0.001 reputation
	}
	
	return nil
}

// GetParams returns the oracle module parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the oracle module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// GetAllAggregatedAssets returns all assets with aggregated prices
func (k Keeper) GetAllAggregatedAssets(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, []byte(types.AggregatedPriceKey))
	defer iterator.Close()
	
	var assets []string
	for ; iterator.Valid(); iterator.Next() {
		// Extract asset from key
		key := string(iterator.Key())
		asset := key[len(types.AggregatedPriceKey):]
		assets = append(assets, asset)
	}
	
	return assets
}

// GetAllAggregatedPrices returns all aggregated prices
func (k Keeper) GetAllAggregatedPrices(ctx sdk.Context) []types.AggregatedPrice {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, []byte(types.AggregatedPriceKey))
	defer iterator.Close()
	
	var prices []types.AggregatedPrice
	for ; iterator.Valid(); iterator.Next() {
		var price types.AggregatedPrice
		k.cdc.MustUnmarshal(iterator.Value(), &price)
		prices = append(prices, price)
	}
	
	return prices
}

// SetAggregatedPrice stores an aggregated price
func (k Keeper) SetAggregatedPrice(ctx sdk.Context, price types.AggregatedPrice) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AggregatedPriceKey + price.Asset)
	value := k.cdc.MustMarshal(&price)
	store.Set(key, value)
}

// GetAllPriceHistories returns all price histories
func (k Keeper) GetAllPriceHistories(ctx sdk.Context) []types.PriceHistory {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, []byte(types.PriceHistoryKey))
	defer iterator.Close()
	
	var histories []types.PriceHistory
	for ; iterator.Valid(); iterator.Next() {
		var history types.PriceHistory
		k.cdc.MustUnmarshal(iterator.Value(), &history)
		histories = append(histories, history)
	}
	
	return histories
}

// SetPriceHistory stores a price history
func (k Keeper) SetPriceHistory(ctx sdk.Context, history types.PriceHistory) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PriceHistoryKey + history.Asset)
	value := k.cdc.MustMarshal(&history)
	store.Set(key, value)
}

// RemovePriceFeed removes a price feed
func (k Keeper) RemovePriceFeed(ctx sdk.Context, feedID string) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PriceFeedKey + feedID)
	store.Delete(key)
}