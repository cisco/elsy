# elsy Templates

The elsy lifecycle manifests itself in subtly different ways depending on the
underlying build tool. elsy ships with a small set of pre-baked templates (e.g.,
mvn, sbt) that define a sensible default lifecycle for the build tool
encapsulated by the template.

A template provides a base set of docker-compose services that implement the elsy
lifecycle for the specific build tool. When you run any elsy command, elsy will
extend the template using the repo's `docker-compose.yml` before executing the
command.

For example, you can see how elsy is extending the template with the repo's compose
file by running elsy in debug mode:

```
## this command was run from a repo using the mvn template
## redacted everything but a single docker-compose call
## lc_docker_compose_template849512109 is the elsy template dumped to a tmp file for use in the command
$  lc --debug bootstrap
DEBU[0000] running command /usr/local/bin/docker-compose with args [/usr/local/bin/docker-compose -f /Users/<user>/dev/code/java-service/lc_docker_compose_template849512109 -f docker-compose.yml kill]
```

It is possible to override template-provided services by following the
[docker-compose extends](https://docs.docker.com/compose/extends/) semantics.

To see the "effective compose" file that elsy ends up using you can run `lc dc config`.

## Configuring a Template

To configure a repo to use a specific template, simply update the
[lc.yml](./configuringlcrepo.md) file to use the `template: <template-name>`
config.

## Viewing a Template

To view the a raw template you can run `lc system view-template <template-name>`

## Supported Templates

The following subsections list the templates that elsy provides. The `mvn` and
`sbt` templates are the most widely used, all other templates have limited known
use (at this time), so may require some improvement.

### lein

To use the lein template, ensure your `lc.yml` has the line:

```
template: lein
```

This template enables the `lc lein` subcommand, you can run any lein command in
your repo by running `lc lein -- <leincmd>`. This template also adds a data
container called `lc_shared_mvndata`. This data container holds the `~/.m2`
cache for the host, meaning that all elsy lein and elsy mvn repos running on a
single host will share the same `~/.m2` cache.

To override the lein image that comes pre-baked into the template, you must
have, at a minimum, the following overrides in the repo's `docker-compose.yml`:

```
lein:
  image: <custom-image>
test:
  image: <custom-image>
package:
  image: <custom-image>
publish:
  image: <custom-image>
clean:
  image: <custom-image>
```

### make

To use the make template, ensure your `lc.yml` has the line:

```
template: make
```

This template enables the `lc make` subcommand, you can run any make command in
your repo by running `lc make -- <makecmd>`.

To override the make image that comes pre-baked into the template, you must
have, at a minimum, the following overrides in the repo's `docker-compose.yml`:

```
make:
  image: <custom-image>
test:
  image: <custom-image>
clean:
  image: <custom-image>

```

### mvn

To use the mvn template, ensure your `lc.yml` has the line:

```
template: mvn
```

This template enables the `lc mvn` subcommand, you can run any mvn command in
your repo by running `lc mvn -- <mvncmd>`. This template also adds a data
container called `lc_shared_mvndata`. This data container holds the `~/.m2`
cache for the host, meaning that all elsy mvn and elsy lein repos running on a
single host will share the same `~/.m2` cache.

To override the mvn image that comes pre-baked into the template, you must
have, at a minimum, the following overrides in the repo's `docker-compose.yml`:

```
mvn:
  image: <custom-image>
test:
  image: <custom-image>
package:
  image: <custom-image>
publish:
  image: <custom-image>
clean:
  image: <custom-image>
```

### sbt

**Currently there is no official sbt Docker image, it is STRONGLY recommended that
you override the default image baked into the elsy sbt template.**

To use the mvn template, ensure your `lc.yml` has the line:

```
template: sbt
```

This template enables the `lc sbt` subcommand, you can run any sbt command in
your repo by running `lc sbt -- <sbtcmd>`. This template also adds a data
container called `lc_shared_sbtdata`. This data container holds the `~/.ivy2`
cache for the host, meaning that all elsy sbt repos running on a single host will
share the same `~/.ivy2` cache.

To override the sbt image that comes pre-baked into the template, you must
have, at a minimum, the following overrides in the repo's `docker-compose.yml`:

```
sbt:
  image: <custom-image>
test:
  image: <custom-image>
package:
  image: <custom-image>
publish:
  image: <custom-image>
clean:
  image: <custom-image>
```
