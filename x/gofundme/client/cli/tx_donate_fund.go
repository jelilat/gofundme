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

func CmdDonateFund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "donate-fund [id] [donation]",
		Short: "Broadcast message donate-fund",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argDonation := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDonateFund(
				clientCtx.GetFromAddress().String(),
				argId,
				argDonation,
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
