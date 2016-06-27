#!/bin/bash

set -e

IMAGE_NAME="analyse"

# Get the parent directory of where this script is.
echo ${BASE_SOURCE}
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do
    SOURCE="$(readlink "$SOURCE")";
done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
echo $DIR

echo "==> Getting dependencies..."
go get ./...

echo "==> Removing old directory..."
rm -f bin/*
rm -f pkg/*
mkdir -p bin/

echo "==> Building..."
go build -o bin/${IMAGE_NAME} main.go
