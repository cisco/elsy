#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $DIR/_helpers

: ${JVM_OPTS:="-Xms1g -Xmx2g"}

set -e

wait_for_services
wait_for_db_migration
exec java $JVM_OPTS -jar java-note-service.jar "$@"
