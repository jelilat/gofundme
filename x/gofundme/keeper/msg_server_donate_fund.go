package keeper

import (
	"context"

	"github.com/cosmonaut/gofundme/x/gofundme/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DonateFund(goCtx context.Context, msg *types.MsgDonateFund) (*types.MsgDonateFundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	gofundme, found := k.GetGofundme(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "gofundme campaign with id %s not found", msg.Id)
	}

	//Get the start date
	start, err := sdk.ParseTimeBytes([]byte(gofundme.Start))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	//Get the end date
	end, err := sdk.ParseTimeBytes([]byte(gofundme.End))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	//Check whether the current time is between the start and end date
	if ctx.BlockTime().Before(start) || ctx.BlockTime().After(end) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "gofundme campaign is not active")
	}

	//Get the donor's address
	donor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid donor address (%s)", donor)
	}
	//Get the donation amount
	donation, _ := sdk.ParseCoinsNormalized(msg.Donation)

	//Ensure that the donation amount is greater than 0
	if donation.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "donation amount must be greater than 0")
	}

	//Update the gofundme status
	gofundme.Donor = append(gofundme.Donor, msg.Creator)
	gofundme.Donation = append(gofundme.Donation, msg.Donation)
	gofundme.Totaldonations += donation[0].Amount.Uint64()

	//Use the module account as an escrow account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, donor, types.ModuleName, donation)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetGofundme(ctx, gofundme)

	return &types.MsgDonateFundResponse{}, nil
}
