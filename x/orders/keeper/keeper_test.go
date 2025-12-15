package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	orderskeeper "github.com/stateset/core/x/orders/keeper"
	orderstypes "github.com/stateset/core/x/orders/types"
)

func setupKeeper(t *testing.T) (orderskeeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(orderstypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := orderskeeper.NewKeeper(
		cdc,
		storeKey,
		"stateset1authority",
		nil, // bankKeeper - nil for basic tests
		nil, // complianceKeeper
		nil, // settlementKeeper
		nil, // accountKeeper
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{
		Height:  1,
		ChainID: "stateset-test",
		Time:    time.Now(),
	}, false, log.NewNopLogger())

	return k, ctx
}

func TestNewKeeper(t *testing.T) {
	k, _ := setupKeeper(t)
	require.NotNil(t, k)
	require.Equal(t, "stateset1authority", k.GetAuthority())
}

func TestGetSetParams(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Get default params
	params := k.GetParams(ctx)
	require.NotNil(t, params)

	// Default values should be sensible
	require.True(t, params.DefaultOrderExpiration > 0, "default order expiration should be positive")
	require.True(t, params.DefaultEscrowExpiration > 0, "default escrow expiration should be positive")
	require.True(t, params.DisputeWindow > 0, "dispute window should be positive")

	// Set custom params
	customParams := orderstypes.Params{
		DefaultOrderExpiration:  172800, // 2 days
		DefaultEscrowExpiration: 1209600, // 14 days
		DisputeWindow:           2419200, // 28 days
		MinOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(500_000)),
		MaxOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(500_000_000_000)),
		DefaultFeeRateBps:       150, // 1.5%
		StablecoinDenom:         "ussUSD",
		AutoCompleteAfterDelivery: true,
		AutoCompleteWindow:      345600, // 4 days
	}
	err := k.SetParams(ctx, customParams)
	require.NoError(t, err)

	// Verify params were stored
	retrieved := k.GetParams(ctx)
	require.Equal(t, int64(172800), retrieved.DefaultOrderExpiration)
	require.Equal(t, int64(1209600), retrieved.DefaultEscrowExpiration)
	require.Equal(t, int64(2419200), retrieved.DisputeWindow)
	require.Equal(t, uint32(150), retrieved.DefaultFeeRateBps)
	require.True(t, retrieved.AutoCompleteAfterDelivery)
}

func TestGetOrder_NotFound(t *testing.T) {
	k, ctx := setupKeeper(t)

	_, found := k.GetOrder(ctx, 999)
	require.False(t, found)
}

func TestIterateOrders_Empty(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Iterate over empty store
	count := 0
	k.IterateOrders(ctx, func(order orderstypes.Order) bool {
		count++
		return false
	})
	require.Equal(t, 0, count)
}

func TestGetDispute_NotFound(t *testing.T) {
	k, ctx := setupKeeper(t)

	_, found := k.GetDispute(ctx, 999)
	require.False(t, found)
}

func TestParamsValidation(t *testing.T) {
	k, ctx := setupKeeper(t)

	tests := []struct {
		name      string
		params    orderstypes.Params
		expectErr bool
	}{
		{
			name: "valid params",
			params: orderstypes.Params{
				DefaultOrderExpiration:  86400,
				DefaultEscrowExpiration: 604800,
				DisputeWindow:           1209600,
				MinOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(100_000)),
				MaxOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(1_000_000_000_000)),
				DefaultFeeRateBps:       100,
				StablecoinDenom:         "ussUSD",
				AutoCompleteAfterDelivery: true,
				AutoCompleteWindow:      259200,
			},
			expectErr: false,
		},
		{
			name: "zero order expiration",
			params: orderstypes.Params{
				DefaultOrderExpiration:  0,
				DefaultEscrowExpiration: 604800,
				DisputeWindow:           1209600,
				MinOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(100_000)),
				MaxOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(1_000_000_000_000)),
				DefaultFeeRateBps:       100,
				StablecoinDenom:         "ussUSD",
			},
			expectErr: true,
		},
		{
			name: "fee rate exceeds max",
			params: orderstypes.Params{
				DefaultOrderExpiration:  86400,
				DefaultEscrowExpiration: 604800,
				DisputeWindow:           1209600,
				MinOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(100_000)),
				MaxOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(1_000_000_000_000)),
				DefaultFeeRateBps:       10001, // > 100%
				StablecoinDenom:         "ussUSD",
			},
			expectErr: true,
		},
		{
			name: "empty stablecoin denom is allowed (uses default)",
			params: orderstypes.Params{
				DefaultOrderExpiration:  86400,
				DefaultEscrowExpiration: 604800,
				DisputeWindow:           1209600,
				MinOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(100_000)),
				MaxOrderAmount:          sdk.NewCoin("ussUSD", sdkmath.NewInt(1_000_000_000_000)),
				DefaultFeeRateBps:       100,
				StablecoinDenom:         "",
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.SetParams(ctx, tt.params)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestOrderStatusConstants(t *testing.T) {
	// Verify all status constants are defined
	require.NotEmpty(t, orderstypes.OrderStatusPending)
	require.NotEmpty(t, orderstypes.OrderStatusConfirmed)
	require.NotEmpty(t, orderstypes.OrderStatusPaid)
	require.NotEmpty(t, orderstypes.OrderStatusShipped)
	require.NotEmpty(t, orderstypes.OrderStatusDelivered)
	require.NotEmpty(t, orderstypes.OrderStatusCompleted)
	require.NotEmpty(t, orderstypes.OrderStatusCancelled)
	require.NotEmpty(t, orderstypes.OrderStatusRefunded)
	require.NotEmpty(t, orderstypes.OrderStatusDisputed)

	// Verify they are distinct
	statuses := []string{
		orderstypes.OrderStatusPending,
		orderstypes.OrderStatusConfirmed,
		orderstypes.OrderStatusPaid,
		orderstypes.OrderStatusShipped,
		orderstypes.OrderStatusDelivered,
		orderstypes.OrderStatusCompleted,
		orderstypes.OrderStatusCancelled,
		orderstypes.OrderStatusRefunded,
		orderstypes.OrderStatusDisputed,
	}
	seen := make(map[string]bool)
	for _, s := range statuses {
		require.False(t, seen[s], "duplicate status: %s", s)
		seen[s] = true
	}
}

func TestDisputeStatusConstants(t *testing.T) {
	require.NotEmpty(t, orderstypes.DisputeStatusOpen)
	require.NotEmpty(t, orderstypes.DisputeStatusResolved)
}

func TestGetAuthority(t *testing.T) {
	k, _ := setupKeeper(t)
	require.Equal(t, "stateset1authority", k.GetAuthority())
}

func TestIterateDisputes_Empty(t *testing.T) {
	k, ctx := setupKeeper(t)

	count := 0
	k.IterateDisputes(ctx, func(dispute orderstypes.Dispute) bool {
		count++
		return false
	})
	require.Equal(t, 0, count)
}

func TestDefaultParams(t *testing.T) {
	params := orderstypes.DefaultParams()

	// Verify default params are valid
	require.NoError(t, params.Validate())

	// Verify reasonable defaults
	require.True(t, params.DefaultOrderExpiration > 0)
	require.True(t, params.DefaultEscrowExpiration > 0)
	require.True(t, params.DisputeWindow > 0)
	require.True(t, params.MinOrderAmount.IsValid())
	require.True(t, params.MaxOrderAmount.IsValid())
	require.True(t, params.MinOrderAmount.Amount.LT(params.MaxOrderAmount.Amount))
	require.True(t, params.DefaultFeeRateBps <= 10000) // Max 100%
}
