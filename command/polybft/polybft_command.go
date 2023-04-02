package polybft

import (
	"https://github.com/FitRTeams/familychain/command/sidechain/registration"
	"https://github.com/FitRTeams/familychain/command/sidechain/staking"
	"https://github.com/FitRTeams/familychain/command/sidechain/unstaking"
	"https://github.com/FitRTeams/familychain/command/sidechain/validators"

	"github.com/spf13/cobra"
	"https://github.com/FitRTeams/familychain/command/sidechain/whitelist"
	"https://github.com/FitRTeams/familychain/command/sidechain/withdraw"
)

func GetCommand() *cobra.Command {
	polybftCmd := &cobra.Command{
		Use:   "polybft",
		Short: "Polybft command",
	}

	polybftCmd.AddCommand(
		staking.GetCommand(),
		unstaking.GetCommand(),
		withdraw.GetCommand(),
		validators.GetCommand(),
		whitelist.GetCommand(),
		registration.GetCommand(),
	)

	return polybftCmd
}
