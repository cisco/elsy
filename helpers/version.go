package helpers

import (
	"fmt"
)

var version string
var build string

// BuildVersionString builds a version string containing both the version and the build
func BuildVersionString() string {
	return fmt.Sprintf("%s (build: %s)", version, build)
}

// Version returns the raw version string
func Version() string {
	return version
}
