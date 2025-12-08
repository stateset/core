package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Ensure message types implement sdk.Msg interface
var (
	_ sdk.Msg = (*MsgPauseSystem)(nil)
	_ sdk.Msg = (*MsgResumeSystem)(nil)
	_ sdk.Msg = (*MsgTripCircuit)(nil)
	_ sdk.Msg = (*MsgResetCircuit)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
)

// MsgPauseSystem pauses the entire system
type MsgPauseSystem struct {
	Authority string `json:"authority" protobuf:"bytes,1,opt,name=authority,proto3"`
	Reason    string `json:"reason" protobuf:"bytes,2,opt,name=reason,proto3"`
	// DurationSeconds is optional; 0 means indefinite until manual resume
	DurationSeconds int64 `json:"duration_seconds,omitempty" protobuf:"varint,3,opt,name=duration_seconds,proto3"`
}

func (m *MsgPauseSystem) Reset()         { *m = MsgPauseSystem{} }
func (m *MsgPauseSystem) String() string { return "" }
func (m *MsgPauseSystem) ProtoMessage()  {}

func NewMsgPauseSystem(authority, reason string, duration int64) *MsgPauseSystem {
	return &MsgPauseSystem{
		Authority:       authority,
		Reason:          reason,
		DurationSeconds: duration,
	}
}

func (msg MsgPauseSystem) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if msg.Reason == "" {
		return errorsmod.Wrap(ErrInvalidParams, "reason is required")
	}
	if msg.DurationSeconds < 0 {
		return errorsmod.Wrap(ErrInvalidDuration, "duration cannot be negative")
	}
	return nil
}

func (msg MsgPauseSystem) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// MsgResumeSystem resumes the system from pause
type MsgResumeSystem struct {
	Authority string `json:"authority" protobuf:"bytes,1,opt,name=authority,proto3"`
}

func (m *MsgResumeSystem) Reset()         { *m = MsgResumeSystem{} }
func (m *MsgResumeSystem) String() string { return "" }
func (m *MsgResumeSystem) ProtoMessage()  {}

func NewMsgResumeSystem(authority string) *MsgResumeSystem {
	return &MsgResumeSystem{Authority: authority}
}

func (msg MsgResumeSystem) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	return nil
}

func (msg MsgResumeSystem) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// MsgTripCircuit trips a specific module's circuit breaker
type MsgTripCircuit struct {
	Authority  string `json:"authority" protobuf:"bytes,1,opt,name=authority,proto3"`
	ModuleName string `json:"module_name" protobuf:"bytes,2,opt,name=module_name,proto3"`
	Reason     string `json:"reason" protobuf:"bytes,3,opt,name=reason,proto3"`
	// DisableMessages are specific message types to disable (empty = all)
	DisableMessages []string `json:"disable_messages,omitempty" protobuf:"bytes,4,rep,name=disable_messages,proto3"`
}

func (m *MsgTripCircuit) Reset()         { *m = MsgTripCircuit{} }
func (m *MsgTripCircuit) String() string { return "" }
func (m *MsgTripCircuit) ProtoMessage()  {}

func NewMsgTripCircuit(authority, moduleName, reason string, disableMessages []string) *MsgTripCircuit {
	return &MsgTripCircuit{
		Authority:       authority,
		ModuleName:      moduleName,
		Reason:          reason,
		DisableMessages: disableMessages,
	}
}

func (msg MsgTripCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if msg.ModuleName == "" {
		return errorsmod.Wrap(ErrInvalidParams, "module name is required")
	}
	if msg.Reason == "" {
		return errorsmod.Wrap(ErrInvalidParams, "reason is required")
	}
	return nil
}

func (msg MsgTripCircuit) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// MsgResetCircuit resets a module's circuit breaker
type MsgResetCircuit struct {
	Authority  string `json:"authority" protobuf:"bytes,1,opt,name=authority,proto3"`
	ModuleName string `json:"module_name" protobuf:"bytes,2,opt,name=module_name,proto3"`
}

func (m *MsgResetCircuit) Reset()         { *m = MsgResetCircuit{} }
func (m *MsgResetCircuit) String() string { return "" }
func (m *MsgResetCircuit) ProtoMessage()  {}

func NewMsgResetCircuit(authority, moduleName string) *MsgResetCircuit {
	return &MsgResetCircuit{
		Authority:  authority,
		ModuleName: moduleName,
	}
}

func (msg MsgResetCircuit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	if msg.ModuleName == "" {
		return errorsmod.Wrap(ErrInvalidParams, "module name is required")
	}
	return nil
}

func (msg MsgResetCircuit) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// MsgUpdateParams updates module parameters
type MsgUpdateParams struct {
	Authority string `json:"authority" protobuf:"bytes,1,opt,name=authority,proto3"`
	Params    Params `json:"params" protobuf:"bytes,2,opt,name=params,proto3"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return "" }
func (m *MsgUpdateParams) ProtoMessage()  {}

func NewMsgUpdateParams(authority string, params Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}
}

func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrap(ErrUnauthorized, "invalid authority address")
	}
	return msg.Params.Validate()
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}
