# simple-docker-image

A dead simple docker image that serves a single web page.

The key things this example illustrates are:

- How lc can manage internal docker images of third-party tools.
- Using lc to do the bare minimum: build a docker image.

Note `lc publish` will always fail for this example since it is configured
to publish to a non-existent docker registry.

## CI

Simply run `lc ci`

## Local Development
```
## bootstrap repo
$ lc bootstrap

## test your code
lc blackbox-test

## start packaged server
$ lc server start -prod

## get logs
$ lc dc -- logs -f

## cleanup
$ lc teardown
```
