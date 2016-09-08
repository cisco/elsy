package server

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
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
