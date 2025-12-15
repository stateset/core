package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/compliance/types"
)

func TestMsgUpsertProfile_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()
	validAddress := sdk.AccAddress("address_____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgUpsertProfile
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgUpsertProfile{
				Authority: validAuthority,
				Profile: types.Profile{
					Address:      validAddress,
					KYCLevel:     types.KYCStandard,
					Risk:         types.RiskLow,
					Jurisdiction: "US",
				},
			},
			expectErr: false,
		},
		{
			name: "invalid authority address",
			msg: &types.MsgUpsertProfile{
				Authority: "invalid",
				Profile: types.Profile{
					Address:      validAddress,
					KYCLevel:     types.KYCStandard,
					Risk:         types.RiskLow,
					Jurisdiction: "US",
				},
			},
			expectErr: true,
		},
		{
			name: "invalid profile address",
			msg: &types.MsgUpsertProfile{
				Authority: validAuthority,
				Profile: types.Profile{
					Address:      "invalid",
					KYCLevel:     types.KYCStandard,
					Risk:         types.RiskLow,
					Jurisdiction: "US",
				},
			},
			expectErr: true,
		},
		{
			name: "empty kyc level",
			msg: &types.MsgUpsertProfile{
				Authority: validAuthority,
				Profile: types.Profile{
					Address:      validAddress,
					KYCLevel:     "",
					Risk:         types.RiskLow,
					Jurisdiction: "US",
				},
			},
			expectErr: true,
		},
		{
			name: "empty risk level",
			msg: &types.MsgUpsertProfile{
				Authority: validAuthority,
				Profile: types.Profile{
					Address:      validAddress,
					KYCLevel:     types.KYCStandard,
					Risk:         "",
					Jurisdiction: "US",
				},
			},
			expectErr: true,
		},
		{
			name: "invalid jurisdiction length",
			msg: &types.MsgUpsertProfile{
				Authority: validAuthority,
				Profile: types.Profile{
					Address:      validAddress,
					KYCLevel:     types.KYCStandard,
					Risk:         types.RiskLow,
					Jurisdiction: "USA", // should be 2 chars
				},
			},
			expectErr: true,
		},
		{
			name: "empty jurisdiction is allowed",
			msg: &types.MsgUpsertProfile{
				Authority: validAuthority,
				Profile: types.Profile{
					Address:  validAddress,
					KYCLevel: types.KYCStandard,
					Risk:     types.RiskLow,
				},
			},
			expectErr: false,
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

func TestMsgSetSanction_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()
	validAddress := sdk.AccAddress("address_____________").String()

	tests := []struct {
		name      string
		msg       *types.MsgSetSanction
		expectErr bool
	}{
		{
			name: "valid sanction message",
			msg: &types.MsgSetSanction{
				Authority: validAuthority,
				Address:   validAddress,
				Sanction:  true,
				Reason:    "OFAC listing",
			},
			expectErr: false,
		},
		{
			name: "valid unsanction message",
			msg: &types.MsgSetSanction{
				Authority: validAuthority,
				Address:   validAddress,
				Sanction:  false,
				Reason:    "Removed from OFAC",
			},
			expectErr: false,
		},
		{
			name: "invalid authority address",
			msg: &types.MsgSetSanction{
				Authority: "invalid",
				Address:   validAddress,
				Sanction:  true,
			},
			expectErr: true,
		},
		{
			name: "invalid target address",
			msg: &types.MsgSetSanction{
				Authority: validAuthority,
				Address:   "invalid",
				Sanction:  true,
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

func TestProfile_ValidateBasic(t *testing.T) {
	validAddress := sdk.AccAddress("address_____________").String()

	tests := []struct {
		name      string
		profile   types.Profile
		expectErr bool
	}{
		{
			name: "valid profile",
			profile: types.Profile{
				Address:      validAddress,
				KYCLevel:     types.KYCStandard,
				Risk:         types.RiskLow,
				Jurisdiction: "US",
			},
			expectErr: false,
		},
		{
			name: "invalid address",
			profile: types.Profile{
				Address:  "invalid",
				KYCLevel: types.KYCStandard,
				Risk:     types.RiskLow,
			},
			expectErr: true,
		},
		{
			name: "empty kyc level",
			profile: types.Profile{
				Address: validAddress,
				Risk:    types.RiskLow,
			},
			expectErr: true,
		},
		{
			name: "empty risk level",
			profile: types.Profile{
				Address:  validAddress,
				KYCLevel: types.KYCStandard,
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.profile.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestProfile_IsBlocked(t *testing.T) {
	validAddress := sdk.AccAddress("address_____________").String()

	tests := []struct {
		name     string
		profile  types.Profile
		expected bool
	}{
		{
			name: "active profile not blocked",
			profile: types.Profile{
				Address:  validAddress,
				Status:   types.StatusActive,
				Sanction: false,
			},
			expected: false,
		},
		{
			name: "sanctioned profile is blocked",
			profile: types.Profile{
				Address:  validAddress,
				Status:   types.StatusActive,
				Sanction: true,
			},
			expected: true,
		},
		{
			name: "pending profile is blocked",
			profile: types.Profile{
				Address:  validAddress,
				Status:   types.StatusPending,
				Sanction: false,
			},
			expected: true,
		},
		{
			name: "suspended profile is blocked",
			profile: types.Profile{
				Address:  validAddress,
				Status:   types.StatusSuspended,
				Sanction: false,
			},
			expected: true,
		},
		{
			name: "rejected profile is blocked",
			profile: types.Profile{
				Address:  validAddress,
				Status:   types.StatusRejected,
				Sanction: false,
			},
			expected: true,
		},
		{
			name: "blocked jurisdiction is blocked",
			profile: types.Profile{
				Address:      validAddress,
				Status:       types.StatusActive,
				Sanction:     false,
				Jurisdiction: "KP", // North Korea
			},
			expected: true,
		},
		{
			name: "Iran is blocked",
			profile: types.Profile{
				Address:      validAddress,
				Status:       types.StatusActive,
				Sanction:     false,
				Jurisdiction: "IR",
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.profile.IsBlocked()
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestProfile_RequiresEnhancedDueDiligence(t *testing.T) {
	validAddress := sdk.AccAddress("address_____________").String()

	tests := []struct {
		name     string
		profile  types.Profile
		expected bool
	}{
		{
			name: "low risk profile does not require EDD",
			profile: types.Profile{
				Address:      validAddress,
				Risk:         types.RiskLow,
				Jurisdiction: "US",
			},
			expected: false,
		},
		{
			name: "medium risk profile does not require EDD",
			profile: types.Profile{
				Address:      validAddress,
				Risk:         types.RiskMedium,
				Jurisdiction: "US",
			},
			expected: false,
		},
		{
			name: "high risk profile requires EDD",
			profile: types.Profile{
				Address:      validAddress,
				Risk:         types.RiskHigh,
				Jurisdiction: "US",
			},
			expected: true,
		},
		{
			name: "high risk jurisdiction requires EDD",
			profile: types.Profile{
				Address:      validAddress,
				Risk:         types.RiskLow,
				Jurisdiction: "AF", // Afghanistan
			},
			expected: true,
		},
		{
			name: "Belarus requires EDD",
			profile: types.Profile{
				Address:      validAddress,
				Risk:         types.RiskLow,
				Jurisdiction: "BY",
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.profile.RequiresEnhancedDueDiligence()
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestProfile_IsExpired(t *testing.T) {
	validAddress := sdk.AccAddress("address_____________").String()
	now := time.Now()

	tests := []struct {
		name     string
		profile  types.Profile
		now      time.Time
		expected bool
	}{
		{
			name: "zero expiry means never expires",
			profile: types.Profile{
				Address: validAddress,
			},
			now:      now,
			expected: false,
		},
		{
			name: "future expiry not expired",
			profile: types.Profile{
				Address:   validAddress,
				ExpiresAt: now.Add(24 * time.Hour),
			},
			now:      now,
			expected: false,
		},
		{
			name: "past expiry is expired",
			profile: types.Profile{
				Address:   validAddress,
				ExpiresAt: now.Add(-24 * time.Hour),
			},
			now:      now,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.profile.IsExpired(tc.now)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestProfile_AddAuditEntry(t *testing.T) {
	validAddress := sdk.AccAddress("address_____________").String()

	profile := types.Profile{
		Address: validAddress,
	}

	// Add first entry
	profile.AddAuditEntry("status_change", "admin", "initial verification", types.StatusPending, types.StatusActive)
	require.Len(t, profile.AuditLog, 1)
	require.Equal(t, "status_change", profile.AuditLog[0].Action)
	require.Equal(t, "admin", profile.AuditLog[0].Actor)
	require.Equal(t, "initial verification", profile.AuditLog[0].Reason)
	require.Equal(t, string(types.StatusPending), profile.AuditLog[0].OldStatus)
	require.Equal(t, string(types.StatusActive), profile.AuditLog[0].NewStatus)

	// Add multiple entries and verify log truncation at 100
	for i := 0; i < 105; i++ {
		profile.AddAuditEntry("test", "actor", "reason", types.StatusActive, types.StatusActive)
	}
	require.Len(t, profile.AuditLog, 100, "audit log should be capped at 100 entries")
}

func TestMsgUpsertProfile_GetSigners(t *testing.T) {
	authority := sdk.AccAddress("authority___________")
	msg := types.MsgUpsertProfile{
		Authority: authority.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, authority, signers[0])
}

func TestMsgSetSanction_GetSigners(t *testing.T) {
	authority := sdk.AccAddress("authority___________")
	msg := types.MsgSetSanction{
		Authority: authority.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, authority, signers[0])
}

func TestNewMsgUpsertProfile(t *testing.T) {
	authority := sdk.AccAddress("authority___________").String()
	profile := types.Profile{
		Address:  sdk.AccAddress("address_____________").String(),
		KYCLevel: types.KYCStandard,
		Risk:     types.RiskLow,
	}

	msg := types.NewMsgUpsertProfile(authority, profile)

	require.Equal(t, authority, msg.Authority)
	require.Equal(t, profile, msg.Profile)
}

func TestNewMsgSetSanction(t *testing.T) {
	authority := sdk.AccAddress("authority___________").String()
	address := sdk.AccAddress("address_____________").String()
	sanction := true
	reason := "OFAC listing"

	msg := types.NewMsgSetSanction(authority, address, sanction, reason)

	require.Equal(t, authority, msg.Authority)
	require.Equal(t, address, msg.Address)
	require.Equal(t, sanction, msg.Sanction)
	require.Equal(t, reason, msg.Reason)
}

func TestBlockedJurisdictions(t *testing.T) {
	blockedCodes := []string{"KP", "IR", "SY", "CU", "RU"}
	for _, code := range blockedCodes {
		require.True(t, types.BlockedJurisdictions[code], "expected %s to be blocked", code)
	}

	require.False(t, types.BlockedJurisdictions["US"])
	require.False(t, types.BlockedJurisdictions["GB"])
	require.False(t, types.BlockedJurisdictions["DE"])
}

func TestHighRiskJurisdictions(t *testing.T) {
	highRiskCodes := []string{"AF", "BY", "MM", "VE", "YE"}
	for _, code := range highRiskCodes {
		require.True(t, types.HighRiskJurisdictions[code], "expected %s to be high risk", code)
	}

	require.False(t, types.HighRiskJurisdictions["US"])
	require.False(t, types.HighRiskJurisdictions["GB"])
	require.False(t, types.HighRiskJurisdictions["DE"])
}
