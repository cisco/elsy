package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdStop(c *cli.Context) error {
	services := []string{}
	if helpers.DockerComposeHasService("devserver") {
		services = append(services, "devserver")
	}
	if helpers.DockerComposeHasService("prodserver") {
		services = append(services, "prodserver")
	}
	if len(services) > 0 {
		logrus.Info("Stopping server")
		cmd := helpers.DockerComposeCommand(append([]string{"kill"}, services...)...)
		return helpers.RunCommand(cmd)
	}
	return nil
}
