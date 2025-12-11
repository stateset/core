package keeper

import (
	"context"
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stablecoin/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns a MsgServer implementation backed by the stablecoin keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return msgServer{keeper: k}
}

func (m msgServer) CreateVault(goCtx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidVault, err.Error())
	}

	vaultID, err := m.keeper.CreateVault(ctx, owner, msg.Collateral, msg.Debt)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVaultCreated,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyVaultID, strconv.FormatUint(vaultID, 10)),
			sdk.NewAttribute(types.AttributeKeyCollateral, msg.Collateral.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Debt.String()),
		),
	)

	return &types.MsgCreateVaultResponse{VaultId: vaultID}, nil
}

func (m msgServer) DepositCollateral(goCtx context.Context, msg *types.MsgDepositCollateral) (*types.MsgDepositCollateralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidVault, err.Error())
	}

	if err := m.keeper.DepositCollateral(ctx, owner, msg.VaultId, msg.Collateral); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCollateralDeposited,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyVaultID, strconv.FormatUint(msg.VaultId, 10)),
			sdk.NewAttribute(types.AttributeKeyCollateral, msg.Collateral.String()),
		),
	)

	return &types.MsgDepositCollateralResponse{}, nil
}

func (m msgServer) WithdrawCollateral(goCtx context.Context, msg *types.MsgWithdrawCollateral) (*types.MsgWithdrawCollateralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidVault, err.Error())
	}

	if err := m.keeper.WithdrawCollateral(ctx, owner, msg.VaultId, msg.Collateral); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCollateralWithdrawn,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyVaultID, strconv.FormatUint(msg.VaultId, 10)),
			sdk.NewAttribute(types.AttributeKeyCollateral, msg.Collateral.String()),
		),
	)

	return &types.MsgWithdrawCollateralResponse{}, nil
}

func (m msgServer) MintStablecoin(goCtx context.Context, msg *types.MsgMintStablecoin) (*types.MsgMintStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidVault, err.Error())
	}

	if err := m.keeper.MintStablecoin(ctx, owner, msg.VaultId, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStablecoinMinted,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyVaultID, strconv.FormatUint(msg.VaultId, 10)),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
	)

	return &types.MsgMintStablecoinResponse{}, nil
}

func (m msgServer) RepayStablecoin(goCtx context.Context, msg *types.MsgRepayStablecoin) (*types.MsgRepayStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidVault, err.Error())
	}

	if err := m.keeper.RepayStablecoin(ctx, owner, msg.VaultId, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStablecoinRepaid,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeKeyVaultID, strconv.FormatUint(msg.VaultId, 10)),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
	)

	return &types.MsgRepayStablecoinResponse{}, nil
}

func (m msgServer) LiquidateVault(goCtx context.Context, msg *types.MsgLiquidateVault) (*types.MsgLiquidateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	liquidator, err := sdk.AccAddressFromBech32(msg.Liquidator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidVault, err.Error())
	}

	collateral, err := m.keeper.LiquidateVault(ctx, liquidator, msg.VaultId)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVaultLiquidated,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Liquidator),
			sdk.NewAttribute(types.AttributeKeyLiquidator, msg.Liquidator),
			sdk.NewAttribute(types.AttributeKeyVaultID, strconv.FormatUint(msg.VaultId, 10)),
			sdk.NewAttribute(types.AttributeKeyCollateral, collateral.String()),
		),
	)

	return &types.MsgLiquidateVaultResponse{}, nil
}

// ============================================================================
// Reserve-Backed Stablecoin Messages
// ============================================================================

func (m msgServer) DepositReserve(goCtx context.Context, msg *types.MsgDepositReserve) (*types.MsgDepositReserveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	depositID, ssusdMinted, err := m.keeper.DepositReserve(ctx, depositor, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositReserveResponse{
		DepositId:   depositID,
		SsusdMinted: ssusdMinted.String(),
	}, nil
}

func (m msgServer) RequestRedemption(goCtx context.Context, msg *types.MsgRequestRedemption) (*types.MsgRequestRedemptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	requester, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	ssusdAmount, ok := sdkmath.NewIntFromString(msg.SsusdAmount)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid ssusd amount")
	}

	redemptionID, err := m.keeper.RequestRedemption(ctx, requester, ssusdAmount, msg.OutputDenom)
	if err != nil {
		return nil, err
	}

	return &types.MsgRequestRedemptionResponse{
		RedemptionId: redemptionID,
	}, nil
}

