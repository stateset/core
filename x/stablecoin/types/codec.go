package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var ModuleCdc = codec.NewLegacyAmino()

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateVault{}, "stateset/stablecoin/MsgCreateVault", nil)
	cdc.RegisterConcrete(&MsgDepositCollateral{}, "stateset/stablecoin/MsgDepositCollateral", nil)
	cdc.RegisterConcrete(&MsgWithdrawCollateral{}, "stateset/stablecoin/MsgWithdrawCollateral", nil)
	cdc.RegisterConcrete(&MsgMintStablecoin{}, "stateset/stablecoin/MsgMintStablecoin", nil)
	cdc.RegisterConcrete(&MsgRepayStablecoin{}, "stateset/stablecoin/MsgRepayStablecoin", nil)
	cdc.RegisterConcrete(&MsgLiquidateVault{}, "stateset/stablecoin/MsgLiquidateVault", nil)
	cdc.RegisterConcrete(&MsgDepositReserve{}, "stateset/stablecoin/MsgDepositReserve", nil)
	cdc.RegisterConcrete(&MsgRequestRedemption{}, "stateset/stablecoin/MsgRequestRedemption", nil)
	cdc.RegisterConcrete(&MsgExecuteRedemption{}, "stateset/stablecoin/MsgExecuteRedemption", nil)
	cdc.RegisterConcrete(&MsgCancelRedemption{}, "stateset/stablecoin/MsgCancelRedemption", nil)
	cdc.RegisterConcrete(&MsgUpdateReserveParams{}, "stateset/stablecoin/MsgUpdateReserveParams", nil)
	cdc.RegisterConcrete(&MsgRecordAttestation{}, "stateset/stablecoin/MsgRecordAttestation", nil)
	cdc.RegisterConcrete(&MsgSetApprovedAttester{}, "stateset/stablecoin/MsgSetApprovedAttester", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	ModuleCdc.Seal()
}
