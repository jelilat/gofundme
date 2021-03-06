package gofundme

import (
	"math/rand"

	"github.com/cosmonaut/gofundme/testutil/sample"
	gofundmesimulation "github.com/cosmonaut/gofundme/x/gofundme/simulation"
	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = gofundmesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateGofundme = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGofundme int = 100

	opWeightMsgDonateFund = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDonateFund int = 100

	opWeightMsgWithdrawDonation = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawDonation int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	gofundmeGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gofundmeGenesis)
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

	var weightMsgCreateGofundme int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGofundme, &weightMsgCreateGofundme, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGofundme = defaultWeightMsgCreateGofundme
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGofundme,
		gofundmesimulation.SimulateMsgCreateGofundme(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDonateFund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDonateFund, &weightMsgDonateFund, nil,
		func(_ *rand.Rand) {
			weightMsgDonateFund = defaultWeightMsgDonateFund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDonateFund,
		gofundmesimulation.SimulateMsgDonateFund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawDonation int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawDonation, &weightMsgWithdrawDonation, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawDonation = defaultWeightMsgWithdrawDonation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawDonation,
		gofundmesimulation.SimulateMsgWithdrawDonation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
