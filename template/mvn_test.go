package template

import "testing"

func TestMvnTemplate(t *testing.T) {
	if dataContainers := GetSharedExternalDataContainers("mvn"); len(dataContainers) != 1 {
		t.Error("expected mvn template to register one shared external data container")
	}

	if _, err := GetV1("mvn", false); err != nil {
		t.Error("expected mvn template to be registered")
	}

	if _, err := GetV2("mvn", false); err != nil {
		t.Error("expected mvn template to be registered")
	}
}
