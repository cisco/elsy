package main

import (
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/command"
)

var GlobalFlags = []cli.Flag{
  cli.StringFlag{
    Name:  "docker-compose",
    Value: "docker-compose.yml",
    Usage: "path to docker-compose.yml",
  },
}

var Commands = []cli.Command{
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
    Flags:  []cli.Flag{},
  },
  {
    Name:   "publish",
    Usage:  "",
    Action: command.CmdPublish,
    Flags:  []cli.Flag{},
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

func CommandNotFound(c *cli.Context, command string) {
  fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
  os.Exit(2)
}
