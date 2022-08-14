package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgFactorInvoice{}, "invoice/FactorInvoice", nil)
	cdc.RegisterConcrete(&MsgCreateSentInvoice{}, "invoice/CreateSentInvoice", nil)
	cdc.RegisterConcrete(&MsgUpdateSentInvoice{}, "invoice/UpdateSentInvoice", nil)
	cdc.RegisterConcrete(&MsgDeleteSentInvoice{}, "invoice/DeleteSentInvoice", nil)
	cdc.RegisterConcrete(&MsgCreateTimedoutInvoice{}, "invoice/CreateTimedoutInvoice", nil)
	cdc.RegisterConcrete(&MsgUpdateTimedoutInvoice{}, "invoice/UpdateTimedoutInvoice", nil)
	cdc.RegisterConcrete(&MsgDeleteTimedoutInvoice{}, "invoice/DeleteTimedoutInvoice", nil)
	cdc.RegisterConcrete(&MsgCreateInvoice{}, "invoice/CreateInvoice", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFactorInvoice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSentInvoice{},
		&MsgUpdateSentInvoice{},
		&MsgDeleteSentInvoice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTimedoutInvoice{},
		&MsgUpdateTimedoutInvoice{},
		&MsgDeleteTimedoutInvoice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInvoice{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
