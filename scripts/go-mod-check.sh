#!/usr/bin/env bash

set -uo pipefail

for f in $(find . -name go.mod)
do (cd $(dirname $f);
    go mod tidy

    if [[ $(git diff --stat "**/go.sum" | wc -l) != 0 ]]; then
        echo "'go.mod' file does not match the module source code. Please run 'go mod tidy' in '$(pwd)'"
        exit 1
    fi
)
done