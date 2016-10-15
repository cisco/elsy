package system

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elsy/template"
)

func CmdViewTemplate(c *cli.Context) error {
	if templateName := c.Args().First(); len(templateName) == 0 {
		return fmt.Errorf("view-template requires an argument of the template to view")
	} else if yaml, err := template.GetV1(templateName, c.GlobalBool("enable-scratch-volumes")); err != nil {
		return err
	} else {
		fmt.Println(yaml)
	}
	return nil
}
