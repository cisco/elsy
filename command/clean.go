package command

import (
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdClean will run a "clean" service that will remove old build artifacts.
func CmdClean(c *cli.Context) error {
	args := append([]string{"run", "--rm", "clean"}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
