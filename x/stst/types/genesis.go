package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesisState returns the default genesis state for the STST module
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:       DefaultParams(),
		StakingState: DefaultStakingState(),
		FeeBurnState: DefaultFeeBurnState(),
		VestingSchedules: []VestingSchedule{
			DefaultTreasuryVestingSchedule(),
			DefaultValidatorRewardsVestingSchedule(),
			DefaultTeamVestingSchedule(),
			DefaultInvestorVestingSchedule(),
			DefaultPartnerEcosystemVestingSchedule(),
			DefaultCommunityVestingSchedule(),
		},
	}
}

// DefaultStakingState returns the default staking state
func DefaultStakingState() StakingState {
	return StakingState{
		TotalStaked:                  sdk.ZeroInt(),
		TotalValidators:              0,
		LastRewardDistributionTime:   0,
	}
}

// DefaultFeeBurnState returns the default fee burn state
func DefaultFeeBurnState() FeeBurnState {
	return FeeBurnState{
		TotalBurned:      sdk.ZeroInt(),
		BurnRateHistory:  []BurnRateEntry{},
	}
}

// Default vesting schedules based on STST tokenomics
func DefaultTreasuryVestingSchedule() VestingSchedule {
	return VestingSchedule{
		Category:        "treasury",
		TotalAmount:     sdk.NewIntFromUint64(300_000_000_000_000), // 300M STST
		CliffDuration:   time.Hour * 24 * 30,                      // 30 days
		VestingDuration: time.Hour * 24 * 365 * 4,                 // 4 years
		StartTime:       0, // Set at genesis
		VestedAmount:    sdk.ZeroInt(),
		Beneficiaries:   []string{}, // To be set by governance
	}
}

func DefaultValidatorRewardsVestingSchedule() VestingSchedule {
	return VestingSchedule{
		Category:        "validator_rewards",
		TotalAmount:     sdk.NewIntFromUint64(250_000_000_000_000), // 250M STST
		CliffDuration:   time.Hour * 24 * 1,                       // 1 day
		VestingDuration: time.Hour * 24 * 365 * 10,                // 10 years linear
		StartTime:       0, // Set at genesis
		VestedAmount:    sdk.ZeroInt(),
		Beneficiaries:   []string{}, // Validators
	}
}

func DefaultTeamVestingSchedule() VestingSchedule {
	return VestingSchedule{
		Category:        "team",
		TotalAmount:     sdk.NewIntFromUint64(150_000_000_000_000), // 150M STST
		CliffDuration:   time.Hour * 24 * 365,                     // 12 months cliff
		VestingDuration: time.Hour * 24 * 365 * 3,                 // 36 months vesting
		StartTime:       0, // Set at genesis
		VestedAmount:    sdk.ZeroInt(),
		Beneficiaries:   []string{}, // Team addresses
	}
}

func DefaultInvestorVestingSchedule() VestingSchedule {
	return VestingSchedule{
		Category:        "investors",
		TotalAmount:     sdk.NewIntFromUint64(150_000_000_000_000), // 150M STST
		CliffDuration:   time.Hour * 24 * 180,                     // 6 months cliff
		VestingDuration: time.Hour * 24 * 365 * 2,                 // 24 months vesting
		StartTime:       0, // Set at genesis
		VestedAmount:    sdk.ZeroInt(),
		Beneficiaries:   []string{}, // Investor addresses
	}
}

func DefaultPartnerEcosystemVestingSchedule() VestingSchedule {
	return VestingSchedule{
		Category:        "partner_ecosystem",
		TotalAmount:     sdk.NewIntFromUint64(100_000_000_000_000), // 100M STST
		CliffDuration:   time.Hour * 24 * 90,                      // 3 months cliff
		VestingDuration: time.Hour * 24 * 365 * 2,                 // 24 months vesting
		StartTime:       0, // Set at genesis
		VestedAmount:    sdk.ZeroInt(),
		Beneficiaries:   []string{}, // Partner addresses
	}
}

func DefaultCommunityVestingSchedule() VestingSchedule {
	return VestingSchedule{
		Category:        "community",
		TotalAmount:     sdk.NewIntFromUint64(50_000_000_000_000), // 50M STST
		CliffDuration:   time.Hour * 0,                           // No cliff
		VestingDuration: time.Hour * 24 * 365,                    // 12 months vesting
		StartTime:       0, // Set at genesis
		VestedAmount:    sdk.ZeroInt(),
		Beneficiaries:   []string{}, // Community addresses
	}
}

