#!/usr/bin/env bash

set -e
rm -rf ./dist-production || true
echo "## Packaging production build"
ember build --environment="production" --output-path=dist-production
