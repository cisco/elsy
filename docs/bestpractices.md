# elsy Best Practices and Workflows

This document enumerates some typical development scenarios and describes how to
use elsy to enable those workflows.

## Cloning a new Repo

After cloning a new repo to hack on some code you can use elsy to quickly build, test,
and run the code.

Simply run the following to setup your local environment:

    lc bootstrap

Now you can run the unit tests:

    lc test

After you have verified that pre-package tests are passing, you can package the repo
artifact and execute blackbox tests:

    lc blackbox-test

Assuming all of the above commands pass (i.e., non-0 exit code) you can now run the
artifact (assuming it is something that runs, e.g., a docker image) by using the
following command:

    lc server start -prod

This command will run the `prodserver` service defined in the repo's
`docker-compose.yml` file and output the `ip:port` that the service is listening
on. Note that some repos also support a `devserver` service that you can run by
leaving off the `-prod` flag to the above command; see the `Hot Reload Workflow` section
for more information.

## Starting Fresh in an Existing Repo

There are certain situations when you need to wipe away all containers and start fresh,
for example:

- When switching to a new branch, you want to ensure you are using the containers
that originate from that branch.
- When debugging a particularly strange issue and you want to make sure that stale
state is not a factor.

elsy provides a simple mechanism to enable this, just call `lc teardown` and then
revisit the "Cloning a new Repo" section.

## Viewing docker-compose Commands

elsy is just an opinionated wrapper around `docker-compose` and as such it sometimes
hides how `docker-compose` is being used.

It is always possible to view the exact `docker-compose` commands that elsy is running by
using the `--debug` flag when running your command (e.g., `lc --debug bootstrap`).

## Executing docker-compose Commands

If you want to run plain old `docker-compose` commands in your repo, you should use
elsy to do it by running `lc dc -- COMMAND` (the `--` ensures that elsy will not
attempt to process any arguments that follow it, it is not required in all cases.)

For example, to run a specific service with a custom entrypoint:

    lc dc -- run --entrypoint=/bin/bash prodserver

You may be asking yourself "Why do I need to use elsy to run docker-compose commands?".
The reason is that elsy is doing some things under the hood (e.g., declaring a
specific compose project-name, linking in a parent compose file) and to correctly
interact with the compose-managed containers you need elsy to setup this
wiring for every call.

## Viewing Logs

To view all logs for running containers in the repo simply run:

    lc dc logs

To view the logs for a specific service just qualify the call with the service name:

    lc dc logs prodserver

## Inspecting a Container

Sometimes you need to get a shell into a running container. You can do this by
using `docker exec`, you just need to find the container name.

For example, to shell into the `prodserver` container:

    container_name=$(lc dc ps | grep prodserver | awk '{print $1}')
    docker exec -it "$container_name" /bin/bash

## Inspecting Build Cache Data

If you are using elsy to build an `sbt` or a `maven` repo, chances are you have
found yourself wondering where your `.ivy` or `.m2` data is. The answer is that elsy
leverages a shared data container (shared across all repos) to hold all of that data.

To view the caches you can run the following:

For sbt:

    ## to view sbt data (inside an sbt repo):
    lc dc -- run --entrypoint=/bin/ls sbt /root/.ivy2/cache

    ## to interactively explore sbt data (inside a mvn repo):
    lc dc -- run --entrypoint=/bin/bash sbt

for mvn:

    ## to view mvn data (inside a mvn repo):
    lc dc -- run --entrypoint=/bin/ls mvn /root/.m2/repository

    ## to interactively explore mvn data (inside a mvn repo):
    lc dc -- run --entrypoint=/bin/bash mvn

Note that the above commands only work if the repo is configured to use the
`sbt` or `mvn` template.

If you find yourself missing actual build data (e.g., files in `./target`) you
probably have `LC_ENABLE_SCRATCH_VOLUMES` set to `true` and need to read the
[Improving Performance](improving-performance.md) doc.

## Hot Reload Workflow

The dream workflow for engineers is to:

1. Make a change to source code
2. Visit the browser (or whatever tool you use to interact with the artifact)
and immediately see the effects of your change

elsy helps enable this type of workflow using a `devserver` service in your
`docker-compose.yml`.

If your repo has a `devserver` service, then you can run the following command to
start your hot-reloading server:

    lc server start

This command is very similar to `lc server start -prod`, it just leverages the
implementation provided by `devserver` instead of `prodserver`. It is up to the
repo code to craft a `devserver` service that enables hot reloading (not in scope of this doc).
