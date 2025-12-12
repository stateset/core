package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/compliance/types"
)

// TestAssertCompliant_EdgeCases tests edge cases for compliance checks
func TestAssertCompliant_SanctionedAddress(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:   addr.String(),
		KYCLevel:  types.KYCStandard,
		Risk:      types.RiskLow,
		Status:    types.StatusActive,
		Sanction:  true, // Sanctioned
		UpdatedBy: newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	err := k.AssertCompliant(wctx, addr)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrSanctionedAddress)
}

func TestAssertCompliant_BlockedStatus(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:   addr.String(),
		KYCLevel:  types.KYCStandard,
		Risk:      types.RiskLow,
		Status:    types.StatusSuspended, // Blocked
		UpdatedBy: newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	err := k.AssertCompliant(wctx, addr)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrComplianceBlocked)
}

func TestAssertCompliant_ExpiredProfile(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	pastTime := ctx.BlockTime().AddDate(-2, 0, 0) // 2 years ago
	profile := types.Profile{
		Address:    addr.String(),
		KYCLevel:   types.KYCStandard,
		Risk:       types.RiskLow,
		Status:     types.StatusActive,
		VerifiedAt: pastTime,
		ExpiresAt:  pastTime.AddDate(1, 0, 0), // Expired 1 year ago
		UpdatedBy:  newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	err := k.AssertCompliant(wctx, addr)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProfileExpired)
}

func TestAssertCompliant_ProfileNotFound(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	wctx := sdk.WrapSDKContext(ctx)

	err := k.AssertCompliant(wctx, addr)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProfileNotFound)
}

func TestAssertCompliant_ValidProfile(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:    addr.String(),
		KYCLevel:   types.KYCStandard,
		Risk:       types.RiskLow,
		Status:     types.StatusActive,
		Sanction:   false,
		VerifiedAt: ctx.BlockTime(),
		ExpiresAt:  ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:  newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	err := k.AssertCompliant(wctx, addr)
	require.NoError(t, err)
}

// TestAssertCompliantForAmount tests amount-based compliance checks
func TestAssertCompliantForAmount_ExceedsDailyLimit(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	dailyLimit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))
	dailyUsed := sdk.NewCoin("ssusd", sdkmath.NewInt(900000))

	profile := types.Profile{
		Address:         addr.String(),
		KYCLevel:        types.KYCStandard,
		Risk:            types.RiskLow,
		Status:          types.StatusActive,
		DailyLimit:      dailyLimit,
		DailyUsed:       dailyUsed,
		LastLimitReset:  ctx.BlockTime(),
		VerifiedAt:      ctx.BlockTime(),
		ExpiresAt:       ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:       newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Try to transact 200000, which would exceed daily limit
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(200000))
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrLimitExceeded)
}

func TestAssertCompliantForAmount_WithinDailyLimit(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	dailyLimit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))
	dailyUsed := sdk.NewCoin("ssusd", sdkmath.NewInt(900000))

	profile := types.Profile{
		Address:         addr.String(),
		KYCLevel:        types.KYCStandard,
		Risk:            types.RiskLow,
		Status:          types.StatusActive,
		DailyLimit:      dailyLimit,
		DailyUsed:       dailyUsed,
		LastLimitReset:  ctx.BlockTime(),
		VerifiedAt:      ctx.BlockTime(),
		ExpiresAt:       ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:       newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Transact 50000, which is within limit
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(50000))
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.NoError(t, err)
}

func TestAssertCompliantForAmount_ExceedsMonthlyLimit(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	monthlyLimit := sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))
	monthlyUsed := sdk.NewCoin("ssusd", sdkmath.NewInt(9500000))

	profile := types.Profile{
		Address:         addr.String(),
		KYCLevel:        types.KYCStandard,
		Risk:            types.RiskLow,
		Status:          types.StatusActive,
		MonthlyLimit:    monthlyLimit,
		MonthlyUsed:     monthlyUsed,
		LastLimitReset:  ctx.BlockTime(),
		VerifiedAt:      ctx.BlockTime(),
		ExpiresAt:       ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:       newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Try to transact 600000, which would exceed monthly limit
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(600000))
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrLimitExceeded)
}

