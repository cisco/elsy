package command

import (
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdTest(c *cli.Context) error {
	args := append([]string{"run", "--rm", "test"}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
