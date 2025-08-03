package types

import (
	"cosmossdk.io/errors"
)

// DONTCOVER

// x/stst module sentinel errors
var (
	ErrInvalidSigner                = errors.Register(ModuleName, 1100, "invalid signer")
	ErrUnauthorized                 = errors.Register(ModuleName, 1101, "unauthorized")
	ErrInvalidTokenDenom            = errors.Register(ModuleName, 1102, "invalid token denomination")
	ErrInsufficientStake            = errors.Register(ModuleName, 1103, "insufficient stake amount")
	ErrInsufficientFunds            = errors.Register(ModuleName, 1104, "insufficient funds")
	ErrValidatorNotFound            = errors.Register(ModuleName, 1105, "validator not found")
	ErrDelegationNotFound           = errors.Register(ModuleName, 1106, "delegation not found")
	ErrVestingScheduleNotFound      = errors.Register(ModuleName, 1107, "vesting schedule not found")
	ErrInvalidVestingSchedule       = errors.Register(ModuleName, 1108, "invalid vesting schedule")
	ErrInvalidProposal              = errors.Register(ModuleName, 1109, "invalid proposal")
	ErrProposalNotFound             = errors.Register(ModuleName, 1110, "proposal not found")
	ErrVotingPeriodEnded            = errors.Register(ModuleName, 1111, "voting period has ended")
	ErrVotingPeriodNotStarted       = errors.Register(ModuleName, 1112, "voting period has not started")
	ErrInvalidVoteOption            = errors.Register(ModuleName, 1113, "invalid vote option")
	ErrAlreadyVoted                 = errors.Register(ModuleName, 1114, "already voted on this proposal")
	ErrInsufficientVotingPower      = errors.Register(ModuleName, 1115, "insufficient voting power")
	ErrInvalidBurnAmount            = errors.Register(ModuleName, 1116, "invalid burn amount")
	ErrTokenSupplyExceeded          = errors.Register(ModuleName, 1117, "token supply would be exceeded")
	ErrVestingNotStarted            = errors.Register(ModuleName, 1118, "vesting has not started")
	ErrNoVestableTokens             = errors.Register(ModuleName, 1119, "no tokens available for vesting")
	ErrInvalidUnstakingPeriod       = errors.Register(ModuleName, 1120, "invalid unstaking period")
	ErrStakingRewardsNotAvailable   = errors.Register(ModuleName, 1121, "staking rewards not available")
	ErrInvalidSlashingAmount        = errors.Register(ModuleName, 1122, "invalid slashing amount")
	ErrFeeBurnDisabled              = errors.Register(ModuleName, 1123, "fee burning is disabled")
	ErrInvalidFeeBurnRate           = errors.Register(ModuleName, 1124, "invalid fee burn rate")
	ErrInvalidAddress               = errors.Register(ModuleName, 1125, "invalid address")
	ErrInvalidAmount                = errors.Register(ModuleName, 1126, "invalid amount")
	ErrInvalidDuration              = errors.Register(ModuleName, 1127, "invalid duration")
	ErrInvalidParams                = errors.Register(ModuleName, 1128, "invalid parameters")
	ErrModuleDisabled               = errors.Register(ModuleName, 1129, "module is disabled")
	ErrOperationNotPermitted        = errors.Register(ModuleName, 1130, "operation not permitted")
	ErrInvalidBlockHeight           = errors.Register(ModuleName, 1131, "invalid block height")
	ErrInvalidTimestamp             = errors.Register(ModuleName, 1132, "invalid timestamp")
	ErrDuplicateEntry               = errors.Register(ModuleName, 1133, "duplicate entry")
	ErrEntryNotFound                = errors.Register(ModuleName, 1134, "entry not found")
	ErrInvalidConfiguration         = errors.Register(ModuleName, 1135, "invalid configuration")
)