func (m msgServer) ExecuteRedemption(goCtx context.Context, msg *types.MsgExecuteRedemption) (*types.MsgExecuteRedemptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.ExecutePendingRedemption(ctx, msg.RedemptionId); err != nil {
		return nil, err
	}

	return &types.MsgExecuteRedemptionResponse{}, nil
}

func (m msgServer) CancelRedemption(goCtx context.Context, msg *types.MsgCancelRedemption) (*types.MsgCancelRedemptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.CancelRedemption(ctx, msg.Authority, msg.RedemptionId); err != nil {
		return nil, err
	}

	return &types.MsgCancelRedemptionResponse{}, nil
}

func (m msgServer) UpdateReserveParams(goCtx context.Context, msg *types.MsgUpdateReserveParams) (*types.MsgUpdateReserveParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Verify authority
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", m.keeper.GetAuthority(), msg.Authority)
	}

	if err := m.keeper.SetReserveParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateReserveParamsResponse{}, nil
}

func (m msgServer) RecordAttestation(goCtx context.Context, msg *types.MsgRecordAttestation) (*types.MsgRecordAttestationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Parse amounts
	parseInt := func(raw string, field string) (sdkmath.Int, error) {
		v, ok := sdkmath.NewIntFromString(raw)
		if !ok {
			return sdkmath.ZeroInt(), errorsmod.Wrapf(types.ErrInvalidReserve, "invalid %s amount", field)
		}
		return v, nil
	}

	totalCash, err := parseInt(msg.TotalCash, "total_cash")
	if err != nil {
		return nil, err
	}
	totalTBills, err := parseInt(msg.TotalTBills, "total_t_bills")
	if err != nil {
		return nil, err
	}
	totalTNotes, err := parseInt(msg.TotalTNotes, "total_t_notes")
	if err != nil {
		return nil, err
	}
	totalTBonds, err := parseInt(msg.TotalTBonds, "total_t_bonds")
	if err != nil {
		return nil, err
	}
	totalRepos, err := parseInt(msg.TotalRepos, "total_repos")
	if err != nil {
		return nil, err
	}
	totalMMF, err := parseInt(msg.TotalMMF, "total_mmf")
	if err != nil {
		return nil, err
	}
	totalValue, err := parseInt(msg.TotalValue, "total_value")
	if err != nil {
		return nil, err
	}

	// Parse report date
	reportDate, err := parseDate(msg.ReportDate)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, "invalid report date format")
	}

	attestation := types.OffChainReserveAttestation{
		Attester:        msg.Attester,
		TotalCash:       totalCash,
		TotalTBills:     totalTBills,
		TotalTNotes:     totalTNotes,
		TotalTBonds:     totalTBonds,
		TotalRepos:      totalRepos,
		TotalMMF:        totalMMF,
		TotalValue:      totalValue,
		CustodianName:   msg.CustodianName,
		AuditFirm:       msg.AuditFirm,
		ReportDate:      reportDate,
		AttestationHash: msg.Hash,
	}

	attestationID, err := m.keeper.RecordAttestation(ctx, attestation)
	if err != nil {
		return nil, err
	}

	return &types.MsgRecordAttestationResponse{
		AttestationId: attestationID,
	}, nil
}

func (m msgServer) SetApprovedAttester(goCtx context.Context, msg *types.MsgSetApprovedAttester) (*types.MsgSetApprovedAttesterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Verify authority
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", m.keeper.GetAuthority(), msg.Authority)
	}

	m.keeper.SetApprovedAttester(ctx, msg.Attester, msg.Approved)

	return &types.MsgSetApprovedAttesterResponse{}, nil
}

func parseDate(dateStr string) (time.Time, error) {
	// Try RFC3339 first
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return t, nil
	}
	// Try simple date format
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return t, nil
	}
	return time.Time{}, errorsmod.Wrap(types.ErrInvalidReserve, "invalid date format")
}
