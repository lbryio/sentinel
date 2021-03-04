#!/bin/bash

 set -euo pipefail

 DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
 cd "$DIR"
 cd ".."
 DIR="$PWD"


echo "== Installing dependencies =="
go mod download


echo "== Checking dependencies =="
go mod verify
set -e


echo "== Compiling =="
export IMPORTPATH="github.com/lbryio/chain-watcher"
mkdir -p "$DIR/bin"
go generate -v
export VERSIONSHORT="${TRAVIS_COMMIT:-"$(git describe --tags --always --dirty)"}"
export VERSIONLONG="${TRAVIS_COMMIT:-"$(git describe --tags --always --dirty --long)"}"
export COMMITMSG="$(echo ${TRAVIS_COMMIT_MESSAGE:-"$(git show -s --format=%s)"} | tr -d '"' | head -n 1)"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o "./bin/watcher" -asmflags -trimpath="$DIR" -ldflags "-X ${IMPORTPATH}/meta.version=${VERSIONSHORT} -X \"${IMPORTPATH}/meta.commitMsg=${COMMITMSG}\""

echo "$(git describe --tags --always --dirty)" > ./bin/watcher.txt
chmod +x ./bin/watcher
exit 0