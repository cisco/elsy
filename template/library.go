package template

import (
  "fmt"

  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

var sharedExternalDataContainers = make(map[string][]helpers.DockerDataContainer)

func AddSharedExternalDataContainer(templateName string, ddc helpers.DockerDataContainer) {
  dataContainers := GetSharedExternalDataContainers(templateName)
  sharedExternalDataContainers[templateName] = append(dataContainers, ddc)
}

func GetSharedExternalDataContainers(templateName string) []helpers.DockerDataContainer {
  return sharedExternalDataContainers[templateName]
}

var templates = make(map[string]string)
func Add(name string, yaml string) error {
  if _, ok := templates[name]; ok {
    return fmt.Errorf("template %q already exists", name)
  }
  templates[name] = yaml
  return nil
}

func Get(name string) (string, error) {
  yaml, present := templates[name]
  if !present {
    return "", fmt.Errorf("template %q is not registered", name)
  }
  return yaml, nil
}
