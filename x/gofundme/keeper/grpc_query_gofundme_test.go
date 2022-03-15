package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/cosmonaut/gofundme/testutil/keeper"
	"github.com/cosmonaut/gofundme/testutil/nullify"
	"github.com/cosmonaut/gofundme/x/gofundme/types"
)

func TestGofundmeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GofundmeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGofundme(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGofundmeRequest
		response *types.QueryGetGofundmeResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetGofundmeRequest{Id: msgs[0].Id},
			response: &types.QueryGetGofundmeResponse{Gofundme: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetGofundmeRequest{Id: msgs[1].Id},
			response: &types.QueryGetGofundmeResponse{Gofundme: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetGofundmeRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Gofundme(wctx, tc.request)
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

func TestGofundmeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GofundmeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGofundme(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGofundmeRequest {
		return &types.QueryAllGofundmeRequest{
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
			resp, err := keeper.GofundmeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Gofundme), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Gofundme),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GofundmeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Gofundme), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Gofundme),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GofundmeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Gofundme),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GofundmeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
