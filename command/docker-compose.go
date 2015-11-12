package command

import (
  "os"
  "os/exec"

  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdDockerCompose(c *cli.Context) error {
  return helpers.RunCommand(dockerComposeCommand(c, c.Args()...))
}

func dockerComposeCommand(c *cli.Context, args ...string) *exec.Cmd {
  if _, err := os.Stat("docker-compose.yml"); err == nil {
    args = append([]string{"-f", "docker-compose.yml"}, args...)
  }

  if baseComposeFile := os.Getenv("LC_BASE_COMPOSE_FILE"); len(baseComposeFile) > 0 {
    args = append([]string{"-f", baseComposeFile}, args...)
  }

  return exec.Command(c.GlobalString("docker-compose"), args...)
}
