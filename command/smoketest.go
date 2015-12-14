package command

import (
  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdSmoketest processes cmd args and then runs smoketests
func CmdSmoketest(c *cli.Context) error {

  if !c.Bool("skip-package") {
    logrus.Info("Running package before executing smoketests")
    if err := CmdPackage(c); err != nil {
      return err
    }
  }

  return RunSmoketest(c)
}

// RunSmoketest will execute the smoketest service and then return
func RunSmoketest(c *cli.Context) error {
  args := append([]string{"run", "--rm", "smoketest"}, c.Args()...)
  return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
