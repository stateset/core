package types

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Legacy vault-based messages
	TypeCreateVault        = "create_vault"
	TypeDepositCollateral  = "deposit_collateral"
	TypeWithdrawCollateral = "withdraw_collateral"
	TypeMintStablecoin     = "mint_stablecoin"
	TypeRepayStablecoin    = "repay_stablecoin"
	TypeLiquidateVault     = "liquidate_vault"

	// Reserve-backed stablecoin messages
	TypeDepositReserve       = "deposit_reserve"
	TypeRequestRedemption    = "request_redemption"
	TypeExecuteRedemption    = "execute_redemption"
	TypeCancelRedemption     = "cancel_redemption"
	TypeUpdateReserveParams  = "update_reserve_params"
	TypeRecordAttestation    = "record_attestation"
	TypeSetApprovedAttester  = "set_approved_attester"
)

var (
	// Legacy vault messages
	_ sdk.Msg = (*MsgCreateVault)(nil)
	_ sdk.Msg = (*MsgDepositCollateral)(nil)
	_ sdk.Msg = (*MsgWithdrawCollateral)(nil)
	_ sdk.Msg = (*MsgMintStablecoin)(nil)
	_ sdk.Msg = (*MsgRepayStablecoin)(nil)
	_ sdk.Msg = (*MsgLiquidateVault)(nil)

	// Reserve-backed stablecoin messages
	_ sdk.Msg = (*MsgDepositReserve)(nil)
	_ sdk.Msg = (*MsgRequestRedemption)(nil)
	_ sdk.Msg = (*MsgExecuteRedemption)(nil)
	_ sdk.Msg = (*MsgCancelRedemption)(nil)
	_ sdk.Msg = (*MsgUpdateReserveParams)(nil)
	_ sdk.Msg = (*MsgRecordAttestation)(nil)
	_ sdk.Msg = (*MsgSetApprovedAttester)(nil)
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
	// Legacy vault operations
	CreateVault(ctx context.Context, msg *MsgCreateVault) (*MsgCreateVaultResponse, error)
	DepositCollateral(ctx context.Context, msg *MsgDepositCollateral) (*MsgDepositCollateralResponse, error)
	WithdrawCollateral(ctx context.Context, msg *MsgWithdrawCollateral) (*MsgWithdrawCollateralResponse, error)
	MintStablecoin(ctx context.Context, msg *MsgMintStablecoin) (*MsgMintStablecoinResponse, error)
	RepayStablecoin(ctx context.Context, msg *MsgRepayStablecoin) (*MsgRepayStablecoinResponse, error)
	LiquidateVault(ctx context.Context, msg *MsgLiquidateVault) (*MsgLiquidateVaultResponse, error)

	// Reserve-backed stablecoin operations
	DepositReserve(ctx context.Context, msg *MsgDepositReserve) (*MsgDepositReserveResponse, error)
	RequestRedemption(ctx context.Context, msg *MsgRequestRedemption) (*MsgRequestRedemptionResponse, error)
	ExecuteRedemption(ctx context.Context, msg *MsgExecuteRedemption) (*MsgExecuteRedemptionResponse, error)
	CancelRedemption(ctx context.Context, msg *MsgCancelRedemption) (*MsgCancelRedemptionResponse, error)
	UpdateReserveParams(ctx context.Context, msg *MsgUpdateReserveParams) (*MsgUpdateReserveParamsResponse, error)
	RecordAttestation(ctx context.Context, msg *MsgRecordAttestation) (*MsgRecordAttestationResponse, error)
	SetApprovedAttester(ctx context.Context, msg *MsgSetApprovedAttester) (*MsgSetApprovedAttesterResponse, error)
}

// ============================================================================
// Reserve-Backed Stablecoin Messages
// ============================================================================

// MsgDepositReserve deposits tokenized treasuries to mint ssUSD
type MsgDepositReserve struct {
	Depositor string   `json:"depositor" yaml:"depositor"`
	Amount    sdk.Coin `json:"amount" yaml:"amount"` // Tokenized treasury (e.g., usdy, stbt, ousg)
}

