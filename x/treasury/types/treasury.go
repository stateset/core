package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TreasuryParams defines the parameters for the treasury module
type TreasuryParams struct {
	// MinTimelockDuration is the minimum timelock for spend proposals (e.g., 24 hours)
	MinTimelockDuration time.Duration `json:"min_timelock_duration" yaml:"min_timelock_duration"`
	// MaxTimelockDuration is the maximum timelock for spend proposals (e.g., 30 days)
	MaxTimelockDuration time.Duration `json:"max_timelock_duration" yaml:"max_timelock_duration"`
	// ProposalExpiryDuration is how long a proposal remains valid after timelock expires
	ProposalExpiryDuration time.Duration `json:"proposal_expiry_duration" yaml:"proposal_expiry_duration"`
	// MaxPendingProposals is the maximum number of pending proposals allowed
	MaxPendingProposals uint32 `json:"max_pending_proposals" yaml:"max_pending_proposals"`
	// EmergencyMultisigThreshold for bypassing timelock (requires N of M signers)
	EmergencyMultisigThreshold uint32 `json:"emergency_multisig_threshold" yaml:"emergency_multisig_threshold"`
	// BaseBurnRate percentage of incoming fees to burn (basis points, e.g., 2500 = 25%)
	BaseBurnRate uint32 `json:"base_burn_rate" yaml:"base_burn_rate"`
	// ValidatorRewardRate percentage to validators (basis points)
	ValidatorRewardRate uint32 `json:"validator_reward_rate" yaml:"validator_reward_rate"`
	// CommunityPoolRate percentage to community pool (basis points)
	CommunityPoolRate uint32 `json:"community_pool_rate" yaml:"community_pool_rate"`
}

// DefaultTreasuryParams returns default treasury parameters
func DefaultTreasuryParams() TreasuryParams {
	return TreasuryParams{
		MinTimelockDuration:        24 * time.Hour,        // 24 hours minimum
		MaxTimelockDuration:        30 * 24 * time.Hour,   // 30 days maximum
		ProposalExpiryDuration:     7 * 24 * time.Hour,    // 7 days to execute after timelock
		MaxPendingProposals:        100,
		EmergencyMultisigThreshold: 3,                     // 3 of 5 for emergency
		BaseBurnRate:               2500,                  // 25% burned
		ValidatorRewardRate:        5000,                  // 50% to validators
		CommunityPoolRate:          2500,                  // 25% to community pool
	}
}

// Validate validates treasury params
func (p TreasuryParams) Validate() error {
	if p.MinTimelockDuration < time.Hour {
		return errorsmod.Wrap(ErrTimelockTooShort, "minimum timelock must be at least 1 hour")
	}
	if p.MaxTimelockDuration < p.MinTimelockDuration {
		return errorsmod.Wrap(ErrTimelockTooShort, "max timelock must be >= min timelock")
	}
	totalRate := p.BaseBurnRate + p.ValidatorRewardRate + p.CommunityPoolRate
	if totalRate != 10000 {
		return errorsmod.Wrapf(ErrInvalidAmount, "fee distribution rates must sum to 10000 (100%%), got %d", totalRate)
	}
	return nil
}

// SpendProposal represents a time-locked treasury spend proposal
type SpendProposal struct {
	Id           uint64    `json:"id" yaml:"id"`
	Proposer     string    `json:"proposer" yaml:"proposer"`
	Recipient    string    `json:"recipient" yaml:"recipient"`
	Amount       sdk.Coins `json:"amount" yaml:"amount"`
	Category     string    `json:"category" yaml:"category"`
	Description  string    `json:"description" yaml:"description"`
	Status       string    `json:"status" yaml:"status"`
	CreatedAt    time.Time `json:"created_at" yaml:"created_at"`
	ExecuteAfter time.Time `json:"execute_after" yaml:"execute_after"`
	ExpiresAt    time.Time `json:"expires_at" yaml:"expires_at"`
	ExecutedAt   time.Time `json:"executed_at,omitempty" yaml:"executed_at,omitempty"`
	TxHash       string    `json:"tx_hash,omitempty" yaml:"tx_hash,omitempty"`
}

