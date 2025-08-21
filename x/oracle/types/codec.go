package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitPriceFeed{}, "oracle/SubmitPriceFeed", nil)
	cdc.RegisterConcrete(&MsgRegisterOracle{}, "oracle/RegisterOracle", nil)
	cdc.RegisterConcrete(&MsgUpdateOracle{}, "oracle/UpdateOracle", nil)
	cdc.RegisterConcrete(&MsgRemoveOracle{}, "oracle/RemoveOracle", nil)
	cdc.RegisterConcrete(&MsgRequestPrice{}, "oracle/RequestPrice", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitPriceFeed{},
		&MsgRegisterOracle{},
		&MsgUpdateOracle{},
		&MsgRemoveOracle{},
		&MsgRequestPrice{},
	)
	
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	Amino.Seal()
}