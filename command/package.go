package command

import (
	"os/exec"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command/system"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdPackage runs package service if present and then attempts to build Dockerfile.
// Unless --skip-tests is passed, it *will* run the tests, and any failures will abort
// the packaging process.
func CmdPackage(c *cli.Context) error {
	if err := system.CmdVerifyLds(c); err != nil {
		return err
	}

	if !c.Bool("skip-tests") && helpers.DockerComposeHasService("test") {
		logrus.Info("Running tests before packaging")
		if err := CmdTest(c); err != nil {
			return err
		}
	}

	return RunPackage(c)
}

// RunPackage runs package service if present and then attempts to build Dockerfile
// This command does *not* attempt to run any tests, nor does it pay attention to the -skip-test flag
func RunPackage(c *cli.Context) error {
	if helpers.DockerComposeHasService("package") {
		err := helpers.RunCommand(helpers.DockerComposeCommand("run", "--rm", "package"))
		if err != nil {
			return err
		}
	} else {
		logrus.Debug("no package service found, skipping")
	}

	// docker build
	commands := []*exec.Cmd{}
	dockerImageName := c.String("docker-image-name")
	if helpers.HasDockerfile() && !c.Bool("skip-docker") {
		logrus.Debug("detected Dockerfile for packaging")

		if !c.GlobalBool("offline") {
			if image, err := helpers.DockerImage("Dockerfile"); err == nil && image.IsRemote() {
				commands = append(commands, exec.Command("docker", "pull", image.String()))
			}
		}

		if len(dockerImageName) == 0 {
			logrus.Panic("you must use `--docker-image-name` to package a docker image")
		}

		commands = append(commands, exec.Command("docker", "build", "-t", dockerImageName, "."))
	}

	if err := helpers.ChainCommands(commands); err != nil {
		return err
	}

	if helpers.HasDockerfile() && !c.Bool("skip-docker") {
		// remove any containers that were created from the previous version of the image
		if err := helpers.RemoveContainersOfImage(dockerImageName); err != nil {
			logrus.Warnf("could not remove containers created from previous version of %q, err: %q", dockerImageName, err)
		}
	}

	return nil
}
