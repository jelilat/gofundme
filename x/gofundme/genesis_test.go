package gofundme_test

import (
	"testing"

	keepertest "github.com/cosmonaut/gofundme/testutil/keeper"
	"github.com/cosmonaut/gofundme/testutil/nullify"
	"github.com/cosmonaut/gofundme/x/gofundme"
	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		GofundmeList: []types.Gofundme{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		GofundmeCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GofundmeKeeper(t)
	gofundme.InitGenesis(ctx, *k, genesisState)
	got := gofundme.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.GofundmeList, got.GofundmeList)
	require.Equal(t, genesisState.GofundmeCount, got.GofundmeCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
