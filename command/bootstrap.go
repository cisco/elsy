package command

import (
  "os/exec"

  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdBootstrap(c *cli.Context) {
  CmdTeardown(c)
  helpers.ChainCommands([]*exec.Cmd{
    dockerComposeCommand(c, "build", "--pull"),
    dockerComposeCommand(c, "pull", "--ignore-pull-failures"),
  })
}
