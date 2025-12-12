package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
)

func (cp CollateralParam) ValidateBasic() error {
	if cp.Denom == "" {
		return errorsmod.Wrap(ErrInvalidVault, "collateral denom required")
	}
	if cp.LiquidationRatio.LT(sdkmath.LegacyMustNewDecFromStr("1.0")) {
		return errorsmod.Wrap(ErrInvalidVault, "liquidation ratio must be >= 1.0")
	}
	if !cp.DebtLimit.IsPositive() {
		return errorsmod.Wrap(ErrInvalidVault, "debt limit must be positive")
	}
	return nil
}

func DefaultParams() Params {
	return Params{
		VaultMintingEnabled: false,
		CollateralParams: []CollateralParam{
			{
				Denom:            "stst",
				LiquidationRatio: sdkmath.LegacyMustNewDecFromStr("1.5"),
				StabilityFee:     sdkmath.LegacyMustNewDecFromStr("0.02"),
				DebtLimit:        sdkmath.NewInt(1_000_000_000_000),
				Active:           true,
			},
		},
	}
}

func (p Params) ValidateBasic() error {
	for _, cp := range p.CollateralParams {
		if err := cp.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}

func (p Params) GetCollateralParam(denom string) (CollateralParam, bool) {
	for _, cp := range p.CollateralParams {
		if cp.Denom == denom {
			return cp, true
		}
	}
	return CollateralParam{}, false
}
