package main

import (
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "lancope"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.Before = func(c *cli.Context) error {
		preReqCheck()
		return nil
	}
	app.Run(os.Args)
}

func preReqCheck() {
	// TODO: replace this with checking presence and version of local-docker-stack
	if _, err := exec.LookPath("docker"); err != nil {
		panic("could not find docker, please install local-docker-stack")
	}
	if _, err := exec.LookPath("docker-compose"); err != nil {
		panic("could not find docker-compose, please install local-docker-stack")
	}
}