// ValidateGenesis validates the genesis state
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	if err := ValidateStakingState(data.StakingState); err != nil {
		return fmt.Errorf("invalid staking state: %w", err)
	}

	if err := ValidateFeeBurnState(data.FeeBurnState); err != nil {
		return fmt.Errorf("invalid fee burn state: %w", err)
	}

	// Validate vesting schedules
	totalVestingAmount := sdk.ZeroInt()
	categories := make(map[string]bool)
	
	for i, schedule := range data.VestingSchedules {
		if err := ValidateVestingSchedule(schedule); err != nil {
			return fmt.Errorf("invalid vesting schedule %d: %w", i, err)
		}
		
		// Check for duplicate categories
		if categories[schedule.Category] {
			return fmt.Errorf("duplicate vesting schedule category: %s", schedule.Category)
		}
		categories[schedule.Category] = true
		
		totalVestingAmount = totalVestingAmount.Add(schedule.TotalAmount)
	}

	// Verify total vesting amount equals total supply
	if !totalVestingAmount.Equal(data.Params.TotalSupply) {
		return fmt.Errorf("total vesting amount (%s) does not equal total supply (%s)", 
			totalVestingAmount, data.Params.TotalSupply)
	}

	return nil
}

// ValidateStakingState validates the staking state
func ValidateStakingState(state StakingState) error {
	if state.TotalStaked.IsNil() || state.TotalStaked.IsNegative() {
		return fmt.Errorf("total staked must be non-negative: %s", state.TotalStaked)
	}

	if state.LastRewardDistributionTime < 0 {
		return fmt.Errorf("last reward distribution time must be non-negative: %d", 
			state.LastRewardDistributionTime)
	}

	return nil
}

// ValidateFeeBurnState validates the fee burn state
func ValidateFeeBurnState(state FeeBurnState) error {
	if state.TotalBurned.IsNil() || state.TotalBurned.IsNegative() {
		return fmt.Errorf("total burned must be non-negative: %s", state.TotalBurned)
	}

	// Validate burn rate history
	var lastBlockHeight int64 = -1
	for i, entry := range state.BurnRateHistory {
		if err := ValidateBurnRateEntry(entry); err != nil {
			return fmt.Errorf("invalid burn rate entry %d: %w", i, err)
		}
		
		// Ensure entries are in chronological order
		if entry.BlockHeight <= lastBlockHeight {
			return fmt.Errorf("burn rate history must be in chronological order: entry %d has block height %d <= %d", 
				i, entry.BlockHeight, lastBlockHeight)
		}
		lastBlockHeight = entry.BlockHeight
	}

	return nil
}

// ValidateBurnRateEntry validates a burn rate entry
func ValidateBurnRateEntry(entry BurnRateEntry) error {
	if entry.BlockHeight < 0 {
		return fmt.Errorf("block height must be non-negative: %d", entry.BlockHeight)
	}

	if entry.BurnRate.IsNil() || entry.BurnRate.IsNegative() || entry.BurnRate.GT(sdk.OneDec()) {
		return fmt.Errorf("burn rate must be between 0 and 1: %s", entry.BurnRate)
	}

	if entry.AmountBurned.IsNil() || entry.AmountBurned.IsNegative() {
		return fmt.Errorf("amount burned must be non-negative: %s", entry.AmountBurned)
	}

	return nil
}

// ValidateVestingSchedule validates a vesting schedule
func ValidateVestingSchedule(schedule VestingSchedule) error {
	if len(schedule.Category) == 0 {
		return fmt.Errorf("category cannot be empty")
	}

	if schedule.TotalAmount.IsNil() || schedule.TotalAmount.IsNegative() || schedule.TotalAmount.IsZero() {
		return fmt.Errorf("total amount must be positive: %s", schedule.TotalAmount)
	}

	if schedule.CliffDuration < 0 {
		return fmt.Errorf("cliff duration must be non-negative: %s", schedule.CliffDuration)
	}

	if schedule.VestingDuration <= 0 {
		return fmt.Errorf("vesting duration must be positive: %s", schedule.VestingDuration)
	}

	if schedule.StartTime < 0 {
		return fmt.Errorf("start time must be non-negative: %d", schedule.StartTime)
	}

	if schedule.VestedAmount.IsNil() || schedule.VestedAmount.IsNegative() {
		return fmt.Errorf("vested amount must be non-negative: %s", schedule.VestedAmount)
	}

	if schedule.VestedAmount.GT(schedule.TotalAmount) {
		return fmt.Errorf("vested amount (%s) cannot exceed total amount (%s)", 
			schedule.VestedAmount, schedule.TotalAmount)
	}

	// Check for duplicate beneficiaries
	beneficiaryMap := make(map[string]bool)
	for _, beneficiary := range schedule.Beneficiaries {
		if len(beneficiary) == 0 {
			return fmt.Errorf("beneficiary address cannot be empty")
		}
		
		if beneficiaryMap[beneficiary] {
			return fmt.Errorf("duplicate beneficiary address: %s", beneficiary)
		}
		beneficiaryMap[beneficiary] = true

		// Validate address format
		if _, err := sdk.AccAddressFromBech32(beneficiary); err != nil {
			return fmt.Errorf("invalid beneficiary address %s: %w", beneficiary, err)
		}
	}

	return nil
}