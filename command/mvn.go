package command

import "github.com/codegangsta/cli"

func CmdMvn(c *cli.Context) {
  args := append([]string{"run", "--rm", "mvn"}, c.Args()...)
  dockerComposeExec(c, args...)
}
