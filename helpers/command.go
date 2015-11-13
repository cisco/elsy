package helpers

import (
  "os"
  "os/exec"

  "github.com/Sirupsen/logrus"
)

func RunCommand(command *exec.Cmd) error {
  command.Stdout = os.Stdout
  command.Stderr = os.Stderr
  command.Stdin  = os.Stdin
  logrus.Debugf("running command %s with args %v", command.Path, command.Args)
  if err := command.Run(); err != nil {
    logrus.Debug("last command was not successful")
    return err
  }
  return nil
}

func RunCommandWithOutput(command *exec.Cmd) (string, error) {
  logrus.Debugf("running command %s with args %v", command.Path, command.Args)

  out, err := command.Output()

  if err != nil {
    logrus.Debug("last command was not successful")
  }
  
  return string(out[:]), err
}

func ChainCommands(commands []*exec.Cmd) error {
  for _, command := range commands {
    if err := RunCommand(command); err != nil {
      return err
    }
  }
  return nil
}
