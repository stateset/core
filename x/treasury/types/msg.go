package types

import (
	"context"
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgRecordReserve    = "record_reserve"
	TypeMsgProposeSpend     = "propose_spend"
	TypeMsgExecuteSpend     = "execute_spend"
	TypeMsgCancelSpend      = "cancel_spend"
	TypeMsgSetBudget        = "set_budget"
	TypeMsgUpdateParams     = "update_treasury_params"
	TypeMsgReceiveRevenue   = "receive_revenue"
)

var (
	_ sdk.Msg = (*MsgRecordReserve)(nil)
	_ sdk.Msg = (*MsgProposeSpend)(nil)
	_ sdk.Msg = (*MsgExecuteSpend)(nil)
	_ sdk.Msg = (*MsgCancelSpend)(nil)
	_ sdk.Msg = (*MsgSetBudget)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
)

// ============================================================================
// MsgRecordReserve
// ============================================================================

// MsgRecordReserve records a new reserve snapshot.
type MsgRecordReserve struct {
	Authority string          `json:"authority" yaml:"authority"`
	Snapshot  ReserveSnapshot `json:"snapshot" yaml:"snapshot"`
}

func (m *MsgRecordReserve) Reset() { *m = MsgRecordReserve{} }
func (m *MsgRecordReserve) String() string {
	return fmt.Sprintf("MsgRecordReserve{%s %s}", m.Authority, m.Snapshot.TotalSupply.String())
}
func (*MsgRecordReserve) ProtoMessage() {}

func NewMsgRecordReserve(authority string, snapshot ReserveSnapshot) *MsgRecordReserve {
	return &MsgRecordReserve{
		Authority: authority,
		Snapshot:  snapshot,
	}
}

func (m MsgRecordReserve) Route() string { return RouterKey }
func (m MsgRecordReserve) Type() string  { return TypeMsgRecordReserve }

func (m MsgRecordReserve) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, err.Error())
	}
	return m.Snapshot.ValidateBasic()
}

func (m MsgRecordReserve) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgRecordReserve) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgRecordReserveResponse struct {
	SnapshotID uint64 `json:"snapshot_id"`
}

// ============================================================================
// MsgProposeSpend - Create a time-locked spend proposal
// ============================================================================

type MsgProposeSpend struct {
	Authority       string        `json:"authority" yaml:"authority"`
	Recipient       string        `json:"recipient" yaml:"recipient"`
	Amount          sdk.Coins     `json:"amount" yaml:"amount"`
	Category        string        `json:"category" yaml:"category"`
	Description     string        `json:"description" yaml:"description"`
	TimelockSeconds uint64        `json:"timelock_seconds" yaml:"timelock_seconds"`
}

func (m *MsgProposeSpend) Reset() { *m = MsgProposeSpend{} }
func (m *MsgProposeSpend) String() string {
	return fmt.Sprintf("MsgProposeSpend{%s -> %s: %s}", m.Authority, m.Recipient, m.Amount.String())
}
func (*MsgProposeSpend) ProtoMessage() {}

func NewMsgProposeSpend(authority, recipient string, amount sdk.Coins, category, description string, timelockSeconds uint64) *MsgProposeSpend {
	return &MsgProposeSpend{
		Authority:       authority,
		Recipient:       recipient,
		Amount:          amount,
		Category:        category,
		Description:     description,
		TimelockSeconds: timelockSeconds,
	}
}

func (m MsgProposeSpend) Route() string { return RouterKey }
func (m MsgProposeSpend) Type() string  { return TypeMsgProposeSpend }

func (m MsgProposeSpend) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be valid and non-zero")
	}
	if !IsValidCategory(m.Category) {
		return errorsmod.Wrapf(ErrInvalidCategory, "unknown category: %s", m.Category)
	}
	if len(m.Description) == 0 || len(m.Description) > 1000 {
		return errorsmod.Wrap(ErrInvalidAmount, "description must be 1-1000 characters")
	}
	return nil
}

