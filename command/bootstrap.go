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
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

// CmdBootstrap pulls and builds the services in the docker-compose file
func CmdBootstrap(c *cli.Context) error {
	CmdTeardown(c)

	if !c.GlobalBool("offline") {
		if err := helpers.RunCommand(helpers.DockerComposeCommand("build", "--pull")); err != nil {
			return err
		}
		pullCmd := helpers.DockerComposeCommand("pull", "--ignore-pull-failures")
		benignError := regexp.MustCompile(fmt.Sprintf(`Error: image library/%s(:latest|) not found`, c.String("docker-image-name")))
		helpers.RunCommandWithFilter(pullCmd, benignError.MatchString)
	}

	return CmdInstallDependencies(c)
}