func (m *MsgDepositReserve) Reset()         { *m = MsgDepositReserve{} }
func (m *MsgDepositReserve) String() string { return fmt.Sprintf("MsgDepositReserve{%s %s}", m.Depositor, m.Amount) }
func (*MsgDepositReserve) ProtoMessage()    {}
func (m MsgDepositReserve) Route() string   { return RouterKey }
func (m MsgDepositReserve) Type() string    { return TypeDepositReserve }

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
	addr, _ := sdk.AccAddressFromBech32(m.Depositor)
	return []sdk.AccAddress{addr}
}

func (m MsgDepositReserve) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgDepositReserveResponse struct {
	DepositId   uint64 `json:"deposit_id"`
	SsusdMinted string `json:"ssusd_minted"`
}

// MsgRequestRedemption requests redemption of ssUSD for tokenized treasuries
type MsgRequestRedemption struct {
	Requester   string `json:"requester" yaml:"requester"`
	SsusdAmount string `json:"ssusd_amount" yaml:"ssusd_amount"` // Amount of ssUSD to redeem
	OutputDenom string `json:"output_denom" yaml:"output_denom"` // Desired output token (e.g., usdy, stbt)
}

func (m *MsgRequestRedemption) Reset()         { *m = MsgRequestRedemption{} }
func (m *MsgRequestRedemption) String() string { return fmt.Sprintf("MsgRequestRedemption{%s %s %s}", m.Requester, m.SsusdAmount, m.OutputDenom) }
func (*MsgRequestRedemption) ProtoMessage()    {}
func (m MsgRequestRedemption) Route() string   { return RouterKey }
func (m MsgRequestRedemption) Type() string    { return TypeRequestRedemption }

func (m MsgRequestRedemption) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Requester); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid requester address")
	}
	if m.SsusdAmount == "" {
		return errorsmod.Wrap(ErrInvalidAmount, "ssusd amount required")
	}
	if m.OutputDenom == "" {
		return errorsmod.Wrap(ErrInvalidAmount, "output denom required")
	}
	return nil
}

func (m MsgRequestRedemption) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Requester)
	return []sdk.AccAddress{addr}
}

func (m MsgRequestRedemption) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgRequestRedemptionResponse struct {
	RedemptionId uint64 `json:"redemption_id"`
}

// MsgExecuteRedemption executes a pending redemption (after timelock)
type MsgExecuteRedemption struct {
	Executor     string `json:"executor" yaml:"executor"`
	RedemptionId uint64 `json:"redemption_id" yaml:"redemption_id"`
}

func (m *MsgExecuteRedemption) Reset()         { *m = MsgExecuteRedemption{} }
func (m *MsgExecuteRedemption) String() string { return fmt.Sprintf("MsgExecuteRedemption{%s %d}", m.Executor, m.RedemptionId) }
func (*MsgExecuteRedemption) ProtoMessage()    {}
func (m MsgExecuteRedemption) Route() string   { return RouterKey }
func (m MsgExecuteRedemption) Type() string    { return TypeExecuteRedemption }

func (m MsgExecuteRedemption) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Executor); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid executor address")
	}
	if m.RedemptionId == 0 {
		return errorsmod.Wrap(ErrInvalidReserve, "redemption id required")
	}
	return nil
}

func (m MsgExecuteRedemption) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Executor)
	return []sdk.AccAddress{addr}
}

func (m MsgExecuteRedemption) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgExecuteRedemptionResponse struct{}

// MsgCancelRedemption cancels a pending redemption (authority only)
type MsgCancelRedemption struct {
	Authority    string `json:"authority" yaml:"authority"`
	RedemptionId uint64 `json:"redemption_id" yaml:"redemption_id"`
}

func (m *MsgCancelRedemption) Reset()         { *m = MsgCancelRedemption{} }
func (m *MsgCancelRedemption) String() string { return fmt.Sprintf("MsgCancelRedemption{%s %d}", m.Authority, m.RedemptionId) }
func (*MsgCancelRedemption) ProtoMessage()    {}
func (m MsgCancelRedemption) Route() string   { return RouterKey }
func (m MsgCancelRedemption) Type() string    { return TypeCancelRedemption }

