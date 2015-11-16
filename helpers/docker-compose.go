package helpers

import (
  "io/ioutil"
  "os"
  "os/exec"
  "errors"
  "strings"
  "strconv"

  "github.com/Sirupsen/logrus"
  "github.com/codegangsta/cli"
  "gopkg.in/yaml.v2"
)

type DockerComposeMap map[string]DockerComposeService

type DockerComposeService struct {
  Build string
  Image string
}

func DockerComposeServices() (services []string) {
  if _, err := os.Stat("docker-compose.yml"); err == nil {
    for k := range getDockerComposeMap("docker-compose.yml") {
      services = append(services, k)
    }
  }
  if file := os.Getenv("LC_BASE_COMPOSE_FILE"); len(file) > 0 {
    for k := range getDockerComposeMap(file) {
      services = append(services, k)
    }
  }
  return
}

// first return value is the human readable version
// second return value is an array of the {majorVersion, minorVersion, patchVersion}
func GetDockerComposeVersion(c *cli.Context) (string, []int, error){
  if out, err := RunCommandWithOutput(exec.Command(c.GlobalString("docker-compose"), "--version")); err != nil {
    return "", nil, err
  } else {

    return parseDockerComposeVersion(out)
  }
}

func parseDockerComposeVersion(versionString string) (string, []int, error){
  // assuming version is last word in string
  firstLine := strings.Split(versionString, "\n")[0]
  words := strings.Split(firstLine, " ")
  version := strings.TrimSpace(words[len(words)-1])
  versionArray := strings.Split(version, ".")

  if len(versionArray) != 3 {
    logrus.Debugf("could not parse version, expected 3 version components, found %d", len(versionArray))
    return version, nil, errors.New("could not parse version")
  }

  versionNumbers := []int{}
  for _, x := range versionArray {
    if val, err := strconv.Atoi(x); err == nil{
      versionNumbers = append(versionNumbers, val)
    } else {
      logrus.Debugf("could not parse integers from version %s", version, err)
      return version, nil, errors.New("could not parse version")
    }
  }
  return version, versionNumbers, nil
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
