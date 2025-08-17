package keeper_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/purchaseorder/types"
)

func TestTimedoutPurchaseorderMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTimedoutPurchaseorder(ctx, &types.MsgCreateTimedoutPurchaseorder{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTimedoutPurchaseorderMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTimedoutPurchaseorder
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTimedoutPurchaseorder{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutPurchaseorder{Creator: "B"},
			err:     errorsmod.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutPurchaseorder{Creator: creator, Id: 10},
			err:     errorsmod.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateTimedoutPurchaseorder(ctx, &types.MsgCreateTimedoutPurchaseorder{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateTimedoutPurchaseorder(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTimedoutPurchaseorderMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTimedoutPurchaseorder
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTimedoutPurchaseorder{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTimedoutPurchaseorder{Creator: "B"},
			err:     errorsmod.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteTimedoutPurchaseorder{Creator: creator, Id: 10},
			err:     errorsmod.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateTimedoutPurchaseorder(ctx, &types.MsgCreateTimedoutPurchaseorder{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTimedoutPurchaseorder(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
