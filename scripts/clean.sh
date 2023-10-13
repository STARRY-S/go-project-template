#!/usr/bin/env bash

set -euo pipefail

cd $(dirname $0)/../
WORKINGDIR=$(pwd)

files=(
    # TODO: edit here
    ".dapper*"
    "go-project-template"
    "bin/"
    "dist/"
)

for f in ${files[@]}; do
    if [[ -e "$f" ]]; then
        echo "Delete: $f"
        rm -rf $WORKINGDIR/$f
    fi
done

exit 0
