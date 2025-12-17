package ante

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	"github.com/stateset/core/x/feemarket/keeper"
	"github.com/stateset/core/x/feemarket/types"
)

// FeeMarketDecorator implements dynamic fee calculation based on the fee market module.
// It validates that the transaction fees meet the minimum base fee requirements.
type FeeMarketDecorator struct {
	feemarketKeeper keeper.Keeper
	accountKeeper   AccountKeeper
	bankKeeper      BankKeeper
	oracleKeeper    OracleKeeper
}

// AccountKeeper defines the expected account keeper interface.
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) sdk.AccountI
}

// BankKeeper defines the expected bank keeper interface.
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

// OracleKeeper defines the expected oracle keeper interface.
type OracleKeeper interface {
	GetPriceDec(ctx context.Context, denom string) (sdkmath.LegacyDec, error)
}

// NewFeeMarketDecorator creates a new fee market ante handler decorator.
func NewFeeMarketDecorator(fmk keeper.Keeper, ak AccountKeeper, bk BankKeeper, ok OracleKeeper) FeeMarketDecorator {
	return FeeMarketDecorator{
		feemarketKeeper: fmk,
		accountKeeper:   ak,
		bankKeeper:      bk,
		oracleKeeper:    ok,
	}
}

// AnteHandle validates transaction fees against the current base fee.
func (fmd FeeMarketDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// Skip fee validation in simulation mode
	if simulate {
		return next(ctx, tx, simulate)
	}

	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "tx must implement FeeTx interface")
	}

	// Get fee market params
	params := fmd.feemarketKeeper.GetParams(ctx)

	// If fee market is not enabled, skip dynamic fee validation
	if !params.Enabled {
		return next(ctx, tx, simulate)
	}

	// Get current base fee (in native denom)
	baseFee := fmd.feemarketKeeper.GetBaseFee(ctx)

	// Get transaction fee and gas limit
	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()

	// Calculate minimum required fee based on base fee (in native denom)
	minRequiredFee := baseFee.MulInt64(int64(gas))

	// Validate fee against minimum
	if err := fmd.validateFee(ctx, feeCoins, minRequiredFee); err != nil {
		return ctx, err
	}

	// Emit fee market event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeFeeValidation,
			sdk.NewAttribute(types.AttributeKeyBaseFee, baseFee.String()),
			sdk.NewAttribute(types.AttributeKeyGasLimit, fmt.Sprintf("%d", gas)),
			sdk.NewAttribute(types.AttributeKeyMinRequiredFee, minRequiredFee.String()),
			sdk.NewAttribute(types.AttributeKeyActualFee, feeCoins.String()),
		),
	)

	return next(ctx, tx, simulate)
}

// validateFee checks if the provided fee meets the minimum requirement.
func (fmd FeeMarketDecorator) validateFee(ctx sdk.Context, feeCoins sdk.Coins, minRequiredFee sdkmath.LegacyDec) error {
	// If no fees provided, return error
	if feeCoins.IsZero() {
		return errorsmod.Wrapf(
			types.ErrInsufficientFee,
			"transaction fee cannot be zero; minimum required: %s (native)",
			minRequiredFee.String(),
		)
	}

	// Get the fee denom (assuming first coin is the fee denom)
	feeCoin := feeCoins[0]
	feeAmount := sdkmath.LegacyNewDecFromInt(feeCoin.Amount)

	// Native Gas Abstraction: Handle ssUSD or other whitelisted stablecoins.
	// Supported fee denoms: ssusd (legacy: ussUSD)
	isAllowedDenom := feeCoin.Denom == "ussUSD" || feeCoin.Denom == "ssusd"

	if isAllowedDenom {
		// Get price of the fee token (ssUSD should be ~$1.00)
		feeTokenPrice, err := fmd.oracleKeeper.GetPriceDec(ctx, feeCoin.Denom)
		if err != nil {
			return errorsmod.Wrapf(types.ErrInsufficientFee, "failed to get %s price for fee validation: %s", feeCoin.Denom, err)
		}

		// Get price of Native Token
		// Try "stake" first, then fallback to "uatom"
		nativeDenom := "stake"
		nativePrice, err := fmd.oracleKeeper.GetPriceDec(ctx, nativeDenom)
		if err != nil {
			// Fallback to uatom if native price not found
			nativePrice, err = fmd.oracleKeeper.GetPriceDec(ctx, "uatom")
			if err != nil {
				return errorsmod.Wrapf(types.ErrInsufficientFee, "failed to get native token price for fee validation: %s", err)
			}
		}

		// Calculate required fee token amount
		// RequiredNative * NativePrice = RequiredUSD
		// RequiredUSD / FeeTokenPrice = RequiredFeeToken
		requiredUSD := minRequiredFee.Mul(nativePrice)
		requiredFeeToken := requiredUSD.Quo(feeTokenPrice)

		if feeAmount.LT(requiredFeeToken) {
			return errorsmod.Wrapf(
				types.ErrInsufficientFee,
				"insufficient fee; got: %s %s, required: %s %s",
				feeAmount.String(), feeCoin.Denom,
				requiredFeeToken.String(), feeCoin.Denom,
			)
		}
		return nil
	}

	// Standard Native Fee Validation
	// Check if fee meets minimum requirement directly
	if feeAmount.LT(minRequiredFee) {
		return errorsmod.Wrapf(
			types.ErrInsufficientFee,
			"insufficient fee; got: %s, required minimum: %s",
			feeAmount.String(),
			minRequiredFee.String(),
		)
	}

	return nil
}

