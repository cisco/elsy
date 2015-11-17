package server

import (
  "github.com/codegangsta/cli"
)

func CmdRestart(c *cli.Context) error {
  if runningServer, err := runningServer(); err != nil {
    return err
  } else if err := CmdStop(c); err != nil {
    return err
  } else {
    var serviceToStart string
    if runningServer == "prodserver" {
      serviceToStart = "prodserver"
    } else {
      serviceToStart = "devserver"
    }
    if err := ensureServiceStarted(serviceToStart); err != nil {
      return err
    }
    return CmdStatus(c)
  }
}
