package vesseloracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"vesseloracle/x/vesseloracle/keeper"
	"vesseloracle/x/vesseloracle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the vessel
	for _, elem := range genState.VesselList {
		k.SetVessel(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.VesselList = k.GetAllVessel(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
