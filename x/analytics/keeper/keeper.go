package keeper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/analytics/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// RecordMetric records a performance metric
func (k Keeper) RecordMetric(ctx sdk.Context, metric types.PerformanceMetric) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetricKeyPrefix))
	
	key := []byte(fmt.Sprintf("%s_%d", metric.Name, ctx.BlockTime().Unix()))
	value, err := json.Marshal(metric)
	if err != nil {
		return err
	}
	
	store.Set(key, value)
	return nil
}

// GetMetrics retrieves metrics for a given time range
func (k Keeper) GetMetrics(ctx sdk.Context, startTime, endTime time.Time) ([]types.PerformanceMetric, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetricKeyPrefix))
	
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	
	var metrics []types.PerformanceMetric
	for ; iterator.Valid(); iterator.Next() {
		var metric types.PerformanceMetric
		err := json.Unmarshal(iterator.Value(), &metric)
		if err != nil {
			continue
		}
		
		if metric.Timestamp.After(startTime) && metric.Timestamp.Before(endTime) {
			metrics = append(metrics, metric)
		}
	}
	
	return metrics, nil
}