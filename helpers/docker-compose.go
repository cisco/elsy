package helpers

import (
  "io/ioutil"
  "os"
  "path/filepath"
  "regexp"
  "strings"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
)

func DockerComposeBeforeHook(c *cli.Context) {
  os.Setenv("COMPOSE_PROJECT_NAME", dockerComposeProjectName(c))

  if yaml := baseYAML(c.GlobalString("template")); len(yaml) > 0 {
    os.Setenv("LC_BASE_COMPOSE_FILE", createTempDockerComposeFile(yaml))
  }

  if dataContainer, ok := dataContainers[c.GlobalString("template")]; ok {
    if err := dataContainer.Ensure(); err != nil {
      logrus.Panic("unable to create data container")
    }
  }
}

func DockerComposeAfterHook(c *cli.Context) error {
  file := os.Getenv("LC_BASE_COMPOSE_FILE")
  if err := os.Remove(file); err != nil {
    return err
  }
  return nil
}

func dockerComposeProjectName(c *cli.Context) string {
  var invalidChars = regexp.MustCompile("[^a-z0-9]")
  projectName := c.GlobalString("project-name")
  if len(projectName) == 0 {
    path, _ := os.Getwd()
    projectName = filepath.Base(path)
  }
  return invalidChars.ReplaceAllString(strings.ToLower(projectName), "")
}

func createTempDockerComposeFile(yaml string) string {
  cwd, _ := os.Getwd()
  fh, err := ioutil.TempFile(cwd, "lc_docker_compose_template")
  if err != nil {
    logrus.Panic("could not create temporary yaml file")
  }
  defer fh.Close()
  _, err = fh.WriteString(yaml)
  if err != nil {
    logrus.Panic("could not write to temporary yaml file")
  }
  return fh.Name()
}

// TEMPLATES
func baseYAML(template string) string {
  switch template {
    case "sbt":
      return `
sbt: &sbt
  image: arch-docker.eng.lancope.local:5000/sbt
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: sbt
  volumes_from:
    - lc_shared_sbtdata
test:
  <<: *sbt
  entrypoint: [sbt, test]
package:
  <<: *sbt
  command: [assembly]
`
    case "mvn":
      return `
mvn: &mvn
  image: maven:3.2-jdk-8
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: mvn
  volumes_from:
    - lc_shared_mvndata
test:
  <<: *mvn
  entrypoint: [mvn, test]
package
  <<: *mvn
  command: [package, "-DskipTests=true"]
`
    default:
      return ""
  }
}

var dataContainers = map[string]DockerDataContainer{
  "sbt": {
    Image: "busybox:latest",
    Name: "lc_shared_sbtdata",
    Volumes: []string{"/root/.ivy2"},
    Resilient: true,
  },
  "mvn": {
    Image: "busybox:latest",
    Name: "lc_shared_mvndata",
    Volumes: []string{"/root/.m2/repository"},
    Resilient: true,
  },
}
