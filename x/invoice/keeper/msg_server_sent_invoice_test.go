package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/invoice/types"
)

func TestSentInvoiceMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateSentInvoice(ctx, &types.MsgCreateSentInvoice{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestSentInvoiceMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateSentInvoice
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateSentInvoice{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateSentInvoice{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateSentInvoice{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateSentInvoice(ctx, &types.MsgCreateSentInvoice{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateSentInvoice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSentInvoiceMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteSentInvoice
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteSentInvoice{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteSentInvoice{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteSentInvoice{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateSentInvoice(ctx, &types.MsgCreateSentInvoice{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteSentInvoice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
