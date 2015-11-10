package main

import (
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command"
)

func GlobalFlags() []cli.Flag {
  return []cli.Flag{
    cli.StringFlag{
      Name:  "project-name",
      Value: GetConfigFileString("project_name"),
      Usage: "docker-compose project name. defaults to name of `root` option",
    },
    cli.StringFlag{
      Name:  "docker-compose",
      Value: GetConfigFileStringWithDefault("docker_compose", "docker-compose"),
      Usage: "command to use for docker-compose",
    },
    cli.StringFlag{
      Name:  "template",
      Value: GetConfigFileString("template"),
      Usage: "project template to include",
    },
    cli.BoolFlag{
      Name:  "debug",
      Usage: "turn on debug level logging",
    },
  }
}

func Commands() []cli.Command {
  return []cli.Command{
    {
      Name:   "bootstrap",
      Usage:  "",
      Action: command.CmdBootstrap,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "ci",
      Usage:  "",
      Action: command.CmdCi,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "dc",
      Usage:  "",
      Action: command.CmdDockerCompose,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "jenkins",
      Usage:  "",
      Action: command.CmdJenkins,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "mvn",
      Usage:  "",
      Action: command.CmdMvn,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "sbt",
      Usage:  "",
      Action: command.CmdSbt,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "bower",
      Usage:  "",
      Action: command.CmdBower,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "npm",
      Usage:  "",
      Action: command.CmdNpm,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "package",
      Usage:  "",
      Action: command.CmdPackage,
      Flags:  []cli.Flag{
        cli.StringFlag{
          Name:  "docker-image-name",
          Value: GetConfigFileString("docker_image_name"),
          Usage: "docker image name to create",
        },
      },
    },
    {
      Name:   "publish",
      Usage:  "",
      Action: command.CmdPublish,
      Flags:  []cli.Flag{
        cli.StringFlag{
          Name:  "docker-image-name",
          Value: GetConfigFileString("docker_image_name"),
          Usage: "local docker image name to publish",
        },
        cli.StringFlag{
          Name:  "docker-registry",
          Value: GetConfigFileString("docker_registry"),
          Usage: "address of docker registry to publish to",
        },
      },
    },
    {
      Name:   "server",
      Usage:  "",
      Action: command.CmdServer,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "teardown",
      Usage:  "",
      Action: command.CmdTeardown,
      Flags:  []cli.Flag{},
    },
    {
      Name:   "test",
      Usage:  "",
      Action: command.CmdTest,
      Flags:  []cli.Flag{},
    },
  }
}

func CommandNotFound(c *cli.Context, command string) {
  fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
  os.Exit(2)
}
