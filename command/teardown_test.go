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

package command

import (
	"reflect"
	"testing"
)

var removeTestData = []struct {
	ids      []string
	remove   []string
	expected []string
}{
	{[]string{"one", "two", "three", "four"}, []string{"three"}, []string{"one", "two", "four"}},
	{[]string{"one", "two", "three", "four"}, []string{"one", "two", "three", "four"}, []string{}},
	{[]string{"one", "two", "three", "four"}, []string{}, []string{"one", "two", "three", "four"}},
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"one"}, []string{}},
}

func TestRemoveIds(t *testing.T) {

	for _, data := range removeTestData {
		result := removeIds(&data.ids, &data.remove)
		if !reflect.DeepEqual(data.expected, result) {
			t.Errorf("expected result to be: %q but got %q instead", data.expected, result)
		}
	}

}
