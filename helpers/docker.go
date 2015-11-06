package helpers

import (
  "github.com/fsouza/go-dockerclient"
)

func DockerContainerExists(name string) bool {
  client := GetDockerClient()
  if containers, err := client.ListContainers(docker.ListContainersOptions{All: true}); err != nil {
    panic("unable to list docker containers")
  } else {
    for _, container := range containers {
      for _, containerName := range container.Names {
        if containerName == "/"+name {
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
    dockerClient, err = docker.NewClientFromEnv()
    if err != nil {
      panic("unable to create docker client")
    }
  } else {
  }
  return dockerClient
}
