package app

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
	circuittypes "github.com/stateset/core/x/circuit/types"
)

// CircuitBreakerDecorator checks circuit breakers and rate limits
type CircuitBreakerDecorator struct {
	circuitKeeper *circuitkeeper.Keeper
}

// NewCircuitBreakerDecorator creates a new circuit breaker decorator
func NewCircuitBreakerDecorator(ck *circuitkeeper.Keeper) CircuitBreakerDecorator {
	return CircuitBreakerDecorator{
		circuitKeeper: ck,
	}
}

func (cbd CircuitBreakerDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (sdk.Context, error) {
	// Skip checks during simulation or if keeper not initialized
	if simulate || cbd.circuitKeeper == nil {
		return next(ctx, tx, simulate)
	}

	// Check global pause first
	if cbd.circuitKeeper.IsGloballyPaused(ctx) {
		return ctx, errorsmod.Wrap(circuittypes.ErrGlobalPause, "system is paused for maintenance")
	}

	// Get signers for rate limiting from signatures
	signers := getSignersFromTx(tx)

	// Check each message
	for _, msg := range tx.GetMsgs() {
		msgType := sdk.MsgTypeURL(msg)
		moduleName := extractModuleName(msgType)

		// Check module circuit breaker
		if cbd.circuitKeeper.IsModuleCircuitOpen(ctx, moduleName) {
			if cbd.circuitKeeper.IsMessageDisabled(ctx, moduleName, msgType) {
				return ctx, errorsmod.Wrapf(
					circuittypes.ErrCircuitOpen,
					"module %s circuit is open for message type %s",
					moduleName,
					msgType,
				)
			}
		}

		// Check rate limits for each signer
		for signer := range signers {
			if err := cbd.circuitKeeper.CheckAllRateLimits(ctx, signer, msgType); err != nil {
				return ctx, errorsmod.Wrapf(
					err,
					"rate limit exceeded for %s on message type %s",
					signer,
					msgType,
				)
			}
		}
	}

	return next(ctx, tx, simulate)
}

// getSignersFromTx extracts signer addresses from transaction signatures
func getSignersFromTx(tx sdk.Tx) map[string]bool {
	signers := make(map[string]bool)

	// Try to get signers from signing tx
	sigTx, ok := tx.(authsigning.Tx)
	if !ok {
		return signers
	}

	// Get signatures
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return signers
	}

	for _, sig := range sigs {
		pubKey := sig.PubKey
		if pubKey != nil {
			addr := sdk.AccAddress(pubKey.Address())
			signers[addr.String()] = true
		}
	}

	return signers
}

// getSignersList returns signers as a slice
func getSignersList(tx sdk.Tx) []string {
	signersMap := getSignersFromTx(tx)
	signers := make([]string, 0, len(signersMap))
	for signer := range signersMap {
		signers = append(signers, signer)
	}
	return signers
}

// extractModuleName extracts the module name from a message type URL
// e.g., "/stateset.stablecoin.v1.MsgMint" -> "stablecoin"
func extractModuleName(msgTypeURL string) string {
	// Remove leading slash
	msgTypeURL = strings.TrimPrefix(msgTypeURL, "/")

	// Split by dots
	parts := strings.Split(msgTypeURL, ".")

	// Expected format: package.module.version.MsgType
	// e.g., stateset.stablecoin.v1.MsgMint
	if len(parts) >= 2 {
		return parts[1] // Return the module name
	}

	return ""
}

// LiquidationSurgeDecorator checks liquidation surge protection
type LiquidationSurgeDecorator struct {
	circuitKeeper *circuitkeeper.Keeper
}

// NewLiquidationSurgeDecorator creates a new liquidation surge decorator
func NewLiquidationSurgeDecorator(ck *circuitkeeper.Keeper) LiquidationSurgeDecorator {
	return LiquidationSurgeDecorator{
		circuitKeeper: ck,
	}
}

