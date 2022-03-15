package keeper

import (
	"context"

	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GofundmeAll(c context.Context, req *types.QueryAllGofundmeRequest) (*types.QueryAllGofundmeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gofundmes []types.Gofundme
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	gofundmeStore := prefix.NewStore(store, types.KeyPrefix(types.GofundmeKey))

	pageRes, err := query.Paginate(gofundmeStore, req.Pagination, func(key []byte, value []byte) error {
		var gofundme types.Gofundme
		if err := k.cdc.Unmarshal(value, &gofundme); err != nil {
			return err
		}

		gofundmes = append(gofundmes, gofundme)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGofundmeResponse{Gofundme: gofundmes, Pagination: pageRes}, nil
}

func (k Keeper) Gofundme(c context.Context, req *types.QueryGetGofundmeRequest) (*types.QueryGetGofundmeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	gofundme, found := k.GetGofundme(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetGofundmeResponse{Gofundme: gofundme}, nil
}
