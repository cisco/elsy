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
