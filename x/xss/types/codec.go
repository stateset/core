package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
	groupcodec "github.com/cosmos/cosmos-sdk/x/group/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/xss interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStakeTokens{}, "xss/MsgStakeTokens", nil)
	cdc.RegisterConcrete(&MsgUnstakeTokens{}, "xss/MsgUnstakeTokens", nil)
	cdc.RegisterConcrete(&MsgWithdrawStakingRewards{}, "xss/MsgWithdrawStakingRewards", nil)
	cdc.RegisterConcrete(&MsgCreateValidator{}, "xss/MsgCreateValidator", nil)
	cdc.RegisterConcrete(&MsgEditValidator{}, "xss/MsgEditValidator", nil)
	cdc.RegisterConcrete(&MsgUnjailValidator{}, "xss/MsgUnjailValidator", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "xss/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgSlashValidator{}, "xss/MsgSlashValidator", nil)
	cdc.RegisterConcrete(&MsgExecuteAgent{}, "xss/MsgExecuteAgent", nil)
	cdc.RegisterConcrete(&MsgBurnTokens{}, "xss/MsgBurnTokens", nil)
}

// RegisterInterfaces registers the x/xss interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStakeTokens{},
		&MsgUnstakeTokens{},
		&MsgWithdrawStakingRewards{},
		&MsgCreateValidator{},
		&MsgEditValidator{},
		&MsgUnjailValidator{},
		&MsgUpdateParams{},
		&MsgSlashValidator{},
		&MsgExecuteAgent{},
		&MsgBurnTokens{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()

	// Register all Amino interfaces and concrete types on the authz and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}