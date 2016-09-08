package command

import (
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdSbt(c *cli.Context) error {
	args := append([]string{"run", "--rm", "sbt"}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
