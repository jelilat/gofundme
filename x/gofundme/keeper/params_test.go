package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/gofundme/testutil/keeper"
	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GofundmeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
