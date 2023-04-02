package polybft

import (
	"github.com/FamilyChain/family/command/sidechain/registration"
	"github.com/FamilyChain/family/command/sidechain/staking"
	"github.com/FamilyChain/family/command/sidechain/unstaking"
	"github.com/FamilyChain/family/command/sidechain/validators"

	"github.com/FamilyChain/family/command/sidechain/whitelist"
	"github.com/FamilyChain/family/command/sidechain/withdraw"
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
