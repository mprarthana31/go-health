#!/usr/bin/bash

set -exo pipefail
BUILD_DATE=$(date -u)

PACKAGE_NAME="main"

LDFLAGS=
LDFLAGS="${LDFLAGS} -X '${PACKAGE_NAME}.BuildTimestamp=${BUILD_DATE}'"

go build -ldflags "${LDFLAGS}" -o ./health