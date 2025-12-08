package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var ModuleCdc = codec.NewLegacyAmino()

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePayment{}, "stateset/payments/MsgCreatePayment", nil)
	cdc.RegisterConcrete(&MsgSettlePayment{}, "stateset/payments/MsgSettlePayment", nil)
	cdc.RegisterConcrete(&MsgCancelPayment{}, "stateset/payments/MsgCancelPayment", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	_ = registry
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
