package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/settlement/keeper"
	"github.com/stateset/core/x/settlement/types"
)

// TestQueryServer_Settlement tests querying a single settlement
func TestQueryServer_Settlement(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create a settlement
	settlementId, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "REF001", "test")
	require.NoError(t, err)

	// Query the settlement
	req := &types.QuerySettlementRequest{
		Id: settlementId,
	}

	resp, err := queryServer.Settlement(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, settlementId, resp.Settlement.Id)
	require.Equal(t, sender.String(), resp.Settlement.Sender)
	require.Equal(t, recipient.String(), resp.Settlement.Recipient)
	require.Equal(t, "REF001", resp.Settlement.Reference)
}

func TestQueryServer_Settlement_NotFound(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QuerySettlementRequest{
		Id: 9999,
	}

	_, err := queryServer.Settlement(sdk.WrapSDKContext(ctx), req)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrSettlementNotFound)
}

func TestQueryServer_Settlement_ZeroId(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QuerySettlementRequest{
		Id: 0,
	}

	_, err := queryServer.Settlement(sdk.WrapSDKContext(ctx), req)
	require.Error(t, err)
}

// TestQueryServer_Settlements tests querying all settlements
func TestQueryServer_Settlements(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create multiple settlements
	for i := 0; i < 5; i++ {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		_, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "", "")
		require.NoError(t, err)
	}

	// Query all settlements
	req := &types.QuerySettlementsRequest{
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.Settlements(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Settlements, 5)
	require.Equal(t, uint64(5), resp.Total)
}

func TestQueryServer_Settlements_Pagination(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create 10 settlements
	for i := 0; i < 10; i++ {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		_, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "", "")
		require.NoError(t, err)
	}

	// Query first page
	req := &types.QuerySettlementsRequest{
		Limit:  3,
		Offset: 0,
	}

	resp, err := queryServer.Settlements(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Settlements, 3)
	require.Equal(t, uint64(10), resp.Total)

	// Query second page
	req.Offset = 3
	resp, err = queryServer.Settlements(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Settlements, 3)
}

