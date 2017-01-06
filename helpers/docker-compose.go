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

package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
)

// ComposeFileVersion inside the repo
type ComposeFileVersion int

const (
	unknown ComposeFileVersion = iota
	V1
	V2
)

func DockerComposeCommand(args ...string) *exec.Cmd {
	if _, err := os.Stat("docker-compose.yml"); err == nil {
		args = append([]string{"-f", "docker-compose.yml"}, args...)
	}

	if baseComposeFile := os.Getenv("LC_BASE_COMPOSE_FILE"); len(baseComposeFile) > 0 {
		args = append([]string{"-f", baseComposeFile}, args...)
	}

	return exec.Command(os.Getenv("DOCKER_COMPOSE_BINARY"), args...)
}

type DockerComposeService struct {
	Build string
	Image string
}

type DockerComposeNetwork struct {
	Driver string
}

type DockerComposeMap map[string]DockerComposeService
type DockerComposeNetworkMap map[string]DockerComposeNetwork

type DockerComposeV2 struct {
	Version  string
	Services DockerComposeMap
	Networks DockerComposeNetworkMap
}

func DockerComposeServices() (services []string) {
	if _, err := os.Stat("docker-compose.yml"); err == nil {
		for k := range getDockerComposeMap("docker-compose.yml") {
			services = append(services, k)
		}
	}
	if file := os.Getenv("LC_BASE_COMPOSE_FILE"); len(file) > 0 {
		for k := range getDockerComposeMap(file) {
			services = append(services, k)
		}
	}
	return
}

// GetDockerComposeVersion returns version of the docker-compose binary
// first return value is the human readable version
// second return value is an array of the {majorVersion, minorVersion, patchVersion}
func GetDockerComposeVersion(c *cli.Context) (string, []int, error) {
	out, err := RunCommandWithOutput(exec.Command(c.GlobalString("docker-compose"), "--version"))
	if err != nil {
		return "", nil, err
	}
	return parseDockerComposeVersion(out)
}

// GetComposeFileVersion of the current repo
// Will return the defaultVersion if it cannot determine the actual version
func GetComposeFileVersion(file string, defaultVersion ComposeFileVersion) ComposeFileVersion {
	version, _, err := parseComposeFile(file)
	if err != nil {
		logrus.Debugf("error parsing compose file found at %q, returning defaultVersion: %q, err: %q", file, defaultVersion, err)
		return defaultVersion
	}

	if version == unknown {
		logrus.Debugf("could not determine compospe version from file %q, returning defaultVersion: %q", file, defaultVersion)
		return defaultVersion
	}
	return version
}

func getDockerComposeMap(file string) DockerComposeMap {
	_, m, err := parseComposeFile(file)
	if err != nil {
		logrus.Errorf("error parsing compose file found at %q, err: %q", file, err)
		panic(err)
	}
	return m
}

func parseComposeFile(file string) (ComposeFileVersion, DockerComposeMap, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return unknown, nil, fmt.Errorf("no docker-comopse file found at %q", file)
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return unknown, nil, fmt.Errorf("error reading compose file at %q, err: %q", file, err)
	}

	var v1Contents DockerComposeMap
	if err := yaml.Unmarshal(b, &v1Contents); err == nil {
		logrus.Debugf("found compose v1 file at %q", file)
		return V1, v1Contents, nil
	}

	var v2Contents DockerComposeV2
	if err := yaml.Unmarshal(b, &v2Contents); err == nil {
		logrus.Debug("found compose v2 file at %q", file)
		return V2, v2Contents.Services, nil
	}

	return unknown, nil, fmt.Errorf("error reading compose file at %q, could not parse as V1 or V2 version", file)
}

func parseDockerComposeVersion(versionString string) (string, []int, error) {
	// assuming version is last word in string
	firstLine := strings.Split(versionString, "\n")[0]
	versionComponent := strings.Split(firstLine, ",")[0]
	words := strings.Split(versionComponent, " ")
	version := strings.TrimSpace(words[len(words)-1])

	versionComponents, err := parseVersionString(version)
	if err != nil {
		return "", nil, err
	}
	return version, versionComponents, nil
}

func DockerComposeHasService(service string) bool {
	for _, v := range DockerComposeServices() {
		if v == service {
			return true
		}
	}
	return false
}

func DockerComposeServiceIsRunning(serviceName string) (bool, error) {
	if containerId, err := DockerComposeServiceId(serviceName); err != nil {
		return false, err
	} else if len(containerId) == 0 {
		return false, nil
	} else if running, err := DockerContainerIsRunning(containerId); err != nil {
		return false, err
	} else {
		return running, nil
	}
}

func DockerComposeServiceId(serviceName string) (string, error) {
	cmd := DockerComposeCommand("ps", "-q", serviceName)
	if out, err := RunCommandWithOutput(cmd); err != nil {
		return "", err
	} else {
		containerId := strings.TrimSpace(out)
		if len(containerId) == 0 {
			return "", nil
		}
		return containerId, nil
	}
}
