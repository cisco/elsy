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

package server

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
)

func CmdStart(c *cli.Context) error {
	var serviceName = "devserver"
	if c.Bool("prod") {
		serviceName = "prodserver"
	}
	if err := ensureServiceStarted(serviceName); err != nil {
		return err
	} else {
		return CmdStatus(c)
	}
}

func ensureServiceStarted(serviceName string) error {
	if !helpers.DockerComposeHasService(serviceName) {
		return fmt.Errorf("no %q service defined", serviceName)
	} else if runningServer, err := runningServer(); err != nil {
		return nil
	} else if len(runningServer) > 0 {
		logrus.Infof("%q already running", runningServer)
		return nil
	}
	logrus.Infof("starting service %q", serviceName)
	cmd := helpers.DockerComposeCommand("up", "-d", serviceName)
	if err := helpers.RunCommand(cmd); err != nil {
		return err
	}
	return nil
}
