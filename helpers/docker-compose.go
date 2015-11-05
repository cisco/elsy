package helpers

import (
  "io/ioutil"
  "path"

  "github.com/codegangsta/cli"
  "gopkg.in/yaml.v2"
)

type DockerCompose map[string]DockerComposeService

type DockerComposeService struct {
  Build string
  Image string
}

func GetDockerCompose(c *cli.Context) DockerCompose {
  return parseYAML(readYAML(path.Join(c.GlobalString("root"), "docker-compose.yml")))
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
