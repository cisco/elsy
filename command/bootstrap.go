package command

import (
  "os"
  "os/exec"

  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdBootstrap(c *cli.Context) {
  images := make(map[string]bool)
  for _, dockerCompose := range helpers.DockerCompose(c.GlobalString("root")) {
    if len(dockerCompose.Image) > 0 {
      images[dockerCompose.Image] = true
    }
  }
  for _, image := range helpers.DockerfileImages(c.GlobalString("root")) {
    images[image] = true
  }
  for image := range images {
    cmd := exec.Command("docker", "pull", image)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
  }

  helpers.DockerComposeExec("build")
}
