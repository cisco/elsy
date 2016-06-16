package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

var releaseTagRegexp = regexp.MustCompile(`^v\d+\.\d+\.\d+(-.{0,120})?$`)
var releaseRegexp = regexp.MustCompile("^origin/release/(.+)$")
var snapshotRegexp = regexp.MustCompile("^origin/(.+)$")

// regex for valid tag name taken from https://github.com/docker/distribution/blob/b07d759241defb2f345e95ed04bfdeb8ac010ab2/reference/regexp.go#L25
var validTagName = regexp.MustCompile(`^[\w][\w.-]{0,127}$`)

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
