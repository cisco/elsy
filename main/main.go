package main

import (
  "io/ioutil"
  "os"
  "os/exec"
  "path/filepath"
  "regexp"
  "strings"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/template"
  "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"
)

func main() {
  if err := LoadConfigFile("lc.yml"); err != nil {
    panic(err)
  }

  app := cli.NewApp()
  app.Name = Name
  app.Version = Version
  app.Author = "lancope"
  app.Email = ""
  app.Usage = ""

  app.Flags = GlobalFlags()
  app.Commands = Commands()
  app.CommandNotFound = CommandNotFound
  app.Before = beforeHook
  app.After = afterHook
  app.RunAndExitOnError()

  if !CommandSuccess {
    os.Exit(1)
  }
}

func beforeHook(c *cli.Context) error {
  setLogLevel(c)
  preReqCheck(c)
  setComposeBinary(c)
  setComposeProjectName(c)
  setComposeTemplate(c)
  return nil
}

func afterHook(c *cli.Context) error {
  // clean up compose template if it exists
  if file := os.Getenv("LC_BASE_COMPOSE_FILE"); len(file) > 0 {
    logrus.Debugf("attempting to remove base compose file: %v", file)
    if err := os.Remove(file); err != nil {
      return err
    }
  }
  return nil
}

func setLogLevel(c *cli.Context) {
  if c.GlobalBool("debug") {
    logrus.SetLevel(logrus.DebugLevel)
  } else {
    logrus.SetLevel(logrus.InfoLevel)
  }
}

func preReqCheck(c *cli.Context) {
  // TODO: replace this with checking presence and version of local-docker-stack
  if _, err := exec.LookPath("docker"); err != nil {
    logrus.Fatal("could not find docker, please install local-docker-stack")
  }
  dockerCompose := c.GlobalString("docker-compose")
  if _, err := exec.LookPath(dockerCompose); err != nil {
    logrus.Fatalf("could not find docker compose binary: %q, please install local-docker-stack", dockerCompose)
  }

  if versionString, versionComponents, err := helpers.GetDockerComposeVersion(c); err != nil {
    logrus.Warnf("failed checking docker-compose version. Note that lc only supports docker-compose 1.5.0 or higher")
  } else {
    major, minor := versionComponents[0], versionComponents[1]
    // assuming we won't see any docker-compose versions less than 1.x
    if major == 1 && minor < 5 {
      logrus.Fatalf("found docker-compose version %s, lc only supports docker-compose 1.5.0 or higher", versionString)
    }
  }
}

func setComposeBinary(c *cli.Context) {
  os.Setenv("DOCKER_COMPOSE_BINARY", c.GlobalString("docker-compose"))
}

func setComposeProjectName(c *cli.Context) {
  var invalidChars = regexp.MustCompile("[^a-z0-9]")
  projectName := c.GlobalString("project-name")
  if len(projectName) == 0 {
    logrus.Debug("using current working directory for compose project name")
    path, _ := os.Getwd()
    projectName = filepath.Base(path)
  } else {
    logrus.Debugf("using configured value: %q for project name", projectName)
  }
  projectName = invalidChars.ReplaceAllString(strings.ToLower(projectName), "")
  os.Setenv("COMPOSE_PROJECT_NAME", projectName)
}

func setComposeTemplate(c *cli.Context) {
  templateName := c.GlobalString("template")
  if len(templateName) > 0 {
    if yaml, err := template.Get(templateName); err == nil {
      file := createTempComposeFile(yaml)
      logrus.Debugf("setting LC_BASE_COMPOSE_FILE to %v", file)
      os.Setenv("LC_BASE_COMPOSE_FILE", file)
    } else {
      logrus.Panicf("template %q does not exist", templateName)
    }
  }

  dataContainers := template.GetSharedExternalDataContainers(templateName)
  for _, dataContainer := range dataContainers {
    if err := dataContainer.Ensure(); err != nil {
      logrus.Panic("unable to create data container")
    }
  }
}

func createTempComposeFile(yaml string) string {
  cwd, _ := os.Getwd()
  fh, err := ioutil.TempFile(cwd, "lc_docker_compose_template")
  if err != nil {
    logrus.Panic("could not create temporary yaml file")
  }
  defer fh.Close()
  _, err = fh.WriteString(yaml)
  if err != nil {
    logrus.Panic("could not write to temporary yaml file")
  }
  return fh.Name()
}
