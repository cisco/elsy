# emberjs-ui

A simple [Ember](http://emberjs.com/) UI that renders notes to the screen.

The key things this example illustrates are:

- Using `lc` to build a javascript project without installing any javascript tooling
- Using the `lc devserver` pattern to enable a hot-reloading dev workflow for
javascript development
- Using a custom `package` phase to build the ember code before packaging the
docker image.
- Running blackbox-tests to artifact the final product works as expected.

Note:`lc publish` will always fail for this example since it is configured
to publish to a non-existent docker registry.

TODO: Currently live reloading only works if developing natively on linux, need
to update this example to hook up https://github.com/leighmcculloch/docker-unison and
sending of mac os x file events to docker-daemon VM.

## CI

Simply run `lc ci`

## Local Development
```
## bootstrap repo
$ lc bootstrap

## test your code
$ lc test && lc blackbox-test

## run ember command
$ lc dc -- run ember <embercmd>

## run npm command
$ lc dc -- run npm <npmcmd>

## run bower command
$ lc dc -- run bower <bowercmd>

## start the gulp development that supports hot-reloading of code
$ lc server start
# restart server
$ lc server restart

## start server as it will appear in prod (i.e., served from nginx)
$ lc package && lc server start -prod

## get logs
$ lc dc -- logs -f

## cleanup
$ lc teardown
```
