# Contributing to LC

This section describes how to build the `lc` code base. If you would like to
submit a patch to lc, please submit a PR to the `master` branch.

## Local Development

Use `lc` to develop `lc`!

This repo exposes all of the core `lc` tasks for ongoing development:

```
## bootstrap repo
$ lc bootstrap

## test your code
$ lc test && lc blackbox-test

## package a new binary, will show up in ./target/
$ lc package
```

### Setting up CI
In order to have Jenkins build and publish `project-lifecycle`, you must
configure the Jenkins server with credentials to upload to
artifactory1.eng.lancope.local. First, create a local file called `netrc` with
the contents:

```
machine artifactory1.eng.lancope.local login <username> password <password>
```

Visit your Jenkins Credentials management page, add a new credential of type
*Secret File* then upload the `netrc` file you just created.

In the `project-lifecycle` job configuration page, add _Use Secret text for
file_ to the _Build Environment_ section. Select the credential you just added
and specify an env variable of `LC_REPO_NETRC`.

### IDE Integration

You must have go version 1.6 or higher installed locally for IDE integration to
fully work.

Follow these instructions to enable IDE integration during development. IDE
integration is purely for speeding local work, developers should still run `lc
test && lc blackbox-test` to validate code before pushing.

[Atom](https://atom.io/) is the recommended editor for `golang` projects and it
is also recommended that you use the [go-plus](https://atom.io/packages/go-plus)
package for live `golinting` and `govetting`

```
$ git clone git@github.com:cisco/elsy.git
$ cd project-lifecycle
$ lc bootstrap
```

Now, open atom by running `atom .`.
