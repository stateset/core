package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var ModuleCdc = codec.NewLegacyAmino()

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateVault{}, "stateset/stablecoin/MsgCreateVault", nil)
	cdc.RegisterConcrete(&MsgDepositCollateral{}, "stateset/stablecoin/MsgDepositCollateral", nil)
	cdc.RegisterConcrete(&MsgWithdrawCollateral{}, "stateset/stablecoin/MsgWithdrawCollateral", nil)
	cdc.RegisterConcrete(&MsgMintStablecoin{}, "stateset/stablecoin/MsgMintStablecoin", nil)
	cdc.RegisterConcrete(&MsgRepayStablecoin{}, "stateset/stablecoin/MsgRepayStablecoin", nil)
	cdc.RegisterConcrete(&MsgLiquidateVault{}, "stateset/stablecoin/MsgLiquidateVault", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// Note: In SDK 0.53+, messages need proper protobuf definitions for interface registration.
	// For now, using legacy amino registration is sufficient for basic functionality.
	// Full protobuf support requires generating .pb.go files from proto definitions.
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
