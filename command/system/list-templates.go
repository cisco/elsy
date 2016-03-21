package system

import (
	"fmt"
	"github.com/codegangsta/cli"
	"stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/template"
)

// CmdListTemplates displays an alphabetical list of all available templates
func CmdListTemplates(c *cli.Context) error {
	fmt.Println("Run `lc system view-template <template-name>` to see the template contents.")
	for _, name := range template.List() {
		fmt.Println(name)
	}

	return nil
}
