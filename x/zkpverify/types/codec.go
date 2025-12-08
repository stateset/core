package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterCircuit{},
		&MsgDeactivateCircuit{},
		&MsgRegisterSymbolicRule{},
		&MsgSubmitProof{},
		&MsgChallengeProof{},
		&MsgUpdateParams{},
	)
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
