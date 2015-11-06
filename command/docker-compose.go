package command

import (
  "os"
  "os/exec"
  "path/filepath"
  "regexp"
  "strings"

  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdDockerCompose(c *cli.Context) {
  dockerComposeExec(c, c.Args().Tail()...)
}

func dockerComposeExec(c *cli.Context, args ...string) error {
  composeFile, _ := filepath.Abs(filepath.Join(c.GlobalString("root"), "docker-compose.yml"))
  args = append([]string{"-f", composeFile}, args...)

  yaml := baseYAML(c)
  if len(yaml) > 0 {
    file := helpers.CreateTempDockerComposeFile(yaml)
    defer os.Remove(file)
    args = append([]string{"-f", file}, args...)
  }

  if dataContainer, ok := dataContainers[c.GlobalString("template")]; ok {
    if err := dataContainer.Ensure(); err != nil {
      panic("unable to create data container")
    }
  }

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

func baseYAML(c *cli.Context) string {
  switch c.GlobalString("template") {
    case "sbt":
      return `
sbt:
  image: arch-docker.eng.lancope.local:5000/sbt
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: sbt
  volumes_from:
    - lc_shared_sbtdata
`
    case "mvn":
      return `
mvn:
  image: maven:3.2-jdk-8
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: mvn
  volumes_from:
    - lc_shared_mvndata
`
    default:
      return ""
  }
}

var dataContainers = map[string]helpers.DockerDataContainer{
  "sbt": {
    Name: "lc_shared_sbtdata",
    Volumes: []string{"/root/.ivy2"},
    Resilient: true,
  },
  "mvn": {
    Name: "lc_shared_mvndata",
    Volumes: []string{"/root/.m2/repository"},
    Resilient: true,
  },
}
