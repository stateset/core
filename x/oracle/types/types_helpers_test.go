package types_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/oracle/types"
)

func TestMsgUpdatePrice_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgUpdatePrice
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgUpdatePrice{
				Authority: validAuthority,
				Denom:     "uatom",
				Price:     sdkmath.LegacyMustNewDecFromStr("10.50"),
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgUpdatePrice{
				Authority: "invalid",
				Denom:     "uatom",
				Price:     sdkmath.LegacyMustNewDecFromStr("10.50"),
			},
			expectErr: true,
		},
		{
			name: "empty denom",
			msg: &types.MsgUpdatePrice{
				Authority: validAuthority,
				Denom:     "",
				Price:     sdkmath.LegacyMustNewDecFromStr("10.50"),
			},
			expectErr: true,
		},
		{
			name: "zero price",
			msg: &types.MsgUpdatePrice{
				Authority: validAuthority,
				Denom:     "uatom",
				Price:     sdkmath.LegacyZeroDec(),
			},
			expectErr: true,
		},
		{
			name: "negative price",
			msg: &types.MsgUpdatePrice{
				Authority: validAuthority,
				Denom:     "uatom",
				Price:     sdkmath.LegacyMustNewDecFromStr("-10.50"),
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

func TestPrice_ValidateBasic(t *testing.T) {
	validUpdater := sdk.AccAddress("updater_____________").String()

	tests := []struct {
		name      string
		price     types.Price
		expectErr bool
	}{
		{
			name: "valid price",
			price: types.Price{
				Denom:       "uatom",
				Amount:      sdkmath.LegacyMustNewDecFromStr("10.50"),
				LastUpdater: validUpdater,
			},
			expectErr: false,
		},
		{
			name: "valid price without updater",
			price: types.Price{
				Denom:  "uatom",
				Amount: sdkmath.LegacyMustNewDecFromStr("10.50"),
			},
			expectErr: false,
		},
		{
			name: "empty denom",
			price: types.Price{
				Denom:  "",
				Amount: sdkmath.LegacyMustNewDecFromStr("10.50"),
			},
			expectErr: true,
		},
		{
			name: "zero price",
			price: types.Price{
				Denom:  "uatom",
				Amount: sdkmath.LegacyZeroDec(),
			},
			expectErr: true,
		},
		{
			name: "negative price",
			price: types.Price{
				Denom:  "uatom",
				Amount: sdkmath.LegacyMustNewDecFromStr("-10.50"),
			},
			expectErr: true,
		},
		{
			name: "invalid updater address",
			price: types.Price{
				Denom:       "uatom",
				Amount:      sdkmath.LegacyMustNewDecFromStr("10.50"),
				LastUpdater: "invalid",
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.price.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPrice_IsStale(t *testing.T) {
	now := time.Now()
	threshold := time.Hour

	tests := []struct {
		name      string
		price     types.Price
		now       time.Time
		threshold time.Duration
		expected  bool
	}{
		{
			name: "zero update time is stale",
			price: types.Price{
				Denom:  "uatom",
				Amount: sdkmath.LegacyOneDec(),
			},
			now:       now,
			threshold: threshold,
			expected:  true,
		},
		{
			name: "recent price is not stale",
			price: types.Price{
				Denom:     "uatom",
				Amount:    sdkmath.LegacyOneDec(),
				UpdatedAt: now.Add(-30 * time.Minute),
			},
			now:       now,
			threshold: threshold,
			expected:  false,
		},
		{
			name: "old price is stale",
			price: types.Price{
				Denom:     "uatom",
				Amount:    sdkmath.LegacyOneDec(),
				UpdatedAt: now.Add(-2 * time.Hour),
			},
			now:       now,
			threshold: threshold,
			expected:  true,
		},
		{
			name: "exactly at threshold is not stale",
			price: types.Price{
				Denom:     "uatom",
				Amount:    sdkmath.LegacyOneDec(),
				UpdatedAt: now.Add(-threshold),
			},
			now:       now,
			threshold: threshold,
			expected:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.price.IsStale(tc.now, tc.threshold)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestDefaultOracleConfig(t *testing.T) {
	config := types.DefaultOracleConfig("uatom")

	require.Equal(t, "uatom", config.Denom)
	require.Equal(t, uint64(500), config.MaxDeviationBps)
	require.Equal(t, int64(3600), config.StalenessThresholdSeconds)
	require.Equal(t, int64(60), config.MinUpdateIntervalSeconds)
	require.Equal(t, uint32(1), config.RequiredConfirmations)
	require.True(t, config.Enabled)
}

func TestDefaultOracleParams(t *testing.T) {
	params := types.DefaultOracleParams()

	require.Equal(t, uint64(500), params.DefaultMaxDeviationBps)
	require.Equal(t, int64(3600), params.DefaultStalenessThreshold)
	require.Equal(t, uint64(1000), params.SlashFractionBps)
	require.Equal(t, uint32(10), params.MaxProviders)
	require.Equal(t, uint32(100), params.PriceHistorySize)
}

func TestCalculateDeviation(t *testing.T) {
	tests := []struct {
		name     string
		oldPrice sdkmath.LegacyDec
		newPrice sdkmath.LegacyDec
		expected sdkmath.LegacyDec
	}{
		{
			name:     "zero old price returns zero",
			oldPrice: sdkmath.LegacyZeroDec(),
			newPrice: sdkmath.LegacyMustNewDecFromStr("10"),
			expected: sdkmath.LegacyZeroDec(),
		},
		{
			name:     "same price returns zero",
			oldPrice: sdkmath.LegacyMustNewDecFromStr("100"),
			newPrice: sdkmath.LegacyMustNewDecFromStr("100"),
			expected: sdkmath.LegacyZeroDec(),
		},
		{
			name:     "10% increase",
			oldPrice: sdkmath.LegacyMustNewDecFromStr("100"),
			newPrice: sdkmath.LegacyMustNewDecFromStr("110"),
			expected: sdkmath.LegacyMustNewDecFromStr("1000"), // 1000 bps = 10%
		},
		{
			name:     "10% decrease",
			oldPrice: sdkmath.LegacyMustNewDecFromStr("100"),
			newPrice: sdkmath.LegacyMustNewDecFromStr("90"),
			expected: sdkmath.LegacyMustNewDecFromStr("1000"), // 1000 bps = 10%
		},
		{
			name:     "50% increase",
			oldPrice: sdkmath.LegacyMustNewDecFromStr("100"),
			newPrice: sdkmath.LegacyMustNewDecFromStr("150"),
			expected: sdkmath.LegacyMustNewDecFromStr("5000"), // 5000 bps = 50%
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := types.CalculateDeviation(tc.oldPrice, tc.newPrice)
			require.True(t, result.Equal(tc.expected), "expected %s, got %s", tc.expected, result)
		})
	}
}

func TestMsgUpdatePrice_GetSigners(t *testing.T) {
	authority := sdk.AccAddress("authority___________")
	msg := types.MsgUpdatePrice{
		Authority: authority.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, authority, signers[0])
}

func TestNewMsgUpdatePrice(t *testing.T) {
	authority := sdk.AccAddress("authority___________").String()
	denom := "uatom"
	price := sdkmath.LegacyMustNewDecFromStr("10.50")

	msg := types.NewMsgUpdatePrice(authority, denom, price)

	require.Equal(t, authority, msg.Authority)
	require.Equal(t, denom, msg.Denom)
	require.True(t, msg.Price.Equal(price))
}
