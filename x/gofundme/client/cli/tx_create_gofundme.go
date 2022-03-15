package cli

import (
	"strconv"

	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateGofundme() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-gofundme [goal] [start] [end]",
		Short: "Broadcast message create-gofundme",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGoal, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argStart := args[1]
			argEnd := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGofundme(
				clientCtx.GetFromAddress().String(),
				argGoal,
				argStart,
				argEnd,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
