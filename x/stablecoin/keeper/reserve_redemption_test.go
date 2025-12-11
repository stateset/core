package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/stablecoin/types"
)

func TestReserveRedemption_ImmediateUpdatesStatsAndIDs(t *testing.T) {
	k, ctx, bank, oracle, _ := setupKeeper(t)

	depositor := newAddress()
	amount := sdk.NewInt64Coin("usdy", 1_000)
	bank.SetBalance(depositor, sdk.NewCoins(amount))
	oracle.SetPrice("usdy", sdkmath.LegacyMustNewDecFromStr("1.0"))

	_, minted, err := k.DepositReserve(ctx, depositor, amount)
	require.NoError(t, err)
	require.True(t, minted.Equal(sdkmath.NewInt(994)))

	redemptionID, err := k.RequestRedemption(ctx, depositor, minted, "usdy")
	require.NoError(t, err)
	require.Equal(t, uint64(1), redemptionID)

	req, found := k.GetRedemptionRequest(ctx, redemptionID)
	require.True(t, found)
	require.Equal(t, types.RedeemStatusExecuted, req.Status)
	require.Equal(t, "usdy", req.OutputAmount.Denom)
	require.True(t, req.OutputAmount.Amount.Equal(sdkmath.NewInt(993)))

	stats := k.GetDailyMintStats(ctx)
	require.True(t, stats.TotalRedeemed.Equal(minted))

	// Second immediate redemption should increment ID.
	depositor2 := newAddress()
	bank.SetBalance(depositor2, sdk.NewCoins(amount))
	_, minted2, err := k.DepositReserve(ctx, depositor2, amount)
	require.NoError(t, err)

	redemptionID2, err := k.RequestRedemption(ctx, depositor2, minted2, "usdy")
	require.NoError(t, err)
	require.Equal(t, uint64(2), redemptionID2)
}

func TestCancelRedemption_RequiresAuthority(t *testing.T) {
	k, ctx, bank, oracle, _ := setupKeeper(t)

	depositor := newAddress()
	amount := sdk.NewInt64Coin("usdy", 1_000)
	bank.SetBalance(depositor, sdk.NewCoins(amount))
	oracle.SetPrice("usdy", sdkmath.LegacyMustNewDecFromStr("1.0"))

	_, minted, err := k.DepositReserve(ctx, depositor, amount)
	require.NoError(t, err)

	// Set a non-zero delay so redemption stays pending.
	params := k.GetReserveParams(ctx)
	params.RedemptionDelay = time.Hour
	require.NoError(t, k.SetReserveParams(ctx, params))

	redemptionID, err := k.RequestRedemption(ctx, depositor, minted, "usdy")
	require.NoError(t, err)

	req, found := k.GetRedemptionRequest(ctx, redemptionID)
	require.True(t, found)
	require.Equal(t, types.RedeemStatusPending, req.Status)

	badAuthority := newAddress().String()
	require.NotEqual(t, k.GetAuthority(), badAuthority)

	err = k.CancelRedemption(ctx, badAuthority, redemptionID)
	require.Error(t, err)

	err = k.CancelRedemption(ctx, k.GetAuthority(), redemptionID)
	require.NoError(t, err)

	req, found = k.GetRedemptionRequest(ctx, redemptionID)
	require.True(t, found)
	require.Equal(t, types.RedeemStatusCancelled, req.Status)

	// Depositor should have their ssUSD refunded.
	require.True(t, bank.Balance(depositor).AmountOf(types.StablecoinDenom).Equal(minted))
}
