package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTimedoutAgreement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTimedoutAgreement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTimedoutAgreement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTimedoutAgreement{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateTimedoutAgreement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTimedoutAgreement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTimedoutAgreement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTimedoutAgreement{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteTimedoutAgreement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTimedoutAgreement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTimedoutAgreement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTimedoutAgreement{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
