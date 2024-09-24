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

func TestVesselQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	msgs := createNVessel(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVesselRequest
		response *types.QueryGetVesselResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVesselRequest{
				Imo: msgs[0].Imo,
				Ts:  msgs[0].Ts,
			},
			response: &types.QueryGetVesselResponse{Vessel: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVesselRequest{
				Imo: msgs[1].Imo,
				Ts:  msgs[1].Ts,
			},
			response: &types.QueryGetVesselResponse{Vessel: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVesselRequest{
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
			response, err := keeper.Vessel(ctx, tc.request)
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

func TestVesselQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VesseloracleKeeper(t)
	msgs := createNVessel(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVesselRequest {
		return &types.QueryAllVesselRequest{
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
			resp, err := keeper.VesselAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Vessel), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Vessel),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VesselAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Vessel), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Vessel),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VesselAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Vessel),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VesselAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
