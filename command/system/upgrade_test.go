package system

import (
  "testing"

  "fmt"
  "io/ioutil"
  "os"
)


func check(t *testing.T, e error) {
  if e != nil {
    t.Error(e)
  }
}

func TestCompareFiles(t *testing.T) {
  noDiff1 := []string{
    "samefile",
  }

  noDiff2 := []string{
    "samefile",
    "samefile",
  }

  noDiff3 := []string{
    "samefile",
    "samefile",
    "samefile",
  }

  noDiff4 := []string{
    "samefile",
    "samefile",
    "samefile",
  }

  diff1 := []string{
    "samefile",
    "differentFile",
  }

  diff2 := []string{
    "samefile",
    "samefile",
    "differentFile",
  }

  diff3 := []string{
    "samefile",
    "samefile",
    "differentFile",
    "samefile",
  }

  doCompare(t, noDiff1, false)
  doCompare(t, noDiff2, false)
  doCompare(t, noDiff3, false)
  doCompare(t, noDiff4, false)
  doCompare(t, diff1, true)
  doCompare(t, diff2, true)
  doCompare(t, diff3, true)
}

func doCompare(t *testing.T, testCase []string, expected bool){
  tmpDir, err := ioutil.TempDir("", "upgradeTestCompareFiles")
  check(t, err)
  defer os.RemoveAll(tmpDir)

  files := make([]string, 0, len(testCase))
  for i, content := range testCase {
    file := fmt.Sprintf("%s/%s%d", tmpDir, "testfile", i)
    d := []byte(content)
    err = ioutil.WriteFile(file, d, 0644)
    check(t, err)
    files = append(files, file)
  }
  result := detectDifferences(files...)
  if result != expected {
    t.Errorf("expected files: %s to produce result: %t but got %t instead", testCase, expected, result)
  }
}
