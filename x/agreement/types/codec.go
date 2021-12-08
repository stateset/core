package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgActivateAgreement{}, "agreement/ActivateAgreement", nil)
	cdc.RegisterConcrete(&MsgExpireAgreement{}, "agreement/ExpireAgreement", nil)
	cdc.RegisterConcrete(&MsgRenewAgreement{}, "agreement/RenewAgreement", nil)
	cdc.RegisterConcrete(&MsgTerminateAgreement{}, "agreement/TerminateAgreement", nil)
	cdc.RegisterConcrete(&MsgCreateSentAgreement{}, "agreement/CreateSentAgreement", nil)
	cdc.RegisterConcrete(&MsgUpdateSentAgreement{}, "agreement/UpdateSentAgreement", nil)
	cdc.RegisterConcrete(&MsgDeleteSentAgreement{}, "agreement/DeleteSentAgreement", nil)
	cdc.RegisterConcrete(&MsgCreateTimedoutAgreement{}, "agreement/CreateTimedoutAgreement", nil)
	cdc.RegisterConcrete(&MsgUpdateTimedoutAgreement{}, "agreement/UpdateTimedoutAgreement", nil)
	cdc.RegisterConcrete(&MsgDeleteTimedoutAgreement{}, "agreement/DeleteTimedoutAgreement", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgActivateAgreement{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExpireAgreement{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRenewAgreement{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTerminateAgreement{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSentAgreement{},
		&MsgUpdateSentAgreement{},
		&MsgDeleteSentAgreement{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTimedoutAgreement{},
		&MsgUpdateTimedoutAgreement{},
		&MsgDeleteTimedoutAgreement{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
