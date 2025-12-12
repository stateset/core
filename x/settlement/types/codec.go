package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInstantTransfer{}, "settlement/InstantTransfer", nil)
	cdc.RegisterConcrete(&MsgCreateEscrow{}, "settlement/CreateEscrow", nil)
	cdc.RegisterConcrete(&MsgReleaseEscrow{}, "settlement/ReleaseEscrow", nil)
	cdc.RegisterConcrete(&MsgRefundEscrow{}, "settlement/RefundEscrow", nil)
	cdc.RegisterConcrete(&MsgCreateBatch{}, "settlement/CreateBatch", nil)
	cdc.RegisterConcrete(&MsgSettleBatch{}, "settlement/SettleBatch", nil)
	cdc.RegisterConcrete(&MsgOpenChannel{}, "settlement/OpenChannel", nil)
	cdc.RegisterConcrete(&MsgCloseChannel{}, "settlement/CloseChannel", nil)
	cdc.RegisterConcrete(&MsgClaimChannel{}, "settlement/ClaimChannel", nil)
	cdc.RegisterConcrete(&MsgRegisterMerchant{}, "settlement/RegisterMerchant", nil)
	cdc.RegisterConcrete(&MsgUpdateMerchant{}, "settlement/UpdateMerchant", nil)
	cdc.RegisterConcrete(&MsgInstantCheckout{}, "settlement/InstantCheckout", nil)
	cdc.RegisterConcrete(&MsgPartialRefund{}, "settlement/PartialRefund", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
