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

package system

import (
	"fmt"

	"github.com/cisco/elsy/template"
	"github.com/codegangsta/cli"
)

// CmdListTemplates displays an alphabetical list of all available templates
func CmdListTemplates(c *cli.Context) error {
	fmt.Println("Run `lc system view-template <template-name>` to see the template contents.")

	fmt.Println()
	fmt.Println("Compose v1 Templates:")
	for _, name := range template.ListV1() {
		fmt.Println(name)
	}

	fmt.Println()
	fmt.Println("Compose v2 Templates:")
	for _, name := range template.ListV2() {
		fmt.Println(name)
	}
	return nil
}
