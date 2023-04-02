package txpool

import (
	"github.com/spf13/cobra"
	"https://github.com/FitRTeams/familychain/command/helper"
	"https://github.com/FitRTeams/familychain/command/txpool/status"
	"https://github.com/FitRTeams/familychain/command/txpool/subscribe"
)

func GetCommand() *cobra.Command {
	txPoolCmd := &cobra.Command{
		Use:   "txpool",
		Short: "Top level command for interacting with the transaction pool. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(txPoolCmd)

	registerSubcommands(txPoolCmd)

	return txPoolCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// txpool status
		status.GetCommand(),
		// txpool subscribe
		subscribe.GetCommand(),
	)
}
