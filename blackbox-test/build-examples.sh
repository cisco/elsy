#!/usr/bin/env bash

# This script is used for building all the example programs, when run from TravisCI.

BBT_LC=$(pwd)/target/lc-blackbox
DIR=$(pwd)/examples

function build() {
  dir=$1
  has_tests=$2
  has_bbtests=$3

  echo "*** Building example: ${dir} ***"
  cd $dir

  # Ensure nothing is left over from previous runs
  $BBT_LC teardown

  if $has_tests; then
    $BBT_LC bootstrap
  fi

  if $has_bbtests; then
    $BBT_LC blackbox-test
  fi
}

set -e

build "${DIR}/c-code" false true

# https://github.com/cisco/elsy/issues/81
# build "${DIR}/emberjs-ui" true true

build "${DIR}/java-library" false false
build "${DIR}/java-note-service" false true
build "${DIR}/sbt-scala" false true
build "${DIR}/simple-docker-image" false true
