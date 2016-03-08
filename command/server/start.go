package server

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
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
