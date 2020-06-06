#!/bin/bash

# get the location of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "running go fmt..."
go fmt "${SCRIPT_DIR}"/...

echo "running go lint..."
golint "${SCRIPT_DIR}"/...

echo "running go vet..."
go vet "${SCRIPT_DIR}"/...

echo "running go test and displaying coverage.."
go test "${SCRIPT_DIR}"/... -v -coverpkg=./...
