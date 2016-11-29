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
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGetConfigFileMethods(t *testing.T) {
	if fh, err := ioutil.TempFile("", "testlcconfigyaml"); err != nil {
		t.Fatal("could not create temporary file")
	} else {
		defer fh.Close()
		fh.WriteString("foo: bar\n")
		fh.WriteString("slicetest:\n")
		fh.WriteString("  - value1\n")
		fh.WriteString("  - value2\n")
		fh.WriteString("slicetest2: ['value1', 'value2']\n")
		fh.Sync()
		LoadConfigFile(fh.Name())
	}
	// get configured value
	if v := GetConfigFileString("foo"); v != "bar" {
		t.Errorf("expected to get 'bar' for 'foo' but got %q instead", v)
	}
	// get value containing slice:
	expected := []string{"value1", "value2"}
	if v := GetConfigFileSlice("slicetest"); !reflect.DeepEqual(expected, v) {
		t.Errorf("expected to get %q for 'arraytest' but got %q instead", expected, v)
	}
	if v := GetConfigFileSlice("slicetest2"); !reflect.DeepEqual(expected, v) {
		t.Errorf("expected to get %q for 'arraytest' but got %q instead", expected, v)
	}
	// get empty for unset values
	if v := GetConfigFileString("bar"); v != "" {
		t.Errorf("expected to get '' for 'bar' but got %q instead", v)
	}
	// when given a default and a configured value, get configured value
	if v := GetConfigFileStringWithDefault("foo", "baz"); v != "bar" {
		t.Errorf("expected to get 'bar' for 'foo' but got %q instead", v)
	}
	// when given a default and no configured value, get default
	if v := GetConfigFileStringWithDefault("bar", "foo"); v != "foo" {
		t.Errorf("expected to get 'foo' for 'bar' but got %q instead", v)
	}
}
