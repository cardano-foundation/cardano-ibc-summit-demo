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

func createNConsolidatedDataReport(keeper keeper.Keeper, ctx context.Context, n int) []types.ConsolidatedDataReport {
	items := make([]types.ConsolidatedDataReport, n)
	for i := range items {
		items[i].Imo = strconv.Itoa(i)
		items[i].Ts = uint64(i)

		keeper.SetConsolidatedDataReport(ctx, items[i])
	}
	return items
}

func TestConsolidatedDataReportGet(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	items := createNConsolidatedDataReport(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetConsolidatedDataReport(ctx,
			item.Imo,
			item.Ts,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestConsolidatedDataReportRemove(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	items := createNConsolidatedDataReport(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveConsolidatedDataReport(ctx,
			item.Imo,
			item.Ts,
		)
		_, found := keeper.GetConsolidatedDataReport(ctx,
			item.Imo,
			item.Ts,
		)
		require.False(t, found)
	}
}

func TestConsolidatedDataReportGetAll(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	items := createNConsolidatedDataReport(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllConsolidatedDataReport(ctx)),
	)
}
