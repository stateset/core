package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	compliancekeeper "github.com/stateset/core/x/compliance/keeper"
	compliancetypes "github.com/stateset/core/x/compliance/types"
)

// genTestAddress generates a random test address
func genTestAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

func setupKeeper(t *testing.T) (compliancekeeper.Keeper, sdk.Context) {
	// Call the shared config setup from keeper_test.go
	setupBasicConfig()

	storeKey := storetypes.NewKVStoreKey(compliancetypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	authority := genTestAddress()
	k := compliancekeeper.NewKeeper(
		cdc,
		storeKey,
		authority.String(),
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	return k, ctx
}

func createTestProfile(addr string) compliancetypes.Profile {
	return compliancetypes.Profile{
		Address:      addr,
		KYCLevel:     compliancetypes.KYCStandard,
		Risk:         compliancetypes.RiskLow,
		Status:       compliancetypes.StatusActive,
		Sanction:     false,
		Jurisdiction: "US",
		BusinessType: "individual",
		DailyLimit:   sdk.NewCoin("uusdc", sdkmath.NewInt(100000000)), // 100 USDC
		MonthlyLimit: sdk.NewCoin("uusdc", sdkmath.NewInt(1000000000)), // 1000 USDC
		DailyUsed:    sdk.NewCoin("uusdc", sdkmath.ZeroInt()),
		MonthlyUsed:  sdk.NewCoin("uusdc", sdkmath.ZeroInt()),
	}
}

func TestSetGetProfile(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())

	// Set profile
	k.SetProfile(ctx, profile)

	// Get profile
	retrieved, found := k.GetProfile(ctx, addr)
	require.True(t, found)
	require.Equal(t, profile.Address, retrieved.Address)
	require.Equal(t, profile.KYCLevel, retrieved.KYCLevel)
	require.Equal(t, profile.Risk, retrieved.Risk)
	require.Equal(t, profile.Status, retrieved.Status)
}

func TestProfileNotFound(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()

	// Profile should not exist
	_, found := k.GetProfile(ctx, addr)
	require.False(t, found)

	// AssertCompliant should fail
	err := k.AssertCompliant(ctx, addr)
	require.Error(t, err)
	require.ErrorIs(t, err, compliancetypes.ErrProfileNotFound)
}

func TestAssertCompliantSanctioned(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.Sanction = true

	k.SetProfile(ctx, profile)

	// Should fail due to sanction
	err := k.AssertCompliant(ctx, addr)
	require.Error(t, err)
	require.ErrorIs(t, err, compliancetypes.ErrSanctionedAddress)
}

func TestAssertCompliantBlockedStatus(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.Status = compliancetypes.StatusSuspended

	k.SetProfile(ctx, profile)

	// Should fail due to suspended status
	err := k.AssertCompliant(ctx, addr)
	require.Error(t, err)
	require.ErrorContains(t, err, "blocked")
}

func TestAssertCompliantBlockedJurisdiction(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.Jurisdiction = "KP" // North Korea - blocked

	k.SetProfile(ctx, profile)

	// Should fail due to blocked jurisdiction
	err := k.AssertCompliant(ctx, addr)
	require.Error(t, err)
	require.ErrorContains(t, err, "blocked")
}

func TestAssertCompliantExpired(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.ExpiresAt = ctx.BlockTime().Add(-24 * time.Hour) // Expired yesterday

	k.SetProfile(ctx, profile)

	// Should fail due to expiration
	err := k.AssertCompliant(ctx, addr)
	require.Error(t, err)
	require.ErrorContains(t, err, "expired")
}

func TestAssertCompliantForAmountDailyLimit(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.DailyLimit = sdk.NewCoin("uusdc", sdkmath.NewInt(1000))
	profile.DailyUsed = sdk.NewCoin("uusdc", sdkmath.NewInt(900))
	// Set LastLimitReset to current time so limits don't get auto-reset
	profile.LastLimitReset = ctx.BlockTime()

	k.SetProfile(ctx, profile)

	// Small amount should pass
	err := k.AssertCompliantForAmount(ctx, addr, sdk.NewCoin("uusdc", sdkmath.NewInt(50)))
	require.NoError(t, err)

	// Amount exceeding limit should fail
	err = k.AssertCompliantForAmount(ctx, addr, sdk.NewCoin("uusdc", sdkmath.NewInt(200)))
	require.Error(t, err)
	require.ErrorContains(t, err, "daily limit exceeded")
}

func TestAssertCompliantForAmountMonthlyLimit(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.DailyLimit = sdk.NewCoin("uusdc", sdkmath.NewInt(10000))    // High daily
	profile.MonthlyLimit = sdk.NewCoin("uusdc", sdkmath.NewInt(1000))  // Low monthly
	profile.MonthlyUsed = sdk.NewCoin("uusdc", sdkmath.NewInt(900))
	// Set LastLimitReset to current time so limits don't get auto-reset
	profile.LastLimitReset = ctx.BlockTime()

	k.SetProfile(ctx, profile)

	// Amount exceeding monthly limit should fail
	err := k.AssertCompliantForAmount(ctx, addr, sdk.NewCoin("uusdc", sdkmath.NewInt(200)))
	require.Error(t, err)
	require.ErrorContains(t, err, "monthly limit exceeded")
}

func TestRecordTransaction(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())

	k.SetProfile(ctx, profile)

	// Record a transaction
	amount := sdk.NewCoin("uusdc", sdkmath.NewInt(1000))
	err := k.RecordTransaction(ctx, addr, amount)
	require.NoError(t, err)

	// Verify usage was updated
	updated, _ := k.GetProfile(ctx, addr)
	require.Equal(t, sdkmath.NewInt(1000), updated.DailyUsed.Amount)
	require.Equal(t, sdkmath.NewInt(1000), updated.MonthlyUsed.Amount)

	// Record another transaction
	err = k.RecordTransaction(ctx, addr, amount)
	require.NoError(t, err)

	// Verify cumulative usage
	updated, _ = k.GetProfile(ctx, addr)
	require.Equal(t, sdkmath.NewInt(2000), updated.DailyUsed.Amount)
	require.Equal(t, sdkmath.NewInt(2000), updated.MonthlyUsed.Amount)
}

