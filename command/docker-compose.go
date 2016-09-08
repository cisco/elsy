package command

import (
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdDockerCompose(c *cli.Context) error {
	return helpers.RunCommand(helpers.DockerComposeCommand(c.Args()...))
}
