package helpers

import (
  "io/ioutil"
  "os"

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
