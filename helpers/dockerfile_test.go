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
	"path"
	"reflect"
	"testing"
)

func TestDockerfileImages(t *testing.T) {
	dir, err := ioutil.TempDir("", "DockerfileImagesTest")
	if err != nil {
		t.Fatal("unable to create temporary directory")
	}
	defer os.RemoveAll(dir)
	nestedDockerDir := path.Join(dir, "dev-env")
	os.MkdirAll(nestedDockerDir, os.FileMode(int(0755)))
	otherDir := path.Join(dir, "target")
	os.MkdirAll(otherDir, os.FileMode(int(0755)))

	files := map[string]string{
		path.Join(dir, "Dockerfile"):             "FROM bar",
		path.Join(nestedDockerDir, "Dockerfile"): "FROM library/foo",
		path.Join(otherDir, "Dockerfile"):        "BLERG",
	}
	for file, content := range files {
		if err := ioutil.WriteFile(file, []byte(content), os.FileMode(int(0644))); err != nil {
			t.Fatalf("could not create file %v", path.Join(dir, "Dockerfile"))
		}
	}

	if images := DockerfileImages(dir); !reflect.DeepEqual(images, []string{"bar", "library/foo"}) {
		t.Errorf("did not get expected string slice. got %v", images)
	}

	// test DockerImage()
	if image, err := DockerImage(path.Join(dir, "Dockerfile")); err != nil {
		t.Fatal(err)
	} else if image.IsRemote() {
		t.Errorf("expected image 'bar' to not be remote")
	}
	if image, err := DockerImage(path.Join(nestedDockerDir, "Dockerfile")); err != nil {
		t.Fatal(err)
	} else if !image.IsRemote() {
		t.Errorf("expected image 'library/foo' to be remote")
	}
}
