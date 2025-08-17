package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTimedoutPurchaseorder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTimedoutPurchaseorder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTimedoutPurchaseorder{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTimedoutPurchaseorder{
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

func TestMsgUpdateTimedoutPurchaseorder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTimedoutPurchaseorder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTimedoutPurchaseorder{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTimedoutPurchaseorder{
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

func TestMsgDeleteTimedoutPurchaseorder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTimedoutPurchaseorder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTimedoutPurchaseorder{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTimedoutPurchaseorder{
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
