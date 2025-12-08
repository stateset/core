package app

import (
	math "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
)

// HandlerOptions extends the SDK's AnteHandler options by requiring the IBC keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper     *ibckeeper.Keeper
	CircuitKeeper *circuitkeeper.Keeper
	// WasmConfig        wasmTypes.WasmConfig // Temporarily commented out due to dependency conflicts
}

type MinCommissionDecorator struct{}

func NewMinCommissionDecorator() MinCommissionDecorator {
	return MinCommissionDecorator{}
}

func (MinCommissionDecorator) AnteHandle(
	ctx sdk.Context, tx sdk.Tx,
	simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	msgs := tx.GetMsgs()
	minCommissionRate := math.LegacyNewDecWithPrec(5, 2)
	for _, m := range msgs {
		switch msg := m.(type) {
		case *stakingtypes.MsgCreateValidator:
			c := msg.Commission
			if c.Rate.LT(minCommissionRate) {
				return ctx, sdkerrors.ErrUnauthorized.Wrap("commission can't be lower than 5%")
			}
		case *stakingtypes.MsgEditValidator:
			if msg.CommissionRate == nil {
				continue
			}
			if msg.CommissionRate.LT(minCommissionRate) {
				return ctx, sdkerrors.ErrUnauthorized.Wrap("commission can't be lower than 5%")
			}
		default:
			continue
		}
	}
	return next(ctx, tx, simulate)
}

// NewAnteHandler returns an ante handler configured for the application.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, sdkerrors.ErrLogic.Wrap("sign mode handler is required for ante builder")
	}

	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(),
		NewMinCommissionDecorator(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	// Add circuit breaker security decorator
	if options.CircuitKeeper != nil {
		anteDecorators = append(anteDecorators, NewCombinedSecurityDecorator(options.CircuitKeeper))
	}

	if options.IBCKeeper != nil {
		anteDecorators = append(anteDecorators, ibcante.NewRedundantRelayDecorator(options.IBCKeeper))
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
