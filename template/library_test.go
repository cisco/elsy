package template

import (
  "testing"

  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func TestSharedExternalDataContainer(t *testing.T) {
  ddcTmp := helpers.DockerDataContainer{
    Image: "busybox:latest",
    Name: "lc_test_shared_datacontainer_tmp",
    Volumes: []string{"/tmp"},
    Resilient: true,
  }
  ddcVar := helpers.DockerDataContainer{
    Image: "busybox:latest",
    Name: "lc_test_shared_datacontainer_var",
    Volumes: []string{"/var"},
    Resilient: true,
  }

  if dataContainers := GetSharedExternalDataContainers("foo"); len(dataContainers) != 0 {
    t.Error("expected shared external data containers to start empty")
  }

  AddSharedExternalDataContainer("foo", ddcTmp)
  AddSharedExternalDataContainer("foo", ddcVar)
  if dataContainers := GetSharedExternalDataContainers("foo"); len(dataContainers) != 2 {
    t.Error("expected shared external data containers to contain added containers")
  }

  AddSharedExternalDataContainer("bar", ddcVar)
  if dataContainers := GetSharedExternalDataContainers("bar"); len(dataContainers) != 1 {
    t.Error("expected shared external data containers to contain added containers")
  }
}

func TestTemplateRegistration(t *testing.T) {
  if _, err := Get("foo", false); err == nil {
    t.Error("expected Get to return an error for a non-existant template")
  }
  if err := Add("foo", "someyaml"); err != nil {
    t.Error("expected Add to register a template")
  }
  if _, err := Get("foo", false); err != nil {
    t.Error("expected Get to return yaml after registering a template")
  }
  if err := Add("foo", "someyaml"); err == nil {
    t.Error("expected Add to return an error when registering an existing template")
  }
}
