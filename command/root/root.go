package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/FitRTeams/familychain/command/backup"
	"github.com/FitRTeams/familychain/command/bridge"
	"github.com/FitRTeams/familychain/command/genesis"
	"github.com/FitRTeams/familychain/command/helper"
	"github.com/FitRTeams/familychain/command/ibft"
	"github.com/FitRTeams/familychain/command/license"
	"github.com/FitRTeams/familychain/command/monitor"
	"github.com/FitRTeams/familychain/command/peers"
	"github.com/FitRTeams/familychain/command/polybft"
	"github.com/FitRTeams/familychain/command/polybftmanifest"
	"github.com/FitRTeams/familychain/command/polybftsecrets"
	"github.com/FitRTeams/familychain/command/regenesis"
	"github.com/FitRTeams/familychain/command/rootchain"
	"github.com/FitRTeams/familychain/command/secrets"
	"github.com/FitRTeams/familychain/command/server"
	"github.com/FitRTeams/familychain/command/status"
	"github.com/FitRTeams/familychain/command/txpool"
	"github.com/FitRTeams/familychain/command/version"
	"github.com/FitRTeams/familychain/command/whitelist"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Family Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		polybftmanifest.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
