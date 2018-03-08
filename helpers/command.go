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
	"os"
	"os/exec"
	"syscall"

	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
)

func RunCommand(command *exec.Cmd) error {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	logrus.Debugf("running command %s with args %v", command.Path, command.Args)

	if err := command.Start(); err != nil {
		logrus.Error("unable to start command")
		return err
	}

	if err := command.Wait(); err != nil {
		logrus.Debug("last command was not successful")
		// do the dance to get the exitcode
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return cli.NewExitError(err.Error(), status.ExitStatus())
			} else {
				logrus.Error("unable to cast to WaitStatus")
				return err
			}
		} else {
			logrus.Error("unable to cast to ExitError")
			return err
		}
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
