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
	"github.com/cisco/elsy/helpers"
)

// CmdCi runs ci loop
func CmdCi(c *cli.Context) error {
	defer CmdTeardown(c)

	logrus.Info("Running bootstrap")
	if err := CmdBootstrap(c); err != nil {
		return err
	}

	if helpers.DockerComposeHasService("test") {
		logrus.Info("Running test")
		if err := CmdTest(c); err != nil {
			return err
		}
	}

	logrus.Info("Running package")
	if err := RunPackage(c); err != nil {
		return err
	}

	if helpers.DockerComposeHasService("blackbox-test") {
		logrus.Infof("Running blackbox-test")
		if err := RunBlackboxTest(c); err != nil {
			return err
		}
	}

  if helpers.DockerComposeHasService("publish") || helpers.HasDockerfile() {
    logrus.Info("Running publish")
    return CmdPublish(c)
  } else {
    logrus.Info("No publish service defined, and no Dockerfile present.")
    return nil
  }
}
