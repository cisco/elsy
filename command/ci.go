package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdCi(c *cli.Context) error {
  if err := CmdBootstrap(c); err != nil {
    return err
  }
  if err := CmdTest(c); err != nil {
    return err
  }
  if err := CmdPackage(c); err != nil {
    return err
  }
  if helpers.DockerComposeHasService("smoketest") {
    if err := CmdSmoketest(c); err != nil {
      return err
    }
  }
  return CmdPublish(c)
}
