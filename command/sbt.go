package command

import "github.com/codegangsta/cli"

func CmdSbt(c *cli.Context) {
  args := append([]string{"run", "--rm", "sbt"}, c.Args()...)
  dockerComposeExec(c, args...)
}
