package polybft

import (
	"github.com/FitRTeams/familychain/command/sidechain/registration"
	"github.com/FitRTeams/familychain/command/sidechain/staking"
	"github.com/FitRTeams/familychain/command/sidechain/unstaking"
	"github.com/FitRTeams/familychain/command/sidechain/validators"

	"github.com/FitRTeams/familychain/command/sidechain/whitelist"
	"github.com/FitRTeams/familychain/command/sidechain/withdraw"
	"github.com/spf13/cobra"
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
