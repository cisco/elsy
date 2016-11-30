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
	"os/exec"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/cisco/elsy/helpers"
)

// CmdRelease will create, and push a release tag
func CmdRelease(c *cli.Context) error {
	version := c.String("version")
	commit := c.String("git-commit")
	if len(version) == 0 {
		return fmt.Errorf("--version flag required")
	}
	if len(commit) == 0 {
		return fmt.Errorf("--git-commit flag required")
	}
	if err := helpers.CheckTag(version); err != nil {
		return err
	}

	// TODO: we might want to allow a '-f' option to support re-running a tag build
	// since if a user pushes a tag and the build fails, it is not simple to rerun that build without
	// repushing the tag
	logrus.Infof("creating, and pushing, git tag %s at commit %s", version, commit)
	return helpers.ChainCommands([]*exec.Cmd{
		exec.Command("git", "tag", "-a", version, commit, "-m", fmt.Sprintf("add release tag for %s", version)),
		exec.Command("git", "push", "origin", version),
	})
}
