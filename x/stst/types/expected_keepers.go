package types

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	// Methods imported from account should be defined here
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx context.Context, name string) authtypes.ModuleAccountI
	SetModuleAccount(context.Context, authtypes.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx context.Context) []BankBalance
	GetSupply(ctx context.Context, denom string) sdk.Coin
	GetPaginatedTotalSupply(ctx context.Context, pagination *QueryPaginatedTotalSupplyRequest) (*QueryPaginatedTotalSupplyResponse, error)
	IterateAccountBalances(ctx context.Context, addr sdk.AccAddress, fn func(coin sdk.Coin) (stop bool))
	IterateAllBalances(ctx context.Context, fn func(address sdk.AccAddress, coin sdk.Coin) (stop bool))
	SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	DelegateCoins(ctx context.Context, delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error
	UndelegateCoins(ctx context.Context, moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error
}

// StakingKeeper defines the expected staking keeper used for simulations (noalias)
type StakingKeeper interface {
	GetValidator(ctx context.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	GetValidators(ctx context.Context, maxRetrieve uint32) (validators []stakingtypes.Validator)
	GetDelegation(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingtypes.Delegation, found bool)
	GetAllValidators(ctx context.Context) (validators []stakingtypes.Validator)
	GetAllDelegations(ctx context.Context) (delegations []stakingtypes.Delegation)
	Delegate(ctx context.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc stakingtypes.BondStatus, validator stakingtypes.Validator, subtractAccount bool) (newShares sdk.Dec, err error)
	ValidateUnbondAmount(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, amt sdk.Int) (shares sdk.Dec, err error)
	Undelegate(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, sharesAmount sdk.Dec) (time.Time, error)
	GetDelegatorDelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) (delegations []stakingtypes.Delegation)
	GetValidatorDelegations(ctx context.Context, valAddr sdk.ValAddress) (delegations []stakingtypes.Delegation)
	GetBondedValidatorsByPower(ctx context.Context) []stakingtypes.Validator
	TotalBondedTokens(ctx context.Context) sdk.Int
	BondDenom(ctx context.Context) string
}

// DistributionKeeper defines the expected distribution keeper used for simulations (noalias)
type DistributionKeeper interface {
	WithdrawDelegationRewards(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error)
	GetDelegatorWithdrawAddr(ctx context.Context, delAddr sdk.AccAddress) sdk.AccAddress
}

// Additional types for bank keeper compatibility
type BankBalance struct {
	Address string
	Coins   sdk.Coins
}

type QueryPaginatedTotalSupplyRequest struct {
	// Pagination defines an optional pagination for the request.
	Pagination *PageRequest
}

type QueryPaginatedTotalSupplyResponse struct {
	// Supply is the supply of the paginated coins.
	Supply sdk.Coins
	// Pagination defines the pagination in the response.
	Pagination *PageResponse
}

type PageRequest struct {
	Key        []byte
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

type PageResponse struct {
	NextKey []byte
	Total   uint64
}