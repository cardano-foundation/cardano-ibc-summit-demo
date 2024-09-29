package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "vesseloracle/testutil/keeper"
	"vesseloracle/x/vesseloracle/keeper"
	"vesseloracle/x/vesseloracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestConsolidatedDataReportMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.VesseloracleKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateConsolidatedDataReport{Creator: creator,
			Imo: strconv.Itoa(i),
			Ts:  uint64(i),
		}
		_, err := srv.CreateConsolidatedDataReport(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetConsolidatedDataReport(ctx,
			expected.Imo,
			expected.Ts,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestConsolidatedDataReportMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateConsolidatedDataReport
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateConsolidatedDataReport{Creator: creator,
				Imo: strconv.Itoa(0),
				Ts:  0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateConsolidatedDataReport{Creator: "B",
				Imo: strconv.Itoa(0),
				Ts:  0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateConsolidatedDataReport{Creator: creator,
				Imo: strconv.Itoa(100000),
				Ts:  100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VesseloracleKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateConsolidatedDataReport{Creator: creator,
				Imo: strconv.Itoa(0),
				Ts:  0,
			}
			_, err := srv.CreateConsolidatedDataReport(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateConsolidatedDataReport(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetConsolidatedDataReport(ctx,
					expected.Imo,
					expected.Ts,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestConsolidatedDataReportMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteConsolidatedDataReport
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteConsolidatedDataReport{Creator: creator,
				Imo: strconv.Itoa(0),
				Ts:  0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteConsolidatedDataReport{Creator: "B",
				Imo: strconv.Itoa(0),
				Ts:  0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteConsolidatedDataReport{Creator: creator,
				Imo: strconv.Itoa(100000),
				Ts:  100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VesseloracleKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateConsolidatedDataReport(ctx, &types.MsgCreateConsolidatedDataReport{Creator: creator,
				Imo: strconv.Itoa(0),
				Ts:  0,
			})
			require.NoError(t, err)
			_, err = srv.DeleteConsolidatedDataReport(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetConsolidatedDataReport(ctx,
					tc.request.Imo,
					tc.request.Ts,
				)
				require.False(t, found)
			}
		})
	}
}
