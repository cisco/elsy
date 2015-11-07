package command

import (
  "os"
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdPackage(c *cli.Context) {
  commands := []*exec.Cmd{dockerComposeCommand(c, "run", "--rm", "package")}

  // docker build
  if _, err := os.Stat("Dockerfile"); err == nil {
    logrus.Debug("detected Dockerfile for packaging")
    if image, err := helpers.DockerImage("Dockerfile"); err == nil {
      commands = append(commands, exec.Command("docker", "pull", image))
    }
    commands = append(commands, exec.Command("docker", "build", "-t", os.Getenv("COMPOSE_PROJECT_NAME"), "."))
  }
  helpers.ChainCommands(commands)
}
