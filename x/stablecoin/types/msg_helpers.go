package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmath "cosmossdk.io/math"
)

func NewMsgCreateVault(owner string, collateral, debt sdk.Coin) *MsgCreateVault {
	return &MsgCreateVault{Owner: owner, Collateral: collateral, Debt: debt}
}

func (m MsgCreateVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return errorsmod.Wrap(ErrInvalidVault, err.Error())
	}
	if !m.Collateral.IsValid() || m.Collateral.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "collateral must be positive")
	}
	if !m.Debt.IsZero() && m.Debt.Denom != StablecoinDenom {
		return errorsmod.Wrap(ErrInvalidAmount, "debt denom must be stablecoin")
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

func NewMsgDepositCollateral(owner string, vaultID uint64, collateral sdk.Coin) *MsgDepositCollateral {
	return &MsgDepositCollateral{Owner: owner, VaultId: vaultID, Collateral: collateral}
}

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

func NewMsgWithdrawCollateral(owner string, vaultID uint64, collateral sdk.Coin) *MsgWithdrawCollateral {
	return &MsgWithdrawCollateral{Owner: owner, VaultId: vaultID, Collateral: collateral}
}

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

func NewMsgMintStablecoin(owner string, vaultID uint64, amount sdk.Coin) *MsgMintStablecoin {
	return &MsgMintStablecoin{Owner: owner, VaultId: vaultID, Amount: amount}
}

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

func NewMsgRepayStablecoin(owner string, vaultID uint64, amount sdk.Coin) *MsgRepayStablecoin {
	return &MsgRepayStablecoin{Owner: owner, VaultId: vaultID, Amount: amount}
}

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

func NewMsgLiquidateVault(liquidator string, vaultID uint64) *MsgLiquidateVault {
	return &MsgLiquidateVault{Liquidator: liquidator, VaultId: vaultID}
}

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

func NewMsgDepositReserve(depositor string, amount sdk.Coin) *MsgDepositReserve {
	return &MsgDepositReserve{Depositor: depositor, Amount: amount}
}

func (m MsgDepositReserve) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Depositor); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid depositor address")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	return nil
}

func (m MsgDepositReserve) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgRequestRedemption(requester string, ssusdAmount string, outputDenom string) *MsgRequestRedemption {
	return &MsgRequestRedemption{Requester: requester, SsusdAmount: ssusdAmount, OutputDenom: outputDenom}
}

func (m MsgRequestRedemption) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Requester); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid requester address")
	}
	if m.OutputDenom == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "output denom required")
	}
	ssusdAmount, ok := sdkmath.NewIntFromString(m.SsusdAmount)
	if !ok || !ssusdAmount.IsPositive() {
		return errorsmod.Wrap(ErrInvalidAmount, "ssusd amount must be positive")
	}
	return nil
}

func (m MsgRequestRedemption) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Requester)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgExecuteRedemption(executor string, redemptionID uint64) *MsgExecuteRedemption {
	return &MsgExecuteRedemption{Executor: executor, RedemptionId: redemptionID}
}

func (m MsgExecuteRedemption) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Executor); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid executor address")
	}
	if m.RedemptionId == 0 {
		return errorsmod.Wrap(ErrRedemptionNotFound, "redemption id required")
	}
	return nil
}

func (m MsgExecuteRedemption) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Executor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgCancelRedemption(authority string, redemptionID uint64) *MsgCancelRedemption {
	return &MsgCancelRedemption{Authority: authority, RedemptionId: redemptionID}
}

func (m MsgCancelRedemption) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid authority address")
	}
	if m.RedemptionId == 0 {
		return errorsmod.Wrap(ErrRedemptionNotFound, "redemption id required")
	}
	return nil
}

func (m MsgCancelRedemption) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgUpdateReserveParams(authority string, params ReserveParams) *MsgUpdateReserveParams {
	return &MsgUpdateReserveParams{Authority: authority, Params: params}
}

func (m MsgUpdateReserveParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid authority address")
	}
	return m.Params.ValidateBasic()
}

func (m MsgUpdateReserveParams) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgRecordAttestation) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Attester); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid attester address")
	}
	if m.CustodianName == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "custodian name required")
	}
	if m.ReportDate == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "report date required")
	}
	if m.Hash == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "hash required")
	}

	parseNonNegativeInt := func(raw string, field string) error {
		v, ok := sdkmath.NewIntFromString(raw)
		if !ok {
			return errorsmod.Wrapf(ErrInvalidReserve, "invalid %s amount", field)
		}
		if v.IsNegative() {
			return errorsmod.Wrapf(ErrInvalidReserve, "%s cannot be negative", field)
		}
		return nil
	}

	if err := parseNonNegativeInt(m.TotalCash, "total_cash"); err != nil {
		return err
	}
	if err := parseNonNegativeInt(m.TotalTbills, "total_tbills"); err != nil {
		return err
	}
	if err := parseNonNegativeInt(m.TotalTnotes, "total_tnotes"); err != nil {
		return err
	}
	if err := parseNonNegativeInt(m.TotalTbonds, "total_tbonds"); err != nil {
		return err
	}
	if err := parseNonNegativeInt(m.TotalRepos, "total_repos"); err != nil {
		return err
	}
	if err := parseNonNegativeInt(m.TotalMmf, "total_mmf"); err != nil {
		return err
	}
	if err := parseNonNegativeInt(m.TotalValue, "total_value"); err != nil {
		return err
	}

	return nil
}

func (m MsgRecordAttestation) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Attester)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgSetApprovedAttester(authority, attester string, approved bool) *MsgSetApprovedAttester {
	return &MsgSetApprovedAttester{Authority: authority, Attester: attester, Approved: approved}
}

func (m MsgSetApprovedAttester) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid authority address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Attester); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid attester address")
	}
	return nil
}

func (m MsgSetApprovedAttester) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
