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
	"github.com/codegangsta/cli"
)

func CmdRestart(c *cli.Context) error {
	if runningServer, err := runningServer(); err != nil {
		return err
	} else if err := CmdStop(c); err != nil {
		return err
	} else {
		var serviceToStart string
		if runningServer == "prodserver" {
			serviceToStart = "prodserver"
		} else {
			serviceToStart = "devserver"
		}
		if err := ensureServiceStarted(serviceToStart); err != nil {
			return err
		}
		return CmdStatus(c)
	}
}
