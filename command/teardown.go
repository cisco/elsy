package command

import (
  "os/exec"

  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdTeardown(c *cli.Context) error {
  return helpers.ChainCommands([]*exec.Cmd{
    dockerComposeCommand(c, "kill"),
    dockerComposeCommand(c, "rm", "-f", "-v"),
  })
}
