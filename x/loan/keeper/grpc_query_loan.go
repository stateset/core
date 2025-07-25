package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/loan/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LoanAll(c context.Context, req *types.QueryAllLoanRequest) (*types.QueryAllLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var loans []types.Loan
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	loanStore := prefix.NewStore(store, types.KeyPrefix(types.LoanKey))

	pageRes, err := query.Paginate(loanStore, req.Pagination, func(key []byte, value []byte) error {
		var loan types.Loan
		if err := k.cdc.Unmarshal(value, &loan); err != nil {
			return err
		}

		loans = append(loans, loan)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLoanResponse{Loan: loans, Pagination: pageRes}, nil
}

func (k Keeper) Loan(c context.Context, req *types.QueryGetLoanRequest) (*types.QueryGetLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	loan, found := k.GetLoan(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLoanResponse{Loan: loan}, nil
}
