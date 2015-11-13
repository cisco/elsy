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
      EnvVar: "COMPOSE_PROJECT_NAME",
    },
    cli.StringFlag{
      Name:  "docker-compose",
      Value: GetConfigFileStringWithDefault("docker_compose", "docker-compose"),
      Usage: "command to use for docker-compose",
      EnvVar: "LC_DOCKER_COMPOSE",
    },
    cli.StringFlag{
      Name:  "template",
      Value: GetConfigFileString("template"),
      Usage: "project template to include",
    },
    cli.BoolFlag{
      Name:  "debug",
      Usage: "turn on debug level logging",
      EnvVar: "LC_DEBUG",
    },
  }
}

func Commands() []cli.Command {
  return []cli.Command{
    {
      Name:   "bootstrap",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdBootstrap(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "install",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdInstallDependencies(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "ci",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdCi(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "dc",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdDockerCompose(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "jenkins",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdJenkins(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "mvn",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdMvn(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "sbt",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdSbt(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "bower",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdBower(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "npm",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdNpm(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "package",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdPackage(c) },
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
      Action: func(c *cli.Context) { command.CmdPublish(c) },
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
      Action: func(c *cli.Context) { command.CmdServer(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "smoketest",
      Usage:  "run smoketest service. forwards arguments",
      Action: func(c *cli.Context) { command.CmdSmoketest(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "teardown",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdTeardown(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "test",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdTest(c) },
      Flags:  []cli.Flag{},
    },
    {
      Name:   "upgrade",
      Usage:  "",
      Action: func(c *cli.Context) { command.CmdUpgrade(c) },
      Flags:  []cli.Flag{},
    },
  }
}

func CommandNotFound(c *cli.Context, command string) {
  fmt.Fprintf(os.Stderr, "ERROR: %s: %q is not a valid command.\n\n", c.App.Name, command)
  cli.ShowAppHelp(c)
  os.Exit(2)
}
