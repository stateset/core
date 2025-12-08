package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// ModuleCdc references the global module codec.
var ModuleCdc = codec.NewLegacyAmino()

// RegisterLegacyAminoCodec registers concrete types on the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdatePrice{}, "stateset/oracle/MsgUpdatePrice", nil)
}

// RegisterInterfaces registers module interfaces to protobuf type registry.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// Note: Full interface registration requires protobuf-generated code.
	// Legacy amino registration handles basic functionality.
	_ = registry
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
