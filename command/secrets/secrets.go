package secrets

import (
	"github.com/spf13/cobra"
	"https://github.com/FitRTeams/familychain/command/helper"
	"https://github.com/FitRTeams/familychain/command/secrets/generate"
	initCmd "https://github.com/FitRTeams/familychain/command/secrets/init"
	"https://github.com/FitRTeams/familychain/command/secrets/output"
)

func GetCommand() *cobra.Command {
	secretsCmd := &cobra.Command{
		Use:   "secrets",
		Short: "Top level SecretsManager command for interacting with secrets functionality. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(secretsCmd)

	registerSubcommands(secretsCmd)

	return secretsCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// secrets init
		initCmd.GetCommand(),
		// secrets generate
		generate.GetCommand(),
		// secrets output public data
		output.GetCommand(),
	)
}
