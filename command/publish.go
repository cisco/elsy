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
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
)

// CmdPublish will publish all artifacts associated with the current repo
func CmdPublish(c *cli.Context) error {
	tagName, err := helpers.ExtractTag(c.String("git-tag"), c.String("git-branch"))
	if err != nil {
		return err
	}
	// don't run custom publish on non stable branches because custom publishes almost
	// always require some modification of the source code (e.g., pom.xml version update) to change
	// the identifier of the published artifact. We don't want to accidentally overwrite a previously
	// published artifact because the developer forgot to change the version number in source code.
	if len(c.String("git-tag")) > 0 || helpers.IsStableBranch(c.String("git-branch")) {
		if err := customPublish(tagName); err != nil {
			return err
		}
	} else {
		logrus.Infof("skipping custom publish task because HEAD is not tagged or tip of a stable branch")
	}

	return publishImage(tagName, c)
}

// customPublish runs publish service found in template, if found
// pass along tagName in ENV var LC_PUBLISH_DOCKER_TAG
func customPublish(tagName string) error {
	if helpers.DockerComposeHasService("publish") {
		cmd := helpers.DockerComposeCommand("run", "--rm", "publish")
		env := os.Environ()
		env = append(env, fmt.Sprintf("LC_PUBLISH_DOCKER_TAG=%s", tagName))
		cmd.Env = env
		return helpers.RunCommand(cmd)
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

	registries := c.StringSlice("docker-registry")

	// fail fast if a single publish fails
	for _, dockerRegistry := range registries {
		if len(dockerRegistry) == 0 {
			return errors.New("cannot publish to empty docker_registry, is `--docker-registry` specified?")
		}

		remoteSpec := fmt.Sprintf("%s/%s:%s", dockerRegistry, dockerImageName, tagName)
		err := helpers.ChainCommands([]*exec.Cmd{
			exec.Command("docker", "tag", dockerImageName, remoteSpec),
			exec.Command("docker", "push", remoteSpec),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
