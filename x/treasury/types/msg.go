package types

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgRecordReserve = "record_reserve"

var _ sdk.Msg = (*MsgRecordReserve)(nil)

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

type MsgServer interface {
	RecordReserve(ctx context.Context, msg *MsgRecordReserve) (*MsgRecordReserveResponse, error)
}
