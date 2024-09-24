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

func TestVesselMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.VesseloracleKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateVessel{Creator: creator,
			Imo:    strconv.Itoa(i),
			Ts:     uint64(i),
			Source: strconv.Itoa(i),
		}
		_, err := srv.CreateVessel(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetVessel(ctx,
			expected.Imo,
			expected.Ts,
			expected.Source,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestVesselMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateVessel
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateVessel{Creator: creator,
				Imo:    strconv.Itoa(0),
				Ts:     0,
				Source: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateVessel{Creator: "B",
				Imo:    strconv.Itoa(0),
				Ts:     0,
				Source: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateVessel{Creator: creator,
				Imo:    strconv.Itoa(100000),
				Ts:     100000,
				Source: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VesseloracleKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateVessel{Creator: creator,
				Imo:    strconv.Itoa(0),
				Ts:     0,
				Source: strconv.Itoa(0),
			}
			_, err := srv.CreateVessel(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateVessel(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetVessel(ctx,
					expected.Imo,
					expected.Ts,
					expected.Source,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestVesselMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteVessel
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteVessel{Creator: creator,
				Imo:    strconv.Itoa(0),
				Ts:     0,
				Source: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteVessel{Creator: "B",
				Imo:    strconv.Itoa(0),
				Ts:     0,
				Source: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteVessel{Creator: creator,
				Imo:    strconv.Itoa(100000),
				Ts:     100000,
				Source: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VesseloracleKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateVessel(ctx, &types.MsgCreateVessel{Creator: creator,
				Imo:    strconv.Itoa(0),
				Ts:     0,
				Source: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteVessel(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetVessel(ctx,
					tc.request.Imo,
					tc.request.Ts,
					tc.request.Source,
				)
				require.False(t, found)
			}
		})
	}
}
