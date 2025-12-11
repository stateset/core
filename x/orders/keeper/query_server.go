package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stateset/core/x/orders/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl returns a QueryServer implementation backed by keeper logic.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{keeper: k}
}

func (q queryServer) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.keeper.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

func (q queryServer) Order(goCtx context.Context, req *types.QueryOrderRequest) (*types.QueryOrderResponse, error) {
	if req == nil || req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	order, found := q.keeper.GetOrder(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	return &types.QueryOrderResponse{Order: order}, nil
}

func (q queryServer) Orders(goCtx context.Context, req *types.QueryOrdersRequest) (*types.QueryOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var all []types.Order
	q.keeper.IterateOrders(ctx, func(order types.Order) bool {
		if req.Customer != "" && order.Customer != req.Customer {
			return false
		}
		if req.Merchant != "" && order.Merchant != req.Merchant {
			return false
		}
		if req.Status != "" && string(order.Status) != req.Status {
			return false
		}
		all = append(all, order)
		return false
	})

	total := uint64(len(all))
	offset := req.Offset
	if offset > total {
		offset = total
	}
	limit := req.Limit
	if limit == 0 || offset+limit > total {
		limit = total - offset
	}
	orders := all[offset : offset+limit]

	return &types.QueryOrdersResponse{Orders: orders, Total: total}, nil
}
