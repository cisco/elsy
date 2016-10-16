package system

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elsy/template"
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
