# Lancope Lifecycle

This repo contains the lifecycle (`lc`) binary. It provides a standardized development workflow for common types of
projects. It is primarily a wrapper around docker-compose.

## Core Patterns

`lc` works with repos which only require docker and docker-compose to build, test, and package. It does *not* require
that the primary artifact be a docker image. Customization of lifecycle phases are done by customizing docker-compose
services.

The core `lc` tasks are:

* bootstrap: pulls and builds all services in the docker-compose project
* test: calls the `test` service, forwarding arguments
* package: calls the `package` service. It will also build a docker image if the repo contains a `Dockerfile` at its
  root directory.
* publish: calls the `publish` service. It will also publish the docker image if the repo contains a `Dockerfile` at its
  root directory.
* ci: calls `bootstrap`, `test`, `package`, then `publish`
* release: allows user to create a release of the repo
* teardown: kills and removes all containers for the docker-compose project

`lc` also supports tasks for commonly used project automation tools;

* sbt: calls the `sbt` service, forwarding arguments
* mvn: calls the `mvn` service, forwarding arguments
* npm: calls the `npm` service, forwarding arguments
* bower: calls the `bower` service, forwarding arguments
* ember: calls the `ember` service, forwarding arguments

Sometimes you just need to run `docker-compose` commands using the composite `docker-compose.yml` files that `lc
creates`. `lc` supports this using `lc dc --`, where everything after the
[double dash](http://unix.stackexchange.com/a/11382) are the arguments passed to `docker-compose`. Some examples:

```
# get command help for 'docker-compose ps'
$ lc dc -- ps --help

# get docker-compose version
$ lc dc -- --version

# get an ssh shell into one of your services
$ lc dc -- run --entrypoint=bash package -c bash
```

## Project Templates

`lc` contains docker-compose templates for the most commonly used build tools used in Lancope. A project template
provides a base set of docker-compose services which you may override in a project's `docker-compose.yml` file. The
_overlaying_ of docker-compose services is accomplished by passing multiple `-f` arguments to docker-compose where each
subsequent yaml file may extend services defined in previous yaml files. For example, the `sbt` template provides a
`test` service which calls `sbt test`. If you wanted the test task to also include code coverage, you would add this to
your repo's `docker-compose.yml` file:

```
test:
  command: [coverage test coverageReport]
```

Then run the lifecycle test task like so:

```
lc --template=sbt test
```

## Configuration

See the [Creating a New `lc` Project](docs/createnewlcproject.md) document.

## Improving `lc` performance (experimental)

See the [Improving Performance](docs/improving-performance.md) doc.

## Keeping `lc` up-to-date

Simply run `lc system upgrade` to download the latest `lc` binary.

## Local Development

Use `lc` to develop `lc`!

This repo exposes all of the core `lc` tasks for ongoing development:

```
## bootstrap repo
$ lc bootstrap

## test your code
$ lc test && lc smoketest

## package a new binary, will show up in ./target/
$ lc package
```

### Adding a dependency

To add a dependency, just add a new entry to the `dependencies` array in `./dev-env/dependencies`. After adding the new
dependency you will need to run `lc bootstrap` to import it into the project.

### Setting up CI
In order to have Jenkins build and publish `project-lifecycle`, you must
configure the Jenkins server with credentials to upload to
artifactory1.eng.lancope.local. First, create a local file called `netrc` with
the contents:

```
machine artifactory1.eng.lancope.local login <username> password <password>
```

Visit your Jenkins Credentials management page, add a new credential of type
*Secret File* then upload the `netrc` file you just created.

In the `project-lifecycle` job configuration page, add _Use Secret text for
file_ to the _Build Environment_ section. Select the credential you just added
and specify an env variable of `LC_REPO_NETRC`.

### IDE Integration

**TODO: Figure out if we can share the full $GOPATH from the container**

Follow these instructions to enable IDE integration during development. IDE integration is purely for speeding
local work, developers should still run `lc test && lc smoketest` to validate code before pushing.

[Atom](https://atom.io/) is the recommended editor for `golang` projects and it is also recommended that you use the
[go-plus](https://atom.io/packages/go-plus) package for live `golinting` and `govetting`

To setup your IDE to work with this repo ensure that you have cloned the repo into your `gopath`:

```
$ cd $GOPATH
$ mkdir -p stash0.eng.lancope.local/dev-infrastructure && cd stash0.eng.lancope.local/dev-infrastructure
$ git clone ssh://git@stash0.eng.lancope.local/dev-infrastructure/project-lifecycle.git
```

With the repo cloned, all you need to do is make sure to `go get` all the dependencies:

```
$ cd project-lifecycle
$ ./dev-env/go-gets
```

Now, open atom by running `atom .`
