package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateStablecoin{}, "stablecoins/CreateStablecoin", nil)
	cdc.RegisterConcrete(&MsgUpdateStablecoin{}, "stablecoins/UpdateStablecoin", nil)
	cdc.RegisterConcrete(&MsgMintStablecoin{}, "stablecoins/MintStablecoin", nil)
	cdc.RegisterConcrete(&MsgBurnStablecoin{}, "stablecoins/BurnStablecoin", nil)
	cdc.RegisterConcrete(&MsgPauseStablecoin{}, "stablecoins/PauseStablecoin", nil)
	cdc.RegisterConcrete(&MsgUnpauseStablecoin{}, "stablecoins/UnpauseStablecoin", nil)
	cdc.RegisterConcrete(&MsgUpdatePriceData{}, "stablecoins/UpdatePriceData", nil)
	cdc.RegisterConcrete(&MsgUpdateReserves{}, "stablecoins/UpdateReserves", nil)
	cdc.RegisterConcrete(&MsgWhitelistAddress{}, "stablecoins/WhitelistAddress", nil)
	cdc.RegisterConcrete(&MsgBlacklistAddress{}, "stablecoins/BlacklistAddress", nil)
	cdc.RegisterConcrete(&MsgRemoveFromWhitelist{}, "stablecoins/RemoveFromWhitelist", nil)
	cdc.RegisterConcrete(&MsgRemoveFromBlacklist{}, "stablecoins/RemoveFromBlacklist", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStablecoin{},
		&MsgUpdateStablecoin{},
		&MsgMintStablecoin{},
		&MsgBurnStablecoin{},
		&MsgPauseStablecoin{},
		&MsgUnpauseStablecoin{},
		&MsgUpdatePriceData{},
		&MsgUpdateReserves{},
		&MsgWhitelistAddress{},
		&MsgBlacklistAddress{},
		&MsgRemoveFromWhitelist{},
		&MsgRemoveFromBlacklist{},
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