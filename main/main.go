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
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/elsy/helpers"
	"github.com/elsy/template"
)

func main() {
	if err := LoadConfigFile("lc.yml"); err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "lc"
	app.Version = helpers.BuildVersionString()
	app.Author = "Cisco"
	app.Usage = "an opinionated, multi-language, build-tool based on Docker and Docker Compose"

	app.Flags = GlobalFlags()
	app.Commands = Commands()
	app.CommandNotFound = CommandNotFound
	app.Before = beforeHook
	app.After = afterHook
	app.RunAndExitOnError()

	if !CommandSuccess {
		os.Exit(1)
	}
}

func beforeHook(c *cli.Context) error {
	setLogLevel(c)
	preReqCheck(c)
	setComposeBinary(c)
	setComposeProjectName(c)
	setComposeTemplate(c)
	addSignalListener()
	return nil
}

func afterHook(c *cli.Context) error {
	return removeComposeTemplate()
}

func addSignalListener() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		removeComposeTemplate()
		os.Exit(2)
	}()
}

func removeComposeTemplate() error {
	// clean up compose template if it exists
	if file := os.Getenv("LC_BASE_COMPOSE_FILE"); len(file) > 0 {
		logrus.Debugf("attempting to remove base compose file: %v", file)
		if err := os.Remove(file); err != nil {
			return err
		}
	}
	return nil
}

func setLogLevel(c *cli.Context) {
	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func preReqCheck(c *cli.Context) {
	if len(c.Args()) == 0 || c.Args()[0] == "system" {
		// system commands do not need docker
		return
	}

	// TODO: replace this with checking presence and version of local-docker-stack
	if _, err := exec.LookPath("docker"); err != nil {
		logrus.Fatal("could not find docker, please install local-docker-stack")
	}
	dockerCompose := c.GlobalString("docker-compose")
	if _, err := exec.LookPath(dockerCompose); err != nil {
		logrus.Fatalf("could not find docker compose binary: %q, please install local-docker-stack", dockerCompose)
	}

	if versionString, versionComponents, err := helpers.GetDockerComposeVersion(c); err != nil {
		logrus.Warnf("failed checking docker-compose version. Note that lc only supports docker-compose 1.5.0 or higher")
	} else {
		major, minor := versionComponents[0], versionComponents[1]
		// assuming we won't see any docker-compose versions less than 1.x
		if major == 1 && minor < 5 {
			logrus.Fatalf("found docker-compose version %s, lc only supports docker-compose 1.5.0 or higher", versionString)
		}
	}

	if err := helpers.EnsureDockerConnectivity(); err != nil {
		ip, _ := helpers.DockerIp()
		logrus.Fatalf("could not connect to docker daemon at %q, err: %q.", ip, err)

	}
}

func setComposeBinary(c *cli.Context) {
	os.Setenv("DOCKER_COMPOSE_BINARY", c.GlobalString("docker-compose"))
}

func setComposeProjectName(c *cli.Context) {
	var invalidChars = regexp.MustCompile("[^a-z0-9]")
	projectName := c.GlobalString("project-name")
	if len(projectName) == 0 {
		logrus.Debug("using current working directory for compose project name")
		path, _ := os.Getwd()
		projectName = filepath.Base(path)
	} else {
		logrus.Debugf("using configured value: %q for project name", projectName)
	}
	projectName = invalidChars.ReplaceAllString(strings.ToLower(projectName), "")
	os.Setenv("COMPOSE_PROJECT_NAME", projectName)
}

func setComposeTemplate(c *cli.Context) {
	templateName := c.GlobalString("template")
	enableScratchVolume := c.GlobalBool("enable-scratch-volumes")
	if len(templateName) > 0 {
		if yaml, err := template.GetTemplate(templateName, enableScratchVolume); err == nil {
			file := createTempComposeFile(yaml)
			logrus.Debugf("setting LC_BASE_COMPOSE_FILE to %v", file)
			os.Setenv("LC_BASE_COMPOSE_FILE", file)
		} else {
			logrus.Panicf("error finding template %q, err: %q", templateName, err)
		}

		dataContainers := template.GetSharedExternalDataContainers(templateName)
		for _, dataContainer := range dataContainers {
			if err := dataContainer.Ensure(c.GlobalBool("offline")); err != nil {
				logrus.Panic("unable to create data container")
			}
		}
	}
}

func createTempComposeFile(yaml string) string {
	cwd, _ := os.Getwd()
	fh, err := ioutil.TempFile(cwd, "lc_docker_compose_template")
	if err != nil {
		logrus.Panic("could not create temporary yaml file")
	}
	defer fh.Close()
	_, err = fh.WriteString(yaml)
	if err != nil {
		logrus.Panic("could not write to temporary yaml file")
	}
	return fh.Name()
}
