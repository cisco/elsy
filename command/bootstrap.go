package command

import (
  "github.com/codegangsta/cli"
)

func CmdBootstrap(c *cli.Context) {
  dockerComposeExec(c, "build", "--pull")
  dockerComposeExec(c, "pull", "--ignore-pull-failures")
}
