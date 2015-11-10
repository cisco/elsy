package main

import (
  "io/ioutil"
  "testing"
)

func TestGetConfigFileMethods(t *testing.T) {
  if fh, err := ioutil.TempFile("", "testlcconfigyaml"); err != nil {
    t.Fatal("could not create temporary file")
  } else {
    defer fh.Close()
    fh.WriteString(`foo: bar`)
    LoadConfigFile(fh.Name())
  }
  // get configured value
  if v := GetConfigFileString("foo"); v != "bar" {
    t.Errorf("expected to get 'bar' for 'foo' but got %q instead", v)
  }
  // get empty for unset values
  if v := GetConfigFileString("bar"); v != "" {
    t.Errorf("expected to get '' for 'bar' but got %q instead", v)
  }
  // when given a default and a configured value, get configured value
  if v := GetConfigFileStringWithDefault("foo", "baz"); v != "bar" {
    t.Errorf("expected to get 'bar' for 'foo' but got %q instead", v)
  }
  // when given a default and no configured value, get default
  if v := GetConfigFileStringWithDefault("bar", "foo"); v != "foo" {
    t.Errorf("expected to get 'foo' for 'bar' but got %q instead", v)
  }
}
