package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var ModuleCdc = codec.NewLegacyAmino()

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRecordReserve{}, "stateset/treasury/MsgRecordReserve", nil)
	cdc.RegisterConcrete(&MsgProposeSpend{}, "stateset/treasury/MsgProposeSpend", nil)
	cdc.RegisterConcrete(&MsgExecuteSpend{}, "stateset/treasury/MsgExecuteSpend", nil)
	cdc.RegisterConcrete(&MsgCancelSpend{}, "stateset/treasury/MsgCancelSpend", nil)
	cdc.RegisterConcrete(&MsgSetBudget{}, "stateset/treasury/MsgSetBudget", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "stateset/treasury/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
