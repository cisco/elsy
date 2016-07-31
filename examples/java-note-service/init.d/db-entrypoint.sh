#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $DIR/_helpers

set -e

wait_for_services
exec java -jar java-note-service.jar "$@"
