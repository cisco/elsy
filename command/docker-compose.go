package command

import (
  "os"
  "os/exec"
  "path/filepath"
  "regexp"
  "strings"

  "github.com/codegangsta/cli"
)

func CmdDockerCompose(c *cli.Context) {
  dockerComposeExec(c, c.Args().Tail()...)
}

func dockerComposeExec(c *cli.Context, args ...string) error {
  cmd := exec.Command(c.GlobalString("docker-compose"), args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func DockerComposeSetenv(c *cli.Context) {
  os.Setenv("COMPOSE_PROJECT_NAME", dockerComposeProjectName(c))
}

func dockerComposeProjectName(c *cli.Context) string {
  var invalidChars = regexp.MustCompile("[^a-z0-9]")
  projectName := c.GlobalString("project-name")
  if len(projectName) == 0 {
    path, _ := filepath.Abs(c.GlobalString("root"))
    projectName = filepath.Base(path)
  }
  return invalidChars.ReplaceAllString(strings.ToLower(projectName), "")
}
