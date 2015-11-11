package command

import (
  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdInstallDependencies(c *cli.Context) {
  if helpers.DockerComposeHasService("installdependencies") {
    helpers.RunCommand(dockerComposeCommand(c, "run", "--rm", "installdependencies"))
  } else {
    logrus.Debugf("no installdependencies service found, skipping installdependencies")
  }
}
