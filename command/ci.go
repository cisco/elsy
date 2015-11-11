package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdCi(c *cli.Context) {
  CmdBootstrap(c)
  if !helpers.LastCommandSuccess {
    return
  }
  CmdTest(c)
  if !helpers.LastCommandSuccess {
    return
  }
  CmdPackage(c)
  if !helpers.LastCommandSuccess {
    return
  }
  if helpers.DockerComposeHasService("smoketest") {
    CmdSmoketest(c)
    if !helpers.LastCommandSuccess {
      return
    }
  }
  CmdPublish(c)
}
