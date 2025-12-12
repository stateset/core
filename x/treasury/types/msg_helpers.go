package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgRecordReserve(authority string, snapshot ReserveSnapshot) *MsgRecordReserve {
	return &MsgRecordReserve{Authority: authority, Snapshot: snapshot}
}

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
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgExecuteSpend(authority string, proposalID uint64) *MsgExecuteSpend {
	return &MsgExecuteSpend{Authority: authority, ProposalID: proposalID}
}

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
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgCancelSpend(authority string, proposalID uint64, reason string) *MsgCancelSpend {
	return &MsgCancelSpend{Authority: authority, ProposalID: proposalID, Reason: reason}
}

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
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

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
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgUpdateParams(authority string, params TreasuryParams) *MsgUpdateParams {
	return &MsgUpdateParams{Authority: authority, Params: params}
}

func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	return m.Params.Validate()
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
