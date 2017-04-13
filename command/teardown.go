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

package command

import (
	"fmt"
	"os"

	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"github.com/sirupsen/logrus"
	"strings"
)

func CmdTeardown(c *cli.Context) error {
	if err := helpers.RunCommand(helpers.DockerComposeCommand("kill")); err != nil {
		return err
	}

	if err := removeNetworks(); err != nil {
		return err
	}

	if c.Bool("force") {
		logrus.Debugf("found -f flag on teardown, removing all containers")
		if err := helpers.RunCommand(helpers.DockerComposeCommand("down", "--remove-orphans", "-v")); err != nil {
			return err
		}
		return nil
	} else {
		return removeContainersWithoutGcLabel()
	}
}

func removeNetworks() error {
	projectName := os.Getenv("COMPOSE_PROJECT_NAME")

	client := helpers.GetDockerClient()

	networks, err := client.ListNetworks()
	if err != nil {
		return err
	}

	for _, network := range networks {
		if strings.HasPrefix(network.Name, projectName) {
			if err := client.RemoveNetwork(network.ID); err != nil {
				logrus.Errorf("error removing network: %v", err)
			}
		}
	}

	return nil
}

func removeContainersWithoutGcLabel() error {
	// only remove containers that don't have the com.lancope.docker-gc.keep set
	client := helpers.GetDockerClient()
	project := fmt.Sprintf("com.docker.compose.project=%s", os.Getenv("COMPOSE_PROJECT_NAME"))
	queryAll := docker.ListContainersOptions{All: true, Filters: map[string][]string{"label": []string{project}}}
	queryGc := docker.ListContainersOptions{All: true,
		Filters: map[string][]string{"label": []string{project, "com.lancope.docker-gc.keep"}}}

	containers, err := client.ListContainers(queryAll)
	if err != nil {
		logrus.Error("could not query containers to remove", err)
		return err
	}
	logrus.Debugf("found %d container(s) for possible removal", len(containers))

	gcSafeContainers, err := client.ListContainers(queryGc)
	if err != nil {
		logrus.Error("could not query containers to remove", err)
		return err
	}
	logrus.Debugf("found %d container(s) with gc protection", len(gcSafeContainers))

	allIds := getContainerIds(&containers)
	gcSafeIds := getContainerIds(&gcSafeContainers)

	idsToRemove := removeIds(&allIds, &gcSafeIds)
	logrus.Debugf("removing %d containers", len(idsToRemove))
	for _, id := range idsToRemove {
		if err := client.RemoveContainer(docker.RemoveContainerOptions{ID: id, RemoveVolumes: true}); err != nil {
			logrus.Errorf("error removing container with ID: %q, err: %q", id, err)
		}
	}
	return nil
}

func getContainerIds(contaners *[]docker.APIContainers) []string {
	ids := []string{}
	for _, container := range *contaners {
		ids = append(ids, container.ID)
	}
	return ids
}

// remove the items in the 2nd argument from the first
func removeIds(allIds *[]string, idsToRemove *[]string) []string {
	if len(*idsToRemove) <= 0 {
		return *allIds
	}

	ids := []string{}
	for _, id := range *allIds {
		keep := true
		for _, r := range *idsToRemove {
			if id == r {
				keep = false
			}
		}
		if keep {
			ids = append(ids, id)
		}
	}
	return ids
}