func (m MsgCancelRedemption) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid authority address")
	}
	if m.RedemptionId == 0 {
		return errorsmod.Wrap(ErrInvalidReserve, "redemption id required")
	}
	return nil
}

func (m MsgCancelRedemption) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgCancelRedemption) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCancelRedemptionResponse struct{}

// MsgUpdateReserveParams updates reserve parameters (governance only)
type MsgUpdateReserveParams struct {
	Authority string        `json:"authority" yaml:"authority"`
	Params    ReserveParams `json:"params" yaml:"params"`
}

func (m *MsgUpdateReserveParams) Reset()         { *m = MsgUpdateReserveParams{} }
func (m *MsgUpdateReserveParams) String() string { return fmt.Sprintf("MsgUpdateReserveParams{%s}", m.Authority) }
func (*MsgUpdateReserveParams) ProtoMessage()    {}
func (m MsgUpdateReserveParams) Route() string   { return RouterKey }
func (m MsgUpdateReserveParams) Type() string    { return TypeUpdateReserveParams }

func (m MsgUpdateReserveParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid authority address")
	}
	return m.Params.ValidateBasic()
}

func (m MsgUpdateReserveParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateReserveParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgUpdateReserveParamsResponse struct{}

// MsgRecordAttestation records an off-chain reserve attestation
type MsgRecordAttestation struct {
	Attester      string `json:"attester" yaml:"attester"`
	TotalCash     string `json:"total_cash" yaml:"total_cash"`
	TotalTBills   string `json:"total_tbills" yaml:"total_tbills"`
	TotalTNotes   string `json:"total_tnotes" yaml:"total_tnotes"`
	TotalTBonds   string `json:"total_tbonds" yaml:"total_tbonds"`
	TotalRepos    string `json:"total_repos" yaml:"total_repos"`
	TotalMMF      string `json:"total_mmf" yaml:"total_mmf"`
	TotalValue    string `json:"total_value" yaml:"total_value"`
	CustodianName string `json:"custodian_name" yaml:"custodian_name"`
	AuditFirm     string `json:"audit_firm" yaml:"audit_firm"`
	ReportDate    string `json:"report_date" yaml:"report_date"` // RFC3339 format
	Hash          string `json:"hash" yaml:"hash"`               // Attestation document hash
}

func (m *MsgRecordAttestation) Reset()         { *m = MsgRecordAttestation{} }
func (m *MsgRecordAttestation) String() string { return fmt.Sprintf("MsgRecordAttestation{%s %s}", m.Attester, m.TotalValue) }
func (*MsgRecordAttestation) ProtoMessage()    {}
func (m MsgRecordAttestation) Route() string   { return RouterKey }
func (m MsgRecordAttestation) Type() string    { return TypeRecordAttestation }

func (m MsgRecordAttestation) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Attester); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid attester address")
	}
	if m.TotalValue == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "total value required")
	}
	if m.CustodianName == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "custodian name required")
	}
	return nil
}

func (m MsgRecordAttestation) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Attester)
	return []sdk.AccAddress{addr}
}

func (m MsgRecordAttestation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgRecordAttestationResponse struct {
	AttestationId uint64 `json:"attestation_id"`
}

// MsgSetApprovedAttester adds or removes an approved attester (authority only)
type MsgSetApprovedAttester struct {
	Authority string `json:"authority" yaml:"authority"`
	Attester  string `json:"attester" yaml:"attester"`
	Approved  bool   `json:"approved" yaml:"approved"`
}

func (m *MsgSetApprovedAttester) Reset()         { *m = MsgSetApprovedAttester{} }
func (m *MsgSetApprovedAttester) String() string { return fmt.Sprintf("MsgSetApprovedAttester{%s %s %t}", m.Authority, m.Attester, m.Approved) }
func (*MsgSetApprovedAttester) ProtoMessage()    {}
func (m MsgSetApprovedAttester) Route() string   { return RouterKey }
func (m MsgSetApprovedAttester) Type() string    { return TypeSetApprovedAttester }

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
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgSetApprovedAttester) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgSetApprovedAttesterResponse struct{}
