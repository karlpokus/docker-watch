#!/bin/bash

# copy binary for the specified os to dest

# USAGE
# ./copy.sh <os>

OS=$1
DEST=~/.docker/cli-plugins/

if test -z $OS; then
  echo "missing arg"
  exit 1
fi

if test $OS != "darwin" && test $OS != "linux"; then
  echo "unknown os"
  exit 1
fi

mkdir -p $DEST
cp -v bin/$OS/docker-watch $DEST
