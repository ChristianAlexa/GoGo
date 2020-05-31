#!/bin/bash

# get the location of this script.
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# recursively check for tests in project and run them.
go test -v ${SCRIPT_DIR}/...
