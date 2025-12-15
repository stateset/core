package types

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamSubspace defines the expected parameter subspace interface.
type ParamSubspace = paramtypes.Subspace

// BankKeeper defines the expected bank keeper subset.
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx context.Context, sender sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipient sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	GetSupply(ctx context.Context, denom string) sdk.Coin
}

// AccountKeeper describes the subset used to ensure module accounts are set up.
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	SetModuleAccount(ctx context.Context, macc sdk.ModuleAccountI)
}

// OracleKeeper exposes price lookups.
type OracleKeeper interface {
	GetPriceDec(ctx context.Context, denom string) (sdkmath.LegacyDec, error)
	// GetPriceDecSafe returns the current price with staleness checks enforced.
	GetPriceDecSafe(ctx context.Context, denom string) (sdkmath.LegacyDec, error)
}

// ComplianceKeeper ensures addresses are cleared.
type ComplianceKeeper interface {
	AssertCompliant(ctx context.Context, addr sdk.AccAddress) error
}
