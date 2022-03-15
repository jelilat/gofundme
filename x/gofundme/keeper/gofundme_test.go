package keeper_test

import (
	"testing"

	keepertest "github.com/cosmonaut/gofundme/testutil/keeper"
	"github.com/cosmonaut/gofundme/testutil/nullify"
	"github.com/cosmonaut/gofundme/x/gofundme/keeper"
	"github.com/cosmonaut/gofundme/x/gofundme/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNGofundme(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Gofundme {
	items := make([]types.Gofundme, n)
	for i := range items {
		items[i].Id = keeper.AppendGofundme(ctx, items[i])
	}
	return items
}

func TestGofundmeGet(t *testing.T) {
	keeper, ctx := keepertest.GofundmeKeeper(t)
	items := createNGofundme(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetGofundme(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestGofundmeRemove(t *testing.T) {
	keeper, ctx := keepertest.GofundmeKeeper(t)
	items := createNGofundme(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGofundme(ctx, item.Id)
		_, found := keeper.GetGofundme(ctx, item.Id)
		require.False(t, found)
	}
}

func TestGofundmeGetAll(t *testing.T) {
	keeper, ctx := keepertest.GofundmeKeeper(t)
	items := createNGofundme(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGofundme(ctx)),
	)
}

func TestGofundmeCount(t *testing.T) {
	keeper, ctx := keepertest.GofundmeKeeper(t)
	items := createNGofundme(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetGofundmeCount(ctx))
}
