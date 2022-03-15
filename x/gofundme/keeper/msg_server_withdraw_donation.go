package keeper

import (
	"context"
	"fmt"

	"strconv"

	"github.com/cosmonaut/gofundme/x/gofundme/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) WithdrawDonation(goCtx context.Context, msg *types.MsgWithdrawDonation) (*types.MsgWithdrawDonationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	gofundme, found := k.GetGofundme(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "gofundme with id %s not found", msg.Id)
	}

	//Confirm that the donations have not been claimed
	if gofundme.Claim == "yes" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "gofundme with id %s has already been claimed", msg.Id)
	}

	//Get the end date
	_, err := sdk.ParseTimeBytes([]byte(gofundme.End))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Can't withdraw donations until the end of the gofundme campaign.")
	}

	//Get gofundme creator and message sender
	creator := sdk.AccAddress(gofundme.Creator)
	messageSender := sdk.AccAddress(msg.Creator)

	//Get total donations and gofundme goal
	totalDonations := gofundme.Totaldonations
	goal := gofundme.Goal
	stringTotalDonation := strconv.FormatUint(totalDonations, 10)
	normalizedTotalDonation, _ := sdk.ParseCoinsNormalized(stringTotalDonation)
	fmt.Print(creator, messageSender)

	//Ensure that creator can withdraw only if the goal was met
	if creator.Equals(messageSender) && totalDonations < goal {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Can't withdraw donations because goal wasn't met.")
	}

	if creator.Equals(messageSender) && totalDonations >= goal {
		gofundme.Claim = "yes"
		sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, normalizedTotalDonation)
		if sdkError != nil {
			return nil, sdkError
		}
	}

	//Ensure that donor can withdraw only if goal wasn't met
	if !creator.Equals(messageSender) && totalDonations >= goal {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only the owner of the crowdfund can withdraw donations")
	}

	if !creator.Equals(messageSender) && totalDonations < goal {
		var index int
		for idx, v := range gofundme.Donation {
			if v == msg.Creator {
				index = idx
				break
			}
		}
		donation := gofundme.Donation[index]
		normalizedDonation, _ := sdk.ParseCoinsNormalized(donation)

		//Ensure that donor can only withdraw once
		if normalizedDonation.IsZero() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Donation already withdrawn")
		}
		//Update the donation amount to zero to reflect that the donation was withdrawn
		gofundme.Donation[index] = "0"
		sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, messageSender, normalizedDonation)
		if sdkError != nil {
			return nil, sdkError
		}
	}

	k.SetGofundme(ctx, gofundme)

	return &types.MsgWithdrawDonationResponse{}, nil
}
