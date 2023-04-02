package peers

import (
	"github.com/spf13/cobra"
	"https://github.com/FitRTeams/familychain/command/helper"
	"https://github.com/FitRTeams/familychain/command/peers/add"
	"https://github.com/FitRTeams/familychain/command/peers/list"
	"https://github.com/FitRTeams/familychain/command/peers/status"
)

func GetCommand() *cobra.Command {
	peersCmd := &cobra.Command{
		Use:   "peers",
		Short: "Top level command for interacting with the network peers. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(peersCmd)

	registerSubcommands(peersCmd)

	return peersCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// peers status
		status.GetCommand(),
		// peers list
		list.GetCommand(),
		// peers add
		add.GetCommand(),
	)
}
