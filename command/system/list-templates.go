package system

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elsy/template"
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
