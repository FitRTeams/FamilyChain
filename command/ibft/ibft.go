package ibft

import (
	"github.com/spf13/cobra"
	"https://github.com/FitRTeams/familychain/command/helper"
	"https://github.com/FitRTeams/familychain/command/ibft/candidates"
	"https://github.com/FitRTeams/familychain/command/ibft/propose"
	"https://github.com/FitRTeams/familychain/command/ibft/quorum"
	"https://github.com/FitRTeams/familychain/command/ibft/snapshot"
	"https://github.com/FitRTeams/familychain/command/ibft/status"
	_switch "https://github.com/FitRTeams/familychain/command/ibft/switch"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
