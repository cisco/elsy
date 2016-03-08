package command

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdPublish will publish all artifacts associated with the current repo
func CmdPublish(c *cli.Context) error {
	// first try to publish gitTag
	gitTag := c.String("git-tag")
	if len(gitTag) != 0 {
		logrus.Infof("attempting to publish git tag %q", gitTag)
		return publishTag(gitTag, c)
	}

	// if no tag was found, attempt to publish the branch
	gitBranch := c.String("git-branch")
	if len(gitBranch) == 0 {
		return fmt.Errorf("The publish task requires that either a git branch or git tag be set, found neither")
	}
	logrus.Infof("attempting to publish git branch %q", gitBranch)
	return publishBranch(gitBranch, c)
}

func publishTag(tag string, c *cli.Context) error {
	tagName, err := extractTagFromTag(tag)
	if err != nil {
		return err
	}
	if err := customPublish(); err != nil {
		return err
	}
	return publishImage(tagName, c)
}

func publishBranch(branch string, c *cli.Context) error {
	tagName, err := extractTagFromBranch(branch)
	if err != nil {
		return err
	}

	// don't run custom publish on non stable branches because custom publishes almost
	// always require some modification of the source code (e.g., pom.xml version update) to change
	// the identifier of the published artifact. We don't want to accidentally overwrite a previously
	// published artifact because the developer forgot to change the version number in source code.
	if !isStableBranch(branch) {
		logrus.Infof("skipping custom publish task because %q is not a stable branch", branch)
	} else {
		if err := customPublish(); err != nil {
			return err
		}
	}

	return publishImage(tagName, c)
}

// customPublish runs publish service found in template, if found
func customPublish() error {
	if helpers.DockerComposeHasService("publish") {
		return helpers.RunCommand(helpers.DockerComposeCommand("run", "--rm", "publish"))
	}
	logrus.Debug("no publish service found, skipping")
	return nil
}

// publishImage will publish the docker image if a Dockerfile is found
func publishImage(tagName string, c *cli.Context) error {
	if !helpers.HasDockerfile() {
		logrus.Debug("no Dockerfile found, skipping")
		return nil
	}
	// check required flags
	dockerImageName := c.String("docker-image-name")
	if len(dockerImageName) == 0 {
		return errors.New("you must use `--docker-image-name` to publish a docker image")
	}
	dockerRegistry := c.String("docker-registry")
	if len(dockerRegistry) == 0 {
		return errors.New("you must use `--docker-registry` to publish a docker image")
	}

	remoteSpec := fmt.Sprintf("%s/%s:%s", dockerRegistry, dockerImageName, tagName)
	return helpers.ChainCommands([]*exec.Cmd{
		exec.Command("docker", "tag", "-f", dockerImageName, remoteSpec),
		exec.Command("docker", "push", remoteSpec),
	})
}

var releaseTagRegexp = regexp.MustCompile(`^v\d+\.\d+\.\d(?:([-]).{0,120}|$)`)
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
func extractTagFromBranch(gitBranch string) (string, error) {
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
func extractTagFromTag(gitTag string) (string, error) {
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

func isStableBranch(gitBranch string) bool {
	if gitBranch == "origin/master" {
		return true
	} else if releaseRegexp.MatchString(gitBranch) {
		return true
	}
	return false
}
