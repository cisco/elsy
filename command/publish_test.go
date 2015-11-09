package command

import "testing"

var testData = []struct{
  Input string
  TagName string
  ErrExpected bool
}{
  {"origin/master", "latest", false},
  {"origin/release", "snapshot.release", false},
  {"origin/release/1.0", "1.0", false},
  {"origin/foo", "snapshot.foo", false},
  {"origin/foo/bar", "snapshot.foo.bar", false},
  {"origin", "", true},
  {"release", "", true},
  {"foo", "", true},
}
func TestExtractTagFromBranch(t *testing.T) {
  for _, d := range testData {
    if tagName, err := extractTagFromBranch(d.Input); d.TagName != tagName {
      t.Errorf("expected input: %q to produce tag: %q but got %q instead", d.Input, d.TagName, tagName)
    } else if d.ErrExpected && err == nil  {
      t.Errorf("expected input: %q to produce an error but did not receive an error", d.Input)
    } else if !d.ErrExpected && err != nil  {
      t.Errorf("expected input: %q to not produce an error but received error: %v", d.Input, err)
    }
  }
}