func TestSuspendProfile(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())

	k.SetProfile(ctx, profile)

	// Suspend the profile
	err := k.SuspendProfile(ctx, addr, "operator", "suspicious activity")
	require.NoError(t, err)

	// Verify suspension
	updated, _ := k.GetProfile(ctx, addr)
	require.Equal(t, compliancetypes.StatusSuspended, updated.Status)
	require.Equal(t, "operator", updated.UpdatedBy)
	require.Len(t, updated.AuditLog, 1)
	require.Equal(t, "suspended", updated.AuditLog[0].Action)
	require.Equal(t, "suspicious activity", updated.AuditLog[0].Reason)

	// AssertCompliant should now fail
	err = k.AssertCompliant(ctx, addr)
	require.Error(t, err)
}

func TestReactivateProfile(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.Status = compliancetypes.StatusSuspended

	k.SetProfile(ctx, profile)

	// Reactivate the profile
	err := k.ReactivateProfile(ctx, addr, "admin", "cleared by review")
	require.NoError(t, err)

	// Verify reactivation
	updated, _ := k.GetProfile(ctx, addr)
	require.Equal(t, compliancetypes.StatusActive, updated.Status)
	require.Len(t, updated.AuditLog, 1)
	require.Equal(t, "reactivated", updated.AuditLog[0].Action)

	// AssertCompliant should now pass
	err = k.AssertCompliant(ctx, addr)
	require.NoError(t, err)
}

func TestReactivateNonSuspendedProfile(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())

	k.SetProfile(ctx, profile)

	// Try to reactivate an active profile
	err := k.ReactivateProfile(ctx, addr, "admin", "test")
	require.Error(t, err)
	require.ErrorContains(t, err, "not suspended")
}

func TestUpdateKYCLevel(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.KYCLevel = compliancetypes.KYCBasic

	k.SetProfile(ctx, profile)

	// Update to enhanced KYC
	err := k.UpdateKYCLevel(ctx, addr, compliancetypes.KYCEnhanced, "compliance_officer", "documents verified")
	require.NoError(t, err)

	// Verify update
	updated, _ := k.GetProfile(ctx, addr)
	require.Equal(t, compliancetypes.KYCEnhanced, updated.KYCLevel)
	require.False(t, updated.VerifiedAt.IsZero())
	require.False(t, updated.ExpiresAt.IsZero())
	require.Len(t, updated.AuditLog, 1)
	require.Equal(t, "kyc_updated", updated.AuditLog[0].Action)
}

func TestEnhancedDueDiligenceRequired(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.Risk = compliancetypes.RiskHigh
	profile.KYCLevel = compliancetypes.KYCBasic

	k.SetProfile(ctx, profile)

	// High-risk profile with basic KYC should fail EDD check
	err := k.AssertCompliantForAmount(ctx, addr, sdk.NewCoin("uusdc", sdkmath.NewInt(100)))
	require.Error(t, err)
	require.ErrorContains(t, err, "enhanced KYC required")

	// Upgrade KYC
	err = k.UpdateKYCLevel(ctx, addr, compliancetypes.KYCEnhanced, "admin", "EDD completed")
	require.NoError(t, err)

	// Now should pass
	err = k.AssertCompliantForAmount(ctx, addr, sdk.NewCoin("uusdc", sdkmath.NewInt(100)))
	require.NoError(t, err)
}

