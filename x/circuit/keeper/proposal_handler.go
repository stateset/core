package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/stateset/core/x/circuit/types"
)

// NewCircuitProposalHandler creates a governance handler for circuit proposals
func NewCircuitProposalHandler(k Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.PauseSystemProposal:
			return handlePauseSystemProposal(ctx, k, c)
		case *types.ResumeSystemProposal:
			return handleResumeSystemProposal(ctx, k, c)
		case *types.TripCircuitProposal:
			return handleTripCircuitProposal(ctx, k, c)
		case *types.ResetCircuitProposal:
			return handleResetCircuitProposal(ctx, k, c)
		case *types.UpdateCircuitParamsProposal:
			return handleUpdateCircuitParamsProposal(ctx, k, c)
		default:
			return errorsmod.Wrapf(types.ErrInvalidParams, "unrecognized circuit proposal content type: %T", c)
		}
	}
}

func handlePauseSystemProposal(ctx sdk.Context, k Keeper, p *types.PauseSystemProposal) error {
	// Pause the system via governance
	k.PauseSystem(ctx, "governance", p.Reason, p.DurationSeconds)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"governance_pause_system",
			sdk.NewAttribute("reason", p.Reason),
			sdk.NewAttribute("title", p.Title),
		),
	)

	return nil
}

func handleResumeSystemProposal(ctx sdk.Context, k Keeper, p *types.ResumeSystemProposal) error {
	// Resume the system via governance
	if err := k.ResumeSystem(ctx, "governance"); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"governance_resume_system",
			sdk.NewAttribute("title", p.Title),
		),
	)

	return nil
}

func handleTripCircuitProposal(ctx sdk.Context, k Keeper, p *types.TripCircuitProposal) error {
	// Trip the circuit via governance
	k.TripCircuit(ctx, p.ModuleName, "governance", p.Reason, p.DisableMessages)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"governance_trip_circuit",
			sdk.NewAttribute("module", p.ModuleName),
			sdk.NewAttribute("reason", p.Reason),
			sdk.NewAttribute("title", p.Title),
		),
	)

	return nil
}

func handleResetCircuitProposal(ctx sdk.Context, k Keeper, p *types.ResetCircuitProposal) error {
	// Reset the circuit via governance
	if err := k.ResetCircuit(ctx, p.ModuleName, "governance"); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"governance_reset_circuit",
			sdk.NewAttribute("module", p.ModuleName),
			sdk.NewAttribute("title", p.Title),
		),
	)

	return nil
}

func handleUpdateCircuitParamsProposal(ctx sdk.Context, k Keeper, p *types.UpdateCircuitParamsProposal) error {
	// Update params via governance
	k.SetParams(ctx, p.Params)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"governance_update_circuit_params",
			sdk.NewAttribute("title", p.Title),
		),
	)

	return nil
}
