package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/settlement/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the settlement QueryServer interface
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

// Settlement returns a settlement by ID
func (q queryServer) Settlement(goCtx context.Context, req *types.QuerySettlementRequest) (*types.QuerySettlementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	settlement, found := q.Keeper.GetSettlement(ctx, req.Id)
	if !found {
		return nil, types.ErrSettlementNotFound
	}

	return &types.QuerySettlementResponse{
		Settlement: settlement,
	}, nil
}

// Settlements returns all settlements with pagination
func (q queryServer) Settlements(goCtx context.Context, req *types.QuerySettlementsRequest) (*types.QuerySettlementsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.MaxQueryLimit)
	if maxLimit == 0 {
		maxLimit = 100 // fallback default
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var settlements []types.Settlement
	var total uint64

	q.Keeper.IterateSettlements(ctx, func(s types.Settlement) bool {
		if total >= offset && uint64(len(settlements)) < limit {
			settlements = append(settlements, s)
		}
		total++
		return false
	})

	return &types.QuerySettlementsResponse{
		Settlements: settlements,
		Total:       total,
	}, nil
}

// SettlementsByStatus returns settlements filtered by status
func (q queryServer) SettlementsByStatus(goCtx context.Context, req *types.QuerySettlementsByStatusRequest) (*types.QuerySettlementsByStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.MaxQueryLimit)
	if maxLimit == 0 {
		maxLimit = 100
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var settlements []types.Settlement
	var total uint64
	var matched uint64

	q.Keeper.IterateSettlements(ctx, func(s types.Settlement) bool {
		if s.Status == req.Status {
			if matched >= offset && uint64(len(settlements)) < limit {
				settlements = append(settlements, s)
			}
			matched++
		}
		total++
		return false
	})

	return &types.QuerySettlementsByStatusResponse{
		Settlements: settlements,
		Total:       matched,
	}, nil
}

// Batch returns a batch by ID
func (q queryServer) Batch(goCtx context.Context, req *types.QueryBatchRequest) (*types.QueryBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	batch, found := q.Keeper.GetBatch(ctx, req.Id)
	if !found {
		return nil, types.ErrBatchNotFound
	}

	return &types.QueryBatchResponse{
		Batch: batch,
	}, nil
}

// Batches returns all batches with pagination
func (q queryServer) Batches(goCtx context.Context, req *types.QueryBatchesRequest) (*types.QueryBatchesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.MaxQueryLimit)
	if maxLimit == 0 {
		maxLimit = 100
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var batches []types.BatchSettlement
	var total uint64

	q.Keeper.IterateBatches(ctx, func(b types.BatchSettlement) bool {
		if total >= offset && uint64(len(batches)) < limit {
			batches = append(batches, b)
		}
		total++
		return false
	})

	return &types.QueryBatchesResponse{
		Batches: batches,
		Total:   total,
	}, nil
}

// Channel returns a payment channel by ID
func (q queryServer) Channel(goCtx context.Context, req *types.QueryChannelRequest) (*types.QueryChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	channel, found := q.Keeper.GetChannel(ctx, req.Id)
	if !found {
		return nil, types.ErrChannelNotFound
	}

	return &types.QueryChannelResponse{
		Channel: channel,
	}, nil
}

// Channels returns all payment channels with pagination
func (q queryServer) Channels(goCtx context.Context, req *types.QueryChannelsRequest) (*types.QueryChannelsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.MaxQueryLimit)
	if maxLimit == 0 {
		maxLimit = 100
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var channels []types.PaymentChannel
	var total uint64

	q.Keeper.IterateChannels(ctx, func(c types.PaymentChannel) bool {
		if total >= offset && uint64(len(channels)) < limit {
			channels = append(channels, c)
		}
		total++
		return false
	})

	return &types.QueryChannelsResponse{
		Channels: channels,
		Total:    total,
	}, nil
}

// ChannelsByParty returns channels for a specific sender or recipient
func (q queryServer) ChannelsByParty(goCtx context.Context, req *types.QueryChannelsByPartyRequest) (*types.QueryChannelsByPartyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.MaxQueryLimit)
	if maxLimit == 0 {
		maxLimit = 100
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var channels []types.PaymentChannel
	var matched uint64

	q.Keeper.IterateChannels(ctx, func(c types.PaymentChannel) bool {
		if c.Sender == req.Address || c.Recipient == req.Address {
			if matched >= offset && uint64(len(channels)) < limit {
				channels = append(channels, c)
			}
			matched++
		}
		return false
	})

	return &types.QueryChannelsByPartyResponse{
		Channels: channels,
		Total:    matched,
	}, nil
}

// Merchant returns a merchant configuration by address
func (q queryServer) Merchant(goCtx context.Context, req *types.QueryMerchantRequest) (*types.QueryMerchantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	merchant, found := q.Keeper.GetMerchant(ctx, req.Address)
	if !found {
		return nil, types.ErrMerchantNotFound
	}

	return &types.QueryMerchantResponse{
		Merchant: merchant,
	}, nil
}

// Merchants returns all merchants with pagination
func (q queryServer) Merchants(goCtx context.Context, req *types.QueryMerchantsRequest) (*types.QueryMerchantsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)
	maxLimit := uint64(params.MaxQueryLimit)
	if maxLimit == 0 {
		maxLimit = 100
	}

	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var merchants []types.MerchantConfig
	var total uint64

	q.Keeper.IterateMerchants(ctx, func(m types.MerchantConfig) bool {
		if total >= offset && uint64(len(merchants)) < limit {
			merchants = append(merchants, m)
		}
		total++
		return false
	})

	return &types.QueryMerchantsResponse{
		Merchants: merchants,
		Total:     total,
	}, nil
}

// Params returns the module parameters
func (q queryServer) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
