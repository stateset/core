package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stablecointypes "github.com/stateset/core/x/stablecoins/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetMetadata(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
}

// StablecoinsKeeper defines the expected interface needed for stablecoin operations
type StablecoinsKeeper interface {
	GetStablecoin(ctx sdk.Context, denom string) (stablecoin stablecointypes.Stablecoin, found bool)
	GetAllStablecoins(ctx sdk.Context) (list []stablecointypes.Stablecoin)
	IsValidStablecoin(ctx sdk.Context, denom string) bool
	TransferStablecoin(ctx sdk.Context, from, to sdk.AccAddress, amount sdk.Coin) error
	EscrowStablecoin(ctx sdk.Context, from sdk.AccAddress, orderId string, amount sdk.Coin) error
	ReleaseEscrow(ctx sdk.Context, orderId string, to sdk.AccAddress) error
	RefundEscrow(ctx sdk.Context, orderId string, to sdk.AccAddress) error
	GetEscrowBalance(ctx sdk.Context, orderId string) (sdk.Coin, bool)
	ValidateStablecoinPayment(ctx sdk.Context, denom string, amount sdk.Int) error
}