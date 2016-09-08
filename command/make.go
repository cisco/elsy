package command

import (
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdMake(c *cli.Context) error {
	args := append([]string{"run", "--rm", "make"}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
