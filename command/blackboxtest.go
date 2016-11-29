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
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

// CmdBlackbox processes cmd args and then runs blackbox tests
func CmdBlackbox(c *cli.Context) error {
	if !c.Bool("skip-package") {
		logrus.Info("Running package before executing blackbox tests")
		if err := RunPackage(c); err != nil {
			return err
		}
	}

	return RunBlackboxTest(c)
}

// RunBlackboxTest will execute the blackbox service and then return
func RunBlackboxTest(c *cli.Context) error {
	service := "blackbox-test"
	args := append([]string{"run", "--rm", service}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
