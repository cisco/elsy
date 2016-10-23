package helpers

import (
	"reflect"
	"testing"
)

var testVersionStringData = []struct {
	Version            string
	ExpectedComponents []int
}{
	{"1.9.1", []int{1, 9, 1}},
	{"1.10.3", []int{1, 10, 3}},
	{"1.11.2", []int{1, 11, 2}},
	{"", nil},
}

func TestParseVersionString(t *testing.T) {
	for _, data := range testVersionStringData {
		components, _ := parseVersionString(data.Version)
		if !reflect.DeepEqual(data.ExpectedComponents, components) {
			t.Errorf("expected version to be: %q but got %q instead", data.ExpectedComponents, components)
		}
	}
}
