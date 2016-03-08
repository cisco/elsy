package server

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
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
