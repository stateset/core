package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"cosmossdk.io/log"

	"github.com/stateset/core/x/ssusd/types"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	ststKeeper    types.STSTKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	ststKeeper types.STSTKeeper,
) *Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		ststKeeper:    ststKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetSTSTPrice(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.OraclePriceKey("stst"))
	if bz == nil {
		return sdk.NewDec(1)
	}
	
	var price types.OraclePrice
	k.cdc.MustUnmarshal(bz, &price)
	return price.Price
}

func (k Keeper) SetOraclePrice(ctx sdk.Context, price types.OraclePrice) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&price)
	store.Set(types.OraclePriceKey(price.Asset), bz)
}

func (k Keeper) GetOraclePrice(ctx sdk.Context, asset string) (types.OraclePrice, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.OraclePriceKey(asset))
	if bz == nil {
		return types.OraclePrice{}, false
	}
	
	var price types.OraclePrice
	k.cdc.MustUnmarshal(bz, &price)
	return price, true
}

func (k Keeper) SetCollateralPosition(ctx sdk.Context, position types.CollateralPosition) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&position)
	store.Set(types.CollateralPositionKey(position.Owner), bz)
}

func (k Keeper) GetCollateralPosition(ctx sdk.Context, owner string) (types.CollateralPosition, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.CollateralPositionKey(owner))
	if bz == nil {
		return types.CollateralPosition{}, false
	}
	
	var position types.CollateralPosition
	k.cdc.MustUnmarshal(bz, &position)
	return position, true
}

func (k Keeper) DeleteCollateralPosition(ctx sdk.Context, owner string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.CollateralPositionKey(owner))
}

func (k Keeper) IterateCollateralPositions(ctx sdk.Context, cb func(position types.CollateralPosition) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.CollateralPositionKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var position types.CollateralPosition
		k.cdc.MustUnmarshal(iterator.Value(), &position)
		if cb(position) {
			break
		}
	}
}

func (k Keeper) SetAgentWallet(ctx sdk.Context, wallet types.AgentWallet) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&wallet)
	store.Set(types.AgentWalletKey(wallet.AgentId), bz)
}

func (k Keeper) GetAgentWallet(ctx sdk.Context, agentID string) (types.AgentWallet, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AgentWalletKey(agentID))
	if bz == nil {
		return types.AgentWallet{}, false
	}
	
	var wallet types.AgentWallet
	k.cdc.MustUnmarshal(bz, &wallet)
	return wallet, true
}

func (k Keeper) IterateAgentWallets(ctx sdk.Context, cb func(wallet types.AgentWallet) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AgentWalletKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var wallet types.AgentWallet
		k.cdc.MustUnmarshal(iterator.Value(), &wallet)
		if cb(wallet) {
			break
		}
	}
}

func (k Keeper) SetLiquidationAuction(ctx sdk.Context, auction types.LiquidationAuction) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&auction)
	store.Set(types.LiquidationAuctionKey(auction.Id), bz)
}

func (k Keeper) GetLiquidationAuction(ctx sdk.Context, auctionID string) (types.LiquidationAuction, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LiquidationAuctionKey(auctionID))
	if bz == nil {
		return types.LiquidationAuction{}, false
	}
	
	var auction types.LiquidationAuction
	k.cdc.MustUnmarshal(bz, &auction)
	return auction, true
}

func (k Keeper) DeleteLiquidationAuction(ctx sdk.Context, auctionID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.LiquidationAuctionKey(auctionID))
}

func (k Keeper) IterateLiquidationAuctions(ctx sdk.Context, cb func(auction types.LiquidationAuction) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LiquidationAuctionKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var auction types.LiquidationAuction
		k.cdc.MustUnmarshal(iterator.Value(), &auction)
		if cb(auction) {
			break
		}
	}
}

func (k Keeper) GetStabilityPool(ctx sdk.Context) types.StabilityPool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.StabilityPoolKey)
	if bz == nil {
		return types.StabilityPool{
			TotalDeposits:     sdk.NewCoin("ssusd", sdk.ZeroInt()),
			TotalDebtAbsorbed: sdk.NewCoin("ssusd", sdk.ZeroInt()),
			Providers:         []types.StabilityProvider{},
		}
	}
	
	var pool types.StabilityPool
	k.cdc.MustUnmarshal(bz, &pool)
	return pool
}

