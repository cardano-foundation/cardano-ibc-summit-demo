package vesseloracle

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"vesseloracle/testutil/sample"
	vesseloraclesimulation "vesseloracle/x/vesseloracle/simulation"
	"vesseloracle/x/vesseloracle/types"
)

// avoid unused import issue
var (
	_ = vesseloraclesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateVessel = "op_weight_msg_vessel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVessel int = 100

	opWeightMsgUpdateVessel = "op_weight_msg_vessel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateVessel int = 100

	opWeightMsgDeleteVessel = "op_weight_msg_vessel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVessel int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vesseloracleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		VesselList: []types.Vessel{
			{
				Creator: sample.AccAddress(),
				Imo:     "0",
				Ts:      0,
			},
			{
				Creator: sample.AccAddress(),
				Imo:     "1",
				Ts:      1,
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vesseloracleGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateVessel int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateVessel, &weightMsgCreateVessel, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVessel = defaultWeightMsgCreateVessel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVessel,
		vesseloraclesimulation.SimulateMsgCreateVessel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateVessel int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateVessel, &weightMsgUpdateVessel, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVessel = defaultWeightMsgUpdateVessel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateVessel,
		vesseloraclesimulation.SimulateMsgUpdateVessel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteVessel int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteVessel, &weightMsgDeleteVessel, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVessel = defaultWeightMsgDeleteVessel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVessel,
		vesseloraclesimulation.SimulateMsgDeleteVessel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateVessel,
			defaultWeightMsgCreateVessel,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vesseloraclesimulation.SimulateMsgCreateVessel(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateVessel,
			defaultWeightMsgUpdateVessel,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vesseloraclesimulation.SimulateMsgUpdateVessel(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteVessel,
			defaultWeightMsgDeleteVessel,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vesseloraclesimulation.SimulateMsgDeleteVessel(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