func TestAssertCompliantForAmount_ZeroAmount(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:    addr.String(),
		KYCLevel:   types.KYCStandard,
		Risk:       types.RiskLow,
		Status:     types.StatusActive,
		VerifiedAt: ctx.BlockTime(),
		ExpiresAt:  ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:  newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Zero amount should pass
	amount := sdk.NewCoin("ssusd", sdkmath.ZeroInt())
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.NoError(t, err)
}

func TestAssertCompliantForAmount_HighRiskRequiresEnhancedKYC(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:    addr.String(),
		KYCLevel:   types.KYCStandard, // Not enhanced
		Risk:       types.RiskHigh,    // High risk
		Status:     types.StatusActive,
		VerifiedAt: ctx.BlockTime(),
		ExpiresAt:  ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:  newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrEnhancedDueDiligenceRequired)
}

func TestAssertCompliantForAmount_EnhancedKYCHighRisk(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:    addr.String(),
		KYCLevel:   types.KYCEnhanced, // Enhanced KYC
		Risk:       types.RiskHigh,    // High risk is OK with enhanced KYC
		Status:     types.StatusActive,
		VerifiedAt: ctx.BlockTime(),
		ExpiresAt:  ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:  newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.NoError(t, err)
}

// TestRecordTransaction tests transaction recording
func TestRecordTransaction_UpdatesUsage(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:         addr.String(),
		KYCLevel:        types.KYCStandard,
		Risk:            types.RiskLow,
		Status:          types.StatusActive,
		DailyUsed:       sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		MonthlyUsed:     sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		LastLimitReset:  ctx.BlockTime(),
		UpdatedBy:       newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Record transaction
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
	err := k.RecordTransaction(wctx, addr, amount)
	require.NoError(t, err)

	// Check updated usage
	updatedProfile, found := k.GetProfile(wctx, addr)
	require.True(t, found)
	require.Equal(t, amount, updatedProfile.DailyUsed)
	require.Equal(t, amount, updatedProfile.MonthlyUsed)
}

func TestRecordTransaction_AccumulatesUsage(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	profile := types.Profile{
		Address:         addr.String(),
		KYCLevel:        types.KYCStandard,
		Risk:            types.RiskLow,
		Status:          types.StatusActive,
		DailyUsed:       sdk.NewCoin("ssusd", sdkmath.NewInt(50000)),
		MonthlyUsed:     sdk.NewCoin("ssusd", sdkmath.NewInt(50000)),
		LastLimitReset:  ctx.BlockTime(),
		UpdatedBy:       newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Record additional transaction
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(30000))
	err := k.RecordTransaction(wctx, addr, amount)
	require.NoError(t, err)

	// Check accumulated usage
	updatedProfile, found := k.GetProfile(wctx, addr)
	require.True(t, found)
	require.Equal(t, sdk.NewCoin("ssusd", sdkmath.NewInt(80000)), updatedProfile.DailyUsed)
	require.Equal(t, sdk.NewCoin("ssusd", sdkmath.NewInt(80000)), updatedProfile.MonthlyUsed)
}

func TestRecordTransaction_ProfileNotFound(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	wctx := sdk.WrapSDKContext(ctx)

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
	err := k.RecordTransaction(wctx, addr, amount)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProfileNotFound)
}

