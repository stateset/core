package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/compliance/keeper"
	"github.com/stateset/core/x/compliance/types"
)

func TestMsgUpsertProfile(t *testing.T) {
	k, ctx := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	addr := newAddress()
	msg := types.NewMsgUpsertProfile(k.GetAuthority(), types.Profile{
		Address:  addr.String(),
		KYCLevel: "enhanced",
		Risk:     types.RiskMedium,
	})

	_, err := msgServer.UpsertProfile(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)

	stored, found := k.GetProfile(ctx, addr)
	require.True(t, found)
	require.Equal(t, msg.Authority, stored.UpdatedBy)
	require.Equal(t, msg.Profile.KYCLevel, stored.KYCLevel)
}

func TestMsgUpsertProfile_Unauthorized(t *testing.T) {
	k, ctx := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := types.NewMsgUpsertProfile(newAddress().String(), types.Profile{
		Address:  newAddress().String(),
		KYCLevel: "standard",
		Risk:     types.RiskLow,
	})

	_, err := msgServer.UpsertProfile(sdk.WrapSDKContext(ctx), msg)
	require.ErrorIs(t, err, types.ErrUnauthorized)
}

func TestMsgSetSanction(t *testing.T) {
	k, ctx := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	target := newAddress()
	authority := k.GetAuthority()
	upsert := types.NewMsgUpsertProfile(authority, types.Profile{
		Address:  target.String(),
		KYCLevel: "standard",
		Risk:     types.RiskLow,
	})
	_, err := msgServer.UpsertProfile(sdk.WrapSDKContext(ctx), upsert)
	require.NoError(t, err)

	setSanction := types.NewMsgSetSanction(authority, target.String(), true, "fraud detection")
	_, err = msgServer.SetSanction(sdk.WrapSDKContext(ctx), setSanction)
	require.NoError(t, err)

	stored, found := k.GetProfile(ctx, target)
	require.True(t, found)
	require.True(t, stored.Sanction)
	require.Equal(t, authority, stored.UpdatedBy)
	require.Equal(t, "fraud detection", stored.Metadata)

	setSanction = types.NewMsgSetSanction(authority, target.String(), false, "")
	_, err = msgServer.SetSanction(sdk.WrapSDKContext(ctx), setSanction)
	require.NoError(t, err)

	stored, found = k.GetProfile(ctx, target)
	require.True(t, found)
	require.False(t, stored.Sanction)
}

func TestMsgSetSanctionRequiresProfile(t *testing.T) {
	k, ctx := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	target := newAddress().String()
	msg := types.NewMsgSetSanction(k.GetAuthority(), target, true, "test")

	_, err := msgServer.SetSanction(sdk.WrapSDKContext(ctx), msg)
	require.ErrorIs(t, err, types.ErrProfileNotFound)
}

func TestMsgSetSanction_Unauthorized(t *testing.T) {
	k, ctx := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	target := newAddress()
	upsert := types.NewMsgUpsertProfile(k.GetAuthority(), types.Profile{
		Address:  target.String(),
		KYCLevel: "standard",
		Risk:     types.RiskLow,
	})
	_, err := msgServer.UpsertProfile(sdk.WrapSDKContext(ctx), upsert)
	require.NoError(t, err)

	msg := types.NewMsgSetSanction(newAddress().String(), target.String(), true, "test")
	_, err = msgServer.SetSanction(sdk.WrapSDKContext(ctx), msg)
	require.ErrorIs(t, err, types.ErrUnauthorized)

	// Ensure the sanction flag was not mutated.
	stored, found := k.GetProfile(ctx, target)
	require.True(t, found)
	require.False(t, stored.Sanction)
}
