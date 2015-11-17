package helpers

import (
  "fmt"
  "os"
  "regexp"
  "runtime"

  "github.com/fsouza/go-dockerclient"
  "github.com/Sirupsen/logrus"
)

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

type DockerDataContainer struct {
  Image string
  Name string
  Volumes []string
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
      Image: ddc.Image,
      Labels: labels,
      Volumes: volumeMap,
      Entrypoint: []string{"/bin/true"},
    },
  })
  return err
}

func (ddc *DockerDataContainer) Ensure() error {
  if !DockerContainerExists(ddc.Name) {
    return ddc.Create()
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
  status, err := GetDockerClient().InspectContainer(id);
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
