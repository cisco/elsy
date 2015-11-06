package command

import "github.com/codegangsta/cli"

func CmdTest(c *cli.Context) {
  args := append([]string{"run", "--rm", "test"}, c.Args().Tail()...)
  dockerComposeExec(c, args...)
}
