package refund

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stateset/core/testutil/sample"
	refundsimulation "github.com/stateset/core/x/refund/simulation"
	"github.com/stateset/core/x/refund/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = refundsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestRefund = "op_weight_msg_request_refund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestRefund int = 100

	opWeightMsgApproveRefund = "op_weight_msg_approve_refund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveRefund int = 100

	opWeightMsgRejectRefund = "op_weight_msg_reject_refund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRejectRefund int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	refundGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&refundGenesis)
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

	var weightMsgRequestRefund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestRefund, &weightMsgRequestRefund, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRefund = defaultWeightMsgRequestRefund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestRefund,
		refundsimulation.SimulateMsgRequestRefund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveRefund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveRefund, &weightMsgApproveRefund, nil,
		func(_ *rand.Rand) {
			weightMsgApproveRefund = defaultWeightMsgApproveRefund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveRefund,
		refundsimulation.SimulateMsgApproveRefund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRejectRefund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRejectRefund, &weightMsgRejectRefund, nil,
		func(_ *rand.Rand) {
			weightMsgRejectRefund = defaultWeightMsgRejectRefund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRejectRefund,
		refundsimulation.SimulateMsgRejectRefund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
