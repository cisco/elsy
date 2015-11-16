package command

import (
  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdInstallDependencies(c *cli.Context) error {
  if helpers.DockerComposeHasService("installdependencies") {
    return helpers.RunCommand(helpers.DockerComposeCommand("run", "--rm", "installdependencies"))
  } else {
    logrus.Debugf("no installdependencies service found, skipping installdependencies")
  }
  return nil
}
