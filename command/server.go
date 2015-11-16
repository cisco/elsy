package command

import (
  "errors"
  "fmt"
  "regexp"

  "github.com/codegangsta/cli"
  "github.com/fatih/color"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

const usageMessage = "You must specify one of [start|stop|status|restart|logs] [-prod]"

func CmdServer(c *cli.Context) error {
  if len(c.Args()) == 0 {
    logrus.Info(usageMessage)
    return errors.New(usageMessage)
  }

  cmd := c.Args()[0]
  var serverName = "devserver"

  if c.Bool("prod") {
    serverName = "prodserver"
  }

  if ! helpers.DockerComposeHasService(serverName) {
    return fmt.Errorf("no %q service defined", serverName)
  }

  if cmd == "start" {
    logrus.Info("Starting ", serverName)
    if err := helpers.RunCommand(helpers.DockerComposeCommand("up", "-d", serverName)); err != nil {
      logrus.Fatalf("Unable to start %s: %s", serverName, err)
      return err
    }
    showServerAddress(c, serverName, "80")
  } else if cmd == "stop" {
    logrus.Info("Stopping ", serverName)
    if err := helpers.RunCommand(helpers.DockerComposeCommand("stop", serverName)); err != nil {
      return err
    }
  } else if cmd == "restart" {
    logrus.Info("Restarting ", serverName)
    if err := helpers.RunCommand(helpers.DockerComposeCommand("stop", serverName)); err != nil {
      return err
    }
    if err := helpers.RunCommand(helpers.DockerComposeCommand("up", "-d", serverName)); err != nil {
      return err
    }
  } else if cmd == "status" || cmd == "stat" {
    if status, err := helpers.DockerComposeServiceIsRunning(serverName); err != nil {
      return err
    } else if status {
      logrus.Infof("%s: up", serverName)
    } else {
      logrus.Infof("%s: down", serverName)
    }
  } else if cmd == "logs" || cmd == "log" {
    logrus.Info("Press Ctrl-C to stop...")
    if err := helpers.RunCommand(helpers.DockerComposeCommand("logs", serverName)); err != nil {
      logrus.Fatalf("Unable to get logs for %s: %s", serverName, err)
      return err
    }
  } else {
    logrus.Fatal(usageMessage)
    return errors.New(usageMessage)
  }

  return nil
}

func servicePort(c *cli.Context, serviceName string, containerPort string) string {
  var port = ""

  cmd := helpers.DockerComposeCommand("port", serviceName, containerPort)

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
  ip, err := helpers.DockerIp()
  if err != nil {
    logrus.Fatal(err)
    return
  }
  port := servicePort(c, serviceName, containerPort)

  red := color.New(color.FgRed).SprintFunc()
  green := color.New(color.FgGreen).SprintFunc()

  msg := fmt.Sprintf("%s running at http://%s:%s", serviceName, ip, port)
  var logsMsg = "lc server logs"

  if serviceName == "prodserver" {
    logsMsg = "lc server logs --prod"
  }

  logrus.Info(green(msg))
  logrus.Infof("%s %s\n", green("to view the server log, run"), red(logsMsg))
}
