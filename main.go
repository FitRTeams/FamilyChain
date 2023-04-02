package main

import (
	_ "embed"

	"https://github.com/FitRTeams/familychain/command/root"
	"https://github.com/FitRTeams/familychain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
