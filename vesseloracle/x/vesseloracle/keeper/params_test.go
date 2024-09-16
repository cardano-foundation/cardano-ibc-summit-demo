package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "vesseloracle/testutil/keeper"
	"vesseloracle/x/vesseloracle/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VesseloracleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
