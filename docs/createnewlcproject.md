# Creating A New `lc` Project

`lc` is a application that calls `docker-compose` services. As such, the system uses `docker-compose` to read a file to
know how to run various services corresponding to build lifecycle steps; `test`, `package`, etc.

## `lc.yml`

First, create an `lc.yml` file in the root of your project.  This is the root configuration file.  (See "Configuration",
below.)

## Using Templates

Once you have `lc.yml` set up, you must determine what of the default templates (which constitute a default set of
`docker-compose` services) you can use as-is, or need to change.

The template you use will broadly be named the same as the build tool you are using to build the project.  Here is a
list of currently supported templates.

* `ember`
* `mvn`
* `sbt`

See [Templates](docs/templates.md) for an overview of the templating system, and how to change or modify the templates. 



## Configuration

An `lc` project may be configured via `lc.yml` file at the root of a repo. It supports the following configuration
options:

* `project_name`: name of your `docker-compose` project which is used as __COMPOSE_PROJECT_NAME__.
* `docker_compose`: basename or fully qualified path to the docker-compose binary.
* `template`: [compose template](docs/templates.md) to include.
* `docker_image_name`: name of docker image to build.
* `docker_registry`: address of docker registry to publish to.

Some configuration *per command* may also be specified as command line arguments.  If a command line argument is
present, it will take precedence and override any value in the configuration file.

## `lc` and Command Command Line Arguments

To see a list of commands for `lc`, run `lc` with no arguments.

Example:

```
> lc help
```

The output will contain a list of commands (and other information).

To see what options are available for a given command, run `lc help command` with no arguments.

Example:
```
> lc help system
```

The output will contain a help screen for the command listed ("`system`" in the above example).


