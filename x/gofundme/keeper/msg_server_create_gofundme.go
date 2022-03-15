package keeper

import (
	"context"

	"github.com/cosmonaut/gofundme/x/gofundme/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGofundme(goCtx context.Context, msg *types.MsgCreateGofundme) (*types.MsgCreateGofundmeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Start a gofundme with the following user input
	var gofundme = types.Gofundme{
		Creator: msg.Creator,
		Goal:    msg.Goal,
		Start:   msg.Start,
		End:     msg.End,
		Claim:   "no",
	}

	//Allow users to set their gofundme campaign start date to be current time
	if gofundme.Start == "now" {
		gofundme.Start = ctx.BlockTime().Format("2006-01-02T15:04:05.000000000")
	}

	// Add the gofundme to the keeper
	k.AppendGofundme(ctx, gofundme)

	return &types.MsgCreateGofundmeResponse{}, nil
}
