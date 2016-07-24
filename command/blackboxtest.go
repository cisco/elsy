package command

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdBlackbox processes cmd args and then runs blackbox tests
func CmdBlackbox(c *cli.Context) error {
	if !c.Bool("skip-package") {
		logrus.Info("Running package before executing blackbox tests")
		if err := RunPackage(c); err != nil {
			return err
		}
	}

	return RunBlackboxTest(c)
}

// RunBlackboxTest will execute the blackbox service and then return
func RunBlackboxTest(c *cli.Context) error {
	service := "blackbox-test"
	args := append([]string{"run", "--rm", service}, c.Args()...)
	return helpers.RunCommand(helpers.DockerComposeCommand(args...))
}
