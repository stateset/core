package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec registers the account types and interface.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateOrder{}, "orders/CreateOrder", nil)
	cdc.RegisterConcrete(&MsgConfirmOrder{}, "orders/ConfirmOrder", nil)
	cdc.RegisterConcrete(&MsgPayOrder{}, "orders/PayOrder", nil)
	cdc.RegisterConcrete(&MsgShipOrder{}, "orders/ShipOrder", nil)
	cdc.RegisterConcrete(&MsgDeliverOrder{}, "orders/DeliverOrder", nil)
	cdc.RegisterConcrete(&MsgCompleteOrder{}, "orders/CompleteOrder", nil)
	cdc.RegisterConcrete(&MsgCancelOrder{}, "orders/CancelOrder", nil)
	cdc.RegisterConcrete(&MsgRefundOrder{}, "orders/RefundOrder", nil)
	cdc.RegisterConcrete(&MsgOpenDispute{}, "orders/OpenDispute", nil)
	cdc.RegisterConcrete(&MsgResolveDispute{}, "orders/ResolveDispute", nil)
}

// RegisterInterfaces registers the x/orders interfaces types with the interface registry.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrder{},
		&MsgConfirmOrder{},
		&MsgPayOrder{},
		&MsgShipOrder{},
		&MsgDeliverOrder{},
		&MsgCompleteOrder{},
		&MsgCancelOrder{},
		&MsgRefundOrder{},
		&MsgOpenDispute{},
		&MsgResolveDispute{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
