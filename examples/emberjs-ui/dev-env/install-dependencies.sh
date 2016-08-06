#!/usr/bin/env bash

set -e
npm install
bower --allow-root install
bower --allow-root update
