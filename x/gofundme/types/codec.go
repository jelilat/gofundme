package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGofundme{}, "gofundme/CreateGofundme", nil)
	cdc.RegisterConcrete(&MsgDonateFund{}, "gofundme/DonateFund", nil)
	cdc.RegisterConcrete(&MsgWithdrawDonation{}, "gofundme/WithdrawDonation", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGofundme{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDonateFund{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawDonation{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
