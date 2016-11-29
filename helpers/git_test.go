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

import "testing"

type publishTestData struct {
	Input       string
	TagName     string
	ErrExpected bool
}

type extractor func(string) (string, error)

func TestExtractTagFromBranch(t *testing.T) {
	data := []publishTestData{
		{"origin/master", "latest", false},
		{"origin/release", "snapshot.release", false},
		{"origin/release/1.0", "1.0", false},
		{"origin/foo", "snapshot.foo", false},
		{"origin/foo/bar", "snapshot.foo.bar", false},
		{"origin", "", true},
		{"release", "", true},
		{"foo", "", true},
	}
	doTest(t, &data, ExtractTagFromBranch)
}

func TestExtractTagFromTag(t *testing.T) {
	data := []publishTestData{
		{"v0.0.0", "v0.0.0", false},
		{"v9.9.9", "v9.9.9", false},
		{"v0.0.0-rc1", "v0.0.0-rc1", false},
		{"v0.0.0-r/c1", "v0.0.0-r.c1", false},
		{"v99.99.99", "v99.99.99", false},
		{"v99.99.99-rc1", "v99.99.99-rc1", false},
		{"v99.99", "snapshot.v99.99", false},
		{"v0.0.0rc1", "snapshot.v0.0.0rc1", false},
		{"9.9.9", "snapshot.9.9.9", false},
		{"9", "snapshot.9", false},
		{"foo-test", "snapshot.foo-test", false},
		{"testing-1.2.3", "snapshot.testing-1.2.3", false},
		{"x/[z", "", true},
		{"", "", true},
	}
	doTest(t, &data, ExtractTagFromTag)
}

func doTest(t *testing.T, data *[]publishTestData, f extractor) {
	for _, d := range *data {
		if tagName, err := f(d.Input); d.TagName != tagName {
			t.Errorf("expected input: %q to produce tag: %q but got %q instead", d.Input, d.TagName, tagName)
		} else if d.ErrExpected && err == nil {
			t.Errorf("expected input: %q to produce an error but did not receive an error", d.Input)
		} else if !d.ErrExpected && err != nil {
			t.Errorf("expected input: %q to not produce an error but received error: %v", d.Input, err)
		}
	}
}
