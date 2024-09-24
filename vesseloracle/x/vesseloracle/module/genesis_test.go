package vesseloracle_test

import (
	"testing"

	keepertest "vesseloracle/testutil/keeper"
	"vesseloracle/testutil/nullify"
	vesseloracle "vesseloracle/x/vesseloracle/module"
	"vesseloracle/x/vesseloracle/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		VesselList: []types.Vessel{
			{
				Imo: "0",
				Ts:  0,
			},
			{
				Imo: "1",
				Ts:  1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VesseloracleKeeper(t)
	vesseloracle.InitGenesis(ctx, k, genesisState)
	got := vesseloracle.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VesselList, got.VesselList)
	// this line is used by starport scaffolding # genesis/test/assert
}
