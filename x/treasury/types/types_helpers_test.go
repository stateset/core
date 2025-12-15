package types_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/treasury/types"
)

func TestDefaultTreasuryParams(t *testing.T) {
	params := types.DefaultTreasuryParams()

	require.Equal(t, 24*time.Hour, params.MinTimelockDuration)
	require.Equal(t, 30*24*time.Hour, params.MaxTimelockDuration)
	require.Equal(t, 7*24*time.Hour, params.ProposalExpiryDuration)
	require.Equal(t, uint32(100), params.MaxPendingProposals)
	require.Equal(t, uint32(3), params.EmergencyMultisigThreshold)
	require.Equal(t, uint32(2500), params.BaseBurnRate)
	require.Equal(t, uint32(5000), params.ValidatorRewardRate)
	require.Equal(t, uint32(2500), params.CommunityPoolRate)

	// Validate returns no error for default params
	require.NoError(t, params.Validate())
}

func TestTreasuryParams_Validate(t *testing.T) {
	tests := []struct {
		name      string
		params    types.TreasuryParams
		expectErr bool
	}{
		{
			name:      "valid default params",
			params:    types.DefaultTreasuryParams(),
			expectErr: false,
		},
		{
			name: "min timelock too short",
			params: types.TreasuryParams{
				MinTimelockDuration: 30 * time.Minute, // less than 1 hour
				MaxTimelockDuration: 24 * time.Hour,
				BaseBurnRate:        2500,
				ValidatorRewardRate: 5000,
				CommunityPoolRate:   2500,
			},
			expectErr: true,
		},
		{
			name: "max timelock less than min",
			params: types.TreasuryParams{
				MinTimelockDuration: 24 * time.Hour,
				MaxTimelockDuration: 1 * time.Hour, // less than min
				BaseBurnRate:        2500,
				ValidatorRewardRate: 5000,
				CommunityPoolRate:   2500,
			},
			expectErr: true,
		},
		{
			name: "fee distribution rates dont sum to 10000",
			params: types.TreasuryParams{
				MinTimelockDuration: 24 * time.Hour,
				MaxTimelockDuration: 30 * 24 * time.Hour,
				BaseBurnRate:        2500,
				ValidatorRewardRate: 5000,
				CommunityPoolRate:   3000, // 2500 + 5000 + 3000 = 10500
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgProposeSpend_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()
	validRecipient := sdk.AccAddress("recipient___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgProposeSpend
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgProposeSpend{
				Authority:   validAuthority,
				Recipient:   validRecipient,
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
				Category:    types.CategoryDevelopment,
				Description: "Development funding",
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgProposeSpend{
				Authority:   "invalid",
				Recipient:   validRecipient,
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
				Category:    types.CategoryDevelopment,
				Description: "Development funding",
			},
			expectErr: true,
		},
		{
			name: "invalid recipient",
			msg: &types.MsgProposeSpend{
				Authority:   validAuthority,
				Recipient:   "invalid",
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
				Category:    types.CategoryDevelopment,
				Description: "Development funding",
			},
			expectErr: true,
		},
		{
			name: "zero amount",
			msg: &types.MsgProposeSpend{
				Authority:   validAuthority,
				Recipient:   validRecipient,
				Amount:      sdk.NewCoins(),
				Category:    types.CategoryDevelopment,
				Description: "Development funding",
			},
			expectErr: true,
		},
		{
			name: "invalid category",
			msg: &types.MsgProposeSpend{
				Authority:   validAuthority,
				Recipient:   validRecipient,
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
				Category:    "invalid_category",
				Description: "Development funding",
			},
			expectErr: true,
		},
		{
			name: "empty description",
			msg: &types.MsgProposeSpend{
				Authority:   validAuthority,
				Recipient:   validRecipient,
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
				Category:    types.CategoryDevelopment,
				Description: "",
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

func TestMsgExecuteSpend_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgExecuteSpend
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgExecuteSpend{
				Authority:  validAuthority,
				ProposalID: 1,
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgExecuteSpend{
				Authority:  "invalid",
				ProposalID: 1,
			},
			expectErr: true,
		},
		{
			name: "zero proposal id",
			msg: &types.MsgExecuteSpend{
				Authority:  validAuthority,
				ProposalID: 0,
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

func TestMsgSetBudget_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgSetBudget
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgSetBudget{
				Authority:      validAuthority,
				Category:       types.CategoryMarketing,
				TotalLimit:     sdk.NewCoins(sdk.NewInt64Coin("ssusd", 10000)),
				PeriodLimit:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 1000)),
				PeriodDuration: 24 * time.Hour,
				Enabled:        true,
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgSetBudget{
				Authority: "invalid",
				Category:  types.CategoryMarketing,
			},
			expectErr: true,
		},
		{
			name: "invalid category",
			msg: &types.MsgSetBudget{
				Authority: validAuthority,
				Category:  "invalid_category",
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

func TestSpendProposal_CanExecute(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		proposal  types.SpendProposal
		now       time.Time
		expectErr error
	}{
		{
			name: "can execute",
			proposal: types.SpendProposal{
				Status:       types.SpendStatusPending,
				ExecuteAfter: now.Add(-time.Hour),
				ExpiresAt:    now.Add(time.Hour),
			},
			now:       now,
			expectErr: nil,
		},
		{
			name: "not pending",
			proposal: types.SpendProposal{
				Status:       types.SpendStatusExecuted,
				ExecuteAfter: now.Add(-time.Hour),
				ExpiresAt:    now.Add(time.Hour),
			},
			now:       now,
			expectErr: types.ErrProposalNotPending,
		},
		{
			name: "timelock not expired",
			proposal: types.SpendProposal{
				Status:       types.SpendStatusPending,
				ExecuteAfter: now.Add(time.Hour),
				ExpiresAt:    now.Add(2 * time.Hour),
			},
			now:       now,
			expectErr: types.ErrTimelockNotExpired,
		},
		{
			name: "proposal expired",
			proposal: types.SpendProposal{
				Status:       types.SpendStatusPending,
				ExecuteAfter: now.Add(-2 * time.Hour),
				ExpiresAt:    now.Add(-time.Hour),
			},
			now:       now,
			expectErr: types.ErrProposalExpired,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.proposal.CanExecute(tc.now)
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBudget_CanSpend(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		budget    types.Budget
		amount    sdk.Coins
		expectErr bool
	}{
		{
			name: "can spend within limits",
			budget: types.Budget{
				Category:       types.CategoryDevelopment,
				TotalLimit:     sdk.NewCoins(sdk.NewInt64Coin("ssusd", 10000)),
				TotalSpent:     sdk.NewCoins(sdk.NewInt64Coin("ssusd", 1000)),
				PeriodLimit:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 5000)),
				PeriodSpent:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 500)),
				PeriodStart:    now.Add(-time.Hour),
				PeriodDuration: 24 * time.Hour,
				Enabled:        true,
			},
			amount:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
			expectErr: false,
		},
		{
			name: "budget disabled",
			budget: types.Budget{
				Category: types.CategoryDevelopment,
				Enabled:  false,
			},
			amount:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
			expectErr: true,
		},
		{
			name: "exceeds period limit",
			budget: types.Budget{
				Category:       types.CategoryDevelopment,
				TotalLimit:     sdk.NewCoins(sdk.NewInt64Coin("ssusd", 10000)),
				TotalSpent:     sdk.NewCoins(),
				PeriodLimit:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
				PeriodSpent:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 90)),
				PeriodStart:    now.Add(-time.Hour),
				PeriodDuration: 24 * time.Hour,
				Enabled:        true,
			},
			amount:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 20)), // 90 + 20 > 100
			expectErr: true,
		},
		{
			name: "exceeds total limit",
			budget: types.Budget{
				Category:       types.CategoryDevelopment,
				TotalLimit:     sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100)),
				TotalSpent:     sdk.NewCoins(sdk.NewInt64Coin("ssusd", 90)),
				PeriodLimit:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 1000)),
				PeriodSpent:    sdk.NewCoins(),
				PeriodStart:    now.Add(-time.Hour),
				PeriodDuration: 24 * time.Hour,
				Enabled:        true,
			},
			amount:    sdk.NewCoins(sdk.NewInt64Coin("ssusd", 20)), // 90 + 20 > 100
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.budget.CanSpend(tc.amount, now)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsValidCategory(t *testing.T) {
	validCategories := []string{
		types.CategoryDevelopment,
		types.CategoryMarketing,
		types.CategoryOperations,
		types.CategoryGrants,
		types.CategorySecurity,
		types.CategoryInfrastructure,
		types.CategoryReserve,
	}

	for _, cat := range validCategories {
		require.True(t, types.IsValidCategory(cat), "expected %s to be valid", cat)
	}

	invalidCategories := []string{"", "invalid", "unknown", "test"}
	for _, cat := range invalidCategories {
		require.False(t, types.IsValidCategory(cat), "expected %s to be invalid", cat)
	}
}

func TestValidCategories(t *testing.T) {
	categories := types.ValidCategories()
	require.Len(t, categories, 7)
	require.Contains(t, categories, types.CategoryDevelopment)
	require.Contains(t, categories, types.CategoryMarketing)
	require.Contains(t, categories, types.CategoryOperations)
	require.Contains(t, categories, types.CategoryGrants)
	require.Contains(t, categories, types.CategorySecurity)
	require.Contains(t, categories, types.CategoryInfrastructure)
	require.Contains(t, categories, types.CategoryReserve)
}

func TestCalculateFeeDistribution(t *testing.T) {
	params := types.TreasuryParams{
		BaseBurnRate:        2500, // 25%
		ValidatorRewardRate: 5000, // 50%
		CommunityPoolRate:   2500, // 25%
	}

	fees := sdk.NewCoins(sdk.NewInt64Coin("ssusd", 1000))
	dist := types.CalculateFeeDistribution(fees, params)

	// 25% of 1000 = 250
	require.Equal(t, sdkmath.NewInt(250), dist.BurnAmount.AmountOf("ssusd"))
	// 50% of 1000 = 500
	require.Equal(t, sdkmath.NewInt(500), dist.ValidatorAmount.AmountOf("ssusd"))
	// 25% of 1000 = 250
	require.Equal(t, sdkmath.NewInt(250), dist.CommunityPoolAmount.AmountOf("ssusd"))
}

func TestMsgProposeSpend_GetSigners(t *testing.T) {
	authority := sdk.AccAddress("authority___________")
	msg := types.MsgProposeSpend{
		Authority: authority.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, authority, signers[0])
}

func TestNewMsgProposeSpend(t *testing.T) {
	authority := sdk.AccAddress("authority___________").String()
	recipient := sdk.AccAddress("recipient___________").String()
	amount := sdk.NewCoins(sdk.NewInt64Coin("ssusd", 100))
	category := types.CategoryDevelopment
	description := "test proposal"
	timelockSeconds := uint64(3600)

	msg := types.NewMsgProposeSpend(authority, recipient, amount, category, description, timelockSeconds)

	require.Equal(t, authority, msg.Authority)
	require.Equal(t, recipient, msg.Recipient)
	require.Equal(t, amount, msg.Amount)
	require.Equal(t, category, msg.Category)
	require.Equal(t, description, msg.Description)
	require.Equal(t, timelockSeconds, msg.TimelockSeconds)
}
