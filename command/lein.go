package command

import (
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdLein(c *cli.Context) error {
	args := append([]string{"run", "--rm", "lein"}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
