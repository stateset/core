package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
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
