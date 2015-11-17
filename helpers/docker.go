package helpers

import (
  "fmt"
  "os"
  "strings"
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
  imageComponents := strings.Split(ddc.Image, ":")
  image, tag := imageComponents[0], ""
  if (len(imageComponents) > 1){
    tag = imageComponents[1]
  }
  if err := ensureImageExists(image, tag); err != nil {
    logrus.Debugf("error ensuring image exits, attempting to create anyway, err: %q", err)
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
