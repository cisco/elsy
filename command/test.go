package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command/system"
)

func CmdTest(c *cli.Context) error {
  if err := system.CmdVerifyLds(c); err != nil {
    return err
  }
  args := append([]string{"run", "--rm", "test"}, c.Args()...)
  return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
