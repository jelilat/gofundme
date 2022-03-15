package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGofundme = "create_gofundme"

var _ sdk.Msg = &MsgCreateGofundme{}

func NewMsgCreateGofundme(creator string, goal uint64, start string, end string) *MsgCreateGofundme {
	return &MsgCreateGofundme{
		Creator: creator,
		Goal:    goal,
		Start:   start,
		End:     end,
	}
}

func (msg *MsgCreateGofundme) Route() string {
	return RouterKey
}

func (msg *MsgCreateGofundme) Type() string {
	return TypeMsgCreateGofundme
}

func (msg *MsgCreateGofundme) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGofundme) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGofundme) ValidateBasic() error {
	// Validate the creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Get the goal amount
	goal := sdk.NewIntFromUint64(msg.Goal)

	// Ensure that the goal amount is greater than 0
	if goal.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "goal amount must be greater than 0")
	}

	var start time.Time
	//Get the start date
	if msg.Start != "now" {
		startDate, err := sdk.ParseTimeBytes([]byte(msg.Start))
		if err != nil {
			return err
		}
		start = startDate
	} else {
		start = time.Now()
	}

	//Get the end date
	end, err := sdk.ParseTimeBytes([]byte(msg.End))
	if err != nil {
		return err
	}

	// Get the current block time
	blockTime := time.Now()

	// Ensure that end date is in the future
	if end.Before(blockTime) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "end date must be in the future")
	}

	//Ensure that start date is not in the past
	if start.Before(blockTime) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "start date cannot be in the past")
	}

	// Ensure that start date is before end date
	if start.After(end) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "start date must be before end date")
	}
	return nil
}
