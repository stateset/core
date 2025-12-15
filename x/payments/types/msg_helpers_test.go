package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/payments/types"
)

func TestMsgCreatePayment_ValidateBasic(t *testing.T) {
	validPayer := sdk.AccAddress("payer_______________").String()
	validPayee := sdk.AccAddress("payee_______________").String()

	tests := []struct {
		name      string
		msg       *types.MsgCreatePayment
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgCreatePayment{
				Payer:  validPayer,
				Payee:  validPayee,
				Amount: sdk.NewInt64Coin("ssusd", 100),
			},
			expectErr: false,
		},
		{
			name: "invalid payer address",
			msg: &types.MsgCreatePayment{
				Payer:  "invalid",
				Payee:  validPayee,
				Amount: sdk.NewInt64Coin("ssusd", 100),
			},
			expectErr: true,
		},
		{
			name: "invalid payee address",
			msg: &types.MsgCreatePayment{
				Payer:  validPayer,
				Payee:  "invalid",
				Amount: sdk.NewInt64Coin("ssusd", 100),
			},
			expectErr: true,
		},
		{
			name: "zero amount",
			msg: &types.MsgCreatePayment{
				Payer:  validPayer,
				Payee:  validPayee,
				Amount: sdk.NewInt64Coin("ssusd", 0),
			},
			expectErr: true,
		},
		{
			name: "negative amount",
			msg: &types.MsgCreatePayment{
				Payer:  validPayer,
				Payee:  validPayee,
				Amount: sdk.Coin{Denom: "ssusd", Amount: sdkmath.NewInt(-100)},
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgSettlePayment_ValidateBasic(t *testing.T) {
	validPayee := sdk.AccAddress("payee_______________").String()

	tests := []struct {
		name      string
		msg       *types.MsgSettlePayment
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgSettlePayment{
				Payee:     validPayee,
				PaymentId: 1,
			},
			expectErr: false,
		},
		{
			name: "invalid payee address",
			msg: &types.MsgSettlePayment{
				Payee:     "invalid",
				PaymentId: 1,
			},
			expectErr: true,
		},
		{
			name: "zero payment id",
			msg: &types.MsgSettlePayment{
				Payee:     validPayee,
				PaymentId: 0,
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCancelPayment_ValidateBasic(t *testing.T) {
	validPayer := sdk.AccAddress("payer_______________").String()

	tests := []struct {
		name      string
		msg       *types.MsgCancelPayment
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgCancelPayment{
				Payer:     validPayer,
				PaymentId: 1,
				Reason:    "Changed my mind",
			},
			expectErr: false,
		},
		{
			name: "invalid payer address",
			msg: &types.MsgCancelPayment{
				Payer:     "invalid",
				PaymentId: 1,
				Reason:    "Changed my mind",
			},
			expectErr: true,
		},
		{
			name: "zero payment id",
			msg: &types.MsgCancelPayment{
				Payer:     validPayer,
				PaymentId: 0,
				Reason:    "Changed my mind",
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCreatePayment_GetSigners(t *testing.T) {
	payer := sdk.AccAddress("payer_______________")
	msg := types.MsgCreatePayment{
		Payer: payer.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, payer, signers[0])
}

func TestMsgSettlePayment_GetSigners(t *testing.T) {
	payee := sdk.AccAddress("payee_______________")
	msg := types.MsgSettlePayment{
		Payee: payee.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, payee, signers[0])
}

func TestMsgCancelPayment_GetSigners(t *testing.T) {
	payer := sdk.AccAddress("payer_______________")
	msg := types.MsgCancelPayment{
		Payer: payer.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, payer, signers[0])
}

func TestNewMsgCreatePayment(t *testing.T) {
	payer := sdk.AccAddress("payer_______________").String()
	payee := sdk.AccAddress("payee_______________").String()
	amount := sdk.NewInt64Coin("ssusd", 100)
	metadata := "test payment"

	msg := types.NewMsgCreatePayment(payer, payee, amount, metadata)

	require.Equal(t, payer, msg.Payer)
	require.Equal(t, payee, msg.Payee)
	require.Equal(t, amount, msg.Amount)
	require.Equal(t, metadata, msg.Metadata)
}

func TestNewMsgSettlePayment(t *testing.T) {
	payee := sdk.AccAddress("payee_______________").String()
	paymentId := uint64(123)

	msg := types.NewMsgSettlePayment(payee, paymentId)

	require.Equal(t, payee, msg.Payee)
	require.Equal(t, paymentId, msg.PaymentId)
}

func TestNewMsgCancelPayment(t *testing.T) {
	payer := sdk.AccAddress("payer_______________").String()
	paymentId := uint64(123)
	reason := "test reason"

	msg := types.NewMsgCancelPayment(payer, paymentId, reason)

	require.Equal(t, payer, msg.Payer)
	require.Equal(t, paymentId, msg.PaymentId)
	require.Equal(t, reason, msg.Reason)
}
