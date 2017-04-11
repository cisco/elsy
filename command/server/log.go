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

	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
)

func CmdLog(c *cli.Context) error {
	runningServer, err := runningServer()
	if err != nil {
		return err
	} else if len(runningServer) == 0 {
		return fmt.Errorf("Server is not running")
	}

	logrus.Info("Press Ctrl-C to stop...")
	cmd := helpers.DockerComposeCommand("logs", runningServer)
	if err := helpers.RunCommand(cmd); err != nil {
		return err
	}
	return nil
}
