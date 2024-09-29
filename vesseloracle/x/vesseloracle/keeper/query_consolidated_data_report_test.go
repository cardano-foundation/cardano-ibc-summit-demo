package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "vesseloracle/testutil/keeper"
	"vesseloracle/testutil/nullify"
	"vesseloracle/x/vesseloracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestConsolidatedDataReportQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	msgs := createNConsolidatedDataReport(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetConsolidatedDataReportRequest
		response *types.QueryGetConsolidatedDataReportResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetConsolidatedDataReportRequest{
				Imo: msgs[0].Imo,
				Ts:  msgs[0].Ts,
			},
			response: &types.QueryGetConsolidatedDataReportResponse{ConsolidatedDataReport: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetConsolidatedDataReportRequest{
				Imo: msgs[1].Imo,
				Ts:  msgs[1].Ts,
			},
			response: &types.QueryGetConsolidatedDataReportResponse{ConsolidatedDataReport: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetConsolidatedDataReportRequest{
				Imo: strconv.Itoa(100000),
				Ts:  100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ConsolidatedDataReport(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestConsolidatedDataReportQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	msgs := createNConsolidatedDataReport(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllConsolidatedDataReportRequest {
		return &types.QueryAllConsolidatedDataReportRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ConsolidatedDataReportAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ConsolidatedDataReport), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ConsolidatedDataReport),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ConsolidatedDataReportAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ConsolidatedDataReport), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ConsolidatedDataReport),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ConsolidatedDataReportAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ConsolidatedDataReport),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ConsolidatedDataReportAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
