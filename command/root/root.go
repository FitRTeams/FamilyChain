package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/FamilyChain/family/command/backup"
	"github.com/FamilyChain/family/command/bridge"
	"github.com/FamilyChain/family/command/genesis"
	"github.com/FamilyChain/family/command/helper"
	"github.com/FamilyChain/family/command/ibft"
	"github.com/FamilyChain/family/command/license"
	"github.com/FamilyChain/family/command/monitor"
	"github.com/FamilyChain/family/command/peers"
	"github.com/FamilyChain/family/command/polybft"
	"github.com/FamilyChain/family/command/polybftmanifest"
	"github.com/FamilyChain/family/command/polybftsecrets"
	"github.com/FamilyChain/family/command/regenesis"
	"github.com/FamilyChain/family/command/rootchain"
	"github.com/FamilyChain/family/command/secrets"
	"github.com/FamilyChain/family/command/server"
	"github.com/FamilyChain/family/command/status"
	"github.com/FamilyChain/family/command/txpool"
	"github.com/FamilyChain/family/command/version"
	"github.com/FamilyChain/family/command/whitelist"
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
