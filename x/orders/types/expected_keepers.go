package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	settlementtypes "github.com/stateset/core/x/settlement/types"
)

// BankKeeper defines the expected bank keeper interface.
type BankKeeper interface {
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// ComplianceKeeper defines the expected compliance keeper interface.
type ComplianceKeeper interface {
	AssertCompliant(ctx context.Context, addr sdk.AccAddress) error
}

// SettlementKeeper defines the expected settlement keeper interface.
type SettlementKeeper interface {
	InstantTransfer(ctx sdk.Context, sender, recipient string, amount sdk.Coin, reference, metadata string) (uint64, error)
	CreateEscrow(ctx sdk.Context, sender, recipient string, amount sdk.Coin, reference, metadata string, expirationSeconds int64) (uint64, error)
	ReleaseEscrow(ctx sdk.Context, settlementId uint64, sender sdk.AccAddress) error
	RefundEscrow(ctx sdk.Context, settlementId uint64, recipient sdk.AccAddress, reason string) error
	GetMerchant(ctx sdk.Context, address string) (settlementtypes.MerchantConfig, bool)
}

// AccountKeeper defines the expected account keeper interface.
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}
