package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	// Register messages
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStakeTokens{},
		&MsgUnstakeTokens{},
		&MsgClaimStakingRewards{},
		&MsgSubmitProposal{},
		&MsgVote{},
		&MsgClaimVestedTokens{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}