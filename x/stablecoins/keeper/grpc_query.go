package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stateset/core/x/stablecoins/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) Stablecoin(goCtx context.Context, req *types.QueryStablecoinRequest) (*types.QueryStablecoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStablecoin(ctx, req.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, req.Denom)
	}

	return &types.QueryStablecoinResponse{Stablecoin: val}, nil
}

func (k Keeper) Stablecoins(goCtx context.Context, req *types.QueryStablecoinsRequest) (*types.QueryStablecoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stablecoins []types.Stablecoin
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	stablecoinStore := prefix.NewStore(store, types.StablecoinKeyPrefix)

	pageRes, err := query.Paginate(stablecoinStore, req.Pagination, func(key []byte, value []byte) error {
		var stablecoin types.Stablecoin
		if err := k.cdc.Unmarshal(value, &stablecoin); err != nil {
			return err
		}

		stablecoins = append(stablecoins, stablecoin)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryStablecoinsResponse{Stablecoins: stablecoins, Pagination: pageRes}, nil
}

func (k Keeper) StablecoinsByIssuer(goCtx context.Context, req *types.QueryStablecoinsByIssuerRequest) (*types.QueryStablecoinsByIssuerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stablecoins []types.Stablecoin
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	stablecoinStore := prefix.NewStore(store, types.StablecoinKeyPrefix)

	pageRes, err := query.Paginate(stablecoinStore, req.Pagination, func(key []byte, value []byte) error {
		var stablecoin types.Stablecoin
		if err := k.cdc.Unmarshal(value, &stablecoin); err != nil {
			return err
		}

		if stablecoin.Issuer == req.Issuer {
			stablecoins = append(stablecoins, stablecoin)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryStablecoinsByIssuerResponse{Stablecoins: stablecoins, Pagination: pageRes}, nil
}

func (k Keeper) StablecoinSupply(goCtx context.Context, req *types.QueryStablecoinSupplyRequest) (*types.QueryStablecoinSupplyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stablecoin, found := k.GetStablecoin(ctx, req.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, req.Denom)
	}

	// Get actual supply from bank module
	actualSupply := k.bankKeeper.GetSupply(ctx, req.Denom)

	return &types.QueryStablecoinSupplyResponse{
		TotalSupply:       stablecoin.TotalSupply,
		MaxSupply:         stablecoin.MaxSupply,
		CirculatingSupply: actualSupply.Amount,
	}, nil
}

func (k Keeper) PriceData(goCtx context.Context, req *types.QueryPriceDataRequest) (*types.QueryPriceDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	priceDataStore := prefix.NewStore(store, types.PriceDataKeyPrefix)

	var priceDataList []types.PriceData
	iterator := priceDataStore.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var priceData types.PriceData
		k.cdc.MustUnmarshal(iterator.Value(), &priceData)
		if priceData.Denom == req.Denom {
			priceDataList = append(priceDataList, priceData)
		}
	}

	return &types.QueryPriceDataResponse{PriceData: priceDataList}, nil
}

func (k Keeper) ReserveInfo(goCtx context.Context, req *types.QueryReserveInfoRequest) (*types.QueryReserveInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stablecoin, found := k.GetStablecoin(ctx, req.Denom)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, req.Denom)
	}

	if stablecoin.ReserveInfo == nil {
		return &types.QueryReserveInfoResponse{
			ReserveInfo: &types.ReserveInfo{},
		}, nil
	}

	return &types.QueryReserveInfoResponse{
		ReserveInfo: stablecoin.ReserveInfo,
	}, nil
}

func (k Keeper) IsWhitelisted(goCtx context.Context, req *types.QueryIsWhitelistedRequest) (*types.QueryIsWhitelistedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	whitelisted := k.Keeper.IsWhitelisted(ctx, req.Denom, req.Address)

	return &types.QueryIsWhitelistedResponse{
		Whitelisted: whitelisted,
	}, nil
}

func (k Keeper) IsBlacklisted(goCtx context.Context, req *types.QueryIsBlacklistedRequest) (*types.QueryIsBlacklistedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	blacklisted := k.Keeper.IsBlacklisted(ctx, req.Denom, req.Address)

	return &types.QueryIsBlacklistedResponse{
		Blacklisted: blacklisted,
	}, nil
}

func (k Keeper) StablecoinStats(goCtx context.Context, req *types.QueryStablecoinStatsRequest) (*types.QueryStablecoinStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	allStablecoins := k.GetAllStablecoin(ctx)
	
	totalStablecoins := uint64(len(allStablecoins))
	activeStablecoins := uint64(0)
	totalMarketCap := sdk.ZeroDec()
	totalVolume24h := sdk.ZeroDec()
	
	var stablecoinStats []types.StablecoinStat

	for _, stablecoin := range allStablecoins {
		if stablecoin.Active {
			activeStablecoins++
		}

		// Calculate market cap (simplified - just using total supply)
		marketCap := sdk.NewDecFromInt(stablecoin.TotalSupply)
		totalMarketCap = totalMarketCap.Add(marketCap)

		// Create individual stablecoin stat
		stat := types.StablecoinStat{
			Denom:           stablecoin.Denom,
			MarketCap:       marketCap,
			Volume24h:       sdk.ZeroDec(), // In a real implementation, this would come from trading data
			Price:           sdk.OneDec(),  // Assuming stable price of 1.0
			PriceChange24h:  sdk.ZeroDec(),
		}
		stablecoinStats = append(stablecoinStats, stat)
	}

	return &types.QueryStablecoinStatsResponse{
		TotalStablecoins:  totalStablecoins,
		ActiveStablecoins: activeStablecoins,
		TotalMarketCap:    totalMarketCap,
		TotalVolume24h:    totalVolume24h,
		StablecoinStats:   stablecoinStats,
	}, nil
}

func (k Keeper) MintRequests(goCtx context.Context, req *types.QueryMintRequestsRequest) (*types.QueryMintRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mintRequests []types.MintRequest
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	mintRequestStore := prefix.NewStore(store, types.MintRequestKeyPrefix)

	pageRes, err := query.Paginate(mintRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var mintRequest types.MintRequest
		if err := k.cdc.Unmarshal(value, &mintRequest); err != nil {
			return err
		}

		// Apply filters if specified
		include := true
		if req.Denom != "" && mintRequest.Denom != req.Denom {
			include = false
		}
		if req.Status != "" && mintRequest.Status != req.Status {
			include = false
		}

		if include {
			mintRequests = append(mintRequests, mintRequest)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryMintRequestsResponse{MintRequests: mintRequests, Pagination: pageRes}, nil
}

func (k Keeper) BurnRequests(goCtx context.Context, req *types.QueryBurnRequestsRequest) (*types.QueryBurnRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var burnRequests []types.BurnRequest
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	burnRequestStore := prefix.NewStore(store, types.BurnRequestKeyPrefix)

	pageRes, err := query.Paginate(burnRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var burnRequest types.BurnRequest
		if err := k.cdc.Unmarshal(value, &burnRequest); err != nil {
			return err
		}

		// Apply filters if specified
		include := true
		if req.Denom != "" && burnRequest.Denom != req.Denom {
			include = false
		}
		if req.Status != "" && burnRequest.Status != req.Status {
			include = false
		}

		if include {
			burnRequests = append(burnRequests, burnRequest)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBurnRequestsResponse{BurnRequests: burnRequests, Pagination: pageRes}, nil
}