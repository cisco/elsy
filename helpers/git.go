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
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var releaseTagRegexp = regexp.MustCompile(`^v\d+\.\d+\.\d+(-.{0,120})?$`)
var releaseRegexp = regexp.MustCompile("^origin/release/(.+)$")
var snapshotRegexp = regexp.MustCompile("^origin/(.+)$")

// regex for valid tag name taken from https://github.com/docker/distribution/blob/b07d759241defb2f345e95ed04bfdeb8ac010ab2/reference/regexp.go#L25
var validTagName = regexp.MustCompile(`^[\w][\w.-]{0,127}$`)

const (
	TagSearch    byte = iota
	BranchSearch byte = iota
)

/*
*  extract tag name from git tag or git branch
*  gives priority to git tag
 */
func ExtractTag(gitTag string, gitBranch string) (string, error) {
	if len(gitTag) > 0 {
		return ExtractTagFromTag(gitTag)
	} else if len(gitBranch) > 0 {
		return ExtractTagFromBranch(gitBranch)
	} else {
		return "", fmt.Errorf("expecting a git branch and/or a git tag be set, found neither")
	}
}

/*
*  extract tag name from branch
*  branch: `master` becomes tag `latest`
*  branch: `origin/release/xxx` becomes tag `xxx`
*  branch: `origin/aaa/xxx` becomes tag `snapshot.aaa.xxx`
*  branch: `origin/xxx` becomes tag `snapshot.xxx`
 */
func ExtractTagFromBranch(gitBranch string) (string, error) {
	var tagName string
	if gitBranch == "origin/master" {
		tagName = "latest"
	} else if matches := releaseRegexp.FindStringSubmatch(gitBranch); matches != nil {
		tagName = matches[1]
	} else if matches := snapshotRegexp.FindStringSubmatch(gitBranch); matches != nil {
		tagName = "snapshot." + matches[1]
	} else {
		return "", fmt.Errorf("could not determine tag from git branch: %q", gitBranch)
	}

	return validateTag(tagName)
}

// extractTagFromTag will extract the docker tag from the git tag
//
// gitTag must be of format 'v.X.Y.Z-q', where X, Y, and Z are ints and q is some character-baed qualifier. example: v0.2.2, v0.2.3-rc1
func ExtractTagFromTag(gitTag string) (string, error) {
	var tagName string

	if match := releaseTagRegexp.MatchString(gitTag); match {
		tagName = gitTag
	} else if len(gitTag) > 0 {
		tagName = "snapshot." + gitTag
	}

	return validateTag(tagName)
}

func validateTag(tag string) (string, error) {
	tagName := strings.Replace(tag, "/", ".", -1)
	if !validTagName.MatchString(tagName) {
		return "", fmt.Errorf("tagName: %q is not valid", tagName)
	}
	return tagName, nil
}

func IsStableBranch(gitBranch string) bool {
	if gitBranch == "origin/master" {
		return true
	} else if releaseRegexp.MatchString(gitBranch) {
		return true
	}
	return false
}

func CheckTag(v string) error {
	if match := releaseTagRegexp.MatchString(v); !match {
		return fmt.Errorf("release value syntax was not valid, it must adhere to: %s", releaseTagRegexp)
	}
	return nil
}

func IsTagNameAlreadyUsed(tag string) (bool, error) {
	return doesTagExist(TagSearch, tag)
}

func IsTagNameAlreadyUsedAsABranchName(tag string) (bool, error) {
	return doesTagExist(BranchSearch, tag)
}

func doesTagExist(searchType byte, tag string) (bool, error) {
	var cmd *exec.Cmd

	if searchType == TagSearch {
		cmd = exec.Command("git", "tag")
	} else {
		cmd = exec.Command("git", "branch", "-a")
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return false, err
	}

	for _, item := range strings.Fields(out.String()) {
		if item == tag ||
			strings.Contains(item, "/"+tag) {
			return true, nil
		}
	}

	return false, nil
}
