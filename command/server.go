package command

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "runtime"
  "strings"
  "github.com/codegangsta/cli"
  "github.com/fatih/color"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func CmdServer(c *cli.Context) error {
  if len(c.Args()) == 0 {
    logrus.Info("You must specify one of [start|stop|status|restart] [-prod]")
    return
  }

  cmd := c.Args()[0]
  var serverName = "devserver"

  if c.Bool("prod") {
    serverName = "prodserver"
  }

  if ! helpers.DockerComposeHasService(serverName) {
    return
  }

  if cmd == "start" {
    logrus.Info("Starting ", serverName)
    if err := helpers.RunCommand(dockerComposeCommand(c, "up", "-d", serverName)); err != nil {
      logrus.Fatalf("Unable to start %s: %s", serverName, err)
      return
    }

    showServerAddress(c, serverName, "80")
  } else if cmd == "stop" {
    logrus.Info("Stopping ", serverName)
    if err := helpers.RunCommand(dockerComposeCommand(c, "stop", serverName)); err != nil {
      return
    }
  } else if cmd == "restart" {
    logrus.Info("Restarting ", serverName)
    if err := helpers.RunCommand(dockerComposeCommand(c, "stop", serverName)); err != nil {
      return
    }
    if err := helpers.RunCommand(dockerComposeCommand(c, "up", "-d", serverName)); err != nil {
      return
    }
  } else if cmd == "status" || cmd == "stat" {
    serviceStatus(c, serverName)
  } else {
    logrus.Fatal("The only options are [start|stop|status|restart]")
  }
}

func dockerIp() string {
  var ip = ""

  if dockerHost := os.Getenv("DOCKER_HOST"); dockerHost != "" {
    pattern := regexp.MustCompile("^tcp://([^:]+).*$")
    if matches := pattern.FindStringSubmatch(dockerHost); len(matches) > 0 {
      ip = matches[1]
    } else {
      logrus.Fatal("DOCKER_HOST environment variable is in the wrong format")
    }
  } else if runtime.GOOS == "linux" {
    ip = "127.0.0.1"
  } else {
    logrus.Fatal("You do not have a DOCKER_HOST environment variable set")
  }

  return ip
}

func servicePort(c *cli.Context, serviceName string, containerPort string) string {
  var port = ""

  cmd := dockerComposeCommand(c, "port", serviceName, containerPort)

  out, err := helpers.RunCommandWithOutput(cmd)

  if err != nil {
    logrus.Fatal("Unable to get port", err)
  } else {
    pattern := regexp.MustCompile("^.+:([0-9]+)")

    if matches := pattern.FindStringSubmatch(out); matches != nil {
      port = matches[1]
    } else {
      logrus.Fatal("docker-compose did not return a port")
    }
  }

  return port
}

func showServerAddress(c *cli.Context, serviceName string, containerPort string) {
  ip := dockerIp()
  port := servicePort(c, serviceName, containerPort)

  red := color.New(color.FgRed).SprintFunc()
  green := color.New(color.FgGreen).SprintFunc()

  msg := fmt.Sprintf("%s running at http://%s:%s", serviceName, ip, port)

  logrus.Info(green(msg))
  logrus.Infof("%s %s\n", green("to view the server log, run"), red("./script/logs"))
}

func serviceStatus(c *cli.Context, serviceName string) {
  out, err := helpers.RunCommandWithOutput(dockerComposeCommand(c, "ps", "-q", serviceName))
  if err != nil {
    logrus.Fatal("Unable to get server status: ", err)
  } else {
    containerId := strings.TrimSpace(out)

    if (containerId == "") {
      logrus.Info(serviceName, ": down")
      return
    }

    cmd := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", containerId)
    out, err = helpers.RunCommandWithOutput(cmd)
    if err != nil {
      logrus.Fatal("Unable to get server details: ", err)
    } else {
      if strings.TrimSpace(out) == "true" {
        logrus.Info(serviceName, ": up")
      } else {
        logrus.Info(serviceName, ": down")
      }
    }
  }
}
