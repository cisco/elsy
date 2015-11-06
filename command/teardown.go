package command

import "github.com/codegangsta/cli"

func CmdTeardown(c *cli.Context) {
  dockerComposeExec(c, "kill")
  dockerComposeExec(c, "rm", "-f", "-v")
}
