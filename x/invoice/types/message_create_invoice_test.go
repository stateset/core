package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateInvoice_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateInvoice
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateInvoice{
				Creator: "invalid_address",
			},
			err: errorsmod.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateInvoice{
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
