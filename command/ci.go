package command

import (
  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdCi runs ci loop
func CmdCi(c *cli.Context) error {
  defer CmdTeardown(c)

  logrus.Info("Running bootstrap")
  if err := CmdBootstrap(c); err != nil {
    return err
  }

  if helpers.DockerComposeHasService("test") {
    logrus.Info("Running test")
    if err := CmdTest(c); err != nil {
      return err
    }
  }

  logrus.Info("Running package")
  if err := CmdPackage(c); err != nil {
    return err
  }

  if helpers.DockerComposeHasService("smoketest") {
    logrus.Info("Running smoketest")

    if err := RunSmoketest(c); err != nil {
      return err
    }
  }

  logrus.Info("Running publish")
  return CmdPublish(c)
}
