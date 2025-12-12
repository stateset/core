package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