func TestQueryServer_Settlements_DefaultLimit(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QuerySettlementsRequest{
		Limit:  0, // Should use default
		Offset: 0,
	}

	resp, err := queryServer.Settlements(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestQueryServer_Settlements_ExceedsMaxLimit(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QuerySettlementsRequest{
		Limit:  10000, // Exceeds max, should be capped
		Offset: 0,
	}

	resp, err := queryServer.Settlements(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

// TestQueryServer_SettlementsByStatus tests querying settlements by status
func TestQueryServer_SettlementsByStatus(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create instant transfers (completed)
	for i := 0; i < 3; i++ {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		_, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "", "")
		require.NoError(t, err)
	}

	// Create escrows (pending)
	for i := 0; i < 2; i++ {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		_, err := k.CreateEscrow(ctx, sender.String(), recipient.String(), amount, "", "", 86400)
		require.NoError(t, err)
	}

	// Query completed settlements
	req := &types.QuerySettlementsByStatusRequest{
		Status: types.SettlementStatusCompleted,
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.SettlementsByStatus(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Settlements, 3)
	require.Equal(t, uint64(3), resp.Total)

	// Query pending settlements
	req.Status = types.SettlementStatusPending
	resp, err = queryServer.SettlementsByStatus(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Settlements, 2)
	require.Equal(t, uint64(2), resp.Total)
}

func TestQueryServer_SettlementsByStatus_NoMatches(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QuerySettlementsByStatusRequest{
		Status: types.SettlementStatusCancelled,
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.SettlementsByStatus(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Settlements, 0)
	require.Equal(t, uint64(0), resp.Total)
}

// TestQueryServer_Batch tests querying a single batch
func TestQueryServer_Batch(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()
	sender2 := newSettlementAddress()

	senders := []string{sender1.String(), sender2.String()}
	amounts := []sdk.Coin{
		sdk.NewCoin("ssusd", sdkmath.NewInt(100000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(200000)),
	}
	references := []string{"REF1", "REF2"}

	for _, sender := range senders {
		bankKeeper.SetBalance(sender, sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))
	}

	// Create a batch
	batchId, _, err := k.CreateBatch(ctx, merchant.String(), senders, amounts, references)
	require.NoError(t, err)

	// Query the batch
	req := &types.QueryBatchRequest{
		Id: batchId,
	}

	resp, err := queryServer.Batch(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, batchId, resp.Batch.Id)
	require.Equal(t, merchant.String(), resp.Batch.Merchant)
	require.Equal(t, uint64(2), resp.Batch.Count)
}

func TestQueryServer_Batch_NotFound(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QueryBatchRequest{
		Id: 9999,
	}

	_, err := queryServer.Batch(sdk.WrapSDKContext(ctx), req)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrBatchNotFound)
}

// TestQueryServer_Batches tests querying all batches
func TestQueryServer_Batches(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()

	bankKeeper.SetBalance(sender1.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create multiple batches
	for i := 0; i < 3; i++ {
		senders := []string{sender1.String()}
		amounts := []sdk.Coin{sdk.NewCoin("ssusd", sdkmath.NewInt(100000))}
		references := []string{"REF"}

		_, _, err := k.CreateBatch(ctx, merchant.String(), senders, amounts, references)
		require.NoError(t, err)
	}

	// Query all batches
	req := &types.QueryBatchesRequest{
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.Batches(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Batches, 3)
	require.Equal(t, uint64(3), resp.Total)
}

func TestQueryServer_Batches_Empty(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QueryBatchesRequest{
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.Batches(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Batches, 0)
	require.Equal(t, uint64(0), resp.Total)
}

// TestQueryServer_Channel tests querying a single channel
func TestQueryServer_Channel(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Open a channel
	channelId, err := k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
	require.NoError(t, err)

	// Query the channel
	req := &types.QueryChannelRequest{
		Id: channelId,
	}

	resp, err := queryServer.Channel(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, channelId, resp.Channel.Id)
	require.Equal(t, sender.String(), resp.Channel.Sender)
	require.Equal(t, recipient.String(), resp.Channel.Recipient)
	require.True(t, resp.Channel.IsOpen)
}

func TestQueryServer_Channel_NotFound(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QueryChannelRequest{
		Id: 9999,
	}

	_, err := queryServer.Channel(sdk.WrapSDKContext(ctx), req)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrChannelNotFound)
}

// TestQueryServer_Channels tests querying all channels
func TestQueryServer_Channels(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Open multiple channels
	for i := 0; i < 3; i++ {
		_, err := k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
		require.NoError(t, err)
	}

	// Query all channels
	req := &types.QueryChannelsRequest{
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.Channels(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Channels, 3)
	require.Equal(t, uint64(3), resp.Total)
}

func TestQueryServer_Channels_Pagination(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(20000000))))

	// Open 5 channels
	for i := 0; i < 5; i++ {
		_, err := k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
		require.NoError(t, err)
	}

	// Query first page
	req := &types.QueryChannelsRequest{
		Limit:  2,
		Offset: 0,
	}

	resp, err := queryServer.Channels(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Channels, 2)
	require.Equal(t, uint64(5), resp.Total)

	// Query second page
	req.Offset = 2
	resp, err = queryServer.Channels(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Channels, 2)
}

// TestQueryServer_ChannelsByParty tests querying channels by party
func TestQueryServer_ChannelsByParty(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	sender1 := newSettlementAddress()
	sender2 := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender1.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))
	bankKeeper.SetBalance(sender2.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Open channels from sender1
	for i := 0; i < 2; i++ {
		_, err := k.OpenChannel(ctx, sender1.String(), recipient.String(), deposit, 1000)
		require.NoError(t, err)
	}

	// Open channel from sender2
	_, err := k.OpenChannel(ctx, sender2.String(), recipient.String(), deposit, 1000)
	require.NoError(t, err)

	// Query channels for sender1
	req := &types.QueryChannelsByPartyRequest{
		Address: sender1.String(),
		Limit:   10,
		Offset:  0,
	}

	resp, err := queryServer.ChannelsByParty(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Channels, 2)
	require.Equal(t, uint64(2), resp.Total)

	// Query channels for recipient (should include all 3)
	req.Address = recipient.String()
	resp, err = queryServer.ChannelsByParty(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Channels, 3)
	require.Equal(t, uint64(3), resp.Total)
}

func TestQueryServer_ChannelsByParty_NoChannels(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	addr := newSettlementAddress()

	req := &types.QueryChannelsByPartyRequest{
		Address: addr.String(),
		Limit:   10,
		Offset:  0,
	}

	resp, err := queryServer.ChannelsByParty(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Channels, 0)
	require.Equal(t, uint64(0), resp.Total)
}

// TestQueryServer_Merchant tests querying a single merchant
func TestQueryServer_Merchant(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	merchantAddr := newSettlementAddress()
	config := types.MerchantConfig{
		Address:      merchantAddr.String(),
		Name:         "Test Merchant",
		FeeRateBps:   25,
		BatchEnabled: true,
	}

	err := k.RegisterMerchant(ctx, config)
	require.NoError(t, err)

	// Query the merchant
	req := &types.QueryMerchantRequest{
		Address: merchantAddr.String(),
	}

	resp, err := queryServer.Merchant(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, merchantAddr.String(), resp.Merchant.Address)
	require.Equal(t, "Test Merchant", resp.Merchant.Name)
	require.Equal(t, uint32(25), resp.Merchant.FeeRateBps)
}

func TestQueryServer_Merchant_NotFound(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	addr := newSettlementAddress()

	req := &types.QueryMerchantRequest{
		Address: addr.String(),
	}

	_, err := queryServer.Merchant(sdk.WrapSDKContext(ctx), req)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrMerchantNotFound)
}

// TestQueryServer_Merchants tests querying all merchants
func TestQueryServer_Merchants(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	// Register multiple merchants
	for i := 0; i < 3; i++ {
		merchantAddr := newSettlementAddress()
		config := types.MerchantConfig{
			Address:    merchantAddr.String(),
			Name:       "Test Merchant",
			FeeRateBps: 25,
		}
		err := k.RegisterMerchant(ctx, config)
		require.NoError(t, err)
	}

	// Query all merchants
	req := &types.QueryMerchantsRequest{
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.Merchants(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Merchants, 3)
	require.Equal(t, uint64(3), resp.Total)
}

func TestQueryServer_Merchants_Empty(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QueryMerchantsRequest{
		Limit:  10,
		Offset: 0,
	}

	resp, err := queryServer.Merchants(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Merchants, 0)
	require.Equal(t, uint64(0), resp.Total)
}

func TestQueryServer_Merchants_Pagination(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	// Register 5 merchants
	for i := 0; i < 5; i++ {
		merchantAddr := newSettlementAddress()
		config := types.MerchantConfig{
			Address:    merchantAddr.String(),
			Name:       "Test Merchant",
			FeeRateBps: 25,
		}
		err := k.RegisterMerchant(ctx, config)
		require.NoError(t, err)
	}

	// Query first page
	req := &types.QueryMerchantsRequest{
		Limit:  2,
		Offset: 0,
	}

	resp, err := queryServer.Merchants(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Merchants, 2)
	require.Equal(t, uint64(5), resp.Total)

	// Query third page
	req.Offset = 4
	resp, err = queryServer.Merchants(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Len(t, resp.Merchants, 1)
}

// TestQueryServer_Params tests querying module params
func TestQueryServer_Params(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	req := &types.QueryParamsRequest{}

	resp, err := queryServer.Params(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Params)
}

func TestQueryServer_Params_AfterUpdate(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	queryServer := keeper.NewQueryServerImpl(k)

	// Update params
	params := k.GetParams(ctx)
	params.DefaultFeeRateBps = 100
	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	// Query params
	req := &types.QueryParamsRequest{}

	resp, err := queryServer.Params(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint32(100), resp.Params.DefaultFeeRateBps)
}
