package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrUnsupportedCollateral   = errorsmod.Register(ModuleName, 1, "unsupported collateral")
	ErrVaultNotFound           = errorsmod.Register(ModuleName, 2, "vault not found")
	ErrInvalidVault            = errorsmod.Register(ModuleName, 3, "invalid vault parameters")
	ErrUnderCollateralized     = errorsmod.Register(ModuleName, 4, "vault below minimum collateral ratio")
	ErrUnauthorized            = errorsmod.Register(ModuleName, 5, "unauthorized")
	ErrInvalidAmount           = errorsmod.Register(ModuleName, 6, "invalid amount")
	ErrPriceNotFound           = errorsmod.Register(ModuleName, 7, "price not available")
	ErrPriceStale              = errorsmod.Register(ModuleName, 8, "price is stale")
	ErrLiquidationSurge        = errorsmod.Register(ModuleName, 9, "liquidation surge protection triggered")
	ErrVaultHealthy            = errorsmod.Register(ModuleName, 10, "vault is healthy, cannot liquidate")
	ErrInsufficientCollateral  = errorsmod.Register(ModuleName, 11, "insufficient collateral")
	ErrDebtLimitExceeded       = errorsmod.Register(ModuleName, 12, "debt limit exceeded")
	ErrCollateralInactive      = errorsmod.Register(ModuleName, 13, "collateral type is inactive")
	ErrComplianceCheckFailed   = errorsmod.Register(ModuleName, 14, "compliance check failed")
	ErrModuleAccountNotFound   = errorsmod.Register(ModuleName, 15, "module account not found")
	ErrInvalidCollateralParams = errorsmod.Register(ModuleName, 16, "invalid collateral parameters")
)