// ValidateBasic performs basic validation
func (p SpendProposal) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(p.Proposer); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid proposer address")
	}
	if _, err := sdk.AccAddressFromBech32(p.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if !p.Amount.IsValid() || p.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be valid and non-zero")
	}
	if !IsValidCategory(p.Category) {
		return errorsmod.Wrapf(ErrInvalidCategory, "unknown category: %s", p.Category)
	}
	if len(p.Description) == 0 {
		return errorsmod.Wrap(ErrInvalidAmount, "description cannot be empty")
	}
	if len(p.Description) > 1000 {
		return errorsmod.Wrap(ErrInvalidAmount, "description too long (max 1000 chars)")
	}
	return nil
}

// CanExecute checks if the proposal can be executed now
func (p SpendProposal) CanExecute(now time.Time) error {
	if p.Status != SpendStatusPending {
		return ErrProposalNotPending
	}
	if now.Before(p.ExecuteAfter) {
		return ErrTimelockNotExpired
	}
	if now.After(p.ExpiresAt) {
		return ErrProposalExpired
	}
	return nil
}

// Budget represents a category budget with limits
type Budget struct {
	Category       string         `json:"category" yaml:"category"`
	TotalLimit     sdk.Coins      `json:"total_limit" yaml:"total_limit"`
	PeriodLimit    sdk.Coins      `json:"period_limit" yaml:"period_limit"`
	PeriodSpent    sdk.Coins      `json:"period_spent" yaml:"period_spent"`
	TotalSpent     sdk.Coins      `json:"total_spent" yaml:"total_spent"`
	PeriodStart    time.Time      `json:"period_start" yaml:"period_start"`
	PeriodDuration time.Duration  `json:"period_duration" yaml:"period_duration"`
	Enabled        bool           `json:"enabled" yaml:"enabled"`
}

// ValidateBasic performs basic validation
func (b Budget) ValidateBasic() error {
	if !IsValidCategory(b.Category) {
		return errorsmod.Wrapf(ErrInvalidCategory, "unknown category: %s", b.Category)
	}
	if !b.TotalLimit.IsValid() {
		return errorsmod.Wrap(ErrInvalidAmount, "invalid total limit")
	}
	if !b.PeriodLimit.IsValid() {
		return errorsmod.Wrap(ErrInvalidAmount, "invalid period limit")
	}
	return nil
}

// CanSpend checks if the budget allows spending the amount
func (b Budget) CanSpend(amount sdk.Coins, now time.Time) error {
	if !b.Enabled {
		return errorsmod.Wrapf(ErrBudgetExceeded, "budget for %s is disabled", b.Category)
	}

	// Check period limit
	periodSpent := b.PeriodSpent
	if now.After(b.PeriodStart.Add(b.PeriodDuration)) {
		// New period, reset spent
		periodSpent = sdk.NewCoins()
	}

	newPeriodSpent := periodSpent.Add(amount...)
	for _, coin := range newPeriodSpent {
		limit := b.PeriodLimit.AmountOf(coin.Denom)
		if coin.Amount.GT(limit) {
			return errorsmod.Wrapf(ErrBudgetExceeded, "period limit exceeded for %s in category %s", coin.Denom, b.Category)
		}
	}

	// Check total limit
	newTotalSpent := b.TotalSpent.Add(amount...)
	for _, coin := range newTotalSpent {
		limit := b.TotalLimit.AmountOf(coin.Denom)
		if !limit.IsZero() && coin.Amount.GT(limit) {
			return errorsmod.Wrapf(ErrBudgetExceeded, "total limit exceeded for %s in category %s", coin.Denom, b.Category)
		}
	}

	return nil
}

// Allocation tracks funds allocated to a specific recipient
type Allocation struct {
	Recipient      string    `json:"recipient" yaml:"recipient"`
	TotalAllocated sdk.Coins `json:"total_allocated" yaml:"total_allocated"`
	TotalDisbursed sdk.Coins `json:"total_disbursed" yaml:"total_disbursed"`
	Pending        sdk.Coins `json:"pending" yaml:"pending"`
	LastUpdated    time.Time `json:"last_updated" yaml:"last_updated"`
}

