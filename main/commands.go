package main

import (
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command"
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
      Name:   "jenkins",
      Usage:  "",
      Action: panicOnError(command.CmdJenkins),
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
      Usage:  "",
      Action: panicOnError(command.CmdServer),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "smoketest",
      Usage:  "run smoketest service. forwards arguments",
      Action: panicOnError(command.CmdSmoketest),
      Flags:  []cli.Flag{},
    },
    {
      Name:   "teardown",
      Usage:  "",
      Action: panicOnError(command.CmdTeardown),
      Flags:  []cli.Flag{},
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

type cmdWithError func(c *cli.Context) error
func panicOnError(f cmdWithError) func(c *cli.Context) {
  return func(c *cli.Context) {
    if err := f(c); err != nil {
      panic(err)
    }
  }
}

func CommandNotFound(c *cli.Context, command string) {
  fmt.Fprintf(os.Stderr, "ERROR: %s: %q is not a valid command.\n\n", c.App.Name, command)
  cli.ShowAppHelp(c)
  os.Exit(2)
}
