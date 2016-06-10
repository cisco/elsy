## _Unreleased_

## v0.16.0

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