// RevenueRecord tracks incoming revenue to the treasury
type RevenueRecord struct {
	Id        uint64    `json:"id" yaml:"id"`
	Source    string    `json:"source" yaml:"source"`
	Amount    sdk.Coins `json:"amount" yaml:"amount"`
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
	TxHash    string    `json:"tx_hash" yaml:"tx_hash"`
	Metadata  string    `json:"metadata" yaml:"metadata"`
}

// TreasuryStats provides aggregate statistics
type TreasuryStats struct {
	TotalBalance         sdk.Coins              `json:"total_balance" yaml:"total_balance"`
	TotalAllocated       sdk.Coins              `json:"total_allocated" yaml:"total_allocated"`
	TotalDisbursed       sdk.Coins              `json:"total_disbursed" yaml:"total_disbursed"`
	TotalBurned          sdk.Coins              `json:"total_burned" yaml:"total_burned"`
	TotalRevenue         sdk.Coins              `json:"total_revenue" yaml:"total_revenue"`
	PendingProposals     uint32                 `json:"pending_proposals" yaml:"pending_proposals"`
	ExecutedProposals    uint32                 `json:"executed_proposals" yaml:"executed_proposals"`
	CategorySpending     map[string]sdk.Coins   `json:"category_spending" yaml:"category_spending"`
	LastRevenueTimestamp time.Time              `json:"last_revenue_timestamp" yaml:"last_revenue_timestamp"`
}

// IsValidCategory checks if the category is valid
func IsValidCategory(category string) bool {
	switch category {
	case CategoryDevelopment, CategoryMarketing, CategoryOperations,
		CategoryGrants, CategorySecurity, CategoryInfrastructure, CategoryReserve:
		return true
	}
	return false
}

// ValidCategories returns all valid budget categories
func ValidCategories() []string {
	return []string{
		CategoryDevelopment,
		CategoryMarketing,
		CategoryOperations,
		CategoryGrants,
		CategorySecurity,
		CategoryInfrastructure,
		CategoryReserve,
	}
}

// FeeDistribution represents how fees are distributed
type FeeDistribution struct {
	BurnAmount          sdk.Coins `json:"burn_amount" yaml:"burn_amount"`
	ValidatorAmount     sdk.Coins `json:"validator_amount" yaml:"validator_amount"`
	CommunityPoolAmount sdk.Coins `json:"community_pool_amount" yaml:"community_pool_amount"`
	TreasuryAmount      sdk.Coins `json:"treasury_amount" yaml:"treasury_amount"`
}

// CalculateFeeDistribution calculates how to distribute fees
func CalculateFeeDistribution(fees sdk.Coins, params TreasuryParams) FeeDistribution {
	dist := FeeDistribution{
		BurnAmount:          sdk.NewCoins(),
		ValidatorAmount:     sdk.NewCoins(),
		CommunityPoolAmount: sdk.NewCoins(),
		TreasuryAmount:      sdk.NewCoins(),
	}

	for _, coin := range fees {
		burnAmt := coin.Amount.Mul(sdkmath.NewInt(int64(params.BaseBurnRate))).Quo(sdkmath.NewInt(10000))
		validatorAmt := coin.Amount.Mul(sdkmath.NewInt(int64(params.ValidatorRewardRate))).Quo(sdkmath.NewInt(10000))
		communityAmt := coin.Amount.Mul(sdkmath.NewInt(int64(params.CommunityPoolRate))).Quo(sdkmath.NewInt(10000))

		if burnAmt.IsPositive() {
			dist.BurnAmount = dist.BurnAmount.Add(sdk.NewCoin(coin.Denom, burnAmt))
		}
		if validatorAmt.IsPositive() {
			dist.ValidatorAmount = dist.ValidatorAmount.Add(sdk.NewCoin(coin.Denom, validatorAmt))
		}
		if communityAmt.IsPositive() {
			dist.CommunityPoolAmount = dist.CommunityPoolAmount.Add(sdk.NewCoin(coin.Denom, communityAmt))
		}
	}

	return dist
}
