package system

import (
  "errors"
  "fmt"
  "os"
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// file that we will use to verify volume mounts are working, assumption is that every lc repo should include this
var requiredFile = "/opt/project/lc.yml"

// CmdVerifyLds will ensure lds disk mounting is functioning
// This is mainly to address SAI-32.
func CmdVerifyLds(c *cli.Context) error {
  logrus.Debug("attempting to verify that the all lds components are functioning")
  cwd, err := os.Getwd()
  if err != nil {
    return fmt.Errorf("could not find current working directory to verify repo: %q", err)
  }

  volume := fmt.Sprintf("%v:/opt/project", cwd)
  fileCheck := fmt.Sprintf("if [ ! -e %s ]; then exit 1; fi", requiredFile)
  args := []string{"run", "--rm", "-v", volume, "--entrypoint=/bin/sh", "busybox", "-c", fileCheck}
  cmd := exec.Command("docker", args...)
  if err := helpers.RunCommand(cmd); err != nil {
    return errors.New("It appears that your local disk is not mounted into the LDS VM. See SAI-32 for instructions on how to fix.")
  }
  return nil
}
