# java-note-service

A simple [Dropwizard](http://www.dropwizard.io/) service that stores arbitrary
notes in mysql.

The key things this example illustrates are:

- Using `lc` to build a mvn-based microservice.
- Using the `lc` mvn template with a custom mvn version
- Ensuring specific order of operations when running blackbox-tests; this example
is configured to ensure that the server and blackbox-tests will not start until
the mysql database is ready.

Note `lc publish` will always fail for this example since it is configured
to publish to a non-existent docker registry.

## CI

Simply run `lc ci`

## Local Development
```
## bootstrap repo
$ lc bootstrap

## test your code
$ lc test && lc blackbox-test

## run specific mvn commands:
$ lc mvn <mvncmd>

## package service into a Docker image
$ lc package

## start packaged server
$ lc server start -prod

## get logs
$ lc dc -- logs -f

## cleanup
$ lc teardown
```
