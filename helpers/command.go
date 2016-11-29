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
