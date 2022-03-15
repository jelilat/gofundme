package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/gofundme/testutil/keeper"
	"github.com/cosmonaut/gofundme/x/gofundme/keeper"
	"github.com/cosmonaut/gofundme/x/gofundme/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GofundmeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
