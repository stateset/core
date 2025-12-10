package keeper

import (
	"context"
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/treasury/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns a MsgServer backed by the treasury keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return msgServer{keeper: k}
}

// RecordReserve records a new reserve snapshot
func (m msgServer) RecordReserve(goCtx context.Context, msg *types.MsgRecordReserve) (*types.MsgRecordReserveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	id := m.keeper.RecordSnapshot(goCtx, msg.Snapshot)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSnapshotRecorded,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Authority),
			sdk.NewAttribute(types.AttributeKeySnapshotID, strconv.FormatUint(id, 10)),
			sdk.NewAttribute(types.AttributeKeyReporter, msg.Snapshot.Reporter),
			sdk.NewAttribute(types.AttributeKeyTotal, msg.Snapshot.TotalSupply.String()),
		),
	)

	return &types.MsgRecordReserveResponse{SnapshotID: id}, nil
}

// ProposeSpend creates a new time-locked spend proposal
func (m msgServer) ProposeSpend(goCtx context.Context, msg *types.MsgProposeSpend) (*types.MsgProposeSpendResponse, error) {
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	timelockDuration := time.Duration(msg.TimelockSeconds) * time.Second

	proposalID, err := m.keeper.CreateSpendProposal(
		goCtx,
		msg.Authority,
		msg.Recipient,
		msg.Amount,
		msg.Category,
		msg.Description,
		timelockDuration,
	)
	if err != nil {
		return nil, err
	}

	proposal, _ := m.keeper.GetSpendProposal(goCtx, proposalID)

	return &types.MsgProposeSpendResponse{
		ProposalID:   proposalID,
		ExecuteAfter: proposal.ExecuteAfter,
		ExpiresAt:    proposal.ExpiresAt,
	}, nil
}

// ExecuteSpend executes a spend proposal after timelock expires
func (m msgServer) ExecuteSpend(goCtx context.Context, msg *types.MsgExecuteSpend) (*types.MsgExecuteSpendResponse, error) {
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get proposal before execution to return details
	proposal, found := m.keeper.GetSpendProposal(goCtx, msg.ProposalID)
	if !found {
		return nil, types.ErrProposalNotFound
	}

	if err := m.keeper.ExecuteSpendProposal(goCtx, msg.ProposalID); err != nil {
		return nil, err
	}

	return &types.MsgExecuteSpendResponse{
		Recipient: proposal.Recipient,
		Amount:    proposal.Amount,
	}, nil
}

// CancelSpend cancels a pending spend proposal
func (m msgServer) CancelSpend(goCtx context.Context, msg *types.MsgCancelSpend) (*types.MsgCancelSpendResponse, error) {
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.CancelSpendProposal(goCtx, msg.ProposalID, msg.Reason); err != nil {
		return nil, err
	}

	return &types.MsgCancelSpendResponse{}, nil
}

// SetBudget sets budget limits for a category
func (m msgServer) SetBudget(goCtx context.Context, msg *types.MsgSetBudget) (*types.MsgSetBudgetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	budget := types.Budget{
		Category:       msg.Category,
		TotalLimit:     msg.TotalLimit,
		PeriodLimit:    msg.PeriodLimit,
		PeriodDuration: msg.PeriodDuration,
		PeriodStart:    ctx.BlockTime(),
		PeriodSpent:    nil,
		TotalSpent:     nil,
		Enabled:        msg.Enabled,
	}

	if err := m.keeper.SetBudget(goCtx, budget); err != nil {
		return nil, err
	}

	return &types.MsgSetBudgetResponse{}, nil
}

// UpdateParams updates treasury parameters (governance only)
func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.SetParams(goCtx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
