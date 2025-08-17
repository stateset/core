package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateSentPurchaseorder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSentPurchaseorder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSentPurchaseorder{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSentPurchaseorder{
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

func TestMsgUpdateSentPurchaseorder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSentPurchaseorder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSentPurchaseorder{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSentPurchaseorder{
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

func TestMsgDeleteSentPurchaseorder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSentPurchaseorder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSentPurchaseorder{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSentPurchaseorder{
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
