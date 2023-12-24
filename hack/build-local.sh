#!/usr/bin/env bash
set -e

export RELEASE_VERSION=v0.0.1-dev

export GO111MODULE=on
export GOFLAGS=-mod=vendor

PROVIDER_ROOT=$(git rev-parse --show-toplevel)

if [[ "$(pwd)" != "${PROVIDER_ROOT}" ]]; then
  echo "you are not in the root of the repo" 1>&2
  echo "please cd to ${PROVIDER_ROOT} before running this script" 1>&2
  exit 1
fi

$PROVIDER_ROOT/hack/build.sh

# update provider.yaml for local use
go run -mod vendor "${PROVIDER_ROOT}/hack/local/main.go" ${RELEASE_VERSION} "$PROVIDER_ROOT/release" > "${PROVIDER_ROOT}/release/provider-local.yaml"

