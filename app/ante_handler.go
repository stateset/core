package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// HandlerOptions are the options required for constructing an enhanced ante handler
type HandlerOptions struct {
	AccountKeeper      authkeeper.AccountKeeper
	BankKeeper         bankkeeper.Keeper
	FeegrantKeeper     feegrantkeeper.Keeper
	GovKeeper          govkeeper.Keeper
	SignModeHandler    sdk.SignModeHandler
	SigGasConsumer     ante.SignatureVerificationGasConsumer
	IBCKeeper          *ibckeeper.Keeper
	ParamsKeeper       paramtypes.Keeper
	MaxTxGasWanted     uint64
	MaxTxSizeBytes     uint64
	MinGasPriceCoins   sdk.DecCoins
}

// NewEnhancedAnteHandler creates a new ante handler with performance optimizations
func NewEnhancedAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}
	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}
	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	
	anteDecorators := []sdk.AnteDecorator{
		// Setup context with gas meter
		ante.NewSetUpContextDecorator(),
		
		// Reject extension options
		ante.NewRejectExtensionOptionsDecorator(),
		
		// Priority-based mempool decorator
		NewPriorityMempoolDecorator(),
		
		// Enhanced validation with caching
		NewEnhancedValidateBasicDecorator(),
		
		// Transaction size limit
		NewTxSizeLimitDecorator(options.MaxTxSizeBytes),
		
		// Gas limit decorator
		NewGasLimitDecorator(options.MaxTxGasWanted),
		
		// Minimum gas price enforcement
		NewMinGasPriceDecorator(options.MinGasPriceCoins),
		
		// Memo size limit
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		
		// Consume gas for transaction size
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		
		// Fee and priority calculation
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		
		// SetPubKey decorator
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		
		// Validate signatures with batch verification
		NewBatchSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler, options.SigGasConsumer),
		
		// Sequence validation
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		
		// IBC ante decorator
		ibcante.NewAnteDecorator(options.IBCKeeper),
		
		// Gov expedited proposals
		NewGovExpeditedProposalsDecorator(options.GovKeeper),
	}
	
	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

// PriorityMempoolDecorator implements transaction prioritization
type PriorityMempoolDecorator struct{}

func NewPriorityMempoolDecorator() PriorityMempoolDecorator {
	return PriorityMempoolDecorator{}
}

func (pmd PriorityMempoolDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "tx must implement FeeTx")
	}
	
	// Calculate priority based on fee amount
	priority := sdk.NewDec(1)
	if !feeTx.GetFee().IsZero() {
		gasPrice := sdk.NewDecFromInt(feeTx.GetFee().AmountOf("state")).Quo(sdk.NewDecFromInt(sdk.NewInt(int64(feeTx.GetGas()))))
		priority = gasPrice
	}
	
	// Set priority in context for mempool ordering
	ctx = ctx.WithPriority(priority.TruncateInt64())
	
	return next(ctx, tx, simulate)
}

// EnhancedValidateBasicDecorator adds caching for validation results
type EnhancedValidateBasicDecorator struct {
	cache map[string]error
}

func NewEnhancedValidateBasicDecorator() EnhancedValidateBasicDecorator {
	return EnhancedValidateBasicDecorator{
		cache: make(map[string]error),
	}
}

func (vbd EnhancedValidateBasicDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// Check cache first
	txHash := string(ctx.TxBytes())
	if err, found := vbd.cache[txHash]; found && !simulate {
		if err != nil {
			return ctx, err
		}
		return next(ctx, tx, simulate)
	}
	
	// Validate transaction
	if err := tx.ValidateBasic(); err != nil {
		vbd.cache[txHash] = err
		return ctx, err
	}
	
	vbd.cache[txHash] = nil
	return next(ctx, tx, simulate)
}

// TxSizeLimitDecorator enforces transaction size limits
type TxSizeLimitDecorator struct {
	maxTxSize uint64
}

