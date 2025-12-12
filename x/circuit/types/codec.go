package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types
// on the provided LegacyAmino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPauseSystem{}, "circuit/MsgPauseSystem", nil)
	cdc.RegisterConcrete(&MsgResumeSystem{}, "circuit/MsgResumeSystem", nil)
	cdc.RegisterConcrete(&MsgTripCircuit{}, "circuit/MsgTripCircuit", nil)
	cdc.RegisterConcrete(&MsgResetCircuit{}, "circuit/MsgResetCircuit", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "circuit/MsgUpdateParams", nil)
}

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
