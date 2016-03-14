package command

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// test that when setupDir is called with no args, it returns working directory
func TestSetupDirCurrentDirectory(t *testing.T) {
	// test current directory
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	lcDir, err := setupDir([]string{})
	if err != nil {
		t.Fatal(err)
	}
	if wd != lcDir {
		t.Errorf("expected result to be: %q but got %q instead", wd, lcDir)
	}
}

// test that when setupDir is called with a specific, pre-existing, directory,
// it returns the absolute path to that directory
func TestSetupDirDifferentDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "lc-init_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	lcDir, err := setupDir([]string{dir})
	if err != nil {
		t.Fatal(err)
	}
	if dir != lcDir {
		t.Errorf("expected result to be: %q but got %q instead", dir, lcDir)
	}
}

// test that setupDir will create the directory if it doesn't exist
func TestSetupDirNewDirectory(t *testing.T) {
	dir := filepath.Join(os.TempDir(), "lc-init_test", "manualtest")
	defer os.RemoveAll(dir)

	if _, err := os.Stat(dir); err == nil {
		t.Errorf("dir %q existed before starting test, cannot proceed", dir)
	}

	lcDir, err := setupDir([]string{dir})
	if err != nil {
		t.Fatal(err)
	}
	if dir != lcDir {
		t.Errorf("expected result to be: %q but got %q instead", dir, lcDir)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("expected %q to exist, but it did not", dir)
	}
}
