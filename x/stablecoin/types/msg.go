package types

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeCreateVault        = "create_vault"
	TypeDepositCollateral  = "deposit_collateral"
	TypeWithdrawCollateral = "withdraw_collateral"
	TypeMintStablecoin     = "mint_stablecoin"
	TypeRepayStablecoin    = "repay_stablecoin"
	TypeLiquidateVault     = "liquidate_vault"
)

var (
	_ sdk.Msg = (*MsgCreateVault)(nil)
	_ sdk.Msg = (*MsgDepositCollateral)(nil)
	_ sdk.Msg = (*MsgWithdrawCollateral)(nil)
	_ sdk.Msg = (*MsgMintStablecoin)(nil)
	_ sdk.Msg = (*MsgRepayStablecoin)(nil)
	_ sdk.Msg = (*MsgLiquidateVault)(nil)
)

type MsgCreateVault struct {
	Owner      string   `json:"owner" yaml:"owner"`
	Collateral sdk.Coin `json:"collateral" yaml:"collateral"`
	Debt       sdk.Coin `json:"debt" yaml:"debt"`
}

func (m *MsgCreateVault) Reset() { *m = MsgCreateVault{} }
func (m *MsgCreateVault) String() string {
	return fmt.Sprintf("MsgCreateVault{%s %s %s}", m.Owner, m.Collateral.String(), m.Debt.String())
}
func (*MsgCreateVault) ProtoMessage() {}
func NewMsgCreateVault(owner string, collateral, debt sdk.Coin) *MsgCreateVault {
	return &MsgCreateVault{Owner: owner, Collateral: collateral, Debt: debt}
}

func (m MsgCreateVault) Route() string { return RouterKey }
func (m MsgCreateVault) Type() string  { return TypeCreateVault }

func (m MsgCreateVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if !m.Collateral.IsValid() || m.Collateral.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "collateral must be positive")
	}
	if !m.Debt.IsZero() {
		if m.Debt.Denom != StablecoinDenom {
			return errorsmod.Wrap(ErrInvalidAmount, "debt denom must be stablecoin")
		}
	}
	return nil
}

func (m MsgCreateVault) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgCreateVault) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCreateVaultResponse struct {
	VaultId uint64 `json:"vault_id"`
}

type MsgDepositCollateral struct {
	Owner      string   `json:"owner" yaml:"owner"`
	VaultId    uint64   `json:"vault_id" yaml:"vault_id"`
	Collateral sdk.Coin `json:"collateral" yaml:"collateral"`
}

func (m *MsgDepositCollateral) Reset() { *m = MsgDepositCollateral{} }
func (m *MsgDepositCollateral) String() string {
	return fmt.Sprintf("MsgDepositCollateral{%s %d %s}", m.Owner, m.VaultId, m.Collateral.String())
}
func (*MsgDepositCollateral) ProtoMessage() {}
func NewMsgDepositCollateral(owner string, vaultID uint64, collateral sdk.Coin) *MsgDepositCollateral {
	return &MsgDepositCollateral{Owner: owner, VaultId: vaultID, Collateral: collateral}
}

func (m MsgDepositCollateral) Route() string { return RouterKey }
func (m MsgDepositCollateral) Type() string  { return TypeDepositCollateral }

func (m MsgDepositCollateral) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if m.VaultId == 0 {
		return errorsmod.Wrap(ErrInvalidVault, "vault id required")
	}
	if !m.Collateral.IsValid() || m.Collateral.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "collateral must be positive")
	}
	return nil
}

func (m MsgDepositCollateral) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgDepositCollateral) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgWithdrawCollateral struct {
	Owner      string   `json:"owner" yaml:"owner"`
	VaultId    uint64   `json:"vault_id" yaml:"vault_id"`
	Collateral sdk.Coin `json:"collateral" yaml:"collateral"`
}

func (m *MsgWithdrawCollateral) Reset() { *m = MsgWithdrawCollateral{} }
func (m *MsgWithdrawCollateral) String() string {
	return fmt.Sprintf("MsgWithdrawCollateral{%s %d %s}", m.Owner, m.VaultId, m.Collateral.String())
}
func (*MsgWithdrawCollateral) ProtoMessage()  {}
func (m MsgWithdrawCollateral) Route() string { return RouterKey }
func (m MsgWithdrawCollateral) Type() string  { return TypeWithdrawCollateral }

func (m MsgWithdrawCollateral) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if m.VaultId == 0 {
		return errorsmod.Wrap(ErrInvalidVault, "vault id required")
	}
	if !m.Collateral.IsValid() || m.Collateral.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "collateral must be positive")
	}
	return nil
}

