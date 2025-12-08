package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vault tracks collateralized debt positions for stablecoin issuance.
type Vault struct {
	Id              uint64      `json:"id" yaml:"id"`
	Owner           string      `json:"owner" yaml:"owner"`
	Collateral      sdk.Coin    `json:"collateral" yaml:"collateral"`
	CollateralDenom string      `json:"collateral_denom" yaml:"collateral_denom"`
	Debt            sdkmath.Int `json:"debt" yaml:"debt"`
	LastAccrued     int64       `json:"last_accrued" yaml:"last_accrued"`
}

func (v Vault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(v.Owner); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if !v.Collateral.IsValid() || v.Collateral.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "collateral must be positive")
	}
	if !v.Debt.IsZero() && v.Debt.IsNegative() {
		return errorsmod.Wrap(ErrInvalidAmount, "debt cannot be negative")
	}
	if v.CollateralDenom == "" {
		return errorsmod.Wrap(ErrInvalidVault, "collateral denom required")
	}
	return nil
}
