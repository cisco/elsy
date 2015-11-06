package command

import (
  "os"
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
)

func CmdDockerCompose(c *cli.Context) {
  dockerComposeExec(c, c.Args().Tail()...)
}

func dockerComposeExec(c *cli.Context, args ...string) error {
  args = append([]string{"-f", "docker-compose.yml"}, args...)

  if baseComposeFile := os.Getenv("LC_BASE_COMPOSE_FILE"); len(baseComposeFile) > 0 {
    args = append([]string{"-f", baseComposeFile}, args...)
  }

  cmd := exec.Command(c.GlobalString("docker-compose"), args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  logrus.Debugf("running command %s with args %v", cmd.Path, cmd.Args)
  return cmd.Run()
}
