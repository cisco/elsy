package command

import (
  "fmt"
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

// CmdRelease will create, and push a release tag
func CmdRelease(c *cli.Context) error {
  version := c.String("version")
  commit := c.String("git-commit")
  if len(version) == 0 {
    return fmt.Errorf("--version flag required")
  }
  if len(commit) == 0 {
    return fmt.Errorf("--git-commit flag required")
  }
  if err := checkTag(version); err != nil {
    return err
  }

  // TODO: we might want to allow a '-f' option to support re-running a tag build
  // since if a user pushes a tag and the build fails, it is not simple to rerun that build without
  // repushing the tag
  logrus.Infof("creating, and pushing, git tag %s at commit %s", version, commit)
  return helpers.ChainCommands([]*exec.Cmd{
    exec.Command("git", "tag", "-a", version, commit, "-m", fmt.Sprintf("add release tag for %s", version)),
    exec.Command("git", "push", "origin", version),
  })
}

func checkTag(v string)  error{
  if match := releaseTagRegexp.MatchString(v); !match {
    return fmt.Errorf("release value syntax was not valid, it must adhere to: %q", releaseTagRegexp)
  }
  return nil
}
