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

package main

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	config            string
	expected          []string
	shouldPanic       bool
	expectedErrorText string
}

func TestResolveDockerRegistry(t *testing.T) {
	testCases := []testCase{
		{"", []string{}, false, ""},
		{"docker_registry: registry", []string{"registry"}, false, ""},
		{"docker_registries: [registry1, test]", []string{"registry1", "test"}, false, ""},
		{"docker_registries: [singletonreg]", []string{"singletonreg"}, false, ""},
		{"docker_registries: \n  - registryb\n  - registryc", []string{"registryb", "registryc"}, false, ""},
		{"docker_registry: registry\ndocker_registries: [registry1, test]", []string{"registry"}, true, "multiple docker registry configs"},
		{"docker_registries: \n  -registryb\n", []string{"registryb"}, true, "but did not find any registries, verify that yaml is correct"},
	}

	for _, test := range testCases {
		if test.shouldPanic {
			runAssertPanic(t, test, resolveDockerRegistryFromConfig)
		} else {
			runAssertNoPanic(t, test, resolveDockerRegistryFromConfig)
		}
	}
}

func TestResolveLocalImages(t *testing.T) {
	testCases := []testCase{
		{"", []string{}, false, ""},
		{"local_images: [image1, test]", []string{"image1", "test"}, false, ""},
		{"local_images: [singletonimage]", []string{"singletonimage"}, false, ""},
		{"local_images: \n  - imageb\n  - imagec", []string{"imageb", "imagec"}, false, ""},
		{"local_images: \n  -imageb\n", []string{"imageb"}, true, "but did not find any images, verify that yaml is correct"},
	}

	for _, test := range testCases {
		if test.shouldPanic {
			runAssertPanic(t, test, resolveLocalImagesFromConfig)
		} else {
			runAssertNoPanic(t, test, resolveLocalImagesFromConfig)
		}
	}
}

type retrieveKeyFunc func() *cli.StringSlice

func runAssertPanic(t *testing.T, test testCase, fn retrieveKeyFunc) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic, but was supposed to, test case: %v", test)
		} else if !strings.Contains(r.(error).Error(), test.expectedErrorText) {
			t.Errorf("Expected error text to contain %q, but found %q", test.expectedErrorText, r)
		}
	}()
	loadConfig(t, test.config)
	fn()
}

func runAssertNoPanic(t *testing.T, test testCase, fn retrieveKeyFunc) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked, but was not supposed to. test case: %v, error: %v", test, r)
		}
	}()
	loadConfig(t, test.config)
	if v := fn(); !reflect.DeepEqual(test.expected, v.Value()) {
		t.Errorf("expected to get %q, but got %q instead", test.expected, v.Value())
	}
}

func loadConfig(t *testing.T, config string) {
	if fh, err := ioutil.TempFile("", "testlcconfigyaml"); err != nil {
		t.Fatal("could not create temporary file")
	} else {
		defer fh.Close()
		fh.WriteString(config)
		LoadConfigFile(fh.Name())
	}
}
