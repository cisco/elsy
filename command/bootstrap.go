package command

import (
  "os/exec"

  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdBootstrap(c *cli.Context) error {
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
