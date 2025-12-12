package keeper_test

import (
	"sync"
	"testing"
	"time"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	testutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/compliance/keeper"
	compliancetypes "github.com/stateset/core/x/compliance/types"
)

var basicConfigOnce sync.Once

func setupBasicConfig() {
	basicConfigOnce.Do(func() {
		cfg := sdk.GetConfig()
		cfg.SetBech32PrefixForAccount("stateset", "statesetpub")
		cfg.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
		cfg.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")
		cfg.Seal()
	})
}

func setupBasicKeeper(t *testing.T) (keeper.Keeper, sdk.Context) {
	t.Helper()

	setupBasicConfig()

	storeKey := storetypes.NewKVStoreKey(compliancetypes.StoreKey)

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("compliance-transient"))
	ctx = ctx.WithBlockTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))

	authority := newBasicAddress().String()
	return keeper.NewKeeper(cdc, storeKey, authority), ctx
}

func newBasicAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

// newAddress is an alias for compatibility with other test files in the package
func newAddress() sdk.AccAddress {
	return newBasicAddress()
}

func TestSetAndGetProfile(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := compliancetypes.Profile{
		Address:   addr.String(),
		KYCLevel:  "standard",
		Risk:      compliancetypes.RiskLow,
		Sanction:  false,
		Metadata:  "initial creation",
		UpdatedBy: newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	result, found := k.GetProfile(wctx, addr)
	require.True(t, found)
	require.Equal(t, profile.Address, result.Address)
	require.Equal(t, profile.KYCLevel, result.KYCLevel)
	require.Equal(t, profile.Risk, result.Risk)
}

func TestAssertCompliant(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()

	wctx := sdk.WrapSDKContext(ctx)

	err := k.AssertCompliant(wctx, addr)
	require.ErrorIs(t, err, compliancetypes.ErrProfileNotFound)

	profile := compliancetypes.Profile{
		Address:   addr.String(),
		KYCLevel:  "standard",
		Risk:      compliancetypes.RiskLow,
		Status:    compliancetypes.StatusActive,
		UpdatedBy: newBasicAddress().String(),
	}
	k.SetProfile(wctx, profile)

	require.NoError(t, k.AssertCompliant(wctx, addr))

	profile.Sanction = true
	k.SetProfile(wctx, profile)

	err = k.AssertCompliant(wctx, addr)
	require.ErrorIs(t, err, compliancetypes.ErrSanctionedAddress)
}

func TestRemoveProfileBasic(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := compliancetypes.Profile{
		Address:   addr.String(),
		KYCLevel:  "standard",
		Risk:      compliancetypes.RiskLow,
		UpdatedBy: newBasicAddress().String(),
	}
	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	_, found := k.GetProfile(wctx, addr)
	require.True(t, found)

	k.RemoveProfile(wctx, addr)

	_, found = k.GetProfile(wctx, addr)
	require.False(t, found)
}
