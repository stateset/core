package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgTerminateAgreement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgTerminateAgreement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgTerminateAgreement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgTerminateAgreement{
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
