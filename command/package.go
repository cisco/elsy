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
	"fmt"
	"os/exec"

	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
)

// commitLabel identifies the git commit the image was built from
const commitLabel = "com.elsy.metadata.git-commit"

// repoLabel identifies the git repo the image was built from
const repoLabel = "com.elsy.metadata.git-url"

// CmdPackage runs package service if present and then attempts to build Dockerfile.
// Unless --skip-tests is passed, it *will* run the tests, and any failures will abort
// the packaging process.
func CmdPackage(c *cli.Context) error {
	if !c.Bool("skip-tests") && helpers.DockerComposeHasService("test") {
		logrus.Info("Running tests before packaging")
		if err := CmdTest(c); err != nil {
			return err
		}
	}

	return RunPackage(c)
}

// RunPackage runs package service if present and then attempts to build Dockerfile
// This command does *not* attempt to run any tests, nor does it pay attention to the -skip-test flag
func RunPackage(c *cli.Context) error {
	if helpers.DockerComposeHasService("package") {
		err := helpers.RunCommand(helpers.DockerComposeCommand("run", "--rm", "package"))
		if err != nil {
			return err
		}
	} else {
		logrus.Debug("no package service found, skipping")
	}

	// docker build
	commands := []*exec.Cmd{}
	dockerImageName := c.String("docker-image-name")
	if helpers.HasDockerfile() && !c.Bool("skip-docker") {
		logrus.Debug("detected Dockerfile for packaging")

		if !c.GlobalBool("offline") {
			if image, err := helpers.DockerImage("Dockerfile"); err == nil && image.IsRemote() {
				commands = append(commands, exec.Command("docker", "pull", image.String()))
			}
		}

		if len(dockerImageName) == 0 {
			logrus.Panic("you must use `--docker-image-name` to package a docker image")
		}

		buildArgs := []string{"build", "-t", dockerImageName}
		labelArgs := constructLabelArgs(c)
		if len(labelArgs) > 0 {
			buildArgs = append(buildArgs, labelArgs...)
		}
		buildArgs = append(buildArgs, ".")
		commands = append(commands, exec.Command("docker", buildArgs...))
	}

	if err := helpers.ChainCommands(commands); err != nil {
		return err
	}

	if helpers.HasDockerfile() && !c.Bool("skip-docker") {
		// remove any containers that were created from the previous version of the image
		if err := helpers.RemoveContainersOfImage(dockerImageName); err != nil {
			logrus.Warnf("could not remove containers created from previous version of %q, err: %q", dockerImageName, err)
		}
	}

	return nil
}

func constructLabelArgs(c *cli.Context) (labelArgs []string) {
	versionString, version, err := helpers.GetDockerVersion()
	if err != nil {
		logrus.Warnf("Skipping applying image labels: could not determine docker version, err: %q", err)
		return nil
	}

	// 'docker build --label' was introduced in docker 1.11.1: https://github.com/docker/docker/releases/tag/v1.11.1-rc1
	// assuming we won't see any docker versions less than 1.x
	major, minor, patch := version[0], version[1], version[2]
	if major == 1 && (minor < 11 || (minor == 11 && patch < 1)) {
		logrus.Debugf("Skipping applying image labels: found docker version %s, 'docker build --label' only supported Docker 1.11.1 and higher", versionString)
		return nil
	}

	commit := c.String("git-commit")
	if commit != "" {
		logrus.Infof("Attaching image label: %s=%s", commitLabel, commit)
		labelArgs = append(labelArgs, "--label", fmt.Sprintf("%s=%s", commitLabel, commit))
	}

	gitUrl := c.String("git-url")
	if gitUrl != "" {
		logrus.Infof("Attaching image label: %s=%s", repoLabel, gitUrl)
		labelArgs = append(labelArgs, "--label", fmt.Sprintf("%s=%s", repoLabel, gitUrl))
	}

	return
}
