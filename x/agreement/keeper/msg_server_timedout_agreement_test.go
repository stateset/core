package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/agreement/types"
)

func TestTimedoutAgreementMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTimedoutAgreement(ctx, &types.MsgCreateTimedoutAgreement{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTimedoutAgreementMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTimedoutAgreement
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTimedoutAgreement{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutAgreement{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutAgreement{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateTimedoutAgreement(ctx, &types.MsgCreateTimedoutAgreement{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateTimedoutAgreement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTimedoutAgreementMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTimedoutAgreement
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTimedoutAgreement{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTimedoutAgreement{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteTimedoutAgreement{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateTimedoutAgreement(ctx, &types.MsgCreateTimedoutAgreement{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTimedoutAgreement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
