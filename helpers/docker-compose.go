package helpers

import (
  "io/ioutil"
)

func CreateTempDockerComposeFile(yaml string) string {
  fh, err := ioutil.TempFile("", "lc_docker_compose_template")
  if err != nil {
    panic("could not create temporary yaml file")
  }
  defer fh.Close()
  _, err = fh.WriteString(yaml)
  if err != nil {
    panic("could not write to temporary yaml file")
  }
  return fh.Name()
}
