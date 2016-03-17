package command

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
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

	var service string

	if helpers.DockerComposeHasService("blackbox-test") {
		service = "blackbox-test"
	} else if helpers.DockerComposeHasService("smoketest") {
		service = "smoketest"
	}

	if service != "" {
		logrus.Infof("Running %s", service)

		if err := RunBlackboxTest(c); err != nil {
			return err
		}
	}

	logrus.Info("Running publish")
	return CmdPublish(c)
}
