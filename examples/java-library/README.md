# java-note-service

A java library that provides a single class.

The key things this example illustrates are:

- Using `lc` to build a java library (i.e., not a docker image)
- How to use a custom `lc publish` service to publish non-docker artifacts.

Note `lc publish` will always fail for this example since it is configured
to publish to a non-existent maven registry.

## CI

Simply run `lc ci`

## Local Development
```
## bootstrap repo
$ lc bootstrap

## test your code
$ lc test

## run specific mvn commands:
$ lc mvn <mvncmd>

## package service into a Docker image
$ lc package
```
