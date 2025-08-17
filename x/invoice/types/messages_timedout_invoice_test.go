package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTimedoutInvoice_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTimedoutInvoice
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTimedoutInvoice{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTimedoutInvoice{
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

func TestMsgUpdateTimedoutInvoice_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTimedoutInvoice
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTimedoutInvoice{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTimedoutInvoice{
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

func TestMsgDeleteTimedoutInvoice_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTimedoutInvoice
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTimedoutInvoice{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTimedoutInvoice{
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