func (k Keeper) SetStabilityPool(ctx sdk.Context, pool types.StabilityPool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pool)
	store.Set(types.StabilityPoolKey, bz)
}

func (k Keeper) GetSystemMetrics(ctx sdk.Context) types.SystemMetrics {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SystemMetricsKey)
	if bz == nil {
		return types.SystemMetrics{
			TotalSupply:           sdk.NewCoin("ssusd", sdk.ZeroInt()),
			TotalCollateral:       sdk.NewCoin("stst", sdk.ZeroInt()),
			GlobalCollateralRatio: sdk.NewDec(0),
			TotalPositions:        0,
			TotalLiquidations:     0,
			TotalAgents:           0,
			AverageCollateralRatio: sdk.NewDec(0),
		}
	}
	
	var metrics types.SystemMetrics
	k.cdc.MustUnmarshal(bz, &metrics)
	return metrics
}

func (k Keeper) SetSystemMetrics(ctx sdk.Context, metrics types.SystemMetrics) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&metrics)
	store.Set(types.SystemMetricsKey, bz)
}

func (k Keeper) UpdateSystemMetrics(ctx sdk.Context) {
	var totalSupply sdk.Int = sdk.ZeroInt()
	var totalCollateral sdk.Int = sdk.ZeroInt()
	var totalPositions uint64 = 0
	var collateralRatioSum sdk.Dec = sdk.ZeroDec()
	
	k.IterateCollateralPositions(ctx, func(position types.CollateralPosition) bool {
		totalSupply = totalSupply.Add(position.Debt.Amount)
		totalCollateral = totalCollateral.Add(position.Collateral.Amount)
		totalPositions++
		collateralRatioSum = collateralRatioSum.Add(position.CollateralizationRatio)
		return false
	})
	
	var totalAgents uint64 = 0
	k.IterateAgentWallets(ctx, func(wallet types.AgentWallet) bool {
		if wallet.IsActive {
			totalAgents++
		}
		return false
	})
	
	globalRatio := sdk.ZeroDec()
	avgRatio := sdk.ZeroDec()
	
	if !totalSupply.IsZero() {
		ststPrice := k.GetSTSTPrice(ctx)
		globalRatio = sdk.NewDecFromInt(totalCollateral).Mul(ststPrice).Quo(sdk.NewDecFromInt(totalSupply))
	}
	
	if totalPositions > 0 {
		avgRatio = collateralRatioSum.Quo(sdk.NewDec(int64(totalPositions)))
	}
	
	currentMetrics := k.GetSystemMetrics(ctx)
	
	metrics := types.SystemMetrics{
		TotalSupply:           sdk.NewCoin("ssusd", totalSupply),
		TotalCollateral:       sdk.NewCoin("stst", totalCollateral),
		GlobalCollateralRatio: globalRatio,
		TotalPositions:        totalPositions,
		TotalLiquidations:     currentMetrics.TotalLiquidations,
		TotalAgents:           totalAgents,
		AverageCollateralRatio: avgRatio,
	}
	
	k.SetSystemMetrics(ctx, metrics)
}

func (k Keeper) GetNextPositionID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextPositionIDKey)
	if bz == nil {
		return 1
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetNextPositionID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextPositionIDKey, sdk.Uint64ToBigEndian(id))
}

func (k Keeper) GetNextAuctionID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextAuctionIDKey)
	if bz == nil {
		return 1
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetNextAuctionID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextAuctionIDKey, sdk.Uint64ToBigEndian(id))
}

func (k Keeper) CalculateCollateralRatio(ctx sdk.Context, collateral, debt sdk.Coin) sdk.Dec {
	if debt.IsZero() {
		return sdk.NewDec(1000000)
	}
	
	ststPrice := k.GetSTSTPrice(ctx)
	collateralValue := sdk.NewDecFromInt(collateral.Amount).Mul(ststPrice)
	debtValue := sdk.NewDecFromInt(debt.Amount)
	
	return collateralValue.Quo(debtValue)
}

func (k Keeper) IsPositionHealthy(ctx sdk.Context, position types.CollateralPosition) bool {
	params := k.GetParams(ctx)
	return position.CollateralizationRatio.GTE(params.MinCollateralRatio)
}

func (k Keeper) IsPositionLiquidatable(ctx sdk.Context, position types.CollateralPosition) bool {
	params := k.GetParams(ctx)
	return position.CollateralizationRatio.LT(params.LiquidationThreshold)
}