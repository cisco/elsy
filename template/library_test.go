/*
 *  Copyright 2016 Cisco Systems, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package template

import (
	"testing"

	"github.com/cisco/elsy/helpers"
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
	if _, err := GetV1("foo", false, ""); err == nil {
		t.Error("expected Get to return an error for a non-existant template")
	}
	if err := addV1(template{name: "foo", composeYmlTmpl: "someyaml"}); err != nil {
		t.Error("expected Add to register a template")
	}
	if _, err := GetV1("foo", false, ""); err != nil {
		t.Error("expected Get to return yaml after registering a template")
	}
	if err := addV1(template{name: "foo", composeYmlTmpl: "someyaml"}); err == nil {
		t.Error("expected Add to return an error when registering an existing template")
	}
}

func TestTemplateV2Registration(t *testing.T) {
	if _, err := GetV2("foo", false, ""); err == nil {
		t.Error("expected Get to return an error for a non-existant template")
	}
	if err := addV2(template{name: "foo", composeYmlTmpl: "someyaml"}); err != nil {
		t.Error("expected Add to register a template")
	}
	if _, err := GetV2("foo", false, ""); err != nil {
		t.Error("expected Get to return yaml after registering a template")
	}
	if err := addV2(template{name: "foo", composeYmlTmpl: "someyaml"}); err == nil {
		t.Error("expected Add to return an error when registering an existing template")
	}
}
