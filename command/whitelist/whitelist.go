package whitelist

import (
	"github.com/spf13/cobra"
	"https://github.com/FitRTeams/familychain/command/whitelist/deployment"
	"https://github.com/FitRTeams/familychain/command/whitelist/show"
)

func GetCommand() *cobra.Command {
	whitelistCmd := &cobra.Command{
		Use:   "whitelist",
		Short: "Top level command for modifying the Polygon Edge whitelists within the config. Only accepts subcommands.",
	}

	registerSubcommands(whitelistCmd)

	return whitelistCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		deployment.GetCommand(),
		show.GetCommand(),
	)
}
