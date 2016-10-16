package template

import (
	"testing"

	"github.com/elsy/helpers"
)

func TestSharedExternalDataContainer(t *testing.T) {
	ddcTmp := helpers.DockerDataContainer{
		Image:     "busybox:latest",
		Name:      "lc_test_shared_datacontainer_tmp",
		Volumes:   []string{"/tmp"},
		Resilient: true,
	}
	ddcVar := helpers.DockerDataContainer{
		Image:     "busybox:latest",
		Name:      "lc_test_shared_datacontainer_var",
		Volumes:   []string{"/var"},
		Resilient: true,
	}

	if dataContainers := GetSharedExternalDataContainers("foo"); len(dataContainers) != 0 {
		t.Error("expected shared external data containers to start empty")
	}

	addSharedExternalDataContainer("foo", ddcTmp)
	addSharedExternalDataContainer("foo", ddcVar)
	if dataContainers := GetSharedExternalDataContainers("foo"); len(dataContainers) != 2 {
		t.Error("expected shared external data containers to contain added containers")
	}

	addSharedExternalDataContainer("bar", ddcVar)
	if dataContainers := GetSharedExternalDataContainers("bar"); len(dataContainers) != 1 {
		t.Error("expected shared external data containers to contain added containers")
	}
}

func TestTemplateV1Registration(t *testing.T) {
	if _, err := GetV1("foo", false); err == nil {
		t.Error("expected Get to return an error for a non-existant template")
	}
	if err := addV1(template{name: "foo", composeYmlTmpl: "someyaml"}); err != nil {
		t.Error("expected Add to register a template")
	}
	if _, err := GetV1("foo", false); err != nil {
		t.Error("expected Get to return yaml after registering a template")
	}
	if err := addV1(template{name: "foo", composeYmlTmpl: "someyaml"}); err == nil {
		t.Error("expected Add to return an error when registering an existing template")
	}
}

func TestTemplateV2Registration(t *testing.T) {
	if _, err := GetV2("foo", false); err == nil {
		t.Error("expected Get to return an error for a non-existant template")
	}
	if err := addV2(template{name: "foo", composeYmlTmpl: "someyaml"}); err != nil {
		t.Error("expected Add to register a template")
	}
	if _, err := GetV2("foo", false); err != nil {
		t.Error("expected Get to return yaml after registering a template")
	}
	if err := addV2(template{name: "foo", composeYmlTmpl: "someyaml"}); err == nil {
		t.Error("expected Add to return an error when registering an existing template")
	}
}
