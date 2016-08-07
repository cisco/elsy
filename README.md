# Cisco Lifecycle (lc)

lc is an opinionated, multi-language, build-tool based on
[Docker](https://github.com/docker/docker) and [Docker
Compose](https://github.com/docker/compose). It is allows organizations to
implement a consistent build workflow across many different repos, spanning a
wide array of programming languages

lc will not replace your favorite build tool, it is simply a thin wrapper on that
tool that:

- provides a consistent development workflow across all repos using lc
- provides the ability to fully test Docker images from a blackbox perspective
- reduces local-dev tool requirements to the bare minimum, regardless of programming
language (i.e., you only need to install Docker, Compose, and lc)
- ensures consistent builds regardless of environment (i.e., fixes the "works on
my machine" problem since the repo defines its exact dependency requirements)

With lc, it is possible to build, test, and publish a repo from scratch with just:

```
git clone <repo>
cd repo
lc ci
```

## Getting Started

Follow the below steps to install lc:

```
## install binary for your system and make it executable, for mac:
$ wget -O /usr/local/bin/lc https://artifactory1.lancope.ciscolabs.com/generic-dev-infrastructure/lc/lc-darwin-amd64-1.0.0
$ chmod +x /usr/local/bin/lc

## Tell lc where to find docker-compose and the docker daemon socket
export LC_DOCKER_COMPOSE=/usr/local/bin/docker-compose
export DOCKER_HOST=tcp://docker-daemon-host:2375

```

See the [Using LC in a Project](docs/configuringlcrepo.md) document for
info on how to setup a repo to use lc.

## The Lifecycle
At its core, lc is an implementation of a build lifecycle that generalizes to
any repo producing a software artifact. Running `lc ci` will execute the
full build lifecycle inside a repo, it is made up of the stages defined in the
following sub-sections. `lc ci` operates in a `fail-fast` mode, so if any stage
fails, the following stage will not be run.

See the [examples folder](./examples/README.md) for concrete examples of this
lifecycle in action.

See [LC Best Practices](docs/bestpractices.md) for guidance on how to use `lc`
for typical development workflows.

### lc teardown
Running `lc teardown` simply tells lc to clean up any state that might be left
over from a previous build.

### lc bootstrap
Running `lc bootstrap` will setup a new repo and make sure all dependencies
(e.g., docker images, external software libs) are downloaded and built. Thanks
to Docker caching, this step is only time-intensive the first time it is run.

If present, bootstrap will call the repo's docker-compose `installdependencies`
service that will execute repo-specific command(s) to install external
libraries. See the `docker-compose.yml` file inside the lc repo itself for an
example of this.

### lc test
Running `lc test` will execute the repo's docker-compose `test` service that will
execute repo-specific command(s) to run all unit and integration tests for the
code in that repo.

### lc package

Running `lc package` will do two things. First it will execute the repo's (optional)
docker-compose `package` service that will execute repo-specific command(s) for packaging
the repo's code into the final artifact. Second, if a `Dockerfile` is found in
the root of the repo, lc will build that `Dockerfile` into a new Docker image that
is ready for final testing and publishing.

Note, when run on its own, `lc package` will also run `lc test` to
ensure you are packaging working code, you can prevent this by using the
`--skip-tests` flag.

### lc blackbox-test
This is where the real power of docker-based development comes into play.

Running `lc blackbox-test` will execute the repo's docker-compose
`blackbox-test` service to run repo-specifc logic for testing the final
artifact of the repo. This means that it is possible to test the real container
before releasing it to production.

For example, if the repo is producing a Docker-based microservice that uses a Mysql
database, the `blackbox-teset` service will:

1. stand up the microservice container that was just packaged during `lc package`
1. stand up a mysql container (and initialize the schema) for the microservice to use
1. execute API-level tests against the microservice container to ensure it is
functioning correctly with the database

Note, when run on its own, `lc blackbox-test` will also run `lc package` to
ensure you  are testing the latest code, you can prevent this by using the
`--skip-package` flag.

### lc publish
Running `lc publish` is similar to `lc package`, in that it does two things:
First it will execute the repo's (optional) docker-compose `publish` service
that will run repo-specific command(s) for publishing an artifact. This custom
service is typically used for repos that do not produce Docker images.

Second, if a Docker image was created during the `lc package` phase, lc will
correctly tag and publish that image to the registry defined in the `lc.yml`
file.

`lc publish` uses the following rules when deciding what to publish:

**For running the custom publish service:**

- lc will only run the custom publish service on branches with the pattern of:
`origin/master` or `origin/release/<name>`, or on a valid lc release git tag.

**For tagging Docker images:**

- If the git branch is `origin/master`, lc will apply the tag `latest` to the
docker image.
- If the git branch is `origin/release/<name>` lc will apply the tag `<name>`
- If the git branch is `origin/feature/<name>`, lc will apply the tag
`snapshot.feature.<name>` to the docker image.
- If the git branch is `origin/<name>`, lc will apply the tag `snapshot.<name>`
- If a git tag exists and it is a valid lc release tag, lc will use that tag as
the docker image tag.

**Valid Git Relase Tag:**

lc currently considers a valid git release tag to be any tag following the
schema:

`vX.Y.Z[-Q]`

Where `X`, `Y` and `Z` are integers representing the Major, Minor, and Patch
version (respectively) and `Q` is an optional string qualifier. In the future we
plan to make this schema configurable.

## LC Templates

The lc lifecycle manifests itself in subtly different ways depending on the
underlying build tool. lc ships with a small set of pre-baked templates (e.g.,
mvn, sbt) that define a sensible default lifecycle for the build tool
encapsulated by the template.

See the [lc templates](./docs/templates.md) documentation for more information
on using templates.

## Improving lc Performance

See the [Improving Performance](docs/improving-performance.md) doc.

## Contributing

See the [Contributing to LC](docs/contributing.md) document.
