#!/bin/bash

# build binary for the specified os

# USAGE
# ./build.sh <os>

# set cwd to wherever this script is located
cd "$(dirname "$0")"

OS=$1
VERSION=`cat version`

if test -z $OS; then
  echo "missing arg"
  exit 1
fi

if test $OS != "darwin" && test $OS != "linux"; then
  echo "unknown os"
  exit 1
fi

echo "running tests"
go test
TEST_RESULT=`echo $?`
if test $TEST_RESULT -ne 0; then
  echo "tests failed. Exiting"
  exit $TEST_RESULT
fi

echo "building dockerw $VERSION for $OS"
GOOS=$OS GOARCH=amd64 go build -ldflags "-X dockerw.Version=$VERSION" -o bin/$OS/docker-watch ./cmd/main
echo done
