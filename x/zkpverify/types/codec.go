package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var ModuleCdc = codec.NewLegacyAmino()

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterCircuit{}, "zkpverify/MsgRegisterCircuit", nil)
	cdc.RegisterConcrete(&MsgDeactivateCircuit{}, "zkpverify/MsgDeactivateCircuit", nil)
	cdc.RegisterConcrete(&MsgRegisterSymbolicRule{}, "zkpverify/MsgRegisterSymbolicRule", nil)
	cdc.RegisterConcrete(&MsgSubmitProof{}, "zkpverify/MsgSubmitProof", nil)
	cdc.RegisterConcrete(&MsgChallengeProof{}, "zkpverify/MsgChallengeProof", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "zkpverify/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
