package helpers

import (
  "io/ioutil"
  "os"
  "path"
  "reflect"
  "testing"
)

var testDockerComposeFiles = map[string]DockerComposeMap {
  `myservice:
    build: .`: {"myservice": DockerComposeService{Build: "."}},
  `myservice:
    image: nginx`: {"myservice": DockerComposeService{Image: "nginx"}},
}
func TestDockerCompose(t *testing.T) {
  dir, err := ioutil.TempDir("", "DockerComposeTest")
  if err != nil {
    t.Fatal("unable to create temporary directory")
  }
  defer os.RemoveAll(dir)
  for input, expected := range testDockerComposeFiles {
    ioutil.WriteFile(path.Join(dir, "docker-compose.yml"), []byte(input), os.FileMode(int(0644)))
    if !reflect.DeepEqual(DockerCompose(dir), expected) {
      t.Errorf("did not get expected map from yaml:\n%s\n", input)
    }
  }
}
