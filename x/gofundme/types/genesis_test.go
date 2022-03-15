package types_test

import (
	"testing"

	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
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

				GofundmeList: []types.Gofundme{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				GofundmeCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated gofundme",
			genState: &types.GenesisState{
				GofundmeList: []types.Gofundme{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid gofundme count",
			genState: &types.GenesisState{
				GofundmeList: []types.Gofundme{
					{
						Id: 1,
					},
				},
				GofundmeCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
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
