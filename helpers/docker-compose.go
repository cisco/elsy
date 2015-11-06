package helpers

import (
  "io/ioutil"
  "os"
)

func CreateTempDockerComposeFile(yaml string) string {
  cwd, _ := os.Getwd()
  fh, err := ioutil.TempFile(cwd, "lc_docker_compose_template")
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
