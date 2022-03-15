package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDonateFund = "donate_fund"

var _ sdk.Msg = &MsgDonateFund{}

func NewMsgDonateFund(creator string, id uint64, donation string) *MsgDonateFund {
	return &MsgDonateFund{
		Creator:  creator,
		Id:       id,
		Donation: donation,
	}
}

func (msg *MsgDonateFund) Route() string {
	return RouterKey
}

func (msg *MsgDonateFund) Type() string {
	return TypeMsgDonateFund
}

func (msg *MsgDonateFund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDonateFund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDonateFund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
