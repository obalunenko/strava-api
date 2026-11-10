#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"

API_DOC_URL="${API_DOC_URL:-https://developers.strava.com/swagger/swagger.json}"

# Pinned: a different flatten output format would conflict across the whole file.
SWAGGER="github.com/go-swagger/go-swagger/cmd/swagger@v0.36.7"

SPEC="docs/swagger.json"
UPSTREAM="docs/swagger.upstream.json"

cd "${REPO_ROOT}"

echo "${SCRIPT_NAME} is running... "

if [[ -n "$(git status --porcelain -- docs)" ]]; then
  echo "${SCRIPT_NAME}: commit or stash the changes in docs/ first" >&2
  exit 1
fi

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT

go run "${SWAGGER}" flatten --format=json --with-flatten=minimal --output="${TMP_DIR}/upstream.json" "${API_DOC_URL}"

if cmp -s "${TMP_DIR}/upstream.json" "${UPSTREAM}"; then
  echo "${SCRIPT_NAME}: the Strava spec has not changed"
  exit 0
fi

# Only what Strava changed between its two versions is applied, so the local
# changes in ${SPEC} survive. The exit code is the number of conflicts.
conflicts=0
git merge-file -L "${SPEC}" -L "previous upstream" -L "new upstream" \
  "${SPEC}" "${UPSTREAM}" "${TMP_DIR}/upstream.json" || conflicts=$?

if [[ "${conflicts}" -gt 127 ]]; then
  echo "${SCRIPT_NAME}: merge failed" >&2
  exit 1
fi

# The snapshot moves on even with conflicts: it is the base of the next update.
cp "${TMP_DIR}/upstream.json" "${UPSTREAM}"

if [[ "${conflicts}" -gt 0 ]]; then
  echo "${SCRIPT_NAME}: ${conflicts} conflict(s) in ${SPEC}, resolve the <<<<<<< markers and run 'make codegen'" >&2
  exit 1
fi

# Fails on broken JSON and on references to definitions that Strava removed.
if ! go run "${SWAGGER}" expand --format=json --output="${TMP_DIR}/check.json" "${SPEC}"; then
  echo "${SCRIPT_NAME}: ${SPEC} is broken after the merge, fix it and run 'make codegen'" >&2
  exit 1
fi

echo "${SCRIPT_NAME} done. Review 'git diff docs/' and run 'make codegen'."
