#!/bin/bash
CODE_DIR=$1
APP_NMAE=$2

[ -n "$CODE_DIR" ] || {
    echo "CODE_DIR is null"
    exit 1
}
[ -n "$APP_NMAE" ] || {
    echo "APP_NMAE is null"
    exit 1
}

cd $CODE_DIR
if [ -e go.mod ]; then
    sed -i 's/go 1.15/go 1.17/g' go.mod
    go get github.com/prometheus/procfs
    go mod tidy || go mod tidy -compat=1.17
    go mod vendor
fi