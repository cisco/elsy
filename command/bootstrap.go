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
	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
)

// CmdBootstrap pulls and builds the services in the docker-compose file
func CmdBootstrap(c *cli.Context) error {
	CmdTeardown(c)

	if !c.GlobalBool("offline") {
		if err := helpers.RunCommand(helpers.DockerComposeCommand("build", "--pull")); err != nil {
			return err
		}

		var args []string
		if c.GlobalBool("disable-parallel-pull") {
			args = append(args, "pull")
		} else {
			args = append(args, "pull", "--parallel")
		}

		if c.String("docker-image-name") != "" || len(c.StringSlice("local-images")) > 0 {
			excludes := c.StringSlice("local-images")
			if c.String("docker-image-name") != "" {
				excludes = append(excludes, c.String("docker-image-name"))
			}
			logrus.WithField("docker-image-name", excludes).Debug("not pulling services using repo's docker artifact")
			args = append(args, helpers.DockerComposeServicesExcluding(excludes)...)
		}

		pullCmd := helpers.DockerComposeCommand(args...)
		if err := helpers.RunCommand(pullCmd); err != nil {
			return err
		}
	}

	return CmdInstallDependencies(c)
}
