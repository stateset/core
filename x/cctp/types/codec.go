package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDepositForBurn{}, "cctp/MsgDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgDepositForBurnWithCaller{}, "cctp/MsgDepositForBurnWithCaller", nil)
	cdc.RegisterConcrete(&MsgReceiveMessage{}, "cctp/MsgReceiveMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessage{}, "cctp/MsgSendMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessageWithCaller{}, "cctp/MsgSendMessageWithCaller", nil)
	cdc.RegisterConcrete(&MsgReplaceDepositForBurn{}, "cctp/MsgReplaceDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgReplaceMessage{}, "cctp/MsgReplaceMessage", nil)
	cdc.RegisterConcrete(&MsgAcceptOwner{}, "cctp/MsgAcceptOwner", nil)
	cdc.RegisterConcrete(&MsgAddRemoteTokenMessenger{}, "cctp/MsgAddRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgDisableAttester{}, "cctp/MsgDisableAttester", nil)
	cdc.RegisterConcrete(&MsgEnableAttester{}, "cctp/MsgEnableAttester", nil)
	cdc.RegisterConcrete(&MsgLinkTokenPair{}, "cctp/MsgLinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgPauseBurningAndMinting{}, "cctp/MsgPauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgPauseSendingAndReceivingMessages{}, "cctp/MsgPauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgRemoveRemoteTokenMessenger{}, "cctp/MsgRemoveRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgSetMaxBurnAmountPerMessage{}, "cctp/MsgSetMaxBurnAmountPerMessage", nil)
	cdc.RegisterConcrete(&MsgUnlinkTokenPair{}, "cctp/MsgUnlinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgUnpauseBurningAndMinting{}, "cctp/MsgUnpauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgUnpauseSendingAndReceivingMessages{}, "cctp/MsgUnpauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgUpdateAttesterManager{}, "cctp/MsgUpdateAttesterManager", nil)
	cdc.RegisterConcrete(&MsgUpdateMaxMessageBodySize{}, "cctp/MsgUpdateMaxMessageBodySize", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "cctp/MsgUpdateOwner", nil)
	cdc.RegisterConcrete(&MsgUpdatePauser{}, "cctp/MsgUpdatePauser", nil)
	cdc.RegisterConcrete(&MsgUpdateSignatureThreshold{}, "cctp/MsgUpdateSignatureThreshold", nil)
	cdc.RegisterConcrete(&MsgUpdateTokenController{}, "cctp/MsgUpdateTokenController", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositForBurn{},
		&MsgDepositForBurnWithCaller{},
		&MsgReceiveMessage{},
		&MsgSendMessage{},
		&MsgSendMessageWithCaller{},
		&MsgReplaceDepositForBurn{},
		&MsgReplaceMessage{},
		&MsgAcceptOwner{},
		&MsgAddRemoteTokenMessenger{},
		&MsgDisableAttester{},
		&MsgEnableAttester{},
		&MsgLinkTokenPair{},
		&MsgPauseBurningAndMinting{},
		&MsgPauseSendingAndReceivingMessages{},
		&MsgRemoveRemoteTokenMessenger{},
		&MsgSetMaxBurnAmountPerMessage{},
		&MsgUnlinkTokenPair{},
		&MsgUnpauseBurningAndMinting{},
		&MsgUnpauseSendingAndReceivingMessages{},
		&MsgUpdateAttesterManager{},
		&MsgUpdateMaxMessageBodySize{},
		&MsgUpdateOwner{},
		&MsgUpdatePauser{},
		&MsgUpdateSignatureThreshold{},
		&MsgUpdateTokenController{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}