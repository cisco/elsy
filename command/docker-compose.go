package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdDockerCompose(c *cli.Context) error {
  return helpers.RunCommand(helpers.DockerComposeCommand(c.Args()...))
}
