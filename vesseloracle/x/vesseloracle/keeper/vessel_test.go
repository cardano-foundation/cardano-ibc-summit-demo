package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "vesseloracle/testutil/keeper"
	"vesseloracle/testutil/nullify"
	"vesseloracle/x/vesseloracle/keeper"
	"vesseloracle/x/vesseloracle/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVessel(keeper keeper.Keeper, ctx context.Context, n int) []types.Vessel {
	items := make([]types.Vessel, n)
	for i := range items {
		items[i].Imo = strconv.Itoa(i)
		items[i].Ts = uint64(i)
		items[i].Source = strconv.Itoa(i)

		keeper.SetVessel(ctx, items[i])
	}
	return items
}

func TestVesselGet(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	items := createNVessel(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVessel(ctx,
			item.Imo,
			item.Ts,
			item.Source,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVesselRemove(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	items := createNVessel(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVessel(ctx,
			item.Imo,
			item.Ts,
			item.Source,
		)
		_, found := keeper.GetVessel(ctx,
			item.Imo,
			item.Ts,
			item.Source,
		)
		require.False(t, found)
	}
}

func TestVesselGetAll(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	items := createNVessel(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVessel(ctx)),
	)
}
