package command

import (
  "os/exec"

  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command/system"
)

func CmdBootstrap(c *cli.Context) error {
  if err := system.CmdVerifyLds(c); err != nil {
    return err
  }
  CmdTeardown(c)
  commands := []*exec.Cmd{
    helpers.DockerComposeCommand("build", "--pull"),
    helpers.DockerComposeCommand("pull", "--ignore-pull-failures"),
  }
  if err := helpers.ChainCommands(commands); err != nil {
    return err
  }
  return CmdInstallDependencies(c)
}
