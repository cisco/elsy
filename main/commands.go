/*
 *  Copyright 2016 Cisco Systems, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/cisco/elsy/command"
	"github.com/cisco/elsy/command/server"
	"github.com/cisco/elsy/command/system"
	"github.com/codegangsta/cli"
)

// GlobalFlags sets up flags on the lc command proper
func GlobalFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "project-name",
			Value:  GetConfigFileString("project_name"),
			Usage:  "the docker-compose project name. defaults to name of `root` option",
			EnvVar: "COMPOSE_PROJECT_NAME",
		},
		cli.StringFlag{
			Name:   "docker-compose",
			Value:  GetConfigFileStringWithDefault("docker_compose", "docker-compose"),
			Usage:  "the command to use for docker-compose",
			EnvVar: "LC_DOCKER_COMPOSE",
		},
		cli.StringFlag{
			Name:  "template",
			Value: GetConfigFileString("template"),
			Usage: "the project template to include",
		},
		cli.StringFlag{
			Name:  "template-image",
			Value: GetConfigFileString("template_image"),
			Usage: "the image to override what's in the template",
		},
		cli.BoolFlag{
			Name:   "enable-scratch-volumes",
			Usage:  "EXPERIMENTAL: if true, will put scratch resources in a data container; defaults to 'false'. Turn this on to speed up local builds.",
			EnvVar: "LC_ENABLE_SCRATCH_VOLUMES",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "turns on debug level logging",
			EnvVar: "LC_DEBUG",
		},
		cli.BoolFlag{
			Name:   "offline",
			Usage:  "will not attempt to pull any Docker images",
			EnvVar: "LC_OFFLINE",
		},
	}
}

// Commands sets up the main commands for the system
func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:   "bootstrap",
			Usage:  "Builds all local images and pulls remote images found in docker-compose.yml",
			Action: panicOnError(command.CmdBootstrap),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "docker-image-name",
					Value: GetConfigFileString("docker_image_name"),
					Usage: "local docker image name to publish",
				},
			},
		},
		{
			Name:   "init",
			Usage:  "Initializes an lc repo. If a directory is not provided as the first (and only) argument, then the current directory will be used.",
			Action: panicOnError(command.CmdInit),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "docker-image-name",
					Usage: "Will setup the lc repo using this name to tag the docker-image, only use this flag if the repo produces a Docker image.",
				},
				cli.StringSliceFlag{
					Name:  "docker-registry",
					Usage: "Will setup the lc repo to publish to this docker-registry, only use this flag if the repo produces a Docker image.",
				},
				cli.StringFlag{
					Name:  "project-name",
					Usage: "The value to use for the 'project_name' in the lc.yml file. If not found, the init command will generate this dynamically based on the directory (recommended).",
				},
				cli.StringFlag{
					Name:  "template",
					Usage: "The lc template to use for the repo (not required). Valid values are 'mvn', 'sbt'",
				},
			},
		},
		{
			Name:   "install-dependencies",
			Usage:  "Installs any dependencies the project has. relies on an `installdependencies` service in docker-compose.yml",
			Action: panicOnError(command.CmdInstallDependencies),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "ci",
			Usage:  "Builds, and possibly publishes, the project's artifact. used by the Jenkins job",
			Action: panicOnError(command.CmdCi),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "docker-image-name",
					Value: GetConfigFileString("docker_image_name"),
					Usage: "local docker image name to publish",
				},
				cli.StringSliceFlag{
					Name:  "docker-registry",
					Value: resolveDockerRegistryFromConfig(),
					Usage: "address of docker registry to publish to",
				},
				cli.StringFlag{
					Name:   "git-branch",
					Usage:  "Git branch which is being published",
					EnvVar: "GIT_BRANCH",
				},
				cli.StringFlag{
					Name:   "git-tag",
					Usage:  "Git tag which is being published",
					EnvVar: "GIT_TAG_NAME",
				},
				cli.StringFlag{
					Name:   "git-commit",
					Usage:  "Git commit that is being built",
					EnvVar: "GIT_COMMIT",
				},
				cli.StringFlag{
					Name:  "build-logs-dir",
					Value: GetConfigFileString("build_logs_dir"),
					Usage: "If populated, elsy will dump ALL docker-compose service logs into this directory, directory must be relative to the repo root.",
				},
			},
		},
		{
			Name:   "clean",
			Usage:  "Executes a project-specific `clean` command. Depends on a `clean` service in docker-compose.yml",
			Action: panicOnError(command.CmdClean),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "dc",
			Usage:  "Executes a specific docker-compose command",
			Action: panicOnError(command.CmdDockerCompose),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "run",
			Usage:  "Runs a service from the docker-compose YAML file",
			Action: panicOnError(command.CmdRun),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "make",
			Usage:  "Executes a specific make command. Depends on a `make` service in docker-compose.yml",
			Action: panicOnError(command.CmdMake),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "lein",
			Usage:  "Executes a specific Leiningen command. Depends on a `lein` service in docker-compose.yml",
			Action: panicOnError(command.CmdLein),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "mvn",
			Usage:  "Executes a specific Maven command. Depends on a `mvn` service in docker-compose.yml",
			Action: panicOnError(command.CmdMvn),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "sbt",
			Usage:  "Executes a specific Sbt command. Depends on a `sbt` service in docker-compose.yml",
			Action: panicOnError(command.CmdSbt),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "bower",
			Usage:  "Executes a specific Bower command. Depends on a `bower` service in docker-compose.yml",
			Action: panicOnError(command.CmdBower),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "npm",
			Usage:  "Executes a specific npm command. Depends on an `npm` service in docker-compose.yml",
			Action: panicOnError(command.CmdNpm),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "package",
			Usage:  "Packages the artifact using the `package` service in docker-compose.yml; if not present, will use Dockerfile",
			Action: panicOnError(command.CmdPackage),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "docker-image-name",
					Value: GetConfigFileString("docker_image_name"),
					Usage: "docker image name to create",
				},
				cli.BoolFlag{
					Name:  "skip-docker",
					Usage: "skip building of Dockerfile",
				},
				cli.BoolFlag{
					Name:  "skip-tests",
					Usage: "skip running of tests before packaging",
				},
				cli.StringFlag{
					Name:   "git-commit",
					Usage:  "Git commit that is being built",
					EnvVar: "GIT_COMMIT",
				},
			},
		},
		{
			Name:   "publish",
			Usage:  "Publishes the artifact to Artifactory, a Docker registry, etc., using the `publish` service in docker-compose.yml",
			Action: panicOnError(command.CmdPublish),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "docker-image-name",
					Value: GetConfigFileString("docker_image_name"),
					Usage: "local docker image name to publish",
				},
				cli.StringSliceFlag{
					Name:  "docker-registry",
					Value: resolveDockerRegistryFromConfig(),
					Usage: "address of docker registry to publish to",
				},
				cli.StringFlag{
					Name:   "git-branch",
					Usage:  "Git branch which is being published",
					EnvVar: "GIT_BRANCH",
				},
				cli.StringFlag{
					Name:   "git-tag",
					Usage:  "Git tag which is being published",
					EnvVar: "GIT_TAG_NAME",
				},
			},
		},
		{
			Name:   "release",
			Usage:  "Creates a release tag for the current repo",
			Action: panicOnError(command.CmdRelease),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "git-commit",
					Usage: "commit to tag",
				},
				cli.StringFlag{
					Name:  "version",
					Usage: "version to release, must be of the format vX.Y.Z[-Q], where X, Y, and Z are ints and Q is a string qualifier.",
				},
			},
		},
		{
			Name:  "server",
			Usage: "Manages the project's server (default is devserver)",
			Subcommands: []cli.Command{
				{
					Name:   "status",
					Usage:  "Gets status of server. exits 0 if up, non-zero if down. prints out status as well as dynamic ports",
					Action: panicOnError(server.CmdStatus),
				},
				{
					Name:   "start",
					Usage:  "Starts the devserver or prodserver",
					Action: panicOnError(server.CmdStart),
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "prod, p",
							Usage: "operate on the production server",
						},
					},
				},
				{
					Name:   "stop",
					Usage:  "Stops any running devserver or prodserver",
					Action: panicOnError(server.CmdStop),
				},
				{
					Name:   "restart",
					Usage:  "Calls stop then start",
					Action: panicOnError(server.CmdRestart),
				},
				{
					Name:   "log",
					Usage:  "Follows the log of the running server",
					Action: panicOnError(server.CmdLog),
				},
			},
		},
		{
			Name:    "blackbox-test",
			Aliases: []string{"bbtest"},
			Usage:   "Runs blackbox-test service. forwards arguments",
			Action:  panicOnError(command.CmdBlackbox),
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "skip-package",
					Usage: "do not run package service prior to executing blackbox tests",
				},
				cli.StringFlag{
					Name:  "docker-image-name",
					Value: GetConfigFileString("docker_image_name"),
					Usage: "docker image name to create",
				},
			},
		},
		{
			Name:   "teardown",
			Usage:  "Kills all running services and removes services that do not have gc protection",
			Action: panicOnError(command.CmdTeardown),
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "will remove all services, even those with gc protection",
				},
			},
		},
		{
			Name:   "test",
			Usage:  "Executes project's `test` service. this should run the unit tests",
			Action: panicOnError(command.CmdTest),
			Flags:  []cli.Flag{},
		},
		{
			Name:  "system",
			Usage: "Manages lc itself",
			Subcommands: []cli.Command{
				{
					Name:   "view-template",
					Usage:  "Displays the YAML of a template",
					Action: panicOnError(system.CmdViewTemplate),
					Flags:  []cli.Flag{},
				},
				{
					Name:   "verify-install",
					Usage:  "Runs a series of checks to verify that docker is running correctly. This must be run inside a repo",
					Action: panicOnError(system.CmdVerifyInstall),
					Flags:  []cli.Flag{},
				},
				{
					Name:   "list-templates",
					Usage:  "Displays the name of all available templates",
					Action: panicOnError(system.CmdListTemplates),
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
			if c.GlobalBool("debug") {
				panic(err)
			} else {
				logrus.Error(err)
				logrus.Error("command failed. use --debug to see full stacktrace")
			}
		}
	}
}

// CommandNotFound knows what to do when a command isn't found
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "ERROR: %s: %q is not a valid command.\n\n", c.App.Name, command)
	cli.ShowAppHelp(c)
	os.Exit(2)
}

// resolveDockerRegistryFromConfig handles resolving the docker registry (or set of registries)
// in a backwards compatible way.
//
// Spcificially (because of backwards compatibility reasons) we support two config fields:
// 	- 'docker_registry' -> will hold a single string
// 	- 'docker_registries' -> will hold a yml sequece
//
// 	This function will panic if both fields are defined.
func resolveDockerRegistryFromConfig() *cli.StringSlice {
	singleK := "docker_registry"
	seqK := "docker_registries"

	if IsKeyInConfig(singleK) && IsKeyInConfig(seqK) {
		panic(fmt.Errorf("Error parsing 'lc.yml': multiple docker registry configs found, pick either %q or %q", singleK, seqK))
	}

	if IsKeyInConfig(singleK) {
		return &cli.StringSlice{GetConfigFileString(singleK)}
	}

	if IsKeyInConfig(seqK) {
		v := cli.StringSlice(GetConfigFileSlice(seqK))
		if len(v) == 0 {
			// this will happen if the yaml containing the sequence is not perfectly formatted (e.g., if '-value' instead of '- value')
			// eventually we need to make our parsing logic more forgiving, but until then just make it crystal clear when we can't parse something.
			panic(fmt.Errorf("Error parsing 'lc.yml': found %q key, but did not find any registries, verify that yaml is correct", seqK))
		}

		return &v
	}

	return &cli.StringSlice{}
}
