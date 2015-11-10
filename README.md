# Lancope Lifecycle

This repo contains the lifecycle (`lc`) binary. It provides a standardized development workflow for common types of projects. It is primarily a wrapper around docker-compose.

## Core Patterns

`lc` works with repos which only require docker and docker-compose to build, test, and package. It does *not* require that the primary artifact be a docker image. Customization of lifecycle phases are done by customizing docker-compose services.

The core `lc` tasks are:

* bootstrap: pulls and builds all services in the docker-compose project
* test: calls the `test` service, forwarding arguments
* package: calls the `package` service. It will also build a docker image if the repo contains a `Dockerfile` at its root directory.
* publish: calls the `publish` service. It will also publish the docker image if the repo contains a `Dockerfile` at its root directory.
* ci: calls `bootstrap`, `test`, `package`, then `publish`
* teardown: kills and removes all containers for the docker-compose project

`lc` also supports tasks for commonly used project automation tools;

* sbt: calls the `sbt` service, forwarding arguments
* mvn: calls the `mvn` service, forwarding arguments
* npm: calls the `npm` service, forwarding arguments
* bower: calls the `bower` service, forwarding arguments
* ember: calls the `ember` service, forwarding arguments

## Project Templates

`lc` contains docker-compose templates for the most commonly used build tools used in Lancope. A project template provides a base set of docker-compose services which you may override in a project's `docker-compose.yml` file. The _overlaying_ of docker-compose services is accomplished by passing multiple `-f` arguments to docker-compose where each subsequent yaml file may extend services defined in previous yaml files. For example, the `sbt` template provides a `test` service which calls `sbt test`. If you wanted the test task to also include code coverage, you would add this to your repo's `docker-compose.yml` file:

```
test:
  command: [coverage test coverageReport]
```

Then run the lifecycle test task like so:

```
lc --template=sbt test
```

## Configuration

`lc` may be configured via `.lc.yaml` file at the root of a repo. It supports the following configuration options:

```
project_name: name of your docker-compose project which is used as COMPOSE_PROJECT_NAME
docker_compose: basename or fully qualified path to the docker-compose binary.
template: compose template to include
docker_image_name: name of docker image to build
docker_registry: address of docker registry to publish to
```

Configuration may also be specified as command line arguments in which case they take precedence over values in the configuration file.

## Keeping `lc` up-to-date

TODO: figure this out
