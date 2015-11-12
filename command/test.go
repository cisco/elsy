package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdTest(c *cli.Context) error {
  args := append([]string{"run", "--rm", "test"}, c.Args()...)
  return helpers.RunCommand(dockerComposeCommand(c, args...))
}