func (m MsgWithdrawCollateral) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgWithdrawCollateral) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgMintStablecoin struct {
	Owner   string   `json:"owner" yaml:"owner"`
	VaultId uint64   `json:"vault_id" yaml:"vault_id"`
	Amount  sdk.Coin `json:"amount" yaml:"amount"`
}

func (m *MsgMintStablecoin) Reset() { *m = MsgMintStablecoin{} }
func (m *MsgMintStablecoin) String() string {
	return fmt.Sprintf("MsgMintStablecoin{%s %d %s}", m.Owner, m.VaultId, m.Amount.String())
}
func (*MsgMintStablecoin) ProtoMessage()  {}
func (m MsgMintStablecoin) Route() string { return RouterKey }
func (m MsgMintStablecoin) Type() string  { return TypeMintStablecoin }

func (m MsgMintStablecoin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if m.VaultId == 0 {
		return errorsmod.Wrap(ErrInvalidVault, "vault id required")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	if m.Amount.Denom != StablecoinDenom {
		return errorsmod.Wrap(ErrInvalidAmount, "amount denom must be stablecoin")
	}
	return nil
}

func (m MsgMintStablecoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgMintStablecoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgMintStablecoinResponse struct{}

type MsgRepayStablecoin struct {
	Owner   string   `json:"owner" yaml:"owner"`
	VaultId uint64   `json:"vault_id" yaml:"vault_id"`
	Amount  sdk.Coin `json:"amount" yaml:"amount"`
}

func (m *MsgRepayStablecoin) Reset() { *m = MsgRepayStablecoin{} }
func (m *MsgRepayStablecoin) String() string {
	return fmt.Sprintf("MsgRepayStablecoin{%s %d %s}", m.Owner, m.VaultId, m.Amount.String())
}
func (*MsgRepayStablecoin) ProtoMessage()  {}
func (m MsgRepayStablecoin) Route() string { return RouterKey }
func (m MsgRepayStablecoin) Type() string  { return TypeRepayStablecoin }

func (m MsgRepayStablecoin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if m.VaultId == 0 {
		return errorsmod.Wrap(ErrInvalidVault, "vault id required")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	if m.Amount.Denom != StablecoinDenom {
		return errorsmod.Wrap(ErrInvalidAmount, "amount denom must be stablecoin")
	}
	return nil
}

func (m MsgRepayStablecoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgRepayStablecoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgRepayStablecoinResponse struct{}

// MsgLiquidateVault repays an unhealthy vault's debt in exchange for its collateral.
type MsgLiquidateVault struct {
	Liquidator string `json:"liquidator" yaml:"liquidator"`
	VaultId    uint64 `json:"vault_id" yaml:"vault_id"`
}

func (m *MsgLiquidateVault) Reset() { *m = MsgLiquidateVault{} }
func (m *MsgLiquidateVault) String() string {
	return fmt.Sprintf("MsgLiquidateVault{%s %d}", m.Liquidator, m.VaultId)
}
func (*MsgLiquidateVault) ProtoMessage()  {}
func (m MsgLiquidateVault) Route() string { return RouterKey }
func (m MsgLiquidateVault) Type() string  { return TypeLiquidateVault }

func (m MsgLiquidateVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Liquidator); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if m.VaultId == 0 {
		return errorsmod.Wrap(ErrInvalidVault, "vault id required")
	}
	return nil
}

func (m MsgLiquidateVault) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Liquidator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgLiquidateVault) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgLiquidateVaultResponse struct{}

type MsgDepositCollateralResponse struct{}

type MsgWithdrawCollateralResponse struct{}

type MsgServer interface {
	CreateVault(ctx context.Context, msg *MsgCreateVault) (*MsgCreateVaultResponse, error)
	DepositCollateral(ctx context.Context, msg *MsgDepositCollateral) (*MsgDepositCollateralResponse, error)
	WithdrawCollateral(ctx context.Context, msg *MsgWithdrawCollateral) (*MsgWithdrawCollateralResponse, error)
	MintStablecoin(ctx context.Context, msg *MsgMintStablecoin) (*MsgMintStablecoinResponse, error)
	RepayStablecoin(ctx context.Context, msg *MsgRepayStablecoin) (*MsgRepayStablecoinResponse, error)
	LiquidateVault(ctx context.Context, msg *MsgLiquidateVault) (*MsgLiquidateVaultResponse, error)
}
