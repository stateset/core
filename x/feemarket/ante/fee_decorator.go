package ante

import (
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

// NewFeeMarketDecorator creates a new fee market ante handler decorator.
func NewFeeMarketDecorator(fmk keeper.Keeper, ak AccountKeeper, bk BankKeeper) FeeMarketDecorator {
	return FeeMarketDecorator{
		feemarketKeeper: fmk,
		accountKeeper:   ak,
		bankKeeper:      bk,
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

	// Get current base fee
	baseFee := fmd.feemarketKeeper.GetBaseFee(ctx)

	// Get transaction fee and gas limit
	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()

	// Calculate minimum required fee based on base fee
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
			"transaction fee cannot be zero; minimum required: %s",
			minRequiredFee.String(),
		)
	}

	// Get the fee denom (assuming first coin is the fee denom)
	// In a multi-denom scenario, you'd need more sophisticated logic
	feeCoin := feeCoins[0]

	// Convert fee amount to LegacyDec for comparison
	feeAmount := sdkmath.LegacyNewDecFromInt(feeCoin.Amount)

	// Check if fee meets minimum requirement
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
func DynamicFeeChecker(fmk keeper.Keeper) authante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
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
func FeeMarketCheckTxFeeWithMinGasPrices(fmk keeper.Keeper) authante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
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

		// Also check against validator's minimum gas prices (if set)
		minGasPrices := ctx.MinGasPrices()
		if !minGasPrices.IsZero() {
			// Calculate minimum fee from validator settings
			validatorMinFee := minGasPrices.AmountOf(feeCoins[0].Denom).MulInt64(int64(gas))

			// Use the higher of the two
			if validatorMinFee.GT(minFee) {
				minFee = validatorMinFee
			}
		}

		// Validate fee
		if feeCoins.IsZero() {
			return nil, 0, errorsmod.Wrapf(
				types.ErrInsufficientFee,
				"transaction fee cannot be zero; minimum required: %s",
				minFee.String(),
			)
		}

		feeCoin := feeCoins[0]
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
