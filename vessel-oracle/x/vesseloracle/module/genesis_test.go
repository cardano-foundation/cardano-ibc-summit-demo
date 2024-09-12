package vesseloracle_test

import (
	"testing"

	keepertest "vessel-oracle/testutil/keeper"
	"vessel-oracle/testutil/nullify"
	vesseloracle "vessel-oracle/x/vesseloracle/module"
	"vessel-oracle/x/vesseloracle/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VesseloracleKeeper(t)
	vesseloracle.InitGenesis(ctx, k, genesisState)
	got := vesseloracle.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
