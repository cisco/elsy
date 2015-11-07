package command

import (
  "os"
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdPackage(c *cli.Context) {
  dockerComposeExec(c, "run", "--rm", "package")
  // docker build
  if _, err := os.Stat("Dockerfile"); err == nil {
    if image, err := helpers.DockerImage("Dockerfile"); err == nil {
      pullCmd := exec.Command("docker", "pull", image)
      pullCmd.Stdout = os.Stdout
      pullCmd.Stderr = os.Stderr
      pullCmd.Stdin  = os.Stdin
      logrus.Debugf("running command %s with args %v", pullCmd.Path, pullCmd.Args)
      pullCmd.Run()
    }
    cmd := exec.Command("docker", "build", "-t", os.Getenv("COMPOSE_PROJECT_NAME"), ".")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin  = os.Stdin
    logrus.Debugf("running command %s with args %v", cmd.Path, cmd.Args)
    cmd.Run()
  }
}
