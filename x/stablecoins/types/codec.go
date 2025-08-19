package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateStablecoin{}, "stablecoins/CreateStablecoin", nil)
	cdc.RegisterConcrete(&MsgUpdateStablecoin{}, "stablecoins/UpdateStablecoin", nil)
	cdc.RegisterConcrete(&MsgMintStablecoin{}, "stablecoins/MintStablecoin", nil)
	cdc.RegisterConcrete(&MsgBurnStablecoin{}, "stablecoins/BurnStablecoin", nil)
	cdc.RegisterConcrete(&MsgPauseStablecoin{}, "stablecoins/PauseStablecoin", nil)
	// TODO: Define these message types
	// cdc.RegisterConcrete(&MsgUnpauseStablecoin{}, "stablecoins/UnpauseStablecoin", nil)
	// cdc.RegisterConcrete(&MsgUpdatePriceData{}, "stablecoins/UpdatePriceData", nil)
	// cdc.RegisterConcrete(&MsgUpdateReserves{}, "stablecoins/UpdateReserves", nil)
	// cdc.RegisterConcrete(&MsgWhitelistAddress{}, "stablecoins/WhitelistAddress", nil)
	// cdc.RegisterConcrete(&MsgBlacklistAddress{}, "stablecoins/BlacklistAddress", nil)
	// cdc.RegisterConcrete(&MsgRemoveFromWhitelist{}, "stablecoins/RemoveFromWhitelist", nil)
	// cdc.RegisterConcrete(&MsgRemoveFromBlacklist{}, "stablecoins/RemoveFromBlacklist", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStablecoin{},
		&MsgUpdateStablecoin{},
		&MsgMintStablecoin{},
		&MsgBurnStablecoin{},
		&MsgPauseStablecoin{},
		// TODO: Define these message types
		// &MsgUnpauseStablecoin{},
		// &MsgUpdatePriceData{},
		// &MsgUpdateReserves{},
		// &MsgWhitelistAddress{},
		// &MsgBlacklistAddress{},
		// &MsgRemoveFromWhitelist{},
		// &MsgRemoveFromBlacklist{},
	)

	// TODO: Uncomment when _Msg_serviceDesc is generated
	// msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}