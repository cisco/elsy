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

	"github.com/codegangsta/cli"
	"github.com/cisco/elsy/template"
)

// CmdViewTemplate prints out a given template (all versions)
func CmdViewTemplate(c *cli.Context) error {
	templateName := c.Args().First()
	if len(templateName) == 0 {
		return fmt.Errorf("view-template requires an argument of the template to view")
	}

	yamlV1, err := template.GetV1(templateName, c.GlobalBool("enable-scratch-volumes"))
	if err != nil {
		return err
	}
	fmt.Println("Compose V1 Version:")
	fmt.Println(yamlV1)

	yamlV2, err := template.GetV2(templateName, c.GlobalBool("enable-scratch-volumes"))
	if err != nil {
		return err
	}

	fmt.Println("Compose V2 Version:")
	fmt.Println(yamlV2)
	return nil
}
