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

package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
)

// CmdCi runs ci loop
func CmdCi(c *cli.Context) error {
	defer cleanup(c)

	logrus.Info("Running bootstrap")
	if err := CmdBootstrap(c); err != nil {
		return err
	}

	if helpers.DockerComposeHasService("test") {
		logrus.Info("Running test")
		if err := CmdTest(c); err != nil {
			return err
		}
	}

	logrus.Info("Running package")
	if err := RunPackage(c); err != nil {
		return err
	}

	if helpers.DockerComposeHasService("blackbox-test") {
		logrus.Infof("Running blackbox-test")
		if err := RunBlackboxTest(c); err != nil {
			return err
		}
	}

	if helpers.DockerComposeHasService("publish") || helpers.HasDockerfile() {
		logrus.Info("Running publish")
		return CmdPublish(c)
	} else {
		logrus.Info("No publish service defined, and no Dockerfile present.")
		return nil
	}
}

func cleanup(c *cli.Context) {
	defer CmdTeardown(c)

	//write build logs
	logDir := c.String("build-logs-dir")
	if len(logDir) > 0 {
		logrus.WithFields(logrus.Fields{"build-logs-dir": logDir}).Info("writing build logs to build-logs-dir")

		absDir, err := ensureLogDir(logDir)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"build-logs-dir": logDir,
				"error":          err,
			}).Error("failure creating build-logs-dir, skipping writing logs")
			return
		}

		// NOTE: we may want to filter these service logs and not write them iff
		// the log just contains "Attaching to...", but that might be confusing to
		// the user as well...so keeping it as simple as possible for now, just dump it all.
		writeServiceLogs(absDir, helpers.DockerComposeServices())
	}
}

func ensureLogDir(logDir string) (string, error) {
	absDir, err := filepath.Abs(logDir)
	if err != nil {
		return "", fmt.Errorf("error taking abs path of build-logs-dir %v", err)
	}

	logrus.WithFields(logrus.Fields{"build-logs-dir": logDir}).Debug("wiping old logDir")
	if err := os.RemoveAll(absDir); err != nil {
		return "", fmt.Errorf("error removing old build-logs-dir %v", err)
	}

	logrus.WithFields(logrus.Fields{"build-logs-dir": logDir}).Debug("creating build-logs-dir")
	if err := os.MkdirAll(absDir, os.FileMode(int(0766))); err != nil {
		return "", fmt.Errorf("error creating build-logs-dir %v", err)
	}

	return absDir, nil
}

// writeServiceLogs assumes dir is the absolute path to write log files and
// that it exists and is writable by current process
func writeServiceLogs(dir string, services []string) {
	var wg sync.WaitGroup
	for _, s := range services {
		wg.Add(1)
		go func(service string) {
			defer wg.Done()
			p := filepath.Join(dir, service)

			logs, err := helpers.ServiceLogs(service)
			if err != nil {
				logrus.WithFields(logrus.Fields{"service": service, "error": err}).Error("could not get logs for service")
			}

			logrus.WithFields(logrus.Fields{"service": service, "file": p}).Debug("writing logs for service")

			if err := ioutil.WriteFile(p, logs, os.FileMode(int(0766))); err != nil {
				logrus.WithFields(logrus.Fields{"service": service, "error": err}).Error("could not write logs for service")
			}
		}(s)
	}

	wg.Wait()
}
