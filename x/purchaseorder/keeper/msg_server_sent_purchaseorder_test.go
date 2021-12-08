package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/purchaseorder/types"
)

func TestSentPurchaseorderMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateSentPurchaseorder(ctx, &types.MsgCreateSentPurchaseorder{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestSentPurchaseorderMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateSentPurchaseorder
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateSentPurchaseorder{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateSentPurchaseorder{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateSentPurchaseorder{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateSentPurchaseorder(ctx, &types.MsgCreateSentPurchaseorder{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateSentPurchaseorder(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSentPurchaseorderMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteSentPurchaseorder
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteSentPurchaseorder{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteSentPurchaseorder{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteSentPurchaseorder{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateSentPurchaseorder(ctx, &types.MsgCreateSentPurchaseorder{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteSentPurchaseorder(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
