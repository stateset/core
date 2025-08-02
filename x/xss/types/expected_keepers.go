package types

import (
	"context"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx context.Context, name string) sdk.ModuleAccountI
	SetModuleAccount(context.Context, sdk.ModuleAccountI)
	
	// Methods for fee collection
	IterateAccounts(ctx context.Context, process func(sdk.AccountI) (stop bool))
	GetParams(ctx context.Context) authtypes.Params
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error
	
	UndelegateCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	
	BurnCoins(ctx context.Context, name string, amt sdk.Coins) error
	MintCoins(ctx context.Context, name string, amt sdk.Coins) error
	
	GetSupply(ctx context.Context, denom string) sdk.Coin
	GetPaginatedTotalSupply(ctx context.Context, pagination *PageRequest) (sdk.Coins, *PageResponse, error)
	
	// Additional methods for module accounts
	GetModuleBalance(ctx context.Context, moduleName string, denom string) sdk.Coin
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	
	SetDenomMetaData(ctx context.Context, denomMetaData banktypes.Metadata)
	GetDenomMetaData(ctx context.Context, denom string) (banktypes.Metadata, bool)
}

// StakingKeeper defines the expected staking keeper for hooks
type StakingKeeper interface {
	// Validator operations
	GetValidator(ctx context.Context, addr sdk.ValAddress) (validator Validator, found bool)
	GetAllValidators(ctx context.Context) (validators []Validator)
	GetValidatorsByPower(ctx context.Context) []Validator
	GetValidatorByConsAddr(ctx context.Context, consAddr sdk.ConsAddress) (validator Validator, found bool)
	
	// Delegation operations
	GetDelegation(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation Delegation, found bool)
	GetAllDelegatorDelegations(ctx context.Context, delegator sdk.AccAddress) []Delegation
	GetValidatorDelegations(ctx context.Context, valAddr sdk.ValAddress) []Delegation
	
	// Unbonding operations
	GetUnbondingDelegation(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (ubd UnbondingDelegation, found bool)
	GetAllUnbondingDelegations(ctx context.Context, delegator sdk.AccAddress) []UnbondingDelegation
	
	// Params and info
	BondDenom(ctx context.Context) string
	GetParams(ctx context.Context) Params
	
	// Hooks
	AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress)
	AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress)
	AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress)
	AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress)
	AfterDelegationModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
	BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, fraction math.LegacyDec)
	BeforeDelegationCreated(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
	BeforeDelegationSharesModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
	BeforeDelegationRemoved(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
	BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress)
}

// SlashingKeeper defines the expected slashing keeper for validator signing info
type SlashingKeeper interface {
	GetValidatorSigningInfo(ctx context.Context, address sdk.ConsAddress) (info ValidatorSigningInfo, found bool)
	SetValidatorSigningInfo(ctx context.Context, address sdk.ConsAddress, info ValidatorSigningInfo)
	
	IsTombstoned(ctx context.Context, consAddr sdk.ConsAddress) bool
	HasValidatorSigningInfo(ctx context.Context, consAddr sdk.ConsAddress) bool
	
	JailUntil(ctx context.Context, consAddr sdk.ConsAddress, jailTime time.Time)
	Slash(ctx context.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor math.LegacyDec)
	SlashWithInfractionReason(ctx context.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor math.LegacyDec, infraction Infraction)
	
	GetParams(ctx context.Context) (params SlashingParams)
	SignedBlocksWindow(ctx context.Context) (res int64)
	MinSignedPerWindow(ctx context.Context) (res math.LegacyDec)
	DowntimeJailDuration(ctx context.Context) (res time.Duration)
	SlashFractionDoubleSign(ctx context.Context) (res math.LegacyDec)
	SlashFractionDowntime(ctx context.Context) (res math.LegacyDec)
}

// DistrKeeper defines the expected distribution keeper
type DistrKeeper interface {
	FundCommunityPool(ctx context.Context, amount sdk.Coins, sender sdk.AccAddress) error
	GetFeePool(ctx context.Context) (feePool FeePool)
	SetFeePool(ctx context.Context, feePool FeePool)
	
	AllocateTokens(ctx context.Context, sumPreviousPower int64, totalPreviousPower int64, consAddr sdk.ConsAddress, tokens sdk.DecCoins)
	AllocateTokensToValidator(ctx context.Context, val ValidatorI, tokens sdk.DecCoins)
	
	GetValidatorOutstandingRewards(ctx context.Context, val sdk.ValAddress) (rewards ValidatorOutstandingRewards)
	SetValidatorOutstandingRewards(ctx context.Context, val sdk.ValAddress, rewards ValidatorOutstandingRewards)
	GetValidatorAccumulatedCommission(ctx context.Context, val sdk.ValAddress) (commission ValidatorAccumulatedCommission)
	SetValidatorAccumulatedCommission(ctx context.Context, val sdk.ValAddress, commission ValidatorAccumulatedCommission)
	
	GetDelegatorWithdrawAddr(ctx context.Context, delAddr sdk.AccAddress) sdk.AccAddress
	WithdrawDelegationRewards(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error)
	WithdrawValidatorCommission(ctx context.Context, valAddr sdk.ValAddress) (sdk.Coins, error)
	
	GetParams(ctx context.Context) DistrParams
	SetParams(ctx context.Context, params DistrParams)
}

// GovKeeper defines the expected governance keeper for proposal handling
type GovKeeper interface {
	GetProposal(ctx context.Context, proposalID uint64) (proposal Proposal, found bool)
	SetProposal(ctx context.Context, proposal Proposal)
	GetDeposit(ctx context.Context, proposalID uint64, depositorAddr sdk.AccAddress) (deposit Deposit, found bool)
	SetDeposit(ctx context.Context, deposit Deposit)
	GetVote(ctx context.Context, proposalID uint64, voterAddr sdk.AccAddress) (vote Vote, found bool)
	SetVote(ctx context.Context, vote Vote)
	GetProposalStatus(ctx context.Context, proposalID uint64) ProposalStatus
	SetProposalStatus(ctx context.Context, proposalID uint64, status ProposalStatus)
}

// EvidenceKeeper defines the expected evidence keeper for handling validator misbehavior
type EvidenceKeeper interface {
	HandleEquivocationEvidence(ctx context.Context, evidence Equivocation)
}

// Placeholder types (these would normally be imported from respective modules)
type (
	PageRequest  = struct{}
	PageResponse = struct{}
	
	Validator               = interface{}
	ValidatorI              = interface{}
	Delegation              = interface{}
	UnbondingDelegation     = interface{}
	Params                  = interface{}
	ValidatorSigningInfo    = interface{}
	SlashingParams          = interface{}
	Infraction              = interface{}
	FeePool                 = interface{}
	ValidatorOutstandingRewards = interface{}
	ValidatorAccumulatedCommission = interface{}
	DistrParams             = interface{}
	Proposal                = interface{}
	Deposit                 = interface{}
	Vote                    = interface{}
	ProposalStatus          = interface{}
	Equivocation            = interface{}
)