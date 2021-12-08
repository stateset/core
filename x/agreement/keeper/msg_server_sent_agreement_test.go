package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/agreement/types"
)

func TestSentAgreementMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateSentAgreement(ctx, &types.MsgCreateSentAgreement{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestSentAgreementMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateSentAgreement
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateSentAgreement{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateSentAgreement{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateSentAgreement{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateSentAgreement(ctx, &types.MsgCreateSentAgreement{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateSentAgreement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSentAgreementMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteSentAgreement
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteSentAgreement{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteSentAgreement{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteSentAgreement{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateSentAgreement(ctx, &types.MsgCreateSentAgreement{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteSentAgreement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
