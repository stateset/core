package types

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgUpsertProfile = "upsert_profile"
	TypeMsgSetSanction   = "set_sanction"
)

var (
	_ sdk.Msg = (*MsgUpsertProfile)(nil)
	_ sdk.Msg = (*MsgSetSanction)(nil)
)

// MsgUpsertProfile allows compliance operators to create or update a profile.
type MsgUpsertProfile struct {
	Authority string  `json:"authority" yaml:"authority"`
	Profile   Profile `json:"profile" yaml:"profile"`
}

func (m *MsgUpsertProfile) Reset() { *m = MsgUpsertProfile{} }
func (m *MsgUpsertProfile) String() string {
	return fmt.Sprintf("MsgUpsertProfile{%s %s}", m.Authority, m.Profile.Address)
}
func (*MsgUpsertProfile) ProtoMessage() {}

func NewMsgUpsertProfile(authority string, profile Profile) *MsgUpsertProfile {
	return &MsgUpsertProfile{
		Authority: authority,
		Profile:   profile,
	}
}

func (m MsgUpsertProfile) Route() string { return RouterKey }
func (m MsgUpsertProfile) Type() string  { return TypeMsgUpsertProfile }

func (m MsgUpsertProfile) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidAddress, fmt.Sprintf("authority: %v", err))
	}
	return m.Profile.ValidateBasic()
}

func (m MsgUpsertProfile) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgUpsertProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// MsgUpsertProfileResponse indicates profile update success.
type MsgUpsertProfileResponse struct{}

// MsgSetSanction toggles the sanction flag for a profile.
type MsgSetSanction struct {
	Authority string `json:"authority" yaml:"authority"`
	Address   string `json:"address" yaml:"address"`
	Sanction  bool   `json:"sanction" yaml:"sanction"`
	Reason    string `json:"reason" yaml:"reason"`
}

func (m *MsgSetSanction) Reset() { *m = MsgSetSanction{} }
func (m *MsgSetSanction) String() string {
	return fmt.Sprintf("MsgSetSanction{%s %s %t}", m.Authority, m.Address, m.Sanction)
}
func (*MsgSetSanction) ProtoMessage() {}

func NewMsgSetSanction(authority, address string, sanction bool, reason string) *MsgSetSanction {
	return &MsgSetSanction{
		Authority: authority,
		Address:   address,
		Sanction:  sanction,
		Reason:    reason,
	}
}

func (m MsgSetSanction) Route() string { return RouterKey }
func (m MsgSetSanction) Type() string  { return TypeMsgSetSanction }

func (m MsgSetSanction) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidAddress, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errorsmod.Wrap(ErrInvalidAddress, err.Error())
	}
	return nil
}

func (m MsgSetSanction) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgSetSanction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// MsgSetSanctionResponse indicates sanction flag update.
type MsgSetSanctionResponse struct{}

// MsgServer exposes the gRPC Msg service contract for future protobuf migration.
type MsgServer interface {
	UpsertProfile(ctx context.Context, msg *MsgUpsertProfile) (*MsgUpsertProfileResponse, error)
	SetSanction(ctx context.Context, msg *MsgSetSanction) (*MsgSetSanctionResponse, error)
}
