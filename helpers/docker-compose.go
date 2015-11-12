package helpers

import (
  "io/ioutil"
  "os"
  "path/filepath"
  "regexp"
  "strings"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "gopkg.in/yaml.v2"
)

func DockerComposeBeforeHook(c *cli.Context) {
  os.Setenv("COMPOSE_PROJECT_NAME", dockerComposeProjectName(c))

  if yaml := baseYAML(c.GlobalString("template")); len(yaml) > 0 {
    file := createTempDockerComposeFile(yaml)
    logrus.Debugf("setting LC_BASE_COMPOSE_FILE to %v", file)
    os.Setenv("LC_BASE_COMPOSE_FILE", file)
  }

  if dataContainer, ok := dataContainers[c.GlobalString("template")]; ok {
    if err := dataContainer.Ensure(); err != nil {
      logrus.Panic("unable to create data container")
    }
  }
}

func DockerComposeAfterHook(c *cli.Context) error {
  if file := os.Getenv("LC_BASE_COMPOSE_FILE"); len(file) > 0 {
    logrus.Debugf("attempting to remove base compose file: %v", file)
    if err := os.Remove(file); err != nil {
      return err
    }
  }
  return nil
}

func dockerComposeProjectName(c *cli.Context) string {
  var invalidChars = regexp.MustCompile("[^a-z0-9]")
  projectName := c.GlobalString("project-name")
  if len(projectName) == 0 {
    logrus.Debug("using current working directory for compose project name")
    path, _ := os.Getwd()
    projectName = filepath.Base(path)
  } else {
    logrus.Debugf("using configured value: %q for project name", projectName)
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
    case "ember":
      return `
emberdata:
  image: arch-docker.eng.lancope.local:5000/ember
  volumes:
   - /opt/app/bower_components
   - /opt/app/dist
   - /opt/app/node_modules
   - /opt/app/vendor
   - /opt/app/tmp
  entrypoint: "/bin/true"
ember: &ember
  image: arch-docker.eng.lancope.local:5000/ember
  volumes:
   - .:/opt/app
  working_dir: /opt/app
  entrypoint: /usr/local/bin/ember
  volumes_from:
   - emberdata
npm:
  <<: *ember
  entrypoint: /usr/local/bin/npm
bower:
  <<: *ember
  entrypoint: /usr/local/bin/bower
installdependencies:
  <<: *ember
  entrypoint: bash
  command: -c "npm install && npm update && bower install && bower update"
test:
  <<: *ember
  command: [test]
package:
  <<: *ember
  command: [build, "--environment='production'",  "--output-path=dist-production"]
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

type DockerComposeMap map[string]DockerComposeService

type DockerComposeService struct {
  Build string
  Image string
}

func DockerComposeServices() (services []string) {
  for k := range getDockerComposeMap("docker-compose.yml") {
    services = append(services, k)
  }
  if file := os.Getenv("LC_BASE_COMPOSE_FILE"); len(file) > 0 {
    for k := range getDockerComposeMap(file) {
      services = append(services, k)
    }
  }
  return
}

func getDockerComposeMap(file string) (m DockerComposeMap) {
  if s, err := ioutil.ReadFile(file); err != nil {
    panic(err)
  } else if err := yaml.Unmarshal(s, &m); err != nil {
    panic(err)
  }
  return
}

func DockerComposeHasService(service string) bool {
  for _, v := range DockerComposeServices() {
    if v == service {
      return true
    }
  }
  return false
}
