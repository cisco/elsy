package helpers

import (
  "io/ioutil"

  "github.com/codegangsta/cli"
  "gopkg.in/yaml.v2"
)

type DockerCompose map[string]DockerComposeService

type DockerComposeService struct {
  Build string
  Image string
}

func GetDockerCompose(c *cli.Context) DockerCompose {
  return parseYAML(readYAML(c.GlobalString("docker-compose")))
}

func parseYAML(s []byte) (d DockerCompose) {
  if err := yaml.Unmarshal(s, &d); err != nil {
    panic(err)
  }
  return
}

func readYAML(path string) []byte {
  if s, err := ioutil.ReadFile(path); err != nil {
    panic(err)
  } else {
    return s
  }
}
