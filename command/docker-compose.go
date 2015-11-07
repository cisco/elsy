package command

import (
  "os"
  "os/exec"

  "github.com/codegangsta/cli"
)

func CmdDockerCompose(c *cli.Context) {
  dockerComposeCommand(c, c.Args()...)
}

func dockerComposeCommand(c *cli.Context, args ...string) *exec.Cmd {
  args = append([]string{"-f", "docker-compose.yml"}, args...)

  if baseComposeFile := os.Getenv("LC_BASE_COMPOSE_FILE"); len(baseComposeFile) > 0 {
    args = append([]string{"-f", baseComposeFile}, args...)
  }

  return exec.Command(c.GlobalString("docker-compose"), args...)
}
