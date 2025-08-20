package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateOrder{}, "orders/CreateOrder", nil)
	cdc.RegisterConcrete(&MsgUpdateOrder{}, "orders/UpdateOrder", nil)
	cdc.RegisterConcrete(&MsgCancelOrder{}, "orders/CancelOrder", nil)
	cdc.RegisterConcrete(&MsgFulfillOrder{}, "orders/FulfillOrder", nil)
	cdc.RegisterConcrete(&MsgRefundOrder{}, "orders/RefundOrder", nil)
	cdc.RegisterConcrete(&MsgUpdateOrderStatus{}, "orders/UpdateOrderStatus", nil)
	cdc.RegisterConcrete(&MsgPayWithStablecoin{}, "orders/PayWithStablecoin", nil)
	cdc.RegisterConcrete(&MsgConfirmStablecoinPayment{}, "orders/ConfirmStablecoinPayment", nil)
	cdc.RegisterConcrete(&MsgRefundStablecoinPayment{}, "orders/RefundStablecoinPayment", nil)
	cdc.RegisterConcrete(&MsgReleaseEscrow{}, "orders/ReleaseEscrow", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrder{},
		&MsgUpdateOrder{},
		&MsgCancelOrder{},
		&MsgFulfillOrder{},
		&MsgRefundOrder{},
		&MsgUpdateOrderStatus{},
		&MsgPayWithStablecoin{},
		&MsgConfirmStablecoinPayment{},
		&MsgRefundStablecoinPayment{},
		&MsgReleaseEscrow{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)