func (lsd LiquidationSurgeDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (sdk.Context, error) {
	// Skip during simulation or if keeper not initialized
	if simulate || lsd.circuitKeeper == nil {
		return next(ctx, tx, simulate)
	}

	// Check each message for liquidation operations
	for _, msg := range tx.GetMsgs() {
		msgType := sdk.MsgTypeURL(msg)

		// Check if this is a liquidation message
		if strings.Contains(msgType, "Liquidate") {
			// Get the protection state
			protection := lsd.circuitKeeper.GetLiquidationProtection(ctx)

			// Reset if new block
			if ctx.BlockHeight() > protection.LastResetHeight {
				protection.CurrentBlockLiquidations = 0
				protection.CurrentBlockValue = sdkmath.ZeroInt()
				protection.LastResetHeight = ctx.BlockHeight()
				lsd.circuitKeeper.SetLiquidationProtection(ctx, protection)
			}

			// Check if liquidations are allowed
			if protection.CurrentBlockLiquidations >= protection.MaxLiquidationsPerBlock {
				return ctx, errorsmod.Wrap(
					circuittypes.ErrLiquidationSurge,
					"too many liquidations this block",
				)
			}
		}
	}

	return next(ctx, tx, simulate)
}

// OracleValidationDecorator validates oracle price updates
type OracleValidationDecorator struct {
	circuitKeeper *circuitkeeper.Keeper
}

// NewOracleValidationDecorator creates a new oracle validation decorator
func NewOracleValidationDecorator(ck *circuitkeeper.Keeper) OracleValidationDecorator {
	return OracleValidationDecorator{
		circuitKeeper: ck,
	}
}

func (ovd OracleValidationDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (sdk.Context, error) {
	// This decorator is handled in the oracle keeper itself
	// since it needs access to current prices
	return next(ctx, tx, simulate)
}

// CombinedSecurityDecorator combines all security checks in one decorator
type CombinedSecurityDecorator struct {
	circuitKeeper *circuitkeeper.Keeper
}

// NewCombinedSecurityDecorator creates a combined security decorator
func NewCombinedSecurityDecorator(ck *circuitkeeper.Keeper) CombinedSecurityDecorator {
	return CombinedSecurityDecorator{
		circuitKeeper: ck,
	}
}

func (csd CombinedSecurityDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (sdk.Context, error) {
	// Skip during simulation or if keeper not initialized
	if simulate || csd.circuitKeeper == nil {
		return next(ctx, tx, simulate)
	}

	// 1. Check global pause
	if csd.circuitKeeper.IsGloballyPaused(ctx) {
		return ctx, errorsmod.Wrap(circuittypes.ErrGlobalPause, "system is paused")
	}

	// Get all signers once
	signers := getSignersList(tx)

	// Process each message
	for _, msg := range tx.GetMsgs() {
		msgType := sdk.MsgTypeURL(msg)
		moduleName := extractModuleName(msgType)

		// 2. Check module circuit breaker
		if csd.circuitKeeper.IsModuleCircuitOpen(ctx, moduleName) {
			if csd.circuitKeeper.IsMessageDisabled(ctx, moduleName, msgType) {
				return ctx, errorsmod.Wrapf(
					circuittypes.ErrCircuitOpen,
					"operations disabled for module %s",
					moduleName,
				)
			}
		}

		// 3. Check rate limits
		for _, signer := range signers {
			if err := csd.circuitKeeper.CheckAllRateLimits(ctx, signer, msgType); err != nil {
				return ctx, errorsmod.Wrapf(circuittypes.ErrRateLimitExceeded, "rate limit: %v", err)
			}
		}

		// 4. Check liquidation surge protection
		if strings.Contains(msgType, "Liquidate") {
			protection := csd.circuitKeeper.GetLiquidationProtection(ctx)
			if ctx.BlockHeight() <= protection.LastResetHeight &&
				protection.CurrentBlockLiquidations >= protection.MaxLiquidationsPerBlock {
				return ctx, errorsmod.Wrap(circuittypes.ErrLiquidationSurge, "liquidation limit reached")
			}
		}
	}

	return next(ctx, tx, simulate)
}

// Ensure decorators implement sdk.AnteDecorator
var (
	_ sdk.AnteDecorator = CircuitBreakerDecorator{}
	_ sdk.AnteDecorator = LiquidationSurgeDecorator{}
	_ sdk.AnteDecorator = OracleValidationDecorator{}
	_ sdk.AnteDecorator = CombinedSecurityDecorator{}
)
