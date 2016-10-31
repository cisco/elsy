package command

import (
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdRun(c *cli.Context) error {
  args := append([]string{"run"}, c.Args()...)

	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
