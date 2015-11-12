package command

import (
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdPackage(c *cli.Context) error {
  commands := []*exec.Cmd{dockerComposeCommand(c, "run", "--rm", "package")}

  // docker build
  if helpers.HasDockerfile() {
    logrus.Debug("detected Dockerfile for packaging")
    if image, err := helpers.DockerImage("Dockerfile"); err == nil {
      commands = append(commands, exec.Command("docker", "pull", image))
    }
    dockerImageName := c.String("docker-image-name")
    if len(dockerImageName) == 0 {
      logrus.Panic("you must use `--docker-image-name` to package a docker image")
    }
    commands = append(commands, exec.Command("docker", "build", "-t", dockerImageName, "."))
  }
  return helpers.ChainCommands(commands)
}
