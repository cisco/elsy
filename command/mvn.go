package command

import (
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdMvn(c *cli.Context) error {
	args := append([]string{"run", "--rm", "mvn"}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
