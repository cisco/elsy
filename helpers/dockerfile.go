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
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func DockerfileImages(root string) (images []string) {
	for _, file := range dockerfiles(root) {
		if image, err := DockerImage(file); err != nil {
			fmt.Println(err)
		} else {
			images = append(images, image.String())
		}
	}
	return
}

func dockerfiles(root string) (d []string) {
	root, _ = filepath.Abs(root)
	filepath.Walk(root, func(path string, _ os.FileInfo, _ error) error {
		switch filepath.Base(path) {
		case ".git":
			return filepath.SkipDir
		case "Dockerfile":
			d = append(d, path)
		}
		return nil
	})
	return
}

type DockerImageString string

func (s DockerImageString) IsRemote() bool {
	return strings.Contains(string(s), "/")
}
func (s DockerImageString) String() string {
	return string(s)
}

func DockerImage(dockerfile string) (DockerImageString, error) {
	dockerFrom := regexp.MustCompile("^FROM\\s+?(.+)")
	fh, err := os.Open(dockerfile)
	if err != nil {
		return "", err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		if matches := dockerFrom.FindStringSubmatch(scanner.Text()); matches != nil {
			return DockerImageString(matches[1]), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", errors.New(fmt.Sprintf("could not find FROM statement in %s", dockerfile))
}

func HasDockerfile() bool {
	if _, err := os.Stat("Dockerfile"); err == nil {
		return true
	} else {
		return false
	}
}
