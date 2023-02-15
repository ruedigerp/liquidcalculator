#!/bin/bash

set -eux

export GOPATH="$(pwd)/.gobuild"
SRCDIR="${GOPATH}/src/github.com/9IT-Full-Service/podloxx-hb/"

[ -d ${GOPATH} ] && rm -rf ${GOPATH}
mkdir -p ${GOPATH}/{src,pkg,bin}
mkdir -p ${SRCDIR}
# cp main.go ${SRCDIR}
cp -av go.mod main.go  ${SRCDIR}
(
    echo ${GOPATH}
    cd ${SRCDIR}
    go mod tidy
    go get .
    go install .
)
