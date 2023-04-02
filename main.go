package main

import (
	_ "embed"

	"github.com/FamilyChain/family/command/root"
	"github.com/FamilyChain/family/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
