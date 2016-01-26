# Improving `lc` performance (experimental)

If you are using `lc` with `lds` there is a lot of I/O that happens between your repo folders
on your host OS and the corresponding folder on the `lds` VM; this I/O may sometimes slow down local builds.

To reduce this I/O it is possible to set the following in your environment:

```
export LC_ENABLE_SCRATCH_VOLUMES=true
```

Exporting the above environment variable will tell `lc` to keep all temporary build resources
(i.e., the data most often found in `./target` subdirectories for jvm builds) inside the `lds`
VM; this is done using data containers. To see how this effects your favorite template try the following:

```
# print the template without scratch volumes enabled:
$ unset LC_ENABLE_SCRATCH_VOLUMES
$ lc system view-template sbt

# print the template with scratch volumes
$ lc --enable-scratch-volumes system view-template sbt
```

Note in the last line in the above example, using the `--enable-scratch-volumes` flag is equivalent
to setting the `LC_ENABLE_SCRATCH_VOLUMES` environment variable.

The following subsections provide more details on using this setting:

## Inspecting build data inside the data container

When using `LC_ENABLE_SCRATCH_VOLUMES`, it is still possible to view all build data. You simply
need to inspect the data container.

To inspect the data container, run the following to spawn a container that mounts the volumes from the data container:

```
## For repos using the sbt template:
$ lc dc -- run --entrypoint=bash sbt -c bash

## For repos using the maven template:
$ lc dc -- run --entrypoint=bash mvn -c bash
```

After running the above you will have a shell in with a working directory of `/opt/project`, which
will contain all of your repo data, including the data inside the data container.

## Technical Details

This section explains the technical mechanisms underpinning `LC_ENABLE_SCRATCH_VOLUMES`.
Feel free to ignore this.

Lets use the `sbt` template, with `--enable-scratch-volumes` to explain how things are working:

```
$ lc --enable-scratch-volumes system view-template sbt
sbtscratch:
  image: busybox
  command: /bin/true
  volumes:
    - /opt/project/target/resolution-cache
    - /opt/project/target/scala-2.10/classes
    - /opt/project/target/scala-2.10/test-classes
    - /opt/project/target/scala-2.11/classes
    - /opt/project/target/scala-2.11/test-classes
    - /opt/project/target/streams
    - /opt/project/project/project
    - /opt/project/project/target
sbt: &sbt
  image: arch-docker.eng.lancope.local:5000/sbt
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: sbt
  volumes_from:
    - lc_shared_sbtdata
    - sbtscratch
test:
  <<: *sbt
  entrypoint: [sbt, test]
package:
  <<: *sbt
  command: [assembly]
```

The important parts in the above template are the `sbtscratch` and `sbt` services.
The `sbt` service is the primary service that is building the code. The `sbtscratch`
service is the data container that declares, as volumes, all repo paths that hold transient
build resources (i.e., files that are re-built every time the build runs).

Notice that `sbt` uses the [docker-compose volumes_from](https://docs.docker.com/compose/compose-file/#volumes-from)
directive to include the volumes defined in `sbtscratch`. It also uses the
[docker-compose volume](https://docs.docker.com/compose/compose-file/#volumes-volume-driver)
directive to mount the root of the repo (on the `lds` VM) to `/opt/project` inside the container.
When `docker-compose` processes these directives, it sets up the volumes in the following way:

1. Process `volume` directive first and mount repo's root (on the `lds` VM) to `/opt/project` inside the container.
2. Now process `volumes_from` and overlay the volumes from `sbtscratch` onto `sbt`. This means
that within the `sbt` container, for each directory listed in `sbtscratch` the data
inside that directory will be stored inside the `sbtscratch` container, and NOT on the `lds` VM.

Remember that all files that end up on the `lds` VM, also get mirrored to the host OS, which is where the excessive I/O comes from. By bypassing the storage on the `lds` VM, via the data container (e.g., `sbtscratch`), the mirroring never occurs.
