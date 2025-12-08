package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var ModuleCdc = codec.NewLegacyAmino()

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRecordReserve{}, "stateset/treasury/MsgRecordReserve", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	_ = registry
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
