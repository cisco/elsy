# sbt-scala-helloworld

A simple REST hello-world service implemented in Akka-http to illustrate
docker container deployment of sbt based project with Scala.

## CI

Simply run `lc ci`

## Local Development
```
## bootstrap repo
$ lc bootstrap

## unit-test your code
$ lc test

## test your code
$ lc blackbox-test

## start packaged server
$ lc server start -prod

## get logs
$ lc dc -- logs -f

## cleanup
$ lc teardown
```
