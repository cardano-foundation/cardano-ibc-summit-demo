package types_test

import (
	"testing"

	"vesseloracle/x/vesseloracle/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated vessel",
			genState: &types.GenesisState{
				VesselList: []types.Vessel{
					{
						Imo: "0",
						Ts:  0,
					},
					{
						Imo: "0",
						Ts:  0,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