// DynamicFeeChecker returns a fee checker function that uses the fee market base fee.
// This can be used with the standard Cosmos SDK fee decorator.
func DynamicFeeChecker(fmk keeper.Keeper, oracleK OracleKeeper) authante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		feeTx, isFeeTx := tx.(sdk.FeeTx)
		if !isFeeTx {
			return nil, 0, errorsmod.Wrap(sdkerrors.ErrTxDecode, "tx must implement FeeTx interface")
		}

		// Get fee market params
		params := fmk.GetParams(ctx)

		// If fee market is not enabled, use standard validation
		if !params.Enabled {
			return feeTx.GetFee(), int64(feeTx.GetGas()), nil
		}

		// Get current base fee
		baseFee := fmk.GetBaseFee(ctx)

		// Get transaction gas limit
		gas := feeTx.GetGas()

		// Calculate minimum required fee
		minRequiredFee := baseFee.MulInt64(int64(gas))

		// Get actual fee
		feeCoins := feeTx.GetFee()

		// Validate
		if feeCoins.IsZero() {
			return nil, 0, errorsmod.Wrapf(
				types.ErrInsufficientFee,
				"transaction fee cannot be zero; minimum required: %s",
				minRequiredFee.String(),
			)
		}

		// Get fee amount
		feeCoin := feeCoins[0]
		feeAmount := sdkmath.LegacyNewDecFromInt(feeCoin.Amount)

		// Logic duplication for DynamicFeeChecker (simplified)
		if feeCoin.Denom == "ssusd" {
			ssusdPrice, err := oracleK.GetPriceDec(ctx, "ssusd")
			if err != nil {
				return nil, 0, err
			}
			nativePrice, err := oracleK.GetPriceDec(ctx, "stake") // Assumed native
			if err != nil {
				// Fallback
				nativePrice, err = oracleK.GetPriceDec(ctx, "uatom")
				if err != nil {
					return nil, 0, err
				}
			}
			requiredSSUSD := minRequiredFee.Mul(nativePrice).Quo(ssusdPrice)
			if feeAmount.LT(requiredSSUSD) {
				return nil, 0, errorsmod.Wrapf(types.ErrInsufficientFee, "insufficient ssusd fee")
			}
			return feeCoins, int64(gas), nil
		}

		// Check minimum
		if feeAmount.LT(minRequiredFee) {
			return nil, 0, errorsmod.Wrapf(
				types.ErrInsufficientFee,
				"insufficient fee; got: %s, required minimum: %s",
				feeAmount.String(),
				minRequiredFee.String(),
			)
		}

		return feeCoins, int64(gas), nil
	}
}

// FeeMarketCheckTxFeeWithMinGasPrices implements a fee checker that uses the fee market
// in combination with validator minimum gas prices.
func FeeMarketCheckTxFeeWithMinGasPrices(fmk keeper.Keeper, oracleK OracleKeeper) authante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		feeTx, isFeeTx := tx.(sdk.FeeTx)
		if !isFeeTx {
			return nil, 0, errorsmod.Wrap(sdkerrors.ErrTxDecode, "tx must implement FeeTx interface")
		}

		feeCoins := feeTx.GetFee()
		gas := feeTx.GetGas()

		// Get fee market params
		params := fmk.GetParams(ctx)

		// Calculate minimum fee from fee market
		var minFee sdkmath.LegacyDec
		if params.Enabled {
			baseFee := fmk.GetBaseFee(ctx)
			minFee = baseFee.MulInt64(int64(gas))
		} else {
			minFee = params.MinBaseFee.MulInt64(int64(gas))
		}

		// Validate fee early to avoid denom lookups on empty fees.
		if feeCoins.IsZero() {
			return nil, 0, errorsmod.Wrapf(
				types.ErrInsufficientFee,
				"transaction fee cannot be zero; minimum required: %s",
				minFee.String(),
			)
		}

		feeCoin := feeCoins[0]

		// Special handling for ssusd
		if feeCoin.Denom == "ssusd" {
			ssusdPrice, err := oracleK.GetPriceDec(ctx, "ssusd")
			if err != nil {
				return nil, 0, err
			}
			nativePrice, err := oracleK.GetPriceDec(ctx, "stake")
			if err != nil {
				nativePrice, err = oracleK.GetPriceDec(ctx, "uatom")
				if err != nil {
					return nil, 0, err
				}
			}

			// Convert minFee (native) to ssusd
			requiredSSUSD := minFee.Mul(nativePrice).Quo(ssusdPrice)

			feeAmount := sdkmath.LegacyNewDecFromInt(feeCoin.Amount)
			if feeAmount.LT(requiredSSUSD) {
				return nil, 0, errorsmod.Wrapf(types.ErrInsufficientFee, "insufficient ssusd fee")
			}
			return feeCoins, int64(gas), nil
		}

		// Also check against validator's minimum gas prices (if set)
		minGasPrices := ctx.MinGasPrices()
		if !minGasPrices.IsZero() {
			// Calculate minimum fee from validator settings
			validatorMinFee := minGasPrices.AmountOf(feeCoin.Denom).MulInt64(int64(gas))

			// Use the higher of the two
			if validatorMinFee.GT(minFee) {
				minFee = validatorMinFee
			}
		}

		feeAmount := sdkmath.LegacyNewDecFromInt(feeCoin.Amount)

		if feeAmount.LT(minFee) {
			return nil, 0, errorsmod.Wrapf(
				types.ErrInsufficientFee,
				"insufficient fee; got: %s, required minimum: %s",
				feeAmount.String(),
				minFee.String(),
			)
		}

		return feeCoins, int64(gas), nil
	}
}
