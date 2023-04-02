package main

import (
	_ "embed"

	"github.com/FitRTeams/familychain/command/root"
	"github.com/FitRTeams/familychain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
