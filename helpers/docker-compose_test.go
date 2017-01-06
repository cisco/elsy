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

package helpers

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var testComposeVersionData = []struct {
	Version            string
	ExpectedString     string
	ExpectedComponents []int
}{
	{"docker-compose 1.2.0", "1.2.0", []int{1, 2, 0}},
	{"docker-compose version: 1.3.3\nCPython version: 2.7.9\nOpenSSL version: OpenSSL 1.0.1j 15 Oct 2014", "1.3.3", []int{1, 3, 3}},
	{"docker-compose 1.4.1", "1.4.1", []int{1, 4, 1}},
	{"docker-compose version: 1.5.0", "1.5.0", []int{1, 5, 0}},
	{"docker-compose version 1.5.2, build 7240ff3", "1.5.2", []int{1, 5, 2}},
	{"docker-compose version: 1.6.0", "1.6.0", []int{1, 6, 0}},
	{"docker-compose version 1.8.1, build 878cff1", "1.8.1", []int{1, 8, 1}},
}

func TestParseDockerCompseVersion(t *testing.T) {
	for _, data := range testComposeVersionData {
		version, components, _ := parseDockerComposeVersion(data.Version)
		if data.ExpectedString != version {
			t.Errorf("expected version to be: %q but got %q instead", data.ExpectedString, version)
		}
		if !reflect.DeepEqual(data.ExpectedComponents, components) {
			t.Errorf("expected version to be: %q but got %q instead", data.ExpectedComponents, components)
		}
	}
}

var ymlBad = `
notvalidyml
`

func TestParseComposeFileV1(t *testing.T) {
	yml := `
testservice1:
  image: image1
`
	expectedMap := DockerComposeMap{"testservice1": DockerComposeService{Image: "image1", Build: ""}}
	doTestParseComposeFile(t, yml, V1, expectedMap)
}

func TestParseComposeFileV2(t *testing.T) {
	yml := `
version: '2'
services:
  testservice2:
    image: image2
`
	expectedMap := DockerComposeMap{"testservice2": DockerComposeService{Image: "image2", Build: ""}}
	doTestParseComposeFile(t, yml, V2, expectedMap)
}

func TestParseComposeFileInvalid(t *testing.T) {
	yml := "invalid"
	doTestParseComposeFile(t, yml, unknown, nil)
}

func doTestParseComposeFile(t *testing.T, yml string, expectedVersion ComposeFileVersion, expectedMap DockerComposeMap) {
	file, err := ioutil.TempFile(os.TempDir(), "helpers")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())
	file.WriteString(yml)

	version, composeMap, _ := parseComposeFile(file.Name())

	if version != expectedVersion {
		t.Errorf("expected version to be: %q but got %q instead", expectedVersion, version)
	}

	if !reflect.DeepEqual(expectedMap, composeMap) {
		t.Errorf("expected composeMap to be: %q but got %q instead", expectedMap, composeMap)
	}
}

func TestGetComposeFileVersionV2(t *testing.T) {
	yml := `
version: '2'
services:
  testservice2:
    image: image2
`
	testVersion(t, yml, V2, V1)
}

func TestGetComposeFileVersionInvalid(t *testing.T) {
	yml := "invalid"
	testVersion(t, yml, V2, V2)
}

func testVersion(t *testing.T, yml string, expectedVersion ComposeFileVersion, defaultVersion ComposeFileVersion) {
	file, err := ioutil.TempFile(os.TempDir(), "helpers")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())
	file.WriteString(yml)

	version := GetComposeFileVersion(file.Name(), defaultVersion)
	if version != expectedVersion {
		t.Errorf("expected version to be: %q but got %q instead", expectedVersion, version)
	}
}
