package secrets

import (
	"github.com/FamilyChain/family/command/helper"
	"github.com/FamilyChain/family/command/secrets/generate"
	initCmd "github.com/FamilyChain/family/command/secrets/init"
	"github.com/FamilyChain/family/command/secrets/output"
	"github.com/spf13/cobra"
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
