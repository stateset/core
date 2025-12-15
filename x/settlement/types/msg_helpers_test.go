package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/settlement/types"
)

func TestValidateWebhookURL(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		expectErr bool
	}{
		{
			name:      "empty URL is allowed",
			url:       "",
			expectErr: false,
		},
		{
			name:      "valid HTTPS URL",
			url:       "https://example.com/webhook",
			expectErr: false,
		},
		{
			name:      "HTTP URL rejected",
			url:       "http://example.com/webhook",
			expectErr: true,
		},
		{
			name:      "localhost rejected",
			url:       "https://localhost/webhook",
			expectErr: true,
		},
		{
			name:      "127.0.0.1 rejected",
			url:       "https://127.0.0.1/webhook",
			expectErr: true,
		},
		{
			name:      "0.0.0.0 rejected",
			url:       "https://0.0.0.0/webhook",
			expectErr: true,
		},
		{
			name:      "private network 10.x rejected",
			url:       "https://10.0.0.1/webhook",
			expectErr: true,
		},
		{
			name:      "private network 172.16.x rejected",
			url:       "https://172.16.0.1/webhook",
			expectErr: true,
		},
		{
			name:      "private network 192.168.x rejected",
			url:       "https://192.168.1.1/webhook",
			expectErr: true,
		},
		{
			name:      "IPv6 loopback rejected",
			url:       "https://[::1]/webhook",
			expectErr: true,
		},
		{
			name:      "invalid URL",
			url:       "not a url",
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateWebhookURL(tc.url)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgInstantTransfer_ValidateBasic(t *testing.T) {
	validSender := sdk.AccAddress("sender______________").String()
	validRecipient := sdk.AccAddress("recipient___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgInstantTransfer
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgInstantTransfer{
				Sender:    validSender,
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
			},
			expectErr: false,
		},
		{
			name: "invalid sender address",
			msg: &types.MsgInstantTransfer{
				Sender:    "invalid",
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
			},
			expectErr: true,
		},
		{
			name: "invalid recipient address",
			msg: &types.MsgInstantTransfer{
				Sender:    validSender,
				Recipient: "invalid",
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
			},
			expectErr: true,
		},
		{
			name: "sender equals recipient",
			msg: &types.MsgInstantTransfer{
				Sender:    validSender,
				Recipient: validSender,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
			},
			expectErr: true,
		},
		{
			name: "zero amount",
			msg: &types.MsgInstantTransfer{
				Sender:    validSender,
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 0),
			},
			expectErr: true,
		},
		{
			name: "wrong denom",
			msg: &types.MsgInstantTransfer{
				Sender:    validSender,
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin("uatom", 100),
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

func TestMsgCreateEscrow_ValidateBasic(t *testing.T) {
	validSender := sdk.AccAddress("sender______________").String()
	validRecipient := sdk.AccAddress("recipient___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgCreateEscrow
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgCreateEscrow{
				Sender:    validSender,
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresIn: time.Hour,
			},
			expectErr: false,
		},
		{
			name: "invalid sender",
			msg: &types.MsgCreateEscrow{
				Sender:    "invalid",
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresIn: time.Hour,
			},
			expectErr: true,
		},
		{
			name: "sender equals recipient",
			msg: &types.MsgCreateEscrow{
				Sender:    validSender,
				Recipient: validSender,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresIn: time.Hour,
			},
			expectErr: true,
		},
		{
			name: "wrong denom",
			msg: &types.MsgCreateEscrow{
				Sender:    validSender,
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin("uatom", 100),
				ExpiresIn: time.Hour,
			},
			expectErr: true,
		},
		{
			name: "negative expiration",
			msg: &types.MsgCreateEscrow{
				Sender:    validSender,
				Recipient: validRecipient,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresIn: -time.Hour,
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

func TestMsgReleaseEscrow_ValidateBasic(t *testing.T) {
	validSender := sdk.AccAddress("sender______________").String()

	tests := []struct {
		name      string
		msg       *types.MsgReleaseEscrow
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgReleaseEscrow{
				Sender:       validSender,
				SettlementId: 1,
			},
			expectErr: false,
		},
		{
			name: "invalid sender",
			msg: &types.MsgReleaseEscrow{
				Sender:       "invalid",
				SettlementId: 1,
			},
			expectErr: true,
		},
		{
			name: "zero settlement id",
			msg: &types.MsgReleaseEscrow{
				Sender:       validSender,
				SettlementId: 0,
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

func TestMsgRefundEscrow_ValidateBasic(t *testing.T) {
	validRecipient := sdk.AccAddress("recipient___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgRefundEscrow
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgRefundEscrow{
				Recipient:    validRecipient,
				SettlementId: 1,
				Reason:       "Refund reason",
			},
			expectErr: false,
		},
		{
			name: "invalid recipient",
			msg: &types.MsgRefundEscrow{
				Recipient:    "invalid",
				SettlementId: 1,
				Reason:       "Refund reason",
			},
			expectErr: true,
		},
		{
			name: "zero settlement id",
			msg: &types.MsgRefundEscrow{
				Recipient:    validRecipient,
				SettlementId: 0,
				Reason:       "Refund reason",
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

func TestMsgCreateBatch_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()
	validMerchant := sdk.AccAddress("merchant____________").String()
	validSender := sdk.AccAddress("sender______________").String()

	tests := []struct {
		name      string
		msg       *types.MsgCreateBatch
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgCreateBatch{
				Authority:  validAuthority,
				Merchant:   validMerchant,
				Senders:    []string{validSender},
				Amounts:    []sdk.Coin{sdk.NewInt64Coin(types.StablecoinDenom, 100)},
				References: []string{"ref1"},
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgCreateBatch{
				Authority:  "invalid",
				Merchant:   validMerchant,
				Senders:    []string{validSender},
				Amounts:    []sdk.Coin{sdk.NewInt64Coin(types.StablecoinDenom, 100)},
				References: []string{"ref1"},
			},
			expectErr: true,
		},
		{
			name: "empty senders",
			msg: &types.MsgCreateBatch{
				Authority:  validAuthority,
				Merchant:   validMerchant,
				Senders:    []string{},
				Amounts:    []sdk.Coin{},
				References: []string{},
			},
			expectErr: true,
		},
		{
			name: "mismatched lengths",
			msg: &types.MsgCreateBatch{
				Authority:  validAuthority,
				Merchant:   validMerchant,
				Senders:    []string{validSender, validSender},
				Amounts:    []sdk.Coin{sdk.NewInt64Coin(types.StablecoinDenom, 100)},
				References: []string{"ref1"},
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

func TestMsgOpenChannel_ValidateBasic(t *testing.T) {
	validSender := sdk.AccAddress("sender______________").String()
	validRecipient := sdk.AccAddress("recipient___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgOpenChannel
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgOpenChannel{
				Sender:          validSender,
				Recipient:       validRecipient,
				Deposit:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresInBlocks: 100,
			},
			expectErr: false,
		},
		{
			name: "invalid sender",
			msg: &types.MsgOpenChannel{
				Sender:          "invalid",
				Recipient:       validRecipient,
				Deposit:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresInBlocks: 100,
			},
			expectErr: true,
		},
		{
			name: "sender equals recipient",
			msg: &types.MsgOpenChannel{
				Sender:          validSender,
				Recipient:       validSender,
				Deposit:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresInBlocks: 100,
			},
			expectErr: true,
		},
		{
			name: "wrong denom",
			msg: &types.MsgOpenChannel{
				Sender:          validSender,
				Recipient:       validRecipient,
				Deposit:         sdk.NewInt64Coin("uatom", 100),
				ExpiresInBlocks: 100,
			},
			expectErr: true,
		},
		{
			name: "zero expiration blocks",
			msg: &types.MsgOpenChannel{
				Sender:          validSender,
				Recipient:       validRecipient,
				Deposit:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresInBlocks: 0,
			},
			expectErr: true,
		},
		{
			name: "negative expiration blocks",
			msg: &types.MsgOpenChannel{
				Sender:          validSender,
				Recipient:       validRecipient,
				Deposit:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				ExpiresInBlocks: -1,
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

func TestMsgClaimChannel_ValidateBasic(t *testing.T) {
	validRecipient := sdk.AccAddress("recipient___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgClaimChannel
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgClaimChannel{
				Recipient: validRecipient,
				ChannelId: 1,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				Nonce:     1,
				Signature: "a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2",
			},
			expectErr: false,
		},
		{
			name: "invalid recipient",
			msg: &types.MsgClaimChannel{
				Recipient: "invalid",
				ChannelId: 1,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				Nonce:     1,
				Signature: "a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2",
			},
			expectErr: true,
		},
		{
			name: "zero channel id",
			msg: &types.MsgClaimChannel{
				Recipient: validRecipient,
				ChannelId: 0,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				Nonce:     1,
				Signature: "a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2",
			},
			expectErr: true,
		},
		{
			name: "zero nonce",
			msg: &types.MsgClaimChannel{
				Recipient: validRecipient,
				ChannelId: 1,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				Nonce:     0,
				Signature: "a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2c3d4e5f6g7h8i9j0a1b2",
			},
			expectErr: true,
		},
		{
			name: "empty signature",
			msg: &types.MsgClaimChannel{
				Recipient: validRecipient,
				ChannelId: 1,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				Nonce:     1,
				Signature: "",
			},
			expectErr: true,
		},
		{
			name: "short signature",
			msg: &types.MsgClaimChannel{
				Recipient: validRecipient,
				ChannelId: 1,
				Amount:    sdk.NewInt64Coin(types.StablecoinDenom, 100),
				Nonce:     1,
				Signature: "short",
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

func TestMsgRegisterMerchant_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()
	validMerchant := sdk.AccAddress("merchant____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgRegisterMerchant
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgRegisterMerchant{
				Authority:  validAuthority,
				Merchant:   validMerchant,
				Name:       "Test Merchant",
				FeeRateBps: 100,
				WebhookUrl: "https://example.com/webhook",
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgRegisterMerchant{
				Authority:  "invalid",
				Merchant:   validMerchant,
				Name:       "Test Merchant",
				FeeRateBps: 100,
			},
			expectErr: true,
		},
		{
			name: "empty name",
			msg: &types.MsgRegisterMerchant{
				Authority:  validAuthority,
				Merchant:   validMerchant,
				Name:       "",
				FeeRateBps: 100,
			},
			expectErr: true,
		},
		{
			name: "fee rate too high",
			msg: &types.MsgRegisterMerchant{
				Authority:  validAuthority,
				Merchant:   validMerchant,
				Name:       "Test Merchant",
				FeeRateBps: 10001,
			},
			expectErr: true,
		},
		{
			name: "invalid webhook URL",
			msg: &types.MsgRegisterMerchant{
				Authority:  validAuthority,
				Merchant:   validMerchant,
				Name:       "Test Merchant",
				FeeRateBps: 100,
				WebhookUrl: "http://localhost/webhook",
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

func TestMsgInstantCheckout_ValidateBasic(t *testing.T) {
	validCustomer := sdk.AccAddress("customer____________").String()
	validMerchant := sdk.AccAddress("merchant____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgInstantCheckout
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgInstantCheckout{
				Customer:       validCustomer,
				Merchant:       validMerchant,
				Amount:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				OrderReference: "ORDER-123",
			},
			expectErr: false,
		},
		{
			name: "invalid customer",
			msg: &types.MsgInstantCheckout{
				Customer:       "invalid",
				Merchant:       validMerchant,
				Amount:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				OrderReference: "ORDER-123",
			},
			expectErr: true,
		},
		{
			name: "customer equals merchant",
			msg: &types.MsgInstantCheckout{
				Customer:       validCustomer,
				Merchant:       validCustomer,
				Amount:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				OrderReference: "ORDER-123",
			},
			expectErr: true,
		},
		{
			name: "wrong denom",
			msg: &types.MsgInstantCheckout{
				Customer:       validCustomer,
				Merchant:       validMerchant,
				Amount:         sdk.NewInt64Coin("uatom", 100),
				OrderReference: "ORDER-123",
			},
			expectErr: true,
		},
		{
			name: "empty order reference",
			msg: &types.MsgInstantCheckout{
				Customer:       validCustomer,
				Merchant:       validMerchant,
				Amount:         sdk.NewInt64Coin(types.StablecoinDenom, 100),
				OrderReference: "",
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

func TestMsgPartialRefund_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgPartialRefund
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgPartialRefund{
				Authority:    validAuthority,
				SettlementId: 1,
				RefundAmount: sdk.NewInt64Coin(types.StablecoinDenom, 50),
				Reason:       "Partial refund reason",
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgPartialRefund{
				Authority:    "invalid",
				SettlementId: 1,
				RefundAmount: sdk.NewInt64Coin(types.StablecoinDenom, 50),
				Reason:       "Partial refund reason",
			},
			expectErr: true,
		},
		{
			name: "zero settlement id",
			msg: &types.MsgPartialRefund{
				Authority:    validAuthority,
				SettlementId: 0,
				RefundAmount: sdk.NewInt64Coin(types.StablecoinDenom, 50),
				Reason:       "Partial refund reason",
			},
			expectErr: true,
		},
		{
			name: "zero refund amount",
			msg: &types.MsgPartialRefund{
				Authority:    validAuthority,
				SettlementId: 1,
				RefundAmount: sdk.NewInt64Coin(types.StablecoinDenom, 0),
				Reason:       "Partial refund reason",
			},
			expectErr: true,
		},
		{
			name: "empty reason",
			msg: &types.MsgPartialRefund{
				Authority:    validAuthority,
				SettlementId: 1,
				RefundAmount: sdk.NewInt64Coin(types.StablecoinDenom, 50),
				Reason:       "",
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

func TestMsgInstantTransfer_GetSigners(t *testing.T) {
	sender := sdk.AccAddress("sender______________")
	msg := types.MsgInstantTransfer{
		Sender: sender.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, sender, signers[0])
}

func TestNewMsgInstantTransfer(t *testing.T) {
	sender := sdk.AccAddress("sender______________").String()
	recipient := sdk.AccAddress("recipient___________").String()
	amount := sdk.NewInt64Coin(types.StablecoinDenom, 100)
	reference := "ref123"
	metadata := "test transfer"

	msg := types.NewMsgInstantTransfer(sender, recipient, amount, reference, metadata)

	require.Equal(t, sender, msg.Sender)
	require.Equal(t, recipient, msg.Recipient)
	require.Equal(t, amount, msg.Amount)
	require.Equal(t, reference, msg.Reference)
	require.Equal(t, metadata, msg.Metadata)
}
