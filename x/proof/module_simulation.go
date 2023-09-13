package proof

import (
	"math/rand"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stateset/core/testutil/sample"
	proofsimulation "github.com/stateset/core/x/proof/simulation"
	"github.com/stateset/core/x/proof/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = proofsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProof = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProof int = 100

	opWeightMsgVerifyProof = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyProof int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	proofGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&proofGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProof int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProof, &weightMsgCreateProof, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProof = defaultWeightMsgCreateProof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProof,
		proofsimulation.SimulateMsgCreateProof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyProof int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVerifyProof, &weightMsgVerifyProof, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyProof = defaultWeightMsgVerifyProof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyProof,
		proofsimulation.SimulateMsgVerifyProof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
