package system

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/elsy/template"
)

// CmdListTemplates displays an alphabetical list of all available templates
func CmdListTemplates(c *cli.Context) error {
	fmt.Println("Run `lc system view-template <template-name>` to see the template contents.")
	for _, name := range template.List() {
		fmt.Println(name)
	}

	return nil
}
