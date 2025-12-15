package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	orderskeeper "github.com/stateset/core/x/orders/keeper"
	orderstypes "github.com/stateset/core/x/orders/types"
)

func TestQueryServer_Params(t *testing.T) {
	k, ctx := setupKeeper(t)
	qs := orderskeeper.NewQueryServerImpl(k)

	// Test nil request
	_, err := qs.Params(ctx, nil)
	require.Error(t, err)

	// Test valid request
	resp, err := qs.Params(ctx, &orderstypes.QueryParamsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Params)
	require.True(t, resp.Params.DefaultOrderExpiration > 0)
}

func TestQueryServer_Order(t *testing.T) {
	k, ctx := setupKeeper(t)
	qs := orderskeeper.NewQueryServerImpl(k)

	// Test nil request
	_, err := qs.Order(ctx, nil)
	require.Error(t, err)

	// Test invalid request (id=0)
	_, err = qs.Order(ctx, &orderstypes.QueryOrderRequest{Id: 0})
	require.Error(t, err)

	// Test order not found
	_, err = qs.Order(ctx, &orderstypes.QueryOrderRequest{Id: 999})
	require.Error(t, err)
}

func TestQueryServer_Orders(t *testing.T) {
	k, ctx := setupKeeper(t)
	qs := orderskeeper.NewQueryServerImpl(k)

	// Test nil request
	_, err := qs.Orders(ctx, nil)
	require.Error(t, err)

	// Test empty list
	resp, err := qs.Orders(ctx, &orderstypes.QueryOrdersRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint64(0), resp.Total)
	require.Empty(t, resp.Orders)

	// Test filtering - should work even with empty list
	resp, err = qs.Orders(ctx, &orderstypes.QueryOrdersRequest{
		Customer: "stateset1customer",
	})
	require.NoError(t, err)
	require.Equal(t, uint64(0), resp.Total)

	resp, err = qs.Orders(ctx, &orderstypes.QueryOrdersRequest{
		Merchant: "stateset1merchant",
	})
	require.NoError(t, err)
	require.Equal(t, uint64(0), resp.Total)

	resp, err = qs.Orders(ctx, &orderstypes.QueryOrdersRequest{
		Status: orderstypes.OrderStatusPending,
	})
	require.NoError(t, err)
	require.Equal(t, uint64(0), resp.Total)
}

func TestQueryServer_Orders_Pagination(t *testing.T) {
	k, ctx := setupKeeper(t)
	qs := orderskeeper.NewQueryServerImpl(k)

	// Test pagination with empty list
	resp, err := qs.Orders(ctx, &orderstypes.QueryOrdersRequest{
		Offset: 0,
		Limit:  10,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint64(0), resp.Total)

	// Test with offset > total
	resp, err = qs.Orders(ctx, &orderstypes.QueryOrdersRequest{
		Offset: 100,
		Limit:  10,
	})
	require.NoError(t, err)
	require.Equal(t, uint64(0), resp.Total)
	require.Empty(t, resp.Orders)
}

func TestNewQueryServerImpl(t *testing.T) {
	k, _ := setupKeeper(t)
	qs := orderskeeper.NewQueryServerImpl(k)
	require.NotNil(t, qs)
}
