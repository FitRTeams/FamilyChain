package main

import (
	_ "embed"

	"github.com/familychain/family/command/root"
	"github.com/familychain/family/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
