package command

import (
	"fmt"
	"regexp"

	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command/system"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdBootstrap(c *cli.Context) error {
	if err := system.CmdVerifyLds(c); err != nil {
		return err
	}
	CmdTeardown(c)

	if err := helpers.RunCommand(helpers.DockerComposeCommand("build", "--pull")); err != nil {
		return err
	}
	pullCmd := helpers.DockerComposeCommand("pull", "--ignore-pull-failures")
	benignError := regexp.MustCompile(fmt.Sprintf(`Error: image library/%s:latest not found`, c.String("docker-image-name")))
	helpers.RunCommandWithFilter(pullCmd, benignError.MatchString)
	return CmdInstallDependencies(c)
}
