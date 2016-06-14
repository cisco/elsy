package helpers

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/fsouza/go-dockerclient"
)

// DockerDebugMsg contains a string to provide to the user for help debugging docker connectivity issues
const DockerDebugMsg = `
Can't connect to 'docker' daemon. The process cannot continue without the daemon running. Possible causes and solutions follow:

- On Mac OS X:
  - Make sure you have lds installed, see: https://stash0.eng.lancope.local/projects/DEV-INFRASTRUCTURE/repos/local-docker-stack/browse
  - Ensure lds environment variables are in your $PATH:
    - Running 'lds' should print something like 'lds only supports commands' followed by a list of commands
    - If running 'lds' prints 'command not found: lds', then you need to run 'source shell/docker.sh' from within the local-docker-stack repo
  - Run 'lds troubleshoot' for help troubleshooting lds
- On Linux, Docker daemon hasn't been started or has crashed

If none of the above works, surface your issue in the #arch room in slack
`

// EnsureDockerConnectivity will return an error if the docker daemon is not accessible
func EnsureDockerConnectivity() error {
	client := GetDockerClient()
	if err := client.Ping(); err != nil {
		return err
	}
	return nil
}

func DockerContainerExists(name string) bool {
	client := GetDockerClient()
	if containers, err := client.ListContainers(docker.ListContainersOptions{All: true}); err != nil {
		logrus.Panic("unable to list docker containers")
	} else {
		for _, container := range containers {
			for _, containerName := range container.Names {
				if containerName == "/"+name {
					logrus.Debugf("found existing container: %s", name)
					return true
				}
			}
		}
	}
	return false
}

// DockerImageExists returns true if docker image:tag exists in local docker daemon
func DockerImageExists(image string, tag string) (bool, error) {
	client := GetDockerClient()
	images, err := client.ListImages(docker.ListImagesOptions{All: true, Filter: image})
	if err != nil {
		return false, err
	}
	expectedRepoTag := fmt.Sprintf("%s:%s", image, tag)
	for _, img := range images {
		for _, tg := range img.RepoTags {
			if tg == expectedRepoTag {
				return true, nil
			}
		}
	}
	return false, nil
}

// RemoveContainersOfImage will remove all containers created from the provided image
func RemoveContainersOfImage(image string) error {
	logrus.Debugf("removing all containers created from image %s", image)
	client := GetDockerClient()
	if containers, err := client.ListContainers(docker.ListContainersOptions{All: true}); err != nil {
		return err
	} else {
		for _, container := range containers {
			// It is unclear when docker uses the human readable image name vs the sha hash
			// in the api responses, so for now just check all locations where docker
			// lists the image value.
			// It seems that client.InspectContainer.Config.Image is the most trustworthy though
			inspection, err := client.InspectContainer(container.ID)
			if err != nil {
				logrus.Debugf("could not inspect container %q", container.ID)
			}
			if container.Image == image || inspection.Image == image || inspection.Config.Image == image {
				logrus.Debugf("removing container: %s", container.ID)
				options := docker.RemoveContainerOptions{ID: container.ID, RemoveVolumes: true, Force: true}
				if err := client.RemoveContainer(options); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// PullDockerImage blindly pulls the image:tag
func PullDockerImage(image string, tag string) error {
	// TODO: support more advanced PullImageOptions
	client := GetDockerClient()
	if err := client.PullImage(docker.PullImageOptions{Repository: image, Tag: tag}, docker.AuthConfiguration{}); err != nil {
		return err
	}
	return nil
}

type DockerDataContainer struct {
	Image     string
	Name      string
	Volumes   []string
	Resilient bool
}

func (ddc *DockerDataContainer) Create() error {
	client := GetDockerClient()
	var volumeMap = make(map[string]struct{})
	for _, volume := range ddc.Volumes {
		volumeMap[volume] = struct{}{}
	}
	labels := make(map[string]string)
	if ddc.Resilient {
		labels["com.lancope.docker-gc.keep"] = ""
	}
	logrus.Debugf("creating data container: %s", ddc.Name)
	_, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: ddc.Name,
		Config: &docker.Config{
			Image:      ddc.Image,
			Labels:     labels,
			Volumes:    volumeMap,
			Entrypoint: []string{"/bin/true"},
		},
	})
	return err
}

func (ddc *DockerDataContainer) Ensure(offline bool) error {
	imageComponents := strings.Split(ddc.Image, ":")
	image, tag := imageComponents[0], ""
	if len(imageComponents) > 1 {
		tag = imageComponents[1]
	}

	if !offline {
		if err := ensureImageExists(image, tag); err != nil {
			logrus.Debugf("error ensuring image exits, attempting to create anyway, err: %q", err)
		}
	}

	if !DockerContainerExists(ddc.Name) {
		return ddc.Create()
	}

	return nil
}

func ensureImageExists(image string, tag string) error {
	logrus.Debugf("checking if image '%s:%s' exists", image, tag)
	if exists, _ := DockerImageExists(image, tag); !exists {
		logrus.Debugf("image ''%s:%s' does not exist locally, pulling", image, tag)
		return PullDockerImage(image, tag)
	}

	return nil
}

var dockerClient *docker.Client

func GetDockerClient() *docker.Client {
	if dockerClient == nil {
		var err error
		logrus.Debug("creating docker client")
		dockerClient, err = docker.NewClientFromEnv()
		if err != nil {
			logrus.Panic("unable to create docker client")
		}
	} else {
	}
	return dockerClient
}

func DockerIp() (string, error) {
	var ip string
	if dockerHost := os.Getenv("DOCKER_HOST"); dockerHost != "" {
		pattern := regexp.MustCompile("^tcp://([^:]+).*$")
		if matches := pattern.FindStringSubmatch(dockerHost); len(matches) > 0 {
			ip = matches[1]
		} else {
			return "", fmt.Errorf("DOCKER_HOST environment variable is in the wrong format")
		}
	} else if runtime.GOOS == "linux" {
		ip = "127.0.0.1"
	} else {
		return "", fmt.Errorf("Unable to determine Docker daemon IP")
	}
	return ip, nil
}

func DockerContainerIsRunning(id string) (bool, error) {
	if status, err := GetDockerClient().InspectContainer(id); err != nil {
		return false, err
	} else {
		return status.State.Running, nil
	}
}

func DockerContainerDyanmicPorts(id string) (map[string]string, error) {
	status, err := GetDockerClient().InspectContainer(id)
	if err != nil {
		return nil, err
	}
	portBindings := make(map[string]string)
	for port, bindings := range status.HostConfig.PortBindings {
		// NB: PortBindings is an array of docker.PortBinding. Currently we only read the first one.
		// not sure when we would have more than 1
		binding := bindings[0]
		if len(binding.HostPort) == 0 {
			portBindings[string(port)] = status.NetworkSettings.Ports[port][0].HostPort
		}
	}
	return portBindings, nil
}
