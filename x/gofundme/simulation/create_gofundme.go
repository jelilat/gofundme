package simulation

import (
	"math/rand"

	"github.com/cosmonaut/gofundme/x/gofundme/keeper"
	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateGofundme(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateGofundme{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateGofundme simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateGofundme simulation not implemented"), nil, nil
	}
}
