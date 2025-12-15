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

	// Verify authority FIRST before validating params
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "invalid authority: expected %s, got %s", m.keeper.GetAuthority(), msg.Authority)
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
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
	totalTBills, err := parseInt(msg.TotalTbills, "total_t_bills")
	if err != nil {
		return nil, err
	}
	totalTNotes, err := parseInt(msg.TotalTnotes, "total_t_notes")
	if err != nil {
		return nil, err
	}
	totalTBonds, err := parseInt(msg.TotalTbonds, "total_t_bonds")
	if err != nil {
		return nil, err
	}
	totalRepos, err := parseInt(msg.TotalRepos, "total_repos")
	if err != nil {
		return nil, err
	}
	totalMMF, err := parseInt(msg.TotalMmf, "total_mmf")
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
		TotalTbills:     totalTBills,
		TotalTnotes:     totalTNotes,
		TotalTbonds:     totalTBonds,
		TotalRepos:      totalRepos,
		TotalMmf:        totalMMF,
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

// ============================================================================
// Peg Stability Module (PSM) Messages
// ============================================================================

func (m msgServer) PSMSwapIn(goCtx context.Context, msg *types.MsgPSMSwapIn) (*types.MsgPSMSwapInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	ssusdMinted, feeCharged, err := m.keeper.PSMSwapIn(ctx, sender, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgPSMSwapInResponse{
		SsusdMinted: ssusdMinted.String(),
		FeeCharged:  feeCharged.String(),
	}, nil
}

func (m msgServer) PSMSwapOut(goCtx context.Context, msg *types.MsgPSMSwapOut) (*types.MsgPSMSwapOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	ssusdAmount, ok := sdkmath.NewIntFromString(msg.SsusdAmount)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid ssusd amount")
	}

	outputAmount, feeCharged, err := m.keeper.PSMSwapOut(ctx, sender, ssusdAmount, msg.OutputDenom)
	if err != nil {
		return nil, err
	}

	return &types.MsgPSMSwapOutResponse{
		OutputAmount: outputAmount.String(),
		FeeCharged:   feeCharged.String(),
	}, nil
}

func (m msgServer) UpdatePSMConfig(goCtx context.Context, msg *types.MsgUpdatePSMConfig) (*types.MsgUpdatePSMConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.UpdatePSMConfigs(ctx, msg.Authority, msg.Configs); err != nil {
		return nil, err
	}

	return &types.MsgUpdatePSMConfigResponse{}, nil
}

// ============================================================================
// Savings Rate (ssSR) Messages
// ============================================================================

func (m msgServer) DepositSavings(goCtx context.Context, msg *types.MsgDepositSavings) (*types.MsgDepositSavingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	amount, ok := sdkmath.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid amount")
	}

	totalDeposit, err := m.keeper.DepositSavings(ctx, depositor, amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositSavingsResponse{
		TotalDeposit: totalDeposit.String(),
	}, nil
}

func (m msgServer) WithdrawSavings(goCtx context.Context, msg *types.MsgWithdrawSavings) (*types.MsgWithdrawSavingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	amount, ok := sdkmath.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid amount")
	}

	amountWithdrawn, interestEarned, err := m.keeper.WithdrawSavings(ctx, depositor, amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawSavingsResponse{
		AmountWithdrawn: amountWithdrawn.String(),
		InterestEarned:  interestEarned.String(),
	}, nil
}

func (m msgServer) ClaimSavingsInterest(goCtx context.Context, msg *types.MsgClaimSavingsInterest) (*types.MsgClaimSavingsInterestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	interestClaimed, err := m.keeper.ClaimSavingsInterest(ctx, depositor)
	if err != nil {
		return nil, err
	}

	return &types.MsgClaimSavingsInterestResponse{
		InterestClaimed: interestClaimed.String(),
	}, nil
}

func (m msgServer) UpdateSavingsParams(goCtx context.Context, msg *types.MsgUpdateSavingsParams) (*types.MsgUpdateSavingsParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.UpdateSavingsParams(ctx, msg.Authority, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateSavingsParamsResponse{}, nil
}

// ============================================================================
// Dutch Auction Messages
// ============================================================================

func (m msgServer) BidAuction(goCtx context.Context, msg *types.MsgBidAuction) (*types.MsgBidAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	bidder, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	maxCollateral, ok := sdkmath.NewIntFromString(msg.MaxCollateralAmount)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid max collateral amount")
	}

	maxSSUSD, ok := sdkmath.NewIntFromString(msg.MaxSsusdToSpend)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid max ssusd amount")
	}

	collateralPurchased, ssusdSpent, pricePerUnit, err := m.keeper.BidAuction(ctx, bidder, msg.AuctionId, maxCollateral, maxSSUSD)
	if err != nil {
		return nil, err
	}

	return &types.MsgBidAuctionResponse{
		CollateralPurchased: collateralPurchased.String(),
		SsusdSpent:          ssusdSpent.String(),
		PricePerUnit:        pricePerUnit.String(),
	}, nil
}

func (m msgServer) UpdateAuctionParams(goCtx context.Context, msg *types.MsgUpdateAuctionParams) (*types.MsgUpdateAuctionParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.UpdateAuctionParams(ctx, msg.Authority, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateAuctionParamsResponse{}, nil
}

// ============================================================================
// Flash Minting Messages
// ============================================================================

func (m msgServer) FlashMint(goCtx context.Context, msg *types.MsgFlashMint) (*types.MsgFlashMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	amount, ok := sdkmath.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid amount")
	}

	amountMinted, feePaid, err := m.keeper.FlashMint(ctx, sender, amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgFlashMintResponse{
		AmountMinted: amountMinted.String(),
		FeePaid:      feePaid.String(),
	}, nil
}

func (m msgServer) FlashMintCallback(goCtx context.Context, msg *types.MsgFlashMintCallback) (*types.MsgFlashMintCallbackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidReserve, err.Error())
	}

	amountToReturn, ok := sdkmath.NewIntFromString(msg.AmountToReturn)
	if !ok {
		return nil, errorsmod.Wrap(types.ErrInvalidAmount, "invalid amount to return")
	}

	if err := m.keeper.FlashMintCallback(ctx, sender, amountToReturn); err != nil {
		return nil, err
	}

	return &types.MsgFlashMintCallbackResponse{}, nil
}

func (m msgServer) UpdateFlashMintParams(goCtx context.Context, msg *types.MsgUpdateFlashMintParams) (*types.MsgUpdateFlashMintParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.UpdateFlashMintParams(ctx, msg.Authority, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateFlashMintParamsResponse{}, nil
}
