package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

// FiatTokenFactoryKeeper defines the expected interface for the fiattokenfactory module
type FiatTokenFactoryKeeper interface {
	GetMintingDenom(ctx sdk.Context) (string, bool)
	GetMasterMinter(ctx sdk.Context) (string, bool)
	GetMinters(ctx sdk.Context) []string
	GetMinterController(ctx sdk.Context, minter string) (string, bool)
	GetMinterAllowance(ctx sdk.Context, minter string) (sdk.Int, bool)
}