func TestHighRiskJurisdictionRequiresEDD(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.Jurisdiction = "AF" // Afghanistan - high risk
	profile.KYCLevel = compliancetypes.KYCBasic

	k.SetProfile(ctx, profile)

	// High-risk jurisdiction with basic KYC should fail
	err := k.AssertCompliantForAmount(ctx, addr, sdk.NewCoin("uusdc", sdkmath.NewInt(100)))
	require.Error(t, err)
	require.ErrorContains(t, err, "enhanced KYC required")
}

func TestRemoveProfile(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())

	k.SetProfile(ctx, profile)

	// Verify exists
	_, found := k.GetProfile(ctx, addr)
	require.True(t, found)

	// Remove
	k.RemoveProfile(ctx, addr)

	// Verify removed
	_, found = k.GetProfile(ctx, addr)
	require.False(t, found)
}

func TestIterateProfiles(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Create multiple profiles with generated addresses
	var addrs []sdk.AccAddress
	for i := 0; i < 3; i++ {
		addrs = append(addrs, genTestAddress())
	}

	for _, addr := range addrs {
		profile := createTestProfile(addr.String())
		k.SetProfile(ctx, profile)
	}

	// Iterate and count
	var count int
	k.IterateProfiles(ctx, func(profile compliancetypes.Profile) bool {
		count++
		return false
	})

	require.Equal(t, 3, count)
}

func TestGenesisExportImport(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up profiles
	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	k.SetProfile(ctx, profile)

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.Len(t, genesis.Profiles, 1)
	require.NotEmpty(t, genesis.Authority)

	// Create new keeper and import
	k2, ctx2 := setupKeeper(t)
	k2.InitGenesis(ctx2, genesis)

	// Verify imported
	imported, found := k2.GetProfile(ctx2, addr)
	require.True(t, found)
	require.Equal(t, profile.Address, imported.Address)
	require.Equal(t, profile.KYCLevel, imported.KYCLevel)
}

func TestAuditLogMaxSize(t *testing.T) {
	k, ctx := setupKeeper(t)

	addr := genTestAddress()
	profile := createTestProfile(addr.String())
	profile.Status = compliancetypes.StatusSuspended

	k.SetProfile(ctx, profile)

	// Perform many actions to fill audit log
	for i := 0; i < 150; i++ {
		if i%2 == 0 {
			k.ReactivateProfile(ctx, addr, "admin", "test reactivate")
			k.SuspendProfile(ctx, addr, "admin", "test suspend")
		}
	}

	// Verify audit log is capped at 100
	updated, _ := k.GetProfile(ctx, addr)
	require.LessOrEqual(t, len(updated.AuditLog), 100)
}

func TestProfileIsBlocked(t *testing.T) {
	tests := []struct {
		name     string
		profile  compliancetypes.Profile
		expected bool
	}{
		{
			name: "active profile not blocked",
			profile: compliancetypes.Profile{
				Status:       compliancetypes.StatusActive,
				Sanction:     false,
				Jurisdiction: "US",
			},
			expected: false,
		},
		{
			name: "sanctioned profile blocked",
			profile: compliancetypes.Profile{
				Status:       compliancetypes.StatusActive,
				Sanction:     true,
				Jurisdiction: "US",
			},
			expected: true,
		},
		{
			name: "suspended profile blocked",
			profile: compliancetypes.Profile{
				Status:       compliancetypes.StatusSuspended,
				Sanction:     false,
				Jurisdiction: "US",
			},
			expected: true,
		},
		{
			name: "blocked jurisdiction",
			profile: compliancetypes.Profile{
				Status:       compliancetypes.StatusActive,
				Sanction:     false,
				Jurisdiction: "KP", // North Korea
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.profile.IsBlocked())
		})
	}
}

func TestProfileRequiresEnhancedDueDiligence(t *testing.T) {
	tests := []struct {
		name     string
		profile  compliancetypes.Profile
		expected bool
	}{
		{
			name: "low risk US profile",
			profile: compliancetypes.Profile{
				Risk:         compliancetypes.RiskLow,
				Jurisdiction: "US",
			},
			expected: false,
		},
		{
			name: "high risk profile",
			profile: compliancetypes.Profile{
				Risk:         compliancetypes.RiskHigh,
				Jurisdiction: "US",
			},
			expected: true,
		},
		{
			name: "high risk jurisdiction",
			profile: compliancetypes.Profile{
				Risk:         compliancetypes.RiskLow,
				Jurisdiction: "AF", // Afghanistan
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.profile.RequiresEnhancedDueDiligence())
		})
	}
}
