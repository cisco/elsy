package helpers

import (
  "os"
  "os/exec"

  "github.com/Sirupsen/logrus"
)

var LastCommandSuccess = true

func RunCommand(command *exec.Cmd) error {
  command.Stdout = os.Stdout
  command.Stderr = os.Stderr
  command.Stdin  = os.Stdin
  logrus.Debugf("running command %s with args %v", command.Path, command.Args)
  if err := command.Run(); err != nil {
    logrus.Debug("last command was not successful")
    LastCommandSuccess = false
    return err
  }
  LastCommandSuccess = true
  return nil
}

func ChainCommands(commands []*exec.Cmd) error {
  for _, command := range commands {
    if err := RunCommand(command); err != nil {
      return err
    }
  }
  return nil
}
