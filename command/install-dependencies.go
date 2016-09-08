package command

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
)

func CmdInstallDependencies(c *cli.Context) error {
	if helpers.DockerComposeHasService("installdependencies") {
		return helpers.RunCommand(helpers.DockerComposeCommand("run", "--rm", "installdependencies"))
	} else {
		logrus.Debugf("no installdependencies service found, skipping installdependencies")
	}
	return nil
}
