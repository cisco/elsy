package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// markerFiles holds files that mark the repo as "already initialized"
var markerFiles = []string{"lc.yml"}

// CmdInit will setup a blank repo to work with lc.
func CmdInit(c *cli.Context) error {
	dir, err := setupDir(c.Args())
	if err != nil {
		return err
	}

	if err := failIfPresent(dir, markerFiles...); err != nil {
		return fmt.Errorf("repo already initialized, err: %q", err)
	}

	// determine project-name
	projectName := c.String("project-name")
	if len(projectName) == 0 {
		logrus.Debugf("--project-name not used, creating project-name from directory")
		p, err := createProjectName(dir)
		if err != nil {
			return err
		}
		projectName = p
	}

	logrus.Infof("Initializing lc project at %s using project-name: %s", dir, projectName)

	if err := initLcFile(dir, projectName, c.String("template"), c.String("docker-image-name"), c.String("docker-registry")); err != nil {
		return fmt.Errorf("failure writing lc.yml: %q", err)
	}

	if err := initDockerComposeFile(dir); err != nil {
		return fmt.Errorf("failure writing docker-compose.yml: %q", err)
	}

	if len(c.String("docker-image-name")) > 0 {
		if err := initDockerFile(dir); err != nil {
			return fmt.Errorf("failure writing Dockerfile: %q", err)
		}
	}

	return nil
}

// initLcFile creates the lc file
func initLcFile(dir, projectName, template, dockerImageName string, dockerRegistry string) error {
	lc := filepath.Join(dir, "lc.yml")
	mode := os.FileMode(int(0755))

	contents := fmt.Sprintf("project_name: %s", projectName)

	if len(template) > 0 {
		contents = contents + fmt.Sprintf("\ntemplate: %s", template)
	}

	// TODO: should we write the imageName if registry is not specified, and vice versa?
	if len(dockerImageName) > 0 {
		contents = contents + fmt.Sprintf("\ndocker_image_name: %s", dockerImageName)
	}

	if len(dockerRegistry) > 0 {
		contents = contents + fmt.Sprintf("\ndocker_registry: %s", dockerRegistry)
	}

	logrus.Debugf("Creating lc.yml")
	if err := ioutil.WriteFile(lc, []byte(contents), mode); err != nil {
		return err
	}
	return nil
}

// initDockerComposeFile will create the docker-compose file if not already present
func initDockerComposeFile(dir string) error {
	dockerCompose := filepath.Join(dir, "docker-compose.yml")

	if _, err := os.Stat(dockerCompose); err == nil {
		logrus.Debugf("docker-compose.yml already exists, not modifying it")
		return nil
	}

	logrus.Debugf("Creating docker-compose.yml")
	contents := fmt.Sprintf("noop:\n   image: alpine")
	mode := os.FileMode(int(0755))
	if err := ioutil.WriteFile(dockerCompose, []byte(contents), mode); err != nil {
		return err
	}
	return nil
}

// initDockerFile will create the Dockerfile if not already present
func initDockerFile(dir string) error {
	docker := filepath.Join(dir, "Dockerfile")

	if _, err := os.Stat(docker); err == nil {
		logrus.Debugf("Dockerfile already exists, not modifying it")
		return nil
	}

	logrus.Debugf("Creating Dockerfile")
	contents := fmt.Sprintf("FROM scratch")
	mode := os.FileMode(int(0755))
	if err := ioutil.WriteFile(docker, []byte(contents), mode); err != nil {
		return err
	}
	return nil
}

// setupDir will return the string containing the absolute directory path
// to use for the lc repo; this function gaurentees that the directory exists and
// contains a non-initialized lc repo.
func setupDir(commandArgs []string) (string, error) {
	var dir string
	if len(commandArgs) > 0 {
		d, err := createDir(commandArgs[0])
		if err != nil {
			return "", fmt.Errorf("could not create directory at: %q, err: %q", commandArgs[0], err)
		}
		dir = d
	} else {
		logrus.Debugf("no directory specified, initializing using current directory")
		d, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("could not find current working directory: %q", err)
		}
		dir = d
	}
	return dir, nil
}

// create the given directory if it doesn't already exist
// returns the absolute path to that directory
func createDir(dir string) (string, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	logrus.Debugf("creating directory: %q", absDir)
	if err := os.MkdirAll(absDir, os.FileMode(int(0755))); err != nil {
		return "", err
	}
	return absDir, nil
}

// failIfPresent will return an error if any of the given files are found in
// the current directory
func failIfPresent(cwd string, files ...string) error {
	for _, file := range files {
		if _, err := os.Stat(filepath.Join(cwd, file)); err == nil {
			return fmt.Errorf("found file: %q in current directory", file)
		}
	}
	return nil
}

// createProjectName will create an lc project name based off the cwd
func createProjectName(cwd string) (string, error) {
	parts := strings.Split(strings.TrimSuffix(cwd, "/"), "/")
	dir := parts[len(parts)-1]

	logrus.Debugf("stripping non-alphanumeric characters from directory: %q", dir)
	cleanDir := strings.Map(func(r rune) rune {
		// TODO: verify this matches the docker-compose rules for a valid project-name
		if unicode.IsSpace(r) || !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			return -1
		}
		return r
	}, dir)

	if len(cleanDir) == 0 {
		return "", fmt.Errorf("could not create project name from cwd: %q, current dir blank after removing special characters", cwd)
	}

	return strings.ToLower(cleanDir), nil
}
