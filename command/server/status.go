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
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func CmdStatus(c *cli.Context) error {
	if service, err := runningServer(); err != nil {
		return err
	} else if len(service) == 0 {
		return fmt.Errorf("Server is not running")
	} else {
		logrus.Infof("Server is running using %q service", service)
		showDynamicPorts(service)
	}
	return nil
}

func runningServer() (string, error) {
	hasDevserver := helpers.DockerComposeHasService("devserver")
	hasProdserver := helpers.DockerComposeHasService("prodserver")
	if !hasDevserver && !hasProdserver {
		return "", fmt.Errorf("No devserver or prodserver service defined")
	}

	if hasDevserver {
		if running, err := helpers.DockerComposeServiceIsRunning("devserver"); err != nil {
			return "", err
		} else if running {
			return "devserver", nil
		}
	}

	if hasProdserver {
		if running, err := helpers.DockerComposeServiceIsRunning("prodserver"); err != nil {
			return "", err
		} else if running {
			return "prodserver", nil
		}
	}

	return "", nil
}

func showDynamicPorts(serviceName string) error {
	if containerId, err := helpers.DockerComposeServiceId(serviceName); err != nil {
		return err
	} else if portBindings, err := helpers.DockerContainerDyanmicPorts(containerId); err != nil {
		return err
	} else if len(portBindings) == 0 {
		return nil
	} else if dockerIp, err := helpers.DockerIp(); err != nil {
		return err
	} else {
		for port, binding := range portBindings {
			logrus.Info(color.GreenString("%s is available at %s:%s", port, dockerIp, binding))
		}
	}
	return nil
}
