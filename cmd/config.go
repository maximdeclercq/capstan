package cmd

import (
	"github.com/cloudius-systems/capstan/util"
	"github.com/urfave/cli"
)

// ConfigPrint prints current capstan configuration to console.
func ConfigPrint(c *cli.Context) error {
	repo := util.NewRepo(c.GlobalString("u"))
	return repo.PrintRepo()
}
