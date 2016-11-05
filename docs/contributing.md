# Contributing to elsy

This section describes how to build the elsy code base. If you would like to
submit a patch to elsy, please submit a PR to the `master` branch.

## Local Development

Use elsy to develop elsy! Currently docker 1.11 is required to run
blackbox-tests.

This repo exposes all of the core elsy tasks for ongoing development:

```
## bootstrap repo
$ lc bootstrap

## test your code
$ lc test && lc blackbox-test

## package a new binary, will show up in ./target/
$ lc package
```

### IDE Integration

You must have go version 1.6 or higher installed locally for IDE integration to
fully work.

Follow these instructions to enable IDE integration during development. IDE
integration is purely for speeding local work, developers should still run `lc
test && lc blackbox-test` to validate code before pushing.

[Atom](https://atom.io/) is the recommended editor for this project and it
is also recommended that you use the [go-plus](https://atom.io/packages/go-plus)
package for live `golinting` and `govetting`

```
$ git clone git@github.com:cisco/elsy.git
$ cd project-lifecycle
$ lc bootstrap
$ atom .
```

Now, open atom by running `atom .`.
