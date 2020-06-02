#!/bin/bash

# get the location of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "running go fmt..."
go fmt "${SCRIPT_DIR}"/...

echo "running go lint..."
golint "${SCRIPT_DIR}"/...

echo "running go vet..."
go vet "${SCRIPT_DIR}"/...

echo "running go test..."
go test -v "${SCRIPT_DIR}"/...

echo "displaying test coverage..."
go tool cover -func=c.out

