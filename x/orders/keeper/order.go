package keeper

import (
	"strconv"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/orders/types"
)

// SetOrder stores an order in the KVStore
func (k Keeper) SetOrder(ctx sdk.Context, order types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKey))
	b := k.cdc.MustMarshal(&order)
	store.Set([]byte(order.Id), b)
}

// GetOrder retrieves an order from the KVStore
func (k Keeper) GetOrder(ctx sdk.Context, id string) (val types.Order, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKey))

	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrder removes an order from the KVStore
func (k Keeper) RemoveOrder(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKey))
	store.Delete([]byte(id))
}

// GetAllOrders returns all orders
func (k Keeper) GetAllOrders(ctx sdk.Context) (list []types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Order
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOrdersByCustomer returns all orders for a specific customer
func (k Keeper) GetOrdersByCustomer(ctx sdk.Context, customer string) (list []types.Order) {
	allOrders := k.GetAllOrders(ctx)
	for _, order := range allOrders {
		if order.Customer == customer {
			list = append(list, order)
		}
	}
	return
}

// GetOrdersByMerchant returns all orders for a specific merchant
func (k Keeper) GetOrdersByMerchant(ctx sdk.Context, merchant string) (list []types.Order) {
	allOrders := k.GetAllOrders(ctx)
	for _, order := range allOrders {
		if order.Merchant == merchant {
			list = append(list, order)
		}
	}
	return
}

// GetOrdersByStatus returns all orders with a specific status
func (k Keeper) GetOrdersByStatus(ctx sdk.Context, status string) (list []types.Order) {
	allOrders := k.GetAllOrders(ctx)
	for _, order := range allOrders {
		if order.Status == status {
			list = append(list, order)
		}
	}
	return
}

// GetNextOrderID returns the next order ID
func (k Keeper) GetNextOrderID(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderCountKey))
	
	bz := store.Get([]byte("count"))
	if bz == nil {
		store.Set([]byte("count"), []byte("1"))
		return 1
	}

	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		panic(err)
	}

	nextCount := count + 1
	store.Set([]byte("count"), []byte(strconv.FormatUint(nextCount, 10)))
	
	return nextCount
}

// GetOrderCount returns the total number of orders
func (k Keeper) GetOrderCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderCountKey))
	
	bz := store.Get([]byte("count"))
	if bz == nil {
		return 0
	}

	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		return 0
	}

	return count
}