func (m MsgProposeSpend) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgProposeSpend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgProposeSpendResponse struct {
	ProposalID   uint64    `json:"proposal_id"`
	ExecuteAfter time.Time `json:"execute_after"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// ============================================================================
// MsgExecuteSpend - Execute a time-locked spend proposal after timelock expires
// ============================================================================

type MsgExecuteSpend struct {
	Authority  string `json:"authority" yaml:"authority"`
	ProposalID uint64 `json:"proposal_id" yaml:"proposal_id"`
}

func (m *MsgExecuteSpend) Reset() { *m = MsgExecuteSpend{} }
func (m *MsgExecuteSpend) String() string {
	return fmt.Sprintf("MsgExecuteSpend{%s proposal:%d}", m.Authority, m.ProposalID)
}
func (*MsgExecuteSpend) ProtoMessage() {}

func NewMsgExecuteSpend(authority string, proposalID uint64) *MsgExecuteSpend {
	return &MsgExecuteSpend{
		Authority:  authority,
		ProposalID: proposalID,
	}
}

func (m MsgExecuteSpend) Route() string { return RouterKey }
func (m MsgExecuteSpend) Type() string  { return TypeMsgExecuteSpend }

func (m MsgExecuteSpend) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if m.ProposalID == 0 {
		return errorsmod.Wrap(ErrProposalNotFound, "proposal ID must be non-zero")
	}
	return nil
}

func (m MsgExecuteSpend) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgExecuteSpend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgExecuteSpendResponse struct {
	Recipient string    `json:"recipient"`
	Amount    sdk.Coins `json:"amount"`
	TxHash    string    `json:"tx_hash"`
}

// ============================================================================
// MsgCancelSpend - Cancel a pending spend proposal
// ============================================================================

type MsgCancelSpend struct {
	Authority  string `json:"authority" yaml:"authority"`
	ProposalID uint64 `json:"proposal_id" yaml:"proposal_id"`
	Reason     string `json:"reason" yaml:"reason"`
}

func (m *MsgCancelSpend) Reset() { *m = MsgCancelSpend{} }
func (m *MsgCancelSpend) String() string {
	return fmt.Sprintf("MsgCancelSpend{%s proposal:%d}", m.Authority, m.ProposalID)
}
func (*MsgCancelSpend) ProtoMessage() {}

func NewMsgCancelSpend(authority string, proposalID uint64, reason string) *MsgCancelSpend {
	return &MsgCancelSpend{
		Authority:  authority,
		ProposalID: proposalID,
		Reason:     reason,
	}
}

func (m MsgCancelSpend) Route() string { return RouterKey }
func (m MsgCancelSpend) Type() string  { return TypeMsgCancelSpend }

func (m MsgCancelSpend) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if m.ProposalID == 0 {
		return errorsmod.Wrap(ErrProposalNotFound, "proposal ID must be non-zero")
	}
	return nil
}

func (m MsgCancelSpend) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgCancelSpend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCancelSpendResponse struct{}

// ============================================================================
// MsgSetBudget - Set budget limits for a category
// ============================================================================

type MsgSetBudget struct {
	Authority      string        `json:"authority" yaml:"authority"`
	Category       string        `json:"category" yaml:"category"`
	TotalLimit     sdk.Coins     `json:"total_limit" yaml:"total_limit"`
	PeriodLimit    sdk.Coins     `json:"period_limit" yaml:"period_limit"`
	PeriodDuration time.Duration `json:"period_duration" yaml:"period_duration"`
	Enabled        bool          `json:"enabled" yaml:"enabled"`
}

func (m *MsgSetBudget) Reset() { *m = MsgSetBudget{} }
func (m *MsgSetBudget) String() string {
	return fmt.Sprintf("MsgSetBudget{%s category:%s}", m.Authority, m.Category)
}
func (*MsgSetBudget) ProtoMessage() {}

func NewMsgSetBudget(authority, category string, totalLimit, periodLimit sdk.Coins, periodDuration time.Duration, enabled bool) *MsgSetBudget {
	return &MsgSetBudget{
		Authority:      authority,
		Category:       category,
		TotalLimit:     totalLimit,
		PeriodLimit:    periodLimit,
		PeriodDuration: periodDuration,
		Enabled:        enabled,
	}
}

func (m MsgSetBudget) Route() string { return RouterKey }
func (m MsgSetBudget) Type() string  { return TypeMsgSetBudget }

func (m MsgSetBudget) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if !IsValidCategory(m.Category) {
		return errorsmod.Wrapf(ErrInvalidCategory, "unknown category: %s", m.Category)
	}
	return nil
}

func (m MsgSetBudget) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgSetBudget) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgSetBudgetResponse struct{}

// ============================================================================
// MsgUpdateParams - Update treasury parameters (governance only)
// ============================================================================

type MsgUpdateParams struct {
	Authority string         `json:"authority" yaml:"authority"`
	Params    TreasuryParams `json:"params" yaml:"params"`
}

func (m *MsgUpdateParams) Reset() { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string {
	return fmt.Sprintf("MsgUpdateParams{%s}", m.Authority)
}
func (*MsgUpdateParams) ProtoMessage() {}

func NewMsgUpdateParams(authority string, params TreasuryParams) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}
}

func (m MsgUpdateParams) Route() string { return RouterKey }
func (m MsgUpdateParams) Type() string  { return TypeMsgUpdateParams }

func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	return m.Params.Validate()
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgUpdateParamsResponse struct{}

// ============================================================================
// MsgServer interface
// ============================================================================

type MsgServer interface {
	RecordReserve(ctx context.Context, msg *MsgRecordReserve) (*MsgRecordReserveResponse, error)
	ProposeSpend(ctx context.Context, msg *MsgProposeSpend) (*MsgProposeSpendResponse, error)
	ExecuteSpend(ctx context.Context, msg *MsgExecuteSpend) (*MsgExecuteSpendResponse, error)
	CancelSpend(ctx context.Context, msg *MsgCancelSpend) (*MsgCancelSpendResponse, error)
	SetBudget(ctx context.Context, msg *MsgSetBudget) (*MsgSetBudgetResponse, error)
	UpdateParams(ctx context.Context, msg *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}
