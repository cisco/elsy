# Improving elsy performance

NOTE: the improvements described in this section only work if you are running
with and elsy template that supports SCRATCH_VOLUMES (currently the mvn, sbt,
and lein templates support scratch volumes).

If you are using elsy on a non-linux host, then there is a lot of I/O that
happens between your repo folders on your host OS and the corresponding folder
on the VM where your docker-daemon runs; this I/O may sometimes slow down local
builds.

To reduce this I/O it is possible to set the following in your environment:

```
export LC_ENABLE_SCRATCH_VOLUMES=true
```

Exporting the above environment variable will tell elsy to keep all temporary
build resources (i.e., the data most often found in `./target` subdirectories
for jvm builds) inside the docker-daemon VM; this is done using data containers.
To see how this affects your favorite template try the following:

```
# print the template without scratch volumes enabled:
$ unset LC_ENABLE_SCRATCH_VOLUMES
$ lc system view-template mvn

# print the template with scratch volumes
$ lc --enable-scratch-volumes system view-template mvn
```

Note in the last line in the above example, using the `--enable-scratch-volumes`
flag is equivalent to setting the `LC_ENABLE_SCRATCH_VOLUMES` environment
variable.

The following subsections provide more details on using this setting:

## Inspecting build data inside the data container

When using `LC_ENABLE_SCRATCH_VOLUMES`, it is still possible to view all build
data. You simply need to inspect the data container.

To inspect the data container, run the following to spawn a container that
mounts the volumes from the data container:

```
## For repos using the sbt template:
$ lc dc -- run --entrypoint=bash sbt -c bash

## For repos using the maven template:
$ lc dc -- run --entrypoint=bash mvn -c bash
```

After running the above you will have a shell in with a working directory of
`/opt/project`, which will contain all of your repo data, including the data
inside the data container.

## Inspecting build tool data

When using `LC_ENABLE_SCRATCH_VOLUMES`, it is still possible to view the
artifacts downloaded by your build tool.

For `sbt`, (assuming the `sbt` template), you can run the following command:

```
lc --enable-scratch-volumes dc -- run --entrypoint=sh sbt
# optionally without --enable-scratch-volume if you have the env var set
```

then `cd` and/or `ls` to `/root/.ivy2/cache` to see the things `sbt` downloaded.  

What this command is doing is running a `sh` process in the `sbt` image from
`docker-compose`, which _links_ to the shared `sbt` download data container.
`entrypoint` is specified to override the given `entrypoint` in the
`docker-compose` file.

For `mvn`, (assuming the `mvn` template), the procedure is the same, but with a
different directory (since it's a different build tool).

```
lc --enable-scratch-volumes dc -- run --entrypoint=sh mvn
# optionally without --enable-scratch-volume if you have the env var set
```

then `cd` to `/root/.m2/repository` and peruse at your liesure.


## Technical Details

This section explains the technical mechanisms underpinning
`LC_ENABLE_SCRATCH_VOLUMES`. Feel free to ignore this.

Lets use the `mvn` template, with `--enable-scratch-volumes` to explain how
things are working:

```
$ lc --enable-scratch-volumes system view-template mvn
mvnscratch:
  image: busybox
  volumes:
  - /opt/project/target/classes
  - /opt/project/target/journal
  - /opt/project/target/maven-archiver
  - /opt/project/target/maven-status
  - /opt/project/target/snapshots
  - /opt/project/target/test-classes
  - /opt/project/target/war/work
  - /opt/project/target/webappDirectory
  entrypoint: /bin/true
mvn: &mvn
  image: maven:3.2-jdk-8
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: mvn
  volumes_from:
    - lc_shared_mvndata
    - mvnscratch
test:
  <<: *mvn
  entrypoint: [mvn, test]
package:
  <<: *mvn
  command: [package, "-DskipTests=true"]
publish:
  <<: *mvn
  entrypoint: /bin/true
clean:
  <<: *mvn
  entrypoint: [mvn, clean, "-Dmaven.clean.failOnError=false"]
```

The important parts in the above template are the `mvnscratch` and `mvn`
services. The `mvn` service is the primary service that is building the code.
The `mvnscratch` service is the data container that declares, as volumes, all
repo paths that hold transient build resources (i.e., files that are re-built
every time the build runs).

Notice that `mvn` uses the [docker-compose
volumes_from](https://docs.docker.com/compose/compose-file/#volumes-from)
directive to include the volumes defined in `mvnscratch`. It also uses the
[docker-compose
volume](https://docs.docker.com/compose/compose-file/#volumes-volume-driver)
directive to mount the root of the repo (on the `docker-daemon` VM) to `/opt/project`
inside the container. When `docker-compose` processes these directives, it sets
up the volumes in the following way:

1. Process `volume` directive first and mount repo's root (on the `docker-daemon` VM) to
`/opt/project` inside the container.
1. Now process `volumes_from` and overlay the volumes from `mvnscratch` onto
`mvn`. This means that within the `mvn` container, for each directory listed in
`mvnscratch` the data inside that directory will be stored inside the
`mvnscratch` container, and NOT on the `docker-daemon` VM.

Remember that all files that end up on the `docker-daemon` VM, also get mirrored to the
host OS, which is where the excessive I/O comes from. By bypassing the storage
on the `docker-daemon` VM, via the data container (e.g., `mvnscratch`), the mirroring
never occurs.
