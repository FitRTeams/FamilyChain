package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"https://github.com/FitRTeams/familychain/command/backup"
	"https://github.com/FitRTeams/familychain/command/bridge"
	"https://github.com/FitRTeams/familychain/command/genesis"
	"https://github.com/FitRTeams/familychain/command/helper"
	"https://github.com/FitRTeams/familychain/command/ibft"
	"https://github.com/FitRTeams/familychain/command/license"
	"https://github.com/FitRTeams/familychain/command/monitor"
	"https://github.com/FitRTeams/familychain/command/peers"
	"https://github.com/FitRTeams/familychain/command/polybft"
	"https://github.com/FitRTeams/familychain/command/polybftmanifest"
	"https://github.com/FitRTeams/familychain/command/polybftsecrets"
	"https://github.com/FitRTeams/familychain/command/regenesis"
	"https://github.com/FitRTeams/familychain/command/rootchain"
	"https://github.com/FitRTeams/familychain/command/secrets"
	"https://github.com/FitRTeams/familychain/command/server"
	"https://github.com/FitRTeams/familychain/command/status"
	"https://github.com/FitRTeams/familychain/command/txpool"
	"https://github.com/FitRTeams/familychain/command/version"
	"https://github.com/FitRTeams/familychain/command/whitelist"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
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
