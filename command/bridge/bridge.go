package bridge

import (
	"github.com/spf13/cobra"

	"github.com/FitRTeams/familychain/command/bridge/deposit"
	"github.com/FitRTeams/familychain/command/bridge/exit"
	"github.com/FitRTeams/familychain/command/bridge/withdraw"
)

// GetCommand creates "bridge" helper command
func GetCommand() *cobra.Command {
	bridgeCmd := &cobra.Command{
		Use:   "bridge",
		Short: "Top level bridge command.",
	}

	registerSubcommands(bridgeCmd)

	return bridgeCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// bridge deposit
		deposit.GetCommand(),
		// bridge withdraw
		withdraw.GetCommand(),
		// bridge exit
		exit.GetCommand(),
	)
}
