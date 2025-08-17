package keeper_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/invoice/types"
)

func TestTimedoutInvoiceMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTimedoutInvoice(ctx, &types.MsgCreateTimedoutInvoice{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTimedoutInvoiceMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTimedoutInvoice
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTimedoutInvoice{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutInvoice{Creator: "B"},
			err:     errorsmod.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutInvoice{Creator: creator, Id: 10},
			err:     errorsmod.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateTimedoutInvoice(ctx, &types.MsgCreateTimedoutInvoice{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateTimedoutInvoice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTimedoutInvoiceMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTimedoutInvoice
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTimedoutInvoice{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTimedoutInvoice{Creator: "B"},
			err:     errorsmod.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteTimedoutInvoice{Creator: creator, Id: 10},
			err:     errorsmod.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateTimedoutInvoice(ctx, &types.MsgCreateTimedoutInvoice{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTimedoutInvoice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
