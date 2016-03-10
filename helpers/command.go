package helpers

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/Sirupsen/logrus"
)

func RunCommand(command *exec.Cmd) error {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
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

type dropFilterFunc func(string) bool

func RunCommandWithFilter(command *exec.Cmd, filter dropFilterFunc) error {
	if pipe, err := command.StdoutPipe(); err != nil {
		return err
	} else {
		go filterPipe(bufio.NewScanner(pipe), filter, os.Stdout)
	}
	if pipe, err := command.StderrPipe(); err != nil {
		return err
	} else {
		go filterPipe(bufio.NewScanner(pipe), filter, os.Stderr)
	}
	logrus.Debugf("running command %s with args %v", command.Path, command.Args)
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}
func filterPipe(scanner *bufio.Scanner, filter dropFilterFunc, dst *os.File) {
	for scanner.Scan() {
		if !filter(scanner.Text()) {
			dst.Write(scanner.Bytes())
			dst.WriteString("\n")
		}
	}
}

func ChainCommands(commands []*exec.Cmd) error {
	for _, command := range commands {
		if err := RunCommand(command); err != nil {
			return err
		}
	}
	return nil
}
