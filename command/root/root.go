package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/familychain/family/command/backup"
	"github.com/familychain/family/command/bridge"
	"github.com/familychain/family/command/genesis"
	"github.com/familychain/family/command/helper"
	"github.com/familychain/family/command/ibft"
	"github.com/familychain/family/command/license"
	"github.com/familychain/family/command/monitor"
	"github.com/familychain/family/command/peers"
	"github.com/familychain/family/command/polybft"
	"github.com/familychain/family/command/polybftmanifest"
	"github.com/familychain/family/command/polybftsecrets"
	"github.com/familychain/family/command/regenesis"
	"github.com/familychain/family/command/rootchain"
	"github.com/familychain/family/command/secrets"
	"github.com/familychain/family/command/server"
	"github.com/familychain/family/command/status"
	"github.com/familychain/family/command/txpool"
	"github.com/familychain/family/command/version"
	"github.com/familychain/family/command/whitelist"
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
