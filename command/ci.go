package command

import (
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdCi(c *cli.Context) error {
  // call 'teardown' at the end of 'ci', even if a failure occurs
  // note, we don't want the LastCommandSuccess from teardown otherwise
  // it would pollute the exit code.
  defer func() {
    orig := helpers.LastCommandSuccess
    CmdTeardown(c)
    helpers.LastCommandSuccess = orig
  }()
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