// TestLimitReset tests limit reset functionality
func TestLimitReset_DailyReset(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	oldTime := ctx.BlockTime().Add(-25 * time.Hour) // 25 hours ago
	profile := types.Profile{
		Address:         addr.String(),
		KYCLevel:        types.KYCStandard,
		Risk:            types.RiskLow,
		Status:          types.StatusActive,
		DailyUsed:       sdk.NewCoin("ssusd", sdkmath.NewInt(500000)),
		MonthlyUsed:     sdk.NewCoin("ssusd", sdkmath.NewInt(500000)),
		LastLimitReset:  oldTime,
		VerifiedAt:      ctx.BlockTime(),
		ExpiresAt:       ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:       newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Check compliance - should reset daily limit
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.NoError(t, err)

	// Verify daily limit was reset
	updatedProfile, _ := k.GetProfile(wctx, addr)
	require.Equal(t, sdk.NewCoin("ssusd", sdkmath.ZeroInt()), updatedProfile.DailyUsed)
}

func TestLimitReset_MonthlyReset(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	oldTime := ctx.BlockTime().AddDate(0, -1, -5) // Over a month ago
	profile := types.Profile{
		Address:         addr.String(),
		KYCLevel:        types.KYCStandard,
		Risk:            types.RiskLow,
		Status:          types.StatusActive,
		DailyUsed:       sdk.NewCoin("ssusd", sdkmath.NewInt(100000)),
		MonthlyUsed:     sdk.NewCoin("ssusd", sdkmath.NewInt(5000000)),
		LastLimitReset:  oldTime,
		VerifiedAt:      ctx.BlockTime(),
		ExpiresAt:       ctx.BlockTime().AddDate(1, 0, 0),
		UpdatedBy:       newBasicAddress().String(),
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Check compliance - should reset both limits
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
	err := k.AssertCompliantForAmount(wctx, addr, amount)
	require.NoError(t, err)

	// Verify both limits were reset
	updatedProfile, _ := k.GetProfile(wctx, addr)
	require.Equal(t, sdk.NewCoin("ssusd", sdkmath.ZeroInt()), updatedProfile.DailyUsed)
	require.Equal(t, sdk.NewCoin("ssusd", sdkmath.ZeroInt()), updatedProfile.MonthlyUsed)
}

// TestSuspendProfile tests profile suspension
func TestSuspendProfile_EdgeCases(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	actor := newBasicAddress().String()
	profile := types.Profile{
		Address:   addr.String(),
		KYCLevel:  types.KYCStandard,
		Risk:      types.RiskLow,
		Status:    types.StatusActive,
		UpdatedBy: actor,
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Suspend profile
	err := k.SuspendProfile(wctx, addr, actor, "suspicious activity")
	require.NoError(t, err)

	// Verify suspended
	suspended, found := k.GetProfile(wctx, addr)
	require.True(t, found)
	require.Equal(t, types.StatusSuspended, suspended.Status)
	require.Equal(t, actor, suspended.UpdatedBy)
}

func TestSuspendProfile_NotFound(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	actor := newBasicAddress().String()
	wctx := sdk.WrapSDKContext(ctx)

	err := k.SuspendProfile(wctx, addr, actor, "test")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProfileNotFound)
}

// TestReactivateProfile tests profile reactivation
func TestReactivateProfile_EdgeCases(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	actor := newBasicAddress().String()
	profile := types.Profile{
		Address:   addr.String(),
		KYCLevel:  types.KYCStandard,
		Risk:      types.RiskLow,
		Status:    types.StatusSuspended,
		UpdatedBy: actor,
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Reactivate profile
	err := k.ReactivateProfile(wctx, addr, actor, "cleared investigation")
	require.NoError(t, err)

	// Verify reactivated
	reactivated, found := k.GetProfile(wctx, addr)
	require.True(t, found)
	require.Equal(t, types.StatusActive, reactivated.Status)
}

func TestReactivateProfile_NotSuspended(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	actor := newBasicAddress().String()
	profile := types.Profile{
		Address:   addr.String(),
		KYCLevel:  types.KYCStandard,
		Risk:      types.RiskLow,
		Status:    types.StatusActive, // Not suspended
		UpdatedBy: actor,
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	err := k.ReactivateProfile(wctx, addr, actor, "test")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidProfileStatus)
}

// TestUpdateKYCLevel tests KYC level updates
func TestUpdateKYCLevel_EdgeCases(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	actor := newBasicAddress().String()
	profile := types.Profile{
		Address:   addr.String(),
		KYCLevel:  types.KYCBasic,
		Risk:      types.RiskLow,
		Status:    types.StatusActive,
		UpdatedBy: actor,
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Update to standard KYC
	err := k.UpdateKYCLevel(wctx, addr, types.KYCStandard, actor, "upgraded verification")
	require.NoError(t, err)

	// Verify update
	updated, found := k.GetProfile(wctx, addr)
	require.True(t, found)
	require.Equal(t, types.KYCStandard, updated.KYCLevel)
	require.False(t, updated.VerifiedAt.IsZero())
	require.False(t, updated.ExpiresAt.IsZero())
}

func TestUpdateKYCLevel_ToEnhanced(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	actor := newBasicAddress().String()
	profile := types.Profile{
		Address:   addr.String(),
		KYCLevel:  types.KYCStandard,
		Risk:      types.RiskMedium,
		Status:    types.StatusActive,
		UpdatedBy: actor,
	}

	wctx := sdk.WrapSDKContext(ctx)
	k.SetProfile(wctx, profile)

	// Update to enhanced KYC
	err := k.UpdateKYCLevel(wctx, addr, types.KYCEnhanced, actor, "high-value customer")
	require.NoError(t, err)

	// Verify update
	updated, found := k.GetProfile(wctx, addr)
	require.True(t, found)
	require.Equal(t, types.KYCEnhanced, updated.KYCLevel)
}

func TestUpdateKYCLevel_ProfileNotFound(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	addr := newBasicAddress()
	actor := newBasicAddress().String()
	wctx := sdk.WrapSDKContext(ctx)

	err := k.UpdateKYCLevel(wctx, addr, types.KYCStandard, actor, "test")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProfileNotFound)
}

// TestIterateProfiles tests profile iteration
func TestIterateProfiles_EdgeCases(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	wctx := sdk.WrapSDKContext(ctx)

	// Create multiple profiles
	for i := 0; i < 5; i++ {
		addr := newBasicAddress()
		profile := types.Profile{
			Address:   addr.String(),
			KYCLevel:  types.KYCStandard,
			Risk:      types.RiskLow,
			Status:    types.StatusActive,
			UpdatedBy: newBasicAddress().String(),
		}
		k.SetProfile(wctx, profile)
	}

	// Iterate and count
	count := 0
	k.IterateProfiles(wctx, func(p types.Profile) bool {
		count++
		return false
	})
	require.Equal(t, 5, count)
}

func TestIterateProfiles_EarlyStop(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	wctx := sdk.WrapSDKContext(ctx)

	// Create multiple profiles
	for i := 0; i < 10; i++ {
		addr := newBasicAddress()
		profile := types.Profile{
			Address:   addr.String(),
			KYCLevel:  types.KYCStandard,
			Risk:      types.RiskLow,
			Status:    types.StatusActive,
			UpdatedBy: newBasicAddress().String(),
		}
		k.SetProfile(wctx, profile)
	}

	// Iterate and stop after 3
	count := 0
	k.IterateProfiles(wctx, func(p types.Profile) bool {
		count++
		return count >= 3
	})
	require.Equal(t, 3, count)
}

// TestGenesisExportImport tests genesis state
func TestGenesisExportImport_EdgeCases(t *testing.T) {
	k, ctx := setupBasicKeeper(t)

	wctx := sdk.WrapSDKContext(ctx)

	// Create profiles
	profiles := make([]types.Profile, 3)
	for i := 0; i < 3; i++ {
		addr := newBasicAddress()
		profiles[i] = types.Profile{
			Address:   addr.String(),
			KYCLevel:  types.KYCStandard,
			Risk:      types.RiskLow,
			Status:    types.StatusActive,
			UpdatedBy: newBasicAddress().String(),
		}
		k.SetProfile(wctx, profiles[i])
	}

	// Export genesis
	genesis := k.ExportGenesis(wctx)
	require.NotNil(t, genesis)
	require.Len(t, genesis.Profiles, 3)

	// Create new keeper and import
	k2, ctx2 := setupBasicKeeper(t)
	wctx2 := sdk.WrapSDKContext(ctx2)
	k2.InitGenesis(wctx2, genesis)

	// Verify imported profiles
	count := 0
	k2.IterateProfiles(wctx2, func(p types.Profile) bool {
		count++
		return false
	})
	require.Equal(t, 3, count)
}
