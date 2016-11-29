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
	"reflect"
	"testing"
)

var testVersionStringData = []struct {
	Version            string
	ExpectedComponents []int
}{
	{"1.9.1", []int{1, 9, 1}},
	{"1.10.3", []int{1, 10, 3}},
	{"1.11.2", []int{1, 11, 2}},
	{"", nil},
}

func TestParseVersionString(t *testing.T) {
	for _, data := range testVersionStringData {
		components, _ := parseVersionString(data.Version)
		if !reflect.DeepEqual(data.ExpectedComponents, components) {
			t.Errorf("expected version to be: %q but got %q instead", data.ExpectedComponents, components)
		}
	}
}
