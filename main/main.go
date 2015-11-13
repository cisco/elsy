package main

import (
  "os/exec"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func main() {
  if err := LoadConfigFile(".lc.yaml"); err != nil {
    panic(err)
  }

  app := cli.NewApp()
  app.Name = Name
  app.Version = Version
  app.Author = "lancope"
  app.Email = ""
  app.Usage = ""

  app.Flags = GlobalFlags()
  app.Commands = Commands()
  app.CommandNotFound = CommandNotFound
  app.Before = func(c *cli.Context) error {
    setLogLevel(c)
    preReqCheck(c)
    helpers.DockerComposeBeforeHook(c)
    return nil
  }
  app.After = func(c *cli.Context) error {
    if err := helpers.DockerComposeAfterHook(c); err != nil {
      return err
    }
    return nil
  }
  app.RunAndExitOnError()
}

func preReqCheck(c *cli.Context) {
  // TODO: replace this with checking presence and version of local-docker-stack
  if _, err := exec.LookPath("docker"); err != nil {
    logrus.Panic("could not find docker, please install local-docker-stack")
  }
  dockerCompose := c.GlobalString("docker-compose")
  if _, err := exec.LookPath(dockerCompose); err != nil {
    logrus.Panicf("could not find docker compose binary: %q, please install local-docker-stack", dockerCompose)
  }
}

func setLogLevel(c *cli.Context) {
  if c.GlobalBool("debug") {
    logrus.SetLevel(logrus.DebugLevel)
  } else {
    logrus.SetLevel(logrus.InfoLevel)
  }
}
