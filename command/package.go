package command

import (
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdPackage runs package service if present and then attempts to build Dockerfile
func CmdPackage(c *cli.Context) error {
  commands := []*exec.Cmd{}

  if helpers.DockerComposeHasService("package") {
    commands = append(commands, helpers.DockerComposeCommand("run", "--rm", "package"))
  } else {
    logrus.Debug("no package service found, skipping")
  }

  // docker build
  if helpers.HasDockerfile() {
    logrus.Debug("detected Dockerfile for packaging")
    if image, err := helpers.DockerImage("Dockerfile"); err == nil {
      commands = append(commands, exec.Command("docker", "pull", image.String()))
    }
    dockerImageName := c.String("docker-image-name")
    if len(dockerImageName) == 0 {
      logrus.Panic("you must use `--docker-image-name` to package a docker image")
    }
    commands = append(commands, exec.Command("docker", "build", "-t", dockerImageName, "."))
  }
  return helpers.ChainCommands(commands)
}
