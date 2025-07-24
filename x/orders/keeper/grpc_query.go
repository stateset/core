package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/orders/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Order(goCtx context.Context, req *types.QueryOrderRequest) (*types.QueryOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetOrder(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryOrderResponse{Order: &val}, nil
}

func (k Keeper) Orders(goCtx context.Context, req *types.QueryOrdersRequest) (*types.QueryOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orders []types.Order
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKey))

	pageRes, err := query.Paginate(orderStore, req.Pagination, func(key []byte, value []byte) error {
		var order types.Order
		if err := k.cdc.Unmarshal(value, &order); err != nil {
			return err
		}

		orders = append(orders, order)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryOrdersResponse{Orders: orders, Pagination: pageRes}, nil
}

func (k Keeper) OrdersByCustomer(goCtx context.Context, req *types.QueryOrdersByCustomerRequest) (*types.QueryOrdersByCustomerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	orders := k.GetOrdersByCustomer(ctx, req.Customer)

	return &types.QueryOrdersByCustomerResponse{Orders: orders}, nil
}

func (k Keeper) OrdersByMerchant(goCtx context.Context, req *types.QueryOrdersByMerchantRequest) (*types.QueryOrdersByMerchantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	orders := k.GetOrdersByMerchant(ctx, req.Merchant)

	return &types.QueryOrdersByMerchantResponse{Orders: orders}, nil
}

func (k Keeper) OrdersByStatus(goCtx context.Context, req *types.QueryOrdersByStatusRequest) (*types.QueryOrdersByStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	orders := k.GetOrdersByStatus(ctx, req.Status)

	return &types.QueryOrdersByStatusResponse{Orders: orders}, nil
}

func (k Keeper) OrderStats(goCtx context.Context, req *types.QueryOrderStatsRequest) (*types.QueryOrderStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	allOrders := k.GetAllOrders(ctx)

	// Calculate stats
	var totalOrders, pendingOrders, confirmedOrders, shippedOrders, deliveredOrders, cancelledOrders, refundedOrders uint64

	for _, order := range allOrders {
		totalOrders++
		switch order.Status {
		case "pending":
			pendingOrders++
		case "confirmed":
			confirmedOrders++
		case "shipped":
			shippedOrders++
		case "delivered":
			deliveredOrders++
		case "cancelled":
			cancelledOrders++
		case "refunded":
			refundedOrders++
		}
	}

	return &types.QueryOrderStatsResponse{
		TotalOrders:     totalOrders,
		PendingOrders:   pendingOrders,
		ConfirmedOrders: confirmedOrders,
		ShippedOrders:   shippedOrders,
		DeliveredOrders: deliveredOrders,
		CancelledOrders: cancelledOrders,
		RefundedOrders:  refundedOrders,
		TotalRevenue:    k.calculateTotalRevenue(ctx),
		AverageOrderValue: k.calculateAverageOrderValue(ctx, totalOrders),
	}, nil
}

// calculateTotalRevenue calculates the total revenue from all successful orders
func (k Keeper) calculateTotalRevenue(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.OrderPrefix)
	defer iterator.Close()

	totalRevenue := sdk.ZeroInt()
	
	for ; iterator.Valid(); iterator.Next() {
		var order types.Order
		k.cdc.MustUnmarshal(iterator.Value(), &order)
		
		// Only count revenue from delivered orders
		if order.OrderStatus == "delivered" {
			if orderAmount, ok := sdk.NewIntFromString(order.Amount); ok {
				totalRevenue = totalRevenue.Add(orderAmount)
			}
		}
	}
	
	return totalRevenue.String()
}

// calculateAverageOrderValue calculates the average value per order
func (k Keeper) calculateAverageOrderValue(ctx sdk.Context, totalOrders int64) string {
	if totalOrders == 0 {
		return "0"
	}
	
	totalRevenueStr := k.calculateTotalRevenue(ctx)
	totalRevenue, ok := sdk.NewIntFromString(totalRevenueStr)
	if !ok {
		return "0"
	}
	
	averageValue := totalRevenue.Quo(sdk.NewInt(totalOrders))
	return averageValue.String()
}