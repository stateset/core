package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/stablecoin/types"
)

// EndBlocker executes per-block logic for the stablecoin module
func (k Keeper) EndBlocker(ctx sdk.Context) error {
	// 1. Update Reserve Value (re-calculate based on latest oracle prices)
	if err := k.UpdateReserveValue(ctx); err != nil {
		k.Logger(ctx).Error("failed to update reserve value", "error", err)
	}

	// 2. Solvency Check
	reserve := k.GetReserve(ctx)
	params := k.GetReserveParams(ctx)

	// Skip if no stablecoins minted
	if reserve.TotalMinted.IsZero() {
		return nil
	}

	reserveRatio := reserve.GetReserveRatio()

	// Critical Threshold: 90% (9000 bps) - Hardcoded safety net
	// If reserves drop below 90%, we enter "Panic Mode"
	const CriticalThreshold = int64(9000)

	if reserveRatio < CriticalThreshold {
		// Only trigger if not already paused, to avoid spamming events/writes
		if !params.MintPaused || !params.RedeemPaused {
			k.Logger(ctx).Error("CRITICAL: Reserve ratio below safety threshold. Pausing module.", "ratio", reserveRatio)

			// Auto-Pause
			params.MintPaused = true
			params.RedeemPaused = true
			if err := k.SetReserveParams(ctx, params); err != nil {
				k.Logger(ctx).Error("failed to set panic mode params", "error", err)
			}

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeSolvencyEmergency,
					sdk.NewAttribute(types.AttributeKeyReserveRatio, sdk.NewInt(reserveRatio).String()),
					sdk.NewAttribute(types.AttributeKeyAction, "system_paused"),
				),
			)
		}
	}

	return nil
}
