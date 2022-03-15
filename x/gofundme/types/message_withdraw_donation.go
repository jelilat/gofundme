package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawDonation = "withdraw_donation"

var _ sdk.Msg = &MsgWithdrawDonation{}

func NewMsgWithdrawDonation(creator string, id uint64) *MsgWithdrawDonation {
	return &MsgWithdrawDonation{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgWithdrawDonation) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawDonation) Type() string {
	return TypeMsgWithdrawDonation
}

func (msg *MsgWithdrawDonation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawDonation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawDonation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
