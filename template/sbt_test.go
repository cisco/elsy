package template

import "testing"

func TestSbtTemplate(t *testing.T) {
  if dataContainers := GetSharedExternalDataContainers("sbt"); len(dataContainers) != 1 {
    t.Error("expected sbt template to register one shared external data container")
  }

  if _, err := Get("sbt", false); err != nil {
    t.Error("expected sbt template to be registered")
  }
}
