package command

import (
  "github.com/codegangsta/cli"
)

func CmdBootstrap(c *cli.Context) {
  CmdTeardown(c)
  dockerComposeExec(c, "build", "--pull")
  dockerComposeExec(c, "pull", "--ignore-pull-failures")
}
