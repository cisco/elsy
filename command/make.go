package command

import (
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdMake(c *cli.Context) error {
	args := append([]string{"run", "--rm", "make"}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