func NewTxSizeLimitDecorator(maxTxSize uint64) TxSizeLimitDecorator {
	if maxTxSize == 0 {
		maxTxSize = 1048576 // 1MB default
	}
	return TxSizeLimitDecorator{maxTxSize: maxTxSize}
}

func (tsl TxSizeLimitDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	txSize := uint64(len(ctx.TxBytes()))
	if txSize > tsl.maxTxSize {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrTxTooLarge,
			"transaction size %d exceeds maximum %d bytes", txSize, tsl.maxTxSize)
	}
	return next(ctx, tx, simulate)
}

// GasLimitDecorator enforces maximum gas limits
type GasLimitDecorator struct {
	maxGasWanted uint64
}

func NewGasLimitDecorator(maxGasWanted uint64) GasLimitDecorator {
	if maxGasWanted == 0 {
		maxGasWanted = 10000000 // 10M default
	}
	return GasLimitDecorator{maxGasWanted: maxGasWanted}
}

func (gld GasLimitDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "tx must implement FeeTx")
	}
	
	if feeTx.GetGas() > gld.maxGasWanted {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrOutOfGas,
			"gas wanted %d exceeds maximum %d", feeTx.GetGas(), gld.maxGasWanted)
	}
	
	return next(ctx, tx, simulate)
}

// MinGasPriceDecorator enforces minimum gas prices
type MinGasPriceDecorator struct {
	minGasPrices sdk.DecCoins
}

func NewMinGasPriceDecorator(minGasPrices sdk.DecCoins) MinGasPriceDecorator {
	return MinGasPriceDecorator{minGasPrices: minGasPrices}
}

func (mgp MinGasPriceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// Skip during simulation
	if simulate || ctx.IsCheckTx() {
		return next(ctx, tx, simulate)
	}
	
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "tx must implement FeeTx")
	}
	
	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()
	
	// Check minimum gas prices
	if !mgp.minGasPrices.IsZero() {
		requiredFees := make(sdk.Coins, len(mgp.minGasPrices))
		for i, gp := range mgp.minGasPrices {
			fee := gp.Amount.Mul(sdk.NewDec(int64(gas)))
			requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().TruncateInt())
		}
		
		if !feeCoins.IsAnyGTE(requiredFees) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee,
				"insufficient fees; got: %s required: %s", feeCoins, requiredFees)
		}
	}
	
	return next(ctx, tx, simulate)
}

// BatchSigVerificationDecorator implements batch signature verification
type BatchSigVerificationDecorator struct {
	ak             authkeeper.AccountKeeper
	signModeHandler sdk.SignModeHandler
	sigGasConsumer ante.SignatureVerificationGasConsumer
}

func NewBatchSigVerificationDecorator(ak authkeeper.AccountKeeper, signModeHandler sdk.SignModeHandler, sigGasConsumer ante.SignatureVerificationGasConsumer) BatchSigVerificationDecorator {
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}
	return BatchSigVerificationDecorator{
		ak:             ak,
		signModeHandler: signModeHandler,
		sigGasConsumer: sigGasConsumer,
	}
}

func (svd BatchSigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// Use standard sig verification for now
	// TODO: Implement batch verification for multiple signatures
	decorator := ante.NewSigVerificationDecorator(svd.ak, svd.signModeHandler)
	return decorator.AnteHandle(ctx, tx, simulate, next)
}

// GovExpeditedProposalsDecorator handles expedited governance proposals
type GovExpeditedProposalsDecorator struct {
	govKeeper govkeeper.Keeper
}

func NewGovExpeditedProposalsDecorator(gk govkeeper.Keeper) GovExpeditedProposalsDecorator {
	return GovExpeditedProposalsDecorator{govKeeper: gk}
}

func (g GovExpeditedProposalsDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// Check if this is a governance proposal submission
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		if _, ok := msg.(*govtypes.MsgSubmitProposal); ok {
			// Apply expedited processing logic here if needed
			ctx = ctx.WithPriority(100) // High priority for gov proposals
		}
	}
	
	return next(ctx, tx, simulate)
}