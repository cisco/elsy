package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdSbt(c *cli.Context) {
  args := append([]string{"run", "--rm", "sbt"}, c.Args()...)
  helpers.RunCommand(dockerComposeCommand(c, args...))
}
