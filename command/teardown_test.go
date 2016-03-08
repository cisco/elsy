package command

import (
	"reflect"
	"testing"
)

var removeTestData = []struct {
	ids      []string
	remove   []string
	expected []string
}{
	{[]string{"one", "two", "three", "four"}, []string{"three"}, []string{"one", "two", "four"}},
	{[]string{"one", "two", "three", "four"}, []string{"one", "two", "three", "four"}, []string{}},
	{[]string{"one", "two", "three", "four"}, []string{}, []string{"one", "two", "three", "four"}},
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"one"}, []string{}},
}

func TestRemoveIds(t *testing.T) {

	for _, data := range removeTestData {
		result := removeIds(&data.ids, &data.remove)
		if !reflect.DeepEqual(data.expected, result) {
			t.Errorf("expected result to be: %q but got %q instead", data.expected, result)
		}
	}

}
