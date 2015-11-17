package main

import (
  "fmt"
  "os"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command/server"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command/system"
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
      Action: panicOnError(command.CmdBootstrap),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "install",
      Usage:  "",
      Action: panicOnError(command.CmdInstallDependencies),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "ci",
      Usage:  "",
      Action: panicOnError(command.CmdCi),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "dc",
      Usage:  "",
      Action: panicOnError(command.CmdDockerCompose),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "mvn",
      Usage:  "",
      Action: panicOnError(command.CmdMvn),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "sbt",
      Usage:  "",
      Action: panicOnError(command.CmdSbt),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "bower",
      Usage:  "",
      Action: panicOnError(command.CmdBower),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "npm",
      Usage:  "",
      Action: panicOnError(command.CmdNpm),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "package",
      Usage:  "",
      Action: panicOnError(command.CmdPackage),
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
      Action: panicOnError(command.CmdPublish),
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
      Usage:  "manage the project's server (default is devserver)",
      Subcommands: []cli.Command{
        {
          Name: "status",
          Usage: "get status of server. exits 0 if up, non-zero if down. prints out status as well as dynamic ports",
          Action: panicOnError(server.CmdStatus),
        },
        {
          Name: "start",
          Usage: "start the devserver or prodserver",
          Action: panicOnError(server.CmdStart),
          Flags:  []cli.Flag{
            cli.BoolFlag{
              Name:  "prod, p",
              Usage: "operate on the production server",
            },
          },
        },
        {
          Name: "stop",
          Usage: "stops any running devserver or prodserver",
          Action: panicOnError(server.CmdStop),
        },
        {
          Name: "restart",
          Usage: "calls stop then start",
          Action: panicOnError(server.CmdRestart),
        },
        {
          Name: "log",
          Usage: "follows the log of the running server",
          Action: panicOnError(server.CmdLog),
        },
      },
    },
    {
      Name:   "smoketest",
      Usage:  "run smoketest service. forwards arguments",
      Action: panicOnError(command.CmdSmoketest),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "teardown",
      Usage:  "kill all running containers and remove containers that do not have gc protection",
      Action: panicOnError(command.CmdTeardown),
      Flags:  []cli.Flag{
        cli.BoolFlag{
          Name:  "force, f",
          Usage: "will remove all containers, even those with gc protection",
        },
      },
    },
    {
      Name:   "test",
      Usage:  "",
      Action: panicOnError(command.CmdTest),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "system",
      Usage:  "commands for managing lc",
      Subcommands: []cli.Command{
        {
          Name:  "upgrade",
          Usage: "upgrade this lc binary",
          Action: panicOnError(system.CmdUpgrade),
          Flags:  []cli.Flag{},
        },
      },
    },
  }
}

var CommandSuccess = true
type cmdWithError func(c *cli.Context) error
func panicOnError(f cmdWithError) func(c *cli.Context) {
  return func(c *cli.Context) {
    if err := f(c); err != nil {
      CommandSuccess = false
      if c.GlobalBool("debug"){
        panic(err)
      } else {
        logrus.Error(err)
        logrus.Error("command failed. use --debug to see full stacktrace")
      }
    }
  }
}

func CommandNotFound(c *cli.Context, command string) {
  fmt.Fprintf(os.Stderr, "ERROR: %s: %q is not a valid command.\n\n", c.App.Name, command)
  cli.ShowAppHelp(c)
  os.Exit(2)
}
