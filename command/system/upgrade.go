package system

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// CmdUpgrade will upgrade the current lc binary
// This command is now deprecated in favor of upgrading
// through LDS.
func CmdUpgrade(c *cli.Context) error {
	logrus.Infof("lc no longer updates itself. Run `lds upgrade`, to upgrade it.")
	return nil
}
