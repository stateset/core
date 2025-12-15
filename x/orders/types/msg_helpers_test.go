package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/orders/types"
)

func TestMsgCreateOrder_ValidateBasic(t *testing.T) {
	validCustomer := sdk.AccAddress("customer____________").String()
	validMerchant := sdk.AccAddress("merchant____________").String()
	validItems := []types.OrderItem{{ProductId: "PROD001", Quantity: 1, UnitPrice: sdk.NewInt64Coin("ssusd", 100)}}

	tests := []struct {
		name      string
		msg       *types.MsgCreateOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgCreateOrder{
				Customer: validCustomer,
				Merchant: validMerchant,
				Items:    validItems,
			},
			expectErr: nil,
		},
		{
			name: "invalid customer address",
			msg: &types.MsgCreateOrder{
				Customer: "invalid",
				Merchant: validMerchant,
				Items:    validItems,
			},
			expectErr: types.ErrInvalidCustomer,
		},
		{
			name: "invalid merchant address",
			msg: &types.MsgCreateOrder{
				Customer: validCustomer,
				Merchant: "invalid",
				Items:    validItems,
			},
			expectErr: types.ErrInvalidMerchant,
		},
		{
			name: "empty items",
			msg: &types.MsgCreateOrder{
				Customer: validCustomer,
				Merchant: validMerchant,
				Items:    []types.OrderItem{},
			},
			expectErr: types.ErrEmptyItems,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgConfirmOrder_ValidateBasic(t *testing.T) {
	validMerchant := sdk.AccAddress("merchant____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgConfirmOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgConfirmOrder{
				Merchant: validMerchant,
				OrderId:  1,
			},
			expectErr: nil,
		},
		{
			name: "invalid merchant address",
			msg: &types.MsgConfirmOrder{
				Merchant: "invalid",
				OrderId:  1,
			},
			expectErr: types.ErrInvalidMerchant,
		},
		{
			name: "zero order id",
			msg: &types.MsgConfirmOrder{
				Merchant: validMerchant,
				OrderId:  0,
			},
			expectErr: types.ErrInvalidOrder,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgPayOrder_ValidateBasic(t *testing.T) {
	validCustomer := sdk.AccAddress("customer____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgPayOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgPayOrder{
				Customer: validCustomer,
				OrderId:  1,
				Amount:   sdk.NewInt64Coin("ssusd", 100),
			},
			expectErr: nil,
		},
		{
			name: "invalid customer address",
			msg: &types.MsgPayOrder{
				Customer: "invalid",
				OrderId:  1,
				Amount:   sdk.NewInt64Coin("ssusd", 100),
			},
			expectErr: types.ErrInvalidCustomer,
		},
		{
			name: "zero order id",
			msg: &types.MsgPayOrder{
				Customer: validCustomer,
				OrderId:  0,
				Amount:   sdk.NewInt64Coin("ssusd", 100),
			},
			expectErr: types.ErrInvalidOrder,
		},
		{
			name: "zero amount",
			msg: &types.MsgPayOrder{
				Customer: validCustomer,
				OrderId:  1,
				Amount:   sdk.NewInt64Coin("ssusd", 0),
			},
			expectErr: types.ErrInvalidAmount,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgShipOrder_ValidateBasic(t *testing.T) {
	validMerchant := sdk.AccAddress("merchant____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgShipOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgShipOrder{
				Merchant:       validMerchant,
				OrderId:        1,
				Carrier:        "UPS",
				TrackingNumber: "1Z999AA10123456784",
			},
			expectErr: nil,
		},
		{
			name: "invalid merchant address",
			msg: &types.MsgShipOrder{
				Merchant:       "invalid",
				OrderId:        1,
				Carrier:        "UPS",
				TrackingNumber: "1Z999AA10123456784",
			},
			expectErr: types.ErrInvalidMerchant,
		},
		{
			name: "zero order id",
			msg: &types.MsgShipOrder{
				Merchant:       validMerchant,
				OrderId:        0,
				Carrier:        "UPS",
				TrackingNumber: "1Z999AA10123456784",
			},
			expectErr: types.ErrInvalidOrder,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgDeliverOrder_ValidateBasic(t *testing.T) {
	validSigner := sdk.AccAddress("signer______________").String()

	tests := []struct {
		name      string
		msg       *types.MsgDeliverOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgDeliverOrder{
				Signer:  validSigner,
				OrderId: 1,
			},
			expectErr: nil,
		},
		{
			name: "invalid signer address",
			msg: &types.MsgDeliverOrder{
				Signer:  "invalid",
				OrderId: 1,
			},
			expectErr: types.ErrUnauthorized,
		},
		{
			name: "zero order id",
			msg: &types.MsgDeliverOrder{
				Signer:  validSigner,
				OrderId: 0,
			},
			expectErr: types.ErrInvalidOrder,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCompleteOrder_ValidateBasic(t *testing.T) {
	validCustomer := sdk.AccAddress("customer____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgCompleteOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgCompleteOrder{
				Customer: validCustomer,
				OrderId:  1,
			},
			expectErr: nil,
		},
		{
			name: "invalid customer address",
			msg: &types.MsgCompleteOrder{
				Customer: "invalid",
				OrderId:  1,
			},
			expectErr: types.ErrInvalidCustomer,
		},
		{
			name: "zero order id",
			msg: &types.MsgCompleteOrder{
				Customer: validCustomer,
				OrderId:  0,
			},
			expectErr: types.ErrInvalidOrder,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCancelOrder_ValidateBasic(t *testing.T) {
	validSigner := sdk.AccAddress("signer______________").String()

	tests := []struct {
		name      string
		msg       *types.MsgCancelOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgCancelOrder{
				Signer:  validSigner,
				OrderId: 1,
				Reason:  "Customer request",
			},
			expectErr: nil,
		},
		{
			name: "invalid signer address",
			msg: &types.MsgCancelOrder{
				Signer:  "invalid",
				OrderId: 1,
				Reason:  "Customer request",
			},
			expectErr: types.ErrUnauthorized,
		},
		{
			name: "zero order id",
			msg: &types.MsgCancelOrder{
				Signer:  validSigner,
				OrderId: 0,
				Reason:  "Customer request",
			},
			expectErr: types.ErrInvalidOrder,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgRefundOrder_ValidateBasic(t *testing.T) {
	validMerchant := sdk.AccAddress("merchant____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgRefundOrder
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgRefundOrder{
				Merchant:     validMerchant,
				OrderId:      1,
				RefundAmount: sdk.NewInt64Coin("ssusd", 100),
				Reason:       "Defective product",
			},
			expectErr: nil,
		},
		{
			name: "invalid merchant address",
			msg: &types.MsgRefundOrder{
				Merchant:     "invalid",
				OrderId:      1,
				RefundAmount: sdk.NewInt64Coin("ssusd", 100),
				Reason:       "Defective product",
			},
			expectErr: types.ErrInvalidMerchant,
		},
		{
			name: "zero order id",
			msg: &types.MsgRefundOrder{
				Merchant:     validMerchant,
				OrderId:      0,
				RefundAmount: sdk.NewInt64Coin("ssusd", 100),
				Reason:       "Defective product",
			},
			expectErr: types.ErrInvalidOrder,
		},
		{
			name: "zero refund amount",
			msg: &types.MsgRefundOrder{
				Merchant:     validMerchant,
				OrderId:      1,
				RefundAmount: sdk.NewInt64Coin("ssusd", 0),
				Reason:       "Defective product",
			},
			expectErr: types.ErrInvalidAmount,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgOpenDispute_ValidateBasic(t *testing.T) {
	validCustomer := sdk.AccAddress("customer____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgOpenDispute
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgOpenDispute{
				Customer:    validCustomer,
				OrderId:     1,
				Reason:      "Item not received",
				Description: "Package never arrived",
			},
			expectErr: nil,
		},
		{
			name: "invalid customer address",
			msg: &types.MsgOpenDispute{
				Customer:    "invalid",
				OrderId:     1,
				Reason:      "Item not received",
				Description: "Package never arrived",
			},
			expectErr: types.ErrInvalidCustomer,
		},
		{
			name: "zero order id",
			msg: &types.MsgOpenDispute{
				Customer:    validCustomer,
				OrderId:     0,
				Reason:      "Item not received",
				Description: "Package never arrived",
			},
			expectErr: types.ErrInvalidOrder,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgResolveDispute_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgResolveDispute
		expectErr error
	}{
		{
			name: "valid message",
			msg: &types.MsgResolveDispute{
				Authority:    validAuthority,
				DisputeId:    1,
				Resolution:   "Refund issued",
				RefundAmount: sdk.NewInt64Coin("ssusd", 100),
				ToCustomer:   true,
			},
			expectErr: nil,
		},
		{
			name: "invalid authority address",
			msg: &types.MsgResolveDispute{
				Authority:    "invalid",
				DisputeId:    1,
				Resolution:   "Refund issued",
				RefundAmount: sdk.NewInt64Coin("ssusd", 100),
				ToCustomer:   true,
			},
			expectErr: types.ErrUnauthorized,
		},
		{
			name: "zero dispute id",
			msg: &types.MsgResolveDispute{
				Authority:    validAuthority,
				DisputeId:    0,
				Resolution:   "Refund issued",
				RefundAmount: sdk.NewInt64Coin("ssusd", 100),
				ToCustomer:   true,
			},
			expectErr: types.ErrDisputeNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCreateOrder_GetSigners(t *testing.T) {
	customer := sdk.AccAddress("customer____________")
	msg := types.MsgCreateOrder{
		Customer: customer.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, customer, signers[0])
}

func TestMsgConfirmOrder_GetSigners(t *testing.T) {
	merchant := sdk.AccAddress("merchant____________")
	msg := types.MsgConfirmOrder{
		Merchant: merchant.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, merchant, signers[0])
}

func TestNewMsgCreateOrder(t *testing.T) {
	customer := sdk.AccAddress("customer____________").String()
	merchant := sdk.AccAddress("merchant____________").String()
	items := []types.OrderItem{{ProductId: "PROD001", Quantity: 1}}
	shipping := types.ShippingInfo{Method: "standard"}
	metadata := "test metadata"

	msg := types.NewMsgCreateOrder(customer, merchant, items, shipping, metadata)

	require.Equal(t, customer, msg.Customer)
	require.Equal(t, merchant, msg.Merchant)
	require.Equal(t, items, msg.Items)
	require.Equal(t, shipping, msg.ShippingInfo)
	require.Equal(t, metadata, msg.Metadata)
}
