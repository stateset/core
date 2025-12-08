package types

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdatePrice = "update_price"

var _ sdk.Msg = (*MsgUpdatePrice)(nil)

// MsgUpdatePrice defines a message for updating the price of a denom.
type MsgUpdatePrice struct {
	Authority string            `json:"authority" yaml:"authority"`
	Denom     string            `json:"denom" yaml:"denom"`
	Price     sdkmath.LegacyDec `json:"price" yaml:"price"`
}

func (m *MsgUpdatePrice) Reset() { *m = MsgUpdatePrice{} }
func (m *MsgUpdatePrice) String() string {
	return fmt.Sprintf("MsgUpdatePrice{%s %s %s}", m.Authority, m.Denom, m.Price.String())
}
func (*MsgUpdatePrice) ProtoMessage() {}

// MsgUpdatePriceResponse is returned on successful price update.
type MsgUpdatePriceResponse struct{}

// NewMsgUpdatePrice creates a new MsgUpdatePrice instance.
func NewMsgUpdatePrice(authority, denom string, price sdkmath.LegacyDec) *MsgUpdatePrice {
	return &MsgUpdatePrice{
		Authority: authority,
		Denom:     denom,
		Price:     price,
	}
}

func (m MsgUpdatePrice) Route() string { return RouterKey }

func (m MsgUpdatePrice) Type() string { return TypeMsgUpdatePrice }

func (m MsgUpdatePrice) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgUpdatePrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpdatePrice) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrapf(ErrInvalidAuthority, "authority %s", err)
	}
	if len(m.Denom) == 0 {
		return errorsmod.Wrap(ErrInvalidDenom, "denom cannot be empty")
	}
	if !m.Price.IsPositive() {
		return errorsmod.Wrap(ErrInvalidPrice, "price must be positive")
	}
	return nil
}

// MsgServer defines the gRPC interface for oracle messages. We keep it lightweight
// to ease eventual protobuf migration.
type MsgServer interface {
	UpdatePrice(ctx context.Context, msg *MsgUpdatePrice) (*MsgUpdatePriceResponse, error)
}
