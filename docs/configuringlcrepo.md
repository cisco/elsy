# Using elsy in a Project

The simplest way to incorporate elsy into your project (either new or existing)
is to use the `lc init` command:

For example, to create a new elsy repo that publishes a docker image, and uses the
built-in mvn template, do:

```
$ lc init --docker-image-name=java-service --docker-registry=internal-registry.companyfoo.com --template=mvn java-service
$ cd java-service
$ lc bootstrap
```

You will now see an `lc.yml`, `Dockerfile`, and `docker-compose.yml` file in
your repo. The following sections describe how elsy uses these files.

## lc.yml

elsy uses the `lc.yml` file at the root of the repo to figure out how it should
work inside the repo. `lc.yml` supports the following configuration options:

* `project_name`: name of your `docker-compose` project which is used as
__COMPOSE_PROJECT_NAME__. This name should not contain any spaces, underscores
or dashes.
* `docker_compose`: basename or fully qualified path to the docker-compose binary.
This will override the `LC_DOCKER_COMPOSE` environment variable to enable you
to use a custom docker-compose version for your repo.
* `template`: [elsy template](./templates.md) to use for the project.
* `template_image`: specify a docker image to [override the template image](./templates.md#overriding-the-image-specified-by-a-template).
* `docker_image_name`: name of docker image to build.
* `docker_registry`: address of docker registry to publish to.
* `docker_registries`: takes a yaml sequence containing multiple registries to publish to. Use either
`docker_registry` or `docker_registries`, not both.
* `build_logs_dir`: If populated, elsy will dump ALL docker-compose service logs into this
directory, directory must be relative to the repo root.
* `local_images`: takes a yaml sequence containing images to not pull during
bootstrap. This allows repo owners to provide images to the build using some
external process. The image declared in `docker_image_name` is automatically
included in this sequence.

Some configuration options may also be specified as command line arguments.
If a command line argument is present, it will take precedence and override any
value in the configuration file.

## docker-compose.yml

elsy heavily relies on Docker Compose to figure out how it should build the repo.
If you are using a built-in elsy template and do not need to make any modifications,
then your `docker-compose.yml` can be as simple as:

```
## default compose file created by lc init
noop:
   image: alpine
```

You can also [override portions of the template](./templates.md) if needed.

If you are not using a built-in template, then you will need to define a
docker-compose service for the portions of the elsy lifecycle you want to run in
your repo (e.g., test, package, blackbox-test, publish). See the [elsy
examples](../examples/README.md) to get an idea of different ways of doing this.

## Dockerfile (optional)

elsy will use the `Dockerfile` at the root of the repo when packaging and publishing
the Docker image for the repo. The default image created when `lc init` is run
with the `--docker-image-name` flag is:

```
FROM scratch
```

You will need to override this to implement the image you need for your repo.
