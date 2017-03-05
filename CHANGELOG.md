## Unreleased

- Fixed [issue #50](https://github.com/cisco/elsy/issues/50).
- Fixed [issue #51](https://github.com/cisco/elsy/issues/51).
- Fixed [issue #56](https://github.com/cisco/elsy/issues/56).
- Fixed [issue #59](https://github.com/cisco/elsy/issues/59).

## v1.7.0

- Fixed [https://github.com/cisco/elsy/issues/46](https://github.com/cisco/elsy/issues/46).
- Added support for dumping build logs
(See [#43](https://github.com/cisco/elsy/issues/43)).

## v1.6.1

- support for v2 expanded build syntax
- When releasing, it now checks for an existing tag or branch that is the same as the `--version` argument
and gives a better explanation of what happened.
- Fixed a bug that happened in packaging if the verison of Docker had something like `-rc2` at the end.
- Fixed issue #4, which would cause a panic in rare cases during packaging.

## v1.6.0

- Added the ability to override the image that a template uses by setting
`template_image` to the desired image in the project's `lc.yml` file.

## v1.5.1

- Fix `publish` for docker 1.12 (see
[https://github.com/cisco/elsy/issues/25](https://github.com/cisco/elsy/issues/25)).
Note this breaks support for docker 1.9.x and below.

## v1.5.0

- If there's no `publish` service _and_ no Dockerfile, then instead of an error
when using `lc ci`, it will simply print out an info log message that `publish`
isn't being called.

- sbt template now uses the `paulcichonski/sbt` Docker image. This is not an
official Docker image (no official sbt image exists) so users should still
override this image with an image they have vetted.

## v1.4.0

- Added `lc run` to keep from having to type `lc dc run`. Any aguments that `lc dc run` takes,
`lc run` can also take.
- Added `lc bbtest` as an alias for `lc blackbox-test`.

## v1.3.0

- Add initial support for Compose v2 syntax when using templates [\#11](https://github.com/cisco/elsy/pull/11)
- Update `lc package` to apply the label:
`com.elsy.metadata.git-commit=<git-commit>` to the Docker image if the
`GIT_COMMIT` env var is populated. Note that this only works for Docker 1.11.1 and
higher. [\#13](https://github.com/cisco/elsy/pull/13)

## v1.2.1

- fixing release issue with 1.2.0 (`./VERSION`) did not get updated.

## v1.2.0 (NEVER RELEASED)

- remove docker-compose v2 volumes on teardown

## v1.1.1

- reduce logging noise introduced by d6b9310

## v1.1.0

- rename `verify-lds` to `verify-install`
- remove `lc system upgrade` command. This hasn't done anything since 6333e0d.
- initial support for compose v2 file formats (lc won't fail if you use a v2
file format without a built-in lc template).

## v1.0.0

- Fix lc release error message to stop escaping regex.
- (breaking) Remove `lc smoketest`
- Hide benign error when docker-compose service uses the primary docker image
artifact (now works with docker 1.10)

## v0.16.2

- Fix bug where `lc package` was not always removing all previous containers
created from previous versions of the docker image.

## v0.16.1

- no-op release to fix issue around v0.16.0 release.

## v0.16.0

(NEVER RELEASED)
- Stop 'lds-verify' from `bootstrap`, `package` and `test`

## v0.15.1

- Fixed bug in `lc release` where it would not allow multi-digit patch numbers.

## v0.15.0

- Added `-Dmaven.clean.failOnError=false` to the `mvn` template's default `clean` service
so that running `lc clean` when scratch volumes are enabled won't cause the build to fail.
- Deprecated `system upgrade`, since users should upgrade `lc` when they upgrade `lds`
by running `lds upgrade`.
- Update build to use `govendor` to lock in dependencies
- Updated `lc package` to cleanup any containers created from previous versions of the
docker image.

## 0.14.0

- After #DumpsterFireApril2016, we felt an offline mode would be useful. So, if the VM
infrastructure ever erupts in flames again, _and_ you have already pulled down the images
that you need for building, then adding `--offline` will make the build work. If you had
not already pulled the requisite images, you could build them yourself from their sources,
and tag them so that they would be available in your local Docker image cache, and then
use the `--offline` switch.

## 0.13.0

- allow custom package service to generate the Dockerfile

## 0.12.0

- support passing computed docker tag name to the custom publish service

## 0.11.0

- Added a `clean` command, which will remove old build artifacts.

## 0.10.0

- Added support for publishing to multiple docker registries [#102](https://stash0.eng.lancope.local/projects/DEV-INFRASTRUCTURE/repos/project-lifecycle/pull-requests/102/overview).
- Added a `lein` command, which facilitates building Clojure projects.

## 0.9.0

- Added a `make` command, which is intended for building C/C++ projects, which have a Makefile.
- Added `list-templates` to the `system` command, to list all the built-in templates.

## 0.8.0

- `lc package` will now run the `test` service, if present, before packaging. If you do not want those tests to be run,
    run `lc package --skip-tests`.

## 0.7.0

- Added `lc init` command for initializing repos to use lc [#95](https://stash0.eng.lancope.local/projects/DEV-INFRASTRUCTURE/repos/project-lifecycle/pull-requests/95/overview).

## 0.6.1

- Added usage information to every command.
- Added best practices documentation.

## 0.6.0

- Added --skip-docker option to `package`
- Hid benign error when docker-compose service uses the primary docker image artifact

## 0.5.0

- Renamed `smoketest` to `blackbox-test`. Tests should now go in `./blackbox-test`. Existing smoketests are still supported, but the `smoketest` command will be removed at some point.

## 0.4.2

- Source formatted with `gofmt`.
- Now treats non-release tags the same as non-release branches, and should no longer blow up.

## v0.4.1

- Correctly parses docker-compose version strings that include build info.

## v0.4.0

- `lc system upgrade` now requires a `--version` flag to tell `lc` the target
upgrade version
- reworked `lc system upgrade` to reduce possible errors during the install

## v0.3.1

- Fix publishing task

## v0.3.0

- Include version in binary that gets published.
- Include build hash in `lc --version`
