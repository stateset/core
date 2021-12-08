package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgFinancePurchaseorder{}, "purchaseorder/FinancePurchaseorder", nil)
	cdc.RegisterConcrete(&MsgCancelPurchaseorder{}, "purchaseorder/CancelPurchaseorder", nil)
	cdc.RegisterConcrete(&MsgCompletePurchaseorder{}, "purchaseorder/CompletePurchaseorder", nil)
	cdc.RegisterConcrete(&MsgCreateSentPurchaseorder{}, "purchaseorder/CreateSentPurchaseorder", nil)
	cdc.RegisterConcrete(&MsgUpdateSentPurchaseorder{}, "purchaseorder/UpdateSentPurchaseorder", nil)
	cdc.RegisterConcrete(&MsgDeleteSentPurchaseorder{}, "purchaseorder/DeleteSentPurchaseorder", nil)
	cdc.RegisterConcrete(&MsgCreateTimedoutPurchaseorder{}, "purchaseorder/CreateTimedoutPurchaseorder", nil)
	cdc.RegisterConcrete(&MsgUpdateTimedoutPurchaseorder{}, "purchaseorder/UpdateTimedoutPurchaseorder", nil)
	cdc.RegisterConcrete(&MsgDeleteTimedoutPurchaseorder{}, "purchaseorder/DeleteTimedoutPurchaseorder", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinancePurchaseorder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelPurchaseorder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCompletePurchaseorder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSentPurchaseorder{},
		&MsgUpdateSentPurchaseorder{},
		&MsgDeleteSentPurchaseorder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTimedoutPurchaseorder{},
		&MsgUpdateTimedoutPurchaseorder{},
		&MsgDeleteTimedoutPurchaseorder{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
