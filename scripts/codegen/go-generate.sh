#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

checkInstalled 'swagger'

GEN_DIR="${REPO_ROOT}/internal/gen/strava-api-go"

rm -rf "${GEN_DIR}"
mkdir -p "${GEN_DIR}"

go generate -x ./...

echo "${SCRIPT_NAME} done."
