package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgPauseSystem(authority, reason string, duration int64) *MsgPauseSystem {
	return &MsgPauseSystem{
		Authority:       authority,
		Reason:          reason,
		DurationSeconds: duration,
	}
}

func (m MsgPauseSystem) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if m.Reason == "" {
		return errorsmod.Wrap(ErrInvalidParams, "reason is required")
	}
	if m.DurationSeconds < 0 {
		return errorsmod.Wrap(ErrInvalidDuration, "duration cannot be negative")
	}
	return nil
}

func (m MsgPauseSystem) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgResumeSystem(authority string) *MsgResumeSystem {
	return &MsgResumeSystem{Authority: authority}
}

func (m MsgResumeSystem) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	return nil
}

func (m MsgResumeSystem) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgTripCircuit(authority, moduleName, reason string, disableMessages []string) *MsgTripCircuit {
	return &MsgTripCircuit{
		Authority:       authority,
		ModuleName:      moduleName,
		Reason:          reason,
		DisableMessages: disableMessages,
	}
}

func (m MsgTripCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if m.ModuleName == "" {
		return errorsmod.Wrap(ErrInvalidParams, "module name is required")
	}
	if m.Reason == "" {
		return errorsmod.Wrap(ErrInvalidParams, "reason is required")
	}
	return nil
}

func (m MsgTripCircuit) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgResetCircuit(authority, moduleName string) *MsgResetCircuit {
	return &MsgResetCircuit{
		Authority:  authority,
		ModuleName: moduleName,
	}
}

func (m MsgResetCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if m.ModuleName == "" {
		return errorsmod.Wrap(ErrInvalidParams, "module name is required")
	}
	return nil
}

func (m MsgResetCircuit) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgUpdateParams(authority string, params Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}
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
