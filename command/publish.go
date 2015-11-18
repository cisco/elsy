package command

import (
  "fmt"
  "os/exec"
  "regexp"
  "strings"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdPublish(c *cli.Context) error {
  gitBranch := c.String("git-branch")
  if len(gitBranch) == 0 {
    return fmt.Errorf("The publish task requires that the git branch to be set.")
  }
  if helpers.DockerComposeHasService("publish") {
    if isStableBranch(gitBranch) {
      if err := helpers.RunCommand(helpers.DockerComposeCommand("run", "--rm", "publish")); err != nil {
        return err
      }
    } else {
      logrus.Infof("skipping publish task because %q is not a stable branch", gitBranch)
    }
  }
  if helpers.HasDockerfile() {
    // check required flags
    dockerImageName := c.String("docker-image-name")
    if len(dockerImageName) == 0 {
      logrus.Panic("you must use `--docker-image-name` to publish a docker image")
    }
    dockerRegistry := c.String("docker-registry")
    if len(dockerRegistry) == 0 {
      logrus.Panic("you must use `--docker-registry` to publish a docker image")
    }

    tagName, err := extractTagFromBranch(c.String("git-branch"))
    if err != nil {
      logrus.Panic(err)
    }
    remoteSpec := fmt.Sprintf("%s/%s:%s", dockerRegistry, dockerImageName, tagName)
    return helpers.ChainCommands([]*exec.Cmd{
      exec.Command("docker", "tag", "-f", dockerImageName, remoteSpec),
      exec.Command("docker", "push", remoteSpec),
    })
  }
  return nil
}

var releaseRegexp = regexp.MustCompile("^origin/release/(.+)$")
var snapshotRegexp = regexp.MustCompile("^origin/(.+)$")
// regex for valid tag name taken from https://github.com/docker/distribution/blob/b07d759241defb2f345e95ed04bfdeb8ac010ab2/reference/regexp.go#L25
var validTagName = regexp.MustCompile(`^[\w][\w.-]{0,127}$`)
/*
*  extract tag name from branch
*  branch: `master` becomes tag `latest`
*  branch: `origin/release/xxx` becomes tag `xxx`
*  branch: `origin/aaa/xxx` becomes tag `snapshot.aaa.xxx`
*  branch: `origin/xxx` becomes tag `snapshot.xxx`
*/
func extractTagFromBranch(gitBranch string) (string, error) {
  var tagName string
  if gitBranch == "origin/master" {
    tagName = "latest"
  } else if matches := releaseRegexp.FindStringSubmatch(gitBranch); matches != nil {
    tagName = matches[1]
  } else if matches := snapshotRegexp.FindStringSubmatch(gitBranch); matches != nil {
    tagName = "snapshot."+matches[1]
  } else {
    return "", fmt.Errorf("could not determine tag from git branch: %q", gitBranch)
  }
  tagName = strings.Replace(tagName, "/", ".", -1)
  if !validTagName.MatchString(tagName) {
    return "", fmt.Errorf("tagName: %q is not valid", tagName)
  }
  return tagName, nil
}

func isStableBranch(gitBranch string) bool {
  if gitBranch == "origin/master" {
    return true
  } else if releaseRegexp.MatchString(gitBranch) {
    return true
  }
  return false
}
