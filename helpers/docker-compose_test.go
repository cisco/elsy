package helpers

import (
	"reflect"
	"testing"
)

var testComposeVersionData = []struct {
	Version            string
	ExpectedString     string
	ExpectedComponents []int
}{
	{"docker-compose 1.2.0", "1.2.0", []int{1, 2, 0}},
	{"docker-compose version: 1.3.3\nCPython version: 2.7.9\nOpenSSL version: OpenSSL 1.0.1j 15 Oct 2014", "1.3.3", []int{1, 3, 3}},
	{"docker-compose 1.4.1", "1.4.1", []int{1, 4, 1}},
	{"docker-compose version: 1.5.0", "1.5.0", []int{1, 5, 0}},
	{"docker-compose version 1.5.2, build 7240ff3", "1.5.2", []int{1, 5, 2}},
	{"docker-compose version: 1.6.0", "1.6.0", []int{1, 6, 0}},
}

func TestParseDockerCompseVersion(t *testing.T) {
	for _, data := range testComposeVersionData {
		version, components, _ := parseDockerComposeVersion(data.Version)
		if data.ExpectedString != version {
			t.Errorf("expected version to be: %q but got %q instead", data.ExpectedString, version)
		}
		if !reflect.DeepEqual(data.ExpectedComponents, components) {
			t.Errorf("expected version to be: %q but got %q instead", data.ExpectedComponents, components)
		}
	}
}
