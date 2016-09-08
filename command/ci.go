package command

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
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

	logrus.Info("Running publish")
	return CmdPublish(c)
}
