package main

import (
  "fmt"
)

const name string = "lc"
var version string
var build string

func getVersion() string {
  return fmt.Sprintf("%s (build: %s)", version, build)
}
